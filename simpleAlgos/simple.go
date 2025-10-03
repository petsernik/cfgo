package main

import (
	"cmp"
	"fmt"
	"sort"
)

// return first that >= target in Asc slice
func lowerBoundAsc[T cmp.Ordered](arr []T, target T) int {
	return sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})
}

// return first that >= target in Asc slice
func upperBoundAsc[T cmp.Ordered](arr []T, target T) int {
	return sort.Search(len(arr), func(i int) bool {
		return arr[i] > target
	})
}

// return first that <= target in Desc slice
func lowerBoundDesc[T cmp.Ordered](arr []T, target T) int {
	return sort.Search(len(arr), func(i int) bool {
		return arr[i] <= target
	})
}

// return first that < target in Desc slice
func upperBoundDesc[T cmp.Ordered](arr []T, target T) int {
	return sort.Search(len(arr), func(i int) bool {
		return arr[i] < target
	})
}

func checkBounds(l, r int, a ...int) bool {
	for _, x := range a {
		if !(x >= l && x < r) {
			return false
		}
	}
	return true
}

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

// >= go1.21
func isMonotone[T cmp.Ordered](arr []T) bool {
	isAsc, isDesc := false, false
	n := len(arr)
	for i := range n - 1 {
		switch {
		case isAsc:
			if arr[i] > arr[i+1] {
				return false
			}
		case isDesc:
			if arr[i] < arr[i+1] {
				return false
			}
		default:
			if arr[i] < arr[i+1] {
				isAsc = true
			} else if arr[i] > arr[i+1] {
				isDesc = true
			}
		}
	}
	return true
}

// < go1.21
func isMonotoneInt64(arr []int64) bool {
	isAsc, isDesc := false, false
	n := len(arr)
	for i := range n - 1 {
		switch {
		case isAsc:
			if arr[i] > arr[i+1] {
				return false
			}
		case isDesc:
			if arr[i] < arr[i+1] {
				return false
			}
		default:
			if arr[i] < arr[i+1] {
				isAsc = true
			} else if arr[i] > arr[i+1] {
				isDesc = true
			}
		}
	}
	return true
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
