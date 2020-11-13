package main

import "fmt"

// 占位符

func main() {
	demo8()
}

func demo8() {
	// %v 任何类型变量的占位符
	f1 := 1.23
	fmt.Printf("浮点型变量%v 其类型为%T\n", f1, f1)

	s1 := "你好"
	fmt.Printf("字符串变量1%v 其类型为%T\n", s1, s1)

	// 带上#号的字符串会有描述符
	s2 := "你好"
	fmt.Printf("字符串变量2（带井号）%#v 其类型为%T\n", s2, s2)

	b1 := true
	fmt.Printf("整型变量%#v 其类型为%T\n", b1, b1)

	print("demo8!!!")
}
