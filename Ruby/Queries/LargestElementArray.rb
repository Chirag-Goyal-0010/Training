def large_element(arr , num)
    bubble_sort(arr,num)
    arr.last
end
def bubble_sort(arr,num)
    (0...num).each do |i|
        (0...(num - i - 1)).each do|j|
            if arr[j] > arr[j+1]
                temp = arr[j+1]
                arr[j+1] = arr[j]
                arr[j] = temp
            end
        end
    end
end
puts "no of elements"
element = gets.to_i
arr = []
(0...element).each do|i|
    arr[i] = gets.to_i
end
puts "Largest element of array is #{large_element(arr, element)}"
