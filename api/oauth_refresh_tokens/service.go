package oauth_refresh_tokens

import "go-oauth/global"

func Create(row OAuthRefreshToken) error {
	return global.MySqlDb.Create(&row).Error
}
