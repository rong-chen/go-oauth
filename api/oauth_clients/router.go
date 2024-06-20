package oauth_clients

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(r *gin.RouterGroup) {
	router := r.Group("clients")
	{
		router.POST("/client", Created)
	}
}
