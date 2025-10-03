package set

import "testing"

func TestSet(t *testing.T) {
	tests := []struct {
		name     string
		prepare  func(s *Set[int])
		check    func(s *Set[int]) bool
		expected bool
	}{
		{
			name:     "empty set has no elements",
			prepare:  func(s *Set[int]) {},
			check:    func(s *Set[int]) bool { return s.Exists(1) },
			expected: false,
		},
		{
			name: "insert single element exists",
			prepare: func(s *Set[int]) {
				s.Set(10)
			},
			check:    func(s *Set[int]) bool { return s.Exists(10) },
			expected: true,
		},
		{
			name: "insert then erase element",
			prepare: func(s *Set[int]) {
				s.Set(42)
				s.Erase(42)
			},
			check:    func(s *Set[int]) bool { return s.Exists(42) },
			expected: false,
		},
		{
			name: "double insert element still exists",
			prepare: func(s *Set[int]) {
				s.Set(7)
				s.Set(7)
			},
			check:    func(s *Set[int]) bool { return s.Exists(7) },
			expected: true,
		},
		{
			name: "erase non-existing element does nothing",
			prepare: func(s *Set[int]) {
				s.Erase(123)
			},
			check:    func(s *Set[int]) bool { return true }, // just should not panic
			expected: true,
		},
		{
			name: "multiple elements inserted",
			prepare: func(s *Set[int]) {
				for _, v := range []int{1, 2, 3, 4, 5} {
					s.Set(v)
				}
			},
			check:    func(s *Set[int]) bool { return s.Exists(3) && s.Exists(5) && !s.Exists(10) },
			expected: true,
		},
		{
			name: "insert and erase subset of elements",
			prepare: func(s *Set[int]) {
				for _, v := range []int{1, 2, 3, 4} {
					s.Set(v)
				}
				s.Erase(2)
				s.Erase(4)
			},
			check: func(s *Set[int]) bool {
				return s.Exists(1) && !s.Exists(2) && s.Exists(3) && !s.Exists(4)
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := New[int]()
			tc.prepare(s)
			got := tc.check(s)
			if got != tc.expected {
				t.Errorf("%s: got %v, want %v", tc.name, got, tc.expected)
			}
		})
	}
}
