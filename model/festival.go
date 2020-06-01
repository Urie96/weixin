package model

const FestivalTable = "festival"

type Festival struct {
	ID   int
	Date int64
	Type string
	Msg  string
}

func (_ *Festival) TableName() string {
	return FestivalTable
}
