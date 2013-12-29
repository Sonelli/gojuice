package auth

type PermisionRequiredError struct{}

func (err *PermisionRequiredError) Error() string {
	return "Error: Google OAUTH2 permission not granted"
}
