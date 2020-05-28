package dao

import (
	"log"

	"github.com/Urie96/weixin/model"
)

func InsertJoke(joke *model.Joke) error {
	err := db.Create(joke).Error
	if err != nil {
		log.Println("insert joke to db error:", err)
	}
	return err
}

func GetRandomJokeContent() (joke string, err error) {
	err = db.Table(model.JokeTable).Order("RAND()").Limit(1).Pluck("content", joke).Error
	if err != nil {
		log.Println("get joke from db error(likely empty):", err)
	}
	return joke, err
}
