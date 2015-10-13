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
	expireOffset = 3600
)

var authBackendInstance *JwtAuthBackend = nil

// InitJwtAuthBackend instantiate a thread-safe JwtAuthBackend instance if
// it has not been started.
func InitJwtAuthBackend() *JwtAuthBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JwtAuthBackend{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}
	return authBackendInstance
}

// GenerateToken creates an encrypted token including the user's ID using
// signing method RS512 and the public/private key
func (authBackend *JwtAuthBackend) GenerateToken(uid int) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims["exp"] = time.Now().Add(
		time.Hour * time.Duration(settings.Get().JWTExpirationDelta),
	).Unix()
	token.Claims["iat"] = time.Now().Unix()
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

// getTokenRemaining get the remaining valid duration of the token
func (authBackend *JwtAuthBackend) getTokenRemaining(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

// getPrivateKey returns the private key provided in settings
func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	return privateKeyImported
}

// getPublicKey returns the public key provided in settings
func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open(settings.Get().PublicKeyPath)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)
	if !ok {
		panic(err)
	}
	return rsaPub
}
