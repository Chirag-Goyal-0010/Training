# my_car.display will output: Car: Toyota Corolla


class Car
    def initialize(make, model)
      @make = make
      @model = model
    end
  
    def display
      puts "Car: #{@make} #{@model}"
    end
  end
  
  my_car = Car.new("Toyota", "Corolla")
  my_car.display
  