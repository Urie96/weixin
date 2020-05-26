package handler

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"

	"github.com/Urie96/weixin/model"
	"github.com/gin-gonic/gin"
)

const (
	token = "ilovehuyue" //设置token
)

func procSignature(c *gin.Context) {
	verify := &model.Verify{}
	c.BindQuery(verify)
	signatureGen := makeSignature(verify.Timestamp, verify.Nonce)
	log.Println(signatureGen)
	if signatureGen != verify.Signature {
		c.String(200, "")
		return
	}
	c.String(200, verify.Echostr)
}

func makeSignature(timestamp, nonce string) string { //本地计算signature
	si := []string{token, timestamp, nonce}
	sort.Strings(si)            //字典序排序
	str := strings.Join(si, "") //组合字符串
	s := sha1.New()             //返回一个新的使用SHA1校验的hash.Hash接口
	io.WriteString(s, str)      //WriteString函数将字符串数组str中的内容写入到s中
	return fmt.Sprintf("%x", s.Sum(nil))
}
