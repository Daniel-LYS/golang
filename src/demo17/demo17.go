package main

import "fmt"

// array
// 遍历

func main() {
	demo17()
}

func demo17() {
	a1 := [...]string{"北京", "上海", "深圳"}
	// 索引遍历
	for i := 0; i < len(a1); i++ {
		fmt.Println(a1[i])
	}
	// 键值
	for i, v := range a1 {
		fmt.Println(i, v)
	}

	print("demo17!!!")
}
