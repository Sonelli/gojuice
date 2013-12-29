package models

import (
	"encoding/json"
	"reflect"

	"github.com/Sonelli/gojuice/crypto/aes"
)

type EncryptedCloudSyncManifest struct {
	Configs []ConfigItem        `json:"configs"`
	Objects EncryptedObjectList `json:"objects"`
}

func NewEncryptedCloudSyncManifest() *EncryptedCloudSyncManifest {
	return &EncryptedCloudSyncManifest{
		Configs: make([]ConfigItem, 0),
		Objects: *NewEncryptedObjectList(),
	}
}

// Decrypt an encrypted CloudSync manifest using the passphrase provided
func (e *EncryptedCloudSyncManifest) Decrypt(passphrase string) *DecryptedCloudSyncManifest {

	// Decrypt the configuration items
	var decryptedManifest *DecryptedCloudSyncManifest = NewDecryptedCloudSyncManifest()
	for _, value := range e.Configs {

		decrypted, err := aes.Decrypt(value.Value, passphrase)
		if err != nil {
			continue
		}

		value.Value = string(decrypted[:])
		decryptedManifest.Configs = append(decryptedManifest.Configs, value)
	}

	for _, object := range e.Objects.Connections {
		var item *Connection = &Connection{}
		MapEncryptedObject(passphrase, &object, item)
		if *item != (Connection{}) {
			decryptedManifest.Objects.Connections = append(decryptedManifest.Objects.Connections, *item)
		}
	}

	for _, object := range e.Objects.ConnectionGroups {
		var item *ConnectionGroup = &ConnectionGroup{}
		MapEncryptedObject(passphrase, &object, item)
		if *item != (ConnectionGroup{}) {
			decryptedManifest.Objects.ConnectionGroups = append(decryptedManifest.Objects.ConnectionGroups, *item)
		}
	}

	for _, object := range e.Objects.ConnectionGroupMemberships {
		var item *ConnectionGroupMembership = &ConnectionGroupMembership{}
		MapEncryptedObject(passphrase, &object, item)
		if *item != (ConnectionGroupMembership{}) {
			decryptedManifest.Objects.ConnectionGroupMemberships = append(decryptedManifest.Objects.ConnectionGroupMemberships, *item)
		}
	}

	for _, object := range e.Objects.Snippets {
		var item *Snippet = &Snippet{}
		MapEncryptedObject(passphrase, &object, item)
		if *item != (Snippet{}) {
			decryptedManifest.Objects.Snippets = append(decryptedManifest.Objects.Snippets, *item)
		}
	}

	for _, object := range e.Objects.Identities {
		var item *Identity = &Identity{}
		MapEncryptedObject(passphrase, &object, item)

		// The password and/or private key is encrypted again within
		// the decrypted identity record, so decrypt those too
		password, _ := aes.Decrypt(item.Password, passphrase)
		item.Password = string(password[:])

		privateKey, _ := aes.Decrypt(item.PrivateKey, passphrase)
		item.PrivateKey = string(privateKey[:])

		privateKeyPassword, _ := aes.Decrypt(item.PrivateKeyPassword, passphrase)
		item.PrivateKeyPassword = string(privateKeyPassword[:])

		if *item != (Identity{}) {
			decryptedManifest.Objects.Identities = append(decryptedManifest.Objects.Identities, *item)
		}

	}

	for _, object := range e.Objects.ConnectionIdentities {
		var item *ConnectionIdentity = &ConnectionIdentity{}
		MapEncryptedObject(passphrase, &object, item)
		if *item != (ConnectionIdentity{}) {
			decryptedManifest.Objects.ConnectionIdentities = append(decryptedManifest.Objects.ConnectionIdentities, *item)
		}
	}

	for _, object := range e.Objects.PortForwards {
		var item *PortForward = &PortForward{}
		MapEncryptedObject(passphrase, &object, item)
		if *item != (PortForward{}) {
			decryptedManifest.Objects.PortForwards = append(decryptedManifest.Objects.PortForwards, *item)
		}
	}

	return decryptedManifest

}

func MapEncryptedObject(passphrase string, input *EncryptedItem, output interface{}) {

	if len(input.Data) > 0 {
		decrypted, err := aes.Decrypt(input.Data, passphrase)
		if err != nil {
			return
		}
		json.Unmarshal(decrypted, output)
	}

	// Set the standard base fields on the output
	for i := 0; i < reflect.TypeOf(input).Elem().NumField(); i++ {
		inputField := reflect.TypeOf(output).Elem().FieldByIndex([]int{i})

		outputField, found := reflect.TypeOf(output).Elem().FieldByName(inputField.Name)
		if found {

			inputFieldValue := reflect.ValueOf(input).Elem().FieldByName(inputField.Name)
			outputFieldValue := reflect.ValueOf(output).Elem().FieldByName(outputField.Name)

			if inputFieldValue.IsValid() && outputFieldValue.IsValid() && outputFieldValue.CanSet() {
				outputFieldValue.Set(inputFieldValue)
			}

		}

	}

}

type DecryptedCloudSyncManifest struct {
	Configs []ConfigItem        `json:"configs"`
	Objects DecryptedObjectList `json:"objects"`
}

func NewDecryptedCloudSyncManifest() *DecryptedCloudSyncManifest {
	return &DecryptedCloudSyncManifest{
		Configs: make([]ConfigItem, 0),
		Objects: *NewDecryptedObjectList(),
	}
}
