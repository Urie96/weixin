package util

import (
	"fmt"
	"time"

	"github.com/Urie96/weixin/dao"

	"github.com/Urie96/weixin/model"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

var day = int64(3600 * 24)
var msg = map[int64]string{
	day * 333:  "333天",
	day * 444:  "444天",
	day * 555:  "555天",
	day * 666:  "666天",
	day * 777:  "777天",
	day * 888:  "888天",
	day * 999:  "999天",
	day * 1000: "1000天",
	day * 400:  "400天",
	day * 500:  "500天",
	day * 600:  "600天",
	day * 700:  "700天",
	day * 800:  "800天",
	day * 900:  "900天",
	day * 1000: "1000天",
	day * 520:  "520天",
	day * 1314: "1314天",
	day * 1111: "1111天",
	day * 2222: "2222天",
	day * 3333: "3333天",
	day * 4444: "4444天",
	day * 5555: "5555天",
	day * 6666: "6666天",
	day * 7777: "7777天",
	day * 8888: "8888天",
	day * 9999: "9999天",
	30000000:   "三千万秒",
	40000000:   "四千万秒",
	50000000:   "五千万秒",
	60000000:   "六千万秒",
	70000000:   "七千万秒",
	80000000:   "八千万秒",
	90000000:   "九千万秒",
	100000000:  "一亿秒",
}

func InsertTogetherFestival() {
	kownTime, _ := time.ParseInLocation(TIME_LAYOUT, "2019-05-08 16:16:00", time.Local)
	loveTime, _ := time.ParseInLocation(TIME_LAYOUT, "2019-07-13 23:45:00", time.Local)
	meetTime, _ := time.ParseInLocation(TIME_LAYOUT, "2019-06-30 09:45:00", time.Local)
	insertTogetherFestivalToDB(kownTime.Unix(), "相识")
	insertTogetherFestivalToDB(loveTime.Unix(), "相恋")
	insertTogetherFestivalToDB(meetTime.Unix(), "见到你")
}

func insertTogetherFestivalToDB(date int64, typ string) {
	for k, v := range msg {
		festival := &model.Festival{
			Type: typ + v + "纪念日",
			Date: timestampZero(date + k),
		}
		dao.InsertFestival(festival)
	}
	for y := 1; y <= 100; y++ {
		festival := &model.Festival{
			Type: fmt.Sprintf("%s%d周年纪念日", typ, y),
			Date: timestampZero(date + day*365*int64(y)),
		}
		dao.InsertFestival(festival)
	}
}

func timestampZero(timestamp int64) int64 {
	t := time.Unix(timestamp, 0)
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return t1.Unix()
}
