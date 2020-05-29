package main

import (
	"fmt"

	"github.com/Urie96/weixin/crawler"
)

func main() {
	for y := 2020; y < 2100; y++ {
		solar := crawler.GetSolar(y, 6, 27)
		fmt.Println(`"%s":true,`, solar)
	}

}
