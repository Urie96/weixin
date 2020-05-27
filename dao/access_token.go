package dao

import (
	"log"

	"github.com/Urie96/weixin/model"
)

func InsertAccessToken(accesstoken *model.AccessToken) error {
	err := db.Table(model.AccessTokenTable).Create(accesstoken).Error
	if err != nil {
		log.Println("insert access token error:", err)
	}
	return err
}
