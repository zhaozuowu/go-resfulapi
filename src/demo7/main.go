package main

import "fmt"

func main() {

	c1 := make(chan string, 2)
	c1 <- "message"
	c1 <- "hello"
	close(c1)
	for item := range c1 {
		fmt.Println(item)
	}

}
