package pkcs7

type CorruptPaddingError struct{}

func (err *CorruptPaddingError) Error() string {
	return "Error: Invalid padding descriptor"
}
