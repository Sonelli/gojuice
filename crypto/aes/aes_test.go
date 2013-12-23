package aes

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func BenchmarkEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encrypt([]byte("Hello"), "password123")
	}
}

func BenchmarkDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decrypt("HIY6P+MIeWs=#GAYJei56awocGBLvmUhSGA==#AmEmLlHNoMZpwTeL1b8vBg==", "password123")
	}
}

func TestDecrypt(t *testing.T) {

	Convey("Given an encrypted input and a passphrase", t, func() {

		input := "HIY6P+MIeWs=#GAYJei56awocGBLvmUhSGA==#AmEmLlHNoMZpwTeL1b8vBg=="
		passphrase := "password123"
		expected := []byte("Hello")

		result, err := Decrypt(input, passphrase)
		if err != nil {
			t.Error(err)
		}

		Convey("The output should be properly decrypted", func() {
			So(string(result[:]), ShouldEqual, string(expected[:]))
		})

	})

	Convey("Given an incorrect passphrase", t, func() {
		_, err := Decrypt("HIY6P+MIeWs=#GAYJei56awocGBLvmUhSGA==#AmEmLlHNoMZpwTeL1b8vBg==", "incorrect-passphrase")
		Convey("An IncorrectPasswordError should be thrown", func() {
			So(err, ShouldHaveSameTypeAs, &IncorrectPassphraseError{})
		})
	})

	Convey("Given an invalid salt", t, func() {
		_, err := Decrypt("invalid-salt#GAYJei56awocGBLvmUhSGA==#AmEmLlHNoMZpwTeL1b8vBg==", "password123")
		Convey("An InvalidSaltError should be thrown", func() {
			So(err, ShouldHaveSameTypeAs, &InvalidSaltError{})
		})
	})

	Convey("Given an invalid IV", t, func() {
		_, err := Decrypt("HIY6P+MIeWs=#invalid-iv#AmEmLlHNoMZpwTeL1b8vBg==", "password123")
		Convey("An InvalidIVError should be thrown", func() {
			So(err, ShouldHaveSameTypeAs, &InvalidIVError{})
		})
	})

	Convey("Given an empty decryption passphrase", t, func() {
		_, err := Decrypt("HIY6P+MIeWs=#GAYJei56awocGBLvmUhSGA==#AmEmLlHNoMZpwTeL1b8vBg==", "")
		Convey("An InvalidPassphraseError should be thrown", func() {
			So(err, ShouldHaveSameTypeAs, &InvalidPassphraseError{})
		})
	})

	Convey("Given an invalid string to decrypt", t, func() {
		_, err := Decrypt("some-invalid-encrypted-string", "some-passphrase")
		Convey("An InvalidEncryptedDataError should be thrown", func() {
			So(err, ShouldHaveSameTypeAs, &InvalidEncryptedDataError{})
		})
	})

}

func TestEncrypt(t *testing.T) {

	Convey("Given an ASCII string to encrypt", t, func() {

		input := []byte("some-string-to-encrypt")
		passphrase := "some-passphrase"

		encrypted, err := Encrypt(input, passphrase)

		Convey("It correctly encrypts", func() {
			So(err, ShouldBeNil)

			decrypted, err := Decrypt(encrypted, passphrase)

			Convey("The resulting encrypted data correctly decrypts", func() {
				So(err, ShouldBeNil)
			})

			Convey("The resulting decrypted data matches the original input", func() {
				So(string(decrypted[:]), ShouldEqual, string(input[:]))
			})

		})

	})

	Convey("Given a UTF-8 string to encrypt", t, func() {

		input := []byte("簥翷臒 礛簼繰 蚙迻 跿 蒛 溛滁溒")
		passphrase := "鶐 貆賌 莦莚虙"

		encrypted, err := Encrypt(input, passphrase)

		Convey("It correctly encrypts", func() {
			So(err, ShouldBeNil)

			decrypted, err := Decrypt(encrypted, passphrase)

			Convey("The resulting encrypted data correctly decrypts", func() {
				So(err, ShouldBeNil)
			})

			Convey("The resulting decrypted data matches the original input", func() {
				So(string(decrypted[:]), ShouldEqual, string(input[:]))
			})

		})

	})

}
