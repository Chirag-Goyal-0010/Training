# Control Structures in Ruby

## 1. If Statements
```ruby
# Basic if
if condition
  # code
end

# If-else
if condition
  # code
else
  # code
end

# If-elsif-else
if condition1
  # code
elsif condition2
  # code
else
  # code
end

# Modifier form
puts "Hello" if condition
```

## 2. Unless Statement
```ruby
# Basic unless
unless condition
  # code
end

# Unless-else
unless condition
  # code
else
  # code
end

# Modifier form
puts "Hello" unless condition
```

## 3. Case Statement
```ruby
# Basic case
case expression
when value1
  # code
when value2
  # code
else
  # code
end

# Case with ranges
case number
when 1..5
  puts "Between 1 and 5"
when 6..10
  puts "Between 6 and 10"
else
  puts "Other"
end

# Case with multiple values
case value
when "a", "e", "i", "o", "u"
  puts "Vowel"
else
  puts "Consonant"
end
```

## 4. Loops

### While Loop
```ruby
# Basic while
while condition
  # code
end

# Modifier form
begin
  # code
end while condition
```

### Until Loop
```ruby
# Basic until
until condition
  # code
end

# Modifier form
begin
  # code
end until condition
```

### For Loop
```ruby
# Basic for
for i in 1..5
  puts i
end

# For with array
for item in array
  puts item
end
```

### Each Loop
```ruby
# Each with range
(1..5).each do |i|
  puts i
end

# Each with array
array.each do |item|
  puts item
end

# Each with index
array.each_with_index do |item, index|
  puts "#{index}: #{item}"
end
```

### Times Loop
```ruby
# Times loop
5.times do |i|
  puts i
end
```

## 5. Loop Control

### Break
```ruby
# Break from loop
while true
  break if condition
  # code
end
```

### Next
```ruby
# Skip to next iteration
for i in 1..5
  next if i == 3
  puts i
end
```

### Redo
```ruby
# Restart current iteration
for i in 1..5
  puts i
  redo if i == 3
end
```

### Retry
```ruby
# Retry from beginning
begin
  # code
rescue
  retry
end
```

## 6. Return Statement
```ruby
def method
  return value if condition
  # code
end
```

## 7. Throw/Catch
```ruby
# Throw and catch
catch :done do
  loop do
    throw :done if condition
    # code
  end
end
```

## 8. Begin/End Blocks
```ruby
# Basic begin/end
begin
  # code
end

# With rescue
begin
  # code
rescue
  # handle error
end

# With ensure
begin
  # code
ensure
  # always executed
end
```

## 9. Conditional Expressions
```ruby
# Ternary operator
result = condition ? true_value : false_value

# Or operator
value = other_value || default_value

# And operator
value = other_value && other_value.method
```

## 10. Guard Clauses
```ruby
# Early return
def method
  return if condition
  # code
end

# Guard with value
def method
  return value if condition
  # code
end
``` 