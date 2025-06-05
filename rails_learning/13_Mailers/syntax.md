# Mailer Views Syntax

This document provides syntax examples for creating and using mailer views in Rails.

## HTML Mailer View (`.html.erb`)

Located in `app/views/[mailer_name]/[method_name].html.erb`. You can use standard HTML and ERB.

```erb
<%# app/views/user_mailer/welcome_email.html.erb %>

<!DOCTYPE html>
<html>
<head>
  <meta content="text/html; charset=UTF-8" http-equiv="Content-Type" />
  <style>
    /* Add your email-specific CSS here. Consider inlining for compatibility. */
    .container {
      font-family: sans-serif;
      line-height: 1.5;
      color: #333;
    }
    .button {
      display: inline-block;
      background-color: #007bff;
      color: white !important; /* Use !important to override some client styles */
      padding: 10px 20px;
      text-decoration: none; /* Remove underline from links */
      border-radius: 5px;
    }
  </style>
</head>
<body>
  <%# Using a mailer layout (optional) %>
  <div class="container">
    <h1>Welcome, <%= @user.name %>!</h1>

    <p>
      Thank you for joining My Awesome Site. We're excited to have you!
    </p>

    <p>
      To get started, please log in here:
    </p>

    <p>
      <%= link_to 'Login to Your Account', @url, class: 'button' %>
    </p>

    <p>If the button above doesn't work, you can copy and paste the following URL into your browser:</p>
    <p><%= @url %></p>

    <p>Thanks,<br>The My Awesome Site Team</p>
  </div>
</body>
</html>
```

## Plain Text Mailer View (`.text.erb`)

Located in `app/views/[mailer_name]/[method_name].text.erb`. Should contain only plain text.

```erb
<%# app/views/user_mailer/welcome_email.text.erb %>

Welcome, <%= @user.name %>!

Thank you for joining My Awesome Site. We're excited to have you!

To get started, please log in here:
<%= @url %>

Thanks,
The My Awesome Site Team
```

## Accessing Instance Variables

Instance variables set in the mailer method (e.g., `@user`, `@url`) are directly accessible in the views:

```ruby
# In your mailer method (e.g., UserMailer#welcome_email)
def welcome_email
  @user = params[:user]
  @url  = 'http://example.com/login'
  mail(to: @user.email, subject: 'Welcome!')
end
```

```erb
<%# In your view %>
<p>User Name: <%= @user.name %></p>
<p>Login URL: <%= @url %></p>
```

## Using Layouts

Mailer views can use layouts, typically located in `app/views/layouts/mailer.html.erb` and `app/views/layouts/mailer.text.erb`.

```erb
<%# app/views/layouts/mailer.html.erb %>

<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <%# Add global email styles or includes here %>
</head>
<body>
  <%# The content from the specific email view is rendered here %>
  <%= yield %>
</body>
</html>
```

Creating both HTML and plain text views ensures wider compatibility for your emails. The next topic will cover how to trigger your mailers to send these emails.

# Sending Emails Syntax

This document provides syntax examples for triggering mailer methods and sending emails in Rails.

## Calling a Mailer Method

Call the mailer method on the mailer class. This returns a `Mail::Message` object.

```ruby
# In a controller action, service object, or background job
@user = User.find(params[:id])

# Call the mailer method and pass parameters using 'with'
email = UserMailer.with(user: @user).welcome_email

# The 'email' variable now holds the Mail::Message object, but the email hasn't been sent yet.
```

## Sending Immediately (`deliver_now`)

Call `deliver_now` on the `Mail::Message` object to send the email synchronously.

```ruby
# Send the welcome email immediately
UserMailer.with(user: @user).welcome_email.deliver_now

# Example in a controller create action (less common for welcome emails due to blocking)
def create
  @user = User.new(user_params)
  if @user.save
    # This will pause the request until the email is sent
    UserMailer.with(user: @user).welcome_email.deliver_now
    redirect_to @user, notice: 'User created successfully.'
  else
    render :new
  end
end
```

## Sending in the Background (`deliver_later`)

Call `deliver_later` on the `Mail::Message` object to enqueue the email for asynchronous delivery via Active Job.

```ruby
# Enqueue the welcome email for background processing (recommended)
UserMailer.with(user: @user).welcome_email.deliver_later

# Example in a controller create action (preferred for responsiveness)
def create
  @user = User.new(user_params)
  if @user.save
    # Enqueue the email job without blocking the response
    UserMailer.with(user: @user).welcome_email.deliver_later
    redirect_to @user, notice: 'User created successfully. Welcome email enqueued.'
  else
    render :new
  end
end
```

