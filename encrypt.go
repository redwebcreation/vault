package vault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

func encrypt(password, value []byte) ([]byte, error) {
	block, err := deriveBlockCypherFromPassword(password)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, block.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return block.Seal(nonce, nonce, value, nil), nil
}

func decrypt(password, value []byte) ([]byte, error) {
	block, err := deriveBlockCypherFromPassword(password)
	if err != nil {
		return nil, err
	}

	nonceSize := block.NonceSize()
	nonce, encrypted := value[:nonceSize], value[nonceSize:]

	return block.Open(nil, nonce, encrypted, nil)
}

func deriveBlockCypherFromPassword(password []byte) (cipher.AEAD, error) {
	key := sha256.Sum256(password)
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	return cipher.NewGCM(block)
}
