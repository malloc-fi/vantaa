package sessions

import (
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/jmcvetta/neoism"
	"github.com/nathandao/vantaa/model"
	"github.com/nathandao/vantaa/neo"
)

const (
	MaxAge = 48 * time.Hour
)

type Session struct {
	Sid       []byte    `json:"sid"`
	Uid       int       `json:"id(u)"`
	Created   time.Time `json:"created"`
	LastLogin time.Time `json:"last_login"`
}

// Login authenticates using the user name and password and creates a new
// session uppon successful login
func Login(name string, password string) (*Session, error) {
	u, err := model.FindUser(neoism.Props{"name": name})
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword(u.PasswordDigest, []byte(password)); err != nil {
		return nil, err
	}
	s, err := CreateSession(u.Id)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CreateSessions writes a new session to the database using the User's info
func CreateSession(uid int) (*Session, error) {
	sid, err := SaltedHash([]byte(SecretKey))
	if err != nil {
		return nil, err
	}
	created := time.Now().Local()
	last_login := time.Now().Local()
	s := []Session{}
	db := neo.Connect()
	cq := neoism.CypherQuery{
		Statement: `MATCH (u:User) WHERE id(u) = {uid}
                CREATE (s:Session { sid: {sid}, created: {created}, last_login: {last_login} })
                CREATE (s)-[r:BELONGS_TO]->(u)
                RETURN s.sid, s.created, s.last_login, id(u)`,
		Parameters: neoism.Props{"sid": sid, "created": created, "last_login": last_login, "uid": uid},
		Result:     &s,
	}
	if err := db.Cypher(&cq); err != nil {
		return nil, err
	}
	return &s[0], nil
}
