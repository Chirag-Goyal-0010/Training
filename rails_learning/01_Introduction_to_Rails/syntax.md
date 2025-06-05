# Introduction to Rails - Syntax

## Basic Rails Syntax

### Creating a New Rails Application
```bash
# Create a new Rails application
rails new my_app

# Create with specific database
rails new my_app --database=postgresql

# Create with specific Rails version
rails _7.0.0_ new my_app
```

### Rails Server
```bash
# Start the Rails server
rails server
# or
rails s

# Start on specific port
rails server -p 3001

# Start in specific environment
rails server -e production
```

### Rails Console
```bash
# Start the Rails console
rails console
# or
rails c

# Start in sandbox mode (changes are rolled back)
rails console --sandbox
```

### Rails Generate
```bash
# Generate a model
rails generate model Post title:string content:text
# or
rails g model Post title:string content:text

# Generate a controller
rails generate controller Posts index show
# or
rails g controller Posts index show

# Generate a scaffold
rails generate scaffold Post title:string content:text
# or
rails g scaffold Post title:string content:text
```

### Database Commands
```bash
# Create database
rails db:create

# Run migrations
rails db:migrate

# Rollback last migration
rails db:rollback

# Seed database
rails db:seed

# Reset database
rails db:reset
```

## File Naming Conventions

### Models
```ruby
# app/models/post.rb
class Post < ApplicationRecord
end

# app/models/comment.rb
class Comment < ApplicationRecord
end
```

### Controllers
```ruby
# app/controllers/posts_controller.rb
class PostsController < ApplicationController
end

# app/controllers/comments_controller.rb
class CommentsController < ApplicationController
end
```

### Views
```
app/views/posts/
├── index.html.erb
├── show.html.erb
├── new.html.erb
├── edit.html.erb
└── _form.html.erb
```

## Basic Controller Syntax
```ruby
class PostsController < ApplicationController
  def index
    @posts = Post.all
  end

  def show
    @post = Post.find(params[:id])
  end

  def new
    @post = Post.new
  end

  def create
    @post = Post.new(post_params)
    if @post.save
      redirect_to @post
    else
      render :new
    end
  end

  private

  def post_params
    params.require(:post).permit(:title, :content)
  end
end
```

## Basic Model Syntax
```ruby
class Post < ApplicationRecord
  # Validations
  validates :title, presence: true
  validates :content, presence: true

  # Associations
  belongs_to :user
  has_many :comments

  # Scopes
  scope :published, -> { where(published: true) }
  scope :recent, -> { order(created_at: :desc) }
end
```

## Basic View Syntax
```erb
<%# app/views/posts/index.html.erb %>
<h1>Posts</h1>

<% @posts.each do |post| %>
  <article>
    <h2><%= post.title %></h2>
    <p><%= post.content %></p>
    <%= link_to 'Read more', post_path(post) %>
  </article>
<% end %>
```

## Routes Syntax
```ruby
# config/routes.rb
Rails.application.routes.draw do
  # RESTful routes
  resources :posts

  # Custom routes
  get 'about', to: 'pages#about'
  post 'contact', to: 'pages#contact'

  # Root route
  root 'posts#index'
end
```

## Helper Methods
```ruby
# app/helpers/posts_helper.rb
module PostsHelper
  def format_date(date)
    date.strftime("%B %d, %Y")
  end
end

# Usage in view
<%= format_date(@post.created_at) %>
```

## Common View Helpers
```erb
<%# Links %>
<%= link_to 'Home', root_path %>
<%= link_to 'Edit', edit_post_path(@post) %>

<%# Forms %>
<%= form_with(model: @post) do |f| %>
  <%= f.label :title %>
  <%= f.text_field :title %>
  
  <%= f.label :content %>
  <%= f.text_area :content %>
  
  <%= f.submit %>
<% end %>

<%# Images %>
<%= image_tag 'logo.png' %>

<%# JavaScript %>
<%= javascript_include_tag 'application' %>

<%# Stylesheets %>
<%= stylesheet_link_tag 'application' %>
``` 