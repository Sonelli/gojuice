package pkcs7

import (
	"crypto/rand"
	"fmt"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPadPKCS7(t *testing.T) {

	tests := []struct {
		input     []byte
		blocksize int
		output    int
	}{
		{[]byte(""), 16, 16},
		{[]byte("SmallTest"), 16, 16},
		{[]byte("This is a test larger than the block size"), 16, 48},
	}

	for _, test := range tests {

		Convey(fmt.Sprintf("Given an unpadded input of length %d", len(test.input)), t, func() {

			padded := PadPKCS7(test.input, test.blocksize)

			Convey("The output should be padded to a valid block size", func() {
				mod := len(padded) % test.blocksize
				So(mod, ShouldEqual, 0)
			})

			Convey(fmt.Sprintf("The output should be padded to the correct block size (%d)", test.output), func() {
				So(len(padded), ShouldEqual, test.output)
			})

		})
	}

}

func TestUnpadPKCS7(t *testing.T) {
	Convey("Given an corrupt padding input", t, func() {

		corrupt := make([]byte, 1024)
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
