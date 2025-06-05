# child.access_protected(parent) will output: Protected method in parent class

class Parent
  protected

  def protected_method
    puts "Protected method in parent class"
  end
end

class Child < Parent
  def access_protected(other)
    other.protected_method
  end
end

parent = Parent.new
child = Child.new
child.access_protected(parent)
