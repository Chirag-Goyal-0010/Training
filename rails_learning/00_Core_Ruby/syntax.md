# Core Ruby Syntax Examples

## Variables and Data Types

### Variables
```ruby
# Local variables
name = "John"
age = 25

# Instance variables
@user_name = "John"

# Class variables
@@total_users = 0

# Global variables
$debug_mode = true

# Constants
MAX_USERS = 100
```

### Numbers
```ruby
# Integers
age = 25
big_number = 1_000_000

# Floating point
price = 19.99
pi = 3.14159

# Complex numbers
complex = Complex(3, 4)

# Rational numbers
rational = Rational(1, 3)  # 1/3
```

### Strings
```ruby
# String creation
name = "John"
greeting = 'Hello'

# String interpolation
message = "Hello, #{name}!"

# String methods
"hello".upcase        # "HELLO"
"HELLO".downcase     # "hello"
" hello ".strip      # "hello"
"hello".reverse      # "olleh"
```

### Arrays
```ruby
# Array creation
numbers = [1, 2, 3, 4, 5]
mixed = [1, "hello", 3.14, true]

# Array methods
numbers.push(6)      # [1, 2, 3, 4, 5, 6]
numbers.pop          # [1, 2, 3, 4, 5]
numbers.shift        # [2, 3, 4, 5]
numbers.unshift(0)   # [0, 2, 3, 4, 5]
```

### Hashes
```ruby
# Hash creation
person = { name: "John", age: 25 }
person_old = { :name => "John", :age => 25 }

# Hash methods
person[:name]        # "John"
person[:age] = 26    # Update value
person[:city] = "NY" # Add new key-value
```

## Methods and Blocks

### Method Definition
```ruby
# Basic method
def greet(name)
  "Hello, #{name}!"
end

# Method with default parameter
def greet(name = "Guest")
  "Hello, #{name}!"
end

# Method with variable arguments
def sum(*numbers)
  numbers.sum
end

# Method with keyword arguments
def create_user(name:, email:, age: 18)
  { name: name, email: email, age: age }
end
```

### Blocks
```ruby
# Array iteration with blocks
[1, 2, 3].each { |num| puts num }

# Multi-line block
[1, 2, 3].each do |num|
  puts "Number: #{num}"
  puts "Square: #{num ** 2}"
end

# Block with yield
def with_logging
  puts "Starting..."
  yield
  puts "Finished!"
end

with_logging { puts "Doing something..." }
```

## Classes and Objects

### Class Definition
```ruby
class Person
  # Class variable
  @@total_people = 0

  # Constructor
  def initialize(name, age)
    @name = name
    @age = age
    @@total_people += 1
  end

  # Instance methods
  def greet
    "Hello, I'm #{@name}!"
  end

  # Class method
  def self.total_people
    @@total_people
  end
end

# Object creation and usage
person = Person.new("John", 25)
puts person.greet
puts Person.total_people
```

## Modules and Mixins

### Module Definition
```ruby
module Greetable
  def greet
    "Hello!"
  end
end

module Workable
  def work
    "Working..."
  end
end

class Employee
  include Greetable
  include Workable
end

employee = Employee.new
puts employee.greet  # "Hello!"
puts employee.work   # "Working..."
```

## Exception Handling

### Basic Exception Handling
```ruby
begin
  # Risky operation
  result = 10 / 0
rescue ZeroDivisionError => e
  puts "Error: #{e.message}"
rescue StandardError => e
  puts "Unexpected error: #{e.message}"
ensure
  puts "This always executes"
end
```

### Custom Exception
```ruby
class ValidationError < StandardError
  def initialize(message = "Validation failed")
    super(message)
  end
end

def validate_age(age)
  raise ValidationError, "Age must be positive" if age < 0
  raise ValidationError, "Age must be less than 150" if age > 150
end
```

## Ruby Gems and Bundler

### Gemfile Example
```ruby
# Gemfile
source 'https://rubygems.org'

gem 'rails', '~> 7.0.0'
gem 'sqlite3'
gem 'puma'
gem 'sass-rails'
gem 'uglifier'
gem 'coffee-rails'
gem 'jquery-rails'
gem 'turbolinks'
gem 'jbuilder'

group :development, :test do
  gem 'byebug'
  gem 'rspec-rails'
end

group :development do
  gem 'web-console'
  gem 'spring'
end
```

### Bundler Commands
```bash
# Install gems
bundle install

# Update gems
bundle update

# Run command with bundled gems
bundle exec rails server

# Check for outdated gems
bundle outdated
```

## Best Practices Examples

### Method Organization
```ruby
class User
  # Public methods first
  def full_name
    "#{first_name} #{last_name}"
  end

  # Protected methods
  protected
  def validate_email
    # email validation logic
  end

  # Private methods last
  private
  def generate_token
    # token generation logic
  end
end
```

### Error Handling
```ruby
def process_file(filename)
  raise ArgumentError, "Filename cannot be empty" if filename.empty?
  
  begin
    File.open(filename) do |file|
      # Process file
    end
  rescue Errno::ENOENT
    puts "File not found: #{filename}"
  rescue StandardError => e
    puts "Error processing file: #{e.message}"
  end
end
```

### Documentation
```ruby
# Class documentation
class Calculator
  # Adds two numbers
  # @param a [Numeric] First number
  # @param b [Numeric] Second number
  # @return [Numeric] Sum of a and b
  def add(a, b)
    a + b
  end
end
``` 