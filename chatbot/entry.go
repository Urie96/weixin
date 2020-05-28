package chatbot

import (
	"strings"
)

const UnrecognizedAnswer = `对不起，我不够聪明，不能理解你的意思。
你可以试试这样：
1、讲笑话`

// 2、即将到来的日子`

func Chat(text string) string {
	switch text {
	case "1":
		return tellAJoke()
	case "2":
	}
	if strings.Contains(text, "笑话") {
		return tellAJoke()
	}
	return UnrecognizedAnswer
}
