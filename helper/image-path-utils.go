package helper

import (
	"net/url"
	"strconv"
	"strings"
)

func IsMD5Path(str string) bool {
	return regexpUrlParse.MatchString(str)
}

func SortPath(str []byte) string {

	// 对 byte 进行排序
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		for j := 1 + i; j < strLen; j++ {
			if str[i] > str[j] {
				str[i], str[j] = str[j], str[i]
			}
		}
	}
	var ret = strings.Builder{}

	for i := 0; i < strLen; i++ {
		ret.WriteString(strconv.Itoa(int(str[i])))
	}

	return ret.String()
}

func JoinPath(md5_str string) string {

	sortPath := SortPath([]byte(md5_str[:5]))

	var str = strings.Builder{}

	str.WriteString("img/") // path where to store the image
	str.WriteString(sortPath)
	str.WriteString("/")
	str.WriteString(md5_str)

	return str.String()

}

func UrlParse(md5_url string) string {

	if md5_url == "" {
		return ""
	}

	if len(md5_url) < 32 {
		return ""
	}

	// 进行 url 解析
	parse, err := url.Parse(md5_url)
	if err != nil {
		return ""
	}

	parsePath := parse.Path

	if len(parsePath) != 32 {
		return ""
	}

	if !IsMD5Path(parsePath) {
		return ""
	}

	return JoinPath(parsePath) + "/" + parsePath

}

func StringToInt(str string) int {
	if str == "" {
		return 0
	}

	toint, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	if toint < 0 {
		return 0
	}

	return toint
}
