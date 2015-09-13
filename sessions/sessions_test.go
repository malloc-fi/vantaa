package sessions

import (
	"testing"

	"github.com/nathandao/vantaa/model"
	"github.com/nathandao/vantaa/testhelper"
)

func TestLogin(t *testing.T) {
	defer testhelper.ClearNeo()

	u := model.DummyUser()
	// need this coz u.Password will be removed automatically after user.Save()
	pw := u.Password
	u.Save()
	_, err := Login(u.Name, pw)
	if err != nil {
		t.Error(
			"Expected user login to not receive error",
			"got error", err,
		)
	}
}
