package main

import (
	"fmt"
	"time"
)

func main() {
	// util.InsertFestival("06", "27", "悦悦的生日", "", false)
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	fmt.Println(tm1.Unix())
}
