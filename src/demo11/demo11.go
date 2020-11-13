package main

import "fmt"

// for 1

func main() {
	demo11()
}

func demo11() {
	// 基本格式
	for i1 := 1; i1 < 10; i1++ {
		fmt.Println(i1)
	}

	i2 := 5
	for ; i2 < 10; i2++ {
		fmt.Println(i2)
	}

	i3 := 5
	for i3 < 10 {
		fmt.Println(i2)
		i3++
	}

	// for {
	// 	fmt.Println("123")
	// }

	// 键值循环 一种特殊的迭代结构
	s1 := "hello你好"
	for i, v := range s1 {
		fmt.Printf("%d %c\n", i, v)
	}

	// go中为了处理非ASCII码类型的字符 定义了rune类型
	// go中字符串不能修改
	s2 := "白萝卜"
	fmt.Println(s2)
	s3 := []rune(s2) // 转换成rune切片 每一份是一个字符型变量
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "黑"
	c2 := '黑' // utf8字符
	fmt.Printf("c1 is %s, it is %T, c2 is %c, it is %T, and value is %d\n", c1, c1, c2, c2, c2)
	c3 := byte('h')
	fmt.Printf("c3 is %c, it is %T\n", c3, c3)

	print("demo11!!!")
}
