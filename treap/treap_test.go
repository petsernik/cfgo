package main

import (
	"math/rand"
	"testing"
)

func newTestTreap() *Treap[int] {
	return newTreap[int](nil)
}

func TestInsertFind(t *testing.T) {
	tr := newTestTreap()
	values := []int{5, 2, 8, 1, 3, 7, 10}

	for _, v := range values {
		tr.Insert(v)
	}

	for _, v := range values {
		got := tr.Find(v)
		if got == nil || *got != v {
			t.Errorf("Find(%d) = %v, want %d", v, got, v)
		}
	}

	if got := tr.Find(100); got != nil {
		t.Errorf("Find(100) = %v, want nil", *got)
	}
}

func TestErase(t *testing.T) {
	tr := newTestTreap()
	values := []int{5, 2, 8, 1, 3, 7, 10}
	for _, v := range values {
		tr.Insert(v)
	}

	tr.Erase(2)
	if got := tr.Find(2); got != nil {
		t.Errorf("Erase(2) failed, Find(2) = %v", *got)
	}

	tr.Erase(5)
	if got := tr.Find(5); got != nil {
		t.Errorf("Erase(5) failed, Find(5) = %v", *got)
	}

	// Ensure others remain
	for _, v := range []int{1, 3, 7, 8, 10} {
		if got := tr.Find(v); got == nil {
			t.Errorf("Value %d should still exist", v)
		}
	}
}

func TestBounds(t *testing.T) {
	tr := newTestTreap()
	for _, v := range []int{1, 3, 5, 7, 9} {
		tr.Insert(v)
	}

	if lb := tr.LowerBound(6); lb == nil || *lb != 7 {
		t.Errorf("LowerBound(6) = %v, want 7", lb)
	}
	if lb := tr.LowerBound(0); lb == nil || *lb != 1 {
		t.Errorf("LowerBound(0) = %v, want 1", lb)
	}
	if lb := tr.LowerBound(10); lb != nil {
		t.Errorf("LowerBound(10) = %v, want nil", *lb)
	}

	if ub := tr.UpperBound(5); ub == nil || *ub != 7 {
		t.Errorf("UpperBound(5) = %v, want 7", ub)
	}
	if ub := tr.UpperBound(9); ub != nil {
		t.Errorf("UpperBound(9) = %v, want nil", *ub)
	}
}

func TestMinMax(t *testing.T) {
	tr := newTestTreap()
	for _, v := range []int{5, 2, 8, 1, 3, 7, 10} {
		tr.Insert(v)
	}

	if min := tr.Min(); min == nil || *min != 1 {
		t.Errorf("Min() = %v, want 1", min)
	}
	if max := tr.Max(); max == nil || *max != 10 {
		t.Errorf("Max() = %v, want 10", max)
	}
}

func TestSize(t *testing.T) {
	tr := newTestTreap()
	if tr.Size() != 0 {
		t.Errorf("Size() = %d, want 0", tr.Size())
	}
	tr.Insert(1)
	tr.Insert(2)
	tr.Insert(3)
	if tr.Size() != 3 {
		t.Errorf("Size() = %d, want 3", tr.Size())
	}
	tr.Erase(2)
	if tr.Size() != 2 {
		t.Errorf("Size() = %d, want 2", tr.Size())
	}
}

func BenchmarkTreapInsert(b *testing.B) {
	tr := newTreap[int](rand.New(rand.NewSource(1)))
	r := rand.New(rand.NewSource(1))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tr.Insert(r.Intn(b.N))
	}
}

func BenchmarkMapInsert(b *testing.B) {
	m := make(map[int]struct{})
	r := rand.New(rand.NewSource(1))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m[r.Intn(b.N)] = struct{}{}
	}
}

func BenchmarkTreapFind(b *testing.B) {
	tr := newTreap[int](rand.New(rand.NewSource(1)))
	r := rand.New(rand.NewSource(1))
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		x := r.Intn(b.N)
		data[i] = x
		tr.Insert(x)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = tr.Find(data[i])
	}
}

func BenchmarkMapFind(b *testing.B) {
	m := make(map[int]struct{})
	r := rand.New(rand.NewSource(1))
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		x := r.Intn(b.N)
		data[i] = x
		m[x] = struct{}{}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = m[data[i]]
	}
}

func BenchmarkTreapOps1e5(b *testing.B) {
	const N = 100_000
	for i := 0; i < b.N; i++ {
		tr := newTreap[int](rand.New(rand.NewSource(1)))
		r := rand.New(rand.NewSource(1))

		// Insert N элементов
		for j := 0; j < N; j++ {
			tr.Insert(r.Intn(N * 10))
		}

		// Find N случайных элементов
		for j := 0; j < N; j++ {
			_ = tr.Find(r.Intn(N * 10))
		}

		// Erase N случайных элементов
		for j := 0; j < N; j++ {
			tr.Erase(r.Intn(N * 10))
		}
	}
}
