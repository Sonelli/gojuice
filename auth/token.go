package auth

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"code.google.com/p/goauth2/oauth"
)

var oauthCfg = &oauth.Config{
	ClientId:       "425124909500-34h1jqmm130mh61chutqn6anaujo054p.apps.googleusercontent.com",
	ClientSecret:   "dIAPPoBeUeOZOWydWjlpSyj0",
	AuthURL:        "https://accounts.google.com/o/oauth2/auth",
	TokenURL:       "https://accounts.google.com/o/oauth2/token",
	RedirectURL:    "http://localhost:25966/oauth2callback",
	Scope:          "https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email",
	AccessType:     "offline",
	ApprovalPrompt: "force",
	TokenCache:     oauth.CacheFile(UserHomeDir() + ".juicessh.auth"),
}

type TokenResponse struct {
	Token *oauth.Token
	Error error
}

var tokenChannel *chan TokenResponse

const port = ":25966"

func StartServer(channel *chan TokenResponse) {

	tokenChannel = channel

	go func() {
		http.HandleFunc("/", handleAuthorize)
		http.HandleFunc("/oauth2callback", handleOAuth2Callback)
		http.ListenAndServe(port, nil)
	}()

}

func GetURL() (url string) {
	url = "http://localhost" + port
	return
}

// Requests a Google OAUTH token from Google.
// If the user has not already authorised us, then launch a local
// HTTP server to retrieve the Google callback and inform the user to launch
// a browser.
func GetToken() (token string, tokenError error) {

	cached, err := oauth.CacheFile(UserHomeDir() + ".juicessh.auth").Token()

	if err != nil || cached == nil {

		channel := make(chan TokenResponse)

		fmt.Printf("\n")
		fmt.Printf("***************************************************\n")
		fmt.Printf(" This utility requires authorisation to access the\n")
		fmt.Printf(" JuiceSSH API via Google OAUTH2 authentication.\n")
		fmt.Printf(" This is a one time operation; OAUTH2 credentials\n")
		fmt.Printf(" will be cached in %s.juicessh.auth\n\n", UserHomeDir())
		fmt.Printf(" Please open the following URL in a browser: \n %s\n", GetURL())
		fmt.Printf("***************************************************\n\n")
		StartServer(&channel)

		response := <-channel

		if response.Error != nil {
			tokenError = response.Error
			return
		}

		token = response.Token.AccessToken
		return

	}

	if cached.Expired() {

		transport := &oauth.Transport{Config: oauthCfg, Token: cached}
		err := transport.Refresh()
		if err != nil {
			tokenError = err
			return
		}

		token = transport.Token.AccessToken
		return

	}

	token = cached.AccessToken
	return

}

// Start the authorization process
func handleAuthorize(w http.ResponseWriter, r *http.Request) {
	url := oauthCfg.AuthCodeURL("")
	http.Redirect(w, r, url, http.StatusFound)
}

// Function that handles the callback from the Google server
func handleOAuth2Callback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")

	transport := &oauth.Transport{Config: oauthCfg}
	tok, err := transport.Exchange(code)

	if err != nil {
		*tokenChannel <- TokenResponse{Error: errors.New("Could not retrieve a valid OAUTH2 token from the Google API")}
	}

	*tokenChannel <- TokenResponse{Token: tok}

	complete := bytes.NewBufferString("<html><body>The JuiceSSH CLI client is now authenticated to use your JuiceSSH account.<br />Please close this browser window and return to the CLI client.<br /></body></html>")

	io.Copy(w, complete)

}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE") + "\\"
		}
		return home
	}
	return os.Getenv("HOME") + "/"
}
