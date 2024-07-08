puts "Enter num to check if it is prime or not"
num = gets.to_i
temp = 2
prime = true
while temp < num 
    if num%temp == 0
        puts "Not Prime"
        prime = false
        break
    end
    temp += 1
end
if prime 
    puts "Prime Num"
end