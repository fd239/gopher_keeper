package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"github.com/fd239/gopher_keeper/config"
	"log"
)

// CipherCrypt struct for Decrypting
type CipherCrypt struct {
	nonce  []byte
	aesGCM cipher.AEAD
}

//Decrypt decrypts some string
func (c *CipherCrypt) Decrypt(entity string) (string, error) {
	b, err := hex.DecodeString(entity)
	if err != nil {
		log.Printf("Decrypt decode string error: %v", err)
		return "", err
	}

	decrypted, terr := c.aesGCM.Open(nil, c.nonce, b, nil)

	if terr != nil {
		log.Printf("aesGCM open error: %v", terr)
	}

	return string(decrypted), err

}

//Encrypt encrypts some string
func (c *CipherCrypt) Encrypt(entity string) (string, error) {
	encrypted := c.aesGCM.Seal(nil, c.nonce, []byte(entity), nil)
	return hex.EncodeToString(encrypted), nil
}

//NewCrypt returns crypt functionality
func NewCrypt(cfg *config.Config) (*CipherCrypt, error) {
	secretKey := []byte(cfg.Keeper.CryptSecret)
	aesblock, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	nonce := secretKey[len(secretKey)-aesgcm.NonceSize():]

	return &CipherCrypt{
		aesGCM: aesgcm,
		nonce:  nonce,
	}, nil

}
