package middleware

import (
	"net"
	"os"
	"strings"

	"github.com/Yan-Bin-Lin/DCreater/app/apperr"
	"github.com/Yan-Bin-Lin/DCreater/app/log"
	"github.com/gin-gonic/gin"
)

// ErrorHandling returns a middleware that recovers from any panics and writes a 500 if there was one
// if no panic but there is error in context error, handle for warning to cleint side
func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			var brokenPipe bool
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// detect error type
				if brokenPipe {
					// connect error, can't send to front
					log.Error(c, apperr.ErrConnectFail, err.(error), 1)
				} else {
					// system error, need to send to front
					log.Error(c, apperr.ErrSystemFail, err.(error), 1, "Sorry, something error of system")
				}

				c.Abort()
			}

			// if no error
			err := c.Errors.Last()
			if err == nil {
				// log success
				log.Success(c)
				return
			}

			// If the connection is dead, we can't write a status to it.
			if !brokenPipe {

				// get apperr meta
				var (
					code int
					msg  string
				)
				switch err.Meta.(type) {
				case *apperr.ErrorDataStruct:
					meta := err.Meta.(*apperr.ErrorDataStruct)
					code = meta.Code
					msg = meta.Msg
				default:
					// worng type or something error
					code = apperr.ErrWrongErrorType
					msg = "Sorry, Something error"
					log.Warn(c, code, nil, msg)
				}

				// return to client
				_, httpStatus, _ := apperr.SplitCode(code)
				c.JSON(httpStatus, gin.H{
					"Code": code,
					"Msg":  msg,
				})
			}
		}()
		c.Next()
	}
}
