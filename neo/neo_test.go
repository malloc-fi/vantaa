package neo

import (
	"github.com/jmcvetta/neoism"
	"testing"
)

func TestConnection(t *testing.T) {
	db := Connect()

	if db.Url != "http://neo4j:admin@localhost:7474/db/data/" {
		t.Error("Expectded http://neo4j:admin@localhost:7474/db/data/, got ", db.Url)
	}
}

func TestPropString(t *testing.T) {
	props := neoism.Props{
		"id":    1234,
		"name":  "admin",
		"email": "admin@example.com",
	}

	expectstr := "id(u) = {id} and u.name = {name} and u.email = {email} "

	if str := PropString("u", props); str != expectstr {
		t.Error(
			"Expect ", expectstr,
			"got", str,
		)
	}
}
