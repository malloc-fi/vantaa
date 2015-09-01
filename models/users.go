package models

import (
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"github.com/jmcvetta/neoism"
	"github.com/nathandao/vantaa/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"regexp"
	"strings"
)

const saltSize = 16

type User struct {
	Id             bson.ObjectId `json:"_id" bson:"_id"`
	Name           string        `json:"name"`
	Email          string        `json:"email"`
	Password       string        `"-"`
	PasswordDigest []byte        `json:"-" bson:"password_digest"`
}

func (u *User) Save(s *mgo.Session) (User, error) {
	// user sanitization
	sanitizeUser(u)

	if valid, err := validateUser(u, s); err != nil {
		return User{}, err
	} else if valid == false {
		return User{}, errors.New("User info validation failed. Make sure your dat is correct.")
	}

	passwordDigest, err := generateSalt([]byte(u.Password))
	if err != nil {
		return User{}, err
	}
	// remove passsword string right after
	u.Password = ""
	u.PasswordDigest = passwordDigest

	u.Id = bson.NewObjectId()
	c := s.DB(db.Dbname).C("Users")
	if err := c.Insert(u); err != nil {
		panic(err)
		return User{}, err
	}

	result := User{}
	if err := c.Find(bson.M{"name": u.Name}).One(&result); err != nil {
		return User{}, err
	}

	return result, nil
}

// private

func validateUser(u *User, s *mgo.Session) (bool, error) {

	// sanitization
	sanitizeUser(u)

	// validate user name
	if u.Name == "" {
		return false, errors.New("Missing user's Name")
	}
	if matched, err := regexp.MatchString("^[a-z0-9_]+$", u.Name); err != nil {
		return false, err
	} else if matched == false {
		return false, errors.New("User name can only contain numbers, letters and \"_\"")
	}

	// validate user email
	if u.Email == "" {
		return false, errors.New("Missing user's Email")
	}
	if matched, err := regexp.MatchString("^[a-z0-9-_+.%]+@[a-z0-9-_]+.+[a-z]+$", u.Email); err != nil {
		return false, err
	} else if matched == false {
		return false, errors.New("Invalid email")
	}

	// unique user name
	if umatch := findUser(bson.M{"name": u.Name}, s); umatch.Name == u.Name {
		return false, errors.New("Username " + u.Name + " was taken.")
	}

	// unique user email
	if umatch := findUser(bson.M{"email": u.Email}, s); umatch.Email == u.Email {
		return false, errors.New("Email " + u.Email + " was taken.")
	}

	// validate user password
	if u.Password == "" {
		return false, errors.New("Missing user's Password")
	}

	return true, nil
}

func findUser(params bson.M, s *mgo.Session) User {
	u := User{}
	s.DB(db.Dbname).C("Users").Find(params).One(&u)
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
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
}
