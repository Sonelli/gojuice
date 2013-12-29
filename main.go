package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sonelli/gojuice/api"
	"github.com/Sonelli/gojuice/auth"
	"github.com/Sonelli/gojuice/models"
)

func main() {

	if len(os.Args) != 2 {
		showUsage()
		return
	}

	fmt.Printf("\nFetching an OAUTH2 token from Google APIs...\n")
	token, err := auth.GetToken()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Authenticating with JuiceSSH API...\n")
	user, err := auth.Authenticate(token)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Successfully authenticated as %s (%s)\n\n", user.Name, user.Email)

	fmt.Printf("Retrieving CloudSync backup...\n")
	result, err := api.GetCloudSync(user)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Processing CloudSync backup...\n")
	encryptedManifest := &models.EncryptedCloudSyncManifest{}
	json.Unmarshal([]byte(result), &encryptedManifest)

	fmt.Printf("Decrypting CloudSync backup...\n")
	decryptedManifest := encryptedManifest.Decrypt(os.Args[1])

	fmt.Printf("\nDecrypted %d configuration options\n", len(decryptedManifest.Configs))
	fmt.Printf("Decrypted %d connections\n", len(decryptedManifest.Objects.Connections))
	fmt.Printf("Decrypted %d connection groups\n", len(decryptedManifest.Objects.ConnectionGroups))
	fmt.Printf("Decrypted %d connection -> connection group links\n", len(decryptedManifest.Objects.ConnectionGroupMemberships))
	fmt.Printf("Decrypted %d identities\n", len(decryptedManifest.Objects.Identities))
	fmt.Printf("Decrypted %d connection -> identity links\n", len(decryptedManifest.Objects.Identities))
	fmt.Printf("Decrypted %d snippets\n", len(decryptedManifest.Objects.Snippets))
	fmt.Printf("Decrypted %d port forwards\n", len(decryptedManifest.Objects.PortForwards))

	fmt.Printf("\nWriting encrypted JSON to ./cloudsync.encrypted.json\n")
	fmt.Printf("This is your data as the JuiceSSH servers see it (encrypted).\n")
	ioutil.WriteFile("cloudsync.encrypted.json", []byte(result), 0600)

	fmt.Printf("\nWriting decrypted JSON to ./cloudsync.decrypted.json\n")
	fmt.Printf("This is your data decrypted client-side with your passphrase.\n")
	fmt.Printf("The JuiceSSH servers *do not* see this decrypted data ever.\n")
	fmt.Printf("Please do not ever try to push an unencrypted CloudSync file\n")
	fmt.Printf("back to our servers, we do not want to see your unencrypted data!\n\n")
	decryptedJson, _ := json.MarshalIndent(decryptedManifest, "", "  ")
	ioutil.WriteFile("cloudsync.decrypted.json", decryptedJson, 0600)

}

func showUsage() {
	fmt.Printf("Usage: \n")
	fmt.Printf("./gojuice <decryption passphrase>\n")
}
