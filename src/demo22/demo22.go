package main

import "fmt"

// slice
// 切片 2
func main() {
	demo22()
}

func demo22() {
	// 定义切片长度和容量 （类型，长度，容量）
	s1 := make([]int, 5, 10)
	fmt.Printf("s1 is %d, length is %d, capcity is %d\n", s1, len(s1), cap(s1))
	s2 := make([]int, 5)
	fmt.Printf("s2 is %d, length is %d, capcity is %d\n", s2, len(s2), cap(s2))
	s3 := make([]int, 0, 10)
	fmt.Printf("s3 is %d, length is %d, capcity is %d\n", s3, len(s3), cap(s3))

	for i, v := range s3 {
		fmt.Println(i, v)
	}

	println("demo22!")
}
