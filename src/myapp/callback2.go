package main

import "fmt"

type callAddNum func([]int) int

func addSum (arr []int) int {

	sum := 0

	for _,val := range arr {
		sum+=val
	}
	return sum
}

func getSum(num int,call callAddNum)  {

	var sum []int

	for i:=0; i<=num;i++ {
		sum = append(sum,i)
	}

	result := call(sum)

	fmt.Printf("result type:#%v\n",result)

	
}

func main()  {

	getSum(10,addSum)
}

