package vault

import (
	"os"
	"testing"
)

func TestNewVault(t *testing.T) {
	_, err := NewVault("/not/found")
	if err == nil {
		t.Error("expected error")
	}

	os.MkdirAll("/tmp/vault", 0700)
	_, err = NewVault("/tmp/vault")
	if err != nil {
		t.Error(err)
	}
}
