package middleware

import (
	"ginShop/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func TokenVer() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")

		if token == "" {
			util.ResponseWithJson(9003, "", "", c)
			c.Abort()
			return
		} else {
			claims, err := util.ParseToken(token)
			if err != nil { //token无法解析
				util.ResponseWithJson(9004, "", "", c)
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt { //token过期
				util.ResponseWithJson(9005, "", "", c)
				c.Abort()
				return
			} else {
				c.Set("ID", claims.ID)
				c.Next()
			}
		}
	}
}
