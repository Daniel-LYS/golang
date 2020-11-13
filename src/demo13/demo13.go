package main

import "fmt"

// for 2

func main() {
	demo13()
}

func demo13() {
	// break 跳出for循环
	for i1 := 1; i1 < 10; i1++ {
		if i1 == 5 {
			break
		}
		fmt.Println(i1)
	}
	// continue 跳出本次循环
	for i2 := 1; i2 < 10; i2++ {
		if i2 == 5 {
			continue
		}
		fmt.Println(i2)
	}

	print("demo13!!!")
}
