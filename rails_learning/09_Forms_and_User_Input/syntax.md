# Introduction to Forms Syntax (HTML)

This document provides basic HTML syntax examples for creating web forms, as an introduction before using Rails' form helpers.

## Basic `<form>` Tag

```html
<form action="/submit-data" method="post">
  <!-- Form elements go here -->
</form>
```

## Input Elements

### Text Input (`<input type="text">`)

```html
<label for="username">Username:</label>
<input type="text" id="username" name="user[username]">
```

### Password Input (`<input type="password">`)

```html
<label for="password">Password:</label>
<input type="password" id="password" name="user[password]">
```

### Textarea (`<textarea>`)

For multi-line text input.

```html
<label for="comment">Comment:</label>
<textarea id="comment" name="article[comment]" rows="4" cols="50"></textarea>
```

### Checkbox (`<input type="checkbox">`)

For selecting one or more options.

```html
<input type="checkbox" id="subscribe" name="user[subscribe]" value="yes">
<label for="subscribe">Subscribe to newsletter</label>
```

### Radio Buttons (`<input type="radio">`)

For selecting exactly one option from a set. Use the same `name` for related radio buttons.

```html
<input type="radio" id="male" name="user[gender]" value="male">
<label for="male">Male</label><br>
<input type="radio" id="female" name="user[gender]" value="female">
<label for="female">Female</label>
```

### Select Dropdown (`<select>` and `<option>`)

For selecting an option from a dropdown list.

```html
<label for="country">Country:</label>
<select id="country" name="user[country]">
  <option value="usa">United States</option>
  <option value="canada">Canada</option>
  <option value="uk">United Kingdom</option>
</select>
```

### File Input (`<input type="file">`)

For uploading files. Note that the `<form>` tag requires `enctype="multipart/form-data"` for file uploads.

```html
<label for="avatar">Upload Avatar:</label>
<input type="file" id="avatar" name="user[avatar]">

<!-- Form tag must include this attribute for file uploads -->
<form action="/upload" method="post" enctype="multipart/form-data">
  <!-- ... file input ... -->
</form>
```

### Hidden Input (`<input type="hidden">`)

For sending data that is not visible to the user.

```html
<input type="hidden" name="article[status]" value="draft">
```

### Submit Button (`<input type="submit">` or `<button type="submit">`)

Triggers the form submission.

```html
<input type="submit" value="Save Changes">

<button type="submit">Create Article</button>
```

These basic HTML form elements are the building blocks that Rails form helpers simplify. The next topic will explore how to use Rails' Form Helpers to generate these elements more efficiently and with Rails conventions.

# Form Helpers Syntax

This document provides common syntax examples for using Rails' built-in form helpers, primarily focusing on `form_with`.

## `form_with` for Model-Backed Forms

When you have a model object (`@article`, `@user`, etc.), `form_with` can automatically set the form's `action` and `method` based on whether the object is new or existing (for creating or updating).

```erb
<%= form_with(model: @article, local: true) do |form|
  <%% if @article.errors.any? %>
    <div id="error_explanation">
      <h2><%= pluralize(@article.errors.count, "error") %> prohibited this article from being saved:</h2>

      <ul>
        <%% @article.errors.full_messages.each do |msg| %>
          <li><%= msg %></li>
        <%% end %>
      </ul>
    </div>
  <%% end %>

  <div class="field">
    <%= form.label :title %>
    <%= form.text_field :title %>
  </div>

  <div class="field">
    <%= form.label :body %>
    <%= form.text_area :body %>
  </div>

  <div class="actions">
    <%= form.submit %>
  </div>
<%% end %>
```
*   If `@article` is a new record (`@article.new_record?` is true), the form will submit to `/articles` using `POST`.
*   If `@article` is an existing record, the form will submit to `/articles/:id` using `PATCH`.

## `form_with` for URL-Based Forms

Use when the form is not directly tied to a single model object, or you need to specify the URL and method explicitly.

```erb
<%= form_with(url: '/search', method: :get, local: true) do |form|
  <div>
    <%= form.label :query, "Search:" %>
    <%= form.text_field :query %>
  </div>

  <div>
    <%= form.submit "Search" %>
  </div>
<%% end %>
```

