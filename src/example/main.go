package main

import (
	"fmt"
	"time"
	"runtime"
)

var num int = 0;

func f(str string, ch chan int) {

	for i := 1; i <= 3; i++ {

		fmt.Printf("%s\n", str)
	}

	ch <- 1

}
var count int = 0

func main() {

	//c1 := make(chan int)
	//go f("goroutime",c1)
	//
	//go func(str string) {
	//	fmt.Printf("%s\n", str)
	//	c1 <- 2
	//
	//}("doning")
	//
	//<-c1
	//<-c1

	//ch := make(chan int)
	//
	//go func() {
	//
	//	ch <- 1
	//}()
	//
	//msg := <-ch
	//fmt.Println("go run %s\n",msg)

	//ch := make(chan int)
	//
	//
	//go func() {
	//
	//	msg := <-ch
	//	fmt.Println("go run %s\n",msg)
	//}()
	//
	//ch <- 1

	//ch :=  make(chan int,3)
	//
	//ch <- 1
	//ch <- 2
	//ch <- 3
	//
	//go func() {
	//	msg := <-ch
	//
	//	fmt.Println(msg,"\n")
	//}()
	//ch <- 4
	//ch := make(chan string,1)
	//ch2 := make(chan string,1)
	//
	// sendChannel(ch,"hello")
	//
	//recevieChannel(ch,ch2)
	//
	//
	//fmt.Println(<-ch2)

	//c1 := make(chan int)
	//c2 := make(chan int)
	//
	//go func() {
	//	c1 <- 20
	//
	//}()
	//
	//go func() {
	//	msg := <-c2
	//
	//	fmt.Println(msg,"\n")
	//
	//}()
	//
	//for i :=0 ; i<2;i++ {
	//	select {
	//	case msg := <-c1:
	//		fmt.Println(msg,"\n")
	//	case c2 <- 10:
	//
	//	}
	//}

	//ch1 := make(chan bool)
	//
	//go func() {
	//
	//	fmt.Println("START:\n")
	//	time.Sleep(time.Second*2)
	//	ch1 <- true
	//
	//}()
	//
	//select {
	//	case msg := <- ch1:
	//		fmt.Println("result:",msg,"\n")
	//
	//	case <-time.After(time.Second * 1):
	//		fmt.Println("time out result1\n")
	//}
	//
	//
	//ch2 := make(chan  bool)
	//
	//go func() {
	//
	//	fmt.Println("START2:\n")
	//	time.Sleep(time.Second*2)
	//	ch2 <- true
	//}()
	//
	//select {
	//case msg := <- ch2:
	//	fmt.Println("result2:",msg,"\n")
	//
	//case <-time.After(time.Second * 3):
	//	fmt.Println("time out result1\n")
	//}

	//ch1 := make(chan string)
	//ch2 := make(chan string)
	//
	//select {
	//	case msg := <-ch1:
	//		fmt.Println("recieved messsage:",msg,"\n")
	//default:
	//	fmt.Println("no recieved message\n")
	//}
	//
	//select {
	//	case ch2 <- "hello":
	//		fmt.Println("send message:hello\n")
	//default:
	//	fmt.Println("no send messsage\n")
	//}

	//ch1 := make(chan int,5)
	//ch2 := make(chan bool)
	//
	//go func() {
	//
	//	for element := range ch1 {
	//
	//		fmt.Printf("recevied job:%d\n",element)
	//	}
	//
	//	ch2 <- true
	//}()
	//
	//for i :=0 ;i<3;i++ {
	//	ch1 <- i
	//	fmt.Printf("send job:%d\n",i)
	//
	//}
	//
	////close(ch1)
	//<-ch2

	//c1 := make(chan int)
	//
	//go func() {
	//
	//	for i:=0;i<3;i++ {
	//		c1 <- i
	//		fmt.Printf("send job:%d\n",i)
	//	}
	//
	//	close(c1)
	//}()
	//
	//for element := range c1 {
	//
	//	fmt.Printf("recevied job:%d\n",element)
	//}

	//c1 := make(chan bool)
	//
	//go func() {
	//
	//	fmt.Println("GO GO GO \n")
	//
	//	c1 <- true
	//	c1 <- false
	//	close(c1)
	//
	//}()
	//
	//for value := range c1 {
	//
	//	fmt.Println(value,"\n")
	//}
	//c1 := make(chan int,2)
	//
	//go func() {
	//	fmt.Println("GO GO GO !")
	//
	//	//<-c1
	//}()
	//
	//c1 <- 1
	//c1 <- 2
	//
	//fmt.Println("END END  !")

	//runtime.GOMAXPROCS(1000)
	//	//c1 := make(chan bool,10)
	//	//
	//	//for i :=0; i<10;i++ {
	//	//
	//	//	go Sum(c1,i)
	//	//}
	//	//
	//	//for i :=0;i<10;i++ {
	//	//	<-c1
	//	//}

	//runtime.GOMAXPROCS(runtime.NumCPU())

	//ch1 ,ch2 := make(chan int),make(chan string)
	//
	//go func() {
	//	for {
	//		select {
	//			case value,ok := <-ch1:
	//				if !ok {
	//
	//					break;
	//				}
	//
	//				fmt.Println("c1:",value,"\n")
	//		case value,ok := <-ch2:
	//			if !ok {
	//
	//				break;
	//			}
	//
	//			fmt.Println("c2:",value,"\n")
	//		}
	//	}
	//}()
	//
	//
	//ch1 <- 1
	//
	//ch2 <- "hello"

	//c1 := make(chan string)
	//
	//go func() {
	//
	//	fmt.Println("START")
	//
	//	time.Sleep(time.Second * 2)
	//
	//	c1 <- "hello"
	//}()
	//
	//select {
	//	case msg := <-c1 :
	//		fmt.Println("result:",msg,"\n")
	//	case <-time.After(time.Second * 1):
	//		fmt.Println("time out result")
	//}

	//theMine := []string{"one", "two", "three"}
	//
	//oreChan := make(chan string)
	//
	//go func(theMine []string) {
	//	for _, itme := range theMine {
	//
	//		oreChan <- itme
	//	}
	//}(theMine)
	//
	//go func() {
	//
	//	for {
	//
	//		select {
	//		case value, ok := <-oreChan:
	//			if !ok {
	//				break;
	//			}
	//
	//			fmt.Println("Miner: Received",value,"from finder\n")
	//		}
	//	}
	//}()
	//
	//<-time.After(time.Second * 5)

	//ch := make(chan string,3)
	//
	//go func() {
	//
	//	ch <- "First"
	//	fmt.Println("SEND FIRST\n")
	//
	//	ch<-"SECOND"
	//	fmt.Println("SEND SECOND\n")
	//
	//	ch <- "THIRD"
	//	fmt.Println("SEND THIRD\n")
	//
	//}()


	//Receivig first,SEND FIRST SEND SECOND SECOND  THIRD SEND THIRD
	//SEND FIRST RECEVING first seoncd send second send third third

	//go func() {
	//	num := 1
	//	for {
	//		select {
	//		case value ,ok := <-ch:
	//
	//			if !ok {
	//				break
	//			}
	//
	//			fmt.Println(num,":Receiving:",value)
	//			num++
	//		}
	//	}
	//
	//}()
	//go func() {
	//	firstRead := <- ch
	//	fmt.Println("Receiving..")
	//	fmt.Println(firstRead)
	//	secondRead := <- ch
	//	fmt.Println(secondRead)
	//	thirdRead := <- ch
	//	fmt.Println(thirdRead)
	//}()

	//<-time.After(time.Second * 5)

	//ch := make(chan string)
	//
	//go func() {
	//	ch <- "message"
	//}()
	//
	//select {
	//	case msg := <-ch:
	//		fmt.Println(msg,"\n")
	//default:
	//	fmt.Println("no message\n")
	//}
	//
	//select {
	//case msg := <-ch:
	//	fmt.Println(msg,"\n")
	//default:
	//	fmt.Println("no message\n")
	//}

	//bufferedChan := make(chan string)
	//go func() {
	//	bufferedChan <- "first"
	//	fmt.Println("Sent 1st")
	//	bufferedChan <- "second"
	//	fmt.Println("Sent 2nd")
	//	bufferedChan <- "third"
	//	fmt.Println("Sent 3rd")
	//}()
	//go func() {
	//	firstRead := <- bufferedChan
	//	fmt.Println("Receiving..")
	//	fmt.Println(firstRead)
	//	secondRead := <- bufferedChan
	//	fmt.Println(secondRead)
	//	thirdRead := <- bufferedChan
	//	fmt.Println(thirdRead)
	//}()
	//
	//<-time.After(time.Second * 3)

	//cha := make(chan int)
	//
	//go func() {
	//
	//	time.Sleep(time.Second * 2)
	//	fmt.Println(" GET RESULT\n")
	//
	//	cha <- 1
	//}()
	//
	//L:
	//	for{
	//		select {
	//			case msg := <-cha:
	//				fmt.Println("msg:",msg,"\n")
	//			case <-time.After(time.Second * 1):
	//				fmt.Println("time out\n")
	//				break L
	//		}
	//	}


	runtime.GOMAXPROCS(runtime.NumCPU())
	for  i:=0; i<20;i++ {
		go incr()
	}

	time.Sleep(time.Second * 10)

}

func incr()  {

	count ++

	fmt.Println(count,"\n")
}

func Sum(cha chan bool, index int) {

	sum := 0

	for i := 0; i <= 100000; i++ {

		sum += i;
	}

	fmt.Println(index, sum, "\n")

	cha <- true
}

func sendChannel(ch chan<- string, msg string) {

	ch <- msg
}

func recevieChannel(re <-chan string, send chan<- string) {

	msg := <-re

	send <- msg

}
