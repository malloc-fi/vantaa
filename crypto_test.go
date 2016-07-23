package main

import (
	"testing"
)

func TestGeneratePasswordDigest(t *testing.T) {
	_, err := GeneratePasswordDigest("vantaaRoacks")

	if err != nil {
		t.Error(err)
	}
}
