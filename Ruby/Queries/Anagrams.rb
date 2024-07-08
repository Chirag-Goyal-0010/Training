def sort_array(str)
    (0...str.length).each do |i|
        (0...(str.length - i - 1)).each do |j|
            if str[j] > str[j+1]
                temp = str[j]
                str[j] = str[j+1]
                str[j+1] = temp
            end
        end
    end
end
puts "enter first string"
str_1 = gets.chomp
puts "enter 2nd string"
str_2 = gets.chomp
sort_array(str_1)
sort_array(str_2)
if str_1 == str_2
    puts "string is anagrams"
else 
    puts "not anagrams"
end