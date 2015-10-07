package vantaadb

import (
	"os"
	"testing"

	"github.com/nathandao/vantaa/settings"
)

var err = os.Setenv("GO_ENV", "test")

func TestConnection(t *testing.T) {
	db := Connect()
	expecturl := settings.Get().DbUrl

	if db.Url != expecturl {
		t.Error(
			"Expected ", expecturl,
			"got ", db.Url)
	}
}
