package client_user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-oauth/api/oauth_user"
	"go-oauth/global"
	"go-oauth/utils"
	"time"
)

func Login(c *gin.Context) {
	var params LoginParams
	err := c.BindJSON(&params)
	if err != nil {
		c.JSON(200, global.BackResp(400, "参数错误", err.Error()))
		return
	}
	userInfo := oauth_user.FindUserRow("username", params.Username)
	if userInfo.Id == uuid.Nil {
		c.JSON(200, global.BackResp(400, "账号不存在", nil))
		return
	}
	if ok := utils.ComparePasswords(userInfo.Password, params.Password); !ok {
		c.JSON(200, global.BackResp(400, "密码错误", nil))
		return
	}
	var tp utils.Params
	tp.UserId = userInfo.Id
	rToken, err := utils.GenerateJWT(tp, utils.ClientsUserRefreshToken, time.Now().Add(time.Hour*24*7))
	aToken, err := utils.GenerateJWT(tp, utils.ClientsUserAccessToken, time.Now().Add(time.Minute*15))
	resp := make(map[string]interface{})
	resp["rToken"] = rToken
	resp["aToken"] = aToken
	resp["userInfo"] = userInfo
	c.JSON(200, global.BackResp(200, "登录成功", resp))
}

//
//func Register(c *gin.Context) {
//	var params RegisterParams
//	err := c.BindJSON(&params)
//
//	if err != nil {
//		c.JSON(200, global.BackResp(400, "参数错误", err.Error()))
//		return
//	}
//
//	row := FindUserRow("username", params.Username)
//	if row.Id != uuid.Nil {
//		c.JSON(200, global.BackResp(400, "账号已存在", err.Error()))
//		return
//	}
//
//	info := &Info{
//		Username: params.Username,
//		Password: params.Password,
//		Nickname: params.Nickname,
//	}
//
//	info.Id, err = uuid.NewUUID()
//	if err != nil {
//		c.JSON(200, global.BackResp(400, "网络错误", err.Error()))
//		return
//	}
//
//	err = Create(info)
//	if err != nil {
//		c.JSON(200, global.BackResp(400, "网络错误", err.Error()))
//		return
//	}
//	c.JSON(200, global.BackResp(200, "注册成功", nil))
//}
