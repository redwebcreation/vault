package vault

import "errors"

type Storage interface {
	Get(key string, password []byte) ([]byte, error)
	// Set encrypts the value and writes it to the vault file.
	// If the key already exists, it will be overwritten.
	Set(key string, value []byte, password []byte) error
	Has(key string) bool
	// Delete does not check if the key exists before deleting it.
	Delete(key string) error
}

var (
	ErrInvalidVaultFile    = errors.New("invalid vault file")
	ErrKeyDoesNotExist     = errors.New("key does not exist")
	ErrInvalidVaultVersion = errors.New("invalid vault version")
	ErrInvalidHashProtocol = errors.New("invalid hash protocol")
)
