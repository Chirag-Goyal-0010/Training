# Introduction to Helpers Syntax

This document provides basic syntax examples for using helper methods in Rails views.

## Calling Helper Methods

Helper methods defined in your helper modules are automatically available in your view templates. You can call them just like any other method.

```erb
<%# app/views/articles/show.html.erb %>

<h1><%= @article.title %></h1>

<p>
  Published: <%= formatted_published_date(@article) %> <%# Calling a custom helper method %>
</p>

<p>
  Summary: <%= truncate(@article.body, length: 100) %> <%# Calling a built-in helper method %>
</p>

<%= link_to 'Edit Article', edit_article_path(@article) %> <%# Calling a built-in link helper %>

```

## Helper Methods That Output HTML

Some helper methods generate HTML tags. By default, ERB (`<%= %>`) automatically escapes the output of helper methods to prevent XSS vulnerabilities. If a helper intentionally outputs raw HTML (like `raw` or `html_safe`), it won't be escaped.

```ruby
# app/helpers/application_helper.rb
module ApplicationHelper
  def danger_button(text)
    # This helper generates raw HTML
    tag.button(text, class: 'btn btn-danger')
  end
end
```

```erb
<%# app/views/articles/show.html.erb %>

<%= danger_button('Delete') %> <%# This will render a button HTML tag %>
```

## Passing Arguments to Helpers

Helper methods can accept arguments, allowing you to make them more dynamic and reusable.

```ruby
# app/helpers/application_helper.rb
module ApplicationHelper
  def status_badge(status)
    case status
    when 'published'
      tag.span('Published', class: 'badge bg-success')
    when 'draft'
      tag.span('Draft', class: 'badge bg-secondary')
    else
      tag.span(status.humanize, class: 'badge bg-light')
    end
  end
end
```

```erb
<%# app/views/articles/index.html.erb %>

<%% @articles.each do |article| %>
  <div>
    <h2><%= article.title %></h2>
    <p>Status: <%= status_badge(article.status) %></p> <%# Passing the article's status to the helper %>
  </div>
<%% end %>
```

Understanding how to call and utilize helper methods is key to writing cleaner and more maintainable Rails views. The next topic will dive into some of the useful built-in helpers provided by Rails. 

# Built-in Helpers Syntax

This document provides common syntax examples for using various built-in Rails helper methods in your views.

## Text Helpers

```erb
<%# Truncate text %>
<p><%= truncate("This is a long piece of text that needs to be shortened.", length: 30) %></p>
<%# Output: <p>This is a long piece of te...</p> %>

<%# Pluralize words %>
<p><%= pluralize(@article.comments.count, "comment") %></p>
<%# If @article.comments.count is 1, output: <p>1 comment</p> %>
<%# If @article.comments.count is 5, output: <p>5 comments</p> %>

<%# Simple format for text areas %>
<div><%= simple_format(@article.body) %></div>
<%# Converts line breaks to <br> and double line breaks to <p> %>

<%# Sanitize potentially unsafe HTML %>
<div><%= sanitize(@user_input_html) %></div>
<%# Removes potentially dangerous tags and attributes %>
```

## Number Helpers

```erb
<%# Format as currency %>
<p>Price: <%= number_to_currency(1234.56) %></p>
<%# Output (default): <p>Price: $1,234.56</p> %>
<p>Price (Euro): <%= number_to_currency(1234.56, unit: "€", separator: ",", delimiter: ".") %></p>

<%# Format as percentage %>
<p>Completion: <%= number_to_percentage(75) %></p>
<%# Output: <p>Completion: 75.000%</p> %>

<%# Add delimiters %>
<p>Population: <%= number_with_delimiter(1234567) %></p>
<%# Output: <p>Population: 1,234,567</p> %>

<%# Format as phone number %>
<p>Contact: <%= number_to_phone(1235551212) %></p>
<%# Output: <p>Contact: 123-555-1212</p> %>
```

## Date and Time Helpers

```erb
<%# Time ago in words %>
<p>Posted: <%= time_ago_in_words(@article.created_at) %> ago</p>
<%# Output: <p>Posted: about 5 hours ago</p> %>

<%# Distance of time in words %>
<p>Duration: <%= distance_of_time_in_words(@project.start_date, @project.end_date) %></p>
<%# Output: <p>Duration: about 3 months</p> %>

<%# Localized date/time formatting %>
<p>Published on: <%= l(@article.published_at, format: :long) %></p>
<%# Requires localization setup in config/locales %>
```

