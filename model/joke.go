package model

const JokeTable = "joke"

type Joke struct {
	ID      string
	Content string
}

func (_ *Joke) TableName() string {
	return JokeTable
}
