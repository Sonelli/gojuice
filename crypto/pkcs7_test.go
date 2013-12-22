package crypto

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPadPKCS7(t *testing.T) {

	tests := []struct {
		input     []byte
		blocksize int
	}{
		{[]byte("SmallTest"), 16},
		{[]byte("This is a test larger than the block size"), 16},
	}

	for _, test := range tests {

		Convey("Given an unpadded input", t, func() {

			padded := PadPKCS7(test.input, test.blocksize)
			Convey("The output should be padded to the correct block size", func() {
				mod := len(padded) % test.blocksize
				So(mod, ShouldEqual, 0)
			})

		})
	}
}
