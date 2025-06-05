# Variables and Data Types in Ruby

## Variables
Variables in Ruby are dynamically typed, meaning you don't need to declare their type. Ruby automatically determines the type based on the value assigned.

```ruby
# Variable declaration
name = "John"
age = 25
height = 1.75
is_student = true

# Multiple variables
first_name, last_name = "John", "Doe"
```

## Data Types

### 1. Numbers
```ruby
# Integers
num1 = 100
num2 = 1_00  # Same as 100
num3 = 10_0  # Same as 100

# Floats
pi = 3.14
e = 2.71828

# Different number systems
hex = 0xaa    # Hexadecimal
oct = 0o77    # Octal
bin = 0b1010  # Binary
```

### 2. Strings
```ruby
# String creation
str1 = "Hello"
str2 = 'World'
str3 = String.new("Ruby")

# String interpolation
name = "John"
puts "Hello, #{name}!"

# String methods
puts str1.length
puts str1.upcase
puts str1.downcase
```

### 3. Arrays
```ruby
# Array creation
arr1 = [1, 2, 3]
arr2 = ["a", "b", "c"]
arr3 = [1, "hello", 3.14, true]

# Array methods
arr1.push(4)
arr1.pop
arr1 << 5
```

### 4. Hashes
```ruby
# Hash creation
hash1 = { "name" => "John", "age" => 25 }
hash2 = { name: "John", age: 25 }  # Symbol keys
hash3 = Hash.new

# Hash methods
hash1["name"] = "Jane"
hash1.key?("name")
hash1.value?(25)
```

### 5. Symbols
```ruby
# Symbol creation
:name
:age
:email

# Symbol characteristics
puts :example.object_id == :example.object_id  # true
```

### 6. Boolean
```ruby
true
false
nil  # Ruby's version of null
```

## Type Conversion
```ruby
# String to Integer
"123".to_i

# Integer to String
123.to_s

# String to Float
"3.14".to_f

# Float to Integer
3.14.to_i
```

## Variable Scope
1. **Local Variables**
   - Start with lowercase letter or underscore
   - Only accessible within their scope

2. **Instance Variables**
   - Start with @
   - Accessible across methods in the same instance

3. **Class Variables**
   - Start with @@
   - Shared among all instances of a class

4. **Global Variables**
   - Start with $
   - Accessible from anywhere in the program

```ruby
# Example of different variable types
class Example
  @@class_var = 0  # Class variable
  
  def initialize
    @instance_var = 0  # Instance variable
  end
  
  def method
    local_var = 0  # Local variable
    $global_var = 0  # Global variable
  end
end
``` 