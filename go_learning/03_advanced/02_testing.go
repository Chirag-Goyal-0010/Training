package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// 1. Basic Function to Test
func add(a, b int) int {
	return a + b
}

// 2. Function with Error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// 3. Function with Timeout
func processWithTimeout(duration time.Duration) error {
	time.Sleep(duration)
	return nil
}

// 4. Function with Dependencies
type Calculator interface {
	Add(a, b int) int
	Subtract(a, b int) int
}

type SimpleCalculator struct{}

func (c *SimpleCalculator) Add(a, b int) int {
	return a + b
}

func (c *SimpleCalculator) Subtract(a, b int) int {
	return a - b
}

func calculate(c Calculator, a, b int, operation string) int {
	switch operation {
	case "add":
		return c.Add(a, b)
	case "subtract":
		return c.Subtract(a, b)
	default:
		return 0
	}
}

// 5. Function with Side Effects
type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) GetCount() int {
	return c.count
}

// Tests

// 1. Basic Test
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"zero", 0, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// 2. Test with Error
func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a           float64
		b           float64
		expected    float64
		expectError bool
	}{
		{"valid division", 6, 2, 3, false},
		{"division by zero", 6, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := divide(tt.a, tt.b)
			if tt.expectError {
				if err == nil {
					t.Error("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("divide(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}

// 3. Test with Timeout
func TestProcessWithTimeout(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		timeout  time.Duration
		wantErr  bool
	}{
		{"short duration", 100 * time.Millisecond, 200 * time.Millisecond, false},
		{"long duration", 300 * time.Millisecond, 200 * time.Millisecond, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			done := make(chan error, 1)
			go func() {
				done <- processWithTimeout(tt.duration)
			}()

			select {
			case err := <-done:
				if (err != nil) != tt.wantErr {
					t.Errorf("processWithTimeout() error = %v, wantErr %v", err, tt.wantErr)
				}
			case <-time.After(tt.timeout):
				if !tt.wantErr {
					t.Error("processWithTimeout() timed out")
				}
			}
		})
	}
}

// 4. Test with Mock
type MockCalculator struct {
	AddFunc      func(a, b int) int
	SubtractFunc func(a, b int) int
}

func (m *MockCalculator) Add(a, b int) int {
	if m.AddFunc != nil {
		return m.AddFunc(a, b)
	}
	return 0
}

func (m *MockCalculator) Subtract(a, b int) int {
	if m.SubtractFunc != nil {
		return m.SubtractFunc(a, b)
	}
	return 0
}

func TestCalculate(t *testing.T) {
	mock := &MockCalculator{
		AddFunc: func(a, b int) int {
			return a + b + 1 // Mock implementation
		},
		SubtractFunc: func(a, b int) int {
			return a - b - 1 // Mock implementation
		},
	}

	tests := []struct {
		name       string
		calculator Calculator
		a          int
		b          int
		operation  string
		expected   int
	}{
		{"add with mock", mock, 2, 3, "add", 6},
		{"subtract with mock", mock, 5, 2, "subtract", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculate(tt.calculator, tt.a, tt.b, tt.operation)
			if result != tt.expected {
				t.Errorf("calculate() = %d; want %d", result, tt.expected)
			}
		})
	}
}

// 5. Test with Setup and Teardown
func TestCounter(t *testing.T) {
	counter := &Counter{}

	// Setup
	t.Run("increment", func(t *testing.T) {
		counter.Increment()
		if counter.GetCount() != 1 {
			t.Errorf("Counter.GetCount() = %d; want 1", counter.GetCount())
		}
	})

	// Teardown
	t.Cleanup(func() {
		counter.count = 0
	})
}

// 6. Benchmark Test
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}
}

// 7. Example Test
func ExampleAdd() {
	result := add(2, 3)
	fmt.Println(result)
	// Output: 5
}

// 8. Test Helper
func assertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestAddWithHelper(t *testing.T) {
	result := add(2, 3)
	assertEqual(t, result, 5)
} 