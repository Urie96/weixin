package model

import (
	"fmt"
	"time"

	"github.com/Urie96/weixin/constant"
)

type Msg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Content      string
	MsgId        int
	Event        string
	EventKey     string
}

func NewMsg(touser string) *Msg {
	return &Msg{
		ToUserName:   touser,
		FromUserName: constant.ID,
		MsgType:      "text",
		CreateTime:   int(time.Now().UTC().Unix()),
	}
}

func (m *Msg) ToXMLBytes() []byte {
	str := fmt.Sprintf(`<xml>
	<ToUserName><![CDATA[%s]]></ToUserName>
	<FromUserName><![CDATA[%s]]></FromUserName>
	<CreateTime>%d</CreateTime>
	<MsgType><![CDATA[%s]]></MsgType>
	<Content><![CDATA[%s]]></Content>
   </xml>`, m.ToUserName, m.FromUserName, m.CreateTime, m.MsgType, m.Content)
	return []byte(str)
}
