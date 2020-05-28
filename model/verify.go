package model

type Verify struct {
	Signature string `form:"signature"`
	Timestamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	OpenID    string `form:"openid"`
}
