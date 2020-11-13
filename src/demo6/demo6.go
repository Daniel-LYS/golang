package main

import "fmt"

// variable type
// float 浮点型

func main() {
	demo6()
}

func demo6() {
	// go默认浮点型 float64
	f1 := 1.23
	fmt.Printf("浮点型变量%f 其类型为%T\n", f1, f1)

	// 不同类型的变量不能互相赋值
	f2 := float32(1.23)
	fmt.Printf("浮点型变量%f 其类型为%T\n", f2, f2)

	i3 := 0x101
	fmt.Printf("十六进制%x 其十进制是%d 其类型为%T\n", i3, i3, i3)

	// 定义其他类型的整型
	// 非默认整型 需要明确指定类型 占据1个字节的有符号型整型
	i4 := int8(101)
	fmt.Printf("八位有符号整型的十进制%d 其类型为%T\n", i4, i4)

	print("demo6!!!")
}
