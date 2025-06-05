# Complete Ruby on Rails Cheatsheet

## 🔹 1. Core Ruby

### Variables and Data Types
```ruby
# Variables
local_var = "local"
@instance_var = "instance"
@@class_var = "class"
$global_var = "global"
CONSTANT = "constant"

# Data Types
string = "Hello"
integer = 42
float = 3.14
boolean = true
array = [1, 2, 3]
hash = { key: "value" }
symbol = :symbol
```

### Methods and Blocks
```ruby
# Method definition
def greet(name)
  "Hello, #{name}!"
end

# Block with yield
def with_logging
  puts "Starting..."
  yield
  puts "Finished!"
end

# Lambda
lambda = ->(x) { x * 2 }
```

### Classes and Objects
```ruby
class Person
  attr_accessor :name, :age
  
  def initialize(name, age)
    @name = name
    @age = age
  end
  
  def greet
    "Hello, I'm #{@name}!"
  end
end
```

## 🔹 2. Rails Basics

### Directory Structure
```
app/
├── controllers/    # Controller classes
├── models/        # Model classes
├── views/         # View templates
├── helpers/       # View helpers
├── mailers/       # Mailer classes
├── jobs/          # Background jobs
└── assets/        # JavaScript, CSS, images

config/            # Configuration files
db/               # Database files
lib/              # Library modules
test/             # Test files
```

### Common Commands
```bash
# Create new Rails app
rails new my_app

# Start server
rails server
rails s

# Console
rails console
rails c

# Generate
rails generate model Post title:string content:text
rails generate controller Posts index show
rails generate scaffold Post title:string content:text
```

## 🔹 3. Routing

### Basic Routes
```ruby
# config/routes.rb
Rails.application.routes.draw do
  # RESTful routes
  resources :posts
  
  # Custom routes
  get 'about', to: 'pages#about'
  post 'contact', to: 'pages#contact'
  
  # Nested routes
  resources :posts do
    resources :comments
  end
  
  # Namespace
  namespace :admin do
    resources :posts
  end
end
```

## 🔹 4. Controllers

### Controller Actions
```ruby
class PostsController < ApplicationController
  before_action :set_post, only: [:show, :edit, :update, :destroy]
  
  def index
    @posts = Post.all
  end
  
  def show
  end
  
  def new
    @post = Post.new
  end
  
  def create
    @post = Post.new(post_params)
    if @post.save
      redirect_to @post, notice: 'Post created!'
    else
      render :new
    end
  end
  
  private
  
  def set_post
    @post = Post.find(params[:id])
  end
  
  def post_params
    params.require(:post).permit(:title, :content)
  end
end
```

## 🔹 5. Models

### ActiveRecord
```ruby
class Post < ApplicationRecord
  # Associations
  belongs_to :user
  has_many :comments, dependent: :destroy
  has_many :tags, through: :post_tags
  
  # Validations
  validates :title, presence: true, length: { minimum: 5 }
  validates :content, presence: true
  
  # Scopes
  scope :published, -> { where(published: true) }
  scope :recent, -> { order(created_at: :desc) }
  
  # Callbacks
  before_save :set_slug
  after_create :notify_admin
  
  private
  
  def set_slug
    self.slug = title.parameterize
  end
end
```

## 🔹 6. Views

### ERB Templates
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

<%= form_with(model: @post) do |f| %>
  <%= f.label :title %>
  <%= f.text_field :title %>
  
  <%= f.label :content %>
  <%= f.text_area :content %>
  
  <%= f.submit %>
<% end %>
```

## 🔹 7. RESTful Architecture

### REST Actions
```ruby
# Routes
resources :posts

# Controller actions
def index    # GET /posts
def show     # GET /posts/:id
def new      # GET /posts/new
def create   # POST /posts
def edit     # GET /posts/:id/edit
def update   # PATCH/PUT /posts/:id
def destroy  # DELETE /posts/:id
```

## 🔹 8. Forms and User Input

### Form Helpers
```ruby
# Strong Parameters
def post_params
  params.require(:post).permit(:title, :content, :category_id)
end

# Form Helpers
form_with(model: @post)
form_for(@post)
form_tag(posts_path)
```

## 🔹 9. ActiveRecord Queries

### Query Methods
```ruby
# Finding records
Post.find(1)
Post.find_by(title: "Hello")
Post.where(published: true)
Post.where("created_at > ?", 1.week.ago)

# Chaining
Post.published.recent.limit(10)

# Joins
Post.joins(:comments).where(comments: { approved: true })

# Includes (eager loading)
Post.includes(:comments).all

# Aggregations
Post.count
Post.average(:rating)
Post.group(:category).count
```

## 🔹 10. Serializers

### JSON Serialization
```ruby
# Basic serialization
post.as_json

