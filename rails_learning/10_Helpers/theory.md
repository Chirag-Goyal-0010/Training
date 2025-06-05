# Introduction to Helpers Theory

In Rails, helper methods are modules that contain methods designed to assist in the view layer of your application. They allow you to encapsulate presentation logic that you might otherwise repeat across multiple views. By moving this logic into helper methods, you keep your view templates clean, readable, and focused on displaying data.

## What are Helpers?

Helpers are essentially Ruby modules whose methods are automatically available in your view templates. When you generate a controller, Rails automatically creates a corresponding helper module (e.g., `ArticlesHelper` for `ArticlesController`). You can also create custom helper modules.

## Why Use Helpers?

*   **Keep Views DRY (Don't Repeat Yourself):** If you have snippets of Ruby code or HTML generation logic that are repeated in several views, you can extract them into a helper method and call that method from your views.
*   **Improve Readability:** Complex logic in view templates can make them difficult to understand. Moving this logic to well-named helper methods makes the view code much cleaner.
*   **Separate Concerns:** Helpers help maintain the separation of concerns by keeping presentation logic out of controllers and models.
*   **Testability:** Helper methods are easier to test in isolation compared to testing logic embedded directly in view templates.

## How Helpers are Included

By default:

*   Helper methods defined in `ApplicationHelper` are available in all views throughout your application.
*   Helper methods defined in a controller-specific helper module (e.g., `ArticlesHelper`) are automatically included and available in the views associated with that controller (e.g., `app/views/articles/*`).

## Examples of Logic Suitable for Helpers

*   Formatting data for display (e.g., dates, currency, text truncation).
*   Generating complex HTML tags or structures based on data.
*   Creating links with specific styling or logic.
*   Conditional display of content.

Instead of having complex Ruby logic directly in your `.erb` or `.html.erb` files:

```erb
<%# Example of logic that could be in a helper - NOT RECOMMENDED in view %>
<p>
  Date Published: <%% if @article.published_at.present? %>
    <%= @article.published_at.strftime("%B %d, %Y") %>
  <%% else %>
    Not published yet
  <%% end %>
</p>
```

You would move that logic to a helper method:

```ruby
# app/helpers/articles_helper.rb

module ArticlesHelper
  def formatted_published_date(article)
    if article.published_at.present?
      article.published_at.strftime("%B %d, %Y")
    else
      "Not published yet"
    end
  end
end
```

And call it from your view:

```erb
<%# app/views/articles/show.html.erb %>

<p>
  Date Published: <%= formatted_published_date(@article) %>
</p>
```

This makes the view much cleaner. The next section will explore some of the useful built-in helpers that Rails provides out of the box.

# Built-in Helpers Theory

Rails comes with a large collection of built-in helper methods that are automatically available in your views. These helpers cover a wide range of common presentation tasks, allowing you to avoid writing repetitive or complex logic directly in your templates.

The built-in helpers are organized into different modules based on their functionality. Some of the key categories include:

## Text Helpers

Helpers for formatting and manipulating text, such as truncation, pluralization, and simple formatting.

*   `truncate`: Truncates a string to a specified length.
*   `pluralize`: Returns the plural form of a word based on a count.
*   `simple_format`: Formats text with simple HTML tags for line breaks and paragraphs.
*   `sanitize`: Cleans up potentially malicious HTML from user input.

## Number Helpers

Helpers for formatting numbers, currencies, percentages, and phone numbers.

*   `number_to_currency`: Formats a number as a currency string.
*   `number_to_percentage`: Formats a number as a percentage string.
*   `number_with_delimiter`: Adds delimiters to a number (e.g., commas).
*   `number_to_phone`: Formats a number as a phone number.

## Date and Time Helpers

Helpers for formatting dates and times and calculating time differences.

*   `time_ago_in_words`: Calculates the time difference between a given time and now in words (e.g., "2 minutes ago").
*   `distance_of_time_in_words`: Calculates the time difference between two times in words.
*   `l` (localize): Formats dates and times according to localization settings.

## Form Helpers

As discussed in the previous section, helpers for generating forms and form elements (`form_with`, `text_field`, `select`, etc.).

## URL Helpers

Helpers for generating URLs and links based on your application's routes. These are crucial for maintaining correct links throughout your application.

*   `link_to`: Creates a link (`<a>` tag) to a specified URL or route.
*   `url_for`: Generates a URL for a given set of options (e.g., controller, action, model object).
*   `_path` and `_url` helpers (e.g., `articles_path`, `edit_article_url`): Automatically generated helpers based on your routes definitions (`config/routes.rb`). Using these is highly recommended.

## Tag Helpers

Helpers for generating common HTML tags.

*   `tag`: A flexible helper for generating arbitrary HTML tags.
*   `content_tag`: Generates an HTML tag with content inside.

These are just a few examples, and Rails provides many more built-in helpers. Utilizing them effectively can greatly simplify your view code. The next section will provide syntax examples for using some of these common built-in helpers.

# Custom Helpers Theory

While Rails provides a rich set of built-in helper methods, you will frequently encounter scenarios where you need to create your own custom helpers to encapsulate presentation logic specific to your application. Custom helpers help keep your views DRY (Don't Repeat Yourself) and improve their readability.

## Where to Define Custom Helpers

Custom helper methods are defined within Ruby modules in the `app/helpers` directory. By default:

*   Methods in `app/helpers/application_helper.rb` are available in all views.
*   Methods in controller-specific helper modules (e.g., `app/helpers/articles_helper.rb`) are automatically included and available only in the views for that controller.

You can also include helper modules in other modules or controllers manually if needed, but sticking to the default conventions is usually sufficient.

## How to Create a Custom Helper Method

A custom helper method is simply a public Ruby method defined within a helper module. These methods can take arguments and can output strings (which will be escaped by default in ERB) or raw HTML (using `html_safe` or helpers like `tag`).

```ruby
# app/helpers/application_helper.rb

module ApplicationHelper
  # A simple helper method
  def formatted_price(price)
    number_to_currency(price, unit: "€") # Using a built-in helper within a custom helper
  end

  # A helper method that outputs HTML
  def user_avatar(user, size: 50)
    # Assuming user has an avatar attachment (e.g., using Active Storage)
    if user.avatar.attached?
      image_tag user.avatar.variant(resize_to_limit: [size, size]), alt: user.name
    else
      image_tag 'default_avatar.png', size: "#{size}x#{size}", alt: user.name
    end
  end
end
```

## Using Custom Helpers in Views

Once defined in a helper module that is included in the view, you can call custom helper methods just like built-in ones.

```erb
<%# app/views/products/show.html.erb %>

<p>Price: <%= formatted_price(@product.price) %></p>

<div>
  <%= user_avatar(@product.user, size: 80) %>
  <span><%= @product.user.name %></span>
</div>
```

## Best Practices for Custom Helpers

*   **Keep them focused:** Helper methods should ideally do one thing well.
*   **Use descriptive names:** Method names should clearly indicate their purpose.
*   **Avoid complex logic:** If the logic becomes too complicated, consider moving it to a presenter or a dedicated service object instead of a helper.
*   **Return HTML when generating HTML:** If a helper is designed to output HTML, use `html_safe` or helpers like `tag` or `content_tag` to ensure it's rendered correctly and not escaped.

Custom helpers are an essential tool for managing presentation logic in your Rails application. They contribute to cleaner views and better code organization. The next section will provide syntax examples for creating and using custom helpers.

# Helper Modules and Organization Theory

In Rails, helper methods are organized into modules within the `app/helpers` directory. This modular structure helps manage the scope and organization of your helper code, preventing naming conflicts and making it clearer where to find specific helper methods.

## Default Helper Organization

By default, Rails follows a convention for organizing helpers:

*   **`ApplicationHelper` (`app/helpers/application_helper.rb`):** This module is included in all views throughout your application. It's the place to put helper methods that are needed globally.
*   **Controller-Specific Helpers (`app/helpers/[controller_name]_helper.rb`):** When you generate a controller (e.g., `rails generate controller Articles`), Rails automatically creates a corresponding helper module (e.g., `ArticlesHelper`). Methods in this module are automatically included and available only in the views associated with that controller (e.g., `app/views/articles/*`).

```
app/
├── helpers/
│   ├── application_helper.rb  # Global helpers
│   ├── articles_helper.rb     # Helpers for ArticlesController views
│   └── users_helper.rb        # Helpers for UsersController views
└── views/
    ├── articles/
    │   ├── index.html.erb
    │   └── show.html.erb
    ├── layouts/
    │   └── application.html.erb
    └── users/
        ├── index.html.erb
        └── show.html.erb
```

In the example above:
*   Methods in `application_helper.rb` are available in `application.html.erb`, `index.html.erb` (articles), `show.html.erb` (articles), `index.html.erb` (users), and `show.html.erb` (users).
*   Methods in `articles_helper.rb` are available only in `index.html.erb` (articles) and `show.html.erb` (articles).
*   Methods in `users_helper.rb` are available only in `index.html.erb` (users) and `show.html.erb` (users).

## Including Helpers Manually

While the default organization covers most cases, you can explicitly include helper modules in controllers if needed. This is less common for view helpers but can be useful if you want to make helper methods available outside of views (e.g., in controllers or mailers).

```ruby
# app/controllers/articles_controller.rb

class ArticlesController < ApplicationController
  include ProductsHelper # Include helpers from ProductsHelper

  def show
    @article = Article.find(params[:id])
    # You can now call methods from ProductsHelper here
  end
end
```

**Note:** Including view helpers in controllers can blur the lines between the view and controller layers and should be done judiciously.

## Organizing Complex Helper Logic

For larger applications, you might have a significant number of helper methods. To keep things organized, you can break down large helper modules into smaller, more focused modules and include them as needed. However, stick to the default `app/helpers` structure unless you have a compelling reason to deviate.

Effective organization of helper methods is key to maintaining a clean, understandable, and scalable Rails application. It ensures that helpers are available where needed without cluttering the global namespace. The next section will provide syntax examples related to helper organization and inclusion. 