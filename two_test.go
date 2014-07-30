package demo

import (
	"bytes"
	"reflect"
	"testing"
)

var s = []string{"t", "o", "b", "y", "i", "s", "a", "w", "e", "s", "o", "m", "e"}

var i = []int{1, 2, 3, 4, 5}

func TestValueOfSlice(t *testing.T) {
	v := reflect.ValueOf(s)

	inspect(v)
}

func TestIterateSlice(t *testing.T) {
	v := reflect.ValueOf(s)

	var buffer bytes.Buffer

	for i := 0; i < v.Len(); i++ {
		c := v.Index(i)
		// buffer.WriteString(c)
		// buffer.WriteString(c.String())
		buffer.WriteString(c.Interface().(string))
	}

	out := buffer.String()
	display(out)
}

func TestMakeSlice(t *testing.T) {
	v := reflect.ValueOf(i)

	display(v.Type())
	out := reflect.MakeSlice(v.Type(), v.Len(), v.Len())

	display(out.Type())
	for i := 0; i < v.Len(); i++ {
		c := v.Index(i)
		if (c.Int() % 2) == 1 {
			out = reflect.Append(out, c)
			display(c.Int())
		}
	}

	display(out)
	display(out.Index(0).Int())
}
