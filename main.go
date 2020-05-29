package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Urie96/weixin/util"
)

type a struct {
	Info string
}

func main() {
	resp, err := http.Get("http://cmd.sweetlove.top?cmd=ls%20-l")
	util.CheckError(err)
	output, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(output))
}
