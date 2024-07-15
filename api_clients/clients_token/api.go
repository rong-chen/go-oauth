package clients_token

import (
	"github.com/gin-gonic/gin"
	"go-oauth/global"
	"go-oauth/utils"
	"time"
)

func RefreshToken(c *gin.Context) {
	rToken := c.Query("refresh_token")
	if rToken == "" {
		c.JSON(200, global.BackResp(400, "refresh_token不能为空", nil))
		return
	}
	t, e := utils.ParseJWT(rToken)
	if e != nil {
		c.JSON(200, global.BackResp(400, "refresh_token无效", nil))
		return
	}

	aToken, err := utils.GenerateJWT(t.Params, utils.ClientsUserAccessToken, time.Now().Add(15*time.Minute))
	if err != nil {
		c.JSON(200, global.BackResp(400, "token生成失败", nil))
		return
	}
	c.JSON(200, global.BackResp(200, "", aToken))
}
