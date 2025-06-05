# dog.speak will output: Woof!

class Animal
  def speak
    puts "Generic animal sound"
  end
end

class Dog < Animal
  def speak
    puts "Woof!"
  end
end

dog = Dog.new
dog.speak
