package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func sum(num int, ch chan int) {

	result := 0

	for i := 0; i < result; i++ {

		result += i
	}
	fmt.Println("starting pro\n");
	ch <- result

}

func main() {

	//ch := make(chan int)
	//num := 100
	//
	//go sum(num, ch)
	//fmt.Println("beign pro\n");
	//
	//result := <-ch
	//
	//fmt.Printf("result is :%d\n",result)
	app := iris.New()

	app.Get("users", func(cxt iris.Context) {

		cxt.HTML("<p>hello</p>")
	})

	app.Run(iris.Addr(":8080"))
}
