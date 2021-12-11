package vault

import (
	"encoding/hex"
	"testing"
)

func TestEncrypt(t *testing.T) {
	password := []byte("password")
	value := []byte("hello world")

	ciphertext, err := encrypt(password, value)
	if err != nil {
		t.Error(err)
	}

	plaintext, err := decrypt(password, ciphertext)
	if err != nil {
		t.Error(err)
	}

	if string(plaintext) != string(value) {
		t.Errorf("expected %s, got %s", string(value), string(plaintext))
	}
}

func TestDecrypt(t *testing.T) {
	// "hello world" encrypted with "password"
	encrypted, _ := hex.DecodeString("2db88929a4742b18787dbf0d44dc74ac95d851abc9709e85dbffe009f4ce507352408e6ab1d4c2")
	password := []byte("password")
	value := "hello world"

	plaintext, err := decrypt(password, encrypted)
	if err != nil {
		t.Error(err)
	}

	if string(plaintext) != value {
		t.Errorf("expected %s, got %s", value, string(plaintext))
	}
}
