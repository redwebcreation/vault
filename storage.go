package vault

import (
	"crypto/sha256"
	"fmt"
	"os"
	gopath "path"
)

// Storage is the backend for a Vault.
type Storage interface {
	Read(key string) ([]byte, error)
	Write(key string, data []byte) error
	Has(key string) bool
	Delete(key string) error
}

// Fs is a storage backend that stores data in the file system.
type Fs string

// Read reads the data from the file system.
func (f Fs) Read(key string) ([]byte, error) {
	path := f.path(key)
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(gopath.Clean(path))
}

// Write writes the data to the file system.
func (f Fs) Write(key string, data []byte) error {
	path := f.path(key)

	return os.WriteFile(path, data, 0600)
}

// Delete deletes the file from the file system.
func (f Fs) Delete(key string) error {
	return os.Remove(f.path(key))
}

// Has checks if the key exists in the file system.
func (f Fs) Has(key string) bool {
	_, err := os.Stat(f.path(key))
	return err == nil
}

func (f Fs) path(key string) string {
	path := sha256.Sum256([]byte(key))

	return gopath.Join(string(f), fmt.Sprintf("%x", path[:]))
}
