# How do you break out of a while loop early?

a = 10
while a >5 do
    if a == 7      #  break if a == 7 (2nd method)
        break
    end
    puts "#{a}"
    a -= 1
end