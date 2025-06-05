# ghost.spooky will output: Method spooky not found

class Ghost
  def method_missing(method_name, *args)
    puts "Method #{method_name} not found"
  end
end

ghost = Ghost.new
ghost.spooky
