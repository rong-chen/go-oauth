package oauth_user

import (
	"fmt"
	"go-oauth/global"
)

func FindUserRow(key string, val string) (info Info) {
	global.MySqlDb.Where(fmt.Sprintf("%s = ?", key), val).Find(&info)
	return
}

func Create(info *Info) error {
	return global.MySqlDb.Create(&info).Error
}
