# How do you define and access an instance variable within methods?

def set_name (arg1)
    @name = arg1
  end
  
  def print_name
    puts @name
  end
  
  set_name("Chirag Goyal")
  print_name