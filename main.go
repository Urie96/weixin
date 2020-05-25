package main

import (
	"log"
	"net/http"

	"github.com/Urie96/weixin/handler"
)

func main() {
	log.Println("Wechat Service: Start!")
	http.HandleFunc("/wx", handler.ProcSignature)
	err := http.ListenAndServe(":7001", nil)
	if err != nil {
		log.Println("Wechat Service: ListenAndServe Error: ", err)
	}
	log.Println("Wechat Service: Stop!")
}
