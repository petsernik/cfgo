package list

import (
	"testing"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		left     int
		right    int
		expected []int
	}{
		{
			name:     "reverse entire list",
			input:    []int{1, 2, 3, 4, 5},
			left:     1,
			right:    5,
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "reverse middle part",
			input:    []int{1, 2, 3, 4, 5},
			left:     2,
			right:    4,
			expected: []int{1, 4, 3, 2, 5},
		},
		{
			name:     "reverse single element",
			input:    []int{1, 2, 3, 4, 5},
			left:     3,
			right:    3,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse first two elements",
			input:    []int{1, 2, 3, 4, 5},
			left:     1,
			right:    2,
			expected: []int{2, 1, 3, 4, 5},
		},
		{
			name:     "reverse last two elements",
			input:    []int{1, 2, 3, 4, 5},
			left:     4,
			right:    5,
			expected: []int{1, 2, 3, 5, 4},
		},
		{
			name:     "single element list",
			input:    []int{1},
			left:     1,
			right:    1,
			expected: []int{1},
		},
		{
			name:     "two element list reverse all",
			input:    []int{1, 2},
			left:     1,
			right:    2,
			expected: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create input list
			head := createList(tt.input)

			// Reverse between
			result := reverseBetween(head, tt.left, tt.right)

			// Convert result to slice for comparison
			resultSlice := listToSlice(result)

			// Compare with expected
			if !slicesEqual(resultSlice, tt.expected) {
				t.Errorf("reverseBetween(%v, %d, %d) = %v, expected %v",
					tt.input, tt.left, tt.right, resultSlice, tt.expected)
			}
		})
	}
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "reverse multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "reverse two elements",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "reverse single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "reverse empty list",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create input list
			head := createList(tt.input)

			// Reverse entire list
			result := reverseList(head)

			// Convert result to slice for comparison
			resultSlice := listToSlice(result)

			// Compare with expected
			if !slicesEqual(resultSlice, tt.expected) {
				t.Errorf("reverseList(%v) = %v, expected %v",
					tt.input, resultSlice, tt.expected)
			}
		})
	}
}

func TestAddAndDelNext(t *testing.T) {
	t.Run("add and delete operations", func(t *testing.T) {
		// Create initial list: 1 -> 2 -> 3
		head := &ListNode{Val: 1}
		node2 := head.add(2)
		node3 := node2.add(3)

		// Verify initial state
		if head.Val != 1 || head.Next.Val != 2 || head.Next.Next.Val != 3 {
			t.Errorf("Initial list creation failed")
		}

		// Delete next node from head (should remove node 2)
		head.delNext()

		// Verify deletion: 1 -> 3
		if head.Val != 1 || head.Next.Val != 3 || head.Next.Next != nil {
			t.Errorf("delNext() failed, expected 1->3, got %v", listToSlice(head))
		}

		// Add new node: 1 -> 3 -> 4
		node4 := node3.add(4)

		// Verify addition
		if node4.Val != 4 || node3.Next.Val != 4 {
			t.Errorf("add() failed, expected 4 at end")
		}
	})
}

// Helper function to create a list from slice
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current = current.add(values[i])
	}

	return head
}

// Helper function to convert list to slice
func listToSlice(head *ListNode) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

// Helper function to compare slices
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
