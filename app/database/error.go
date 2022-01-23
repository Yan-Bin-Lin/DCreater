package database

import "errors"

var (
	ERR_TASK_FAIL     = errors.New("IL TO AFFECT ROW")
	ERR_NAME_CONFLICT = errors.New("NAME CONFLICT")
	ERR_PARAMETER     = errors.New("PARAMETER WRONG")
)