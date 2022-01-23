package apperr

import (
	"github.com/Yan-Bin-Lin/DCreater/app/setting"
	"github.com/Yan-Bin-Lin/DCreater/app/util/debug"
)

// struct of error message to let out
type ErrorDataStruct struct {
	Code int    `zap:"code"`
	Msg  string `zap:"msg"`
}

// error message for debug
type ErrorMetaStruct struct {
	Msg   string
	Error error // origin error
	Stack []*debug.FuncDataStruct
}

// return error for log
type ErrorReturn struct {
	ErrorData *ErrorDataStruct
	ErrorMeta *ErrorMetaStruct
}

func (er *ErrorReturn) Error() string {
	return er.ErrorData.Msg
}

// implement Unwrap for error in go 1.13
func (er *ErrorReturn) Unwrap() error {
	return er.ErrorMeta.Error
}

func (ems *ErrorMetaStruct) GetStack() []*debug.FuncDataStruct {
	return ems.Stack
}

// implement Unwrap for error in go 1.13
func (ems *ErrorMetaStruct) Unwrap() error {
	return ems.Error
}

//generate new error data
func NewErrorData(code int, customMsg string) (eData *ErrorDataStruct) {
	if customMsg != "" {
		eData = &ErrorDataStruct{
			Code: code,
			Msg:  customMsg,
		}
	} else {
		eData = &ErrorDataStruct{
			Code: code,
			Msg:  GetMsg(code),
		}
	}
	return
}

//generate new error meta
func NewErrorMeta(err error, stack []*debug.FuncDataStruct, customMsg string) *ErrorMetaStruct {
	return &ErrorMetaStruct{
		Msg:   customMsg,
		Error: err,
		Stack: stack,
	}
}

// new a error return struct
func NewErrorReturn(code int, err error, stack []*debug.FuncDataStruct, customMsg ...string) *ErrorReturn {
	var (
		errorData *ErrorDataStruct
		errorMeta *ErrorMetaStruct
	)

	// check error msg
	if len(customMsg) > 0 {
		errorData = NewErrorData(code, customMsg[0])
	} else {
		errorData = NewErrorData(code, "")
	}

	if len(customMsg) > 1 {
		errorMeta = NewErrorMeta(err, stack, customMsg[1])
	} else {
		errorMeta = NewErrorMeta(err, stack, "")
	}

	return &ErrorReturn{errorData, errorMeta}
}

// check if error code exist
func GetMsg(code int) string {
	if val, ok := setting.ErrorMap[code]; ok {
		return val
	} else {
		return setting.ErrorMap[0]
	}
}

// split error code
// return errorType, httpStatus, customCode
func SplitCode(code int) (int, int, int) {
	if _, ok := setting.ErrorMap[code]; !ok {
		//error code not found
		return 0, 0, 0
	}

	return code / 1000000, (code % 1000000) / 1000, code % 1000
}