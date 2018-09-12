package main

import "fmt"

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "one"
	}()

	go func() {
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("recevice:%s\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("recevice:%s\n", msg2)
		}
	}
}
