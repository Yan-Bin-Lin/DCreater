package common

import (
	"net/http"
	"net/url"

	"github.com/Yan-Bin-Lin/DCreater/app/apperr"
	"github.com/Yan-Bin-Lin/DCreater/app/log"
	"github.com/gin-gonic/gin"
)

// get the cookie value
func GetCookieParam(c *gin.Context, name string) (url.Values, error) {
	var (
		accessCookie *http.Cookie
		err          error
	)
	// check cookie
	if accessCookie, err = c.Request.Cookie(name); err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, name+" not found in cookie")
		return nil, err
	}

	// check access
	param, err := url.ParseQuery(accessCookie.Value)
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, name+" parse fail")
		return nil, err
	}

	return param, nil
}
