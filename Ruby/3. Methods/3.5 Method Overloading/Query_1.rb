# How do you define a method that can take a variable number of arguments and greet each name?

# Method 1
def greet (*arg)
  arg.each do |element|
    puts "Hello, #{element}"
  end
end

greet("Chirag", "Goyal", "Abcde")

# Method 2
def greet_2(*names)
  names.each { |name| puts "Hello, #{name}!" }
end

greet_2("Chirag", "Goyal", "Efghi")