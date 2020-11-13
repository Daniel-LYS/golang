package main

import "fmt"

// if

func main() {
	demo10()
}

func demo10() {
	i1 := 20

	if i1 > 0 {
		fmt.Println("must born")
	}

	if i1 > 18 {
		fmt.Println("older than 18")
	} else {
		fmt.Println("little!")
	}

	if i1 < 35 {
		fmt.Println("youth")
	} else if i1 < 55 {
		fmt.Println("middle age")
	} else {
		fmt.Println("old")
	}

	// i2作用域在if条件语句
	if i2 := 3; i2 > 1 {
		fmt.Println(i2)
	}
	//fmt.Println(i2) 超过作用域

	print("demo10!!!")
}
