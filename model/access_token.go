package model

const AccessTokenTable = "access_token"

type AccessToken struct {
	ID        int32
	Token     string `json:access_token`
	CreatedAt int32
	ExpiresIn int32 `json:expires_in`
}

func (a *AccessToken) TableName() string {
	return AccessTokenTable
}
