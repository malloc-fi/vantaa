package main

import "testing"

func TestDb(t *testing.T) {
	_, err := Db()

	if err != nil {
		t.Error(err)
	}
}
