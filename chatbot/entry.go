package chatbot

import (
	"log"
	"strings"

	"github.com/Urie96/weixin/model"

	"github.com/Urie96/weixin/dao"

	"github.com/Urie96/weixin/crawler"

	"github.com/Urie96/weixin/wxctx"
)

const UnrecognizedAnswer = `对不起，我不够聪明，不能理解你的意思。
你可以试试这样：
1、讲笑话`

// 2、即将到来的日子`

func Chat(c *wxctx.Context, text string) string {
	if strings.HasPrefix(text, "#") || c.IsInCmdMode {
		return handleCMD(c, text)
	}
	if text == "节日" {
		return getFestivals()
	}
	reply := crawler.AIQA(text)
	go saveChatRecord(c.OpenID, text, reply)
	return reply
}

func saveChatRecord(openID, question, answer string) {
	if answer == "" || question == "" {
		return
	}
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
