package zhconv

import (
	"errors"
	"regexp"
	"strings"
)

const (
	data_Chinese_tas = "chinese"
)

var traditionalChinese map[string]string
var simplifiedChinese map[string]string
var multiPinyinChinese map[string]string
var pinyinChinese map[string]string
var chineseRegex *regexp.Regexp

func init() {
	chineseRegex = regexp.MustCompile("[\u4e00-\u9fa5]")
}

func loadResource(filename string, dest map[string]string, reverse bool) error {
	var str string

	if filename == data_Chinese_tas {
		str = chinese_db
	} else {
		return errors.New("error")
	}

	if len(str) == 0 {
		return errors.New("Resource load failed.")
	}
	if !strings.HasSuffix(str, "\n") {
		return errors.New("Config file does not end with a newline character.")
	}
	re := regexp.MustCompile(".*=.*")
	s2 := re.FindAllString(str, -1)

	for _, tempStr := range s2 {
		arr := strings.Split(tempStr, "=")
		if reverse {
			dest[arr[1]] = arr[0]
		} else {
			dest[arr[0]] = arr[1]
		}
	}

	return nil
}

func ConvertToSimplifiedChinese(source string) string {
	result := ""
	for _, runeValue := range source {
		result += toSimplifiedChinese(string(runeValue))
	}
	return result
}

func ConvertToTraditionalChinese(source string) string {
	result := ""
	for _, runeValue := range source {
		result += toTraditionalChinese(string(runeValue))
	}
	return result
}

func loadMapFromResource(resourceName string, reverse bool) map[string]string {
	v := make(map[string]string)
	err := loadResource(resourceName, v, reverse)
	if err != nil {
		panic(err)
	}
	return v
}

func toSimplifiedChinese(source string) string {
	if simplifiedChinese == nil {
		simplifiedChinese = loadMapFromResource(data_Chinese_tas, false)
	}
	v := simplifiedChinese[source]
	if len(v) == 0 {
		return source
	}
	return v
}

func toTraditionalChinese(source string) string {
	if traditionalChinese == nil {
		traditionalChinese = loadMapFromResource(data_Chinese_tas, true)
	}
	v := traditionalChinese[source]
	if len(v) == 0 {
		return source
	}
	return v
}

func IsChinese(char string) bool {
	return chineseRegex.MatchString(char)
}
