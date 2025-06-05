for a in 1..100 do
    if a%3 == 0 && a%5 == 0
        puts "FizzBuzz \n"
    elsif a%3 == 0
        puts "Fizz \n"
    elsif a%5 == 0
        puts "Buzz \n"
    else 
        puts "#{a} \n"
    end
end