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
	Password       string `"-"`
	PasswordDigest []byte `json:"u.password_digest"`
}

func (u *User) Save() (*User, error) {
	newu, err := CreateUser(u)
	if err != nil {
		return nil, err
	}
	return newu, nil
}

func CreateUser(u *User) (*User, error) {
	// user sanitization
	sanitizeUser(u)

	if valid, err := validateUser(u); err != nil || valid == false {
		return nil, errors.New("user info validation failed. make sure your dat is correct.")
	}

	passwordDigest, err := generatesalt([]byte(u.Password))
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
                RETURN id(u), u.name, u.email, u.password_digest`,
		Parameters: neoism.Props{"name": u.Name, "email": u.Email, "password_digest": u.PasswordDigest},
		Result:     &res,
	}

	if err := db.Cypher(&cq); err != nil {
		return nil, err
	}

	newu := &res[0]

	return newu, nil
}

func FindUser(props neoism.Props) (*User, error) {
	u, err := FindUsers(props)
	if err != nil || len(u) <= 0 {
		return nil, err
	}
	return u[0], nil
}

func FindUsers(props neoism.Props) ([]*User, error) {
	db := neo.Connect()
	res := []User{}

	// generate condition string to be used in the cypher statement
	condstr := ""
	for k, _ := range props {
		condstr += k + ": {" + k + "},"
	}
	if condstr != "" {
		condstr = condstr[:len(condstr)-1]
	}

	cq := neoism.CypherQuery{
		Statement: `MATCH (u:User {` + condstr + `})
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

	return users, nil
}

func validateUser(u *User) (bool, error) {
	// sanitization
	sanitizeUser(u)

	// validate user name
	if u.Name == "" {
		return false, errors.New("Missing user's Name")
	}

	if matched, err := regexp.MatchString("^[a-z0-9_]+$", u.Name); err != nil {
		panic(err)
	} else if matched == false {
		return false, errors.New("User name can only contain numbers, letters and \"_\"")
	}

	// validate user email
	if u.Email == "" {
		return false, errors.New("Missing user's Email")
	}

	if matched, err := regexp.MatchString("^[a-z0-9-_+.%]+@[a-z0-9-_]+\\.+[a-z]+$", u.Email); err != nil {
		panic(err)
	} else if matched == false {
		return false, errors.New("Invalid email")
	}

	// validate uniqueness
	for _, props := range []map[string]string{
		{"name": u.Name},
		{"email": u.Email},
	} {
		for k, v := range props {
			if u, _ := FindUser(neoism.Props{k: v}); u != nil {
				return false, errors.New("User with " + k + " \"" + v + "\" already exists.")
			}
		}
	}

	// validate user password
	if u.Password == "" {
		return false, errors.New("missing user's password")
	}

	return true, nil
}

func generatesalt(secret []byte) ([]byte, error) {
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
