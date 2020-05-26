package model

import (
	"fmt"
	"time"
)

type Msg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Content      string
	MsgId        int
}

const (
	cdata  = `<![CDATA[%s]]>`
	myname = "yang"
)

func NewMsg(touser string) *Msg {
	return &Msg{
		ToUserName:   touser,
		FromUserName: myname,
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