## URL Helpers (`_path` and `_url`)

These are automatically generated based on your `config/routes.rb`. Use them to avoid hardcoding URLs.

```erb
<%# Link to the index page for articles %>
<%= link_to 'Back to Articles', articles_path %>

<%# Link to a specific article's show page %>
<%= link_to @article.title, article_path(@article) %>

<%# Link to the edit page for an article %>
<%= link_to 'Edit', edit_article_path(@article) %>

<%# Link to destroy an article (requires remote: :true or button_to) %>
<%= link_to 'Destroy', article_path(@article), data: { turbo_method: :delete, turbo_confirm: 'Are you sure?' } %>

<%# Generate a full URL %>
<p>Share this link: <%= article_url(@article) %></p>
```

## Tag Helpers

```erb
<%# Using the tag helper %>
<%= tag.div class: 'container', id: 'main-content' do %>
  <p>Content goes here.</p>
<% end %>
<%# Output: <div class="container" id="main-content"><p>Content goes here.</p></div> %>

<%# Using content_tag (older helper, tag is preferred) %>
<%= content_tag(:span, 'New!', class: 'badge bg-info') %>
<%# Output: <span class="badge bg-info">New!</span> %>
```

These examples demonstrate the power and convenience of Rails' built-in helpers. By using them, you can keep your views clean and consistent. The next topic will cover how to create your own custom helper methods. 

# Helper Modules and Organization Syntax

This document provides syntax examples related to organizing and including helper modules in Rails.

## Default Helper Modules

Rails automatically creates and includes helper modules based on controller names. You define methods within these modules.

```ruby
# app/helpers/articles_helper.rb
module ArticlesHelper
  def article_status_label(article)
    # Helper method specific to articles views
    case article.status
    when 'published' then 'Published'
    when 'draft' then 'Draft'
    else article.status.humanize
    end
  end
end
```

```ruby
# app/helpers/users_helper.rb
module UsersHelper
  def user_full_name(user)
    # Helper method specific to users views
    "#{user.first_name} #{user.last_name}"
  end
end
```

Methods in `ArticlesHelper` are available in `app/views/articles/*`, and methods in `UsersHelper` are available in `app/views/users/*`.

Methods in `ApplicationHelper` are available everywhere.

```ruby
# app/helpers/application_helper.rb
module ApplicationHelper
  def formatted_date(date)
    date.strftime("%B %d, %Y") if date.present?
  end
end
```

```erb
<%# app/views/articles/show.html.erb %>

<p>Created by: <%= user_full_name(@article.user) %></p> <%# Calling UsersHelper method (if available via ApplicationHelper or explicit include) %>
<p>Published Status: <%= article_status_label(@article) %></p> <%# Calling ArticlesHelper method %>
<p>Created on: <%= formatted_date(@article.created_at) %></p> <%# Calling ApplicationHelper method %>
```

**Note:** While `user_full_name` is defined in `UsersHelper`, it won't be automatically available in `articles/show.html.erb` unless `UsersHelper` is explicitly included in `ArticlesController` or `ApplicationHelper`.

## Manually Including Helper Modules (Less Common)

You can explicitly include helper modules in controllers or other helper modules.

```ruby
# app/controllers/articles_controller.rb

class ArticlesController < ApplicationController
  include UsersHelper # Now methods from UsersHelper are available in ArticlesController and its views

  def show
    @article = Article.find(params[:id])
    # Can call user_full_name here
  end
end
```

```ruby
# app/helpers/articles_helper.rb

module ArticlesHelper
  include ProductsHelper # Include helpers from ProductsHelper into ArticlesHelper

  def article_price_details(article)
    # Can now call methods from ProductsHelper here
    "Price: #{formatted_price(article.price)}"
  end
end
```

## Excluding Helper Methods

You can use `helper` with `:except` or `:only` in your controller to selectively include/exclude methods from controller-specific helpers, but this is less common than the default behavior.

```ruby
# app/controllers/articles_controller.rb
class ArticlesController < ApplicationController
  # Only make article_status_label available from ArticlesHelper
  helper ArticlesHelper, only: :article_status_label

  # Make all of ArticlesHelper available except article_status_label
  # helper ArticlesHelper, except: :article_status_label
end
```

Understanding helper organization helps you decide where to place your helper code and how to access it from your views. This concludes the Helper Methods section. Let me know if you'd like to review anything or move on to the next major topic, which is typically Assets (Managing JavaScript, CSS, and Images). 