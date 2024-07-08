puts "entyer string"
string = gets.chomp
array = []
array = string.split(' ')
array = array.reverse
puts array
string = array.join(' ')
puts string