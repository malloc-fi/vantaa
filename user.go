package main

import (
	"errors"
	"github.com/satori/go.uuid"
)

type User struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	PasswordDigest string `json:"passwordDigest"`
	Password       string `json:"-"`
}

// Save a new user to the database with his/her own credentials
func (u *User) Save() (*User, error) {
	err := u.Verify()
	if err != nil {
		return u, err
	}

	t := NewTransaction()
	t.AddQuad(MakeQuad(VANTAA_BLOG, HAS_USER, u.Id, USERS_GRAPH))
	t.AddQuad(MakeQuad(u.Id, HAS_NAME, u.Name, USERS_GRAPH))
	t.AddQuad(MakeQuad(u.Id, HAS_EMAIL, u.Email, USERS_GRAPH))
	t.AddQuad(MakeQuad(u.Id, HAS_PASSWORD_DIGEST, u.PasswordDigest, USERS_GRAPH))

	if err := ApplyTransaction(t); err != nil {
		return u, err
	}

	return u, nil
}

// Create a new user.
func (u *User) Create() (*User, error) {
	if u.Password == "" {
		return u, errors.New("Missing user's password")
	}

	passwordDigest, err := GeneratePasswordDigest(u.Password)
	if err != nil {
		return u, err
	}
	u.PasswordDigest = passwordDigest

	if err := u.Verify(); err != nil {
		return u, err
	}

	return u.Save()
}

// Verify makes sure user info is valid.
func (u *User) Verify() error {
	if u.PasswordDigest == "" {
		return errors.New("Missing user's PasswordDigest")
	}

	// TODO: Make sure u.Email is unique.
	if u.Email == "" {
		return errors.New("Missing user's Email")
	}

	// TODO: Make sure u.Name is unique.
	if u.Name == "" {
		return errors.New("Missing user's Name")
	}

	if u.Id == "" {
		uid := uuid.NewV4()
		u.Id = string(uid[:len(uid)])
	}

	return nil
}
