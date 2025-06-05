package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// 1. Basic JSON Marshaling
type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Email    string    `json:"email"`
	Password string    `json:"-"` // Won't be included in JSON
	Created  time.Time `json:"created"`
}

// 2. Custom JSON Marshaling
type CustomUser struct {
	Name     string
	Age      int
	IsActive bool
}

func (u CustomUser) MarshalJSON() ([]byte, error) {
	type Alias CustomUser
	return json.Marshal(&struct {
		Alias
		Status string `json:"status"`
	}{
		Alias:  Alias(u),
		Status: "active",
	})
}

// 3. Nested JSON Structures
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type Employee struct {
	User    User    `json:"user"`
	Address Address `json:"address"`
	Role    string  `json:"role"`
	Salary  float64 `json:"salary"`
}

// 4. JSON with Omitempty
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

// 5. JSON with Custom Types
type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

type Account struct {
	ID     string `json:"id"`
	Status Status `json:"status"`
}

// 6. JSON with Interface
type Document struct {
	Type    string          `json:"type"`
	Content json.RawMessage `json:"content"`
}

// 7. JSON with Validation
type ValidatedUser struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=130"`
}

func main() {
	// 1. Basic JSON Marshaling
	fmt.Println("=== Basic JSON Marshaling ===")
	user := User{
		Name:     "John Doe",
		Age:      30,
		Email:    "john@example.com",
		Password: "secret",
		Created:  time.Now(),
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("Marshaled JSON:", string(jsonData))

	// Unmarshal
	var newUser User
	err = json.Unmarshal(jsonData, &newUser)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}
	fmt.Printf("Unmarshaled user: %+v\n", newUser)

	// 2. Custom JSON Marshaling
	fmt.Println("\n=== Custom JSON Marshaling ===")
	customUser := CustomUser{
		Name:     "Jane Doe",
		Age:      25,
		IsActive: true,
	}

	jsonData, err = json.Marshal(customUser)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("Custom marshaled JSON:", string(jsonData))

	// 3. Nested JSON Structures
	fmt.Println("\n=== Nested JSON Structures ===")
	employee := Employee{
		User: User{
			Name:     "Bob Smith",
			Age:      35,
			Email:    "bob@company.com",
			Created:  time.Now(),
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
		},
		Role:   "Developer",
		Salary: 75000,
	}

	jsonData, err = json.Marshal(employee)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("Employee JSON:", string(jsonData))

	// 4. JSON with Omitempty
	fmt.Println("\n=== JSON with Omitempty ===")
	product := Product{
		ID:    "P001",
		Name:  "Laptop",
		Price: 999.99,
	}

	jsonData, err = json.Marshal(product)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("Product JSON (with omitempty):", string(jsonData))

	// 5. JSON with Custom Types
	fmt.Println("\n=== JSON with Custom Types ===")
	account := Account{
		ID:     "A001",
		Status: StatusActive,
	}

	jsonData, err = json.Marshal(account)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("Account JSON:", string(jsonData))

	// 6. JSON with Interface
	fmt.Println("\n=== JSON with Interface ===")
	document := Document{
		Type: "user",
		Content: json.RawMessage(`{
			"name": "Alice",
			"age": 25
		}`),
	}

	jsonData, err = json.Marshal(document)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("Document JSON:", string(jsonData))

	// 7. Pretty Printing
	fmt.Println("\n=== Pretty Printing ===")
	jsonData, err = json.MarshalIndent(employee, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("Pretty printed JSON:")
	fmt.Println(string(jsonData))

	// 8. JSON with Validation
	fmt.Println("\n=== JSON with Validation ===")
	validatedUser := ValidatedUser{
		Name:  "John",
		Email: "john@example.com",
		Age:   30,
	}

	jsonData, err = json.Marshal(validatedUser)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("Validated user JSON:", string(jsonData))
} 