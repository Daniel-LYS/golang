package main

import "fmt"

// variable type
// int 整型

func main() {
	demo5()
}

func demo5() {
	// 定义当前计算机默认类型的整型
	i1 := 101
	fmt.Printf("十进制%d 其类型为%T ", i1, i1)
	//go中无法直接定义二进制数
	// 通过先定义其他进制的变量再在打印输出转义成二进制成为go中可以看见二进制的方法
	fmt.Printf("该数二进制数为%b 其类型仍然为%T\n", i1, i1)

	i2 := 0101
	fmt.Printf("八进制%o 其十进制是%d 其类型为%T\n", i2, i2, i2)

	i3 := 0x101
	fmt.Printf("十六进制%x 其十进制是%d 其类型为%T\n", i3, i3, i3)

	// 定义其他类型的整型
	// 非默认整型 需要明确指定类型 占据1个字节的有符号型整型
	i4 := int8(101)
	fmt.Printf("八位有符号整型的十进制%d 其类型为%T\n", i4, i4)

	print("demo5!!!")
}
