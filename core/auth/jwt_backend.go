package auth

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmcvetta/neoism"
	"github.com/nathandao/vantaa/services/models/user"
	"github.com/nathandao/vantaa/settings"
)

type JwtAuthBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

const (
	// expOffset is the time that the token should be removed from ledis after
	// its expiration
	expireOffset = 3600
)

var (
	authBackendInstance *JwtAuthBackend = nil
	hashCost            int             = settings.Get().HashCost
)

// InitJwtAuthBackend instantiate a thread-safe JwtAuthBackend instance if
// it has not been started.
func InitJwtAuthBackend() (*JwtAuthBackend, error) {
	privateKey, err := getPrivateKey()
	if err != nil {
		return nil, err
	}

	publicKey, err := getPublicKey()
	if err != nil {
		return nil, err
	}

	if authBackendInstance == nil {
		authBackendInstance = &JwtAuthBackend{
			privateKey: privateKey,
			PublicKey:  publicKey,
		}
	}
	return authBackendInstance, nil
}

// GenerateToken creates an encrypted token including the user's ID using
// signing method RS512 and the public/private key
func (authBackend *JwtAuthBackend) GenerateToken(u *user.User) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims["exp"] = time.Now().Add(
		time.Hour * time.Duration(settings.Get().JWTExpirationDelta),
	).Unix()

	token.Claims["iat"] = time.Now().Unix()
	token.Claims["uid"] = u.Id
	token.Claims["email"] = u.Email

	tokenString, err := token.SignedString(authBackend.privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Authenticate authenticate the provided user login credentials
func (authBackend *JwtAuthBackend) Authenticate(u *user.User) bool {
	checkUser, err := user.FindUser(neoism.Props{"email": u.Email})
	if err != nil {
		return false
	}
	diff := bcrypt.CompareHashAndPassword(
		[]byte(checkUser.PasswordDigest),
		[]byte(u.Password),
	)
	if diff != nil {
		return false
	}
	return true
}

// TerminateToken invalidates the token before its expiration time.
// This function is invoked when the user logged out.
func (authBackend *JwtAuthBackend) TerminateToken(tokenstr string, token *jwt.Token) error {
	ledisExp := token.Claims["exp"].(int64) + expireOffset
	if err := BlacklistToken([]byte(tokenstr), ledisExp); err != nil {
		return err
	}
	return nil
}

// IsTerminated check if the token has been terminated before its expiration
func (authBackend *JwtAuthBackend) IsTerminated(tokenstr string) bool {
	terminated, _ := IsBlacklisted([]byte(tokenstr))
	if terminated {
		return true
	}
	return false
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		return nil, err
	}

	return privateKeyImported, nil
}

func getPublicKey() (*rsa.PublicKey, error) {
	publicKeyFile, err := os.Open(settings.Get().PublicKeyPath)
	if err != nil {
		return nil, err
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		return nil, err
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		return nil, err
	}

	return rsaPub, nil
}
