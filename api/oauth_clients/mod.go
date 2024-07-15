package oauth_clients

import (
	"github.com/google/uuid"
	"go-oauth/global"
)

type OAuthClients struct {
	ClientId     uuid.UUID `json:"-" form:"client_id" gorm:"commit:客户端id"`
	ClientSecret string    `json:"-" form:"client_secret"  gorm:"commit:客户端密钥"`
	Scope        string    `json:"scope" form:"scope"  gorm:"commit:权限" binding:"required"`
	ClientName   string    `json:"client_name" form:"client_name"  gorm:"commit:客户端名称" binding:"required"`
	RedirectUri  string    `json:"redirect_uri" form:"redirect_uri"  gorm:"commit:重定向uri"  binding:"required"`
	GrantType    string    `json:"grant_type" form:"grant_type"  gorm:"commit:授权类型，默认code"  binding:"required"`
	ClientUrl    string    `json:"client_url" form:"client_url"  gorm:"commit:客户端URL"  binding:"required"`
	LogoUri      string    `json:"logo_uri" form:"logo_uri"  gorm:"commit:客户端logo"`
	UserId       string    `json:"user_id" form:"user_id" gorm:"commit:创建者ID;index"`
	global.Model
}

func (receiver OAuthClients) TableName() string {
	return "oauth_clients"
}
