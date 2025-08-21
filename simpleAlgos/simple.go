package main

import (
	"fmt"
)

func binpow(x, n int64) int64 {
	var res int64 = 1
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}

func divUp(a, b int64) int64 {
	if a^b > 0 && a%b != 0 {
		return a/b + 1
	}
	return a / b
}

func divDown(a, b int64) int64 {
	if a^b < 0 && a%b != 0 {
		return a/b - 1
	}
	return a / b
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a %= b
		a, b = b, a
	}
	return a
}

func main() {
	fmt.Println(binpow(2, 30))
	fmt.Println(binpow(3, 5))
	dUp := func(a, b int64) {
		fmt.Printf("divUp %d / %d : %d\n", a, b, divUp(a, b))
	}
	dDown := func(a, b int64) {
		fmt.Printf("divDown %d / %d : %d\n", a, b, divDown(a, b))
	}
	for _, f := range [2]func(a, b int64){dUp, dDown} {
		f(1, 2)
		f(2, 2)
		f(1, -2)
		f(2, -2)
		f(-1, 2)
		f(-2, 2)
		f(-1, -2)
		f(-2, -2)
	}
	fmt.Println(gcd(35*9, 3*25*49))
}
