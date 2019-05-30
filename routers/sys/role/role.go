package role

import (
	models "cooky-go/models/sys"
	"cooky-go/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoleRouter(r *gin.Engine) {
	role := r.Group("/role")
	role.GET("", SelectRole)
	role.POST("", AddRole)
}

func SelectRole(ctx *gin.Context) {
	depts := models.SelectAllRole()
	result := make(map[string]interface{})
	result["data"] = depts
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"data": result,
	})
}

func AddRole(ctx *gin.Context) {
	var role models.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "新增角色失败",
		})
	}
	models.AddRole(role)

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "新增角色成功",
	})
}
