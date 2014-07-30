package demo

import (
	"fmt"
)

func display(i interface{}) {
	fmt.Printf("%v\n", i)
}

func inspect(i interface{}) {
	fmt.Printf("%#v\n", i)
}
