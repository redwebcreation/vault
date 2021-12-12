package vault

import (
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

type LocalStorage struct {
	Path string
}

func (l *LocalStorage) pathFor(key string) string {
	return l.Path + "/" + strings.ReplaceAll(key, ".", "/")
}

func (l *LocalStorage) read(key string) ([]byte, error) {
	keyPath := l.pathFor(key)
	_, err := os.Stat(keyPath)

	if os.IsNotExist(err) {
		return nil, ErrKeyDoesNotExist
	}

	if err != nil {
		return nil, err
	}

	return os.ReadFile(keyPath)
}

func (l *LocalStorage) Get(key string, password []byte) ([]byte, error) {
	bytes, err := l.read(key)
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
		return nil, ErrInvalidHashProtocol
	}

	decoded, err := hex.DecodeString(contents)
	if err != nil {
		return nil, err
	}

	return decrypt(password, decoded)
}

func (l *LocalStorage) Has(key string) bool {
	_, err := l.read(key)

	return err == nil
}

func (l *LocalStorage) Set(key string, value []byte, password []byte) error {
	encrypted, err := encrypt(password, value)
	if err != nil {
		return err
	}

	contents := fmt.Sprintf("$VAULT;1.0;AES256\n%x", encrypted)

	keyPath := l.pathFor(key)

	if strings.Count(key, ".") > 0 {
		parent, _ := path.Split(keyPath)

		err = os.MkdirAll(parent, 0640)
		if err != nil {
			return err
		}
	}

	return os.WriteFile(
		keyPath,
		[]byte(contents),
		0640,
	)
}

func (l *LocalStorage) Delete(key string) error {
	keyPath := l.pathFor(key)

	_ = os.Remove(keyPath)

	return nil
}
