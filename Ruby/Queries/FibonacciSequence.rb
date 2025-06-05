def rec_fabonacci(i)
    return i if i==0 || i==1
    return rec_fabonacci(i-1) + rec_fabonacci(i-2)
end

puts "enter sequence num"
seq_no = gets.to_i
seq_no -= 1
while seq_no >=0 do
    puts "#{rec_fabonacci(seq_no)}"
    seq_no -= 1
end