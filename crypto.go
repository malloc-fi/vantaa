package main

import (
	"golang.org/x/crypto/bcrypt"
)

const HASH_COST = 12

func GeneratePasswordDigest(password string) (string, error) {
	digest, err := bcrypt.GenerateFromPassword([]byte(password), HASH_COST)

	if err != nil {
		return "", err
	}

	return string(digest[:len(digest)]), nil
}