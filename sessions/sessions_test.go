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

func TestSessionCookie(t *testing.T) {
	s := DummySession()
	cv, err := CreateCookieValue(s)
	if err != nil {
		t.Error(
			"When creating session cookie value",
			"expected no error",
			"got error", err,
		)
	}
	dec, uid, err := DecryptCookieValue(cv)
	if err != nil {
		t.Error(
			"When decrypting session cookie value",
			"expected no error",
			"got error", err,
		)
	}
	if !bytes.Equal(dec, s.Sid) {
		t.Error(
			"Expected decrypted sid from cookie to equals sid",
			"Got", string(dec), "and", string(s.Sid),
		)
	}
	if s.Uid != uid {
		t.Error(
			"Expected decrypted uid from cookie to equals s.Uid",
			"Got", uid, "and", s.Uid,
		)
	}
	if auth, err := AuthenticateCookie(cv); err != nil || !auth {
		t.Error(
			"When there is legit auth info",
			"expected cookie authentication to pass",
			"got error", err,
		)
	}
}
