package neo

import (
	"testing"
)

func TestConnection(t *testing.T) {
	db := Connect()

	if db.Url != "http://neo4j:admin@localhost:7474/db/data/" {
		t.Error("Expectded http://neo4j:admin@localhost:7474/db/data/, got ", db.Url)
	}
}
