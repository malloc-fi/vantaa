package neo

import (
	"github.com/nathandao/vantaa/neo"
	"testing"
)

func TestConnection(t *testing.T) {
	db := neo.Connect()

	if db.Url != "http://neo4j:admin@localhost:7474/db/data/" {
		t.Error("Expectded http://neo4j:admin@localhost:7474/db/data/, got ", db.Url)
	}
}
