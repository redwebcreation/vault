package vault

import (
	"crypto/sha256"
	"fmt"
	"os"
	gopath "path"
)

type Fs string

func (f Fs) Read(key string) ([]byte, error) {
	path := f.URL(key)
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(path)
}

func (f Fs) URL(key string) string {
	path := sha256.Sum256([]byte(key))

	return gopath.Join(string(f), fmt.Sprintf("%x", path[:]))
}

func (f Fs) Write(key string, data []byte) error {
	path := f.URL(key)

	return os.WriteFile(path, data, 0600)
}

func (f Fs) Delete(key string) error {
	return os.Remove(f.URL(key))
}

func (f Fs) Has(key string) bool {
	_, err := os.Stat(f.URL(key))
	return err == nil
}
