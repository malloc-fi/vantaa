package auth

import (
	"time"

	"github.com/boltdb/bolt"
)

const (
	dbName = "jwt.db"
)

func GetToken() {
}

// connect and return a *bolt.DB instance
func connect() (*bolt.DB, error) {
	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	return db, nil
}
