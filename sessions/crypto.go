package sessions

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io"
)

const (
	SaltSize  = 16
	SecretKey = "7Pdvgdo6nNCJ41UdXynDgozdATUE1vaI"
)

// SaltedHash generate a random session ID hash.
func SaltedHash(secret []byte) ([]byte, error) {
	buf := make([]byte, SaltSize, SaltSize+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return nil, err
	}
	h := sha1.New()
	h.Write(buf)
	h.Write([]byte(secret))
	return h.Sum(buf), nil
}

// MatchSaltedHash compares and see if 2 byte arrays match or not
func MatchSaltedHash(data, target []byte) (bool, error) {
	if len(data) != SaltSize+sha1.Size {
		return false, nil
	}
	h := sha1.New()
	h.Write(data[:SaltSize])
	h.Write(target)
	if eq := bytes.Equal(h.Sum(nil), data[SaltSize:]); !eq {
		return false, errors.New("not match")
	}
	return true, nil
}

// Encrypt encrypts a string byte based on the SecretKey. This is used to
// encrypt session data on the brower's cookie
func Encrypt(str []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(SecretKey))
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(str))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.
		ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(ciphertext[aes.BlockSize:], str)
	return ciphertext, nil
}

// Decrypt decrypts and []byte array based on SecretKey. This is used to decrypt
// data received from the browser's cookie
func Decrypt(enc []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(SecretKey))
	if err != nil {
		return nil, err
	}
	if len(enc) < aes.BlockSize {
		return nil, errors.New("cipher text too short")
	}
	iv := enc[:aes.BlockSize]
	cenc := enc[aes.BlockSize:]
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(cenc, cenc)
	data, err := base64.StdEncoding.DecodeString(string(cenc))
	if err != nil {
		return nil, err
	}
	return data, nil
}
