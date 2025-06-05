# Metaprogramming in Ruby

## Dynamic Method Definition
```ruby
class DynamicClass
  # Define methods dynamically
  ['add', 'subtract', 'multiply'].each do |operation|
    define_method(operation) do |a, b|
      case operation
      when 'add' then a + b
      when 'subtract' then a - b
      when 'multiply' then a * b
      end
    end
  end
end
```

## Method Missing
```ruby
class MethodMissingExample
  def method_missing(method_name, *args)
    if method_name.to_s.start_with?('find_by_')
      attribute = method_name.to_s.split('_').last
      # Implementation for dynamic finder methods
      "Finding by #{attribute}: #{args.first}"
    else
      super
    end
  end

  def respond_to_missing?(method_name, include_private = false)
    method_name.to_s.start_with?('find_by_') || super
  end
end
```

## Eval and Instance Eval
```ruby
# eval - evaluates a string as Ruby code
code = "puts 'Hello from eval'"
eval(code)

# instance_eval - evaluates code in the context of an object
class Example
  def initialize
    @value = 42
  end
end

obj = Example.new
obj.instance_eval { puts @value }  # => 42
```

## Class Eval and Module Eval
```ruby
class DynamicClass
  # class_eval - evaluates code in the context of a class
  def self.add_method(name, &block)
    class_eval do
      define_method(name, &block)
    end
  end
end

# Usage
DynamicClass.add_method(:greet) { puts "Hello!" }
```

## Method Hooks
```ruby
module MethodHooks
  def self.included(base)
    base.extend(ClassMethods)
  end

  module ClassMethods
    def method_added(method_name)
      puts "Method added: #{method_name}"
    end

    def method_removed(method_name)
      puts "Method removed: #{method_name}"
    end

    def method_undefined(method_name)
      puts "Method undefined: #{method_name}"
    end
  end
end
```

## Dynamic Attributes
```ruby
class DynamicAttributes
  def initialize
    @attributes = {}
  end

  def method_missing(name, *args)
    if name.to_s.end_with?('=')
      @attributes[name.to_s.chop.to_sym] = args.first
    else
      @attributes[name]
    end
  end
end
```

## Class Macros
```ruby
module ClassMacros
  def attr_accessor_with_history(*attrs)
    attrs.each do |attr|
      # Define getter
      define_method(attr) do
        instance_variable_get("@#{attr}")
      end

      # Define setter
      define_method("#{attr}=") do |value|
        instance_variable_set("@#{attr}", value)
        @history ||= {}
        @history[attr] ||= []
        @history[attr] << value
      end

      # Define history getter
      define_method("#{attr}_history") do
        @history[attr]
      end
    end
  end
end
```

## Module Inclusion
```ruby
module IncludedModule
  def self.included(base)
    base.extend(ClassMethods)
    base.class_eval do
      # Add instance methods
      def instance_method
        "Instance method from IncludedModule"
      end
    end
  end

  module ClassMethods
    def class_method
      "Class method from IncludedModule"
    end
  end
end
```

## Method Wrapping
```ruby
module MethodWrapper
  def self.wrap_method(klass, method_name)
    klass.class_eval do
      alias_method :"original_#{method_name}", method_name
      define_method(method_name) do |*args, &block|
        puts "Before #{method_name}"
        result = send(:"original_#{method_name}", *args, &block)
        puts "After #{method_name}"
        result
      end
    end
  end
end
```

## Dynamic Delegation
```ruby
class DelegateExample
  def initialize(target)
    @target = target
  end

  def method_missing(method_name, *args, &block)
    if @target.respond_to?(method_name)
      @target.send(method_name, *args, &block)
    else
      super
    end
  end

  def respond_to_missing?(method_name, include_private = false)
    @target.respond_to?(method_name) || super
  end
end
```

## Best Practices
1. Use metaprogramming sparingly and only when necessary
2. Document metaprogramming code thoroughly
3. Consider performance implications
4. Use method_missing with respond_to_missing?
5. Be careful with eval and similar methods
6. Follow Ruby's principle of least surprise
7. Test metaprogramming code thoroughly
8. Use appropriate method hooks
9. Consider maintainability
10. Use metaprogramming to reduce boilerplate, not to show off 