// AES-256 with PKCS#7 Padding
// Key derived via PBKDF2 with HMAC-SHA1

package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"strings"

	"github.com/Sonelli/gojuice/crypto/pkcs7"

	"code.google.com/p/go.crypto/pbkdf2"
)

const ITERATION_COUNT = 1000
const SALT_LENGTH = 8
const KEY_LENGTH = 32 // 32bytes = 256bits = AES256

func Encrypt(input []byte, passphrase string) (output string, encryptError error) {

	if len(passphrase) < 1 {
		encryptError = &InvalidPassphraseError{"Passphrase length must be greater than zero"}
		return
	}

	salt := make([]byte, SALT_LENGTH)
	readSalt, err := rand.Read(salt)
	if err != nil || readSalt != SALT_LENGTH {
		encryptError = &CouldNotObtainRandomSaltError{err.Error()}
	}

	key := pbkdf2.Key([]byte(passphrase), salt, ITERATION_COUNT, KEY_LENGTH, sha1.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		encryptError = &InvalidAESKeyError{err.Error()}
		return
	}

	iv := make([]byte, block.BlockSize())
	readIV, err := rand.Read(iv)
	if err != nil || readIV != block.BlockSize() {
		encryptError = &CouldNotObtainRandomIVError{err.Error()}
	}

	padded := pkcs7.PadPKCS7(input, block.BlockSize())

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(padded, padded)

	output = base64.StdEncoding.EncodeToString(salt) + "#" + base64.StdEncoding.EncodeToString(iv) + "#" + base64.StdEncoding.EncodeToString(padded)
	return

}

func Decrypt(input, passphrase string) (output []byte, decryptError error) {

	if len(passphrase) < 1 {
		decryptError = &InvalidPassphraseError{"Passphrase length must be greater than zero"}
		return
	}

	items := strings.Split(input, "#")
	if len(items) != 3 {
		decryptError = &InvalidEncryptedDataError{"Encrypted data must contain salt, iv and data"}
		return
	}

	salt, err := base64.StdEncoding.DecodeString(items[0])
	if err != nil || len(salt) != SALT_LENGTH {
		decryptError = &InvalidSaltError{"Could not derive salt from input"}
		return
	}

	inputBytes, err := base64.StdEncoding.DecodeString(items[2])
	if err != nil {
		decryptError = &InvalidEncryptedDataError{"Could not derive cipher text from input"}
		return
	}

	// CBC mode always works in whole blocks.
	if len(inputBytes)%aes.BlockSize != 0 {
		decryptError = &InvalidEncryptedDataError{"Cipher text is not a multiple of AES block size"}
		return
	}

	key := pbkdf2.Key([]byte(passphrase), salt, ITERATION_COUNT, KEY_LENGTH, sha1.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		decryptError = &InvalidAESKeyError{err.Error()}
		return
	}

	// Note - IV length is always 16 bytes for AES, regardless of key size
	iv, err := base64.StdEncoding.DecodeString(items[1])
	if err != nil || len(iv) != block.BlockSize() {
		decryptError = &InvalidIVError{"Could not derive IV from input"}
		return
	}

	cbc := cipher.NewCBCDecrypter(block, iv)
	cipherText := make([]byte, len(inputBytes))
	copy(cipherText, inputBytes)
	cbc.CryptBlocks(cipherText, cipherText)

	output, err = pkcs7.UnpadPKCS7(cipherText, block.BlockSize())
	if err != nil {
		decryptError = &IncorrectPassphraseError{}
		return
	}

	return

}
