package vault

import "os"

type Vault struct {
	storage Storage
}

func NewVault(path string) (*Vault, error) {
	_, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	return &Vault{storage: &LocalStorage{
		Path: path,
	}}, nil
}

func (v *Vault) Get(key string, password []byte) ([]byte, error) {
	return v.storage.Get(key, password)
}

func (v *Vault) Set(key string, value []byte, password []byte) error {
	return v.storage.Set(key, value, password)
}

func (v *Vault) Delete(key string) error {
	return v.storage.Delete(key)
}

func (v *Vault) Has(key string) bool {
	return v.storage.Has(key)
}
