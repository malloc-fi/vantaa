package model

import (
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"io"
	"regexp"
	"strings"

	"github.com/jmcvetta/neoism"
	"github.com/nathandao/vantaa/neo"
)

const saltSize = 16

type User struct {
	Id             int    `json:"id(u)"`
	Name           string `json:"u.name"`
	Email          string `json:"u.email"`
	Password       string `json:"-"`
	PasswordDigest []byte `json:"u.password_digest"`
}

// Save or u.Save calls on CreateUser to write a newu User node in the database.
func (u *User) Save() (*User, error) {
	newu, err := CreateUser(u)
	if err != nil {
		return nil, err
	}
	return newu, nil
}

func (u *User) Delete() error {
	props := neoism.Props{}
	if u.Id != 0 {
		props["id"] = u.Id
	}
	if u.Name != "" {
		props["name"] = u.Name
	}
	if u.Email != "" {
		props["email"] = u.Email
	}

	if err := DeleteUser(props); err != nil {
		return err
	}

	return nil
}

// CreateUser validates the User's information. Upon successful validations,
// it creates a new User node in the database,
func CreateUser(u *User) (*User, error) {
	// user sanitization
	sanitizeUser(u)

	if _, err := validateUser(u); err != nil {
		return nil, err
	}

	passwordDigest, err := GenrerateSalt([]byte(u.Password))
	if err != nil {
		return nil, err
	}
	// remove passsword string right after
	u.Password = ""
	u.PasswordDigest = passwordDigest

	res := []User{}
	db := neo.Connect()
	cq := neoism.CypherQuery{
		Statement: `CREATE (u:User {name:{name}, email:{email}, password_digest:{password_digest}})
                RETURN id(u), u.name, u.email`,
		Parameters: neoism.Props{"name": u.Name, "email": u.Email, "password_digest": u.PasswordDigest},
		Result:     &res,
	}
	if err := db.Cypher(&cq); err != nil {
		return nil, err
	}

	newu := &res[0]

	return newu, nil
}

// DeleteUser removes a User node from the database.
// error is return even when user is not found
// TODO: handling removal of all relationships to
// the User node
func DeleteUser(props neoism.Props) error {

	// return error if user is not found in the database
	if u, _ := FindUser(props); u == nil {
		return errors.New("user not found")
	}

	db := neo.Connect()
	cq := neoism.CypherQuery{
		Statement: `MATCH (u:User)
                WHERE ` + neo.PropString("u", props) + `DELETE u`,
		Parameters: props,
	}
	if err := db.Cypher(&cq); err != nil {
		return err
	}
	return nil
}

// FindUser finds a single User by calling on FindUsers and returns the
// first element of the []*User slice.
func FindUser(props neoism.Props) (*User, error) {
	u, err := FindUsers(props)
	if err != nil || len(u) <= 0 {
		return nil, err
	}
	return u[0], nil
}

// FindUsers contructs and executes a cypher query with the provided parametrs
// and return all User nodes that satisfy the props conditions.
func FindUsers(props neoism.Props) ([]*User, error) {
	db := neo.Connect()
	res := []User{}

	// generate condition string to be used in the cypher statement
	condstr := neo.PropString("u", props)
	cq := neoism.CypherQuery{
		Statement: `MATCH (u:User)
                WHERE ` + condstr + `
                RETURN id(u), u.name, u.email, u.password_digest`,
		Parameters: props,
		Result:     &res,
	}

	err := db.Cypher(&cq)

	if err != nil {
		return []*User{}, nil
	}

	users := []*User{}

	for _, u := range res {
		users = append(users, &u)
	}

	if len(users) < 1 {
		return nil, errors.New("not found")
	}

	return users, nil
}

func validateUser(u *User) (bool, error) {
	// sanitization
	sanitizeUser(u)

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

	if matched, err := regexp.MatchString("^[a-z0-9-_+.%]+@[a-z0-9-_]+\\.+[a-z]+$", u.Email); err != nil {
		panic(err)
	} else if matched == false {
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

func GenrerateSalt(secret []byte) ([]byte, error) {
	buf := make([]byte, saltSize, saltSize+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		return nil, err
	}

	hash := sha1.New()
	hash.Write(buf)
	hash.Write(secret)
	return hash.Sum(buf), nil
}

func sanitizeUser(u *User) {
	u.Name = strings.ToLower(strings.TrimSpace(u.Name))
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
}
