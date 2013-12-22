package aes

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

type CouldNotObtainRandomSaltError struct {
	Msg string
}

func (err *CouldNotObtainRandomSaltError) Error() string {
	return "Could not obtain random salt: " + err.Msg
}

type CouldNotObtainRandomIVError struct {
	Msg string
}

func (err *CouldNotObtainRandomIVError) Error() string {
	return "Could not obtain random IV: " + err.Msg
}

type InvalidAESKeyError struct {
	Msg string
}

func (err *InvalidAESKeyError) Error() string {
	return "Invalid AES key: " + err.Msg
}

type InvalidIVError struct {
	Msg string
}

func (err *InvalidIVError) Error() string {
	return "Invalid IV: " + err.Msg
}

type InvalidSaltError struct {
	Msg string
}

func (err *InvalidSaltError) Error() string {
	return "Invalid Salt: " + err.Msg
}

type IncorrectPassphraseError struct{}

func (err *IncorrectPassphraseError) Error() string {
	return "Invalid decryption passphrase"
}
