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
	Birthday     int        `json:"birthday"  form:"birthday" gorm:"birthday"`
	DisabledTime *time.Time `json:"disabledTime" form:"disabledTime" gorm:"disabledTime"`
	global.Model
}

type RegisterInfo struct {
	Username string `json:"username" form:"username" gorm:"username" binding:"required"`
	Password string `json:"password"  form:"password" gorm:"password" binding:"required"`
	Phone    string `json:"phone"  form:"phone" gorm:"phone" binding:"required"`
	Email    string `json:"email"  form:"email" gorm:"email" binding:"required"`
	Nickname string `json:"nickname"  form:"nickname" gorm:"nickname" binding:"required"`
	Sex      string `json:"sex"  form:"sex" gorm:"sex" binding:"required"`
	Birthday int    `json:"birthday"  form:"birthday" gorm:"birthday" binding:"required"`
	Code     string `json:"code"  form:"code" gorm:"code" binding:"required"`
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
