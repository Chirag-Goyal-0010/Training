puts "Enter a string"
str_1 = gets.chomp
len = str_1.length
indicator = 0
(1...len).each do |i|
    if str_1[i] != str_1[len-i-1]
        puts "Not palindrome"
        indicator += 1
        break
    end
end
puts "Palindrome" if indicator == 0