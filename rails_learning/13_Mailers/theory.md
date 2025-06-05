# Introduction to Mailers Theory

Mailers in Rails are classes that allow you to send emails from your application. They are designed with a structure similar to controllers, with methods that define specific emails to be sent and associated views that contain the email content.

## What are Mailers?

Mailers are Ruby classes that inherit from `ApplicationMailer`. Each public method defined in a mailer class represents a specific email that can be sent. These methods are often called "mailer actions" or "email methods."

```ruby
# app/mailers/user_mailer.rb

class UserMailer < ApplicationMailer
  # Method to send a welcome email
  def welcome_email
    @user = params[:user] # Access data passed to the mailer method
    @url  = 'http://example.com/login'
    mail(to: @user.email, subject: 'Welcome to My Awesome Site')
  end

  # Method to send a password reset email
  def password_reset_email
    @user = params[:user]
    @token = params[:token]
    @url = 'http://example.com/reset_password'
    mail(to: @user.email, subject: 'Password Reset Instructions')
  end
end
```

## How Mailers Work

The process of sending an email using a Rails mailer generally involves the following steps:

1.  **Triggering the Mailer:** You call a mailer method (e.g., `UserMailer.with(user: @user).welcome_email`) from your controllers, models, or other parts of your application. This doesn't send the email immediately but returns a `Mail::Message` object.
2.  **Setting up Email Data:** Inside the mailer method, you set instance variables (e.g., `@user`, `@url`) that will be available in the mailer view. You also use the `mail` method to define the email's recipients (`to`, `cc`, `bcc`), sender (`from`), subject, and other headers.
3.  **Rendering Mailer Views:** Rails looks for view templates that correspond to the mailer class and method name. By default, it will look for `app/views/user_mailer/welcome_email.html.erb` (for HTML content) and `app/views/user_mailer/welcome_email.text.erb` (for plain text content). The content of these views becomes the body of the email.
4.  **Delivering the Email:** After the `Mail::Message` object is created, you call the `deliver_now` or `deliver_later` method on it to send the email immediately or enqueue it for background processing.

```ruby
# Example of triggering a mailer from a controller

def create
  @user = User.new(user_params)
  if @user.save
    # Send the welcome email after successful user creation
    UserMailer.with(user: @user).welcome_email.deliver_later
    redirect_to @user, notice: 'User created successfully. Welcome email sent.'
  else
    render :new
  end
end
```

Mailers are a structured and convenient way to handle email sending in Rails, keeping email logic separate from other parts of your application. The next topic will cover how to generate mailer classes and define email methods within them.

# Generating and Defining Mailers Theory

In Rails, you typically generate mailer classes using a built-in generator. This generator creates the mailer class file and sets up the basic structure. After generation, you define individual methods within the class, each representing a specific email you want to send.

## Generating a Mailer

Use the `rails generate mailer` command followed by the name of your mailer (singular or plural, convention is often singular, like `UserMailer`) and optionally, the names of the first few email methods you want to define within it.

```bash
rails generate mailer UserMailer
# This will create app/mailers/user_mailer.rb and test/mailers/user_mailer_test.rb

rails generate mailer OrderMailer new_order order_confirmation
# This will create app/mailers/order_mailer.rb, test/mailers/order_mailer_test.rb,
# and placeholder view directories: app/views/order_mailer/new_order/ and app/views/order_mailer/order_confirmation/
```

## Defining Mailer Methods (Email Actions)

Each public method you define within a mailer class represents a single type of email that your application can send. Inside these methods, you typically:

1.  **Access Data:** Retrieve any necessary data for the email (e.g., user object, order details). You can pass data to the mailer method using `with`.
2.  **Set Instance Variables:** Assign data to instance variables (e.g., `@user`, `@order`) so they are accessible in the mailer views.
3.  **Use the `mail` Method:** Call the `mail` method to define the email's metadata, such as recipients, sender, subject, and any headers.

