package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func (worker *Worker) process(c chan  int) {

	for {
		msg := <-c
		fmt.Printf("woker %d got %d\n",worker.id,msg)
	}

}
func main()  {

	c := make(chan int)
	for i :=0;i<5;i++ {
		worker := Worker{id:i}
		go worker.process(c)
	}

	for {
		num := rand.Int()
		c <- num
		time.Sleep(time.Second*1)
	}
}