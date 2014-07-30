package demo

import (
	"reflect"
	"testing"
)

func TestValueOfInt(t *testing.T) {
	i := 42
	v := reflect.ValueOf(i)

	inspect(v)
	inspect(v.Int())
	inspect(v.Interface())
	inspect(v.Bool())
}

func TestValueOfString(t *testing.T) {
	s := "42"
	v := reflect.ValueOf(s)

	inspect(v)
}

func TestValueOfBool(t *testing.T) {
	s := true
	v := reflect.ValueOf(s)

	if v.Bool() {
		inspect(v)
	}
}
