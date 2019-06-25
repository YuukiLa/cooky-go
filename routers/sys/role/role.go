package role

import (
	models "cooky-go/models/sys"
	"cooky-go/pkg/e"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoleRouter(r *gin.Engine) {
	role := r.Group("/role")
	role.GET("", SelectRole)
	role.POST("", AddRole)
	role.PUT("", EditRole)
	role.DELETE("/:roleId", DeleteRole)
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
	if err := ctx.Bind(&role); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "新增角色失败",
		})
		return
	}
	models.AddRole(role)

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "新增角色成功",
	})
}

func EditRole(ctx *gin.Context) {
	var role models.Role
	if err := ctx.Bind(&role); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "修改角色失败",
		})
		return
	}
	models.EditRole(role)

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "修改角色成功",
	})
}

func DeleteRole(ctx *gin.Context) {
	roleId := com.StrTo(ctx.Param("roleId")).MustInt()
	models.DeleteRole(roleId)
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "删除角色成功",
	})
}
