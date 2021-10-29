package main

import (
	"fmt"
	"reflect"
)

func Receiver(v interface{}){
	a:=reflect.TypeOf(v)

    fmt.Println("这是一个",a)

}