```ruby
# app/mailers/user_mailer.rb

class UserMailer < ApplicationMailer
  # Default sender for all emails from this mailer
  default from: 'notifications@example.com'

  # Email method to send a welcome email
  # Data like the user object is passed using #with when calling the mailer
  # Example call: UserMailer.with(user: @user).welcome_email.deliver_later
  def welcome_email
    @user = params[:user] # Access the user object passed via #with
    @url  = 'http://example.com/login'
    
    # Define email metadata using the mail method
    mail(to: @user.email, subject: 'Welcome to My Awesome Site')
    # Rails will look for views in app/views/user_mailer/welcome_email.*
  end

  # Email method for a password reset notification
  # Example call: UserMailer.with(user: @user, token: @token).password_reset_email.deliver_now
  def password_reset_email
    @user = params[:user]
    @token = params[:token]
    @url = 'http://example.com/reset_password?token=' + @token
    
    # Define email metadata
    mail(to: @user.email, subject: 'Password Reset Instructions')
    # Rails will look for views in app/views/user_mailer/password_reset_email.*
  end
end
```

*   **`default`:** You can set default options (like the `from` address) for all emails sent by a specific mailer class.
*   **`params`:** Use the `params` hash within the mailer method to access data passed to the mailer via the `with` method chain.
*   **`mail`:** This is the core method to define the email. The `to`, `subject`, `from`, `cc`, `bcc`, and `reply_to` options are common.

Defining clear and focused mailer methods helps keep your email logic organized and easy to manage. The next topic will cover creating the views that contain the actual content of your emails.

# Sending Emails Theory

After defining your mailer class, email methods, and views, the final step is to trigger the mailer to send the email. You typically call mailer methods from your controllers, models, or background jobs. Calling a mailer method does not send the email immediately; instead, it returns a `Mail::Message` object, which represents the email to be sent. You then call a delivery method on this object to initiate sending.

## Triggering Mailer Methods

You call mailer methods on the mailer class, often chaining the `with` method to pass data to the mailer method.

```ruby
# Calling the mailer method (returns a Mail::Message object)
mail_object = UserMailer.with(user: @user, token: @token).password_reset_email

# This mail_object can now be delivered
```

## Delivery Methods

Rails provides two primary methods for delivering emails:

### `deliver_now`

The `deliver_now` method sends the email synchronously, immediately when the method is called. This means your application will wait for the email to be sent before continuing with the rest of the code execution.

*   **Use cases:** Sending emails that must be delivered immediately and where waiting for delivery doesn't impact user experience significantly (e.g., critical error notifications).
*   **Syntax:**
    ```ruby
    UserMailer.with(user: @user).welcome_email.deliver_now
    ```

### `deliver_later`

The `deliver_later` method sends the email asynchronously, by enqueuing a job to be processed in the background. This is the recommended approach for most emails in a web application, as it prevents the user from waiting while the email is being sent, improving application responsiveness.

*   **Use cases:** Sending most types of emails, such as welcome emails, notifications, and order confirmations, where immediate delivery is not strictly necessary and performance is important.
*   **Requirements:** `deliver_later` requires an Active Job backend configured (e.g., Sidekiq, Delayed Job, or the built-in Async adapter for development/testing).
*   **Syntax:**
    ```ruby
    UserMailer.with(user: @user).welcome_email.deliver_later
    ```

## Passing Parameters with `with`

The `with` method is the preferred way to pass data to your mailer methods. The data passed to `with` is available in the mailer method's `params` hash.

```ruby
# In your controller or service object
@user = User.find(params[:id])
@order = @user.orders.find(params[:order_id])

# Pass user and order to the mailer
OrderMailer.with(user: @user, order: @order).order_confirmation.deliver_later
```

```ruby
# In your mailer method (e.g., OrderMailer#order_confirmation)
def order_confirmation
  @user = params[:user]
  @order = params[:order]
  
  mail(to: @user.email, subject: "Your Order Confirmation (##{@order.id})")
end
```

Using `deliver_later` in conjunction with a background job processor is generally the best practice for sending emails in production Rails applications to avoid blocking the main application thread. The next topic will cover configuring your mailer settings.

# Configuring Mailers Theory

