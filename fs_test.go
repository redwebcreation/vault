package vault

import (
	"strconv"
	"testing"
	"time"
)

func TestFs_Has(t *testing.T) {
	fs := Fs("/tmp")

	if fs.Has(strconv.FormatInt(time.Now().UnixMilli(), 10) + ".tmp") {
		t.Error("File should not exist")
	}

	f := strconv.FormatInt(time.Now().UnixMilli(), 10) + ".tmp"
	err := fs.Write(f, []byte("test"))
	if err != nil {
		t.Error(err)
	}

	if !fs.Has(f) {
		t.Error("File should exist")
	}
}

func TestFs_Write(t *testing.T) {
	fs := Fs("/tmp")

	f := strconv.FormatInt(time.Now().UnixMilli(), 10) + ".tmp"
	err := fs.Write(f, []byte("test"))
	if err != nil {
		t.Error(err)
	}

	if !fs.Has(f) {
		t.Error("File should exist")
	}
}

func TestFs_Delete(t *testing.T) {
	fs := Fs("/tmp")

	f := strconv.FormatInt(time.Now().UnixMilli(), 10) + ".tmp"
	err := fs.Write(f, []byte("test"))
	if err != nil {
		t.Error(err)
	}

	if !fs.Has(f) {
		t.Error("File should exist")
	}

	err = fs.Delete(f)
	if err != nil {
		t.Error(err)
	}

	if fs.Has(f) {
		t.Error("File should not exist")
	}
}

func TestFs_Read(t *testing.T) {
	fs := Fs("/tmp")

	f := strconv.FormatInt(time.Now().UnixMilli(), 10) + ".tmp"
	err := fs.Write(f, []byte("test"))
	if err != nil {
		t.Error(err)
	}

	if !fs.Has(f) {
		t.Error("File should exist")
	}

	b, err := fs.Read(f)
	if err != nil {
		t.Error(err)
	}

	if string(b) != "test" {
		t.Error("File should contain 'test'")
	}
}
