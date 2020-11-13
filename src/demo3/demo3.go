package main

// anonymous variable

func main() {
	demo3()
}

func demo3() {
	// s1, s2 := fanhui()
	// println("s1 is ", s1)
	// println("s2 is ", s2)
	s1, _ := fanhui()
	println("s1 is ", s1)

	print("demo3!!!")
}

func fanhui() (string, string) {
	return "返回到变量", "返回到匿名变量"
}
