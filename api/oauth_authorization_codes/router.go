package oauth_authorization_codes

import "github.com/gin-gonic/gin"

type Router struct{}

func (Router) InitRouter(r *gin.RouterGroup) {
	router := r.Group("authorization")
	{
		router.POST("/code", ValidAuthorizationCode)
	}
}
