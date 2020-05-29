package chatbot

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	url := os.Getenv("CMD_URL")
	fmt.Println(url)
	resp, err := http.Get(url + "?cmd=" + cmd)
	util.CheckError(err)
	output, err := ioutil.ReadAll(resp.Body)
	util.CheckError(err)
	// setValue(ctx, string(output))
	ch <- string(output)
}
