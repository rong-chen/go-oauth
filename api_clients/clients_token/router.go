package clients_token

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(r *gin.RouterGroup) {
	router := r.Group("clients_token")
	{
		router.GET("/refresh", RefreshToken)
	}
}
