# The code will output:

# Drawing a circle
# Drawing a square

class Shape
  def draw
    puts "Drawing a shape"
  end
end

class Circle < Shape
  def draw
    puts "Drawing a circle"
  end
end

class Square < Shape
  def draw
    puts "Drawing a square"
  end
end

shapes = [Circle.new, Square.new]
shapes.each { |shape| shape.draw }
