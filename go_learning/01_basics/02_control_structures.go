package main

import "fmt"

func main() {
	// 1. If-Else Statements
	fmt.Println("=== If-Else Statements ===")
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// If with initialization
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// 2. Switch Statement
	fmt.Println("\n=== Switch Statement ===")
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the week")
	default:
		fmt.Println("Middle of the week")
	}

	// Switch with multiple cases
	month := 3
	switch month {
	case 1, 2, 3:
		fmt.Println("Spring")
	case 4, 5, 6:
		fmt.Println("Summer")
	case 7, 8, 9:
		fmt.Println("Fall")
	case 10, 11, 12:
		fmt.Println("Winter")
	default:
		fmt.Println("Invalid month")
	}

	// Switch with fallthrough
	fmt.Println("\n=== Switch with Fallthrough ===")
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	}

	// 3. For Loop
	fmt.Println("\n=== For Loop ===")
	// Basic for loop
	for i := 1; i <= 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// While-like loop
	fmt.Println("\n=== While-like Loop ===")
	count := 1
	for count <= 3 {
		fmt.Printf("Count: %d\n", count)
		count++
	}

	// Infinite loop with break
	fmt.Println("\n=== Infinite Loop with Break ===")
	counter := 1
	for {
		if counter > 3 {
			break
		}
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	// 4. For-Range Loop
	fmt.Println("\n=== For-Range Loop ===")
	// Range with slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Range with map
	person := map[string]string{
		"name":  "John",
		"city":  "New York",
		"state": "NY",
	}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// 5. Continue Statement
	fmt.Println("\n=== Continue Statement ===")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Printf("Number: %d\n", i)
	}

	// 6. Nested Loops
	fmt.Println("\n=== Nested Loops ===")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}

	// 7. Label and Goto (rarely used, but available)
	fmt.Println("\n=== Label and Goto ===")
	i := 0
Start:
	if i < 3 {
		fmt.Printf("i: %d\n", i)
		i++
		goto Start
	}
} 