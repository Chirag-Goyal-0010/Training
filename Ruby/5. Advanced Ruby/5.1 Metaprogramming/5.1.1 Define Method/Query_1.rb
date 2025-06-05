# greet.hello will output: Hello!

class Greeting
  define_method(:hello) do
    puts "Hello!"
  end
end

greet = Greeting.new
greet.hello
