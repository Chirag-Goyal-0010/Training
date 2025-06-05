# How do you define a Proc that takes a name as an argument and prints "Hello, [name]!"?

my_proc = Proc.new { |name| puts "Hello, #{name}!" }
my_proc.call("Chirag")