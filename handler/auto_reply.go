package handler

import (
	"github.com/Urie96/weixin/model"
	"github.com/gin-gonic/gin"
)

func autoReply(c *gin.Context) {
	getmsg := &model.MsgFromUser{}
	c.BindXML(getmsg)

	tmp := getmsg.FromUserName
	getmsg.FromUserName = getmsg.ToUserName
	getmsg.ToUserName = tmp
	getmsg.Content = "test"

	c.XML(200, getmsg)
}
