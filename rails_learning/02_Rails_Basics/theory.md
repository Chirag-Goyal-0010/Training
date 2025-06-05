# Rails Basics - Code Examples

## Environment Configuration

```ruby
# config/environments/development.rb
Rails.application.configure do
  config.cache_classes = false
  config.eager_load = false
  config.consider_all_requests_local = true
  config.server_timing = true
end

# config/database.yml
development:
  adapter: postgresql
  database: myapp_development
  username: postgres
  password: password
```

## Database Setup

```ruby
# db/migrate/YYYYMMDDHHMMSS_create_posts.rb
class CreatePosts < ActiveRecord::Migration[7.0]
  def change
    create_table :posts do |t|
      t.string :title
      t.text :content
      t.references :user, foreign_key: true
      t.timestamps
    end
  end
end

# db/seeds.rb
Post.create!([
  { title: 'First Post', content: 'Hello World' },
  { title: 'Second Post', content: 'Rails is awesome' }
])
```

## Asset Pipeline

```ruby
# app/assets/javascripts/application.js
//= require jquery
//= require bootstrap
//= require_tree .

# app/assets/stylesheets/application.scss
@import "bootstrap";
@import "custom";

# app/assets/images/logo.png
# app/assets/fonts/custom-font.woff
```

## Rails Generators

```bash
# Generate model with attributes
rails g model Post title:string content:text published:boolean

# Generate controller with actions
rails g controller Posts index show new edit

# Generate scaffold
rails g scaffold Post title:string content:text

# Generate migration
rails g migration AddUserToPosts user:references
```

## Rails Console

```ruby
# Start console
rails console

# Create records
Post.create(title: 'New Post', content: 'Content')

# Query records
Post.where(published: true)
Post.find_by(title: 'New Post')

# Update records
post = Post.first
post.update(title: 'Updated Title')

# Delete records
Post.last.destroy
```

## Rails Server

```bash
# Start server
rails server

# Start on specific port
rails server -p 3001

# Start in specific environment
rails server -e production
```

## Routes

```ruby
# config/routes.rb
Rails.application.routes.draw do
  # RESTful routes
  resources :posts do
    resources :comments
  end

  # Custom routes
  get 'about', to: 'pages#about'
  post 'contact', to: 'pages#contact'

  # Nested routes
  resources :users do
    resources :posts, only: [:index, :new, :create]
  end

  # Namespace routes
  namespace :admin do
    resources :posts
  end
end
```

## Controllers

```ruby
# app/controllers/posts_controller.rb
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
      redirect_to @post, notice: 'Post created'
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

## Models

```ruby
# app/models/post.rb
class Post < ApplicationRecord
  # Associations
  belongs_to :user
  has_many :comments, dependent: :destroy

  # Validations
  validates :title, presence: true
  validates :content, length: { minimum: 10 }

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

  def notify_admin
    AdminMailer.new_post(self).deliver_later
  end
end
```

## Views

```erb
<%# app/views/posts/index.html.erb %>
<h1>Posts</h1>

<% @posts.each do |post| %>
  <article>
    <h2><%= post.title %></h2>
    <p><%= truncate(post.content, length: 100) %></p>
    <%= link_to 'Read more', post_path(post) %>
  </article>
<% end %>

<%# app/views/posts/_form.html.erb %>
<%= form_with(model: @post) do |f| %>
  <% if @post.errors.any? %>
    <div class="errors">
      <h2><%= pluralize(@post.errors.count, "error") %> prohibited this post from being saved:</h2>
      <ul>
        <% @post.errors.full_messages.each do |msg| %>
          <li><%= msg %></li>
        <% end %>
      </ul>
    </div>
  <% end %>

  <div class="field">
    <%= f.label :title %>
    <%= f.text_field :title %>
  </div>

  <div class="field">
    <%= f.label :content %>
    <%= f.text_area :content %>
  </div>

  <%= f.submit %>
<% end %>
```

## Helpers

```ruby
# app/helpers/posts_helper.rb
module PostsHelper
  def format_date(date)
    date.strftime("%B %d, %Y")
  end

  def post_status_badge(post)
    content_tag :span, post.status, class: "badge #{post.status}"
  end
end
```

## Forms

```erb
<%# app/views/posts/new.html.erb %>
<%= form_with(model: @post, local: true) do |f| %>
  <%= f.label :title %>
  <%= f.text_field :title %>

  <%= f.label :content %>
  <%= f.text_area :content %>

  <%= f.label :category %>
  <%= f.collection_select :category_id, Category.all, :id, :name %>

  <%= f.label :tags %>
  <%= f.collection_check_boxes :tag_ids, Tag.all, :id, :name %>

  <%= f.submit 'Create Post' %>
<% end %>
```

## Validation

```ruby
# app/models/post.rb
class Post < ApplicationRecord
  # Presence validation
  validates :title, presence: true

  # Length validation
  validates :content, length: { minimum: 10, maximum: 1000 }

  # Format validation
  validates :email, format: { with: URI::MailTo::EMAIL_REGEXP }

  # Uniqueness validation
  validates :slug, uniqueness: true

  # Custom validation
  validate :title_must_be_capitalized

  private

  def title_must_be_capitalized
    if title.present? && title[0] != title[0].upcase
      errors.add(:title, "must be capitalized")
    end
  end
end
```

## Testing

```ruby
# test/models/post_test.rb
class PostTest < ActiveSupport::TestCase
  test "should not save post without title" do
    post = Post.new
    assert_not post.save
  end

  test "should save valid post" do
    post = Post.new(title: "Test Post", content: "Valid content")
    assert post.save
  end
end

# test/controllers/posts_controller_test.rb
class PostsControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get posts_url
    assert_response :success
  end

  test "should create post" do
    assert_difference('Post.count') do
      post posts_url, params: { post: { title: "Test", content: "Content" } }
    end
    assert_redirected_to post_url(Post.last)
  end
end
``` 