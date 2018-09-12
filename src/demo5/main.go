package main

import (
	"time"
	"fmt"
)

func main() {

	c1 := make(chan string)

	go func() {

		time.Sleep(time.Second * 2)

		c1 <- "result1"
	}()

	select {
	case result := <-c1:
		fmt.Printf("result1:%s\n", result)
	case <- time.After(time.Second*1):
		fmt.Printf("time out 1\n")
	}

	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second*2)
		c2 <- "result2"
	}()

	select {
	case result := <-c2:
		fmt.Printf("result2:%s\n", result)
	case <- time.After(time.Second*3):
		fmt.Printf("time out 1\n")
	}
}
