package main

import "fmt"

func test(num *int) (result int) {

	*num = *num  * 10

	result = *num

	return


}

func main() {


	b := 10
	c := test(&b)

	fmt.Printf("b is :%v,c is :%v",b,c)
}
