# Ruby on Rails Code Examples

## 🔹 1. Core Ruby

### Variables, Data Types, and Operators
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

# Operators
# Arithmetic: +, -, *, /, %, **
# Comparison: ==, !=, >, <, >=, <=
# Logical: &&, ||, !
# Assignment: =, +=, -=, *=, /=
```

### Methods and Blocks
```ruby
# Method definition
def greet(name)
  "Hello, #{name}!"
end

# Method with default parameters
def greet_with_time(name, time = Time.now)
  "Good #{time.hour < 12 ? 'morning' : 'afternoon'}, #{name}!"
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

### Modules and Mixins
```ruby
module Payable
  def calculate_pay(hours)
    @hourly_rate * hours
  end
end

class Employee
  include Payable
  attr_accessor :hourly_rate
end
```

### Inheritance and Polymorphism
```ruby
class Animal
  def speak
    raise NotImplementedError
  end
end

class Dog < Animal
  def speak
    "Woof!"
  end
end

class Cat < Animal
  def speak
    "Meow!"
  end
end
```

### Exception Handling
```ruby
begin
  # Risky code
  result = 10 / 0
rescue ZeroDivisionError => e
  puts "Error: #{e.message}"
rescue StandardError => e
  puts "Unexpected error: #{e.message}"
ensure
  puts "This always runs"
end
```

### Ruby Gems and Bundler
```ruby
# Gemfile
source 'https://rubygems.org'

gem 'rails', '~> 7.0.0'
gem 'devise'
gem 'pundit'

# Install gems
bundle install

# Update gems
bundle update
```

## 🔹 2. Rails Basics

### Rails Directory Structure
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

### Common Rails Commands
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

# Database
rails db:create
rails db:migrate
rails db:seed
```

## 🔹 3. Rails Routing

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

### RESTful Routes
```ruby
resources :posts
# Creates:
# GET    /posts          # index
# GET    /posts/new      # new
# POST   /posts          # create
# GET    /posts/:id      # show
# GET    /posts/:id/edit # edit
# PATCH  /posts/:id      # update
# DELETE /posts/:id      # destroy
```

### Custom Routes
```ruby
# HTTP Verb Routes
get 'about', to: 'pages#about'
post 'contact', to: 'pages#contact'
patch 'update', to: 'pages#update'
delete 'remove', to: 'pages#remove'

# Named Routes
get 'login', to: 'sessions#new', as: :login
get 'logout', to: 'sessions#destroy', as: :logout

# Route Constraints
get 'users/:id', to: 'users#show', constraints: { id: /[0-9]+/ }
```

### Nested Resources
```ruby
resources :posts do
  resources :comments
  resources :likes, only: [:create, :destroy]
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

### Filters
```ruby
class ApplicationController < ActionController::Base
  before_action :authenticate_user!
  before_action :set_locale
  skip_before_action :verify_authenticity_token, only: [:create]
end
```

### Params Handling
```ruby
# Strong Parameters
def post_params
  params.require(:post).permit(:title, :content, :category_id)
end

# Nested Parameters
def user_params
  params.require(:user).permit(:name, :email, 
    addresses_attributes: [:id, :street, :city, :_destroy])
end
```

## 🔹 5. Models

### ActiveRecord Models
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

### Migrations
```ruby
# Generate migration
rails generate migration AddTitleToPosts title:string

# Migration file
class AddTitleToPosts < ActiveRecord::Migration[7.0]
  def change
    add_column :posts, :title, :string
    add_index :posts, :title
  end
end
```

### Model Relationships
```ruby
# One-to-One
class User < ApplicationRecord
  has_one :profile
end

class Profile < ApplicationRecord
  belongs_to :user
end

# One-to-Many
class Post < ApplicationRecord
  belongs_to :user
  has_many :comments
end

# Many-to-Many
class Post < ApplicationRecord
  has_many :post_tags
  has_many :tags, through: :post_tags
end
```

### Validations
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
```

### Layouts and Partials
```erb
<%# app/views/layouts/application.html.erb %>
<!DOCTYPE html>
<html>
  <head>
    <title>My App</title>
    <%= csrf_meta_tags %>
    <%= stylesheet_link_tag 'application' %>
  </head>
  <body>
    <%= render 'shared/header' %>
    <%= yield %>
    <%= render 'shared/footer' %>
  </body>
</html>

<%# app/views/shared/_header.html.erb %>
<header>
  <nav>
    <%= link_to 'Home', root_path %>
    <%= link_to 'About', about_path %>
  </nav>
</header>
```

### Helpers
```ruby
# app/helpers/posts_helper.rb
module PostsHelper
  def format_date(date)
    date.strftime("%B %d, %Y")
  end
end

# View
<%= format_date(@post.created_at) %>
```

### View Rendering
```ruby
# Controller actions
def index
  @posts = Post.all
  respond_to do |format|
    format.html
    format.json { render json: @posts }
  end
end

def show
  @post = Post.find(params[:id])
  render :show, status: :ok
