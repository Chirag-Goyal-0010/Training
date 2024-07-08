puts "input first array giving spaces in between"
input = gets.chomp
array_1 = input.split(' ')
array_1 = array_1.map(&:to_i)
puts "input second array giving spaces in between"
input = gets.chomp
array_2 = input.split(' ')
array_2 = array_2.map(&:to_i)
array = array_1 + array_2
puts "after merging array one and two \n #{array}"