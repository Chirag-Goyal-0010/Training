package main

import (
	"fmt"
	"strings"
)

// 1. Basic Function
func greet(name string) string {
	return "Hello, " + name + "!"
}

// 2. Multiple Return Values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// 3. Named Return Values
func calculate(x, y int) (sum, product int) {
	sum = x + y
	product = x * y
	return // naked return
}

// 4. Variadic Functions
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 5. Function as a Type
type MathOperation func(int, int) int

// 6. Function that takes a function as parameter
func applyOperation(x, y int, operation MathOperation) int {
	return operation(x, y)
}

// 7. Anonymous Functions
var multiply = func(x, y int) int {
	return x * y
}

// 8. Closure
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 9. Method
type Person struct {
	Name string
	Age  int
}

func (p Person) introduce() string {
	return fmt.Sprintf("Hi, I'm %s and I'm %d years old.", p.Name, p.Age)
}

// 10. Method with Pointer Receiver
func (p *Person) haveBirthday() {
	p.Age++
}

// 11. Interface
type Speaker interface {
	Speak() string
}

func (p Person) Speak() string {
	return p.introduce()
}

// 12. Function that takes an interface
func makeSpeak(s Speaker) string {
	return s.Speak()
}

func main() {
	// 1. Basic Function
	fmt.Println("=== Basic Function ===")
	fmt.Println(greet("John"))

	// 2. Multiple Return Values
	fmt.Println("\n=== Multiple Return Values ===")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// 3. Named Return Values
	fmt.Println("\n=== Named Return Values ===")
	sum, product := calculate(5, 3)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)

	// 4. Variadic Functions
	fmt.Println("\n=== Variadic Functions ===")
	fmt.Println("Sum of 1, 2, 3, 4, 5:", sum(1, 2, 3, 4, 5))

	// 5. Function as a Type
	fmt.Println("\n=== Function as a Type ===")
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Result:", applyOperation(5, 3, add))

	// 6. Anonymous Functions
	fmt.Println("\n=== Anonymous Functions ===")
	fmt.Println("Multiply 4 and 5:", multiply(4, 5))

	// 7. Closure
	fmt.Println("\n=== Closure ===")
	counter := counter()
	fmt.Println("Count:", counter())
	fmt.Println("Count:", counter())
	fmt.Println("Count:", counter())

	// 8. Method
	fmt.Println("\n=== Method ===")
	person := Person{Name: "Alice", Age: 25}
	fmt.Println(person.introduce())

	// 9. Method with Pointer Receiver
	fmt.Println("\n=== Method with Pointer Receiver ===")
	person.haveBirthday()
	fmt.Println("After birthday:", person.introduce())

	// 10. Interface
	fmt.Println("\n=== Interface ===")
	fmt.Println(makeSpeak(person))

	// 11. Function Literal
	fmt.Println("\n=== Function Literal ===")
	transform := func(s string) string {
		return strings.ToUpper(s)
	}
	fmt.Println(transform("hello"))

	// 12. Defer
	fmt.Println("\n=== Defer ===")
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")

	// 13. Panic and Recover
	fmt.Println("\n=== Panic and Recover ===")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		panic("Something went wrong!")
	}()
} 