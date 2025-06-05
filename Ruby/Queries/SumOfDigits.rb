def sum(num)
    sum = 0
    digits = 0
    temp = num
    while temp>0
        digit = temp % 10
        sum = sum + digit
        temp = temp / 10
    end
    return sum
end
puts "Enter the num of which sum of digits require"
sum_digits = gets.to_i
number = sum(sum_digits)
puts number