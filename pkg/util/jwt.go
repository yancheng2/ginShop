package util

import (
	"ginShop/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret) //jwt秘钥

type Claims struct {
	ID     int
	Mobile string
	jwt.StandardClaims
}

//生成token
func GeteraterToken(id int, mobile string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(setting.AppSetting.JwtExpireTime * time.Hour)

	claims := Claims{
		id,
		mobile,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}

	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := withClaims.SignedString(jwtSecret)
	return token, err
}

//解析以及校验token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
