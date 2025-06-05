# Rails Basic Syntax

## Model Syntax

```ruby
# Basic Model
class Post < ApplicationRecord
end

# With Associations
class Post < ApplicationRecord
  belongs_to :user
  has_many :comments
  has_and_belongs_to_many :tags
end

# With Validations
class Post < ApplicationRecord
  validates :title, presence: true
  validates :content, length: { minimum: 10 }
  validates :email, format: { with: URI::MailTo::EMAIL_REGEXP }
end

# With Scopes
class Post < ApplicationRecord
  scope :published, -> { where(published: true) }
  scope :recent, -> { order(created_at: :desc) }
  scope :search, ->(query) { where("title LIKE ?", "%#{query}%") }
end
```

## Controller Syntax

```ruby
# Basic Controller
class PostsController < ApplicationController
  def index
    @posts = Post.all
  end
end

# With Before Actions
class PostsController < ApplicationController
  before_action :set_post, only: [:show, :edit, :update, :destroy]
  before_action :authenticate_user!, except: [:index, :show]

  private
  def set_post
    @post = Post.find(params[:id])
  end
end

# With Strong Parameters
class PostsController < ApplicationController
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
    params.require(:post).permit(:title, :content, :category_id)
  end
end
```

## View Syntax

```erb
<%# Basic View %>
<h1><%= @post.title %></h1>
<p><%= @post.content %></p>

<%# With Conditionals %>
<% if @post.published? %>
  <span class="status">Published</span>
<% else %>
  <span class="status">Draft</span>
<% end %>

<%# With Loops %>
<% @posts.each do |post| %>
  <article>
    <h2><%= post.title %></h2>
    <p><%= post.content %></p>
  </article>
<% end %>

<%# With Partials %>
<%= render 'shared/header' %>
<%= render @posts %>
<%= render partial: 'post', collection: @posts %>
```

## Form Syntax

```erb
<%# Basic Form %>
<%= form_with(model: @post) do |f| %>
  <%= f.label :title %>
  <%= f.text_field :title %>
  
  <%= f.label :content %>
  <%= f.text_area :content %>
  
  <%= f.submit %>
<% end %>

<%# With Nested Attributes %>
<%= form_with(model: @post) do |f| %>
  <%= f.fields_for :comments do |comment_form| %>
    <%= comment_form.label :content %>
    <%= comment_form.text_area :content %>
  <% end %>
<% end %>

<%# With File Upload %>
<%= form_with(model: @post, multipart: true) do |f| %>
  <%= f.file_field :image %>
<% end %>
```

## Route Syntax

```ruby
# Basic Routes
Rails.application.routes.draw do
  resources :posts
end

# Nested Routes
Rails.application.routes.draw do
  resources :users do
    resources :posts
  end
end

# Custom Routes
Rails.application.routes.draw do
  get 'about', to: 'pages#about'
  post 'contact', to: 'pages#contact'
  get 'search', to: 'search#index'
end

# Namespace Routes
Rails.application.routes.draw do
  namespace :admin do
    resources :posts
  end
end
```

## Helper Syntax

```ruby
# Basic Helper
module PostsHelper
  def format_date(date)
    date.strftime("%B %d, %Y")
  end
end

# With HTML
module PostsHelper
  def post_status_badge(post)
    content_tag :span, post.status, class: "badge #{post.status}"
  end
end

# With Links
module PostsHelper
  def edit_post_link(post)
    link_to 'Edit', edit_post_path(post), class: 'btn btn-primary'
  end
end
```

## Migration Syntax

```ruby
# Create Table
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

# Add Column
class AddPublishedToPosts < ActiveRecord::Migration[7.0]
  def change
    add_column :posts, :published, :boolean, default: false
  end
end

# Add Index
class AddIndexToPosts < ActiveRecord::Migration[7.0]
  def change
    add_index :posts, :title
  end
end
```

## Test Syntax

```ruby
# Model Test
class PostTest < ActiveSupport::TestCase
  test "should not save post without title" do
    post = Post.new
    assert_not post.save
  end
end

# Controller Test
class PostsControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get posts_url
    assert_response :success
  end
end

# System Test
class PostTest < ApplicationSystemTestCase
  test "creating a post" do
    visit posts_path
    click_on "New Post"
    fill_in "Title", with: "Test Post"
    fill_in "Content", with: "Test Content"
    click_on "Create Post"
    assert_text "Post was successfully created"
  end
end
``` 