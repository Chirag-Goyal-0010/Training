# bird.fly will output: I can fly!

module Flyable
  def fly
    puts "I can fly!"
  end
end

class Bird
  include Flyable
end

bird = Bird.new
bird.fly
