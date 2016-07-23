package main

import (
	"testing"
)

func TestSave(t *testing.T) {
	passwordDigest, err := GeneratePasswordDigest("vantaaRocks")

	if err != nil {
		t.Error(err)
	}

	// Test for missing info
	u := User{
		Name:           "testuser",
		PasswordDigest: passwordDigest,
		Email:          "user@example.com",
	}

	u2, err := u.Save()

	if err != nil {
		t.Error(err)
	}

	if u2.Id == "" {
		t.Error("Expected uid to be generated, got none")
	}

	uName, err := GetObjects(u2.Id, HAS_NAME)

	if err != nil {
		t.Error(err)
	}

	if len(uName) != 1 || uName[0] != "testuser" {
		t.Error("Expected user name to be testuser, got", uName)
	}
}

func TestCreate(t *testing.T) {
	u := User{
		Name:     "testuser",
		Email:    "vantaa@example.com",
		Password: "vantaaRocks",
	}

	_, err := u.Create()
	if err != nil {
		t.Error(err)
	}
	if u.Id == "" {
		t.Error("Expected new user to have a unique Id, got none")
	}
}

func TestVerify(t *testing.T) {
	u := User{
		Email:          "vantaa@example.com",
		PasswordDigest: "vantaaRocks",
	}
	err := u.Verify()
	if err == nil {
		t.Error("Expected user without Name to fail verification, got passed")
	}

	u = User{
		Name:           "testuser",
		PasswordDigest: "vantaaRocks",
	}
	err = u.Verify()
	if err == nil {
		t.Error("Expected user without Email to fail verification, got passed")
	}

	u = User{
		Name:  "testuser",
		Email: "vantaa@example.com",
	}
	err = u.Verify()
	if err == nil {
		t.Error("Expected user without PasswordDigest to fail verification, got passed")
	}
}
