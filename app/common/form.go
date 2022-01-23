package common

import (
	"github.com/Yan-Bin-Lin/DCreater/app/apperr"
	"github.com/Yan-Bin-Lin/DCreater/app/log"
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

// check form key match or not
func checkParam(param []string, form *multipart.Form) string {
	for _, v := range param {
		if len(form.Value[v]) == 0 {
			return v
		}
	}

	return ""
}

func BindMultipartForm(c *gin.Context, param []string) (*multipart.Form, error) {
	form, err := c.MultipartForm()
	if err != nil {
		log.Warn(c, apperr.ErrWrongArgument, err, "binding error of put multipart form", "binding error of put multipart form")
		return nil, err
	}
	if v := checkParam(param, form); v != "" {
		errStr := "multi part form miss match key "+v
		log.Warn(c, apperr.ErrWrongArgument, nil, errStr)
		return nil, errors.New(errStr)
	}
	return form, nil
}