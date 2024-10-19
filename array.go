package main

import (
	"fmt"
	"sort"
)

func main() {
	// Define an array of integers
	numbers := []int{10, 3, 5, 8, 1, 9, 2, 6, 7, 4}

	fmt.Println("Unsorted array:", numbers)

	// Sort the array
	sort.Ints(numbers)

	// Print the sorted array
	fmt.Println("Sorted array:", numbers)
}
