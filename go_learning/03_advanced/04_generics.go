package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// 1. Basic Generic Function
func Print[T any](value T) {
	fmt.Println(value)
}

// 2. Generic Function with Constraints
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// 3. Generic Type
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

// 4. Generic Interface
type Container[T any] interface {
	Add(T)
	Get() T
	IsEmpty() bool
}

// 5. Generic Struct with Methods
type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if len(q.items) == 0 {
		return zero, fmt.Errorf("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// 6. Generic Map Function
func Map[T, U any](items []T, f func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = f(item)
	}
	return result
}

// 7. Generic Filter Function
func Filter[T any](items []T, f func(T) bool) []T {
	var result []T
	for _, item := range items {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

// 8. Generic Reduce Function
func Reduce[T, U any](items []T, initial U, f func(U, T) U) U {
	result := initial
	for _, item := range items {
		result = f(result, item)
	}
	return result
}

// 9. Custom Type Constraints
type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[T Number](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// 10. Generic Type with Multiple Type Parameters
type Pair[T, U any] struct {
	First  T
	Second U
}

func (p Pair[T, U]) Swap() Pair[U, T] {
	return Pair[U, T]{First: p.Second, Second: p.First}
}

func main() {
	// 1. Basic Generic Function
	fmt.Println("=== Basic Generic Function ===")
	Print("Hello")
	Print(42)
	Print(3.14)

	// 2. Generic Function with Constraints
	fmt.Println("\n=== Generic Function with Constraints ===")
	fmt.Println("Min of 5 and 3:", Min(5, 3))
	fmt.Println("Min of 3.14 and 2.71:", Min(3.14, 2.71))
	fmt.Println("Min of 'a' and 'b':", Min('a', 'b'))

	// 3. Generic Stack
	fmt.Println("\n=== Generic Stack ===")
	stack := &Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	item, _ := stack.Pop()
	fmt.Println("Popped:", item)

	// 4. Generic Queue
	fmt.Println("\n=== Generic Queue ===")
	queue := &Queue[string]{}
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")
	item2, _ := queue.Dequeue()
	fmt.Println("Dequeued:", item2)

	// 5. Generic Map
	fmt.Println("\n=== Generic Map ===")
	numbers := []int{1, 2, 3, 4, 5}
	doubled := Map(numbers, func(n int) int {
		return n * 2
	})
	fmt.Println("Doubled:", doubled)

	// 6. Generic Filter
	fmt.Println("\n=== Generic Filter ===")
	evenNumbers := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", evenNumbers)

	// 7. Generic Reduce
	fmt.Println("\n=== Generic Reduce ===")
	sum := Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Println("Sum:", sum)

	// 8. Custom Type Constraints
	fmt.Println("\n=== Custom Type Constraints ===")
	ints := []int{1, 2, 3, 4, 5}
	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("Sum of ints:", Sum(ints))
	fmt.Println("Sum of floats:", Sum(floats))

	// 9. Generic Pair
	fmt.Println("\n=== Generic Pair ===")
	pair := Pair[string, int]{"age", 30}
	fmt.Println("Original pair:", pair)
	swapped := pair.Swap()
	fmt.Println("Swapped pair:", swapped)

	// 10. Generic Type with Multiple Parameters
	fmt.Println("\n=== Generic Type with Multiple Parameters ===")
	config := Pair[string, map[string]string]{
		First: "database",
		Second: map[string]string{
			"host": "localhost",
			"port": "5432",
		},
	}
	fmt.Println("Config:", config)
} 