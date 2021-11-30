package main

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	password := "password"
	hash, err := Encrypt(password, 0)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(hash) == 0 {
		t.Errorf("Hash is empty")
	}
}

func TestCompare(t *testing.T) {
	password := "password"
	hash, err := Encrypt(password, 0)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if err := Compare(hash, password); err != nil {
		t.Errorf("Password does not match hash: %s", err.Error())
	}
}
