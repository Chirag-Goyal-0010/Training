# animal.dog_sound will output: Dog makes a sound!

class Animal
  [:dog, :cat, :bird].each do |animal|
    define_method("#{animal}_sound") do
      puts "#{animal.capitalize} makes a sound!"
    end
  end
end

animal = Animal.new
animal.dog_sound
animal.cat_sound
animal.bird_sound
