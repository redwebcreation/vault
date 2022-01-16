package vault

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestEncrypt(t *testing.T) {

	ciphertext, err := Encrypt(Password, Secret)
	if err != nil {
		t.Error(err)
	}

	plaintext, err := Decrypt(Password, ciphertext)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(plaintext, Secret) {
		t.Errorf("expected %s, got %s", Secret, plaintext)
	}
}

func TestDecrypt(t *testing.T) {
	// "hello world" encrypted with "password"
	encrypted, _ := hex.DecodeString("2db88929a4742b18787dbf0d44dc74ac95d851abc9709e85dbffe009f4ce507352408e6ab1d4c2")
	value := "hello world"

	plaintext, err := Decrypt(Password, encrypted)
	if err != nil {
		t.Error(err)
	}

	if string(plaintext) != value {
		t.Errorf("expected %s, got %s", value, string(plaintext))
	}
}
