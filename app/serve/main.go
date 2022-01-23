package serve

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Yan-Bin-Lin/DCreater/app/apperr"
	"github.com/Yan-Bin-Lin/DCreater/app/common"
	"github.com/Yan-Bin-Lin/DCreater/app/database"
	"github.com/Yan-Bin-Lin/DCreater/app/log"
	"github.com/Yan-Bin-Lin/DCreater/app/setting"
	"github.com/Yan-Bin-Lin/DCreater/app/util/file"
	"github.com/gin-gonic/gin"
)

var (
	createBlogParam = []string{"oid", "superid", "descript", "blogType"}
	updateBlogParam = append(createBlogParam, "bid", "newsuperid", "newname", "newsuperUrl")
	blogType        = map[string]string{"project": "1", "article": "2"}
)

/*home page*/

// get root
func GetRoot(c *gin.Context) {
	blogDatas, err := database.GetRoot(c.DefaultQuery("p", "1"))
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error", "database error of getting root page")
		return
	}

	meta, err := json.Marshal(blogDatas)
	if err != nil {
		log.Warn(c, apperr.ErrWrongArgument, err, "Sorry, something error", "parse json error")
		return
	}

	// return root page data
	c.HTML(http.StatusOK, "index", gin.H{
		"meta":        string(meta),
		"title":       "DCreater",
		"description": "create your blog",
		"author":      "林彥賓, https://github.com/Yan-Bin-Lin",
		"root":        true,
		"list":        true,
	})
}

/*owner*/

func GetOwner(c *gin.Context) {
	// check project exist
	ownerData, err := database.GetOwner(c.Param("owner"))
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error", "database error of getting owner")
		return
	}

	// drop uid
	ownerData.Uid = 0

	meta, err := json.Marshal(ownerData)
	if err != nil {
		log.Warn(c, apperr.ErrWrongArgument, err, "Sorry, something error", "parse json error")
		return
	}

	// return owner with blog
	c.HTML(http.StatusOK, "index", gin.H{
		"meta":        string(meta),
		"title":       ownerData.Nickname,
		"description": ownerData.Description,
		"list":        true,
	})
}

// create a new owner
func CreateOwner(c *gin.Context) {
	// get cookie param
	param, err := common.GetCookieParam(c, "AccessToken")
	if err != nil {
		return
	}

	if err := database.CreateOwner(param.Get("uid"), c.Param("owner"), c.PostForm("nickname"), c.PostForm("descript")); err != nil {
		if err == database.ERR_NAME_CONFLICT {
			log.Warn(c, apperr.ErrWrongArgument, err, "Name conflict of create owner")
		} else if err == database.ERR_TASK_FAIL {
			log.Warn(c, apperr.ErrWrongArgument, err, "put owner fail, please check oid and uid correct")
		} else {
			log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error. please try again", "database error of create owner")
		}
		return
	}

	c.Header("Location", "/.")
	c.Status(http.StatusCreated)
}

// update a new owner
func UpdateOwner(c *gin.Context) {
	// get cookie param
	param, err := common.GetCookieParam(c, "AccessToken")
	if err != nil {
		return
	}

	log.Debug("ee", param)

	if err := database.UpdateOwner(param.Get("uid"), c.PostForm("oid"), c.Param("owner"), c.PostForm("newuniname"), c.PostForm("nickname"), c.PostForm("descript")); err != nil {
		if err == database.ERR_NAME_CONFLICT {
			log.Warn(c, apperr.ErrWrongArgument, err, "Name conflict of update owner")
		} else if err == database.ERR_TASK_FAIL {
			log.Warn(c, apperr.ErrWrongArgument, err, "put owner fail, please check oid and uid correct")
		} else {
			log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error. please try again", "database error of update owner")
		}
		return
	}

	c.Header("Location", "/.")
	c.Status(http.StatusCreated)
}

func DelOwner(c *gin.Context) {
	// get cookie param
	param, err := common.GetCookieParam(c, "AccessToken")
	if err != nil {
		return
	}

	err = database.DelOwner(param.Get("uid"), c.PostForm("oid"), c.Param("owner"))
	if err != nil {
		if err == database.ERR_TASK_FAIL {
			log.Warn(c, apperr.ErrWrongArgument, err, "delete owner fail, please check oid and owner name correct")
		} else {
			log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error. please try again", "database error of delete owner")
		}
		return
	}

	if err := os.RemoveAll(setting.Servers["main"].FilePath + "/" + c.PostForm("oid")); err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "something error in delete file")
		return
	}

	c.Status(http.StatusResetContent)
}

