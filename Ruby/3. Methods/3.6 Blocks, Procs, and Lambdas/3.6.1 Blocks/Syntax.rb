# How do you define a method that takes a block and executes it a specified number of times?

def repeat(times)
  times.times { yield }
end

repeat(3) { puts "Hello" }
