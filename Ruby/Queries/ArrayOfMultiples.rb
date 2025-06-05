puts "Enter the no of which you want to find factors"
num = gets.to_i
length = 0
arr = []

(1..num).each do |i|
  if num % i == 0
    arr[length] = i
    print "#{', ' if length > 0}#{arr[length]}"
    length += 1
  end
end
