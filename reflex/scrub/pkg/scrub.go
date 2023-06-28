package pkg

import (
	"fmt"
	"reflect"
)

func Scrub(s interface{}, fields map[string]string) interface{} {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return s
	}

	for k, replaceStr := range fields {
		f := v.FieldByName(k)
		if f.IsValid() && f.CanSet() && f.Kind() == reflect.String {
			fieldLen := len([]rune(f.String()))
			replacedValue := ""
			for i := 0; i < fieldLen; i++ {
				replacedValue += replaceStr
			}
			f.Set(reflect.ValueOf(replacedValue))
		} else {
			fmt.Println("skip field", k)
			fmt.Println("field info", f.IsValid(), f.CanSet(), f.Kind())
		}
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		if reflect.Indirect(f).Kind() == reflect.Struct {
			if v.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			Scrub(f.Interface(), fields)
		}
	}

	return s
}
