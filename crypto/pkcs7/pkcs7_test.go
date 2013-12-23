package pkcs7

import (
	"crypto/rand"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPadPKCS7(t *testing.T) {

	blockSize := 16

	Convey("Given an empty input", t, func() {

		padded := PadPKCS7(make([]byte, 0), blockSize)

		Convey("The output should be padded to a valid block size", func() {
			mod := len(padded) % blockSize
			So(mod, ShouldEqual, 0)
		})

		Convey(fmt.Sprintf("The output should be padded to the correct block size (%d)", blockSize), func() {
			So(len(padded), ShouldEqual, blockSize)
		})

	})

	Convey("Given an unpadded input smaller than the block size", t, func() {

		padded := PadPKCS7(make([]byte, 8), blockSize)

		Convey("The output should be padded to a valid block size", func() {
			mod := len(padded) % blockSize
			So(mod, ShouldEqual, 0)
		})

		Convey(fmt.Sprintf("The output should be padded to the correct block size (%d)", blockSize), func() {
			So(len(padded), ShouldEqual, blockSize)
		})

	})

	Convey("Given an unpadded input equal to the block size", t, func() {

		padded := PadPKCS7(make([]byte, 15), blockSize)

		Convey("The output should be padded to a valid block size", func() {
			mod := len(padded) % blockSize
			So(mod, ShouldEqual, 0)
		})

		Convey(fmt.Sprintf("The output should be padded to the correct block size (%d)", blockSize), func() {
			So(len(padded), ShouldEqual, blockSize)
		})

	})

	Convey("Given an unpadded input bigger than the block size", t, func() {

		padded := PadPKCS7(make([]byte, 18), blockSize)

		Convey("The output should be padded to a valid block size", func() {
			mod := len(padded) % blockSize
			So(mod, ShouldEqual, 0)
		})

		Convey(fmt.Sprintf("The output should be padded to the correct block size (%d)", blockSize*2), func() {
			So(len(padded), ShouldEqual, blockSize*2)
		})

	})

}

func TestUnpadPKCS7(t *testing.T) {

	Convey("Given an valid padding input spanning one block", t, func() {
		corrupt := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 6, 6, 6, 6, 6, 6}
		Convey("An correctly unpadded result should be returned", func() {
			result, _ := UnpadPKCS7(corrupt, 16)
			So(len(result), ShouldEqual, 10)
		})
	})

	Convey("Given an valid padding input spanning multiple blocks", t, func() {
		corrupt := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 6, 6, 6, 6, 6, 6}
		Convey("An correctly unpadded result should be returned", func() {
			result, _ := UnpadPKCS7(corrupt, 16)
			So(len(result), ShouldEqual, 26)
		})
	})

	Convey("Given an invalid padding input", t, func() {
		corrupt := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 33}
		Convey("An CorruptPaddingError should be thrown", func() {
			_, err := UnpadPKCS7(corrupt, 16)
			So(err, ShouldHaveSameTypeAs, &CorruptPaddingError{})
		})
	})

	Convey("Given an input with no padding included", t, func() {
		corrupt := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0}
		Convey("An correctly unpadded result should be returned", func() {
			result, _ := UnpadPKCS7(corrupt, 16)
			So(len(result), ShouldEqual, 15)
		})
	})

	Convey("Given an corrupt padding input", t, func() {

		corrupt := make([]byte, 32)
		read, err := rand.Read(corrupt)
		if err != nil || read != len(corrupt) {
			t.Fatalf("Could not generate corrupt padded block from random source: %s", err)
		}

		Convey("An CorruptPaddingError should be thrown", func() {
			_, err := UnpadPKCS7(corrupt, 16)
			So(err, ShouldHaveSameTypeAs, &CorruptPaddingError{})
		})

	})
}
