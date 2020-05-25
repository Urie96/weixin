package handler

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/Urie96/weixin/util"

	"github.com/Urie96/weixin/model"
)

const (
	token = "ilovehuyue" //设置token
)

func makeSignature(timestamp, nonce string) string { //本地计算signature
	si := []string{token, timestamp, nonce}
	sort.Strings(si)            //字典序排序
	str := strings.Join(si, "") //组合字符串
	s := sha1.New()             //返回一个新的使用SHA1校验的hash.Hash接口
	io.WriteString(s, str)      //WriteString函数将字符串数组str中的内容写入到s中
	return fmt.Sprintf("%x", s.Sum(nil))
}

func validateUrl(w http.ResponseWriter, r *http.Request) bool {
	verify := &model.Verify{}
	util.DecodeURLParamsToStruct(r, verify)
	signatureGen := makeSignature(verify.Timestamp, verify.Nonce)

	if signatureGen != verify.Signature {
		return false
	}
	fmt.Fprintf(w, verify.Echostr)
	return true
}

func ProcSignature(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !validateUrl(w, r) {
		log.Println("Wechat Service: This http request is not from wechat platform")
		return
	}
	log.Println("validateUrl Ok")
}
