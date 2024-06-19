package main

import (
	"github.com/gin-gonic/gin"
	"go-oauth/init/initMigrate"
	"go-oauth/init/initMySql"
	"go-oauth/init/initRedis"
	"go-oauth/init/initRouter"
	"go-oauth/init/initViper"
	"log"
)

func main() {
	//初始化配置文件
	c := initViper.InitViper()
	//	初始化数据库
	initMySql.InitMySQL(c)
	initRedis.InitRedis(c)
	initMigrate.InitMigrate()
	g := gin.New()
	initRouter.InitRouter(g)
	err := g.Run(":999")
	if err != nil {
		log.Fatal("run server error" + err.Error())
	}
}
