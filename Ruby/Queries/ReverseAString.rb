str = gets.chomp
length = str.length
rev_str = ""
(0...length).each do |i|
    rev_str += str[length-1-i]
end
puts rev_str