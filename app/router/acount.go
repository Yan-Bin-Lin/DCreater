package router

import (
	"path/filepath"

	"github.com/Yan-Bin-Lin/DCreater/app/middleware"
	"github.com/Yan-Bin-Lin/DCreater/app/serve"
	"github.com/Yan-Bin-Lin/DCreater/app/setting"
	"github.com/gin-gonic/gin"
)

func AccountRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(setting.Servers["account"].RunMode)

	r.Use(middleware.Logging())
	r.Use(middleware.ErrorHandle())

	r.LoadHTMLGlob(filepath.Join(setting.ExecutePath, setting.Config.Servers["main"].ViewPath, "html/**/*"))

	r.Use(middleware.CheckRobot())
	r.POST("/login", serve.Login)
	user := r.Group("/user")
	{
		user.PUT("", serve.PutUser)
		user.DELETE("", serve.DelUser)
	}

	return r
}
