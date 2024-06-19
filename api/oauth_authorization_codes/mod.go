package oauth_authorization_codes

import (
	"github.com/google/uuid"
	"go-oauth/global"
	"time"
)

type OAuthAuthorizationCodes struct {
	AuthorizationCode string    `json:"authorizationCode"`
	ClientId          uuid.UUID `json:"clientId"`
	UserId            uuid.UUID `json:"userId"`
	RedirectUri       string    `json:"redirectUri"`
	Scope             string    `json:"scope"`
	ExpiresAt         time.Time `json:"expiresAt"`
	global.Model
}
