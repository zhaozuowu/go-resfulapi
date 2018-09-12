package main

import "fmt"

func main()  {

	c1 := make(chan  string)
	c2 := make(chan string)

	select {
		case msg := <-c1:
			fmt.Printf("recieved msg:%s\n",msg)
	default:
		fmt.Printf("no message recevied\n")
	}

	msg := "hi"

	select {
		case c1 <- msg :
			fmt.Printf("send msg:%s\n",msg)
	default:
		fmt.Printf("no send message\n")
	}

	select {
		case msg := <-c1 :
			fmt.Printf("recevied c1:%s\n",msg)
		case sig := <- c2:
			fmt.Printf("recevied c2:%s\n",sig)
	default:
		fmt.Printf("no recevied\n")
	}
}
