package main

import "fmt"

// slice
// 切片 1

func main() {
	demo21()
}

func demo21() {
	// 基本格式
	// 动态数组
	// 只定义不赋值 不分配内存空间 ==nil 
	// var s1 []int{} ,s1 := make([]int, 0) 都是有长度的，会开辟内存空间
	var s1 []int
	var s2 []string
	fmt.Printf("s1 is %d type is %T, s2 is %s type is %T\n", s1, s1, s2, s2)
	fmt.Printf("s1 == nil ? %t\n", s1 == nil)
	// 定义切片长度和容量 （类型，长度，容量）
	// s1 := make([]int, 5, 10)
	
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"北京", "上海", "深圳"}
	fmt.Printf("s1 is %d type is %T, s2 is %s type is %T\n", s1, s1, s2, s2)

	// 长度和容量
	fmt.Printf("s1 length is %d, capcity is %d\n", len(s1), cap(s1))
	fmt.Printf("s2 length is %d, capcity is %d\n", len(s2), cap(s2))

	// 由切割数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	// 左包含 右不包含
	s3 := a1[1:5]
	fmt.Printf("s3 is %d\n", s3)
	// 从底层数组的第零个开始 0 -> n
	s4 := a1[:4]
	fmt.Printf("s4 is %d\n", s4)
	// 到底层数组的最后一个 n -> len(s)
	s5 := a1[3:]
	fmt.Printf("s5 is %d\n", s5)
	// 从底层数组的第零个开始到底层数组的最后一个 0 -> len(s)
	s6 := a1[:]
	fmt.Printf("s6 is %d\n", s6)
	// 切片 长度等于元素的个数 容量底层数组从切片的第零个元素到底层数组的最后一个元素
	// 可以扩容
	fmt.Printf("s3 length is %d, capcity is %d\n", len(s3), cap(s3))
	fmt.Printf("s4 length is %d, capcity is %d\n", len(s4), cap(s4))
	fmt.Printf("s5 length is %d, capcity is %d\n", len(s5), cap(s5))
	fmt.Printf("s6 length is %d, capcity is %d\n", len(s6), cap(s6))
	s7 := s6[2:5]
	fmt.Printf("s7 is %d\n", s7)
	fmt.Printf("s7 length is %d, capcity is %d\n", len(s7), cap(s7))
	s8 := s5[3:]
	fmt.Printf("s8 is %d\n", s8)
	fmt.Printf("s8 length is %d, capcity is %d\n", len(s8), cap(s8))
	// 切片是引用类型（相对于值类型），指向底层数组
	a1[6] = 1300
	fmt.Printf("a1 is %d\n", a1)
	fmt.Printf("s5 is %d\n", s5)

	println("demo21!")
}
