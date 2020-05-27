package constant

import "os"

var (
	TOKEN  = os.Getenv("TOKEN")
	ID     = os.Getenv("ID")
	APPID  = os.Getenv("APPID")
	SECRET = os.Getenv("SECRET")
)

const (
	GET_TOKEN_URL     = "https://api.weixin.qq.com/cgi-bin/token"
	CLIENT_CREDENTIAL = "client_credential"
)
