package util

import (
	"fmt"
	"time"

	"github.com/Urie96/weixin/dao"
	"github.com/Urie96/weixin/model"
	"github.com/nosixtools/solarlunar"
)

func InsertFestival(month, day, typ, msg string, isSolar bool) {
	for y := time.Now().Year(); y <= solarlunar.MAX_YEAR; y++ {
		var timestamp int64
		if isSolar {
			solar := fmt.Sprintf("%d-%s-%s", y, month, day)
			timestamp = solarToTimestamp(solar)
		} else {
			lunar := fmt.Sprintf("%d-%s-%s", y, month, day)
			solar := solarlunar.LunarToSolar(lunar, false)
			timestamp = solarToTimestamp(solar)
		}

		f := &model.Festival{
			Date: timestamp,
			Type: typ,
			Msg:  msg,
		}
		dao.InsertFestival(f)
	}
}

func solarToTimestamp(solar string) int64 {
	solarTime, _ := time.ParseInLocation("2006-01-02", solar, time.Local)
	solarTimestamp := solarTime.Unix()
	return solarTimestamp
}
