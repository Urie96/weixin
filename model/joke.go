package model

const JokeTable = "joke"

type Joke struct {
	ID        string
	Content   string
	URL       string
	CreatedAt int32
}

func (_ *Joke) TableName() string {
	return JokeTable
}