## Common Form Builder Methods

Methods called on the `form` object within the `form_with` block.

### `form.label`

Generates a `<label>` tag.

```erb
<%= form.label :title %>
<%= form.label :published, "Is Published?" %> # Custom label text
```

### `form.text_field`, `form.password_field`, `form.number_field`, etc.

Generates various `<input>` tags.

```erb
<%= form.text_field :username %>
<%= form.password_field :password, placeholder: "Enter your password" %>
<%= form.number_field :quantity, in: 1..10 %>
<%= form.email_field :email %>
<%= form.url_field :website %>
<%= form.phone_field :phone_number %>
<%= form.date_field :start_date %>
```

### `form.text_area`

Generates a `<textarea>` tag.

```erb
<%= form.text_area :body, rows: 5, cols: 40 %>
```

### `form.check_box`

Generates a `<input type="checkbox">`.

```erb
<%= form.check_box :terms_of_service %>
<%= form.label :terms_of_service, "I agree to the terms" %>
```

### `form.radio_button`

Generates a `<input type="radio">`.

```erb
<%= form.radio_button :gender, "male" %>
<%= form.label :gender_male, "Male" %>
<%= form.radio_button :gender, "female" %>
<%= form.label :gender_female, "Female" %>
```

### `form.select`

Generates a `<select>` dropdown.

```erb
<%= form.select :category_id, Category.all.map { |c| [c.name, c.id] }, { prompt: "Select a category" } %>

# Select with a default value
<%= form.select :status, ["draft", "published", "archived"], { selected: "draft" } %>

# Multiple select
<%= form.select :tag_ids, Tag.all.map { |t| [t.name, t.id] }, {}, { multiple: true } %>
```

### `form.file_field`

Generates a `<input type="file">`.

```erb
<%= form.file_field :avatar %>

<%# Remember the enctype attribute for the form_with helper when including file fields %>
<%= form_with(model: @user, local: true, html: { enctype: "multipart/form-data" }) do |form|
  <%= form.file_field :avatar %>
  <%= form.submit %>
<% end %>
```

### `form.hidden_field`

Generates a `<input type="hidden">`.

```erb
<%= form.hidden_field :status, value: "pending" %>
```

### `form.submit`

Generates a submit button.

```erb
<%= form.submit %> # Button text defaults to "Create/Update Model Name"
<%= form.submit "Save Article" %> # Custom button text
```

Using these form helpers, you can efficiently create forms in your Rails views that are integrated with your models and adhere to conventions. The next topic will cover how to handle the data submitted from these forms in your controllers.

# Working with Model Data Syntax

This document provides syntax examples for integrating Rails forms directly with Active Record models.

## Form for a New Record

When creating a new record, you pass a new instance of the model to `form_with`.

```erb
<%# app/views/articles/new.html.erb %>

<h1>New Article</h1>

<%= form_with(model: @article, local: true) do |form|
  <div>
    <%= form.label :title %>
    <%= form.text_field :title %>
  </div>

  <div>
    <%= form.label :body %>
    <%= form.text_area :body %>
  </div>

  <div>
    <%= form.submit "Create Article" %>
  </div>
<% end %>
```

In the corresponding controller action (`ArticlesController#new`):

```ruby
# app/controllers/articles_controller.rb

def new
  @article = Article.new
end
```

When this form is rendered, the input fields will be empty. When submitted, it will send a `POST` request to `/articles`.

## Form for an Existing Record (Editing)

When editing an existing record, you pass the existing instance of the model to `form_with`.

```erb
<%# app/views/articles/edit.html.erb %>

<h1>Editing Article</h1>

<%= form_with(model: @article, local: true) do |form|
  <div>
    <%= form.label :title %>
    <%= form.text_field :title %>
  </div>

  <div>
    <%= form.label :body %>
    <%= form.text_area :body %>
  </div>

  <div>
    <%= form.submit "Update Article" %>
  </div>
<% end %>
```

