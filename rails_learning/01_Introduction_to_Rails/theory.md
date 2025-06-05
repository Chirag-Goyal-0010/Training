# Introduction to Rails - Theory

## What is Ruby on Rails?

Ruby on Rails (often just "Rails") is a web application framework written in Ruby. It follows the Model-View-Controller (MVC) architectural pattern and emphasizes the use of well-known software engineering patterns and paradigms, including:

- Convention over Configuration (CoC)
- Don't Repeat Yourself (DRY)
- Active Record pattern
- RESTful architecture

## MVC Architecture

Rails follows the Model-View-Controller (MVC) pattern, which separates an application into three main components:

### Model (M)
- Represents the data and business logic
- Handles database interactions
- Contains validation rules
- Manages relationships between data

### View (V)
- Represents the user interface
- Displays data to users
- Handles user input
- Renders templates

### Controller (C)
- Processes user requests
- Coordinates between models and views
- Handles routing
- Manages application flow

## Convention over Configuration

Rails emphasizes convention over configuration, which means:

- Follows standard naming conventions
- Uses default configurations
- Reduces the need for explicit configuration
- Makes the code more predictable and maintainable

### Key Conventions
- File naming and organization
- Database table naming
- URL routing
- Controller and model naming

## Rails Directory Structure

A typical Rails application has the following structure:

```
app/                    # Application code
├── controllers/       # Controller classes
├── models/           # Model classes
├── views/            # View templates
├── helpers/          # View helpers
├── mailers/          # Mailer classes
├── jobs/             # Background jobs
└── assets/           # JavaScript, CSS, images

config/               # Configuration files
db/                  # Database files
lib/                 # Library modules
test/                # Test files
```

## Development Environment

### Prerequisites
- Ruby (latest stable version)
- RubyGems
- Node.js and Yarn
- Database (PostgreSQL, MySQL, or SQLite)

### Key Tools
- Rails CLI
- Bundler
- Database management tools
- Version control (Git)

## Common Rails Commands

Essential commands for Rails development:

```bash
# Create new application
rails new my_app

# Start server
rails server

# Generate components
rails generate model Post title:string
rails generate controller Posts index show

# Database operations
rails db:create
rails db:migrate
rails db:seed

# Console
rails console
```

## Best Practices

1. **Follow Conventions**
   - Use standard naming conventions
   - Follow Rails directory structure
   - Use built-in generators

2. **Keep Code DRY**
   - Use partials for views
   - Create reusable helpers
   - Extract common logic to modules

3. **Security First**
   - Use strong parameters
   - Implement proper authentication
   - Follow security best practices

4. **Testing**
   - Write tests from the start
   - Use test-driven development
   - Maintain good test coverage

# Rails Core Concepts - Code Perspective

## MVC Pattern in Code

```ruby
# Model (app/models/post.rb)
class Post < ApplicationRecord
  belongs_to :user
  has_many :comments
  validates :title, presence: true
end

# View (app/views/posts/show.html.erb)
<h1><%= @post.title %></h1>
<p><%= @post.content %></p>

# Controller (app/controllers/posts_controller.rb)
class PostsController < ApplicationController
  def show
    @post = Post.find(params[:id])
  end
end
```

## Rails Conventions

```ruby
# File Naming
app/models/user.rb          # User model
app/controllers/users_controller.rb  # UsersController
app/views/users/index.html.erb       # Users index view

# Database Tables
users                # users table
posts                # posts table
comments             # comments table

# URL Routes
/users               # GET    /users
/users/new          # GET    /users/new
/users/:id          # GET    /users/:id
```

## Directory Structure

```
app/
├── controllers/     # Controllers
├── models/         # Models
├── views/          # Views
├── helpers/        # View helpers
├── mailers/        # Mailers
├── jobs/           # Background jobs
└── assets/         # JS, CSS, images

config/             # Configuration
db/                # Database
lib/               # Libraries
test/              # Tests
```

## Essential Commands

```bash
# Create app
rails new my_app

# Generate components
rails g model Post title:string
rails g controller Posts index show
rails g scaffold Post title:string

# Database
rails db:create
rails db:migrate
rails db:seed

# Server
rails server
rails console
```

## Key Concepts in Code

### Routes
```ruby
# config/routes.rb
Rails.application.routes.draw do
  resources :posts
  root 'posts#index'
end
```

### Controllers
```ruby
class PostsController < ApplicationController
  def index
    @posts = Post.all
  end

  def create
    @post = Post.new(post_params)
    @post.save ? redirect_to(@post) : render(:new)
  end

  private
  def post_params
    params.require(:post).permit(:title, :content)
  end
end
```

### Models
```ruby
class Post < ApplicationRecord
  # Associations
  belongs_to :user
  has_many :comments

  # Validations
  validates :title, presence: true
  validates :content, length: { minimum: 10 }

  # Scopes
  scope :published, -> { where(published: true) }
  scope :recent, -> { order(created_at: :desc) }
end
```

### Views
```erb
<%# app/views/posts/index.html.erb %>
<% @posts.each do |post| %>
  <h2><%= post.title %></h2>
  <p><%= post.content %></p>
  <%= link_to 'Show', post_path(post) %>
<% end %>
```

### Forms
```erb
<%= form_with(model: @post) do |f| %>
  <%= f.label :title %>
  <%= f.text_field :title %>
  
  <%= f.label :content %>
  <%= f.text_area :content %>
  
  <%= f.submit %>
<% end %>
```

### Helpers
```ruby
# app/helpers/posts_helper.rb
module PostsHelper
  def format_date(date)
    date.strftime("%B %d, %Y")
  end
end
```

### Assets
```ruby
# app/assets/javascripts/application.js
//= require jquery
//= require bootstrap

# app/assets/stylesheets/application.scss
@import "bootstrap";
```

### Testing
```ruby
# test/models/post_test.rb
class PostTest < ActiveSupport::TestCase
  test "should not save post without title" do
    post = Post.new
    assert_not post.save
  end
end
``` 