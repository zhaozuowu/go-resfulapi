package main

import "fmt"

func Sum(num int,channel chan int)  {

	 sum := 0

	 for i :=1;i<=num;i++ {
	 	sum +=i
	 }

	 channel <- sum

	 //close(channel)

}
func main()  {

	cha := make(chan int)

	go Sum(10,cha)
	go Sum (20,cha)


	x,y := <-cha,<-cha
	fmt.Println(x,y,x+y)




}
