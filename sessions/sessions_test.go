package sessions

import (
	"bytes"
	"testing"

	"github.com/nathandao/vantaa/model"
	"github.com/nathandao/vantaa/testhelper"
)

func TestValidLogin(t *testing.T) {
	defer testhelper.ClearNeo()

	u := model.DummyUser()
	// need this coz u.Password will be removed automatically after user.Save()
	pw := u.Password
	u2, _ := u.Save()
	s, err := Login(u.Name, pw)

	// normal success login
	if err != nil {
		t.Error(
			"Expected user login to not receive error",
			"got error", err,
		)
	}
	if s.Sid == nil {
		t.Error(
			"Expected a session Id to be created",
			"got empty string",
		)
	}

	// unique sid
	s2, _ := CreateSession(u2.Id)
	if bytes.Equal(s.Sid, s2.Sid) {
		t.Error(
			"Expected session id to be unique",
			"got the same Sid", s2.Sid,
		)
	}
}

func TestInvalidLogin(t *testing.T) {
	defer testhelper.ClearNeo()

	// user not found
	if _, err := Login("hacker", "nowork"); err == nil {
		t.Error(
			"When not found user info entered",
			"expected login to return error",
			"got no error",
		)
	}

	// user found
	u := model.CreateDummyUser()
	if _, err := Login(u.Name, "failedpass"); err == nil {
		t.Error(
			"When user fould but wrong password",
			"expected login to return error",
			"got no error",
		)
	}
}
