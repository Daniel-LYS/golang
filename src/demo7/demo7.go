package main

import "fmt"

// variable type
// bool 布尔型

func main() {
	demo7()
}

func demo7() {
	// go默认布尔型不能与其他类型进行强制转换
	b1 := true
	fmt.Printf("b1 is %t 其类型为%T\n", b1, b1)

	// 默认为false
	var b2 bool
	fmt.Printf("b2 is %t\n", b2)

	print("demo7!!!")
}
