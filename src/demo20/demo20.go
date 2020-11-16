package main

func main() {
	a1 := [5]int{1, 3, 5, 7, 8}
	f1(a1)
	f2(a1)
}

// 求数组[1, 3, 5, 7, 8]所有元素的和
func f1(k1 [5]int) {
	var i1 int

	for _, v := range k1 {
		i1 += v
	}
	println("sum is ", i1)
	println("f1 finsh!")
}

// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
func f2(k1 [5]int) {
	flag := 8
	for i1, v1 := range k1 {
		for i2 := i1 + 1; i2 < len(k1); i2++ {
			if (i1[])
		}
	}

	println("f2 finsh!")
}
