puts "enter the string under 100 characters"
str = gets.chomp
sum = 0
(0...str.length).each do |i|
    if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' || str[i] == 'A' || str[i] == 'E' || str[i] == 'I' || str[i] == 'O' || str[i] == 'U'
        sum = sum +1
    end
end
puts sum