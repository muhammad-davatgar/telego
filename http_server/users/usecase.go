package users

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func EncryptPassword(username, password string) (string, error) {

	password = CreateHash(password)

	c, err := aes.NewCipher([]byte(password))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	res := hex.EncodeToString(gcm.Seal(nonce, nonce, []byte(username), nil))

	return res, nil
}

func VerifyPassword(username, password, encrypted_pass string) (bool, error) {

	password = CreateHash(password)
	decoding, err := hex.DecodeString(encrypted_pass)
	if err != nil {
		return false, fmt.Errorf("DB Decoding : %w", err)
	}

	encrypted_pass = string(decoding)
	block, err := aes.NewCipher([]byte(password))
	if err != nil {
		return false, fmt.Errorf("NewCipher : %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return false, fmt.Errorf("NewGCM : %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(encrypted_pass) < nonceSize {
		return false, fmt.Errorf("NonceSize : %w", err)
	}

	nonce, encrypted_pass := encrypted_pass[:nonceSize], encrypted_pass[nonceSize:]

	decrypted_username, err := gcm.Open(nil, []byte(nonce), []byte(encrypted_pass), nil)
	if err != nil {
		return false, err
	}

	if bytes.Equal(decrypted_username, []byte(username)) {
		return true, nil
	}
	return false, nil
}
