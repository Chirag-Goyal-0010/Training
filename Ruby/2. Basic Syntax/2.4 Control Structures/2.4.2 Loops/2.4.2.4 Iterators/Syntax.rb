# Each iterator 

# Method 1
(1...4).each do |number|
    puts number
end

puts 

# Method 2
range = 1..7
range.each do |number|
    puts number
end

# Method 3
arr = [1,23,"chirag","%",4.5,true]
arr.each do |arr_element|
    puts arr_element
end


# Map iterator 

arr = [1,2,3,4]
puts  arr.map { |element| element * 2 }