package debug

import (
	"os"
	"testing"
)

var (
	pwd string
)

func init() {
	pwd, _ = os.Getwd()
}

func wrap1() []*FuncDataStruct {
	return wrap2()
}

func wrap2() []*FuncDataStruct {
	return wrap3()
}

func wrap3() []*FuncDataStruct {
	return GetCallStack()
}

func TestDebug_GetCallStack(t *testing.T) {
	funcDatas := wrap1()
	t.Log(funcDatas[0].Function)
	t.Log(funcDatas[0].Loacate)

	var expect = []FuncDataStruct{
		{pwd + "/stack_test.go:25", "app/util/debug.wrap3"},
		{pwd + "/stack_test.go:21", "app/util/debug.wrap2"},
		{pwd + "/stack_test.go:17", "app/util/debug.wrap1"},
		{pwd + "/stack_test.go:29", "app/util/debug.TestDebug_GetCallStack"},
	}

	for i, l := 0, len(funcDatas); i < l; i++ {
		if funcDatas[i].Function != expect[i].Function || funcDatas[i].Loacate != expect[i].Loacate {
			t.Fatalf("wrong result. expect function: %s, Locate: %s. Get function: %s, Locate: %s",
				expect[i].Function, expect[i].Loacate, funcDatas[i].Function, funcDatas[i].Loacate)
		}
	}
}

func TestDebug_GetFuncData(t *testing.T) {
	funcData := GetFuncData()

	var expect = &FuncDataStruct{pwd + "/stack_test.go:49", "app/util/debug.TestDebug_GetFuncData"}

	if funcData.Function != expect.Function || funcData.Loacate != expect.Loacate {
		t.Fatalf("wrong result. expect function: %s, Locate: %s. Get function: %s, Locate: %s",
			expect.Function, expect.Loacate, funcData.Function, funcData.Loacate)
	}
}