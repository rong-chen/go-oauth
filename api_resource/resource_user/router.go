package resource_user

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (Router) InitRouter(r *gin.RouterGroup) {
	router := r.Group("resource_user")
	{
		router.GET("/info", GetUserInfo)
	}
}
