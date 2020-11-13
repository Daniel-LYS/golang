package main

import (
	"fmt"
	"strings"
)

// variable type
// char字符
// string 字符串
// 转义符

func main() {
	demo6()
}

func demo6() {
	// go 使用utf8编码 一个字符可以是数字 字母 汉字 ... 字符仅储存字节长度真正占用长度不同
	c1 := 'h'
	fmt.Printf("字符型变量%c 其类型为%T\n", c1, c1)

	s1 := "hello"
	fmt.Printf("s1 %s 其类型为%T\n", s1, s1)
	println("s1长度 is ", len(s1))

	// 转义符
	s2 := "hello"
	fmt.Printf("s2 %s \\n 其类型为%T\\\\\\\\\\\\\\\\\\\\\n", s2, s2)

	s3 := `
	hello
world
	`
	fmt.Printf("s3 %s 其类型为%T\n", s3, s3)

	// 拼接
	ss1 := s2 + s2
	println("ss1 is ", ss1)

	ss2 := fmt.Sprintln("give ss1 to ss2", ss1, s2)
	println("ss2 is ", ss2)

	// 分割
	ret1 := strings.Split(ss2, "ll")
	fmt.Println("ret1 is ", ret1)
	// 拼接
	fmt.Println(strings.Join(ret1, "$"))

	// 包含
	fmt.Println(strings.Contains(ss2, "l"))

	// 前缀
	fmt.Println(strings.HasPrefix(ss2, "x"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss2, "\n"))

	// 子串首次出现的位置
	fmt.Println(strings.Index(ss2, "g"))
	fmt.Println(strings.Index(ss2, "o"))
	// 子串最后一次出现的位置
	fmt.Println(strings.LastIndex(ss2, "o"))

	print("demo6!!!")
}
