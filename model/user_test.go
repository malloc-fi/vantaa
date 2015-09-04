package model

import (
	"github.com/nathandao/vantaa/test_helper"

	"testing"
)

func TestSaveValidUser(t *testing.T) {
	defer test_helper.ClearNeo()

	u := User{
		Name:     "admin",
		Email:    "admin@example.com",
		Password: "password",
	}

	u2, err := u.Save()
	if err != nil {
		t.Error(err)
	}

	if u2 == nil {
		t.Error("Expect saved user to be found in database, not found.")
	}

	if u2.Name != u.Name {
		t.Error("Expect saved user to have name ", u.Name, "got ", u2.Name)
	}

	if u2.Email != u.Email {
		t.Error("Expect saved user to have email ", u.Email, "got ", u2.Email)
	}

	if u2.Password != "" {
		t.Error("Expect saved user to not store raw password, it got stored.")
	}
}

func TestSaveInvalidUser(t *testing.T) {
	defer test_helper.ClearNeo()

	cases := []map[string]User{
		{
			"missing name": User{
				Email:    "admin@example.com",
				Password: "password",
			},
		},
		{
			"missing email": User{
				Name:     "admin",
				Password: "password",
			},
		},
		{
			"missing password": User{
				Name:  "admin",
				Email: "admin@example.com",
			},
		},
		{
			"name \"a admin\"": User{
				Name:     "a dmin",
				Email:    "admin@example.com",
				Password: "password",
			},
		},
		{
			"email \"admin\"": User{
				Name:     "admin",
				Email:    "admin",
				Password: "password",
			},
		},
		{
			"email \"admin@admin\"": User{
				Name:     "admin",
				Email:    "admin@admin",
				Password: "password",
			},
		},
	}

	for _, c := range cases {
		for k, u := range c {
			if u2, _ := u.Save(); u2 != nil {
				t.Error("Expect user with", k, "to NOT be allowed to be created. User with Id", u2.Id, "was created")
			}
		}
	}
}
