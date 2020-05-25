package handler

import (
	"net/http"
	"sort"

	"github.com/Urie96/weixin/model"

	"github.com/Urie96/weixin/util"
)

func VerifyHandler(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	verify := &model.Verify{}
	util.DecodeURLParamsToStruct(req, verify)
	if checkSignature(verify) {
		resp.Write([]byte(verify.Echostr))
	} else {
		resp.Write([]byte{})
	}
}

func checkSignature(data *model.Verify) bool {
	token := "ilovehuyue"
	arr := []string{token, data.Timestamp, data.Nonce}
	sort.Sort(sort.StringSlice(arr))
	return arr[0]+arr[1]+arr[2] == data.Signature
}
