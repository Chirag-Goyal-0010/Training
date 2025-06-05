# responder.say_hello will output: You called a dynamic method: say_hello

class DynamicResponder
  def method_missing(method_name, *args)
    if method_name.to_s.start_with?('say_')
      puts "You called a dynamic method: #{method_name}"
    else
      super
    end
  end
end

responder = DynamicResponder.new
responder.say_hello
responder.jump
