package check

import (
	"fmt"
	"reflect"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)

	switch t.Kind() {
	case reflect.Struct:
		fmt.Println("Struct:", t)
	case reflect.Slice, reflect.Array:
		elem := t.Elem()
		if elem.Kind() == reflect.Struct {
			fmt.Println("Slice/Array of Struct:", elem)
		} else {
			fmt.Println("Slice/Array of:", elem)
		}
	default:
		fmt.Println("Bukan struct atau array/slice of struct:", t)
	}
}

// check if array struct or not
func SliceStruct(v interface{}) bool {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		if t.Elem().Kind() == reflect.Struct {
			return true
		}
	}
	return false
}

// check if array  or not
func Slice(v interface{}) bool {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		return true
	}
	return false
}
