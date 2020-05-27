package dao

import (
	"log"

	"github.com/Urie96/weixin/model"
)

func InsertAccessToken(accesstoken *model.AccessToken) error {
	err := db.Create(accesstoken).Error
	if err != nil {
		log.Println("insert access token to db error:", err)
	}
	return err
}

func GetAccessToken() (*model.AccessToken, error) {
	token := &model.AccessToken{}
	err := db.Table(model.AccessTokenTable).Last(token).Error
	if err != nil {
		log.Println("get access token from db error:", err)
	}
	return token, err
}
