package constant

import "os"

var (
	TOKEN = os.Getenv("TOKEN")
	ID    = os.Getenv("ID")
)
