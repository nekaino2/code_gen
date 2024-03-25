package main

import (
	"fmt"
	"reflect"
)

type pi struct {
	k string
	s int
	D D
}
type D struct {
	k string
	l int
	N N
}
type N struct {
	o string
	k int
}

func main() {
	parseStruct(pi{})
}

func parseStruct(inter interface{}) {
	v := reflect.ValueOf(inter)
	fmt.Println(v)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Struct {
			fmt.Println(field)
			fmt.Printf("%s has a struct type\n", v.Type().Field(i).Name)
			parseStruct(field.Interface())
		}
	}
	return
}
