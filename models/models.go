package models

type DecryptedObjectList struct {
	Connections                []Connection                `json:"com.sonelli.juicessh.models.Connection"`
	ConnectionGroups           []ConnectionGroup           `json:"com.sonelli.juicessh.models.ConnectionGroup"`
	ConnectionGroupMemberships []ConnectionGroupMembership `json:"com.sonelli.juicessh.models.ConnectionGroupMembership"`
	Snippets                   []Snippet                   `json:"com.sonelli.juicessh.models.Snippet"`
	Identities                 []Identity                  `json:"com.sonelli.juicessh.models.Identity"`
	ConnectionIdentities       []ConnectionIdentity        `json:"com.sonelli.juicessh.models.ConnectionIdentity"`
	PortForwards               []PortForward               `json:"com.sonelli.juicessh.models.PortForward"`
}

func NewDecryptedObjectList() *DecryptedObjectList {
	return &DecryptedObjectList{
		Connections:                make([]Connection, 0),
		ConnectionGroups:           make([]ConnectionGroup, 0),
		ConnectionGroupMemberships: make([]ConnectionGroupMembership, 0),
		Snippets:                   make([]Snippet, 0),
		Identities:                 make([]Identity, 0),
		ConnectionIdentities:       make([]ConnectionIdentity, 0),
		PortForwards:               make([]PortForward, 0),
	}
}

type EncryptedObjectList struct {
	Connections                []EncryptedItem `json:"com.sonelli.juicessh.models.Connection"`
	ConnectionGroups           []EncryptedItem `json:"com.sonelli.juicessh.models.ConnectionGroup"`
	ConnectionGroupMemberships []EncryptedItem `json:"com.sonelli.juicessh.models.ConnectionGroupMembership"`
	Snippets                   []EncryptedItem `json:"com.sonelli.juicessh.models.Snippet"`
	Identities                 []EncryptedItem `json:"com.sonelli.juicessh.models.Identity"`
	ConnectionIdentities       []EncryptedItem `json:"com.sonelli.juicessh.models.ConnectionIdentity"`
	PortForwards               []EncryptedItem `json:"com.sonelli.juicessh.models.PortForward"`
}

func NewEncryptedObjectList() *EncryptedObjectList {
	return &EncryptedObjectList{
		Connections:                make([]EncryptedItem, 0),
		ConnectionGroups:           make([]EncryptedItem, 0),
		ConnectionGroupMemberships: make([]EncryptedItem, 0),
		Snippets:                   make([]EncryptedItem, 0),
		Identities:                 make([]EncryptedItem, 0),
		ConnectionIdentities:       make([]EncryptedItem, 0),
		PortForwards:               make([]EncryptedItem, 0),
	}
}

type EncryptedItem struct {
	Id            string `json:"_id"`
	Owner         string `json:"owner"`
	Modified      int64  `json:"modified"`
	LastUpdatedBy string `json:"lastUpdatedBy"`
	Data          string `json:"data"`
}

type Snippet struct {
	Id            string `json:"_id"`
	Owner         string `json:"owner"`
	Modified      int64  `json:"modified"`
	LastUpdatedBy string `json:"lastUpdatedBy"`
	Name          string `json:"name"`
	Content       string `json:"content"`
}

type Identity struct {
	Id                 string `json:"_id"`
	Owner              string `json:"owner"`
	Modified           int64  `json:"modified"`
	LastUpdatedBy      string `json:"lastUpdatedBy"`
	Nickname           string `json:"nickname"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	PrivateKey         string `json:"privateKey"`
	PrivateKeyPassword string `json:"privateKeyPassword"`
	PassUpdated        int64  `json:"passUpdated"`
	KeyUpdated         int64  `json:"keyUpdated"`
}

type ConnectionIdentity struct {
	Id            string `json:"_id"`
	Owner         string `json:"owner"`
	Modified      int64  `json:"modified"`
	LastUpdatedBy string `json:"lastUpdatedBy"`
	Connection    string `json:"connection"`
	Identity      string `json:"identity"`
}

type Connection struct {
	Id            string `json:"_id"`
	Owner         string `json:"owner"`
	Modified      int64  `json:"modified"`
	LastUpdatedBy string `json:"lastUpdatedBy"`
	Nickname      string `json:"nickname"`
	Address       string `json:"address"`
	Type          int64  `json:"type"`
	Port          int64  `json:"port"`
}

type ConnectionGroup struct {
	Id            string `json:"_id"`
	Owner         string `json:"owner"`
	Modified      int64  `json:"modified"`
	LastUpdatedBy string `json:"lastUpdatedBy"`
	Name          string `json:"name"`
}

type ConnectionGroupMembership struct {
	Id              string `json:"_id"`
	Owner           string `json:"owner"`
	Modified        int64  `json:"modified"`
	LastUpdatedBy   string `json:"lastUpdatedBy"`
	Connection      string `json:"connection"`
	ConnectionGroup string `json:"group"`
}

type PortForward struct {
	Id            string `json:"_id"`
	Owner         string `json:"owner"`
	Modified      int64  `json:"modified"`
	LastUpdatedBy string `json:"lastUpdatedBy"`
	Name          string `json:"name"`
	Mode          int64  `json:"mode"`
	Connection    string `json:"connection"`
	Host          string `json:"host"`
	LocalPort     int64  `json:"localPort"`
	RemotePort    int64  `json:"remotePort"`
	OpenInBrowser bool   `json:"openInBrowser"`
}

type ConfigItem struct {
	Id       string `json:"id"`
	Field    string `json:"field"`
	Value    string `json:"value"`
	Modified int64  `json:"modified"`
}
