package main

import (
	"fmt"
	"time"
)

// 1. Basic Struct
type Person struct {
	Name     string
	Age      int
	Email    string
	Birthday time.Time
}

// 2. Nested Struct
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

type Employee struct {
	Person  // Embedded struct
	Address // Embedded struct
	ID      string
	Role    string
	Salary  float64
}

// 3. Methods
func (p Person) introduce() string {
	return fmt.Sprintf("Hi, I'm %s and I'm %d years old.", p.Name, p.Age)
}

// 4. Pointer Receiver Method
func (p *Person) haveBirthday() {
	p.Age++
}

// 5. Interface
type Speaker interface {
	Speak() string
}

type Listener interface {
	Listen() string
}

// 6. Interface Composition
type Communicator interface {
	Speaker
	Listener
}

// 7. Implementing Interfaces
func (p Person) Speak() string {
	return p.introduce()
}

func (p Person) Listen() string {
	return fmt.Sprintf("%s is listening", p.Name)
}

// 8. Type Assertion
func processInterface(i interface{}) {
	switch v := i.(type) {
	case Person:
		fmt.Println("It's a Person:", v.Name)
	case Employee:
		fmt.Println("It's an Employee:", v.ID)
	default:
		fmt.Println("Unknown type")
	}
}

// 9. Empty Interface
func printAnything(v interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

func main() {
	// 1. Creating and using structs
	fmt.Println("=== Basic Struct ===")
	person := Person{
		Name:     "John Doe",
		Age:      30,
		Email:    "john@example.com",
		Birthday: time.Date(1993, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	fmt.Println(person.introduce())

	// 2. Nested structs
	fmt.Println("\n=== Nested Struct ===")
	employee := Employee{
		Person: Person{
			Name:  "Jane Smith",
			Age:   28,
			Email: "jane@company.com",
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
		},
		ID:     "EMP001",
		Role:   "Developer",
		Salary: 75000,
	}
	fmt.Printf("Employee: %s, Role: %s\n", employee.Name, employee.Role)

	// 3. Methods
	fmt.Println("\n=== Methods ===")
	fmt.Println(person.introduce())

	// 4. Pointer receiver
	fmt.Println("\n=== Pointer Receiver ===")
	person.haveBirthday()
	fmt.Println("After birthday:", person.introduce())

	// 5. Interface usage
	fmt.Println("\n=== Interface Usage ===")
	var speaker Speaker = person
	fmt.Println(speaker.Speak())

	// 6. Interface composition
	fmt.Println("\n=== Interface Composition ===")
	var communicator Communicator = person
	fmt.Println(communicator.Speak())
	fmt.Println(communicator.Listen())

	// 7. Type assertion
	fmt.Println("\n=== Type Assertion ===")
	processInterface(person)
	processInterface(employee)

	// 8. Empty interface
	fmt.Println("\n=== Empty Interface ===")
	printAnything(42)
	printAnything("Hello")
	printAnything(true)
	printAnything(person)

	// 9. Struct tags (commonly used with JSON)
	type User struct {
		Name     string `json:"name" validate:"required"`
		Age      int    `json:"age" validate:"gte=0,lte=130"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"-"` // Won't be included in JSON
	}

	// 10. Method sets
	fmt.Println("\n=== Method Sets ===")
	// Value receiver methods can be called on both values and pointers
	// Pointer receiver methods can only be called on pointers
	p1 := Person{Name: "Alice", Age: 25}
	p2 := &Person{Name: "Bob", Age: 30}

	fmt.Println(p1.introduce())  // Value receiver
	fmt.Println(p2.introduce())  // Value receiver
	p2.haveBirthday()           // Pointer receiver
	fmt.Println(p2.introduce())
} 