# Active Model Serializers
class PostSerializer < ActiveModel::Serializer
  attributes :id, :title, :content
  belongs_to :user
  has_many :comments
end

# Jbuilder
json.array! @posts do |post|
  json.id post.id
  json.title post.title
  json.content post.content
end
```

## 🔹 11. Validations

### Model Validations
```ruby
class User < ApplicationRecord
  validates :email, presence: true, 
                   uniqueness: true,
                   format: { with: URI::MailTo::EMAIL_REGEXP }
  validates :password, length: { minimum: 8 }
  validates :age, numericality: { greater_than: 18 }
  
  validate :custom_validation
  
  private
  
  def custom_validation
    if password.present? && password == email
      errors.add(:password, "can't be the same as email")
    end
  end
end
```

## 🔹 12. REST APIs

### API Controller
```ruby
class Api::V1::PostsController < ApplicationController
  def index
    @posts = Post.all
    render json: @posts
  end
  
  def show
    @post = Post.find(params[:id])
    render json: @post
  end
  
  def create
    @post = Post.new(post_params)
    if @post.save
      render json: @post, status: :created
    else
      render json: { errors: @post.errors }, status: :unprocessable_entity
    end
  end
end
```

## 🔹 13. Authentication & Authorization

### Devise Setup
```ruby
# Gemfile
gem 'devise'

# Installation
rails generate devise:install
rails generate devise User
rails db:migrate

# Controller
class ApplicationController < ActionController::Base
  before_action :authenticate_user!
end
```

## 🔹 14. Mailers and Background Jobs

### ActionMailer
```ruby
class UserMailer < ApplicationMailer
  def welcome_email(user)
    @user = user
    mail(to: @user.email, subject: 'Welcome!')
  end
end

# Background Job
class SendWelcomeEmailJob < ApplicationJob
  queue_as :default
  
  def perform(user)
    UserMailer.welcome_email(user).deliver_now
  end
end
```

## 🔹 15. Testing

### RSpec Examples
```ruby
# Model spec
RSpec.describe Post, type: :model do
  it "is valid with valid attributes" do
    expect(Post.new(title: "Test")).to be_valid
  end
end

# Controller spec
RSpec.describe PostsController, type: :controller do
  describe "GET #index" do
    it "returns a success response" do
      get :index
      expect(response).to be_successful
    end
  end
end
```

## 🔹 16. Asset Pipeline

### Asset Management
```ruby
# app/assets/stylesheets/application.scss
@import "bootstrap";
@import "custom";

# app/assets/javascript/application.js
//= require jquery
//= require bootstrap
//= require_tree .
```

## 🔹 17. Active Storage

### File Uploads
```ruby
class Post < ApplicationRecord
  has_one_attached :cover_image
  has_many_attached :photos
end

# Controller
def create
  @post = Post.new(post_params)
  @post.cover_image.attach(params[:post][:cover_image])
  @post.save
end
```

## 🔹 18. Real-Time Features

### ActionCable
```ruby
# app/channels/chat_channel.rb
class ChatChannel < ApplicationCable::Channel
  def subscribed
    stream_from "chat_#{params[:room]}"
  end
  
  def receive(data)
    ActionCable.server.broadcast("chat_#{params[:room]}", data)
  end
end
```

## 🔹 19. Security

### Security Measures
```ruby
# CSRF Protection
protect_from_forgery with: :exception

# Secure Headers
config.action_dispatch.default_headers = {
  'X-Frame-Options' => 'DENY',
  'X-XSS-Protection' => '1; mode=block',
  'X-Content-Type-Options' => 'nosniff'
}
```

## 🔹 20. Caching

### Cache Methods
```ruby
# Fragment caching
<% cache @post do %>
  <%= render @post %>
<% end %>

# Low-level caching
Rails.cache.fetch("posts/#{@post.id}") do
  @post.comments.to_a
end
```

## 🔹 21. Internationalization

### I18n Setup
```ruby
# config/locales/en.yml
en:
  welcome:
    message: "Welcome to our site!"

# View
<%= t('welcome.message') %>
```

## 🔹 22. Performance

### Optimization Techniques
```ruby
# Eager loading
Post.includes(:comments).all

# Database indexes
add_index :posts, :title
add_index :posts, [:user_id, :created_at]

# Counter cache
add_column :posts, :comments_count, :integer, default: 0
```

## 🔹 23. Deployment

### Heroku Deployment
```bash
# Initial setup
heroku create
git push heroku main

# Database
heroku run rails db:migrate

# Environment variables
heroku config:set RAILS_ENV=production
```

## 🔹 24. Monitoring

### Debugging Tools
```ruby
# Console debugging
debugger
binding.pry

# Logging
Rails.logger.debug "Debug message"
Rails.logger.info "Info message"
Rails.logger.error "Error message"
``` 