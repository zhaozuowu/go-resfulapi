package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width,height float64
}


type Square struct {
	radius float64
}

func (r Rectangle) area() float64  {

	return r.width * r.height
}

func (r Square) area() float64  {

	return  r.radius * r.radius * math.Pi
}


func area(r Rectangle)  float64 {

	return r.width * r.height
}

func main()  {
	r := Rectangle{5.4,5}

	area := area(r)

	fmt.Printf("area is :%#v\n",area)

	r1 := Rectangle{10,5}
	r2 := Rectangle{20,5}
	r3 := Square{10}
	r4 := Square{20}

	fmt.Printf("r1:%v,r2:%v,r3:%v,r4:%v\n",r1.area(),r2.area(),r3.area(),r4.area())
}