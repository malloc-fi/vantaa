package auth

import (
	"strings"
	"testing"

	"github.com/nathandao/vantaa/services/models/user"
	"github.com/nathandao/vantaa/testhelpers"
)

var authBackend *JwtAuthBackend = InitJwtAuthBackend()

// Make sure generated token string at least has the correct format:
// xxx.xxx.xxx
func TestTokenGeneration(t *testing.T) {
	uid := 12
	token, _ := authBackend.GenerateToken(uid)
	if len(strings.Split(token, ".")) != 3 {
		t.Error("Wrong token format was generated.")
	}
}

// Test valid and invalid authentication
func TestAuthenticate(t *testing.T) {
	defer testhelpers.ClearDb()

	// First, create a dummy user
	dummyu := user.DummyUser()
	dummyu.Save()

	// Test valid authentication
	u := user.DummyUser()
	loggedin := authBackend.Authenticate(&u)
	if !loggedin {
		t.Error("Expected right user credentials to be valid, got invalid.")
	}

	// Test invalid authentication
	u2 := user.User{
		Name:     "Marc",
		Email:    "marc@marc.com",
		Password: "justhacking",
	}
	loggedin = authBackend.Authenticate(&u2)
	if loggedin {
		t.Error("Expected wrong user authentication to be invalid, go valid.")
	}
}
