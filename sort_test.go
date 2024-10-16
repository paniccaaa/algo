package algo

import (
	"reflect"
	"testing"
)

// TestBubbleSort tests the BubbleSort function.
func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 8, 4, 2}, []int{2, 3, 4, 5, 8}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{[]int{}, []int{}},
		{[]int{7}, []int{7}},
	}

	for _, test := range tests {
		BubbleSort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("BubbleSort(%v) = %v; want %v", test.input, test.input, test.expected)
		}
	}
}

// TestInsertionSort tests the InsertionSort function.
func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 8, 4, 2}, []int{2, 3, 4, 5, 8}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{[]int{}, []int{}},
		{[]int{7}, []int{7}},
	}

	for _, test := range tests {
		InsertionSort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("InsertionSort(%v) = %v; want %v", test.input, test.input, test.expected)
		}
	}
}

// TestSelectionSort tests the SelectionSort function.
func TestSelectionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 8, 4, 2}, []int{2, 3, 4, 5, 8}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{[]int{}, []int{}},
		{[]int{7}, []int{7}},
	}

	for _, test := range tests {
		SelectionSort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("SelectionSort(%v) = %v; want %v", test.input, test.input, test.expected)
		}
	}
}

// TestQuickSort tests the QuickSort function.
func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 8, 4, 2}, []int{2, 3, 4, 5, 8}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{[]int{}, []int{}},
		{[]int{7}, []int{7}},
	}

	for _, test := range tests {
		QuickSort(test.input, 0, len(test.input)-1)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("QuickSort(%v) = %v; want %v", test.input, test.input, test.expected)
		}
	}
}
