package chatbot

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Urie96/weixin/util"
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

func asyncCallCMD(ctx context.Context, cmd string) string {
	// ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*5))
	ch := make(chan string, 2)
	go callCMD(ctx, cmd, ch)
	select {
	case output := <-ch:
		return output
	case <-time.After(time.Duration(time.Second * 4)):
		return "命令正在努力执行中，稍后可通过#get获取输出"
	}
}

func callCMD(ctx context.Context, cmd string, ch chan string) {
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
	// setValue(ctx, string(output))
	ch <- string(output)
}
