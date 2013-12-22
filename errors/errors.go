package errors

type InvalidPassphraseError struct {
	Msg string
}

func (err *InvalidPassphraseError) Error() string {
	return "Invalid encryption or decryption passphrase: " + err.Msg
}

type InvalidEncryptedDataError struct {
	Msg string
}

func (err *InvalidEncryptedDataError) Error() string {
	return "Invalid encrypted data: " + err.Msg
}
