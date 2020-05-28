package model

type Verify struct {
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
	OpenID    string `json:"openid"`
}
