package segtree

import (
	"testing"
)

func getSum(l, r int, st *SegTree[int64], a []int64) (int64, int64) {
	var sum int64 = 0
	for i := l; i <= r; i++ {
		sum += a[i]
	}
	return st.Query(l, r), sum
}

func add(l, r int, x int64, st *SegTree[int64], a []int64) {
	st.Add(l, r, x)
	for i := l; i <= r; i++ {
		a[i] += x
	}
}

func change(l, r int, x int64, st *SegTree[int64], a []int64) {
	st.Set(l, r, x)
	for i := l; i <= r; i++ {
		a[i] = x
	}
}

func intAdd(current *int64, value int64, l int, r int) {
	*current += value * int64(r-l)
}

func intSet(current *int64, value int64, l int, r int) {
	*current = value * int64(r-l)
}

func intMerge(a, b int64) int64 {
	return a + b
}

func TestSegTreeComprehensive(t *testing.T) {
	// Тестовые данные
	data := []int64{1, 2, 3, 4, 5}
	neutral := int64(0)

	checkValues := func(t *testing.T, st *SegTree[int64], a []int64) {
		for i := range a {
			if sum, expected := getSum(i, i, st, a); sum != expected {
				t.Errorf("in pos %d want %d, got %d", i, expected, sum)
			}
		}
	}

	t.Run("Multiple operations", func(t *testing.T) {
		st := NewSegTree(data, intAdd, intSet, intMerge, neutral)
		newData := append([]int64{}, data...)

		add(1, 3, 5, st, newData)
		// change(2, 3, 0, st, newData)
		add(0, 4, 2, st, newData)

		// if sum, expected := getSum(0, 4, st, newData); sum != expected {
		// 	t.Errorf("Multiple operations expected %d, got %d", expected, sum)
		// }
		checkValues(t, st, newData)
	})

	t.Run("Basic operations", func(t *testing.T) {
		st := NewSegTree(data, intAdd, intSet, intMerge, neutral)
		newData := append([]int64{}, data...)

		// Проверка начальной суммы
		if sum, expected := getSum(0, 4, st, newData); sum != expected {
			t.Errorf("Initial sum expected %d, got %d", expected, sum)
		}

		// Добавление значения
		add(1, 3, 2, st, newData)
		if sum, expected := getSum(0, 4, st, newData); sum != expected {
			t.Errorf("After Add(1,3,2) expected %d, got %d", expected, sum)
		}

		// Установка значения
		change(2, 4, 10, st, newData)
		if sum, expected := getSum(0, 4, st, newData); sum != expected {
			t.Errorf("After Set(2,4,10) expected %d, got %d", expected, sum)
		}

		checkValues(t, st, newData)
	})

	t.Run("Mixed operations", func(t *testing.T) {
		st := NewSegTree(data, intAdd, intSet, intMerge, neutral)
		newData := append([]int64{}, data...)

		add(0, 4, 1, st, newData)
		//change(1, 3, 7, st, newData)
		// add(2, 4, 3, st, newData)

		if sum, expected := getSum(0, 4, st, newData); sum != expected {
			t.Errorf("Mixed operations expected %d, got %d", expected, sum)
		}
		checkValues(t, st, newData)
	})

	t.Run("Partial queries", func(t *testing.T) {
		st := NewSegTree(data, intAdd, intSet, intMerge, neutral)
		newData := append([]int64{}, data...)

		add(1, 3, 3, st, newData)
		change(2, 4, 8, st, newData)
		add(0, 2, 1, st, newData)

		testCases := []struct {
			l, r int
		}{
			{0, 2}, {2, 4}, {1, 3}, {0, 0}, {4, 4},
		}

		for _, tc := range testCases {
			if sum, expected := getSum(tc.l, tc.r, st, newData); sum != expected {
				t.Errorf("Query(%d,%d) expected %d, got %d", tc.l, tc.r, expected, sum)
			}
		}
		checkValues(t, st, newData)
	})

	t.Run("Edge cases", func(t *testing.T) {
		st := NewSegTree(data, intAdd, intSet, intMerge, neutral)
		newData := append([]int64{}, data...)

		// Пустой отрезок
		if sum := st.Query(3, 2); sum != neutral {
			t.Errorf("Empty segment should return neutral element %d, got %d", neutral, sum)
		}

		// Один элемент
		change(2, 2, 100, st, newData)
		if sum, expected := getSum(2, 2, st, newData); sum != expected {
			t.Errorf("Single element query expected %d, got %d", expected, sum)
		}

		// Весь массив
		change(0, 4, 42, st, newData)
		if sum, expected := getSum(0, 4, st, newData); sum != expected {
			t.Errorf("Full array set expected %d, got %d", expected, sum)
		}

		checkValues(t, st, newData)
	})
}
