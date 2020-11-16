// 九九乘法表
package main

import "fmt"

func main() {
	var (
		x = 1
		y = 1
	)
	// for
	for i := x; i < 10; i++ {
		for v := y; v < 10; v++ {
			fmt.Printf("%d * %d = %d ", i, v, i*v)
			if i == v {
				break
			}
		}
		fmt.Printf("\n")
	}
}