## Passing Parameters with `with`

Always pass necessary data to your mailer methods using the `with` method. The data is accessed via the `params` hash in the mailer method.

```ruby
# In a controller or service
order = Order.find(params[:id])

# Pass multiple parameters
OrderMailer.with(order: order, recipient: order.user, subject: "Your Order Details").order_confirmation.deliver_later
```

```ruby
# In your mailer method (OrderMailer#order_confirmation)
class OrderMailer < ApplicationMailer
  def order_confirmation
    @order = params[:order]
    @recipient = params[:recipient]
    custom_subject = params[:subject] || "Order Confirmation"

    mail(to: @recipient.email, subject: custom_subject)
  end
end
```

Using `deliver_later` is the standard practice in production to ensure your application remains responsive. The next topic will cover configuring your mailer settings, such as the delivery method (e.g., SMTP, SendGrid, etc.).

# Configuring Mailers Syntax

This document provides syntax examples for configuring Action Mailer settings in your Rails environment files (`config/environments/*.rb`).

## Development Environment (`config/environments/development.rb`)

Configuration for development often involves not sending real emails and using tools like Letter Opener.

```ruby
# config/environments/development.rb

Rails.application.configure do
  # ... existing development configurations ...

  # Configure Action Mailer

  # Set the delivery method. :letter_opener opens emails in the browser.
  # Requires the 'letter_opener' gem.
  # config.action_mailer.delivery_method = :letter_opener
  # config.action_mailer.perform_deliveries = true # Set to true to use the delivery method

  # Alternatively, use :async (default for deliver_later in development)
  # config.action_mailer.delivery_method = :async
  # config.action_mailer.perform_deliveries = true

  # Default URL options needed for generating links in emails
  config.action_mailer.default_url_options = { host: 'localhost', port: 3000 }

  # Raise errors if email delivery fails
  config.action_mailer.raise_delivery_errors = true

  # Print email delivery errors to the Rails logger
  config.action_mailer.logger = Logger.new(STDOUT)

  # ... rest of development configurations ...
end
```

## Test Environment (`config/environments/test.rb`)

Configuration for testing typically uses the `:test` delivery method, which stores emails in memory.

```ruby
# config/environments/test.rb

Rails.application.configure do
  # ... existing test configurations ...

  # Configure Action Mailer

  # Use the :test delivery method, which stores emails in ActionMailer::Base.deliveries
  config.action_mailer.delivery_method = :test
  config.action_mailer.perform_deliveries = true # Set to true to store emails

  # Default URL options for tests
  config.action_mailer.default_url_options = { host: 'localhost', port: 3000 }

  # Raise errors during tests
  config.action_mailer.raise_delivery_errors = true

  # ... rest of test configurations ...
end

# In your test files, you can access delivered emails like this:
# ActionMailer::Base.deliveries # Returns an array of Mail::Message objects
# ActionMailer::Base.deliveries.last # Get the last email sent
# ActionMailer::Base.deliveries.clear # Clear the deliveries array before a test
```

## Production Environment (`config/environments/production.rb`)

Configuration for production typically uses the `:smtp` delivery method with details from your email service provider.

```ruby
# config/environments/production.rb

Rails.application.configure do
  # ... existing production configurations ...

  # Configure Action Mailer

  # Use the :smtp delivery method
  config.action_mailer.delivery_method = :smtp
  config.action_mailer.perform_deliveries = true # Set to true to send emails

  # SMTP settings from your email service provider (using environment variables is recommended)
  config.action_mailer.smtp_settings = {
    address:              ENV['SMTP_HOST'],
    port:                 ENV['SMTP_PORT'],
    domain:               ENV['SMTP_DOMAIN'],
    user_name:            ENV['SMTP_USERNAME'],
    password:             ENV['SMTP_PASSWORD'],
    authentication:       :plain,
    enable_starttls_auto: true
  }

  # Default URL options for production (replace with your actual domain)
  config.action_mailer.default_url_options = { host: ENV['PRODUCTION_HOST'] } # e.g., 'www.your-app.com'

  # Ensure delivery errors are handled (e.g., logged)
  config.action_mailer.raise_delivery_errors = true

  # ... rest of production configurations ...
end
```

Properly configuring these settings for each environment is crucial for successful email sending in your Rails application. The final topic in this section covers testing your mailers. 