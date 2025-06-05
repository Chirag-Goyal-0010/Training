# Hello World Rails Application

This example demonstrates how to create a basic Rails application that displays "Hello, World!".

## Steps to Create

1. Create a new Rails application:
```bash
rails new hello_world
cd hello_world
```

2. Generate a controller for the home page:
```bash
rails generate controller Pages home
```

3. Update the routes in `config/routes.rb`:
```ruby
Rails.application.routes.draw do
  root 'pages#home'
end
```

4. Create the view in `app/views/pages/home.html.erb`:
```erb
<h1>Hello, World!</h1>
<p>Welcome to my first Rails application.</p>
```

5. Start the Rails server:
```bash
rails server
```

6. Visit `http://localhost:3000` in your browser to see the result.

## Files Created

- `app/controllers/pages_controller.rb`: Controller for the home page
- `app/views/pages/home.html.erb`: View template
- `config/routes.rb`: Route configuration

## What We Learned

- Creating a new Rails application
- Generating controllers
- Setting up routes
- Creating views
- Starting the Rails server

## Next Steps

1. Add more pages to the application
2. Style the pages with CSS
3. Add JavaScript functionality
4. Create a database and models 