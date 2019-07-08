package user

import (
	models "cooky-go/models/sys"
	"cooky-go/pkg/e"
	"cooky-go/pkg/setting"
	"cooky-go/pkg/util"
	"fmt"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func InitUserRouter(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("", AddUser)
		user.GET("", SelectUser)
		user.PUT("", EditUser)
		user.DELETE("/:userId", DeleteUser)
	}
}

func SelectUser(ctx *gin.Context) {
	username := ctx.Query("username")
	maps := make(map[string]interface{})
	result := make(map[string]interface{})
	if username != "" {
		maps["username"] = username
	}
	list := models.SelectUser(util.GetPage(ctx), setting.PageSize, maps)
	for i := 0; i < len(list); i++ {
		list[i].Password = "不告诉你！"
	}
	result["list"] = list
	result["total"] = models.GetUserTotal(maps)

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"data": result,
	})
}

func AddUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  e.GetMsg(e.INVALID_PARAMS),
			"data": err,
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	if err != nil {
		fmt.Println(err)
	}
	models.AddUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "新增成功",
		"data": nil,
	})
}

func EditUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  e.GetMsg(e.INVALID_PARAMS),
			"data": err,
		})
		return
	}
	user.Password = ""
	models.EditUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "修改用户成功",
		"data": nil,
	})
}

func DeleteUser(ctx *gin.Context) {
	userId := com.StrTo(ctx.Param("userId")).MustInt()
	models.DeleteUser(userId)
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "删除用户成功",
	})
}
