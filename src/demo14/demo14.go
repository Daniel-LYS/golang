package main

import "fmt"

// switch case

func main() {
	demo14()
}

func demo14() {
	// 基本格式
	// case 最后自带 break 语句，匹配成功后就不会执行其他 case
	i1 := 5
	switch i1 {
	case 1:
		println("1")
	case 2:
		println("2")
	case 3:
		println("3")
	case 4:
		println("4")
	case 5:
		println("5")
	default:
		println("default")
	}

	// i2出了该switch后就销毁
	switch i2 := 6; i2 {
	case 1:
		println("1")
	case 2:
		println("2")
	case 3:
		println("3")
	case 4:
		println("4")
	case 5:
		println("5")
	default:
		println("default")
	}

	switch i3 := 6; i3 {
	case 1, 3, 5, 7, 9:
		println("奇数")
	case 2, 4, 6, 8:
		println("偶数")
	default:
		println("default")
	}

	i4 := 30
	switch {
	case i4 < 25:
		fmt.Println("好好学习吧")
	case i4 > 25 && i4 < 35:
		fmt.Println("好好工作吧")
	case i4 > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	i5 := 10
	fmt.Printf("i5 type is %T", i5)

	print("demo14!!!")
}
