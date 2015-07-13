package katas

import "regexp"

func Unsplice(str string) string {
	r, _ := regexp.Compile(`\\\n`)
	return r.ReplaceAllLiteralString(str, "")
}
