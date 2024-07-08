string = gets.chomp
string_temp = ''
array = []
increment = 0
string.each_char do |i|
    if i in array
        break;
    else 
        array >> i
        increment += 1
    end
end
puts array
