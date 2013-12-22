package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"strings"

	"code.google.com/p/go.crypto/pbkdf2"
	"github.com/Sonelli/gojuice/errors"
)

const ITERATION_COUNT = 1000
const SALT_LENGTH = 8
const KEY_LENGTH = 32 // 32bytes = 256bits = AES256

func Decrypt(input, passphrase string) (output []byte, decryptError error) {

	if len(passphrase) < 1 {
		decryptError = &errors.InvalidPassphraseError{"Passphrase length must be greater than zero"}
		return
	}

	items := strings.Split(input, "#")
	if len(items) != 3 {
		decryptError = &errors.InvalidEncryptedDataError{"Encrypted data must contain salt, iv and data"}
		return
	}

	salt, err := base64.StdEncoding.DecodeString(items[0])
	if err != nil {
		decryptError = &errors.InvalidEncryptedDataError{"Could not derive salt"}
		return
	}

	// Note - IV length is always 16 bytes for AES, regardless of key size
	iv, err := base64.StdEncoding.DecodeString(items[1])
	if err != nil {
		decryptError = &errors.InvalidEncryptedDataError{"Could not derive IV"}
		return
	}

	cipherText, err := base64.StdEncoding.DecodeString(items[2])
	if err != nil {
		decryptError = &errors.InvalidEncryptedDataError{"Could not derive cipher text"}
		return
	}

	// CBC mode always works in whole blocks.
	if len(cipherText)%aes.BlockSize != 0 {
		decryptError = &errors.InvalidEncryptedDataError{"Cipher text is not a multiple of AES block size"}
		return
	}

	key := pbkdf2.Key([]byte(passphrase), salt, ITERATION_COUNT, KEY_LENGTH, sha1.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		decryptError = &errors.InvalidEncryptedDataError{err.Error()}
		return
	}

	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(cipherText, cipherText)

	output = UnpadPKCS7(cipherText)

	return

}
