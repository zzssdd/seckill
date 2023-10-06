package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"seckill/conf"
	"time"
)

type MyClaim struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenToken(id int64, email string) (string, error) {
	claim := MyClaim{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "yogen",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(conf.Salt)
}

func ParseToken(tokenString string) (*MyClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return conf.Salt, nil
	})
	if err != nil || token == nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*MyClaim); ok && token.Valid {
		return claim, nil
	}
	return nil, fmt.Errorf("token不合法")
}
