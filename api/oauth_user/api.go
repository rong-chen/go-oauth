package oauth_user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-oauth/api/oauth_authorization_codes"
	"go-oauth/api/oauth_clients"
	"go-oauth/global"
	"go-oauth/utils"
	"time"
)

func Login(c *gin.Context) {
	var lp LoginParams
	err := c.BindJSON(&lp)
	if err != nil {
		c.JSON(200, global.BackResp(400, "参数错误", nil))
		return
	}
	//校验用户名密码是否一致
	info := FindUserRow("username", lp.Password)
	if ok := utils.ComparePasswords(info.Password, lp.Password); !ok {
		c.JSON(200, global.BackResp(400, "账号密码错误", nil))
		return
	}
	// 验证客户端名称是否在数据库中
	clientInfo := oauth_clients.FindClientsRow("client_id", lp.ClientId)
	if clientInfo.Id == uuid.Nil {
		c.JSON(200, global.BackResp(400, "客户端不存在", nil))
		return
	}
	// 生成一个code(uuid)
	code, _ := uuid.NewUUID()
	rowData := &oauth_authorization_codes.OAuthAuthorizationCodes{
		ClientId:          uuid.MustParse(lp.ClientId),
		UserId:            info.Id,
		RedirectUri:       lp.RedirectUrl,
		Scope:             "",
		ExpiresAt:         time.Now().Add(5 * time.Minute),
		AuthorizationCode: code,
	}

	//生成OAuthAuthorizationCodes的uuid
	rowId, _ := uuid.NewUUID()
	rowData.Id = rowId
	//保存到code数据库中，等待后续校验。该code只能使用一次，就要执行删除操作
	err = oauth_authorization_codes.Create(rowData)
	if err != nil {
		c.JSON(200, global.BackResp(400, "网络出错", err.Error()))
		return
	}
	resp := fmt.Sprintf("%s?authorization_code=%s&state=%s&expires_time=%s", lp.RedirectUrl, code, lp.State, time.Now().Add(5*time.Minute).Format("2006-01-02T15:04:05.000Z"))
	c.JSON(200, global.BackResp(302, "登录成功", resp))
}

func Register(c *gin.Context) {
	// 注册用户
	var cp RegisterInfo
	err := c.BindJSON(&cp)
	if err != nil {
		c.JSON(200, global.BackResp(400, err.Error(), nil))
		return
	}
	code, _ := global.RedisDb.Get(c, cp.Email).Result()
	global.RedisDb.Del(c, cp.Email)
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
		Sex:      cp.Sex,
		Birthday: cp.Birthday,
	}

	info.Id, _ = uuid.NewUUID()
	err = Create(info)
	if err != nil {
		c.JSON(200, global.BackResp(400, err.Error(), nil))
		return
	}
	c.JSON(200, global.BackResp(200, "注册成功", nil))
}
