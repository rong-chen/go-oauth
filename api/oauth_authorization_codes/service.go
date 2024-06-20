package oauth_authorization_codes

import (
	"fmt"
	"go-oauth/global"
)

func Create(data *OAuthAuthorizationCodes) error {
	return global.MySqlDb.Create(data).Error
}
func FindCodeRow(key, val string) (row OAuthAuthorizationCodes) {
	global.MySqlDb.Where(fmt.Sprintf("%s = ?", key), val).Find(&row)
	return
}
func DeleteRow(key, val string) {
	global.MySqlDb.Where(fmt.Sprintf("%s =?", key), val).Delete(&OAuthAuthorizationCodes{})
}
