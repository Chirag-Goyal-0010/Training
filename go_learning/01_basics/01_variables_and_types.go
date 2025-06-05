package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. Variable Declaration
	// Using var keyword
	var name string = "John"
	var age int = 25
	var isStudent bool = true

	// Short declaration (:=)
	score := 95.5
	grade := 'A'

	// Multiple variable declaration
	var (
		firstName = "John"
		lastName  = "Doe"
		height    = 180
	)

	// 2. Basic Data Types
	// Integer types
	var intNum int = 42
	var int8Num int8 = 127
	var int16Num int16 = 32767
	var int32Num int32 = 2147483647
	var int64Num int64 = 9223372036854775807

	// Floating point types
	var float32Num float32 = 3.14
	var float64Num float64 = 3.14159265359

	// String type
	var message string = "Hello, Go!"

	// Boolean type
	var isTrue bool = true

	// 3. Type Inference
	inferredInt := 42
	inferredFloat := 3.14
	inferredString := "Hello"

	// 4. Constants
	const PI = 3.14159
	const MAX_VALUE = 100

	// 5. Type Conversion
	var intValue int = 42
	var floatValue float64 = float64(intValue)
	var stringValue string = string(intValue)

	// 6. Zero Values
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	// Print all variables to see their values
	fmt.Println("=== Basic Variables ===")
	fmt.Printf("name: %v, type: %T\n", name, name)
	fmt.Printf("age: %v, type: %T\n", age, age)
	fmt.Printf("isStudent: %v, type: %T\n", isStudent, isStudent)
	fmt.Printf("score: %v, type: %T\n", score, score)
	fmt.Printf("grade: %v, type: %T\n", grade, grade)

	fmt.Println("\n=== Multiple Variables ===")
	fmt.Printf("firstName: %v, lastName: %v, height: %v\n", firstName, lastName, height)

	fmt.Println("\n=== Integer Types ===")
	fmt.Printf("int: %v\n", intNum)
	fmt.Printf("int8: %v\n", int8Num)
	fmt.Printf("int16: %v\n", int16Num)
	fmt.Printf("int32: %v\n", int32Num)
	fmt.Printf("int64: %v\n", int64Num)

	fmt.Println("\n=== Floating Point Types ===")
	fmt.Printf("float32: %v\n", float32Num)
	fmt.Printf("float64: %v\n", float64Num)

	fmt.Println("\n=== Other Types ===")
	fmt.Printf("string: %v\n", message)
	fmt.Printf("bool: %v\n", isTrue)

	fmt.Println("\n=== Type Inference ===")
	fmt.Printf("inferredInt: %v, type: %T\n", inferredInt, inferredInt)
	fmt.Printf("inferredFloat: %v, type: %T\n", inferredFloat, inferredFloat)
	fmt.Printf("inferredString: %v, type: %T\n", inferredString, inferredString)

	fmt.Println("\n=== Constants ===")
	fmt.Printf("PI: %v\n", PI)
	fmt.Printf("MAX_VALUE: %v\n", MAX_VALUE)

	fmt.Println("\n=== Type Conversion ===")
	fmt.Printf("int to float64: %v\n", floatValue)
	fmt.Printf("int to string: %v\n", stringValue)

	fmt.Println("\n=== Zero Values ===")
	fmt.Printf("zeroInt: %v\n", zeroInt)
	fmt.Printf("zeroFloat: %v\n", zeroFloat)
	fmt.Printf("zeroString: %v\n", zeroString)
	fmt.Printf("zeroBool: %v\n", zeroBool)

	// 7. Type Reflection
	fmt.Println("\n=== Type Reflection ===")
	fmt.Printf("Type of intNum: %v\n", reflect.TypeOf(intNum))
	fmt.Printf("Type of float32Num: %v\n", reflect.TypeOf(float32Num))
	fmt.Printf("Type of message: %v\n", reflect.TypeOf(message))
} 