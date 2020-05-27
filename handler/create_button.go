package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Urie96/weixin/model"

	"github.com/Urie96/weixin/constant"
	"github.com/Urie96/weixin/dao"
	"github.com/gin-gonic/gin"
)

func createButton(c *gin.Context) {
	token, err := dao.GetAccessToken()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	url := fmt.Sprintf("%s/menu/create?access_token=%s", constant.WX_API, token.Token)
	b, err := json.Marshal(drawButton())
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	log.Println("Get response from wx for create button:", string(respBody))
	c.String(200, string(respBody))
}

func drawButton() *model.Buttons {
	button1 := &model.Button{
		Name: "悦悦",
		SubButton: []*model.Button{
			&model.Button{
				Name: "主页",
				Type: "view",
				URL:  "huyue.sweetlove.top",
			},
			&model.Button{
				Name: "检讨",
				Type: "view",
				URL:  "huyue.sweetlove.top/review",
			},
		},
	}
	button2 := &model.Button{
		Name: "锐锐",
		Type: "click",
		Key:  "1",
	}
	buttons := &model.Buttons{
		Button: []*model.Button{button1, button2},
	}
	return buttons
}