/*blog*/

func GetBlog(c *gin.Context) {
	blogData, err := database.GetBlog("/" + c.Param("owner") + c.Param("work"))
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error", "database error")
		return
	} else if blogData == nil {
		log.Warn(c, apperr.ErrWrongArgument, nil, "parmeter error", "parmeter error")
		return
	}

	meta, err := json.Marshal(blogData)
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error", "parse json error")
		return
	}

	data := gin.H{
		"title":       blogData.Name,
		"meta":        string(meta),
		"description": blogData.Description,
		"owner":       blogData.OUniquename,
		"nickname":    blogData.ONickname,
		"updatetime":  blogData.Updatetime,
	}

	if blogData.Type == 1 {
		data["list"] = true
		err = file.ParseTmpl(c.Writer, data)
	} else {
		data["content"] = true
		//get file
		fileName := strconv.Itoa(blogData.Bid)

		err = file.ParseTmpl(c.Writer, data, setting.Servers["main"].FilePath+"/"+strconv.Itoa(blogData.Oid)+"/"+fileName+"/"+fileName+".md")
	}
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error", "read file error")
		return
	}

	c.Status(http.StatusOK)
}

func CreateBlog(c *gin.Context) {
	// check form
	form, err := common.BindMultipartForm(c, createBlogParam)
	if err != nil {
		return
	}

	if form.Value["blogType"][0] != "project" && len(form.File["content"]) == 0 {
		log.Warn(c, apperr.ErrWrongArgument, nil, "multy part form miss match key content")
		return
	}

	// create to database
	superUrl, blog := splitWork("/" + c.Param("owner") + c.Param("work"))
	err = database.CreateBlog(form.Value["oid"][0], form.Value["superid"][0], blog, form.Value["descript"][0], blogType[form.Value["blogType"][0]], strings.Join(superUrl, "/"))
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error. please try again", "database error of update blog")
		return
	}

	// get data
	blogData, err := database.GetBlog("/" + c.Param("owner") + c.Param("work"))
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error", "database error")
		return
	}

	// write file
	if form.Value["blogType"][0] == "article" {
		common.WriteFormFile(c, form, strconv.Itoa(blogData.Bid))
	}

	c.Header("Location", "/.")
	c.Status(http.StatusCreated)
}

func UpdateBlog(c *gin.Context) {
	// check form
	form, err := common.BindMultipartForm(c, updateBlogParam)
	if err != nil {
		return
	}

	// check param update to database
	_, blog := splitWork("/" + c.Param("owner") + c.Param("work"))
	// new super should be -1 if no update super
	// new name should be "" if no update name
	fmt.Println(form.Value["oid"][0], form.Value["superid"][0], form.Value["newsuperid"][0],
		form.Value["bid"][0], blog, form.Value["newname"][0], form.Value["descript"][0], "/"+c.Param("owner")+c.Param("work"),
		form.Value["newsuperUrl"][0])
	err = database.UpdateBlog(form.Value["oid"][0], form.Value["superid"][0], form.Value["newsuperid"][0],
		form.Value["bid"][0], blog, form.Value["newname"][0], form.Value["descript"][0], "/"+c.Param("owner")+c.Param("work"),
		form.Value["newsuperUrl"][0])
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error. please try again", "database error of update blog")
		return
	}

	// write file
	if form.Value["blogType"][0] == "article" && len(form.File["content"]) != 0 {
		common.WriteFormFile(c, form, form.Value["bid"][0])
	}

	c.Header("Location", c.Request.URL.Path)
	c.Status(http.StatusCreated)
}

func DelBlog(c *gin.Context) {
	if err := database.DelBlog(c.PostForm("oid"), c.PostForm("bid"), "/"+c.Param("owner")+c.Param("work")); err != nil {
		if err == database.ERR_TASK_FAIL {
			log.Warn(c, apperr.ErrWrongArgument, err, "delete owner fail, please check oid and owner name correct")
		} else {
			log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error. please try again", "database error of delete owner")
		}
		return
	}

	if err := os.RemoveAll(setting.Servers["main"].FilePath + "/" + c.PostForm("oid") + "/" + c.PostForm("bid")); err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "something error in delete file")
		return
	}

	c.Status(http.StatusResetContent)
}

// split url to slice of projects and last project or blog
func splitWork(url string) ([]string, string) {
	works := strings.Split(url, "/")
	return works[:len(works)-1], works[len(works)-1]
}
