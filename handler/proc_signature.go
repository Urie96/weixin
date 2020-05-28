package handler

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/Urie96/weixin/constant"
	"github.com/gin-gonic/gin"
)

func procSignature(c *gin.Context) {
	verify := &struct {
		Signature string
		Timestamp string
		Nonce     string
		Echostr   string
	}{}
	c.BindQuery(verify)
	if !validateURL(verify.Timestamp, verify.Nonce, verify.Signature) {
		c.String(200, "")
		return
	}
	c.String(200, verify.Echostr)
}

func validateURL(timestamp, nonce, signature string) bool {
	signatureGen := makeSignature(timestamp, nonce)
	return signatureGen == signature
}

func makeSignature(timestamp, nonce string) string { //本地计算signature
	si := []string{constant.TOKEN, timestamp, nonce}
	sort.Strings(si)            //字典序排序
	str := strings.Join(si, "") //组合字符串
	s := sha1.New()             //返回一个新的使用SHA1校验的hash.Hash接口
	io.WriteString(s, str)      //WriteString函数将字符串数组str中的内容写入到s中
	return fmt.Sprintf("%x", s.Sum(nil))
}
