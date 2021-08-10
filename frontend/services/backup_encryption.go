package services

import (
	"crypto/rand"
	"errors"
	"io"
	"mnimidamonbackend/models"
	"os"
)

// Errors that can occur on encryption or decryption.
var (
	ErrInvalidKey      = errors.New("ErrInvalidKey")
	ErrEncrypting      = errors.New("ErrEncrypting")
	ErrDecrypting      = errors.New("ErrDecrypting")
	ErrCalculatingHash = errors.New("ErrCalculatingHash")
)

// Used for encryption, length has to be 32.
type EncryptionKey []byte

func (key EncryptionKey) isValid() bool {
	if len(key) == 32 {
		return true
	}
	return false
}

func NewRandomEncryptionKey() (EncryptionKey, error) {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}
	return key, nil
}

// Encryption and decryption.
type BackupEncryption interface {
	BackupEncryptor
	BackupDecryptor
}

// Encoder
type BackupDecryptor interface {
	Decrypt(backup *models.Backup, key EncryptionKey, targetFilePath string) error
}

// Decoder
type BackupEncryptor interface {
	Encrypt(backup *models.Backup, key EncryptionKey, fileData io.ReadCloser) (*os.File, error)
}
