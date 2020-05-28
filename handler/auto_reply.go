package handler

import (
	"io/ioutil"
	"log"

	"github.com/Urie96/weixin/util"

	"github.com/Urie96/weixin/model"
	"github.com/gin-gonic/gin"
)

func autoReply(c *gin.Context) {
	verify := &model.Verify{}
	c.BindQuery(verify)
	if !validateURL(verify.Timestamp, verify.Nonce, verify.Signature) {
		c.AbortWithStatus(403)
		return
	}
	b, _ := ioutil.ReadAll(c.Request.Body)
	log.Println("msg body:\r\n" + string(b))
	getmsg := &model.Msg{}
	c.BindXML(getmsg)
	util.PrintStruct(getmsg)
	msg := model.NewMsg(getmsg.FromUserName)
	msg.Content = "test"
	c.Data(200, "application/xml", msg.ToXMLBytes())
}
