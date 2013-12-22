package crypto

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

func UnpadPKCS7(src []byte) []byte {
	padLen := src[len(src)-1]
	return src[:len(src)-int(padLen)]
}
