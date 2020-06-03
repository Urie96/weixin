package model

const ChatRecordTable = "chat_record"

type ChatRecord struct {
	ID        int
	Question  string
	Answer    string
	CreatedAt int64
	OpenID    string
}

func (_ *ChatRecord) TableName() string {
	return ChatRecordTable
}
