package client_user

import "go-oauth/global"

type Info struct {
	Username string `json:"username" `
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	global.Model
}

type LoginParams struct {
	Username string `json:"username" form:"username" gorm:"username"  binding:"required"`
	Password string `json:"password" form:"password" gorm:"password"  binding:"required"`
}
