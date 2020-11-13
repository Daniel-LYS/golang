package main

import "fmt"

// array
// 多维数组

func main() {
	demo18()
}

func demo18() {
	// 基本格式 对比一维数组 var a3 = [3]int{1, 5, 9}
	var a1 [3][2]int
	a1 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println("a1 is ", a1)
	var a2 [3][4]int
	a2 = [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	fmt.Println("a2 is ", a2)

	// 遍历
	for i1, v1 := range a2 {
		fmt.Println(i1,v1)
		for i2,v2:=range v1{
			fmt.Println(i2,v2)
		}
	}

	print("demo18!!!")
}
