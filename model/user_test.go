package model

import (
	"github.com/jmcvetta/neoism"
	"github.com/nathandao/vantaa/testhelper"

	"testing"
)

func TestSaveValidUser(t *testing.T) {
	defer testhelper.ClearNeo()

	u := User{
		Name:     "admin",
		Email:    "admin@example.com",
		Password: "password",
	}

	u2, err := u.Save()
	if err != nil {
		t.Error("Expected valid user to be saved successfully, save failed with error", err)
	}

	if u2 == nil {
		t.Error("Expected saved user to be found in database, not found.")
	}

	if u2.Name != u.Name {
		t.Error("Expected saved user to have name ", u.Name, "got ", u2.Name)
	}

	if u2.Email != u.Email {
		t.Error("Expected saved user to have email ", u.Email, "got ", u2.Email)
	}

	if u2.Password != "" {
		t.Error("Expected the database to NOT store raw password, it was stored in the database.")
	}
}

func TestSaveInvalidUser(t *testing.T) {
	defer testhelper.ClearNeo()

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
			if u2, err := u.Save(); err == nil {
				t.Error(
					"For user with", k,
					"expected u.Save() to fail",
					"got user with  Id", u2.Id, "created")
			}
		}
	}
}

func TestDeleteUser(t *testing.T) {
	defer testhelper.ClearNeo()
	factory := Factory{}
	u := factory.DummyUser()

	// call user delete now should return an err not found
	// since the user is not added to the database yes
	if err := u.Delete(); err == nil {
		t.Error(
			"For non-existing user",
			"expect error",
			"got no error",
		)
	}

	u2, _ := u.Save()
	if err := u2.Delete(); err != nil {
		t.Error(
			"For existing user",
			"expect no error",
			"got error", err,
		)
	}

	if u3, _ := FindUser(neoism.Props{"name": u.Name}); u3 != nil {
		t.Error(
			"After deleting user",
			"expect FindUser to return nil",
			"got user Id", u3.Id,
		)
	}
}
