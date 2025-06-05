# person.introduce will output: Hello, my name is Alice and I am 30 years old.

class Person
  def initialize(name, age)
    @name = name
    @age = age
  end

  def introduce
    puts "Hello, my name is #{@name} and I am #{@age} years old."
  end
end

person = Person.new("Alice", 30)
person.introduce
