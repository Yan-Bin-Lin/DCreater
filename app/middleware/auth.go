package middleware

import (
	"github.com/Yan-Bin-Lin/DCreater/app/serve"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		serve.CheckAccessAuth(c)
		if c.IsAborted() {
			return
		}

		c.Next()
	}
}

func TokenExist() gin.HandlerFunc {
	return func(c *gin.Context) {

		serve.CheckAccessToken(c)
		if c.IsAborted() {
			return
		}

		c.Next()
	}
}

func CheckRobot() gin.HandlerFunc {
	return func(c *gin.Context) {
		serve.CheckRobot(c)
		if c.IsAborted() {
			return
		}

		c.Next()
	}
}
