# Car.drive will output: I can drive!

module Drivable
  def drive
    puts "I can drive!"
  end
end

class Car
  extend Drivable
end

Car.drive
