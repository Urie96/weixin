package dao

import (
	"log"

	"github.com/Urie96/weixin/model"
)

func InsertFestival(festival *model.Festival) error {
	err := db.Create(festival).Error
	if err != nil {
		log.Println("insert festival to db error:", err)
	}
	return err
}

func GetFestivals(startTime, limit int64) []*model.Festival {
	var festivals []*model.Festival
	db.Table(model.FestivalTable).Where("date >= ?", startTime).Limit(limit).Order("date").Find(&festivals)
	return festivals
}
