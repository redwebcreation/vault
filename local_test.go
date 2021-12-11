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
}

// TODO: Missing Get, Set, Has tests
