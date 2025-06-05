# Query:

# bike.move will output: Moving...
# bike.ring_bell will output: Ring ring!

class Vehicle
  def move
    puts "Moving..."
  end
end

class Bike < Vehicle
  def ring_bell
    puts "Ring ring!"
  end
end

bike = Bike.new
bike.move
bike.ring_bell
