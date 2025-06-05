# Ruby Methods and Blocks Examples

# Basic method definition
def greet(name)
  "Hello, #{name}!"
end

# Method with default parameter
def greet_with_time(name, time = Time.now)
  "Good #{time.hour < 12 ? 'morning' : 'afternoon'}, #{name}!"
end

# Method with multiple parameters
def calculate_total(price, tax_rate = 0.1, discount = 0)
  subtotal = price * (1 - discount)
  total = subtotal * (1 + tax_rate)
  total.round(2)
end

# Block with yield
def with_logging
  puts "Starting..."
  yield
  puts "Finished!"
end

# Using the logging block
with_logging do
  puts "Doing something..."
  sleep(1)
  puts "Done!"
end

# Lambda examples
double = ->(x) { x * 2 }
square = ->(x) { x ** 2 }

# Using lambdas
numbers = [1, 2, 3, 4, 5]
doubled = numbers.map(&double)
squared = numbers.map(&square)

# Method with block parameter
def process_items(items)
  items.each do |item|
    yield(item) if block_given?
  end
end

# Using the process_items method
fruits = ['apple', 'banana', 'orange']
process_items(fruits) { |fruit| puts "Processing #{fruit}" } 