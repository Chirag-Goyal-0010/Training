# Ruby Variables and Data Types Examples

# Variables
local_var = "local"
@instance_var = "instance"
@@class_var = "class"
$global_var = "global"
CONSTANT = "constant"

# Data Types
string = "Hello"
integer = 42
float = 3.14
boolean = true
array = [1, 2, 3]
hash = { key: "value" }
symbol = :symbol

# String operations
puts string.upcase
puts string.downcase
puts string.reverse

# Array operations
puts array.first
puts array.last
puts array.push(4)
puts array.pop

# Hash operations
puts hash[:key]
hash[:new_key] = "new value"
puts hash.keys
puts hash.values 