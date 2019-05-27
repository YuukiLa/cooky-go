package jwt

import (
	"cooky-go/pkg/e"
	"cooky-go/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		log.Println("jwt invoke")
		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.UN_AUTH
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.UN_AUTH
			}
			c.Set("claims",claims)
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : "token不存在或者失效",
				"data" : data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
