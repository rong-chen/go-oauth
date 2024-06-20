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
