package pkg

import "strings"

// StrToPascalCase 将下划线表名或字段名转为大驼峰
func StrToPascalCase(str string) string {
	var res string
	list := strings.Split(str, "_")
	for _, v := range list {
		bytes := []byte(v)
		firstLetter := bytes[:1]
		otherLetters := bytes[1:]
		res += strings.ToUpper(string(firstLetter)) + string(otherLetters)
	}
	return res
}
