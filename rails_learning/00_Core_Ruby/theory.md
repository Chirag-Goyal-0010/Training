# Core Ruby Theory

## Variables, Data Types, and Operators

### Variables
- Local variables (start with lowercase or underscore)
- Instance variables (start with @)
- Class variables (start with @@)
- Global variables (start with $)
- Constants (start with uppercase)

### Data Types
1. **Numbers**
   - Integers (Fixnum, Bignum)
   - Floating-point numbers
   - Complex numbers
   - Rational numbers

2. **Strings**
   - Single quotes ('') vs double quotes ("")
   - String interpolation (#{variable})
   - String methods (upcase, downcase, strip, etc.)

3. **Symbols**
   - Immutable strings
   - Used as identifiers
   - Memory efficient

4. **Arrays**
   - Ordered collections
   - Zero-based indexing
   - Can hold mixed data types
   - Common methods: push, pop, shift, unshift

5. **Hashes**
   - Key-value pairs
   - Keys can be symbols or strings
   - New syntax: { key: value }
   - Old syntax: { :key => value }

6. **Booleans**
   - true and false
   - nil (represents absence of value)

### Operators
- Arithmetic (+, -, *, /, %, **)
- Comparison (==, !=, >, <, >=, <=)
- Logical (&&, ||, !)
- Assignment (=, +=, -=, *=, /=)
- Ternary (condition ? true_value : false_value)

## Methods and Blocks

### Methods
```ruby
def method_name(parameter1, parameter2)
  # method body
end
```

- Method naming conventions
- Default parameters
- Variable number of arguments (*args)
- Keyword arguments
- Return values (explicit and implicit)

### Blocks
```ruby
# Single-line block
[1, 2, 3].each { |num| puts num }

# Multi-line block
[1, 2, 3].each do |num|
  puts num
end
```

- Block parameters
- yield keyword
- block_given?
- Proc objects
- Lambda functions

## Classes and Objects

### Classes
```ruby
class Person
  def initialize(name)
    @name = name
  end

  def greet
    "Hello, #{@name}!"
  end
end
```

- Class definition
- Instance methods
- Class methods
- Constructor (initialize)
- Instance variables
- Class variables

### Objects
- Object instantiation
- Object state
- Object behavior
- Object identity
- Object comparison

## Modules and Mixins

### Modules
```ruby
module Greetable
  def greet
    "Hello!"
  end
end
```

- Module definition
- Module methods
- Module constants
- Namespacing

### Mixins
```ruby
class Person
  include Greetable
end
```

- include vs extend
- Module inclusion order
- Method lookup path
- Module composition

## Inheritance and Polymorphism

### Inheritance
```ruby
class Employee < Person
  def work
    "Working..."
  end
end
```

- Single inheritance
- Superclass methods
- Method overriding
- super keyword

### Polymorphism
- Duck typing
- Method overriding
- Method overloading
- Interface-like behavior

## Exception Handling

### Basic Exception Handling
```ruby
begin
  # code that might raise an exception
rescue StandardError => e
  # handle the exception
ensure
  # code that always executes
end
```

- Exception hierarchy
- Custom exceptions
- Multiple rescue clauses
- Retry mechanism
- Raise vs fail

## Ruby Gems and Bundler

### Ruby Gems
- What are gems?
- Gem structure
- Gem installation
- Gem usage
- Gem documentation

### Bundler
- Gemfile
- bundle install
- bundle update
- bundle exec
- Gem versioning
- Gem groups

## Best Practices
1. Follow Ruby style guide
2. Use meaningful variable names
3. Keep methods small and focused
4. Use modules for shared functionality
5. Handle exceptions appropriately
6. Document your code
7. Write tests
8. Use version control
9. Follow DRY principle
10. Use appropriate data structures 