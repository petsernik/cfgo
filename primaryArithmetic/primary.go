package main

import "fmt"

const mod = int64(1e9) + 7

func powMod(x, n int64) int64 {
	if n < 0 {
		return powMod(getInvMod(x), -n)
	}
	var res int64 = 1
	for n > 0 {
		if n&1 == 1 {
			res = multMod(res, x)
		}
		x = multMod(x, x)
		n >>= 1
	}
	return res
}

func getMod(x int64) int64 {
	if x >= 0 {
		if x < mod {
			return x
		}
		return x % mod
	}
	if t := mod + x%mod; t < mod {
		return t
	}
	return 0
}

// PRIMES ONLY: Fermat's little theorem
func getInvMod(x int64) int64 {
	return powMod(x, mod-2)
}

func multMod(x, y int64) int64 {
	return getMod(getMod(x) * getMod(y))
}

func divMod(x, y int64) int64 {
	return multMod(x, getInvMod(y))
}

func main() {
	fmt.Println(powMod(2, 10))
	fmt.Println(divMod(1, 2))
}
