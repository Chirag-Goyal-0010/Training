def factorial(num)
    return 1 if num == 0 || num == 1
    return num * factorial(num - 1)
end
puts "Enter the no of which factorial require"
fact = gets.to_i
number = factorial(fact)
puts number