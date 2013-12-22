// AES-256 with PKCS#7 Padding
// Key derived via PBKDF2 with HMAC-SHA1

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"strings"

	"code.google.com/p/go.crypto/pbkdf2"
	"github.com/Sonelli/gojuice/errors"
)

const ITERATION_COUNT = 1000
const SALT_LENGTH = 8
const KEY_LENGTH = 32 // 32bytes = 256bits = AES256

func Encrypt(input []byte, passphrase string) (output string, encryptError error) {

	if len(passphrase) < 1 {
		encryptError = &errors.InvalidPassphraseError{"Passphrase length must be greater than zero"}
		return
	}

	salt := make([]byte, SALT_LENGTH)
	readSalt, err := rand.Read(salt)
	if err != nil || readSalt != SALT_LENGTH {
		encryptError = &errors.CouldNotObtainRandomSaltError{err.Error()}
	}

	key := pbkdf2.Key([]byte(passphrase), salt, ITERATION_COUNT, KEY_LENGTH, sha1.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		encryptError = &errors.InvalidAESKeyError{err.Error()}
		return
	}

	iv := make([]byte, block.BlockSize())
	readIV, err := rand.Read(iv)
	if err != nil || readIV != block.BlockSize() {
		encryptError = &errors.CouldNotObtainRandomIVError{err.Error()}
	}

	padded := PadPKCS7(input, block.BlockSize())

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(padded, padded)

	output = base64.StdEncoding.EncodeToString(salt) + "#" + base64.StdEncoding.EncodeToString(iv) + "#" + base64.StdEncoding.EncodeToString(padded)
	return

}

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
		decryptError = &errors.InvalidEncryptedDataError{"Could not derive salt from input"}
		return
	}

	// Note - IV length is always 16 bytes for AES, regardless of key size
	iv, err := base64.StdEncoding.DecodeString(items[1])
	if err != nil {
		decryptError = &errors.InvalidEncryptedDataError{"Could not derive IV from input"}
		return
	}

	cipherText, err := base64.StdEncoding.DecodeString(items[2])
	if err != nil {
		decryptError = &errors.InvalidEncryptedDataError{"Could not derive cipher text from input"}
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
		decryptError = &errors.InvalidAESKeyError{err.Error()}
		return
	}

	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(cipherText, cipherText)

	output = UnpadPKCS7(cipherText)

	return

}
