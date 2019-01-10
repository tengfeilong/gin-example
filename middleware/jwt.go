package middleware

import (
	"gin-example/utils/msg"
	"gin-example/utils/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = msg.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = msg.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = msg.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = msg.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != msg.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  msg.GetMsg(code),
				"data": data,
			})
			//c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
