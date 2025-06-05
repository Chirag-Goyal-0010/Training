# Operators in Ruby

## 1. Arithmetic Operators
```ruby
# Addition
puts 5 + 3    # 8

# Subtraction
puts 5 - 3    # 2

# Multiplication
puts 5 * 3    # 15

# Division
puts 10 / 2   # 5

# Modulus (Remainder)
puts 10 % 3   # 1

# Exponentiation
puts 2 ** 3   # 8
```

## 2. Comparison Operators
```ruby
# Equal to
puts 5 == 5   # true

# Not equal to
puts 5 != 3   # true

# Greater than
puts 5 > 3    # true

# Less than
puts 3 < 5    # true

# Greater than or equal to
puts 5 >= 5   # true

# Less than or equal to
puts 5 <= 5   # true

# Combined comparison
puts 5 <=> 5  # 0 (equal)
puts 5 <=> 3  # 1 (greater)
puts 3 <=> 5  # -1 (less)

# Case equality
puts 5 === 5.0  # true
```

## 3. Logical Operators
```ruby
# AND
puts true && true    # true
puts true && false   # false

# OR
puts true || false   # true
puts false || false  # false

# NOT
puts !true          # false
puts !false         # true
```

## 4. Assignment Operators
```ruby
# Simple assignment
x = 5

# Add and assign
x += 3  # same as x = x + 3

# Subtract and assign
x -= 2  # same as x = x - 2

# Multiply and assign
x *= 2  # same as x = x * 2

# Divide and assign
x /= 2  # same as x = x / 2

# Modulus and assign
x %= 2  # same as x = x % 2

# Exponent and assign
x **= 2 # same as x = x ** 2
```

## 5. Bitwise Operators
```ruby
# Bitwise AND
puts 5 & 3    # 1

# Bitwise OR
puts 5 | 3    # 7

# Bitwise XOR
puts 5 ^ 3    # 6

# Left shift
puts 5 << 1   # 10

# Right shift
puts 5 >> 1   # 2

# Ones complement
puts ~5       # -6
```

## 6. Ternary Operator
```ruby
# condition ? true_value : false_value
age = 20
status = age >= 18 ? "adult" : "minor"
puts status  # "adult"
```

## 7. Range Operators
```ruby
# Inclusive range
(1..5).to_a   # [1, 2, 3, 4, 5]

# Exclusive range
(1...5).to_a  # [1, 2, 3, 4]
```

## 8. defined? Operator
```ruby
# Check if variable is defined
x = 5
puts defined? x      # "local-variable"
puts defined? y      # nil

# Check if method is defined
puts defined? puts   # "method"
```

## 9. Dot and Double Colon Operators
```ruby
# Dot operator for method calls
"hello".length

# Double colon for constants
Math::PI

# Double colon for module/class methods
File::exist?("file.txt")
```

## Operator Precedence
1. `[ ] [ ]=` - Element reference, element set
2. `**` - Exponentiation
3. `! ~ +` - Boolean NOT, bitwise complement, unary plus
4. `* / %` - Multiplication, division, modulo
5. `+ -` - Addition, subtraction
6. `<< >>` - Bitwise shift
7. `&` - Bitwise AND
8. `| ^` - Bitwise OR, XOR
9. `> >= < <=` - Comparison
10. `<=> == === != =~ !~` - Equality, pattern matching
11. `&&` - Logical AND
12. `||` - Logical OR
13. `.. ...` - Range creation
14. `? :` - Ternary operator
15. `= += -= etc.` - Assignment
16. `defined?` - Check definition
17. `not` - Logical NOT
18. `or and` - Logical OR, AND 