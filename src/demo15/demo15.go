package main

import "fmt"

// 运算符
// ++ -- 不是运算符

func main() {
	demo15()
}

func demo15() {
	var (
		i1 = 5
		i2 = 2
		i3 = true
		i4 = false
		i5 = 9
		i6 = 1
	)
	fmt.Println("i1 = ", i1)
	fmt.Println("i2 = ", i2)
	fmt.Println("i3 = ", i3)
	fmt.Println("i4 = ", i4)
	fmt.Println("i5 = ", i5)
	fmt.Println("i6 = ", i6)

	fmt.Println("i1 + i2 = ", i1+i2)
	fmt.Println("i1 - i2 = ", i1-i2)
	fmt.Println("i1 * i2 = ", i1*i2)
	fmt.Println("i1 / i2 = ", i1/i2)
	fmt.Println("i1 % i2 = ", i1%i2)

	fmt.Println("i1 == i2 = ", i1 == i2)
	fmt.Println("i1 != i2 = ", i1 != i2)
	fmt.Println("i1 >= i2 = ", i1 >= i2)
	fmt.Println("i1 <= i2 = ", i1 <= i2)
	fmt.Println("i1 > i2 = ", i1 > i2)
	fmt.Println("i1 < i2 = ", i1 < i2)

	fmt.Println("i3 && i4 = ", i3 && i4)
	fmt.Println("i3 || i4 = ", i3 || i4)
	// 取反
	fmt.Println("!i3 = ", !i3)
	fmt.Println("i3 = ", i3)

	// 5 二进制 0101
	// 2 二进制 0010
	fmt.Println("i1 = ", i1)
	fmt.Println("i2 = ", i2)
	fmt.Println("i1 & i2 = ", i1&i2) // 0000 同1为1
	fmt.Println("i1 | i2 = ", i1|i2) // 0111 有1为1
	fmt.Println("i1 ^ i2 = ", i1^i2) // 0111 不同为1
	fmt.Println("i1 << 1 = ", i1<<1) // 1010
	fmt.Println("i1 >> 1 = ", i1>>1) // 0010

	// i5 += 1
	// fmt.Println("i5 += ", i5) // i5 = i5 + 1 可以优化成i5++
	// i5 -= 1
	// fmt.Println("i5 -= ", i5) // i5 = i5 - 1 可以优化成i5--
	i5 *= 1
	fmt.Println("i5 *= ", i5) // i5 = i5 * 1
	i5 /= 1
	fmt.Println("i5 /= ", i5) // i5 = i5 / 1
	i5 %= 1
	fmt.Println("i5 %= ", i5) // i5 = i5 % 1
	i5 <<= 1
	fmt.Println("i5 <<= ", i5) // i5 = i5 << 1
	i5 >>= 1
	fmt.Println("i5 >>= ", i5) // i5 = i5 >> 1
	i5 &= 1
	fmt.Println("i5 &= ", i5) // i5 = i5 & 1
	i5 |= 1
	fmt.Println("i5 |= ", i5) // i5 = i5 | 1
	i5 ^= 1
	fmt.Println("i5 ^= ", i5) // i5 = i5 ^ 1

	// 在go中不算运算符 是单独的语句
	i1--
	i2++
	fmt.Println("i1-- = ", i1)
	fmt.Println("i2++ = ", i2)

	print("demo15!!!")
}
