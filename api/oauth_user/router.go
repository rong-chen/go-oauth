package oauth_user

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(r *gin.RouterGroup) {
	router := r.Group("oauth_user")
	{
		router.POST("/login", Login)
	}
}
