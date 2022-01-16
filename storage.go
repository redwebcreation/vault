package vault

import "errors"

type Storage interface {
	Read(key string) ([]byte, error)
	Write(key string, data []byte) error
	Has(key string) bool
	Delete(key string) error
}

var (
	ErrInvalidVaultFile    = errors.New("invalid vault file")
	ErrInvalidVaultVersion = errors.New("invalid vault version")
	ErrInvalidCipher       = errors.New("invalid hash protocol")
)