In the corresponding controller action (`ArticlesController#edit`):

```ruby
# app/controllers/articles_controller.rb

def edit
  @article = Article.find(params[:id])
end
```

When this form is rendered, the input fields will be pre-filled with the existing article's data. When submitted, it will send a `PATCH` request to `/articles/:id`.

## Displaying Validation Errors

When a model object fails validations after a form submission, the `errors` object on the model instance will contain the error messages. You can display these errors in your view.

```erb
<%# Often placed at the top of the form or near the relevant fields %>
<%% if @article.errors.any? %>
  <div id="error_explanation">
    <h2><%= pluralize(@article.errors.count, "error") %> prohibited this article from being saved:</h2>

    <ul>
      <%% @article.errors.full_messages.each do |msg| %>
        <li><%= msg %></li>
      <%% end %>
    </ul>
  </div>
<%% end %>

<div class="field <%= 'field_with_errors' if @article.errors[:title].present? %>">
  <%= form.label :title %>
  <%= form.text_field :title %>
  <%% if @article.errors[:title].present? %>
    <span class="error_message"><%= @article.errors[:title].join(', ') %></span>
  <%% end %>
</div>

<%# Rails also adds a 'field_with_errors' CSS class to the parent div by default %>
```

By associating your forms with model objects using `form_with(model: @object, ...)`, Rails simplifies rendering forms with existing data and displaying validation errors. The next topic will cover handling file uploads in forms.

# Form Security Syntax

This document provides syntax examples related to form security in Rails.

## CSRF Protection - Authenticity Token

When you use Rails form helpers (`form_with`, etc.) for non-GET requests, the authenticity token is automatically included as a hidden field. You typically don't need to manually add this when using helpers.

```erb
<%# Example using form_with - token is automatically generated and included %>
<%= form_with(model: @article, local: true) do |form|
  <%# The authenticity token hidden field is automatically rendered here %>
  <div>
    <%= form.label :title %>
    <%= form.text_field :title %>
  </div>
  <div>
    <%= form.submit %>
  </div>
<% end %>
```

If you were building forms manually without helpers (which is not recommended), you would need to include the token yourself using `form_authenticity_token`.

```erb
<%# Manual form - NOT RECOMMENDED for typical Rails development %>
<form action="/articles" method="post">
  <input type="hidden" name="authenticity_token" value="<%= form_authenticity_token %>">
  <!-- Other form fields -->
  <input type="submit" value="Submit">
</form>
```

In your `ApplicationController`, the `protect_from_forgery` line is usually sufficient to enable CSRF protection:

```ruby
# app/controllers/application_controller.rb
class ApplicationController < ActionController::Base
  protect_from_forgery with: :exception
  # ...
end
```

If you need to skip CSRF verification for specific actions (use with extreme caution, and understand the risks!):

```ruby
# app/controllers/api/v1/webhook_controller.rb
class Api::V1::WebhookController < ApplicationController
  protect_from_forgery with: :null_session, only: :create # Example for an API endpoint receiving external data

  def create
    # Process webhook data
  end
end
```

## Strong Parameters Syntax (Review)

As a reminder, Strong Parameters are crucial for preventing mass assignment vulnerabilities. Define a private method in your controller to permit allowed parameters:

```ruby
# app/controllers/articles_controller.rb

# ... controller actions ...

private

def article_params
  params.require(:article).permit(:title, :body, :status, :published_at, :user_id)
  # To permit nested attributes for associations (e.g., article has_one :seo_meta)
  # params.require(:article).permit(:title, :body, seo_meta_attributes: [:id, :meta_title, :meta_description, :_destroy])
end
```

By default, file uploads are handled as part of the permitted parameters. For example, if you have `form.file_field :avatar` and permit `:avatar` in strong parameters, the uploaded file data will be available in `params[:user][:avatar]` in your controller as an instance of `ActionDispatch::Http::UploadedFile`.

Securing your forms involves using Rails' built-in features like CSRF protection and Strong Parameters, along with general web security practices. This concludes the Forms and User Input section. Let me know if you'd like to review anything or move on to the next major topic! 