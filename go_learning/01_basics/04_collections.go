package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1. Arrays
	fmt.Println("=== Arrays ===")
	// Fixed-length array
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println("Array:", arr)

	// Array with initial values
	arr2 := [3]string{"apple", "banana", "orange"}
	fmt.Println("Array with values:", arr2)

	// Array with size inference
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array with inferred size:", arr3)

	// 2. Slices
	fmt.Println("\n=== Slices ===")
	// Creating slices
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice1)

	// Using make
	slice2 := make([]int, 5)    // length and capacity = 5
	slice3 := make([]int, 0, 5) // length = 0, capacity = 5
	fmt.Println("Slice with make:", slice2)
	fmt.Println("Slice with make (len=0, cap=5):", slice3)

	// Slice operations
	slice1 = append(slice1, 6, 7, 8)
	fmt.Println("After append:", slice1)

	// Slice slicing
	subSlice := slice1[2:5]
	fmt.Println("Sub-slice:", subSlice)

	// Copying slices
	slice4 := make([]int, len(slice1))
	copy(slice4, slice1)
	fmt.Println("Copied slice:", slice4)

	// 3. Maps
	fmt.Println("\n=== Maps ===")
	// Creating maps
	dict1 := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	fmt.Println("Map:", dict1)

	// Using make
	dict2 := make(map[string]int)
	dict2["one"] = 1
	dict2["two"] = 2
	fmt.Println("Map with make:", dict2)

	// Map operations
	// Adding/updating
	dict2["three"] = 3
	fmt.Println("After adding:", dict2)

	// Checking existence
	if value, exists := dict2["two"]; exists {
		fmt.Println("Value exists:", value)
	}

	// Deleting
	delete(dict2, "one")
	fmt.Println("After deleting:", dict2)

	// 4. Range with Collections
	fmt.Println("\n=== Range with Collections ===")
	// Range with slice
	fmt.Println("Range with slice:")
	for index, value := range slice1 {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Range with map
	fmt.Println("\nRange with map:")
	for key, value := range dict1 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// 5. Sorting
	fmt.Println("\n=== Sorting ===")
	// Sorting slice
	unsorted := []int{5, 2, 8, 1, 9}
	sort.Ints(unsorted)
	fmt.Println("Sorted slice:", unsorted)

	// Sorting map keys
	keys := make([]string, 0, len(dict1))
	for k := range dict1 {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("Sorted map keys:", keys)

	// 6. Slice Tricks
	fmt.Println("\n=== Slice Tricks ===")
	// Removing element
	slice := []int{1, 2, 3, 4, 5}
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("After removing element:", slice)

	// Filtering
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenNumbers []int
	for _, n := range numbers {
		if n%2 == 0 {
			evenNumbers = append(evenNumbers, n)
		}
	}
	fmt.Println("Even numbers:", evenNumbers)

	// 7. Map Operations
	fmt.Println("\n=== Map Operations ===")
	// Map of slices
	studentScores := map[string][]int{
		"Alice": {85, 90, 88},
		"Bob":   {92, 87, 95},
	}
	fmt.Println("Student scores:", studentScores)

	// Nested maps
	config := map[string]map[string]string{
		"database": {
			"host": "localhost",
			"port": "5432",
		},
		"server": {
			"host": "0.0.0.0",
			"port": "8080",
		},
	}
	fmt.Println("Nested map:", config)
} 