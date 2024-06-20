package user

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(group *gin.RouterGroup) {
	r := group.Group("resource_user")
	{
		r.POST("/register", Register)
	}
}
