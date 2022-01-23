package common

import (
	"mime/multipart"

	"github.com/Yan-Bin-Lin/DCreater/app/apperr"
	"github.com/Yan-Bin-Lin/DCreater/app/log"
	"github.com/Yan-Bin-Lin/DCreater/app/setting"
	"github.com/Yan-Bin-Lin/DCreater/app/util/file"
	"github.com/gin-gonic/gin"
)

// write file and parse to html
func WriteFormFile(c *gin.Context, form *multipart.Form, fileName string) {
	fileHeader := form.File["content"][0]
	filePath := setting.Servers["main"].FilePath + "/" + form.Value["oid"][0] + "/" + fileName
	// check exist and create
	if err := file.Checkdir(filePath); err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "something error in write file", "something error in create folder")
	}
	go func() {
		if err := file.SaveMarkdown2Tmpl(fileHeader, filePath, fileName+".html"); err != nil {
			log.Warn(c, apperr.ErrWrongArgument, err, "something error in write file", "something error in parse markdown")
		}
	}()
	if err := file.SaveFile(fileHeader, filePath, fileName+".md", true); err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "something error in write file")
		return
	}
}
