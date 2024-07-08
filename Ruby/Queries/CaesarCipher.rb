def encryption(str,shift)
    (0...str.length).each do |i|
        if str[i].ord > 96
            if (str[i].ord + shift) > 122
                str[i] = (str[i].ord + shift - 26).chr
            else
                str[i] = (str[i].ord + shift).chr
            end
        elsif str[i].ord < 91 && str[i].ord != 32
            if (str[i].ord + shift) > 90
                str[i] = (str[i].ord + shift - 26).chr
            else 
                str[i] = (str[i].ord + shift).chr
            end
        end
    end
    str
end


def decryption(str,shift)
    (0...str.length).each do |i|
        if str[i].ord > 96
            if (str[i].ord - shift) < 97
                str[i] = (str[i].ord - shift + 26).chr
            else
                str[i] = (str[i].ord - shift).chr
            end
        elsif str[i].ord < 91 && str[i].ord != 32
            if (str[i].ord - shift) < 65
                str[i] = (str[i].ord - shift + 26).chr
            else 
                str[i] = (str[i].ord - shift).chr
            end
        end
    end
    str
end

puts "Enter the string"
str = gets.chomp
puts "Enter the shift value"
shift_value = gets.to_i
encryption(str,shift_value)
puts str
decryption(str,shift_value)
puts str