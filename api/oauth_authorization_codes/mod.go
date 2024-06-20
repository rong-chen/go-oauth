package oauth_authorization_codes

import (
	"github.com/google/uuid"
	"go-oauth/global"
	"time"
)

type OAuthAuthorizationCodes struct {
	UserId            uuid.UUID `json:"userId"`
	AuthorizationCode uuid.UUID `json:"authorizationCode"`
	ClientId          uuid.UUID `json:"clientId"`
	RedirectUri       string    `json:"redirectUri"`
	Scope             string    `json:"scope"`
	ExpiresAt         time.Time `json:"expiresAt"`
	global.Model
}

type ValidateAuthorizationCodesData struct {
	Code         string `json:"authorization_code" form:"authorization_code" binding:"required"`
	ClientId     string `json:"client_id"  form:"client_id" binding:"required"`
	ClientSecret string `json:"client_secret"  form:"client_secret" binding:"required"`
	GrantType    string `json:"grant_type"  form:"grant_type" binding:"required"`
}

func (receiver OAuthAuthorizationCodes) TableName() string {
	return "oauth_authorization_codes"
}
