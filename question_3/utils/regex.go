package utils

import (
	"regexp"
)

// ใช้ regex ในการคัดแยกข้อความเนื้อออกมาผ่าน pattern ที่เตรียมไว้
func RegexFindAllString(words, pattern string) []string {
	regex := regexp.MustCompile(pattern)
	cleanWord := regex.FindAllString(words, -1)
	return cleanWord
}
