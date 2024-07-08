def isPerfect(num)
    arr = []
    length = 0
    sum = 0
    (1...num).each do |i|
        if num % i == 0
            arr[length] = i
            length += 1
        end
    end
    (0...length).each do |i|
        sum = sum + arr[i]
    end
    if sum == num
        puts "Perfect num"
    else 
        puts "Not Perfect"
    end
end
puts "Enter num"
num = gets.to_i
isPerfect(num)
