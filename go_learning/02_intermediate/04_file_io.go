package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 1. Basic File Operations
func basicFileOperations() {
	// Writing to a file
	data := []byte("Hello, Go!")
	err := ioutil.WriteFile("example.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	// Reading from a file
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(content))
}

// 2. Buffered I/O
func bufferedIO() {
	// Writing with buffer
	file, err := os.Create("buffered.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("Line 1\n")
	writer.WriteString("Line 2\n")
	writer.Flush()

	// Reading with buffer
	file, err = os.Open("buffered.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("Read line:", scanner.Text())
	}
}

// 3. File Operations
func fileOperations() {
	// Create directory
	err := os.Mkdir("testdir", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	}

	// Create nested directories
	err = os.MkdirAll("testdir/nested/deep", 0755)
	if err != nil {
		fmt.Println("Error creating nested directories:", err)
	}

	// List directory contents
	files, err := ioutil.ReadDir("testdir")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("Directory contents:")
	for _, file := range files {
		fmt.Printf("Name: %s, IsDir: %v\n", file.Name(), file.IsDir())
	}

	// Walk directory
	err = filepath.Walk("testdir", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("Walking:", path)
		return nil
	})
	if err != nil {
		fmt.Println("Error walking directory:", err)
	}
}

// 4. Environment Variables
func environmentVariables() {
	// Set environment variable
	os.Setenv("MY_VAR", "my_value")

	// Get environment variable
	value := os.Getenv("MY_VAR")
	fmt.Println("Environment variable:", value)

	// Get all environment variables
	fmt.Println("\nAll environment variables:")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}

// 5. JSON File Operations
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func jsonFileOperations() {
	// Create a person
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write JSON to file
	err = ioutil.WriteFile("person.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	// Read JSON from file
	jsonData, err = ioutil.ReadFile("person.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal JSON
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("Unmarshaled person:", newPerson)
}

// 6. File Permissions
func filePermissions() {
	// Create file with specific permissions
	file, err := os.OpenFile("permissions.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Get file info
	info, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	fmt.Printf("File permissions: %v\n", info.Mode())
}

// 7. Temporary Files
func temporaryFiles() {
	// Create temporary file
	tempFile, err := ioutil.TempFile("", "prefix-*.txt")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Clean up

	// Write to temporary file
	tempFile.WriteString("Temporary content")
	tempFile.Close()

	// Create temporary directory
	tempDir, err := ioutil.TempDir("", "tempdir-*")
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
		return
	}
	defer os.RemoveAll(tempDir) // Clean up

	fmt.Println("Temporary file:", tempFile.Name())
	fmt.Println("Temporary directory:", tempDir)
}

func main() {
	fmt.Println("=== Basic File Operations ===")
	basicFileOperations()

	fmt.Println("\n=== Buffered I/O ===")
	bufferedIO()

	fmt.Println("\n=== File Operations ===")
	fileOperations()

	fmt.Println("\n=== Environment Variables ===")
	environmentVariables()

	fmt.Println("\n=== JSON File Operations ===")
	jsonFileOperations()

	fmt.Println("\n=== File Permissions ===")
	filePermissions()

	fmt.Println("\n=== Temporary Files ===")
	temporaryFiles()
} 