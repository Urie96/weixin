package chatbot

import (
	"context"
	"strings"
)

const UnrecognizedAnswer = `对不起，我不够聪明，不能理解你的意思。
你可以试试这样：
1、讲笑话`

// 2、即将到来的日子`

func Chat(ctx context.Context, text string) string {
	if strings.HasPrefix(text, "#") {
		return asyncCallCMD(ctx, text[len("#"):])
	}
	switch text {
	case "1":
		return tellAJoke()
	case "2":
		return getFestivals()
	}
	if strings.Contains(text, "笑话") {
		return tellAJoke()
	}
	return UnrecognizedAnswer
}
