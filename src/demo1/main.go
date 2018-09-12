package main

import "fmt"

func say(str string)  {

	for i :=1;i<=3;i++ {

		fmt.Println(str,":",i,"\n")
	}
}
func main()  {

	 say("start")

	 go say("goroutime")
	 go say("goning")

	 //go func(str string) {
	 //	fmt.Println(str,"\n")
	 //}("goning")
	 mes := make(chan string)


	go func() {
		mes <- "like"
		mes <- "pong"
		}()
	msg := <-mes

	fmt.Println(msg,"\n")

	 fmt.Println("done")
	
}
