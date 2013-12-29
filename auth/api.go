package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Purchases []Purchase `json:"purchases"`
	Disabled  bool       `json:"disabled"`
	Session   Session    `json:"session"`
	Signature string     `json:"signature"`
}

type Purchase struct {
	Id      string `json:"_id"`
	Time    int64  `json:"time"`
	Order   string `json:"order"`
	Product string `json:"product"`
	State   int    `json:"state"`
}

type Session struct {
	Identifier string `json:"identifier"`
	Expires    int    `json:"expires"`
}

// Authenticates with the JuiceSSH API using a previously obtained
// Google OAUTH2 token. The returned User object contains a session
// identifer that can be used for future API calls via a HTTP cookie header:
// Cookie: session=<identifier>
func Authenticate(token string) (*User, error) {

	url := "https://api.sonelli.com/oauth/" + token

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	user := &User{}
	json.Unmarshal(body, &user)

	return user, nil

}
