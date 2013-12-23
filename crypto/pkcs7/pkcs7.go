package pkcs7

// Returns a new byte array padded with PKCS7

func PadPKCS7(src []byte, blockSize int) []byte {
	missing := blockSize - (len(src) % blockSize)
	newSize := len(src) + missing
	dest := make([]byte, newSize, newSize)
	// copy data
	for i := 0; i < len(src); i++ {
		dest[i] = src[i]
	}
	// fill in the rest
	missingB := byte(missing)
	for i := newSize - missing; i < newSize; i++ {
		dest[i] = missingB
	}
	return dest
}

func UnpadPKCS7(src []byte, blockSize int) (output []byte, err error) {
	var paddingLength int
	paddingLength = int(src[len(src)-1])

	if len(src)%blockSize != 0 {
		err = &CorruptPaddingError{}
		return
	}

	if paddingLength > blockSize {
		err = &CorruptPaddingError{}
		return
	}

	for _, v := range src[len(src)-paddingLength:] {
		if int(v) != paddingLength {
			err = &CorruptPaddingError{}
			return
		}
	}

	if paddingLength == 0 {
		paddingLength = 1
	}

	output = src[:len(src)-int(paddingLength)]
	return

}
