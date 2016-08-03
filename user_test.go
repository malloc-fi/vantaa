package main

import (
	"fmt"
	"testing"
)

func TestSt(t *testing.T) {
	RemoveAllNodes()
	RemoveAllQuads()

	u := User{
		Name:           "name",
		PasswordDigest: "pdigest",
		Email:          "mail",
	}

	u.Save()

	u2 := User{
		Name:           "name2",
		PasswordDigest: "pdigest2",
		Email:          "mail2",
	}

	u2.Save()

	store, _ := Db()

	it := store.QuadsAllIterator()
	defer it.Close()

	for it.Next() {
		token := it.Result()
		fmt.Println("QUADFOUND:", store.Quad(token))
	}
}

// func TestSave(t *testing.T) {
// 	defer RemoveAllNodes()
// 	passwordDigest, err := GeneratePasswordDigest("vantaaRocks")

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// Test for missing info
// 	u := User{
// 		Name:           "testu",
// 		PasswordDigest: passwordDigest,
// 		Email:          "somemail",
// 	}

// 	u2, err := u.Save()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	uEmail, err := GetObjects(u2.Name, HAS_EMAIL)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if len(uEmail) != 1 || uEmail[0] != "somemail" {
// 		t.Error("Expected user name to be testuser, got", uEmail)
// 	}
// }

// func TestCreate(t *testing.T) {
// 	RemoveAllNodes()
// 	defer RemoveAllNodes()
// 	u := User{
// 		Name:     "testuser",
// 		Email:    "vantaaATexampleDOTCOM",
// 		Password: "vantaaRocks",
// 	}

// 	_, err := u.Create()
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestVerify(t *testing.T) {
// 	RemoveAllNodes()
// 	defer RemoveAllNodes()
// 	u := User{
// 		Email:          "vantaaATexampleDOTCOM",
// 		PasswordDigest: "vantaaRocks",
// 	}
// 	err := u.Verify()
// 	if err == nil {
// 		t.Error("Expected user without Name to fail verification, got passed")
// 	}

// 	u = User{
// 		Name:           "testuser",
// 		PasswordDigest: "vantaaRocks",
// 	}
// 	err = u.Verify()
// 	if err == nil {
// 		t.Error("Expected user without Email to fail verification, got passed")
// 	}

// 	u = User{
// 		Name:  "testuser",
// 		Email: "vantaaATexampleDOTCOM",
// 	}
// 	err = u.Verify()
// 	if err == nil {
// 		t.Error("Expected user without PasswordDigest to fail verification, got passed")
// 	}
// }
