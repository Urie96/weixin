package handler

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/Urie96/weixin/chatbot"
	"github.com/Urie96/weixin/constant"
	"github.com/Urie96/weixin/model"
	"github.com/Urie96/weixin/util"
	"github.com/Urie96/weixin/wxctx"
	"github.com/gin-gonic/gin"
)

func autoReply(c *gin.Context) {
	verify := &model.Verify{}
	c.BindQuery(verify)
	if !validateURL(verify.Timestamp, verify.Nonce, verify.Signature) {
		c.AbortWithStatus(403)
		return
	}
	getmsg := &model.Msg{}
	err := c.BindXML(getmsg)
	util.CheckError(err)
	util.PrintStruct(getmsg)
	ctx := wxctx.GetContextByOpenID(verify.OpenID)
	sendMsg := model.NewMsg(getmsg.FromUserName)
	content := getTextAndVoiceContent(getmsg)
	sendMsg.Content = chatbot.Chat(ctx, content)
	c.Data(200, "application/xml", sendMsg.ToXMLBytes())
}

func getTextAndVoiceContent(msg *model.Msg) string {
	switch msg.MsgType {
	case "text":
		return msg.Content
	case "voice":
		return msg.Recognition
	}
	return ""
}

// func withOpenID(c *gin.Context, openid string) context.Context {
// 	ctx := c.Request.Context()
// 	return context.WithValue(ctx, "wx_context", wxctx.GetContextByOpenID(openid))
// }

func procSignature(c *gin.Context) {
	// verify := &struct {
	// 	Signature string
	// 	Timestamp string
	// 	Nonce     string
	// 	Echostr   string
	// }{}
	verify := &model.Verify{}
	c.BindQuery(verify)
	if !validateURL(verify.Timestamp, verify.Nonce, verify.Signature) {
		c.AbortWithStatus(403)
		return
	}
	c.String(200, verify.Echostr)
}

func validateURL(timestamp, nonce, signature string) bool {
	signatureGen := makeSignature(timestamp, nonce)
	return signatureGen == signature
}

func makeSignature(timestamp, nonce string) string { //本地计算signature
	si := []string{constant.TOKEN, timestamp, nonce}
	sort.Strings(si)            //字典序排序
	str := strings.Join(si, "") //组合字符串
	s := sha1.New()             //返回一个新的使用SHA1校验的hash.Hash接口
	io.WriteString(s, str)      //WriteString函数将字符串数组str中的内容写入到s中
	return fmt.Sprintf("%x", s.Sum(nil))
}
