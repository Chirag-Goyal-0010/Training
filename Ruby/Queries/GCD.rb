def rec_gcd(a,b)
    if a % b == 0
        return b
    end
    c = a % b
    return rec_gcd(b, c)
end
puts "Enter first no."
num_1 = gets.to_i
puts "Enter second no."
num_2 = gets.to_i
if num_1 < num_2
    puts rec_gcd(num_2,num_1)
else 
    puts rec_gcd(num_1,num_2)
end
