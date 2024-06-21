package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	AccessToken             = "access_token"
	RefreshToken            = "refresh_token"
	ClientsUserAccessToken  = "ClientsUserAccessToken"
	ClientsUserRefreshToken = "ClientsUserRefreshToken"
)

type Params struct {
	UserId uuid.UUID `gorm:"primaryKey"`
}

type Token struct {
	Params               Params `json:"resource_user"`
	Types                string `json:"types"`
	jwt.RegisteredClaims        // v5版本新加的方法
}

var secretKey = []byte("resource_user-registration-center")

func GenerateJWT(params Params, types string, ex time.Time) (string, error) {
	token := Token{
		params,
		types,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ex),
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
		return secretKey, nil
	})
	if claims, ok := t.Claims.(*Token); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
