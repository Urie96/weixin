package chatbot

import (
	"log"
	"strings"

	"github.com/Urie96/weixin/crawler"
	"github.com/Urie96/weixin/dao"
	"github.com/Urie96/weixin/model"
	"github.com/Urie96/weixin/wxctx"
)

const DefaultReply = `抱歉，现在的我还不能明白您的意思`

// 2、即将到来的日子`

func Chat(c *wxctx.Context, text string) string {
	if strings.HasPrefix(text, "#") || c.IsInCmdMode {
		return handleCMD(c, text)
	}
	if text == "节日" {
		return getFestivals()
	}
	if strings.Contains(text, "情话") {
		return crawler.GetLoveTalk()
	}
	reply := crawler.AIQA(text)
	if reply == "defaultReply" {
		return DefaultReply
	}
	go saveChatRecord(c.OpenID, text, reply)
	return reply
}

func saveChatRecord(openID, question, answer string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover:", err)
		}
	}()
	record := &model.ChatRecord{
		OpenID:   openID,
		Question: question,
		Answer:   answer,
	}
	dao.InsertChatRecord(record)
}
