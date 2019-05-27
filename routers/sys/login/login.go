package login

import (
	models "cooky-go/models/sys"
	"cooky-go/pkg/e"
	"cooky-go/pkg/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func InitLoginRouter(r *gin.Engine) {
	r.POST("/login",Login)
}

func Login(ctx *gin.Context) {
	var user models.User
	if err:=ctx.ShouldBind(&user);err!=nil {

	}
	dbUser := models.FindByUsername(user.Username)
	log.Println(dbUser)
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password),[]byte(user.Password))
	if err!=nil {
		ctx.JSON(http.StatusOK,gin.H{
			"code": e.ERROR,
			"msg": "登录失败",
			"data": nil,
		})
	}else {
		roles := models.FindRoleIdByUserId(user.UserId)
		token, err := util.GenerateToken(user.Username, roles)
		if err!=nil{
			ctx.JSON(http.StatusOK,gin.H{
				"code": e.ERROR,
				"msg": "登录失败,生成token错误",
				"data": err,
			})
			return
		}
		result := make(map[string]interface{})
		result["token"] = token
		ctx.JSON(http.StatusOK,gin.H{
			"code": e.SUCCESS,
			"msg": "登录成功",
			"data": result,
		})
	}
}