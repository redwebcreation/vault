package vault

import (
	"os"
	"testing"
)

func TestPathFor(t *testing.T) {
	s := &LocalStorage{
		Path: "/vault",
	}

	if s.pathFor("foo") != "/vault/foo" {
		t.Errorf("Expected pathFor to return /vault/foo, got %s", s.pathFor("foo"))
	}

	if s.pathFor("foo/bar") != "/vault/foo/bar" {
		t.Errorf("Expected pathFor to return /vault/foo/bar, got %s", s.pathFor("foo/bar"))
	}
}

func TestRead(t *testing.T) {
	s := &LocalStorage{
		Path: "/tmp/vault",
	}

	os.MkdirAll(s.Path, 0777)
	os.WriteFile(s.pathFor("foo"), []byte("bar"), 0777)

	contents, err := s.read("foo")
	if err != nil {
		t.Error(err)
	}

	if string(contents) != "bar" {
		t.Errorf("Expected read to return bar, got %s", contents)
	}

	teardown()
}

func TestDelete(t *testing.T) {
	s := &LocalStorage{
		Path: "/tmp/vault",
	}

	os.MkdirAll(s.Path, 0777)
	os.WriteFile(s.pathFor("foo"), []byte("bar"), 0777)

	if _, err := os.Stat(s.pathFor("foo")); err != nil {
		t.Error(err)
	}

	if err := s.Delete("foo"); err != nil {
		t.Error(err)
	}

	teardown()
}

// TODO: Missing Get, Set, Has tests
func TestGet(t *testing.T) {
	s := &LocalStorage{
		Path: "/tmp/vault",
	}

	os.MkdirAll(s.Path, 0777)
	os.WriteFile("/tmp/vault/foo", []byte("$VAULT;1.0;AES256\n58938210ba2ff0dd3dfd3dbeac176986bd9b4563d27110b4ca547fc8eb3c767a5105"), 0777)

	contents, err := s.Get("foo", []byte("password"))
	if err != nil {
		t.Error(err)
	}

	if string(contents) != "secret" {
		t.Errorf("Expected get to return 'secret', got '%s'", contents)
	}

	teardown()
}

func TestHas(t *testing.T) {
	s := &LocalStorage{
		Path: "/tmp/vault",
	}

	os.MkdirAll(s.Path, 0777)

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

	teardown()
}

func TestSet(t *testing.T) {
	s := &LocalStorage{
		Path: "/tmp/vault",
	}

	os.MkdirAll(s.Path, 0777)

	s.Set("foo", []byte("bar"), []byte("password"))

	if _, err := os.Stat(s.pathFor("foo")); err != nil {
		t.Error(err)
	}

	value, err := s.Get("foo", []byte("password"))
	if err != nil {
		t.Error(err)
	}

	if string(value) != "bar" {
		t.Errorf("Expected get to return 'bar', got '%s'", value)
	}

	teardown()
}
func teardown() {
	os.RemoveAll("/tmp/vault")
}
