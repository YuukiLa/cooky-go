package auth

import (
	"cooky-go/pkg/util"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

//拦截器
func CasbinHandler(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {

		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		if obj == "/user/userInfo" {
			c.Next()
			return
		}
		//根据上下文获取载荷claims 从claims获得role
		claims := c.MustGet("claims").(*util.Claims)
		userName := claims.Username

		//判断策略中是否存在
		if e.Enforce(userName, obj, act) {
			fmt.Println("通过权限")
			c.Next()
		} else {
			fmt.Println("权限没有通过")
			c.Abort()
		}
	}
}
