package resource_user

import (
	"github.com/gin-gonic/gin"
	"go-oauth/api/oauth_user"
	"go-oauth/global"
	"go-oauth/utils"
)

// GetUserInfo 客户端用来获取用户信息
func GetUserInfo(c *gin.Context) {
	Authorization := c.GetHeader("Authorization")
	if Authorization == "" {
		c.JSON(200, global.BackResp(400, "Authorization不能为空白", nil))
		return
	}
	p, err := utils.ParseJWT(Authorization)

	if err != nil || p.Types != utils.AccessToken {
		c.JSON(200, global.BackResp(400, "Authorization错误", err.Error()))
		return
	}
	// 获取用户信息
	userinfo := oauth_user.FindUserRow("id", p.Params.UserId.String())
	resp := make(map[string]interface{})
	resp["Nickname"] = userinfo.Nickname
	resp["Avatar"] = userinfo.Avatar
	resp["Sex"] = userinfo.Sex
	resp["Birthday"] = userinfo.Birthday
	c.JSON(200, global.BackResp(200, "", resp))
}
