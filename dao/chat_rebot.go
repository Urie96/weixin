package dao

import (
	"log"
	"time"

	"github.com/Urie96/weixin/model"
)

func InsertChatRecord(chatRecord *model.ChatRecord) error {
	chatRecord.CreatedAt = time.Now().UTC().Unix()
	err := db.Create(chatRecord).Error
	if err != nil {
		log.Println("insert chat record to db error:", err)
	}
	return err
}
