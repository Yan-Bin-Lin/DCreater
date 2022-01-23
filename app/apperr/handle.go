package apperr

import (
	"github.com/Yan-Bin-Lin/DCreater/app/setting"
	"github.com/Yan-Bin-Lin/DCreater/app/util/debug"
	"errors"
	"github.com/gin-gonic/gin"
)

// return wrap error type
func New(code int, err error, skip int, customMsg ...string) error {
	// NewErrorReturn
	if setting.Servers["main"].RunMode == gin.DebugMode {
		return NewErrorReturn(code, err, debug.GetCallStack(skip+1), customMsg...) // skip this level
	} else {
		return NewErrorReturn(code, err, nil, customMsg...)
	}
}

// set context.error to handle to front end
// skip will be set for stack level(start from 0)
func ErrorHandle(c *gin.Context, code int, err error, skip int, customMsg ...string) (er *ErrorReturn) {
	er = New(code, err, skip, customMsg...).(*ErrorReturn)
	_ = c.Error(errors.New("")).SetMeta(er.ErrorData)
	return
}