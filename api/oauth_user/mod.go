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
	DisabledTime *time.Time `json:"disabledTime" form:"disabledTime" gorm:"disabledTime"`
	global.Model
}
