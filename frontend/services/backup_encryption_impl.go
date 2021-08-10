package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mnimidamonbackend/models"
	"os"
	"path/filepath"
)

var BackupCryptography *backupEncryptionImpl
func init() {
	BackupCryptography = &backupEncryptionImpl{}
}

type backupEncryptionImpl struct {}


func (be *backupEncryptionImpl) Encrypt(payload *models.InitializeGroupBackupPayload, key EncryptionKey, fileData io.ReadCloser, encryptedBytes *int) (*os.File, error) {
	defer fileData.Close()
	*encryptedBytes = 0

	if !key.isValid() {
		return nil, fmt.Errorf("%w", ErrInvalidKey)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("%w %v", ErrEncrypting, err)
	}

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("%w %v", ErrEncrypting, err)
	}

	outFile, err := os.OpenFile(BackupStorage.GetTempPath(*payload.FileName), os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return nil, fmt.Errorf("%w openening file %v", ErrEncrypting, err)
	}

	buf := make([]byte, 1024)
	stream := cipher.NewCTR(block, iv)
	for {
		n, err := fileData.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			// Write into file
			_, err := outFile.Write(buf[:n])
			if err != nil {
				return nil, fmt.Errorf("%w when writting to out file %v", ErrEncrypting, err)
			}

			// Update the process.
			*encryptedBytes += n
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("%w when reading %v", ErrEncrypting, err)
		}
	}

	// Append the IV.
	_, err = outFile.Write(iv)
	if err != nil {
		return nil, fmt.Errorf("%w when writting IV to out file %v", ErrEncrypting, err)
	}

	// Reset the reader.
	_, err = outFile.Seek(0, 0)
	if err != nil {
		return nil, fmt.Errorf("%w when seeking %v", ErrEncrypting, err)
	}

	// Calculate the hash.
	hash, err := calculateReaderHash(outFile)
	if err != nil {
		return nil, fmt.Errorf("%w when hashing %v", ErrEncrypting, err)
	}

	// Save the hash.
	payload.Hash = &hash

	// Reset the reader.
	_, err = outFile.Seek(0, 0)
	if err != nil {
		return nil, fmt.Errorf("%w when seeking %v", ErrEncrypting, err)
	}

	// File stats to get the size.
	fi, err := outFile.Stat()
	if err != nil {
		return nil, fmt.Errorf("%w reading stats %v", ErrEncrypting, err)
	}

	// Convert to kilobytes
	size := fi.Size() / 1024
	payload.Size = &size

	return outFile, nil
}

func (be *backupEncryptionImpl) Decrypt(backup *models.Backup, key EncryptionKey, targetFilePath string) error {
	// Open the decrypted contents of the file.
	infile, err := os.Open(BackupStorage.GetBackupPath(int(backup.BackupID)))
	if err != nil {
		return fmt.Errorf("%w opening file %v", ErrDecrypting, err)
	}
	defer infile.Close()

	// New cipher from the key.
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("%w creating cipher %v", ErrDecrypting, err)
	}

	fi, err := infile.Stat()
	if err != nil {
		return fmt.Errorf("%w reading file info %v", ErrDecrypting, err)
	}

	iv := make([]byte, block.BlockSize())
	msgLen := fi.Size() - int64(len(iv))
	_, err = infile.ReadAt(iv, msgLen)
	if err != nil {
		return fmt.Errorf("%w getting the initializing vector %v", ErrDecrypting, err)
	}

	outfile, err := os.OpenFile(filepath.Join(targetFilePath, backup.Filename), os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return fmt.Errorf("%w opening output file %v", ErrDecrypting, err)
	}
	defer outfile.Close()

	buf := make([]byte, 1024)
	stream := cipher.NewCTR(block, iv)
	for {
		n, err := infile.Read(buf)
		if n > 0 {
			// The last bytes are the IV, don't belong the original message
			if n > int(msgLen) {
				n = int(msgLen)
			}
			msgLen -= int64(n)
			stream.XORKeyStream(buf, buf[:n])
			// Write into file
			_, err := outfile.Write(buf[:n])
			if err != nil {
				return fmt.Errorf("%w when writting to out file %v", ErrDecrypting, err)
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("%w when reading infile %v", ErrDecrypting, err)
		}
	}

	return nil
}

func calculateReaderHash(rc io.Reader) (string, error) {
	h := sha256.New()

	if _, err := io.Copy(h, rc); err != nil {
		return "", fmt.Errorf("%w: error calculating hash of backup %v: %v", ErrCalculatingHash, err)
	}

	calculatedHash := hex.EncodeToString(h.Sum(nil))
	return calculatedHash, nil
}