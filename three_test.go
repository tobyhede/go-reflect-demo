package demo

import (
	"reflect"
	"testing"
)

func fn(i int) int {
	return i + 1
}

func TestValueOfFn(t *testing.T) {
	v := reflect.ValueOf(fn)

	display(v.Type().Out(0))
	display(v.Type().In(0))

	// i := v.Call(1)
	// display(i)

	i := v.Call([]reflect.Value{reflect.ValueOf(1)})
	display(i[0].Int())
}
