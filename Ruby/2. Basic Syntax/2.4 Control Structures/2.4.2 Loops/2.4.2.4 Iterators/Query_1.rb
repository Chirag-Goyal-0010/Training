# How do you use the select iterator to filter an array?

arr = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]

# Method 1
puts arr.select {|element| element % 2 == 0}

# Method 2
puts arr.select { |number| number.even? }