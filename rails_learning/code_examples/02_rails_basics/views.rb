# app/views/posts/index.html.erb
<h1>Posts</h1>

<div class="search-form">
  <%= form_tag search_posts_path, method: :get do %>
    <%= text_field_tag :query, params[:query], placeholder: "Search posts..." %>
    <%= submit_tag "Search" %>
  <% end %>
</div>

<div class="posts">
  <% @posts.each do |post| %>
    <article class="post">
      <h2><%= link_to post.title, post_path(post) %></h2>
      <div class="meta">
        By <%= post.user.name %> | <%= time_ago_in_words(post.created_at) %> ago
      </div>
      <div class="content">
        <%= truncate(post.content, length: 200) %>
      </div>
      <div class="actions">
        <%= link_to "Read more", post_path(post), class: "button" %>
        <% if current_user == post.user %>
          <%= link_to "Edit", edit_post_path(post), class: "button" %>
          <%= button_to "Delete", post_path(post), method: :delete, 
              data: { confirm: "Are you sure?" }, class: "button danger" %>
        <% end %>
      </div>
    </article>
  <% end %>
</div>

<%= paginate @posts %>

# app/views/posts/_form.html.erb
<%= form_with(model: post, local: true) do |f| %>
  <% if post.errors.any? %>
    <div class="error-messages">
      <h2><%= pluralize(post.errors.count, "error") %> prohibited this post from being saved:</h2>
      <ul>
        <% post.errors.full_messages.each do |message| %>
          <li><%= message %></li>
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
    <%= f.text_area :content, rows: 10 %>
  </div>

  <div class="field">
    <%= f.label :category_id %>
    <%= f.collection_select :category_id, Category.all, :id, :name %>
  </div>

  <div class="field">
    <%= f.label :published %>
    <%= f.check_box :published %>
  </div>

  <div class="actions">
    <%= f.submit class: "button" %>
  </div>
<% end %>

# app/views/layouts/application.html.erb
<!DOCTYPE html>
<html>
  <head>
    <title>Blog App</title>
    <%= csrf_meta_tags %>
    <%= csp_meta_tag %>

    <%= stylesheet_link_tag 'application', media: 'all', 'data-turbolinks-track': 'reload' %>
    <%= javascript_pack_tag 'application', 'data-turbolinks-track': 'reload' %>
  </head>

  <body>
    <header>
      <nav>
        <%= link_to "Home", root_path %>
        <%= link_to "Posts", posts_path %>
        <% if current_user %>
          <%= link_to "New Post", new_post_path %>
          <%= link_to "Profile", user_path(current_user) %>
          <%= button_to "Logout", logout_path, method: :delete %>
        <% else %>
          <%= link_to "Login", login_path %>
          <%= link_to "Sign Up", signup_path %>
        <% end %>
      </nav>
    </header>

    <main>
      <% if notice %>
        <div class="notice"><%= notice %></div>
      <% end %>
      <% if alert %>
        <div class="alert"><%= alert %></div>
      <% end %>

      <%= yield %>
    </main>

    <footer>
      <p>&copy; <%= Time.current.year %> Blog App. All rights reserved.</p>
    </footer>
  </body>
</html>

# app/views/shared/_flash_messages.html.erb
<% flash.each do |name, msg| %>
  <div class="alert alert-<%= name == "notice" ? "success" : "danger" %>">
    <%= msg %>
  </div>
<% end %>

# app/views/posts/show.html.erb
<article class="post">
  <h1><%= @post.title %></h1>
  
  <div class="meta">
    By <%= @post.user.name %> | 
    Posted <%= time_ago_in_words(@post.created_at) %> ago |
    <%= pluralize(@post.comments.count, "comment") %>
  </div>

  <div class="content">
    <%= @post.content %>
  </div>

  <div class="actions">
    <% if current_user == @post.user %>
      <%= link_to "Edit", edit_post_path(@post), class: "button" %>
      <%= button_to "Delete", post_path(@post), method: :delete, 
          data: { confirm: "Are you sure?" }, class: "button danger" %>
    <% end %>
  </div>
</article>

<section class="comments">
  <h2>Comments</h2>
  
  <%= render @comments %>
  
  <% if current_user %>
    <h3>Add a comment</h3>
    <%= render "comments/form", comment: @comment, post: @post %>
  <% else %>
    <p>Please <%= link_to "log in", login_path %> to comment.</p>
  <% end %>
</section> 