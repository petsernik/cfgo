package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func main() {
	tr := newTreap[int](nil)
	// вставка
	for _, x := range []int{5, 2, 8, 1, 3, 7, 10} {
		tr.Insert(x)
	}
	fmt.Println("size =", tr.Size())
	fmt.Println("min =", *tr.Min())
	fmt.Println("max =", *tr.Max())

	// поиск
	fmt.Println("find 7 =", tr.Find(7))
	fmt.Println("find 42 =", tr.Find(42))

	// нижняя и верхняя границы
	fmt.Println("lower_bound(6) =", tr.LowerBound(6))
	fmt.Println("upper_bound(6) =", tr.UpperBound(6))
	fmt.Println("upper_bound(10) =", tr.UpperBound(10))

	// удаление
	tr.Erase(1)
	tr.Erase(10)
	tr.Erase(5)

	fmt.Println("size after erases =", tr.Size())
	fmt.Println("min =", *tr.Min())
	fmt.Println("max =", *tr.Max())

	const N = 100_000
	tr = newTreap[int](nil)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// -------- Insert benchmark --------
	start := time.Now()
	for range N {
		tr.Insert(r.Intn(N * 10))
	}
	elapsed := time.Since(start)
	fmt.Printf("Inserted %d elements in %v (%.2f ops/sec)\n",
		N, elapsed, float64(N)/elapsed.Seconds())

	// -------- Find benchmark --------
	keys := make([]int, N)
	for i := range keys {
		keys[i] = r.Intn(N * 10)
	}
	slices.Sort(keys)
	start = time.Now()
	for i := range N {
		_ = tr.Find(keys[i])
	}
	elapsed = time.Since(start)
	fmt.Printf("Searched %d elements in %v (%.2f ops/sec)\n",
		N, elapsed, float64(N)/elapsed.Seconds())
}