To send emails from your Rails application, you need to configure Action Mailer's delivery settings. These settings tell Rails how and where to send emails. Configurations are typically done in the environment files (`config/environments/development.rb`, `config/environments/production.rb`, `config/environments/test.rb`).

## Common Configuration Options

Mailer configurations are set within the `Rails.application.configure do ... end` block using `config.action_mailer`.

*   `config.action_mailer.delivery_method`: Specifies how emails should be delivered. Common options include:
    *   `:smtp`: Sends emails via an SMTP server (most common for production).
    *   `:sendmail`: Sends emails using the local Sendmail program.
    *   `:test`: Does not send emails but stores them in `ActionMailer::Base.deliveries` (useful for testing).
    *   `:async`: Delivers emails using Active Job (default for `deliver_later`).
    *   `:log`: Logs email details instead of sending (useful for debugging).

*   `config.action_mailer.smtp_settings`: A hash of settings for the `:smtp` delivery method (e.g., address, port, domain, user_name, password, authentication, enable_starttls_auto).

*   `config.action_mailer.default_url_options`: A hash to set the default host and port for generating URLs within mailer views (important for links like password reset).

*   `config.action_mailer.perform_caching`: Whether mailer views should be cached.

*   `config.action_mailer.raise_delivery_errors`: Whether to raise errors if email delivery fails.

## Configuration by Environment

Mailer configurations often differ between environments:

### Development (`config/environments/development.rb`)

In development, you usually don't want to send real emails. Common configurations include using the `:test` or `:async` delivery methods, or using a gem like Letter Opener to open emails in your browser instead of sending them.

```ruby
# config/environments/development.rb

Rails.application.configure do
  # ... other settings ...

  # Don't actually send emails in development. Show them in browser with Letter Opener gem.
  config.action_mailer.delivery_method = :letter_opener
  config.action_mailer.perform_deliveries = true # Ensure delivery is attempted

  # Default URL options for mailer views
  config.action_mailer.default_url_options = { host: 'localhost', port: 3000 }

  # Raise an error if email delivery cannot be attempted
  config.action_mailer.raise_delivery_errors = true

  # ... rest of development config ...
end
```
*(Note: Using Letter Opener requires adding the gem to your Gemfile).* Alternatively, you might use `:async` with the built-in Async adapter or even `:smtp` if you have a local mail server configured.

### Test (`config/environments/test.rb`)

In the test environment, you almost always use the `:test` delivery method. This prevents emails from being sent during test runs and stores them in `ActionMailer::Base.deliveries` for inspection.

```ruby
# config/environments/test.rb

Rails.application.configure do
  # ... other settings ...

  # The test delivery method stores emails in the deliveries array.
  config.action_mailer.delivery_method = :test
  config.action_mailer.perform_deliveries = true
  config.action_mailer.raise_delivery_errors = true
  config.action_mailer.default_url_options = { host: 'localhost', port: 3000 }

  # ... rest of test config ...
end
```

### Production (`config/environments/production.rb`)

In production, you configure Rails to use a real email sending service, typically via SMTP. You'll need credentials and server details from your email provider (e.g., SendGrid, Mailgun, Amazon SES, or your own SMTP server).

