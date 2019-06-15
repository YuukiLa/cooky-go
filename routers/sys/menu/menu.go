package menu

import (
	models "cooky-go/models/sys"
	"cooky-go/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitMenuRouter(r *gin.Engine) {
	menu := r.Group("/menu")
	menu.GET("", SelectMenuTree)
}

func SelectMenuTree(ctx *gin.Context) {
	menus := models.SelectAllMenu()
	result := make(map[string]interface{})
	result["data"] = menus
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"data": result,
	})
}
