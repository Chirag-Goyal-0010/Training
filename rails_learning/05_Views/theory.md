# Rails Views - Code Examples

## View Templates

```ruby
# Basic ERB template
<%# app/views/posts/index.html.erb %>
<h1>Posts</h1>

<% @posts.each do |post| %>
  <div class="post">
    <h2><%= post.title %></h2>
    <p><%= post.content %></p>
    <p>By <%= post.user.name %></p>
  </div>
<% end %>

# Layout template
<%# app/views/layouts/application.html.erb %>
<!DOCTYPE html>
<html>
  <head>
    <title>Blog</title>
    <%= csrf_meta_tags %>
    <%= stylesheet_link_tag 'application' %>
    <%= javascript_include_tag 'application' %>
  </head>
  <body>
    <%= render 'shared/navigation' %>
    <%= render 'shared/flash' %>
    <%= yield %>
    <%= render 'shared/footer' %>
  </body>
</html>

# Partial template
<%# app/views/posts/_post.html.erb %>
<div class="post">
  <h2><%= post.title %></h2>
  <p><%= post.content %></p>
  <p>By <%= post.user.name %></p>
  <%= render 'posts/comments', post: post %>
</div>
```

## View Helpers

```ruby
# Form helpers
<%= form_for @post do |f| %>
  <%= f.label :title %>
  <%= f.text_field :title %>
  
  <%= f.label :content %>
  <%= f.text_area :content %>
  
  <%= f.submit %>
<% end %>

# URL helpers
<%= link_to 'View Post', post_path(@post) %>
<%= link_to 'Edit Post', edit_post_path(@post) %>
<%= link_to 'Delete Post', post_path(@post), method: :delete %>

# Asset helpers
<%= stylesheet_link_tag 'application' %>
<%= javascript_include_tag 'application' %>
<%= image_tag 'logo.png' %>
<%= video_tag 'video.mp4' %>
<%= audio_tag 'audio.mp3' %>

# Text helpers
<%= truncate(@post.content, length: 100) %>
<%= simple_format(@post.content) %>
<%= pluralize(@post.comments.count, 'comment') %>
<%= time_ago_in_words(@post.created_at) %>
```

## View Testing

```ruby
# Integration test
require 'test_helper'

class PostsTest < ActionDispatch::IntegrationTest
  test "should display posts" do
    get posts_path
    assert_response :success
    assert_select 'h1', 'Posts'
    assert_select '.post', minimum: 1
  end
  
  test "should create new post" do
    get new_post_path
    assert_response :success
    
    assert_difference 'Post.count' do
      post posts_path, params: {
        post: {
          title: 'Test Post',
          content: 'Test Content'
        }
      }
    end
    
    assert_redirected_to post_path(Post.last)
  end
end

# System test
require 'application_system_test_case'

class PostsTest < ApplicationSystemTestCase
  test "creating a post" do
    visit posts_path
    click_on 'New Post'
    
    fill_in 'Title', with: 'Test Post'
    fill_in 'Content', with: 'Test Content'
    click_on 'Create Post'
    
    assert_text 'Post was successfully created'
    assert_text 'Test Post'
  end
end
```

## View Security

```ruby
# XSS protection
<%= @post.title %>                    # Safe by default
<%= raw @post.content %>              # Unsafe
<%= @post.content.html_safe %>        # Unsafe
<%= sanitize @post.content %>         # Safe
<%= strip_tags @post.content %>       # Safe

# CSRF protection
<%= form_for @post do |f| %>          # Includes CSRF token
  <%= f.text_field :title %>
<% end %>

<%= form_tag posts_path do %>         # Includes CSRF token
  <%= text_field_tag :title %>
<% end %>

# Content security policy
# config/initializers/content_security_policy.rb
Rails.application.config.content_security_policy do |policy|
  policy.default_src :self
  policy.font_src    :self, :https, :data
  policy.img_src     :self, :https, :data
  policy.object_src  :none
  policy.script_src  :self, :https
  policy.style_src   :self, :https
end
```

## View Performance

```ruby
# Fragment caching
<%# app/views/posts/index.html.erb %>
<% @posts.each do |post| %>
  <% cache post do %>
    <%= render 'posts/post', post: post %>
  <% end %>
<% end %>

# Russian doll caching
<%# app/views/posts/show.html.erb %>
<% cache @post do %>
  <h1><%= @post.title %></h1>
  <p><%= @post.content %></p>
  
  <% @post.comments.each do |comment| %>
    <% cache comment do %>
      <%= render 'comments/comment', comment: comment %>
    <% end %>
  <% end %>
<% end %>

# ETag support
<%# app/controllers/posts_controller.rb %>
def show
  @post = Post.find(params[:id])
  fresh_when(@post)
end

# Conditional GET
<%# app/controllers/posts_controller.rb %>
def show
  @post = Post.find(params[:id])
  if stale?(last_modified: @post.updated_at)
    respond_to do |format|
      format.html
      format.json { render json: @post }
    end
  end
end
``` 