```ruby
# config/environments/production.rb

Rails.application.configure do
  # ... other settings ...

  # Configure for a real SMTP server
  config.action_mailer.delivery_method = :smtp
  config.action_mailer.smtp_settings = {
    address:              ENV['SMTP_HOST'],     # e.g., 'smtp.sendgrid.net'
    port:                 ENV['SMTP_PORT'],     # e.g., 587
    domain:               ENV['SMTP_DOMAIN'],   # e.g., 'your-app.com'
    user_name:            ENV['SMTP_USERNAME'], # Your SMTP username
    password:             ENV['SMTP_PASSWORD'], # Your SMTP password
    authentication:       :plain,
    enable_starttls_auto: true
  }
  
  # Default URL options for mailer views
  config.action_mailer.default_url_options = { host: ENV['PRODUCTION_HOST'] } # e.g., 'your-app.com'

  # Ensure delivery errors are logged or handled appropriately
  config.action_mailer.raise_delivery_errors = true

  # By default, emails are delivered later in production via Active Job
  # config.action_mailer.perform_deliveries = true # This is usually true by default in production

  # ... rest of production config ...
end
```
*(Note: It's best practice to store sensitive credentials like SMTP username and password in environment variables, as shown above).* You might also use service-specific gems or configuration methods depending on your email provider.

Properly configuring mailers for each environment is essential to ensure emails are sent correctly in production while providing a convenient development and testing experience. The next section will provide syntax examples for these configurations. 

# Testing Mailers Theory

Testing mailers in Rails involves verifying that your mailer classes generate emails correctly. This includes checking the recipients, sender, subject, headers, and the content of both the HTML and plain text bodies. Rails provides a straightforward way to test mailers by capturing sent emails in the test environment.

## How Mailer Testing Works

In the `test` environment (`config/environments/test.rb`), Action Mailer is typically configured with `config.action_mailer.delivery_method = :test`. This setting intercepts emails and prevents them from being sent over the network. Instead, the generated `Mail::Message` objects are stored in a special array: `ActionMailer::Base.deliveries`.

Your mailer tests will:

1.  Trigger the mailer method you want to test.
2.  Inspect the `ActionMailer::Base.deliveries` array to find the email that was just generated.
3.  Make assertions about the attributes and content of that email object.

## Writing Mailer Tests

Mailer tests are usually located in `test/mailers/` and inherit from `ActionMailer::TestCase`.

```ruby
# test/mailers/user_mailer_test.rb

require "test_helper"

class UserMailerTest < ActionMailer::TestCase
  # Clear deliveries before each test to avoid interference
  setup do
    ActionMailer::Base.deliveries.clear
  end

  test "welcome_email" do
    # Assume a user fixture or factory exists
    user = users(:one)

    # Trigger the mailer method. Use #deliver_now or #deliver_later.
    # In the test environment, both will add the email to ActionMailer::Base.deliveries.
    email = UserMailer.with(user: user).welcome_email.deliver_now

    # Assertions about the sent email

    # Assert that one email was sent
    assert_equal 1, ActionMailer::Base.deliveries.size

    # Assert the recipient
    assert_equal [user.email], email.to

    # Assert the sender (from the mailer or default config)
    assert_equal ['notifications@example.com'], email.from

    # Assert the subject
    assert_equal 'Welcome to My Awesome Site', email.subject

    # Assert content in the HTML body
    assert_match "<h1>Welcome, #{user.name}!</h1>", email.html_part.body.to_s
    assert_match "Login to Your Account", email.html_part.body.to_s # Check for link text

    # Assert content in the plain text body
    assert_match "Welcome, #{user.name}!", email.text_part.body.to_s
    assert_match "http://example.com/login", email.text_part.body.to_s # Check for URL
  end

  test "password_reset_email" do
    user = users(:one)
    token = "fake_token"

    email = UserMailer.with(user: user, token: token).password_reset_email.deliver_now

    assert_equal 1, ActionMailer::Base.deliveries.size
    assert_equal [user.email], email.to
    assert_equal ['notifications@example.com'], email.from
    assert_equal 'Password Reset Instructions', email.subject
    assert_match "reset_password?token=fake_token", email.html_part.body.to_s
    assert_match "reset_password?token=fake_token", email.text_part.body.to_s
  end

  # ... add more tests for other mailer methods ...
end
```

*   **`ActionMailer::Base.deliveries`:** This array holds all the `Mail::Message` objects for emails sent during the test.
*   **`email.to`, `email.from`, `email.subject`:** Access email metadata.
*   **`email.html_part.body.to_s`, `email.text_part.body.to_s`:** Access the content of the HTML and plain text bodies as strings for assertions.
*   **`assert_match`:** Useful for checking if a string (like part of your email body) is present within another string.

Testing mailers gives you confidence that your email communication with users is working correctly. This concludes the Mailers section. The next major section will cover **Background Jobs**. 