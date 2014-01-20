package gosk

import (
	"log"
	"regexp"
)

func trimHTML(str string) string {
	if str == "" {
		return str
	}
	re, _ := regexp.Compile(`\<[\S\s]+?\>`)
	newstr := re.ReplaceAllString(str, "")
	return newstr
}

func subStr(str string, start, end int) string {
	if start < 0 {
		log.Panic("start position is wrong!")
	}
	if end > len(str) {
		log.Panic("end positon is wrong!")
	}
	if start > end {
		log.Panic("wrong position!")
	}

	rs := []rune(str)
	return string(rs[start:end])
}
