package chatbot

import (
	"fmt"
	"time"

	"github.com/Urie96/weixin/dao"
)

func getFestivals() string {
	today := getTodayZeroTimestamp()
	festivals := dao.GetFestivals(getTodayZeroTimestamp(), 10)
	reply := ""
	for _, festival := range festivals {
		gap := computeDayGap(today, festival.Date)
		reply += fmt.Sprintf("%d天后：%s\r\n", gap, festival.Msg)
	}
	return reply
}

func computeDayGap(from, to int64) int {
	day := int64(time.Hour * 24)
	return int((from - to) / day)
}

func getTodayZeroTimestamp() int64 {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return tm1.Unix()
}
