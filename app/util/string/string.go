package string

import "regexp"

// match two or more space and replace to only one space
func CleanSpace(str string) string {
	reg := regexp.MustCompile("\\s\\s+")
	return reg.ReplaceAllString(str, " ")
}