end
```

### Flash Messages
```ruby
# Controller
redirect_to @post, notice: 'Post created!'
redirect_to @post, alert: 'Error creating post!'

# View
<% if notice %>
  <div class="notice"><%= notice %></div>
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

### HTTP Verbs and Status Codes
```ruby
# Common Status Codes
:ok                    # 200
:created              # 201
:no_content           # 204
:bad_request          # 400
:unauthorized         # 401
:forbidden            # 403
:not_found            # 404
:unprocessable_entity # 422
:internal_server_error # 500
```

### CRUD Operations
```ruby
# Create
Post.create(title: "New Post")

# Read
Post.find(1)
Post.find_by(title: "Hello")
Post.where(published: true)

# Update
post.update(title: "Updated Title")

# Delete
post.destroy
```

## 🔹 8. Forms and User Input

### Form Helpers
```ruby
# form_with (preferred)
<%= form_with(model: @post) do |f| %>
  <%= f.label :title %>
  <%= f.text_field :title %>
  
  <%= f.label :content %>
  <%= f.text_area :content %>
  
  <%= f.submit %>
<% end %>

# form_for (legacy)
<%= form_for @post do |f| %>
  # form fields
<% end %>

# form_tag
<%= form_tag posts_path do %>
  <%= text_field_tag :query %>
  <%= submit_tag "Search" %>
<% end %>
```

### Strong Parameters
```ruby
def post_params
  params.require(:post).permit(
    :title, 
    :content, 
    :category_id,
    tags_attributes: [:id, :name, :_destroy]
  )
end
```

### CSRF Protection
```ruby
# In application_controller.rb
protect_from_forgery with: :exception

# In forms
<%= form_with(model: @post) do |f| %>
  <%= csrf_meta_tags %>
  # form fields
<% end %>
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

### Scopes
```ruby
class Post < ApplicationRecord
  scope :published, -> { where(published: true) }
  scope :recent, -> { order(created_at: :desc) }
  scope :popular, -> { joins(:likes).group('posts.id').order('COUNT(likes.id) DESC') }
  scope :by_category, ->(category) { where(category: category) }
end
```

### Callbacks
```ruby
class Post < ApplicationRecord
  before_validation :set_slug
  after_create :notify_admin
  before_destroy :check_dependencies
  
  private
  
  def set_slug
    self.slug = title.parameterize
  end
  
  def notify_admin
    AdminMailer.new_post_notification(self).deliver_later
  end
  
  def check_dependencies
    throw(:abort) if comments.any?
  end
end
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

### API Serialization
```ruby
# Controller
def index
  @posts = Post.all
  render json: PostSerializer.new(@posts).serializable_hash
end

# Serializer with custom methods
class PostSerializer < ActiveModel::Serializer
  attributes :id, :title, :content, :formatted_date
  
  def formatted_date
    object.created_at.strftime("%B %d, %Y")
  end
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

### Custom Validators
```ruby
class EmailValidator < ActiveModel::EachValidator
  def validate_each(record, attribute, value)
    unless value =~ URI::MailTo::EMAIL_REGEXP
      record.errors.add(attribute, 'is not a valid email')
    end
  end
end

class User < ApplicationRecord
  validates :email, email: true
end
```

### Database Constraints
```ruby
class AddNotNullToEmail < ActiveRecord::Migration[7.0]
  def change
    change_column_null :users, :email, false
    add_index :users, :email, unique: true
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

### API Versioning
```ruby
# config/routes.rb
namespace :api do
  namespace :v1 do
    resources :posts
  end
  namespace :v2 do
    resources :posts
  end
end
```

### API Authentication
```ruby
class Api::V1::BaseController < ApplicationController
  before_action :authenticate_api_user!
  
  private
  
  def authenticate_api_user!
    token = request.headers['Authorization']&.split(' ')&.last
    @current_user = User.find_by_api_token(token)
    
    unless @current_user
      render json: { error: 'Unauthorized' }, status: :unauthorized
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

### JWT Authentication
```ruby
# Gemfile
gem 'jwt'

# Controller
class Api::V1::SessionsController < ApplicationController
  def create
    user = User.find_by(email: params[:email])
    if user&.authenticate(params[:password])
      token = JWT.encode({ user_id: user.id }, Rails.application.secrets.secret_key_base)
      render json: { token: token }
    else
      render json: { error: 'Invalid credentials' }, status: :unauthorized
    end
  end
end
```

### Authorization with Pundit
```ruby
# Gemfile
gem 'pundit'

# Policy
class PostPolicy < ApplicationPolicy
  def update?
    user.admin? || record.user_id == user.id
  end
end

# Controller
class PostsController < ApplicationController
  def update
    @post = Post.find(params[:id])
    authorize @post
    # update logic
  end
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

# View
# app/views/user_mailer/welcome_email.html.erb
<h1>Welcome, <%= @user.name %>!</h1>
```

### Background Jobs
```ruby
# Gemfile
gem 'sidekiq'

