package conv

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// convert CamelCase ke snake_case
func CamelCaseToSnakeCase(str string) string {
	regex := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := regex.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

func NameStructToStringSnakeCase(v interface{}, additionalNameRight string) string {
	t := reflect.TypeOf(v)

	// Kalau slice/array ambil elemennya
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		t = t.Elem()
	}

	if t.Kind() == reflect.Struct {
		if additionalNameRight != "" {
			return fmt.Sprintf("%s_%s", additionalNameRight, CamelCaseToSnakeCase(t.Name()))
		}
		return CamelCaseToSnakeCase(t.Name())
	}

	return ""
}

func StrToTime(str string) time.Duration {
	timeConv, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return time.Duration(int64(timeConv))
}

func IntToTime(str int) time.Duration {
	return time.Duration(int64(str))
}
