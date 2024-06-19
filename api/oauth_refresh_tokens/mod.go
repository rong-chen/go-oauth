package oauth_refresh_tokens

import "go-oauth/global"

type OAuthAccessToken struct {
	AccessToken string `json:"access_token"`
	ClientId    string `json:"client_id"`
	UserId      string `json:"user_id"`
	Scope       string `json:"scope"`
	ExpiresAt   string `json:"expires_At"`
	global.Model
}
