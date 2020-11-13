package main

import "fmt"

// array
// 初始化

func main() {
	demo16()
}

func demo16() {
	var a1 [6]bool
	fmt.Println("a1 is ", a1)
	var a2 [3]int
	fmt.Println("a2 is ", a2)
	// go中数组是一种带类型的变量 该变量以元素类型和容量作为类型
	fmt.Printf("a1 type is %T\n", a1) // 类型是 存放3个bool型
	fmt.Printf("a2 type is %T\n", a2)

	// 基本格式
	var a3 = [3]bool{true, false, true}
	fmt.Println("a3 is ", a3)
	a4 := [5]float64{0, 1, 2, 3, 4}
	fmt.Println("a4 is ", a4)
	a5 := [...]float64{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("a5 is ", a5)
	a6 := [10]float64{4: 9, 7: 11}
	fmt.Println("a6 is ", a6)

	// 可以留白
	a7 := [5]int{1}
	fmt.Println("a7 is ", a7)

	print("demo16!!!")
}
