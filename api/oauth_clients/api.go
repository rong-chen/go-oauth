package oauth_clients

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-oauth/global"
	"go-oauth/utils"
)

func Created(c *gin.Context) {
	var clientParams OAuthClients
	userId, _ := c.Get("user_id")

	err := c.BindJSON(&clientParams)
	if err != nil {
		c.JSON(200, global.BackResp(400, "参数错误", nil))
		return
	}
	clientParams.UserId = userId.(string)
	clientParams.ClientId, err = uuid.NewUUID()
	clientParams.Id, err = uuid.NewUUID()
	if err != nil {
		c.JSON(200, global.BackResp(400, "添加失败", err.Error()))
		return
	}
	secret := utils.GenerateRandomCode(36)
	//生成36位随机字符串，使用bcrypt不可逆加密。
	clientParams.ClientSecret, err = utils.HashPassword(secret)
	if err != nil {
		c.JSON(200, global.BackResp(400, "添加失败", err.Error()))
		return
	}
	err = CreateRow(clientParams)
	if err != nil {
		c.JSON(200, global.BackResp(400, "创建失败", nil))
		return
	}
	c.JSON(200, global.BackResp(200, "添加成功", secret))
}
