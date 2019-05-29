package routers

import (
	"cooky-go/middleware/cors"
	"cooky-go/middleware/jwt"
	"cooky-go/pkg/setting"
	"cooky-go/routers/sys/dept"
	"cooky-go/routers/sys/login"
	"cooky-go/routers/sys/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(setting.RunMode)
	r.Use(cors.Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//登录接口
	login.InitLoginRouter(r)
	//配置jwt校验中间件
	r.Use(jwt.JWT())
	//配置casbin权限校验中间件
	//a := gormadapter.NewAdapterByDB(models.DB)
	//e := casbin.NewEnforcer("conf/authz_model.conf", a)
	//r.Use(auth.CasbinHandler(e))

	user.InitUserRouter(r)
	dept.InitDeptRouter(r)
	return r
}
