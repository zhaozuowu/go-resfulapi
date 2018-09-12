package main

import "fmt"

func test(cha chan int,num int ) {

	fmt.Printf("start channel\n")
	cha <- num
}
func main() {

	cha := make(chan  int)

	go test(cha,1)
	go test(cha,2)

	for value:= range cha{
		fmt.Printf("num is:%d\n",value)
	}

	// <-cha
	//cha := make(chan int,1)
	//for j :=0; j<10;j++ {
	//	j++
	//	select {
	//		case cha <- 0:
	//		case cha <- 1:
	//
	//	}
	//
	//	i := <-cha
	//
	//	fmt.Printf("value received:%d\n",i)
	//
	//
	//}
}
