package crypto

import (
	"testing"

	"github.com/Sonelli/gojuice/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDecrypt(t *testing.T) {

	tests := []struct {
		input      string
		passphrase string
		output     []byte
	}{
		{
			"HIY6P+MIeWs=#GAYJei56awocGBLvmUhSGA==#AmEmLlHNoMZpwTeL1b8vBg==",
			"password123",
			[]byte("Hello"),
		},
	}

	for _, test := range tests {

		Convey("Given an encrypted input and a passphrase", t, func() {

			output, err := Decrypt(test.input, test.passphrase)
			if err != nil {
				t.Error(err)
			}

			Convey("The output should be properly decrypted", func() {
				So(string(output[:]), ShouldEqual, string(test.output[:]))
			})

		})
	}

	Convey("Given an empty decryption passphrase", t, func() {
		_, err := Decrypt("something-to-decrypt", "")
		Convey("An InvalidPassphraseError should be thrown", func() {
			So(err, ShouldHaveSameTypeAs, &errors.InvalidPassphraseError{})
		})
	})

	Convey("Given an invalid string to decrypt", t, func() {
		_, err := Decrypt("some-invalid-encrypted-string", "some-passphrase")
		Convey("An InvalidEncryptedDataError should be thrown", func() {
			So(err, ShouldHaveSameTypeAs, &errors.InvalidEncryptedDataError{})
		})
	})

}

func TestEncrypt(t *testing.T) {

	tests := []struct {
		input      []byte
		passphrase string
	}{
		{[]byte("some-string-to-encrypt"), "some-passphrase"},
	}

	for _, test := range tests {

		Convey("Given a plain text string to encrypt", t, func() {

			encrypted, err := Encrypt(test.input, test.passphrase)

			Convey("It correctly encrypts", func() {
				So(err, ShouldBeNil)

				decrypted, err := Decrypt(encrypted, test.passphrase)

				Convey("The resulting encrypted data correctly decrypts", func() {
					So(err, ShouldBeNil)
				})

				Convey("The resulting decrypted data matches the original input", func() {
					So(string(decrypted[:]), ShouldEqual, string(test.input[:]))
				})

			})

		})

	}

}
