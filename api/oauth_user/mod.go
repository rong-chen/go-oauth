package oauth_user

import (
	"go-oauth/global"
	"time"
)

type Info struct {
	Username     string     `json:"username" form:"username" gorm:"username" `
	Password     string     `json:"password"  form:"password" gorm:"password"`
	Phone        string     `json:"phone"  form:"phone" gorm:"phone"`
	Email        string     `json:"email;not null;unique"  form:"email" gorm:"email"`
	Nickname     string     `json:"nickname"  form:"nickname" gorm:"nickname"`
	Avatar       string     `json:"avatar"  form:"avatar" gorm:"avatar"`
	Disabled     bool       `json:"disabled" form:"disabled" gorm:"disabled"`
	Sex          string     `json:"sex"  form:"sex" gorm:"sex"`
	Birthday     string     `json:"birthday"  form:"birthday" gorm:"birthday"`
	DisabledTime *time.Time `json:"disabledTime" form:"disabledTime" gorm:"disabledTime"`
	global.Model
}

type LoginParams struct {
	State       string `json:"state"  form:"state" gorm:"state"`
	Username    string `json:"username" form:"username" gorm:"username" binding:"required" `
	Password    string `json:"password"  form:"password" gorm:"password" binding:"required"`
	ClientId    string `json:"client_id"  form:"client_id" gorm:"client_id" binding:"required"`
	RedirectUrl string `json:"redirect_url"  form:"redirect_url" gorm:"redirect_url" binding:"required"`
}

func (receiver Info) TableName() string {
	return "oauth_user"
}
