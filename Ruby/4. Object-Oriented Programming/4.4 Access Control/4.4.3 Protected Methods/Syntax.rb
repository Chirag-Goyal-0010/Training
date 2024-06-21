class ClassName
  def public_method(other)
    other.protected_method
  end

  protected

  def protected_method
    # Accessible within the same class or subclasses
  end
end
