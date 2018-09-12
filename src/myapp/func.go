package main

import "fmt"

func SumAndProduct(a, b int) (result1, result2 int) {

	result1 = a + b
	result2 = a * b
	return
}


func sumTotal(args ...map[string]string)   {


	for _,argItem := range args {

		fmt.Printf("argItem:%#v\n",argItem)
	}
}

func arrSlice(args ...[3]int)  {

	fmt.Printf("arrSlice type :%#v\n",args)
}
func main() {

	a, b := 2, 3
	x, y := SumAndProduct(a, b)
	//sli1,sli2 := []int{1,2,3,4,5},[]int{6,7,8,9,10}
	sli1,sli2 := map[string]string{"username":"zhangsan","age":"27","sex":"男"},map[string]string{"username":"hanhong","age":"30","sex":"女"}
	fmt.Printf("%d+%d=%d\n", a, b, x)
	fmt.Printf("%d*%d=%d\n", a, b, y)
	sumTotal(sli1,sli2)

	orderList := []map[string]string{map[string]string{"orderId":"1","orderName":"小米手机"},map[string]string{"orderId":"2","orderName":"iPhone手机"}}
	fmt.Printf("orderList type :%#v\n",orderList)

	arr1 := [2][3]int{[3]int{1,2,3},[3]int{4,5,6}}

	fmt.Printf("arr1 type :%#v\n",arr1)

	slice1 := [][]int{[]int{7,8,9},[]int{9,10,11,12},[]int{13}}

	arrSlice1,arrSlice2 := [3]int{100,200,300},[3]int{400,500,600}

	fmt.Printf("slice1 type:%#v\n=----------------------\n",slice1)

	arrSlice(arrSlice1,arrSlice2)


}
