package middleware

import (
	"baseapi/global/response"
	"baseapi/models/util"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var mes string

		mes = "SUCCESS"
		token := c.Query("token")
		if token == "" {
			mes = "INVALID_PARAMS"
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				mes = "ERROR_AUTH_CHECK_TOKEN_FAIL"
			} else if time.Now().Unix() > claims.ExpiresAt {
				mes = "ERROR_AUTH_CHECK_TOKEN_TIMEOUT"
			}
		}

		if mes != "SUCCESS" {
			response.FailWithMessage(mes, c)
			c.Abort()
			return
		}

		c.Next()
	}
}
