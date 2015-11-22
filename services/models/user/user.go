package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"

	"github.com/jmcvetta/neoism"
	"github.com/nathandao/vantaa/core/vantaadb"
	"github.com/nathandao/vantaa/settings"
)

type User struct {
	Id             int    `json:"id"`
	Name           string `json:"name" form:"name"`
	Email          string `json:"email" form:"email"`
	PasswordDigest []byte `json:"-"`
	Password       string `json:"password"`
}

type UserAdapter struct {
	Id             int    `json:"id(u)"`
	Name           string `json:"u.name"`
	Email          string `json:"u.email"`
	PasswordDigest []byte `json:"u.password_digest"`
}

// Transform creates a User struct from the UserAdapter
func (ua *UserAdapter) Transform() *User {
	u := User{
		ua.Id,
		ua.Name,
		ua.Email,
		ua.PasswordDigest,
		"",
	}
	return &u
}

// New initiate a new User struct
func New() *User {
	return &User{}
}

// Save or u.Save calls on CreateUser to write a newu User node in the database.
func (u *User) Save() (*User, error) {
	newu, err := CreateUser(u)
	if err != nil {
		return nil, err
	}
	return newu, nil
}

// Delete removes the use
// func (u *User) Delete() error {
// 	if err := DeleteUser(u); err != nil {
// 		return err
// 	}
// 	return nil
// }

// CreateUser validates the User's information. Upon successful validations,
// it creates a new User node in the database,
func CreateUser(u *User) (*User, error) {
	// user sanitization
	u.sanitize()
	if _, err := ValidateUser(u); err != nil {
		return nil, err
	}
	passwordDigest, err := bcrypt.GenerateFromPassword([]byte(u.Password), settings.Get().HashCost)
	if err != nil {
		return nil, err
	}
	// remove passsword string right after
	u.Password = ""
	u.PasswordDigest = passwordDigest
	res := []UserAdapter{}
	db := vantaadb.Connect()
	cq := neoism.CypherQuery{
		Statement: `CREATE (u:User {
                  name:{name},
                  email:{email},
                  password_digest:{password_digest}
                })
                RETURN id(u), u.name, u.email`,
		Parameters: neoism.Props{"name": u.Name, "email": u.Email, "password_digest": u.PasswordDigest},
		Result:     &res,
	}
	if err := db.Cypher(&cq); err != nil {
		return nil, err
	}
	ua := &res[0]
	return ua.Transform(), nil
}

// DeleteUser removes a User node from the database.
// error is return even when user is not found
// func DeleteUser(u *User) error {
// 	db := vantaadb.Connect()
// 	cq := neoism.CypherQuery{
// 		Statement: `MATCH (u:User)
//                 WHERE ` + vantaadb.PropString("u", props) + `DELETE u, s, r`,
// 		Parameters: props,
// 	}
// 	if err := db.Cypher(&cq); err != nil {
// 		return err
// 	}
// 	return nil
// }

// FindUser finds a single User by calling on FindUsers and returns the
// first element of the []*User slice.
func FindUser(props neoism.Props) (*User, error) {
	u, err := FindUsers(props)
	if err != nil {
		return nil, err
	}
	return u[0], nil
}

// FindUsers contructs and executes a cypher query with the provided parametrs
// and return all User nodes that satisfy the props conditions.
func FindUsers(props map[string]interface{}) ([]*User, error) {
	db := vantaadb.Connect()
	res := []UserAdapter{}
	// generate condition string to be used in the cypher statement
	condstr := vantaadb.PropString("u", props)
	cq := neoism.CypherQuery{
		Statement: `MATCH (u:User)
                WHERE ` + condstr + `
                RETURN
                  id(u),
                  u.name,
                  u.email,
                  u.password_digest`,
		Parameters: props,
		Result:     &res,
	}
	err := db.Cypher(&cq)
	if err != nil {
		return []*User{}, err
	}
	users := []*User{}
	for _, ua := range res {
		users = append(users, ua.Transform())
	}
	if len(users) < 1 {
		return nil, errors.New("not found")
	}
	return users, nil
}

// ValidateUser makes sure all user's information is satisfied
func ValidateUser(u *User) (bool, error) {
	// sanitization
	u.sanitize()
	// validate user name
	if u.Name == "" {
		return false, errors.New("missing user's name")
	}
	if matched, err := regexp.MatchString("^[a-z0-9_]+$", u.Name); err != nil {
		panic(err)
	} else if matched == false {
		return false, errors.New("invalid user's name format")
	}
	// validate user email
	if u.Email == "" {
		return false, errors.New("missing user's email")
	}
	if matched, err := regexp.MatchString(`^\w+([-+.']\w+)*@\w+([-.]\w+).\w+([-.]\w+)*$`, u.Email); err != nil || !matched {
		return false, errors.New("invalid email")
	}
	// validate uniqueness
	for _, props := range []map[string]string{
		{"name": u.Name},
		{"email": u.Email},
	} {
		for k, v := range props {
			if u, _ := FindUser(neoism.Props{k: v}); u != nil {
				return false, errors.New("user with " + k + " \"" + v + "\" already exists.")
			}
		}
	}
	// validate user password
	if u.Password == "" {
		return false, errors.New("missing user's password")
	}
	return true, nil
}

// sanitize convert user's name and email to lower case
func (u *User) sanitize() {
	u.Name = strings.ToLower(strings.TrimSpace(u.Name))
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
}
