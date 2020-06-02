package chatbot

import (
	"fmt"
	"time"

	"github.com/Urie96/weixin/dao"
)

func getFestivals() string {
	today := getTodayZeroTimestamp()
	festivals := dao.GetFestivals(getTodayZeroTimestamp(), 10)
	reply := "节日："
	last := -1
	for _, festival := range festivals {
		gap := computeDayGap(today, festival.Date)
		if gap == last {
			reply += "\r\n        与      " + festival.Type
		} else {
			reply += fmt.Sprintf("\r\n%d天后：%s", gap, festival.Type)
		}
		last = gap
	}
	return reply
}

func computeDayGap(from, to int64) int {
	return int((to - from) / (3600 * 24))
}

func getTodayZeroTimestamp() int64 {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return tm1.Unix()
}
