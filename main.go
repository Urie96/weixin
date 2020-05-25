package main

import (
	"net/http"

	"github.com/Urie96/weixin/handler"
)

func main() {
	http.HandleFunc("/wx", handler.VerifyHandler)
	http.ListenAndServe(":7001", nil)
}
