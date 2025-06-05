# Ruby on Rails Cheatsheet

## Common Commands

### Application
```bash
# Create new Rails application
rails new app_name

# Start Rails server
rails server
# or
rails s

# Start Rails console
rails console
# or
rails c

# Generate database tables
rails db:migrate

# Reset database
rails db:reset

# Create database
rails db:create

# Drop database
rails db:drop
```

### Generators
```bash
# Generate model
rails generate model ModelName field1:type field2:type

# Generate controller
rails generate controller ControllerName action1 action2

# Generate scaffold
rails generate scaffold ModelName field1:type field2:type

# Generate migration
rails generate migration AddFieldToTable field:type
```

## Routes
```ruby
# config/routes.rb

# Basic route
get 'welcome/index'

# Resource routes
resources :posts

# Nested resources
resources :posts do
  resources :comments
end

# Custom routes
get 'about', to: 'pages#about'
```

## Models
```ruby
# app/models/post.rb
class Post < ApplicationRecord
  # Validations
  validates :title, presence: true
  validates :content, length: { minimum: 10 }

  # Associations
  belongs_to :user
  has_many :comments
  has_many :tags, through: :post_tags

  # Scopes
  scope :published, -> { where(published: true) }
  scope :recent, -> { order(created_at: :desc) }
end
```

## Controllers
```ruby
# app/controllers/posts_controller.rb
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

## Views
```erb
<%# app/views/posts/index.html.erb %>

<%# Iteration %>
<% @posts.each do |post| %>
  <h2><%= post.title %></h2>
  <p><%= post.content %></p>
<% end %>

<%# Forms %>
<%= form_with(model: @post) do |f| %>
  <%= f.label :title %>
  <%= f.text_field :title %>

  <%= f.label :content %>
  <%= f.text_area :content %>

  <%= f.submit %>
<% end %>
```

## Active Record Queries
```ruby
# Find
Post.find(1)
Post.find_by(title: "Hello")
Post.where(published: true)

# Create
Post.create(title: "Hello", content: "World")
post = Post.new
post.title = "Hello"
post.save

# Update
post.update(title: "New Title")
post.update_attributes(title: "New Title")

# Delete
post.destroy
Post.destroy_all
```

## Helpers
```ruby
# Link helpers
link_to "Home", root_path
link_to "Edit", edit_post_path(@post)

# Form helpers
form_with(model: @post)
text_field_tag :title
text_area_tag :content
select_tag :category, options_for_select(@categories)

# Asset helpers
image_tag "logo.png"
javascript_include_tag "application"
stylesheet_link_tag "application"
```

## Testing
```ruby
# Model test
class PostTest < ActiveSupport::TestCase
  test "should not save post without title" do
    post = Post.new
    assert_not post.save
  end
end

# Controller test
class PostsControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get posts_url
    assert_response :success
  end
end
```

## Deployment
```bash
# Heroku deployment
heroku create
git push heroku master
heroku run rails db:migrate

# Capistrano deployment
cap production deploy
```

## Debugging
```ruby
# Console debugging
debugger
binding.pry

# Logging
Rails.logger.debug "Debug message"
Rails.logger.info "Info message"
Rails.logger.error "Error message"
``` 