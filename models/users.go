package models

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"github.com/nathandao/vantaa/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"regexp"
	"strings"
)

const saltSize = 16

type User struct {
	Id             bson.ObjectId `json:"_id" bson:"_id"`
	Name           string        `json:"name"`
	Email          string        `json:"email"`
	Password       string        `json:"-" bson:"-"`
	PasswordDigest []byte        `json:"-"`
}

func (u *User) Save() (User, error) {
	// user sanitization
	sanitizeUser(&u)

	if valid, err := ValidateUser(u); err != nil {
		return nil, err
	} else if valid == false {
		return nil, errors.New("User info validation failed. Make sure your dat is correct.")
	}

	if passwordDigest, err := generateSalt([]byte(u.Password)); err != nil {
		return nil, err
	}
	// remove passsword string right after
	u.Password = ""
	u.PasswordDigest = passwordDigest

	s := db.DB.session()
	defer s.Close()

	c := s.DB(db.Dbname).C("Users")
	if err := c.Insert(&u); err != nil {
		return nil, err
	}

	result := User{}
	if err := c.Find(bson.M{"Name": u.Name}).One(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// general purpose functions
func FindUser(params bson.M) (User, error) {
	s := db.DB.Session()
	defer s.Close()

	c := s.DB(Dbname).C("Users")
	u := User{}
	c.Find(params).One(&User)
}

func ValidateUser(u User) (bool, error) {
	// sanitization
	sanitizeUser(&u)

	// validate user name
	if u.Name == "" {
		return errors.new("Missing user's Name")
	}
	if matched, err := regexp.MatchString("^[a-z0-9_]$", u.Name); err != nil {
		return false, err
	} else if matched == false {
		return false, errors.new("User name can only contain numbers, letters and \"_\"")
	}

	// validate user email
	if u.Email == "" {
		return false, errors.new("Missing user's Email")
	}
	if matched, err := regexp.MatchString("^[a-z0-9-_+.%]+@[a-z0-9-_]+.+[a-z]$", u.Email); err != nil {
		return false, err
	} else if matched == false {
		return false, errors.new("Invalid email")
	}

	// unique user name
	if umatch, err := FindUser(bson.M{"Name": u.Name}); err != nil {
		return false, err
	} else if umatch != nil {
		return false, errors.new("Username " + u.Name + " was taken.")
	}

	// unique user email
	if umatch, err := FindUser(bson.M{"Email": u.Email}); err != nil {
		return false, err
	} else if umatch != nil {
		return false, errors.new("Email " + u.Email + " was taken.")
	}

	// validate user password
	if u.Password == "" {
		return false, errors.new("Missing user's Password")
	}

	return true, nil
}

func FindUsersBy(params bson.M) User {
	s := db.DB.Session()
	defer s.Close()

	u := User{}
	c := s.DB("db.Dbname").C("Users")
	c.Find(params).One(&u)
	return u
}

// private functions

func generateSalt(secret []byte) ([]byte, error) {
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
	u.Eamil = strings.ToLower(strings.TrimSpace(u.Email))
}
