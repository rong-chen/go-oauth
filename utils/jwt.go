package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	AccessToken  = "access_token"
	RefreshToken = "refresh_token"
	TicketToken  = "ticket_token"
)

// GetUserInfo 注册JWT的时候用户给第三方服务的选择的权限的常量
var (
	GetUserInfo = "GetUserInfo"
)

type Params struct {
	Id uuid.UUID `gorm:"primaryKey"`
}

type Token struct {
	Params               Params `json:"user"`
	Type                 string `json:"type"`
	SessionId            string `json:"sessionId"`
	jwt.RegisteredClaims        // v5版本新加的方法
}

var secretKey = []byte("user-registration-center")

func GenerateJWT(params Params, types string, sessionId string, ex time.Time) (string, error) {
	token := Token{
		params,
		types,
		sessionId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ex), // 30天过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, token)

	resp, err := t.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func ParseJWT(token string) (*Token, error) {
	t, err := jwt.ParseWithClaims(token, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if claims, ok := t.Claims.(*Token); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
