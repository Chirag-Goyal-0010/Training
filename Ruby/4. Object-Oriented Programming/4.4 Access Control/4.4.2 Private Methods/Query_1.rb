# secret.reveal will output: This is a secret.

class Secret
  def reveal
    secret_message
  end

  private

  def secret_message
    puts "This is a secret."
  end
end

secret = Secret.new
secret.reveal
