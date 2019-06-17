package menu

import (
	models "cooky-go/models/sys"
	"cooky-go/pkg/e"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitMenuRouter(r *gin.Engine) {
	menu := r.Group("/menu")
	menu.GET("", SelectMenuTree)
	menu.POST("", AddMenu)
	menu.PUT("", EditMenu)
	menu.DELETE("/:menuId", DeleteMenu)
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

func AddMenu(ctx *gin.Context) {
	var menu models.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "新增菜单参数错误",
		})
		return
	}
	models.AddMenu(menu)
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "新增菜单成功",
	})
}

func EditMenu(ctx *gin.Context) {
	var menu models.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "编辑菜单参数错误",
		})
		return
	}
	models.EditMenu(menu)
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "编辑菜单成功",
	})
}

func DeleteMenu(ctx *gin.Context) {
	menuId := com.StrTo(ctx.Param("menuId")).MustInt()
	models.DeleteMenu(menuId)
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "删除菜单成功",
	})
}
