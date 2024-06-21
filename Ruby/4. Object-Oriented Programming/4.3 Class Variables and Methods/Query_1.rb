# Dog.species will output: Canine

class Dog
  @@species = "Canine"

  def self.species
    @@species
  end
end

puts Dog.species
