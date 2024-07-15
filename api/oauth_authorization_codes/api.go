package oauth_authorization_codes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-oauth/api/oauth_access_tokens"
	"go-oauth/api/oauth_clients"
	"go-oauth/api/oauth_refresh_tokens"
	"go-oauth/global"
	"go-oauth/utils"
	"time"
)

func ValidAuthorizationCode(c *gin.Context) {
	var data ValidateAuthorizationCodesData
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(200, global.BackResp(400, "参数错误", err.Error()))
		return
	}

	// 查看注册的客户端 是否存在
	client := oauth_clients.FindClientsRow("client_id", data.ClientId)
	if client.Id == uuid.Nil {
		c.JSON(200, global.BackResp(400, "无效的client", ""))
		return
	}

	if ok := utils.ComparePasswords(client.ClientSecret, data.ClientSecret); !ok {
		c.JSON(200, global.BackResp(400, "secret失效", ""))
		return
	}

	//查询code，并且判断code是否过期
	row := FindCodeRow("authorization_code", data.Code)
	if row.Id == uuid.Nil {
		c.JSON(200, global.BackResp(400, "参数错误", ""))
		return
	}

	nowTime := time.Now()
	if !row.ExpiresAt.After(nowTime) {
		c.JSON(200, global.BackResp(400, "参数过期", ""))
		return
	}

	//生成token
	var params utils.Params
	params.UserId = row.UserId
	aToken, err := utils.GenerateJWT(params, utils.AccessToken, time.Now().Add(15*time.Minute))
	rToken, err := utils.GenerateJWT(params, utils.RefreshToken, time.Now().Add(15*24*time.Hour))

	if err != nil {
		c.JSON(200, global.BackResp(400, "网络错误", err.Error()))
		return
	}

	//生成aToken和rToken的表数据
	oat := oauth_access_tokens.OAuthAccessToken{
		AccessToken: aToken,
		ClientId:    row.ClientId,
		UserId:      row.UserId,
		Scope:       "code",
		ExpiresAt:   time.Now().Add(15 * time.Minute),
	}

	oat.Id, _ = uuid.NewUUID()
	ort := oauth_refresh_tokens.OAuthRefreshToken{
		RefreshToken: rToken,
		ClientId:     row.ClientId,
		UserId:       row.UserId,
		Scope:        "code",
		ExpiresAt:    time.Now().Add(15 * 24 * time.Hour),
	}
	ort.Id, _ = uuid.NewUUID()

	err = oauth_access_tokens.Create(oat)
	err = oauth_refresh_tokens.Create(ort)

	if err != nil {
		c.JSON(200, global.BackResp(400, "网络错误", err.Error()))
		return
	}

	DeleteRow("authorization_code", data.Code)
	resp := make(map[string]interface{})
	resp["Refresh_Token"] = rToken
	resp["Access_Token"] = aToken
	c.JSON(200, global.BackResp(200, "", resp))
}
