package oauth_clients

import (
	"github.com/gin-gonic/gin"
	"go-oauth/api_clients/clients_common"
)

type Router struct {
}

func (Router) InitRouter(r *gin.RouterGroup) {
	router := r.Group("clients")
	{
		router.POST("/client", clients_common.ValidToken, Created)
	}
}
