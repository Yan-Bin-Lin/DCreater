package debug

import (
	"runtime"
	"strconv"
)

const WrapLv = 0

type FuncDataStruct struct {
	Loacate  string
	Function string
}

// get the function data of caller, start from 0
func GetFuncData(skipNum ...int) *FuncDataStruct {
	skip := getWrapLv(skipNum...)
	pc, file, line, ok := runtime.Caller(skip + 1) // add one to skip this function
	if !ok {
		return nil
	}

	return &FuncDataStruct{
		Loacate:  file + ":" + strconv.Itoa(line),
		Function: runtime.FuncForPC(pc).Name(),
	}
}

// stack returns a nicely formatted stack frame, skipping skip frames.
func GetCallStack(skipNum ...int) []*FuncDataStruct {
	// check skip number
	skip := getWrapLv(skipNum...)

	var stack []*FuncDataStruct

	for i := skip + 1; ; i++ { // Skip the expected number of frames, add 1 to skip this function
		funcData := GetFuncData(i)
		if funcData == nil {
			break
		}

		stack = append(stack, funcData)
	}

	//skip runtime.main and exist
	return stack[:len(stack)-2]
}

func getWrapLv(skipNum ...int) int {
	// check skip number
	if len(skipNum) > 0 {
		return skipNum[0] + WrapLv
	} else {
		return WrapLv
	}
}