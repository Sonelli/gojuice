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
