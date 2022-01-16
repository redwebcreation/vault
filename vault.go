package vault

import (
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
)

type Vault struct {
	Storage Storage
}

func NewVault(path string) (*Vault, error) {
	_, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	return &Vault{
		Storage: Fs(path),
	}, nil
}

func (v *Vault) Get(key string, password []byte) ([]byte, error) {
	bytes, err := v.Storage.Read(key)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`(?m)^\$VAULT;(?P<version>[0-9.]+);(?P<hash>[a-zA-Z1-9]+)\n(?P<contents>[0-9a-f]+)$`)
	matches := re.FindAllStringSubmatch(string(bytes), -1)

	if len(matches) != 1 {
		return nil, ErrInvalidVaultFile
	}

	vaultVersion, hash, contents := matches[0][1], matches[0][2], matches[0][3]

	if vaultVersion != "1.0" {
		return nil, ErrInvalidVaultVersion
	}

	if hash != "AES256" {
		return nil, ErrInvalidCipher
	}

	decoded, err := hex.DecodeString(contents)
	if err != nil {
		return nil, err
	}

	return Decrypt(password, decoded)
}

func (v *Vault) Set(key string, value []byte, password []byte) error {
	encrypted, err := Encrypt(password, value)
	if err != nil {
		return err
	}

	contents := fmt.Sprintf("$VAULT;1.0;AES256\n%x", encrypted)

	return v.Storage.Write(
		key,
		[]byte(contents),
	)
}

func (v *Vault) Delete(key string) error {
	return v.Storage.Delete(key)
}

func (v *Vault) Has(key string) bool {
	return v.Storage.Has(key)
}
