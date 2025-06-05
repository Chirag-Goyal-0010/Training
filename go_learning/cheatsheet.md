# Go Programming Language Cheatsheet

## Table of Contents
1. [Basic Concepts](#basic-concepts)
2. [Intermediate Concepts](#intermediate-concepts)
3. [Advanced Concepts](#advanced-concepts)

## Basic Concepts

### Variables and Data Types
```go
// Variable declaration
var name string = "John"
age := 30  // Short declaration

// Basic types
var (
    intNum    int     = 42
    floatNum  float64 = 3.14
    boolVal   bool    = true
    stringVal string  = "Hello"
)

// Type conversion
floatNum = float64(intNum)
```

### Control Structures
```go
// If statement
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}

// Switch statement
switch day {
case "Monday":
    fmt.Println("Start of week")
default:
    fmt.Println("Other day")
}

// For loops
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Range loop
for index, value := range slice {
    fmt.Println(index, value)
}
```

### Collections
```go
// Arrays (fixed length)
var arr [5]int
arr := [3]string{"a", "b", "c"}

// Slices (dynamic length)
slice := []int{1, 2, 3}
slice = append(slice, 4)

// Maps
dict := map[string]int{
    "one": 1,
    "two": 2,
}
```

### Functions
```go
// Basic function
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Variadic functions
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

## Intermediate Concepts

### Structs and Interfaces
```go
// Struct definition
type Person struct {
    Name string
    Age  int
}

// Interface definition
type Speaker interface {
    Speak() string
}

// Method implementation
func (p Person) Speak() string {
    return fmt.Sprintf("Hi, I'm %s", p.Name)
}
```

### Error Handling
```go
// Error checking
if err != nil {
    log.Fatal(err)
}

// Custom errors
type ValidationError struct {
    Field string
    Msg   string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}
```

### File I/O
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
// Struct tags
type User struct {
    Name     string    `json:"name"`
    Age      int       `json:"age"`
    Password string    `json:"-"`  // Excluded from JSON
}

// Marshaling
jsonData, err := json.Marshal(user)

// Unmarshaling
var user User
err := json.Unmarshal(jsonData, &user)
```

## Advanced Concepts

### Concurrency
```go
// Goroutines
go func() {
    // Concurrent code
}()

// Channels
ch := make(chan int)
go func() {
    ch <- 42  // Send
}()
value := <-ch  // Receive

// Select statement
select {
case msg := <-ch1:
    fmt.Println(msg)
case <-time.After(time.Second):
    fmt.Println("timeout")
}
```

### Testing
```go
// Basic test
func TestAdd(t *testing.T) {
    result := add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}

// Table-driven tests
func TestDivide(t *testing.T) {
    tests := []struct {
        a, b float64
        want float64
    }{
        {10, 2, 5},
        {10, 0, 0},
    }
    // Test implementation
}
```

### REST API
```go
// Basic handler
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// Middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL)
        next.ServeHTTP(w, r)
    }
}
```

### Database Operations
```go
// Database connection
db, err := sql.Open("postgres", connStr)

// Query
rows, err := db.Query("SELECT * FROM users")

// Transaction
tx, err := db.Begin()
tx.Exec("INSERT INTO users (name) VALUES ($1)", "John")
tx.Commit()
```

### Authentication & Security
```go
// JWT token generation
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, err := token.SignedString(jwtKey)

// Password hashing
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

// HTTPS configuration
config := &tls.Config{
    MinVersion: tls.VersionTLS12,
}
```

### Deployment & DevOps
```go
// Environment configuration
type Config struct {
    Port     string `json:"port"`
    Database struct {
        Host string `json:"host"`
    } `json:"database"`
}

// Graceful shutdown
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
```

### Generics
```go
// Generic function
func Print[T any](value T) {
    fmt.Println(value)
}

// Generic type
type Stack[T any] struct {
    items []T
}

// Generic constraints
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

## Best Practices

1. **Error Handling**
   - Always check errors
   - Use custom error types
   - Wrap errors with context

2. **Concurrency**
   - Use channels for communication
   - Avoid shared memory
   - Use sync package for synchronization

3. **Testing**
   - Write tests for all packages
   - Use table-driven tests
   - Mock external dependencies

4. **Security**
   - Never store sensitive data in code
   - Use HTTPS
   - Implement proper authentication
   - Sanitize user input

5. **Performance**
   - Use appropriate data structures
   - Profile your code
   - Use connection pooling
   - Implement caching where appropriate

## Common Patterns

1. **Dependency Injection**
```go
type Service struct {
    repo Repository
}

func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}
```

2. **Middleware Chain**
```go
func chain(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
    for i := len(middleware) - 1; i >= 0; i-- {
        h = middleware[i](h)
    }
    return h
}
```

3. **Context Usage**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()
    // Use ctx for operations
}
```

4. **Worker Pool**
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- j * 2
    }
}
```

## Resources

1. Official Documentation
   - [Go Documentation](https://golang.org/doc/)
   - [Go Tour](https://tour.golang.org/)
   - [Go Blog](https://blog.golang.org/)

2. Books
   - "The Go Programming Language"
   - "Go in Action"
   - "Concurrency in Go"

3. Tools
   - `go fmt` - Format code
   - `go vet` - Report suspicious constructs
   - `go test` - Run tests
   - `go mod` - Dependency management 