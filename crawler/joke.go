package crawler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Urie96/weixin/dao"
	"github.com/Urie96/weixin/model"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

func CrawlJokes() {
	for i := 1; i < 644; i++ {
		jokeListPage(i)
	}
}

func jokeListPage(pageIndex int) {
	url := fmt.Sprintf("http://www.jokeji.cn/list_%d.htm", pageIndex)
	resp, err := http.Get(url)
	checkError(err)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkError(err)
	doc.Find("div.list_title ul li  b a").Each(func(i int, s *goquery.Selection) {
		path, _ := s.Attr("href")
		jokePage("http://www.jokeji.cn" + path)
	})
}

func jokePage(url string) {
	// url := "http://www.jokeji.cn/jokehtml/bxnn/2020032818121586.htm"
	now := int32(time.Now().UTC().Unix())
	converter, _ := iconv.NewConverter("gb2312", "utf-8")
	resp, err := http.Get(url)
	checkError(err)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkError(err)
	doc.Find("span p").Each(func(i int, s *goquery.Selection) {
		str, _ := converter.ConvertString(s.Text())
		offset := strings.Index(str, "、") + len("、")
		if len(str) < offset {
			return
		}
		joke := &model.Joke{
			Content:   str[offset:],
			URL:       url,
			CreatedAt: now,
		}
		insertToDB(joke)
	})
}

func insertToDB(joke *model.Joke) {
	dao.InsertJoke(joke)
}
