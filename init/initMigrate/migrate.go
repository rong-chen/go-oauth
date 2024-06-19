package initMigrate

import (
	"go-oauth/api_web/user"
	"go-oauth/global"
)

var list = []interface{}{
	&user.Info{},
}

func InitMigrate() {
	// 初始化数据库迁移
	for _, a := range list {
		err := global.MySqlDb.AutoMigrate(a)
		if err != nil {
			panic(err.Error())
			break
		}
	}
}
