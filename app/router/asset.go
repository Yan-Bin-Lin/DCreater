package router

import (
	"net/http"

	"github.com/Yan-Bin-Lin/DCreater/app/middleware"
	"github.com/Yan-Bin-Lin/DCreater/app/serve"
	"github.com/Yan-Bin-Lin/DCreater/app/setting"
	"github.com/gin-gonic/gin"
)

func AssetRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(setting.Servers["account"].RunMode)

	r.Use(middleware.Logging())
	r.Use(middleware.ErrorHandle())

	// serve js and css file
	r.StaticFS("/static", http.Dir(setting.Config.Servers["main"].ViewPath))

	r.GET("/file/:type/:oid/:name", serve.GetFile)

	r.Use(middleware.Auth())
	r.POST("/file/:type/:oid/:name", serve.UploadFile)

	return r
}
