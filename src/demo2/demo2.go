package main

import "fmt"

// demo2 variable
func main() {
	println("before demo2() i1: ", i1)
	demo2_1()
	println("after demo2() i1: ", i1)
}

//全局变量可以只声明
var (
	u1 uint8
	i1 int = 1
	i2     = 2
)

func demo2_1() {
	println("in demo2() i1: ", i1)
	// 局部变量只声明不使用，不能通过
	var i1 int
	var (
		s1 string
		b1 bool
		f1 float64 = 64
		s2 string  = "?"
	)

	// 局部变量只声明和赋值不使用，也不能通过
	i1 = 18
	s1 = "hello"
	b1 = true

	fmt.Print("i1 is ", i1, "\n")
	// 使用占位符%
	fmt.Printf("s1 is %s, b1 is %t\n", s1, b1)
	// 自带换行
	fmt.Println("f1 is", f1, ", s2 is %s", s2)

	fmt.Println("1!")

	var s3 string = "string 3"
	println("s3 is ", s3)

	var s4 = "string 4"
	println("s4 is ", s4)

	// 只能在函数内使用
	s5 := "string 5"
	println("s5 is ", s5)
}
