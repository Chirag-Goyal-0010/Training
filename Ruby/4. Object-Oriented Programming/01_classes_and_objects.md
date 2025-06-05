# Classes and Objects in Ruby

## Basic Class Definition
```ruby
class Person
  # Class variable
  @@count = 0

  # Constructor
  def initialize(name, age)
    @name = name
    @age = age
    @@count += 1
  end

  # Instance methods
  def introduce
    "Hi, I'm #{@name} and I'm #{@age} years old."
  end

  # Class method
  def self.count
    @@count
  end
end
```

## Object Creation and Usage
```ruby
# Creating objects
person1 = Person.new("John", 25)
person2 = Person.new("Jane", 30)

# Calling methods
puts person1.introduce
puts Person.count
```

## Attributes
```ruby
class Person
  # Attribute readers
  attr_reader :name, :age

  # Attribute writers
  attr_writer :name, :age

  # Attribute accessors (both reader and writer)
  attr_accessor :name, :age

  def initialize(name, age)
    @name = name
    @age = age
  end
end
```

## Inheritance
```ruby
class Animal
  def initialize(name)
    @name = name
  end

  def speak
    "Some sound"
  end
end

class Dog < Animal
  def speak
    "Woof!"
  end
end

class Cat < Animal
  def speak
    "Meow!"
  end
end
```

## Modules
```ruby
# Module definition
module Swimmable
  def swim
    "I'm swimming!"
  end
end

# Including module in class
class Fish
  include Swimmable
end

# Extending module
class Person
  extend Swimmable
end
```

## Access Modifiers
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

## Class Methods vs Instance Methods
```ruby
class Calculator
  # Instance method
  def add(a, b)
    a + b
  end

  # Class method
  def self.add(a, b)
    a + b
  end
end

# Usage
calc = Calculator.new
calc.add(2, 3)      # Instance method
Calculator.add(2, 3) # Class method
```

## Method Overriding
```ruby
class Parent
  def greet
    "Hello from parent"
  end
end

class Child < Parent
  def greet
    "Hello from child"
  end
end
```

## Method Overloading
```ruby
class Calculator
  def add(*args)
    if args.length == 2
      args[0] + args[1]
    elsif args.length == 3
      args[0] + args[1] + args[2]
    end
  end
end
```

## Object Initialization
```ruby
class Person
  def initialize(name, age)
    @name = name
    @age = age
  end

  # Alternative initialization
  def self.create(name, age)
    new(name, age)
  end
end
```

## Object Comparison
```ruby
class Person
  include Comparable

  attr_reader :age

  def initialize(name, age)
    @name = name
    @age = age
  end

  def <=>(other)
    age <=> other.age
  end
end
```

## Best Practices
1. Keep classes focused and single-purpose
2. Use meaningful class and method names
3. Follow the DRY (Don't Repeat Yourself) principle
4. Use appropriate access modifiers
5. Document public interfaces
6. Use modules for shared functionality
7. Follow Ruby naming conventions
8. Use attr_accessor, attr_reader, and attr_writer appropriately
9. Implement proper error handling
10. Use inheritance and composition appropriately 