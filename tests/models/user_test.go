package models

import (
	"github.com/nathandao/vantaa/models"
	"github.com/nathandao/vantaa/neo"
	"testing"
)

func TestSaveValidUser(t *testing.T) {
	db := neo.Connect()
	u := models.User{
		Name:     "admin",
		Email:    "admin@example.com",
		Password: "password",
	}

	u2, err := u.Save()
	if err != nil {
		t.Error(error)
	}

	if u2.Id == nil {
		t.Error("Expect saved user to have an Id, got nil")
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
	db := neo.Connect()

	users := []models.User{
		models.User{
			Email:    "admin@example.com",
			Password: "password",
		},
		models.User{
			Name:     "admin",
			Password: "password",
		},
		models.User{
			Name:  "admin",
			Email: "admin@example.com",
		},
		models.User{
			Name:     "a dmin",
			Email:    "admin@example.com",
			Password: "password",
		},
		models.User{
			Name:     "admin",
			Email:    "admin",
			Password: "password",
		},
		models.User{
			Name:     "admin",
			Email:    "admin@admin",
			Password: "password",
		},
		moldels.User{
			Name:     "admin",
			Email:    "admin@example.com",
			Password: "  ",
		},
	}

	for _, u := range users {
		if u2, err := u.Save(); err != nil {
			t.Error(err)
		} else if u2.Id != nil {
			t.Error("Expect saving invalid user to return empty user, got user id ", u2.Id, " instead")
		}
	}
}
