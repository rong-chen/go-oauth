package clients_manage

import (
	"github.com/gin-gonic/gin"
	"go-oauth/api_clients/clients_common"
)

type Router struct {
}

func (Router) InitRouter(router *gin.RouterGroup) {
	r := router.Group("/clients_manage")
	r.Use(clients_common.ValidToken)
	{
		r.GET("/list", List)
		r.DELETE("/delete", Delete)
	}
}
