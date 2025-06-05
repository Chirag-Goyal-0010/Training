# Basic Blog Application

This example demonstrates a simple blog application with posts and comments.

## Features

- Create, read, update, and delete posts
- Add comments to posts
- Basic user authentication
- Simple styling with Bootstrap

## Setup

1. Create a new Rails application:
```bash
rails new blog_app
cd blog_app
```

2. Add required gems to `Gemfile`:
```ruby
gem 'bootstrap'
gem 'devise'
```

3. Install dependencies:
```bash
bundle install
```

4. Generate models:
```bash
rails g model Post title:string content:text user:references
rails g model Comment content:text post:references user:references
rails g devise User
```

5. Run migrations:
```bash
rails db:migrate
```

## Code Structure

### Models

```ruby
# app/models/post.rb
class Post < ApplicationRecord
  belongs_to :user
  has_many :comments, dependent: :destroy
  validates :title, presence: true
  validates :content, presence: true
end

# app/models/comment.rb
class Comment < ApplicationRecord
  belongs_to :post
  belongs_to :user
  validates :content, presence: true
end

# app/models/user.rb
class User < ApplicationRecord
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :validatable
  has_many :posts
  has_many :comments
end
```

### Controllers

```ruby
# app/controllers/posts_controller.rb
class PostsController < ApplicationController
  before_action :authenticate_user!, except: [:index, :show]
  before_action :set_post, only: [:show, :edit, :update, :destroy]

  def index
    @posts = Post.all.order(created_at: :desc)
  end

  def show
    @comment = Comment.new
  end

  def new
    @post = Post.new
  end

  def create
    @post = current_user.posts.build(post_params)
    if @post.save
      redirect_to @post, notice: 'Post was successfully created.'
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

# app/controllers/comments_controller.rb
class CommentsController < ApplicationController
  before_action :authenticate_user!

  def create
    @post = Post.find(params[:post_id])
    @comment = @post.comments.build(comment_params)
    @comment.user = current_user

    if @comment.save
      redirect_to @post, notice: 'Comment was successfully added.'
    else
      redirect_to @post, alert: 'Error adding comment.'
    end
  end

  private

  def comment_params
    params.require(:comment).permit(:content)
  end
end
```

### Views

```erb
<%# app/views/posts/index.html.erb %>
<h1>Blog Posts</h1>

<% if user_signed_in? %>
  <%= link_to 'New Post', new_post_path, class: 'btn btn-primary' %>
<% end %>

<% @posts.each do |post| %>
  <article class="post">
    <h2><%= link_to post.title, post %></h2>
    <p class="meta">
      Posted by <%= post.user.email %> on <%= post.created_at.strftime("%B %d, %Y") %>
    </p>
    <p><%= truncate(post.content, length: 200) %></p>
    <%= link_to 'Read more', post %>
  </article>
<% end %>

<%# app/views/posts/show.html.erb %>
<article class="post">
  <h1><%= @post.title %></h1>
  <p class="meta">
    Posted by <%= @post.user.email %> on <%= @post.created_at.strftime("%B %d, %Y") %>
  </p>
  <div class="content">
    <%= @post.content %>
  </div>
</article>

<section class="comments">
  <h2>Comments</h2>
  <%= render @post.comments %>

  <% if user_signed_in? %>
    <h3>Add a comment:</h3>
    <%= render 'comments/form' %>
  <% end %>
</section>

<%# app/views/comments/_form.html.erb %>
<%= form_with(model: [@post, @comment]) do |f| %>
  <div class="field">
    <%= f.label :content %>
    <%= f.text_area :content %>
  </div>

  <div class="actions">
    <%= f.submit 'Add Comment', class: 'btn btn-primary' %>
  </div>
<% end %>
```

### Routes

```ruby
# config/routes.rb
Rails.application.routes.draw do
  devise_for :users
  resources :posts do
    resources :comments, only: [:create]
  end
  root 'posts#index'
end
```

## Styling

Add Bootstrap to your application:

```ruby
# app/assets/stylesheets/application.scss
@import "bootstrap";

# app/javascript/application.js
import "bootstrap"
```

## Testing

```ruby
# test/models/post_test.rb
class PostTest < ActiveSupport::TestCase
  test "should not save post without title" do
    post = Post.new
    assert_not post.save
  end
end

# test/controllers/posts_controller_test.rb
class PostsControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get posts_url
    assert_response :success
  end
end
```

## Next Steps

1. Add categories to posts
2. Implement post search
3. Add user profiles
4. Add image upload
5. Implement post pagination 