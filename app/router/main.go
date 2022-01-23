package router

import (
	"net/http"
	"path/filepath"

	"github.com/Yan-Bin-Lin/DCreater/app/middleware"
	"github.com/Yan-Bin-Lin/DCreater/app/serve"
	"github.com/Yan-Bin-Lin/DCreater/app/setting"

	"github.com/gin-gonic/gin"
)

func MainRouter() *gin.Engine {
	r := gin.New()
	gin.SetMode(setting.Servers["main"].RunMode)

	r.Use(middleware.Logging())
	r.Use(middleware.ErrorHandle())

	r.LoadHTMLGlob(filepath.Join(setting.ExecutePath, setting.Config.Servers["main"].ViewPath, "html/**/*"))

	//home page
	//r.GET("/", serve.GetRoot)

	//asset
	// serve js and css file
	r.StaticFS("/static", http.Dir(setting.Config.Servers["main"].ViewPath))
	// user file
	file := r.Group("/file/:type/:oid/:name")
	{
		file.GET("", serve.GetFile)
		file.POST("", middleware.Auth(), serve.UploadFile)
	}

	//blog
	blog := r.Group("/blog")
	blog.GET("", serve.GetRoot)

	//owner
	blog.GET("/:owner", serve.GetOwner)

	//works
	blog.GET("/:owner/*work", serve.GetBlog)

	//auth
	owner := blog.Group("/:owner")
	{
		owner.Use(middleware.TokenExist())
		owner.POST("", serve.CreateOwner)
		owner.PUT("", serve.UpdateOwner)
		owner.DELETE("", serve.DelOwner)
	}

	work := blog.Group("/:owner/*work")
	{
		work.Use(middleware.Auth())
		work.POST("", serve.CreateBlog)
		work.PUT("", serve.UpdateBlog)
		work.DELETE("", serve.DelBlog)
	}

	//account
	account := r.Group("/account")
	account.Use(middleware.CheckRobot())
	account.POST("/login", serve.Login)
	user := account.Group("/user")
	{
		user.Use(middleware.TokenExist())
		user.PUT("", serve.PutUser)
		user.DELETE("", serve.DelUser)
	}

	return r
}
