package sqlqb

import (
	"fmt"
	"reflect"
)

func MapStruct(s interface{}) []Element {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic(fmt.Errorf("invalid type %T", s))
	}
	return nil
}
