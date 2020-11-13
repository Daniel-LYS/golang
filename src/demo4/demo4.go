package main

// gobal variable

func main() {
	demo4()
}

func demo4() {
	println("pi is ", pi)
	println("statusOK is ", statusOK)
	println("notfound is ", notfound)

	// 常量可以赋值后不使用
	println("批量赋值常量，常量可以赋值后不使用")
	println("c2 is ", c2)
	println("c3 is ", c3)

	println("iota 计数器")
	println("i1 is ", i1)
	println("i2 is ", i2)
	println("i3 is ", i3)

	println("iota 计数器 空出一行 仍然加一")
	println("a1 is ", a1)
	println("a2 is ", a2)
	println("a3 is ", a3)

	println("iota 计数器 其中两行插入一百")
	println("b1 is ", b1)
	println("b2 is ", b2)
	println("b3 is ", b3)
	println("b4 is ", b4)
	println("b5 is ", b5)

	println("iota 计数器 每新增一行常量声明计数一次")
	println("d1 is ", d1)
	println("d2 is ", d2)
	println("d3 is ", d3)
	println("d4 is ", d4)

	println("定义数量级")
	println("KB is ", KB)
	println("MB is ", MB)
	println("GB is ", GB)
	println("TB is ", TB)
	println("PB is ", PB)

	print("demo4!!!")
}

const pi = 3.14

const (
	statusOK = 200
	notfound = 404
)

// 批量声明常量，如果某一行声明后没有赋值，默认和上一行一致
const (
	c1 = 50
	c2
	c3
)

// iota 常量计算器 遇到关键字const后 从零开始计数
const (
	i1 = iota
	i2
	i3
)

// 每新增一行常量声明计数一次
const (
	a1 = iota
	a2
	_
	a3
)

const (
	b1 = iota
	b2 = 100
	b3
	b4 = iota
	b5
)

const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
