package main

import (
	"log"

	_ "github.com/Urie96/weixin/dao"
	"github.com/Urie96/weixin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Wechat Service: Start!")
	router := gin.Default()
	handler.Handle(router)
	if err := router.Run(); err != nil {
		log.Println("Wechat Service: ListenAndServe Error: ", err)
	}
	log.Println("Wechat Service: Stop!")
}
