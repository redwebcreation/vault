package vault

import (
	"bytes"
	"os"
	"testing"
)

func TestNewVault(t *testing.T) {
	_, err := NewVault("/not/found")
	if err == nil {
		t.Error("expected error")
	}

	err = os.MkdirAll("/tmp/vault", 0700)
	if err != nil {
		t.Error(err)
	}
	_, err = NewVault("/tmp/vault")
	if err != nil {
		t.Error(err)
	}
}

var Password = []byte("password")
var Secret = []byte("secret")
var KeyPath = "/tmp/vault/2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"

func setup(t *testing.T) *Vault {
	s := &Vault{
		Storage: Fs("/tmp/vault"),
	}

	err := os.MkdirAll("/tmp/vault", 0700)
	if err != nil {
		t.Errorf("Failed to create test directory: %s", err)
	}

	return s
}

//func TestLocalStorage_pathFor(t *testing.T) {
//	s := &LocalStorage{
//		Path: "/vault",
//	}
//
//	if s.pathFor("foo") != "/vault/foo" {
//		t.Errorf("Expected pathFor to return /vault/foo, got %s", s.pathFor("foo"))
//	}
//
//	if s.pathFor("foo/bar") != "/vault/foo/bar" {
//		t.Errorf("Expected pathFor to return /vault/foo/bar, got %s", s.pathFor("foo/bar"))
//	}
//}

func TestLocalStorage_Delete(t *testing.T) {
	s := setup(t)
	defer teardown(t)

	err := s.Set("foo", Secret, Password)
	if err != nil {
		t.Errorf("Failed to set value: %s", err)
	}

	if !s.Has("foo") {
		t.Errorf("Key foo should exist")
	}

	if err = s.Delete("foo"); err != nil {
		t.Error(err)
	}

	if s.Has("foo") {
		t.Error("File should have been deleted")
	}
}

func TestLocalStorage_Get(t *testing.T) {
	s := setup(t)
	defer teardown(t)

	err := s.Set("foo", Secret, Password)
	if err != nil {
		t.Errorf("Failed to create test file: %s", err)
	}

	contents, err := s.Get("foo", Password)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(contents, Secret) {
		t.Errorf("Expected get to return 'secret', got '%s'", contents)
	}
}

func TestLocalStorage_Get2(t *testing.T) {
	s := setup(t)
	defer teardown(t)

	_, err := s.Get("non-existent", Password)
	if err == nil {
		t.Error("Expected error when a key does not exist")
	}
}

func TestLocalStorage_Get3(t *testing.T) {
	s := setup(t)
	defer teardown(t)

	err := os.WriteFile(KeyPath, []byte("INVALID VAULT FILE"), 0777)
	if err != nil {
		t.Errorf("Failed to create test file: %s", err)
	}

	_, err = s.Get("foo", Password)
	if err != ErrInvalidVaultFile {
		t.Errorf("Expected get to return ErrInvalidVaultFile, got %s", err)
	}
}

func TestLocalStorage_Get4(t *testing.T) {
	s := setup(t)
	defer teardown(t)

	err := os.WriteFile(KeyPath, []byte("$VAULT;0.0;AES256\n0"), 0777)
	if err != nil {
		t.Errorf("Failed to create test file: %s", err)
	}

	_, err = s.Get("foo", Password)
	if err != ErrInvalidVaultVersion {
		t.Errorf("Expected get to return ErrInvalidVaultVersion, got %s", err)
	}
}

func TestLocalStorage_Get5(t *testing.T) {
	s := setup(t)
	defer teardown(t)

	err := os.WriteFile(KeyPath, []byte("$VAULT;1.0;INVALID\n0"), 0777)
	if err != nil {
		t.Errorf("Failed to create test file: %s", err)
	}

	_, err = s.Get("foo", Password)
	if err != ErrInvalidCipher {
		t.Errorf("Expected get to return ErrInvalidCipher, got %s", err)
	}
}

func TestLocalStorage_Set2(t *testing.T) {
	s := setup(t)
	defer teardown(t)

	err := s.Set("nested.key", Secret, Password)
	if err != nil {
		t.Error(err)
	}

	if !s.Storage.Has("nested.key") {
		t.Error(err)
	}
}

func TestLocalStorage_Set3(t *testing.T) {
	s := &Vault{
		Storage: Fs("/vault"), // permission denied
	}

	err := s.Set("foo", Secret, Password)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestLocalStorage_Has(t *testing.T) {
	s := setup(t)
	defer teardown(t)

	if s.Has("foo") {
		t.Errorf("Expected Has(foo) to return false, got true")
	}

	err := s.Set("foo", []byte("bar"), []byte("password"))

	if err != nil {
		t.Error(err)
	}

	if s.Has("foo") == false {
		t.Errorf("Expected Has(foo) to return true, got false")
	}
}

func TestLocalStorage_Set(t *testing.T) {
	s := setup(t)
	defer teardown(t)

	err := s.Set("foo", []byte("bar"), []byte("password"))
	if err != nil {
		t.Error(err)
	}

	if !s.Has("foo") {
		t.Error(err)
	}

	value, err := s.Get("foo", []byte("password"))
	if err != nil {
		t.Error(err)
	}

	if string(value) != "bar" {
		t.Errorf("Expected get to return 'bar', got '%s'", value)
	}
}

func teardown(t *testing.T) {
	err := os.RemoveAll("/tmp/vault")
	if err != nil {
		t.Error(err)
	}
}
