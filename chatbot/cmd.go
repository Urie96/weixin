package chatbot

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Urie96/weixin/util"
	"github.com/Urie96/weixin/wxctx"
)

// type response struct {
// 	text string
// }

// func withValue(ctx context.Context) context.Context {
// 	return context.WithValue(ctx, "response", &response{})
// }

// func getValue(ctx context.Context) string {
// 	return ctx.Value("response").(*response).text
// }

// func setValue(ctx context.Context, value string) {
// 	ctx.Value("response").(*response).text = value
// }

func handleCMD(ctx *wxctx.Context, cmd string) string {
	if cmd == "#" {
		ctx.IsInCmdMode = true
		return "进入cmd模式"
	}
	if cmd == "exit" {
		ctx.IsInCmdMode = false
		return "退出cmd模式"
	}
	if cmd == "GET" {
		return ctx.LastOutput
	}
	return asyncCallCMD(ctx, cmd)
}

func asyncCallCMD(ctx *wxctx.Context, cmd string) string {
	ctx.LastOutput = "命令正在努力执行中，稍后可通过GET命令获取输出"
	ch := make(chan string, 2)
	go callCMD(ctx, cmd, ch)
	select {
	case output := <-ch:
		return output
	case <-time.After(time.Duration(time.Second * 4)):
		return ctx.LastOutput
	}
}

func callCMD(ctx *wxctx.Context, cmd string, ch chan string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover:", err)
		}
	}()
	url := os.Getenv("CMD_URL")
	resp, err := http.Post(url, "text/plain", strings.NewReader(cmd))
	util.CheckError(err)
	output, err := ioutil.ReadAll(resp.Body)
	util.CheckError(err)
	ch <- string(output)
	ctx.LastOutput = string(output)
}
