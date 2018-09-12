package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func poping(pings <-chan  string,poping chan <- string)  {

	msg := <-pings

	poping <- msg
}
func main()  {

	pings := make(chan string)
	popings := make(chan string)

	ping(pings,"hello")
	poping(pings,popings)
	fmt.Println(<-popings)


}
