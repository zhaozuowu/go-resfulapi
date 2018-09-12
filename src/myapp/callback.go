package main

import "fmt"

type callback func(int) bool


func isOdd(num int ) bool  {

	if num % 2 == 0 {
		return true
	}

	return false
	
}

func isEven(num int ) bool  {

	if num % 2 == 0 {
		return false
	}
	return true
}

func filter(arr []int,cal callback) []int {

	var  result []int

	for _,val :=  range  arr {

		if (cal(val)) {
			result = append(result,val)
		}
	}

	return result

}

func main()  {

	arr := []int{1,2,3,4,5,6,7,8}

	result := filter(arr,isOdd)

	fmt.Printf("result type:%#v\n",result)
	result2 := filter(arr,isEven)

	fmt.Printf("result2 type:%#v\n",result2)




}
