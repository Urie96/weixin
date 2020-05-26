package handler

import (
	"github.com/Urie96/weixin/model"
	"github.com/Urie96/weixin/util"
	"github.com/gin-gonic/gin"
)

func autoReply(c *gin.Context) {
	getmsg := &model.Msg{}
	c.BindXML(getmsg)
	util.PrintStruct(getmsg)
	msg := model.NewMsg(getmsg.FromUserName)
	msg.Content = "test"
	c.Data(200, "application/xml", msg.ToXMLBytes())
}
