# Basic Rails Routing Example

This example demonstrates fundamental Rails routing concepts using a blog post system.

## Setup

1. Create a new Rails application:
```bash
rails new blog
cd blog
```

2. Generate the necessary models:
```bash
rails generate model User name:string email:string
rails generate model Post title:string content:text user:references
rails generate model Comment content:text user:references post:references
rails generate model Tag name:string
rails generate model PostTag post:references tag:references
```

3. Run the migrations:
```bash
rails db:migrate
```

## Code Structure

### Routes Configuration

```ruby
# config/routes.rb
Rails.application.routes.draw do
  # Root route
  root 'posts#index'
  
  # Basic routes
  get 'about', to: 'pages#about'
  get 'contact', to: 'pages#contact'
  
  # User routes
  resources :users, only: [:index, :show, :new, :create] do
    resources :posts, only: [:index]
  end
  
  # Post routes
  resources :posts do
    member do
      get 'preview'
      post 'publish'
      delete 'unpublish'
    end
    
    collection do
      get 'search'
      get 'popular'
      get 'recent'
    end
    
    resources :comments, only: [:create, :destroy]
    resources :tags, only: [:index]
  end
  
  # Tag routes
  resources :tags, only: [:index, :show]
  
  # Admin namespace
  namespace :admin do
    resources :posts
    resources :users
    resources :comments
  end
  
  # API namespace
  namespace :api do
    namespace :v1 do
      resources :posts
      resources :users
    end
  end
end
```

### Controllers

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
      redirect_to @post, notice: 'Post was successfully created.'
    else
      render :new
    end
  end
  
  def edit
    @post = Post.find(params[:id])
  end
  
  def update
    @post = Post.find(params[:id])
    if @post.update(post_params)
      redirect_to @post, notice: 'Post was successfully updated.'
    else
      render :edit
    end
  end
  
  def destroy
    @post = Post.find(params[:id])
    @post.destroy
    redirect_to posts_url, notice: 'Post was successfully deleted.'
  end
  
  def preview
    @post = Post.find(params[:id])
  end
  
  def publish
    @post = Post.find(params[:id])
    @post.update(published: true)
    redirect_to @post, notice: 'Post was successfully published.'
  end
  
  def unpublish
    @post = Post.find(params[:id])
    @post.update(published: false)
    redirect_to @post, notice: 'Post was successfully unpublished.'
  end
  
  def search
    @posts = Post.search(params[:q])
  end
  
  def popular
    @posts = Post.popular
  end
  
  def recent
    @posts = Post.recent
  end
  
  private
  
  def post_params
    params.require(:post).permit(:title, :content, :user_id)
  end
end
```

### Views

```erb
<%# app/views/posts/index.html.erb %>
<h1>Posts</h1>

<div class="actions">
  <%= link_to 'New Post', new_post_path, class: 'button' %>
  <%= link_to 'Popular Posts', popular_posts_path, class: 'button' %>
  <%= link_to 'Recent Posts', recent_posts_path, class: 'button' %>
</div>

<div class="search">
  <%= form_tag search_posts_path, method: :get do %>
    <%= text_field_tag :q, params[:q], placeholder: 'Search posts...' %>
    <%= submit_tag 'Search' %>
  <% end %>
</div>

<div class="posts">
  <% @posts.each do |post| %>
    <div class="post">
      <h2><%= link_to post.title, post_path(post) %></h2>
      <p><%= truncate(post.content, length: 200) %></p>
      <div class="actions">
        <%= link_to 'Preview', preview_post_path(post) %>
        <%= link_to 'Edit', edit_post_path(post) %>
        <%= link_to 'Delete', post_path(post), method: :delete, data: { confirm: 'Are you sure?' } %>
        <% unless post.published? %>
          <%= link_to 'Publish', publish_post_path(post), method: :post %>
        <% else %>
          <%= link_to 'Unpublish', unpublish_post_path(post), method: :delete %>
        <% end %>
      </div>
    </div>
  <% end %>
</div>
```

```erb
<%# app/views/posts/show.html.erb %>
<div class="post">
  <h1><%= @post.title %></h1>
  <p class="meta">
    By <%= link_to @post.user.name, user_path(@post.user) %>
    on <%= @post.created_at.strftime('%B %d, %Y') %>
  </p>
  
  <div class="content">
    <%= @post.content %>
  </div>
  
  <div class="tags">
    <% @post.tags.each do |tag| %>
      <%= link_to tag.name, tag_path(tag), class: 'tag' %>
    <% end %>
  </div>
  
  <div class="actions">
    <%= link_to 'Edit', edit_post_path(@post) %>
    <%= link_to 'Delete', post_path(@post), method: :delete, data: { confirm: 'Are you sure?' } %>
    <% unless @post.published? %>
      <%= link_to 'Publish', publish_post_path(@post), method: :post %>
    <% else %>
      <%= link_to 'Unpublish', unpublish_post_path(@post), method: :delete %>
    <% end %>
  </div>
</div>

<div class="comments">
  <h2>Comments</h2>
  
  <%= render @post.comments %>
  
  <h3>Add a Comment</h3>
  <%= render 'comments/form', post: @post, comment: Comment.new %>
</div>
```

### Testing

```ruby
# test/routes/posts_test.rb
require "test_helper"

class PostsRoutesTest < ActionDispatch::IntegrationTest
  test "routes to index" do
    assert_routing "/posts", controller: "posts", action: "index"
  end
  
  test "routes to show" do
    assert_routing "/posts/1", controller: "posts", action: "show", id: "1"
  end
  
  test "routes to new" do
    assert_routing "/posts/new", controller: "posts", action: "new"
  end
  
  test "routes to create" do
    assert_routing({ method: "post", path: "/posts" },
                  { controller: "posts", action: "create" })
  end
  
  test "routes to edit" do
    assert_routing "/posts/1/edit",
                  controller: "posts", action: "edit", id: "1"
  end
  
  test "routes to update" do
    assert_routing({ method: "patch", path: "/posts/1" },
                  { controller: "posts", action: "update", id: "1" })
  end
  
  test "routes to destroy" do
    assert_routing({ method: "delete", path: "/posts/1" },
                  { controller: "posts", action: "destroy", id: "1" })
  end
  
  test "routes to preview" do
    assert_routing "/posts/1/preview",
                  controller: "posts", action: "preview", id: "1"
  end
  
  test "routes to publish" do
    assert_routing({ method: "post", path: "/posts/1/publish" },
                  { controller: "posts", action: "publish", id: "1" })
  end
  
  test "routes to unpublish" do
    assert_routing({ method: "delete", path: "/posts/1/unpublish" },
                  { controller: "posts", action: "unpublish", id: "1" })
  end
  
  test "routes to search" do
    assert_routing "/posts/search", controller: "posts", action: "search"
  end
  
  test "routes to popular" do
    assert_routing "/posts/popular", controller: "posts", action: "popular"
  end
  
  test "routes to recent" do
    assert_routing "/posts/recent", controller: "posts", action: "recent"
  end
end
```

## Key Features Demonstrated

1. **Basic Routes**
   - Root route
   - Static pages
   - Custom routes

2. **Resource Routes**
   - Full CRUD resources
   - Limited resources
   - Nested resources

3. **Route Options**
   - Member routes
   - Collection routes
   - Namespaces
   - Scopes

4. **Route Testing**
   - Basic route tests
   - Parameter tests
   - Method tests

## Next Steps

1. Add authentication using Devise
2. Implement role-based access control
3. Add API versioning
4. Implement rate limiting
5. Add route caching
6. Implement route monitoring
7. Add more comprehensive tests 