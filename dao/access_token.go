package dao

import (
	"log"

	"github.com/Urie96/weixin/model"
)

func InsertAccessToken(accesstoken *model.AccessToken) error {
	err := db.Create(accesstoken).Error
	if err != nil {
		log.Println("insert access token error:", err)
	}
	return err
}
