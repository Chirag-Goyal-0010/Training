def remove_duplicate(array)
    array_duplicate = []
    array.each do |i|
        if array_duplicate.include?(i)
        else 
            array_duplicate << i
        end
    end
    array_duplicate
end
puts "enter array with spaces in betrween elements"
array = gets.chomp.split(' ').map(&:to_i)
array = remove_duplicate(array)
puts array
