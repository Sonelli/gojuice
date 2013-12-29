package api

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"../auth"
)

// Connect to the JuiceSSH CloudSync API and retrieve the latest CloudSync
// JSON manifest. This JSON contains all of the encrypted records backed up
// sorted by record type.
func GetCloudSync(user *auth.User) (json string, err error) {

	url := "https://api.sonelli.com/cloudsync"
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte("{}")))
	if err != nil {
		return
	}

	req.Header.Add("Cookie", "session="+user.Session.Identifier)
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("Could not retrieve cloudsync backup (HTTP/%d)", resp.StatusCode))
		return
	}

	json = string(body[:])
	return

}
