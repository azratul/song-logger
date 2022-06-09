package gdrive

type Client struct {
	ClientEmail string   `json:"client_email"`
	PrivateKey  string   `json:"private_key"`
	Parents     []string `json:"parents"`
}
