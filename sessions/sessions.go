package sessions

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
	"time"

	"github.com/jmcvetta/neoism"
	"github.com/nathandao/vantaa/model"
	"github.com/nathandao/vantaa/neo"
)

const (
	// Todo: make MaxAge configurable
	MaxAge = 48 * time.Hour
)

type Session struct {
	Sid       []byte    `json:"s.sid"`
	Uid       int       `json:"id(u)"`
	Created   time.Time `json:"s.created"`
	LastLogin time.Time `json:"s.last_login"`
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
                CREATE (s:Session {sid: {sid}, created: {created}, last_login: {last_login}})
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

// FindSession find the Session from the database using sid
func FindSession(sid []byte) (*Session, error) {
	db := neo.Connect()
	r := []Session{}
	cq := neoism.CypherQuery{
		Statement: `MATCH (s:Session {sid: {sid}})-[r:BELONGS_TO]->(u)
                RETURN s.sid, s.created, s.last_login, id(u)`,
		Parameters: neoism.Props{"sid": sid},
		Result:     &r,
	}
	if err := db.Cypher(&cq); err != nil {
		return nil, err
	}
	return &r[0], nil
}

// CreateCookieValue concatinates the session Sid, SecretKey, and user Id,
// then encrypts the resulted []bytes. The encrypted value will be stored in
// the brower's cookies
func CreateCookieValue(s *Session) ([]byte, error) {
	str := string(s.Sid) + SecretKey + string(s.Uid)
	enc, err := Encrypt([]byte(str))
	if err != nil {
		return nil, err
	}
	return enc, nil
}

// DecryptCookieValueValue decrypts the cookie and returns the session Sid
// and user's Id
func DecryptCookieValue(enc []byte) ([]byte, int, error) {
	dec, err := Decrypt(enc)
	if err != nil {
		return nil, 0, err
	}
	strs := make([]string, 2)
	strs = strings.Split(string(dec), SecretKey)
	if len(strs) != 2 {
		return nil, 0, errors.New("bad cookie format. Session cookie value should be encrypted from Sid + SecretKey + Uid")
	}
	sid := []byte(strs[0])
	uid, err := strconv.Atoi(strs[1])
	if err != nil {
		return nil, 0, err
	}
	return sid, uid, nil
}

// DeleteSession removes a session from the database
func DeleteSession(s *Session) error {
	db := neo.Connect()
	cq := neoism.CypherQuery{
		Statement: `MATCH (s:Session) WHERE s.sid = {sid}
                OPTIONAL MATCH (s)-[r]->(u:User)
                DELETE s, r`,
		Parameters: neoism.Props{"sid": s.Sid},
	}
	if err := db.Cypher(&cq); err != nil {
		return err
	}
	return nil
}

// AuthenticateCookie check if the cookie id exists and belongs to the correct
// user id
func AuthenticateCookie(enc []byte) (bool, error) {
	sid, uid, err := DecryptCookieValue(enc)
	if err != nil {
		return false, err
	}
	s, err := FindSession(sid)
	if err != nil {
		return false, err
	}
	if s.Created.Add(MaxAge).After(time.Now().Local()) {
		if err := DeleteSession(s); err != nil {
			return false, err
		}
		return false, errors.New("cookie expired")
	}
	if s.Uid == uid {
		return true, nil
	}
	if err := DeleteSession(s); err != nil {
		return false, err
	}
	return false, errors.New("session and user id not match")
}
