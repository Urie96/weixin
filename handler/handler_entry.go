package handler

import "github.com/gin-gonic/gin"

func Handle(router *gin.Engine) {
	router.GET("/wx", procSignature)
	router.POST("/wx", autoReply)
}
