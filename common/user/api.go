package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-oauth/global"
	"go-oauth/utils"
)

func Register(c *gin.Context) {
	// 注册用户
	var cp createParams
	err := c.BindJSON(&cp)
	if err != nil {
		c.JSON(200, global.BackResp(400, err.Error(), nil))
		return
	}
	ctx := context.Background()
	code, _ := global.RedisDb.Get(ctx, cp.Email).Result()
	global.RedisDb.Del(ctx, cp.Email)
	if code != cp.Code {
		c.JSON(200, global.BackResp(400, "验证码错误", nil))
		return
	}
	password, _ := utils.HashPassword(cp.Password)
	var info = &Info{
		Email:    cp.Email,
		Password: password,
		Phone:    cp.Phone,
		Username: cp.Username,
		Nickname: cp.Nickname,
	}

	info.Id, _ = uuid.NewUUID()
	err = Create(info)
	if err != nil {
		c.JSON(200, global.BackResp(400, err.Error(), nil))
		return
	}
	c.JSON(200, global.BackResp(200, "注册成功", nil))
}
