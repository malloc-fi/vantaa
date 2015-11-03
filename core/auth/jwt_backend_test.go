package auth

import (
	"fmt"
	"strings"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nathandao/vantaa/services/models/user"
	"github.com/nathandao/vantaa/testhelpers"
)

// Make sure generated token string at least has the correct format:
// xxx.xxx.xxx
func TestTokenGeneration(t *testing.T) {
	authBackend, _ := InitJwtAuthBackend()

	uid := 12
	token, err := authBackend.GenerateToken(uid)

	if err != nil {
		t.Error(
			"When token data is valid,",
			"Expected token generation to be valid",
			"got error", err,
		)
	}

	if len(strings.Split(token, ".")) != 3 {
		t.Error("Wrong token format was generated.")
	}
}

// Test valid and invalid authentication
func TestAuthenticate(t *testing.T) {
	authBackend, _ := InitJwtAuthBackend()
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

func TestTokenAuth(t *testing.T) {
	authBackend, _ := InitJwtAuthBackend()
	defer ClearAllTokens()

	// Generate a token string
	uid := 12
	tokenStr, _ := authBackend.GenerateToken(uid)

	// Parse token string and make sure it is valid
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})

	// Make sure token is valid
	if err != nil {
		t.Error(
			"Expected token to have valid signing method,",
			"got invalid with error: ", err,
		)
	}

	if int(token.Claims["uid"].(float64)) != uid {
		t.Error(
			"Expected token to include uid ", uid,
			"got", token.Claims["uid"],
		)
	}
}
