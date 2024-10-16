package algo

// Package gosort provides implementations of various sorting algorithms.
// This package includes Bubble Sort, Insertion Sort, Selection Sort,
// and Quick Sort algorithms to sort arrays of integers in ascending order.

// BubbleSort sorts an array of integers using the bubble sort algorithm.
// This algorithm repeatedly steps through the array, compares adjacent
// elements, and swaps them if they are in the wrong order.
// The process is repeated until the array is sorted.
//
// The average and worst-case time complexity of this algorithm is O(n^2),
// making it inefficient for large datasets. However, it is simple and
// easy to understand, making it useful for educational purposes.
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap elements if they are in the wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// InsertionSort sorts an array of integers using the insertion sort algorithm.
// This algorithm builds a sorted array one element at a time. It iterates
// through the input array and inserts each element into its correct position
// in the sorted part of the array.
//
// The average and worst-case time complexity of this algorithm is O(n^2),
// but it performs well on small or partially sorted datasets.
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		// Move elements of arr[0..i-1] that are greater than key,
		// to one position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

// SelectionSort sorts an array of integers using the selection sort algorithm.
// This algorithm divides the array into a sorted and an unsorted part.
// It repeatedly selects the minimum element from the unsorted part and
// moves it to the end of the sorted part.
//
// The average and worst-case time complexity of this algorithm is O(n^2).
// It is not suitable for large datasets but can be useful in specific
// situations where memory writes are costly.
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		// Find the index of the minimum element in the unsorted part
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Swap the found minimum element with the first element
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// QuickSort sorts an array of integers using the quicksort algorithm.
// This algorithm works by selecting a "pivot" element from the array
// and partitioning the other elements into two subarrays: those less than
// the pivot and those greater than the pivot. The subarrays are then
// sorted recursively. This makes quicksort one of the most efficient
// sorting algorithms for large datasets.
//
// The average time complexity of this algorithm is O(n log n),
// while the worst-case time complexity is O(n^2) when the pivot is
// poorly chosen. However, with proper pivot selection (e.g., using
// the median), the performance can be significantly improved.
func QuickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the pivot index
		pi := partition(arr, low, high)

		// Recursively sort elements before and after partition
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

// partition is a helper function for QuickSort.
// It selects the last element as the pivot, places the pivot at its correct
// position in the sorted array, and places all elements smaller than the pivot
// to the left and all elements larger to the right.
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
