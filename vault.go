package vault

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"regexp"
)

// ErrInvalidVaultFile is returned when the vault file is invalid
var ErrInvalidVaultFile = errors.New("invalid vault file")

var fileRegex = regexp.MustCompile(`(?m)^\$VAULT;1.0;AES256\n(?P<contents>[0-9a-f]+)$`)

// Vault is a simple key-value store with an interoperable backend.
type Vault struct {
	Storage Storage
}

// NewVault creates a new vault using a Fs as the storage backend.
func NewVault(path string) (*Vault, error) {
	_, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	return &Vault{
		Storage: Fs(path),
	}, nil
}

// Get retrieves the value of a key from the vault.
func (v *Vault) Get(key string, password []byte) ([]byte, error) {
	bytes, err := v.Storage.Read(key)
	if err != nil {
		return nil, err
	}

	matches := fileRegex.FindAllStringSubmatch(string(bytes), -1)

	if len(matches) != 1 {
		return nil, ErrInvalidVaultFile
	}

	decoded, err := hex.DecodeString(matches[0][1])
	if err != nil {
		return nil, err
	}

	return decrypt(password, decoded)
}

// Set stores a value in the vault.
func (v *Vault) Set(key string, value []byte, password []byte) error {
	encrypted, err := encrypt(password, value)
	if err != nil {
		return err
	}

	contents := fmt.Sprintf("$VAULT;1.0;AES256\n%x", encrypted)

	return v.Storage.Write(
		key,
		[]byte(contents),
	)
}

// Delete removes a key from the vault.
func (v *Vault) Delete(key string) error {
	return v.Storage.Delete(key)
}

// Has returns true if the vault contains the given key.
func (v *Vault) Has(key string) bool {
	return v.Storage.Has(key)
}
