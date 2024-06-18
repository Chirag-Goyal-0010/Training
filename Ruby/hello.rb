# # Ruby program for calculating the Nth Fibonacci number.
# def Fibonacci(number)
 
#     # Base case :  when N is less than 2.
#     if number < 2
#       number
#     else
   
#       # Recursive call : sum of last two Fibonacci's.
#       return Fibonacci(number - 1) + Fibonacci(number - 2)
#     end
#   end
   
# #   puts Fibonacci(2)
  

# # Ruby program of regular expression 

# a="2mnanc3"
# b="2km.5"
# # . literal matches for all character 
# if(a.match(/\d.....\d/)) 
# 	puts("match found") 
# else
# 	puts("not found") 
# end
# # after escaping it, it matches with only '.' literal 
# if(a.match(/\d\.\d/)) 
# 	puts("match found") 
# else
# 	puts("not found") 
# end

# if(b.match(/\d\.\d/)) 
# 	puts("match found") 
# else
# 	puts("not found") 
# end




# Ruby program of sub and gsub method 
text = "geeks for geeks, is a computer science portal"

# Change "rails" to "Rails" throughout 
text.gsub!("geeks", "Geeks") 

# Capitalize the word "Rails" throughout 
text.gsub!(/\bgeeks\b/, "Geeks") 
puts "#{text}"
