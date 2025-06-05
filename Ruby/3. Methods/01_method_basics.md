# Methods in Ruby

## Basic Method Definition
```ruby
# Simple method
def greet
  puts "Hello!"
end

# Method with parameters
def greet(name)
  puts "Hello, #{name}!"
end

# Method with default parameters
def greet(name = "World")
  puts "Hello, #{name}!"
end
```

## Method Parameters
```ruby
# Required parameters
def add(a, b)
  a + b
end

# Optional parameters
def greet(name, greeting = "Hello")
  puts "#{greeting}, #{name}!"
end

# Variable number of parameters
def sum(*numbers)
  numbers.sum
end

# Keyword arguments
def create_user(name:, age:, email:)
  { name: name, age: age, email: email }
end
```

## Return Values
```ruby
# Explicit return
def calculate(x, y)
  return x + y if x > 0
  return x - y
end

# Implicit return
def square(x)
  x * x
end

# Multiple return values
def get_coordinates
  [x, y]
end
```

## Method Visibility
```ruby
class Example
  # Public methods (default)
  def public_method
    # accessible from anywhere
  end

  # Protected methods
  protected
  def protected_method
    # accessible within the class and subclasses
  end

  # Private methods
  private
  def private_method
    # accessible only within the class
  end
end
```

## Method Aliases
```ruby
class String
  alias_method :old_length, :length
  
  def length
    old_length + 1
  end
end
```

## Method Chaining
```ruby
def process_data
  fetch_data
    .transform
    .validate
    .save
end
```

## Method Missing
```ruby
class DynamicMethods
  def method_missing(name, *args)
    if name.to_s.start_with?('find_by_')
      # Handle dynamic finder methods
    else
      super
    end
  end
end
```

## Method Arguments
```ruby
# Required arguments
def required_args(a, b, c)
  [a, b, c]
end

# Optional arguments
def optional_args(a, b = 2, c = 3)
  [a, b, c]
end

# Rest arguments
def rest_args(a, *rest)
  [a, rest]
end

# Keyword arguments
def keyword_args(a:, b: 2, c: 3)
  [a, b, c]
end

# Double splat for keyword arguments
def double_splat_args(a, **options)
  [a, options]
end
```

## Method Documentation
```ruby
# Method with documentation
# @param name [String] the name of the person
# @param age [Integer] the age of the person
# @return [Hash] a hash containing the person's details
def create_person(name, age)
  { name: name, age: age }
end
```

## Method Best Practices
1. Keep methods small and focused
2. Use descriptive names
3. Follow the single responsibility principle
4. Use appropriate visibility modifiers
5. Document complex methods
6. Handle errors appropriately
7. Use meaningful parameter names
8. Consider method chaining
9. Use default parameters when appropriate
10. Return meaningful values 