# How do you use next to skip to the next iteration in an until loop?

a = 0
until a == 7 do
    a += 1
    next if a == 5
    puts "#{a}"
end