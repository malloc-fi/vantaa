package settings

import (
	"os"
	"testing"
)

var err error = os.Setenv("GO_ENV", "test")

func TestSettings(t *testing.T) {

	if err != nil {
		panic(err)
	}

	s := Get()

	if s.HashCost != 10 {
		t.Error(
			"Expected HashCost to equals 10,",
			"got", s.HashCost,
		)
	}

	if s.JWTExpirationDelta != 72 {
		t.Error(
			"Expected JWTExpirationDelta to equals 72,",
			"got", s.JWTExpirationDelta,
		)
	}

	expecturl := "http://neo4j:foobar@localhost:9290/db/data/"
	if s.DbUrl != expecturl {
		t.Error(
			"Expected DbUrl to equals", expecturl,
			"got", s.DbUrl,
		)
	}
	// TODO: Test private and public keypath
}
