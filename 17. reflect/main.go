package main

import (
	"fmt"
	"reflect"
)


func main() {
    b := 12
    intValue := reflect.TypeOf(b)

    fmt.Println("Value of Int: ", reflect.ValueOf(intValue).String())

    var floatValue float64 = 5.234

    fmt.Println("Value of Float: ", reflect.ValueOf(floatValue).String())

    st := map[interface{}]interface{}{"A": "Hello world", 1: 234, }

    fmt.Println("Value of Interface: ", reflect.ValueOf(st).String())
}
