package handler

import (
	"github.com/Urie96/weixin/crawler"
	"github.com/Urie96/weixin/dao"
	"github.com/gin-gonic/gin"
)

func createJokes(c *gin.Context) {
	crawler.CrawlJokes()
	c.String(200, "finished")
}

func getAJoke(c *gin.Context) {
	joke, err := dao.GetRandomJokeContent()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.String(200, joke)
}