# Job
class SendWelcomeEmailJob < ApplicationJob
  queue_as :default
  
  def perform(user)
    UserMailer.welcome_email(user).deliver_now
  end
end

# Usage
SendWelcomeEmailJob.perform_later(user)
```

## 🔹 15. Testing

### RSpec Setup
```ruby
# Gemfile
group :development, :test do
  gem 'rspec-rails'
  gem 'factory_bot_rails'
  gem 'faker'
end

# Installation
rails generate rspec:install
```

### Model Specs
```ruby
RSpec.describe Post, type: :model do
  describe 'validations' do
    it { should validate_presence_of(:title) }
    it { should validate_presence_of(:content) }
  end
  
  describe 'associations' do
    it { should belong_to(:user) }
    it { should have_many(:comments) }
  end
end
```

### Controller Specs
```ruby
RSpec.describe PostsController, type: :controller do
  describe 'GET #index' do
    it 'returns a success response' do
      get :index
      expect(response).to be_successful
    end
  end
end
```

### System Specs
```ruby
RSpec.describe 'Post creation', type: :system do
  it 'creates a new post' do
    visit new_post_path
    fill_in 'Title', with: 'Test Post'
    fill_in 'Content', with: 'Test Content'
    click_button 'Create Post'
    expect(page).to have_content('Post was successfully created')
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

### Webpacker
```ruby
# app/javascript/packs/application.js
import Rails from "@rails/ujs"
import Turbolinks from "turbolinks"
import * as bootstrap from "bootstrap"

Rails.start()
Turbolinks.start()
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

# View
<%= form_with(model: @post) do |f| %>
  <%= f.file_field :cover_image %>
  <%= f.file_field :photos, multiple: true %>
<% end %>
```

### Image Processing
```ruby
# Gemfile
gem 'image_processing', '~> 1.2'

# Model
class Post < ApplicationRecord
  has_one_attached :cover_image do |attachable|
    attachable.variant :thumb, resize_to_limit: [100, 100]
    attachable.variant :medium, resize_to_limit: [300, 300]
  end
end

# View
<%= image_tag @post.cover_image.variant(:thumb) %>
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

# JavaScript
const chatChannel = consumer.subscriptions.create(
  { channel: "ChatChannel", room: "general" },
  {
    received(data) {
      // Handle received data
    }
  }
)
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

# SQL Injection Prevention
User.where("name = ?", params[:name])
User.where(name: params[:name])

# XSS Protection
<%= sanitize @post.content %>
<%= strip_tags @post.content %>
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

# Russian Doll caching
<% cache [@post, @post.comments.maximum(:updated_at)] do %>
  <%= render @post %>
<% end %>
```

## 🔹 21. Internationalization

### I18n Setup
```ruby
# config/locales/en.yml
en:
  welcome:
    message: "Welcome to our site!"
    hello: "Hello, %{name}!"

# View
<%= t('welcome.message') %>
<%= t('welcome.hello', name: @user.name) %>

# Controller
I18n.locale = params[:locale] || I18n.default_locale
```

## 🔹 22. Performance Optimization

### Query Optimization
```ruby
# Eager loading
Post.includes(:comments).all

# Counter cache
class Post < ApplicationRecord
  belongs_to :user, counter_cache: true
end

# Database indexes
add_index :posts, :title
add_index :posts, [:user_id, :created_at]
```

### Background Processing
```ruby
# Sidekiq configuration
config.active_job.queue_adapter = :sidekiq

# Job
class ProcessPostJob < ApplicationJob
  queue_as :default
  
  def perform(post_id)
    post = Post.find(post_id)
    # Process post
  end
end
```

## 🔹 23. Deployment

### Heroku Deployment
```bash
# Initial setup
heroku create
git push heroku main

# Database setup
heroku run rails db:migrate
heroku run rails db:seed

# Environment variables
heroku config:set RAILS_ENV=production
heroku config:set SECRET_KEY_BASE=your_secret_key
```

### Capistrano Deployment
```ruby
# Gemfile
gem 'capistrano'
gem 'capistrano-rails'
gem 'capistrano-rbenv'

# config/deploy.rb
set :application, 'my_app'
set :repo_url, 'git@github.com:user/my_app.git'
set :deploy_to, '/var/www/my_app'
```

## 🔹 24. Monitoring & Debugging

### Logging
```ruby
# config/environments/production.rb
config.log_level = :info
config.logger = ActiveSupport::Logger.new(STDOUT)

# Custom logging
Rails.logger.info "Processing post #{@post.id}"
Rails.logger.error "Error: #{error.message}"
```

### Debugging
```ruby
# Console debugging
rails console
rails c

# Debugger
require 'pry'
binding.pry

# Logging
Rails.logger.debug "Debug message"
Rails.logger.info "Info message"
Rails.logger.warn "Warning message"
Rails.logger.error "Error message"
```

### Exception Tracking
```ruby
# Gemfile
gem 'sentry-raven'

# config/initializers/sentry.rb
Raven.configure do |config|
  config.dsn = 'your-sentry-dsn'
  config.environments = ['production']
end
``` 