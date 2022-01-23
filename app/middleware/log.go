package middleware

import (
	"fmt"
	"time"

	"github.com/Yan-Bin-Lin/DCreater/app/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	timeFormat = time.RFC3339
	utc        = true
)

// Requests with errors are logged using zap.Error().
// Requests with A known error are logged using zap.Warn().
// Requests without errors are logged using zap.Info().
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		start := time.Now()
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		// get log msg
		var (
			LogMsg *log.LogDataStruct
			ok     = false
		)
		value, exist := c.Get(log.LogCtxKey)
		if exist {
			LogMsg, ok = value.(*log.LogDataStruct)
		}
		if !ok || !exist || LogMsg == nil {
			LogMsg = &log.LogDataStruct{
				TaskStatus: log.TaskUnknow,
			}
		}

		// check error
		var lv zapcore.Level
		if LogMsg.TaskStatus == log.TaskFail {
			lv = zap.WarnLevel
		} else if LogMsg.TaskStatus == log.TaskError {
			lv = zap.ErrorLevel
		} else {
			lv = zap.InfoLevel
		}

		// write log
		fmt.Println(lv, c.Writer.Status(), c.Request.Method, path, query, c.ClientIP(), c.Request.UserAgent(), end.Format(timeFormat), latency, LogMsg)
		lgoData, _ := c.Get("LogData")
		logData := lgoData.(*log.LogDataStruct)
		if logData.ErrorMeta != nil {
			fmt.Println(logData.ErrorMeta)
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err) // 這已經是頂層的 UI 介面了，想以自己的方式呈現錯誤
				fmt.Println(log.Logger)
			}
		}()

		if ce := log.Logger.Check(lv, path); ce != nil {
			ce.Write(
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("time", end.Format(timeFormat)),
				zap.Duration("latency", latency),
				zap.Object("LogMsg", LogMsg),
			)
		}
	}
}
