package serve

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Yan-Bin-Lin/DCreater/app/apperr"
	"github.com/Yan-Bin-Lin/DCreater/app/common"
	"github.com/Yan-Bin-Lin/DCreater/app/database"
	"github.com/Yan-Bin-Lin/DCreater/app/log"
	"github.com/Yan-Bin-Lin/DCreater/app/setting"
	"github.com/Yan-Bin-Lin/DCreater/app/util/random"
	"github.com/gin-gonic/gin"
)

// generate a new token and save it
func newAccessToken(uid string) (string, error) {
	code, err := random.GetRandomString(32)
	if err != nil {
		return "", err
	}
	return code, database.NewAccessToken(uid, code)
}

// check access token vaild
func CheckAccessAuth(c *gin.Context) {
	// get cookie param
	param, err := common.GetCookieParam(c, "AccessToken")
	if err != nil {
		c.Abort()
		return
	}

	if has, err := database.CheckAccessAuth(param.Get("uid"), param.Get("AccessCode"), c.PostForm("oid")); err != nil {
		log.Warn(c, 1500006, err, "Sorry, something error", "database error of check access token")
		c.Abort()
		return
	} else if !has {
		log.Warn(c, apperr.ErrPermissionDenied, err, "access token parse fail")
		c.Abort()
		return
	}
}

// check access token exist
func CheckAccessToken(c *gin.Context) {
	// get cookie param
	param, err := common.GetCookieParam(c, "AccessToken")
	if err != nil {
		c.Abort()
		return
	}

	if has, err := database.CheckAccessToken(param.Get("uid"), param.Get("AccessCode")); err != nil {
		log.Warn(c, 1500006, err, "Sorry, something error", "database error of check access token")
		c.Abort()
		return
	} else if !has {
		log.Warn(c, apperr.ErrPermissionDenied, err, "access token parse fail")
		c.Abort()
		return
	}
}

// check request is by robot
func CheckRobot(c *gin.Context) {
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
		url.Values{
			"secret":   {setting.Servers["main"].ReCAPTCHA["key"]},
			"response": {c.PostForm("response")},
			"remoteip": {c.ClientIP()},
		},
	)
	if err != nil {
		log.Warn(c, apperr.ErrPermissionDenied, err, "Sorry, something error", "error in sent post request to reCAPTCHA")
		c.Abort()
		return
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	// parse google response
	if val, ok := res["success"].(bool); !ok {
		err := errors.New("wrong error type")
		log.Error(c, apperr.ErrSystemFail, err, 0, "Sorry, something error", "assert wrong type")
		c.Abort()
		return
	} else if !val {
		err := errors.New("wrong argument")
		log.Warn(c, apperr.ErrWrongArgument, err, "Sorry, something error", "wrong response token")
		c.Abort()
		return
	}
	// compare score
	score, err := strconv.ParseFloat(setting.Servers["main"].ReCAPTCHA["AcceptScore"], 64)
	if err != nil {
		log.Warn(c, apperr.ErrWrongArgument, err, "Sorry, something error", "parse config string to float fail")
		c.Abort()
		return
	}
	if val, ok := res["score"].(float64); !ok {
		err = errors.New("wrong error type")
		log.Error(c, apperr.ErrSystemFail, err, 0, "Sorry, something error", "assert wrong type")
		c.Abort()
		return
	} else if val < score {
		err := errors.New("robot denied")
		log.Warn(c, apperr.ErrWrongArgument, err, "Sorry, we don't accept robot", "robot denied")
		c.Abort()
		return
	}
}

/*
// generate a new refresh token
func NewRefrshToken(c *gin.Context, userName string, uid uint64) (string, error) {

	rereshToken, err := random.GetRandomString(32)
	if err != nil {
		return "", err
	}

	if err := database.CreateToken(database.RefreshTokenTable, rereshToken, uid); err != nil {
		return "", err
	}

	return rereshToken, nil
}

// generate a new access token
func NewAccessToken(refreshToken string) (string, error) {
	// check refresh token first
	uid, err := GetuidByToken(database.RefreshTokenTable, refreshToken)
	if err != nil {
		return "", err
	}

	accessToken, err := hashToken(refreshToken)
	if err != nil {
		return "", err
	}

	if err := database.CreateToken(database.AccessTokenTable, accessToken, uid); err != nil {
		return "", err
	}

	return accessToken, nil
}

// refresh token uid and uid not same error
var ERR_INVALID_USER error = errors.New("Invalid of user with this refresh token")

// check a refresh token is valid
func GetuidByRefreshToken(refreshToken string) (uint64, error) {
	return GetuidByToken(database.RefreshTokenTable, refreshToken)
}

// check a access token is valid
func GetuidByAccessToken(accessToken string) (uint64, error) {
	return GetuidByToken(database.AccessTokenTable, accessToken)
}

// check a token is valid
func GetuidByToken(authType, token string) (uint64, error) {
	result, err := database.GetuidbyToken(authType, token)
	if err != nil {
		return 0, err
	}

	return result, nil
}
*/
