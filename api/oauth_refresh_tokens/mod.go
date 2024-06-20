package oauth_refresh_tokens

import (
	"github.com/google/uuid"
	"go-oauth/global"
	"time"
)

type OAuthRefreshToken struct {
	RefreshToken string    `json:"refresh_token"`
	ClientId     uuid.UUID `json:"client_id"`
	UserId       uuid.UUID `json:"user_id"`
	Scope        string    `json:"scope"`
	ExpiresAt    time.Time `json:"expires_At"`
	global.Model
}

func (receiver OAuthRefreshToken) TableName() string {
	return "oauth_refresh_token"
}
