package initMigrate

import (
	"go-oauth/api/oauth_access_tokens"
	"go-oauth/api/oauth_authorization_codes"
	"go-oauth/api/oauth_clients"
	"go-oauth/api/oauth_refresh_tokens"
	"go-oauth/api/oauth_user"
	"go-oauth/global"
)

var list = []interface{}{
	&oauth_user.Info{},
	&oauth_clients.OAuthClients{},
	&oauth_authorization_codes.OAuthAuthorizationCodes{},
	&oauth_refresh_tokens.OAuthRefreshToken{},
	&oauth_access_tokens.OAuthAccessToken{},
}

func InitMigrate() {
	// 初始化数据库表
	for _, a := range list {
		err := global.MySqlDb.AutoMigrate(a)
		if err != nil {
			panic(err.Error())
			break
		}
	}
}
