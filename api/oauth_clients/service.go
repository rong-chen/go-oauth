package oauth_clients

import (
	"fmt"
	"go-oauth/global"
)

func CreateRow(data OAuthClients) error {
	return global.MySqlDb.Create(&data).Error
}

func FindClientsRow(key, val string) (client OAuthClients) {
	global.MySqlDb.Where(fmt.Sprintf("%s = ?", key), val).Find(&client)
	return
}

func FindClientsRowList(key, val string, offset, limit int) (client []OAuthClients, counts int64) {
	global.MySqlDb.Model(&OAuthClients{}).Where(fmt.Sprintf("%s = ?", key), val).Count(&counts)
	global.MySqlDb.Where(fmt.Sprintf("%s = ?", key), val).Offset(offset).Limit(limit).Find(&client)
	return
}
func DeleteRow(key []string) error {
	return global.MySqlDb.Delete(&OAuthClients{}, key).Error
}
