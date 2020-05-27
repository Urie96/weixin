package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Urie96/weixin/dao"

	"github.com/Urie96/weixin/model"

	"github.com/Urie96/weixin/constant"
	"github.com/gin-gonic/gin"
)

func insertAccessToken(c *gin.Context) {
	token, err := getAccessTokenFromWX()
	if err != nil {
		c.AbortWithError(500, err)
	}
	err = dao.InsertAccessToken(token)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.AbortWithStatus(200)
}

func getAccessTokenFromWX() (*model.AccessToken, error) {
	token := &model.AccessToken{}
	retry := 50
	for token.Token == "" && retry > 0 {
		token = sendParamToWX()
		retry--
	}
	if retry == 0 {
		return nil, errors.New("Can't get access token from wx")
	}
	return token, nil
}

func sendParamToWX() *model.AccessToken {
	ret := &model.AccessToken{CreatedAt: int32(time.Now().UTC().Unix())}
	url := fmt.Sprintf("%s?grant_type=%s&appid=%s&secret=%s",
		constant.GET_TOKEN_URL, constant.CLIENT_CREDENTIAL, constant.APPID, constant.SECRET)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Get access token from wx error:", err)
		return ret
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Get access token from wx error:", err)
		return ret
	}
	log.Println("Get response from wx for access token:", string(b))
	err = json.Unmarshal(b, ret)
	if err != nil {
		log.Println("Get access token from wx error:", err)
		return ret
	}
	return ret
}
