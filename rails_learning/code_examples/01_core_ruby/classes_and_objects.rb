# Ruby Classes and Objects Examples

# Basic class definition
class Person
  attr_accessor :name, :age
  
  def initialize(name, age)
    @name = name
    @age = age
  end
  
  def greet
    "Hello, I'm #{@name}!"
  end
  
  def adult?
    @age >= 18
  end
end

# Inheritance
class Employee < Person
  attr_accessor :salary, :position
  
  def initialize(name, age, salary, position)
    super(name, age)
    @salary = salary
    @position = position
  end
  
  def work
    "#{@name} is working as a #{@position}"
  end
end

# Module for shared functionality
module Payable
  def calculate_pay(hours)
    @salary * hours
  end
end

# Including module in class
class Contractor < Person
  include Payable
  
  attr_accessor :hourly_rate
  
  def initialize(name, age, hourly_rate)
    super(name, age)
    @hourly_rate = hourly_rate
  end
end

# Class methods and instance methods
class BankAccount
  @@interest_rate = 0.05
  
  def self.interest_rate
    @@interest_rate
  end
  
  def self.interest_rate=(rate)
    @@interest_rate = rate
  end
  
  attr_reader :balance
  
  def initialize(initial_balance = 0)
    @balance = initial_balance
  end
  
  def deposit(amount)
    @balance += amount
  end
  
  def withdraw(amount)
    if amount <= @balance
      @balance -= amount
      amount
    else
      "Insufficient funds"
    end
  end
  
  def add_interest
    @balance += @balance * @@interest_rate
  end
end

# Using the classes
person = Person.new("John", 25)
puts person.greet
puts person.adult?

employee = Employee.new("Alice", 30, 50000, "Developer")
puts employee.work
puts employee.greet

contractor = Contractor.new("Bob", 35, 50)
puts contractor.calculate_pay(40)

account = BankAccount.new(1000)
account.deposit(500)
puts account.balance
account.withdraw(200)
puts account.balance
account.add_interest
puts account.balance 