package oauth_access_tokens

import "go-oauth/global"

func Create(row OAuthAccessToken) error {
	return global.MySqlDb.Create(&row).Error
}
