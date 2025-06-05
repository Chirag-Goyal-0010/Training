# Go Language Cheatsheet

## Table of Contents
1. [Basic Syntax](#basic-syntax)
2. [Data Types](#data-types)
3. [Control Structures](#control-structures)
4. [Functions](#functions)
5. [Structs and Interfaces](#structs-and-interfaces)
6. [Concurrency](#concurrency)
7. [Error Handling](#error-handling)
8. [File I/O and JSON](#file-io-and-json)
9. [Web Development](#web-development)
10. [Database Operations](#database-operations)
11. [Security and Authentication](#security-and-authentication)
12. [Deployment and DevOps](#deployment-and-devops)
13. [Generics](#generics)
14. [Best Practices](#best-practices)

## Basic Syntax

### Package Declaration
```go
package main
```

### Imports
```go
import (
    "fmt"
    "strings"
    "time"
    "encoding/json"
    "net/http"
    "database/sql"
    "crypto/tls"
    "context"
)
```

### Variables
```go
// Declaration
var name string
var age int = 25

// Short declaration
score := 95.5

// Multiple variables
var (
    firstName = "John"
    lastName  = "Doe"
)
```

## Data Types

### Basic Types
```go
// Numeric Types
int     // Platform dependent
int8    // -128 to 127
int16   // -32768 to 32767
int32   // -2147483648 to 2147483647
int64   // -9223372036854775808 to 9223372036854775807
uint    // Unsigned integer
float32 // 32-bit floating point
float64 // 64-bit floating point

// String
string  // UTF-8 encoded string

// Boolean
bool    // true or false

// Complex
complex64  // Complex number with float32 real and imaginary parts
complex128 // Complex number with float64 real and imaginary parts
```

### Composite Types
```go
// Array
var arr [5]int
arr := [3]string{"a", "b", "c"}

// Slice
var slice []int
slice := make([]int, 5)
slice := []int{1, 2, 3}
slice = append(slice, 4)

// Map
var m map[string]int
m := make(map[string]int)
m := map[string]int{"a": 1, "b": 2}

// Struct
type Person struct {
    Name string
    Age  int
}
```

## Control Structures

### If Statement
```go
if condition {
    // code
} else if condition {
    // code
} else {
    // code
}

// With initialization
if value := someFunction(); value > 0 {
    // code
}
```

### Switch Statement
```go
switch value {
case 1:
    // code
case 2, 3:
    // code
default:
    // code
}

// Type switch
switch v := value.(type) {
case int:
    // code
case string:
    // code
}
```

### For Loop
```go
// Basic for
for i := 0; i < 10; i++ {
    // code
}

// While-like
for condition {
    // code
}

// Infinite
for {
    // code
    if condition {
        break
    }
}

// Range
for index, value := range slice {
    // code
}
```

## Functions

### Function Declaration
```go
func name(parameter type) returnType {
    // code
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func calculate(x, y int) (sum, product int) {
    sum = x + y
    product = x * y
    return
}

// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

### Methods
```go
func (p Person) introduce() string {
    return fmt.Sprintf("Hi, I'm %s", p.Name)
}

// Pointer receiver
func (p *Person) haveBirthday() {
    p.Age++
}
```

## Structs and Interfaces

### Struct Definition
```go
type Person struct {
    Name string
    Age  int
    Address struct {
        Street string
        City   string
    }
}

// JSON tags
type User struct {
    Name     string    `json:"name"`
    Age      int       `json:"age"`
    Password string    `json:"-"`  // Excluded from JSON
}
```

### Interface Definition
```go
type Speaker interface {
    Speak() string
    Listen() error
}

// Interface composition
type ReaderWriter interface {
    io.Reader
    io.Writer
}
```

## Concurrency

### Goroutines
```go
go function()

// Anonymous function
go func() {
    // concurrent code
}()
```

### Channels
```go
// Create channel
ch := make(chan int)
ch := make(chan int, 10) // Buffered channel

// Send and receive
ch <- value    // Send
value := <-ch  // Receive

// Select
select {
case msg := <-ch1:
    // handle message
case msg := <-ch2:
    // handle message
case <-time.After(time.Second):
    // handle timeout
default:
    // handle default case
}

// Range over channel
for msg := range ch {
    // handle message
}
```

### Sync Package
```go
// Mutex
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()

// WaitGroup
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()
```

## Error Handling

### Error Checking
```go
if err != nil {
    // handle error
}

// Multiple return values
value, err := someFunction()
if err != nil {
    // handle error
}
```

### Custom Errors
```go
type CustomError struct {
    Code    int
    Message string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Error wrapping
if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}
```

## File I/O and JSON

### File Operations
```go
// Reading file
data, err := ioutil.ReadFile("file.txt")

// Writing file
err := ioutil.WriteFile("file.txt", data, 0644)

// Buffered I/O
file, _ := os.Create("file.txt")
writer := bufio.NewWriter(file)
writer.WriteString("Hello")
writer.Flush()
```

### JSON Handling
```go
// Marshaling
jsonData, err := json.Marshal(user)

// Unmarshaling
var user User
err := json.Unmarshal(jsonData, &user)

// Custom marshaling
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
```

## Web Development

### HTTP Server
```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)
```

### Middleware
```go
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %s", r.Method, r.URL, time.Since(start))
    }
}

// Middleware chain
func chain(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
    for i := len(middleware) - 1; i >= 0; i-- {
        h = middleware[i](h)
    }
    return h
}
```

## Database Operations

### Database Connection
```go
db, err := sql.Open("postgres", connStr)
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Connection pool
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(5 * time.Minute)
```

### Queries
```go
// Query
rows, err := db.Query("SELECT * FROM users")
defer rows.Close()

// Prepared statement
stmt, err := db.Prepare("SELECT * FROM users WHERE id = $1")
defer stmt.Close()

// Transaction
tx, err := db.Begin()
tx.Exec("INSERT INTO users (name) VALUES ($1)", "John")
tx.Commit()
```

## Security and Authentication

### JWT Authentication
```go
// Token generation
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, err := token.SignedString(jwtKey)

// Token validation
token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
    return jwtKey, nil
})
```

### Password Hashing
```go
// Hash password
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

// Check password
err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
```

### HTTPS Configuration
```go
config := &tls.Config{
    MinVersion:               tls.VersionTLS12,
    CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
    PreferServerCipherSuites: true,
}
```

## Deployment and DevOps

### Environment Configuration
```go
type Config struct {
    Port     string `json:"port"`
    Database struct {
        Host     string `json:"host"`
        Port     string `json:"port"`
        User     string `json:"user"`
        Password string `json:"password"`
    } `json:"database"`
}
```

### Graceful Shutdown
```go
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

go func() {
    <-sigChan
    log.Println("Shutting down server...")
    server.Close()
}()
```

## Generics

### Generic Functions
```go
func Print[T any](value T) {
    fmt.Println(value)
}

func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

### Generic Types
```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
    var zero T
    if len(s.items) == 0 {
        return zero, fmt.Errorf("stack is empty")
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, nil
}
```

## Best Practices

1. **Error Handling**
   - Always check errors
   - Use meaningful error messages
   - Consider error wrapping
   - Use custom error types

2. **Concurrency**
   - Use channels for communication
   - Avoid shared memory
   - Use sync package when needed
   - Implement proper synchronization

3. **Code Organization**
   - Keep packages focused
   - Use meaningful names
   - Follow Go idioms
   - Document exported functions

4. **Security**
   - Never store sensitive data in code
   - Use HTTPS
   - Implement proper authentication
   - Sanitize user input
   - Use secure defaults

5. **Performance**
   - Use appropriate data structures
   - Profile your code
   - Use connection pooling
   - Implement caching where appropriate
   - Handle timeouts properly

## Common Standard Library Packages

```go
import (
    "fmt"      // Formatting and printing
    "io"       // I/O primitives
    "net/http" // HTTP client and server
    "os"       // Operating system functions
    "strings"  // String manipulation
    "time"     // Time and date
    "sync"     // Synchronization primitives
    "context"  // Context for deadlines and cancellation
    "encoding/json" // JSON encoding/decoding
    "database/sql"  // Database operations
    "crypto/tls"    // TLS/SSL
)
```

## Common Go Commands

```bash
go run main.go           # Run a Go program
go build main.go        # Build a Go program
go test ./...           # Run tests
go mod init module      # Initialize a new module
go mod tidy            # Add missing and remove unused modules
go get package         # Add a dependency
go fmt ./...           # Format code
go vet ./...           # Report suspicious constructs
go doc package         # Show documentation
go generate           # Run code generators
```

## Resources

1. [Official Go Documentation](https://golang.org/doc/)
2. [Go by Example](https://gobyexample.com/)
3. [Effective Go](https://golang.org/doc/effective_go)
4. [Go Tour](https://tour.golang.org/)
5. [Go Blog](https://blog.golang.org/)
6. [Go Playground](https://play.golang.org/) 