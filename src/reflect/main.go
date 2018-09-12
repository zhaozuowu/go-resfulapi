package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age int
}

func main()  {

	num := 10
	t := reflect.TypeOf(num)
	v := reflect.ValueOf(num)
	s := reflect.TypeOf(Person{})
	fmt.Println("num type:\n",t)
	fmt.Println("num valueof:\n",v.Type())
	fmt.Println("num valueof:\n",s)
}