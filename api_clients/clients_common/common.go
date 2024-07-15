package clients_common

import (
	"github.com/gin-gonic/gin"
	"go-oauth/global"
	"go-oauth/utils"
)

func ValidToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(200, global.BackResp(400, "Authorization不能为空", nil))
		c.Abort()
		return
	}
	t, err := utils.ParseJWT(token)
	if err != nil {
		//修改成执行刷新aToken操作
		c.JSON(200, global.BackResp(401, "token无效", nil))
		c.Abort()
		return
	}
	if t.Types == utils.ClientsUserAccessToken {
		c.Set("user_id", t.Params.UserId.String())
		c.Next()
	} else {
		c.JSON(200, global.BackResp(401, "token类型错误", nil))
		c.Abort()
		return
	}
}
