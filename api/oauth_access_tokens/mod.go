package oauth_access_tokens

import (
	"github.com/google/uuid"
	"go-oauth/global"
	"time"
)

type OAuthAccessToken struct {
	AccessToken string    `json:"access_token"`
	ClientId    uuid.UUID `json:"client_id"`
	UserId      uuid.UUID `json:"user_id"`
	Scope       string    `json:"scope"`
	ExpiresAt   time.Time `json:"expires_At"`
	global.Model
}

func (receiver OAuthAccessToken) TableName() string {
	return "oauth_access_token"
}
