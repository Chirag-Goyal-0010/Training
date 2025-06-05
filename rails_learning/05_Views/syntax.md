# Rails View Syntax

## ERB Tags

```erb
<%# Comments %>
<%# This is a comment %>

<%# Ruby code %>
<% @posts.each do |post| %>
  <%= post.title %>
<% end %>

<%# Output %>
<%= @post.title %>
<%= link_to "Home", root_path %>

<%# Output with HTML safety %>
<%= raw @post.content %>
<%= sanitize @post.content %>
<%= h @post.content %>

<%# Content for %>
<% content_for :title, "Page Title" %>
<%= yield :title %>

<%# Capture %>
<% @greeting = capture do %>
  Hello, <%= @user.name %>!
<% end %>
```

## Layout Syntax

```erb
<%# app/views/layouts/application.html.erb %>
<!DOCTYPE html>
<html>
  <head>
    <title><%= content_for?(:title) ? yield(:title) : "Default Title" %></title>
    <%= csrf_meta_tags %>
    <%= csp_meta_tag %>
    
    <%= stylesheet_link_tag "application" %>
    <%= javascript_include_tag "application" %>
  </head>
  
  <body>
    <%= yield %>
  </body>
</html>

<%# app/views/layouts/admin.html.erb %>
<%= content_for :admin_scripts do %>
  <%= javascript_include_tag "admin" %>
<% end %>
```

## Partial Syntax

```erb
<%# Render with local variables %>
<%= render "posts/post", post: @post %>

<%# Render collection %>
<%= render @posts %>

<%# Render with options %>
<%= render "posts/post", 
    post: @post,
    locals: { show_comments: true },
    layout: "post_layout" %>

<%# Render with block %>
<%= render layout: "post_layout" do %>
  <%= @post.title %>
<% end %>
```

## Form Helpers

```erb
<%# Basic form %>
<%= form_with(model: @post) do |f| %>
  <%= f.text_field :title %>
  <%= f.text_area :content %>
  <%= f.submit %>
<% end %>

<%# Form with options %>
<%= form_with(model: @post,
    local: true,
    html: { class: "post-form" },
    data: { remote: false }) do |f| %>
  <%= f.text_field :title, class: "form-control" %>
  <%= f.text_area :content, rows: 10 %>
  <%= f.submit "Save", class: "btn btn-primary" %>
<% end %>

<%# Form fields %>
<%= f.text_field :title %>
<%= f.text_area :content %>
<%= f.email_field :email %>
<%= f.password_field :password %>
<%= f.number_field :age %>
<%= f.date_field :birthday %>
<%= f.time_field :start_time %>
<%= f.datetime_field :published_at %>
<%= f.color_field :color %>
<%= f.range_field :quantity %>
<%= f.telephone_field :phone %>
<%= f.url_field :website %>
<%= f.search_field :query %>
<%= f.hidden_field :token %>

<%# Form collections %>
<%= f.collection_select :category_id, Category.all, :id, :name %>
<%= f.collection_radio_buttons :category_id, Category.all, :id, :name %>
<%= f.collection_check_boxes :tag_ids, Tag.all, :id, :name %>

<%# Form associations %>
<%= f.fields_for :comments do |comment_form| %>
  <%= comment_form.text_field :content %>
<% end %>
```

## URL Helpers

```erb
<%# Basic links %>
<%= link_to "Home", root_path %>
<%= link_to "Post", post_path(@post) %>
<%= link_to "Edit Post", edit_post_path(@post) %>

<%# Link with options %>
<%= link_to "Delete", post_path(@post),
    method: :delete,
    data: { confirm: "Are you sure?" },
    class: "btn btn-danger" %>

<%# Link with block %>
<%= link_to post_path(@post) do %>
  <h2><%= @post.title %></h2>
  <p><%= @post.excerpt %></p>
<% end %>

<%# Button to %>
<%= button_to "Delete", post_path(@post),
    method: :delete,
    class: "btn btn-danger" %>
```

## Asset Helpers

```erb
<%# Stylesheets %>
<%= stylesheet_link_tag "application" %>
<%= stylesheet_link_tag "application", media: "all" %>
<%= stylesheet_link_tag "application", "custom" %>

<%# JavaScript %>
<%= javascript_include_tag "application" %>
<%= javascript_include_tag "application", defer: true %>

<%# Images %>
<%= image_tag "logo.png" %>
<%= image_tag "logo.png", alt: "Logo" %>
<%= image_tag "logo.png", size: "100x50" %>

<%# Video %>
<%= video_tag "movie.mp4" %>
<%= video_tag "movie.mp4", controls: true %>
<%= video_tag ["movie.mp4", "movie.webm"] %>

<%# Audio %>
<%= audio_tag "music.mp3" %>
<%= audio_tag "music.mp3", controls: true %>
<%= audio_tag ["music.mp3", "music.ogg"] %>
```

## Text Helpers

```erb
<%# Formatting %>
<%= simple_format @post.content %>
<%= truncate @post.content, length: 100 %>
<%= excerpt @post.content, "important", radius: 50 %>

<%# Links %>
<%= auto_link @post.content %>
<%= mail_to "user@example.com" %>
<%= mail_to "user@example.com", "Contact Us" %>

<%# Numbers %>
<%= number_to_currency @product.price %>
<%= number_to_percentage @product.discount %>
<%= number_to_phone @user.phone %>
<%= number_with_delimiter @post.views_count %>

<%# Dates %>
<%= time_ago_in_words @post.created_at %>
<%= distance_of_time_in_words @post.created_at, Time.current %>
<%= time_tag @post.created_at %>
```

## View Testing

```ruby
# test/views/posts/index_test.rb
require "test_helper"

class Posts::IndexTest < ActionView::TestCase
  test "renders posts" do
    @posts = [posts(:one), posts(:two)]
    render
    
    assert_select "h1", "Posts"
    assert_select ".post", count: 2
  end
end

# test/helpers/posts_helper_test.rb
require "test_helper"

class PostsHelperTest < ActionView::TestCase
  test "formats post date" do
    post = posts(:one)
    assert_equal "January 1, 2024", format_post_date(post.created_at)
  end
end
```

## View Security

```erb
<%# CSRF Protection %>
<%= csrf_meta_tags %>
<%= form_authenticity_token %>

<%# Content Security %>
<%= content_security_policy do |policy| %>
  policy.default_src :self
  policy.font_src    :self, :https, :data
  policy.img_src     :self, :https, :data
  policy.script_src  :self, :https
  policy.style_src   :self, :https
<% end %>

<%# XSS Protection %>
<%= sanitize @post.content %>
<%= sanitize @post.content, tags: %w(p br strong em) %>
<%= sanitize @post.content, attributes: %w(href class) %>
```

## View Performance

```erb
<%# Fragment Caching %>
<% cache @post do %>
  <%= render @post %>
<% end %>

<%# Russian Doll Caching %>
<% cache [@post, @post.author] do %>
  <%= render @post %>
<% end %>

<%# Collection Caching %>
<%= render partial: "posts/post", collection: @posts, cached: true %>

<%# Cache Digests %>
<% cache ["v1", @post] do %>
  <%= render @post %>
<% end %>

<%# Cache Keys %>
<% cache [@post, "comments", @post.comments.maximum(:updated_at)] do %>
  <%= render @post.comments %>
<% end %>
``` 