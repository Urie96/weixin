package constant

import "os"

var (
	TOKEN  = os.Getenv("TOKEN")
	ID     = os.Getenv("ID")
	APPID  = os.Getenv("APPID")
	SECRET = os.Getenv("SECRET")
)

const (
	WX_API            = "https://api.weixin.qq.com/cgi-bin"
	CLIENT_CREDENTIAL = "client_credential"
)
