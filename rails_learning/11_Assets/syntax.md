# Referencing Assets in Views Syntax

This document provides common syntax examples for using Rails' built-in asset helper methods in your views (`.erb` files).

## Linking to Stylesheets (`stylesheet_link_tag`)

Typically placed in the `<head>` of your layout file.

```erb
<%# app/views/layouts/application.html.erb %>

<%= stylesheet_link_tag 'application', media: 'all', 'data-turbo-track': 'reload' %>
<%# This links to app/assets/stylesheets/application.css (and potentially other files it imports/requires) %>

<%# Link to a specific stylesheet %>
<%= stylesheet_link_tag 'photos', media: 'all' %>
<%# This links to app/assets/stylesheets/photos.css %>

<%# Link to a stylesheet in a subdirectory %>
<%= stylesheet_link_tag 'components/buttons', media: 'all' %>
<%# This links to app/assets/stylesheets/components/buttons.css %>
```

## Linking to JavaScript Files (`javascript_include_tag` / `javascript_importmap_tags`)

Placement depends on whether scripts should load in the `<head>` or before the closing `</body>` tag. Modern Rails often uses Importmap, which uses `javascript_importmap_tags`.

```erb
<%# app/views/layouts/application.html.erb %>

<head>
  <%# ... stylesheets ... %>
  <%= javascript_importmap_tags %> <%# For Importmap-based JS %>
</head>
<body>
  <%= yield %>
  <%# Alternatively, for traditional JS includes %>
  <%# <%= javascript_include_tag 'application', 'data-turbo-track': 'reload' %> %>
</body>

<%# Linking to a specific JS file (less common with modern JS setups) %>
<%# <%= javascript_include_tag 'charts' %> %>
```

## Displaying Images (`image_tag`)

References images from the asset load paths (usually `app/assets/images`).

```erb
<%# app/views/articles/show.html.erb %>

<%= image_tag 'logo.png', alt: 'Company Logo' %>
<%# Links to app/assets/images/logo.png %>

<%# Image in a subdirectory %>
<%= image_tag 'icons/edit.png', size: '16x16' %>
<%# Links to app/assets/images/icons/edit.png %>

<%# Image from a gem or vendor directory (less common directly) %>
<%= image_tag 'my_gem/icon.png' %>
```

## Generating Asset Paths/URLs (`asset_path`, `asset_url`)

Useful when you need the URL of an asset for use in CSS (e.g., background images) or JavaScript.

```erb
<%# In a CSS file processed by Rails (e.g., .scss, .css.erb) %>
.hero-section {
  background-image: url(<%= asset_path 'hero-background.jpg' %>);
}

<%# In a JavaScript file processed by Rails (e.g., .js.erb) %>
const imageUrl = "<%= asset_path 'loading.gif' %>";

<%# Get the full URL %>
<p>Download our guide: <a href="<%= asset_url 'guide.pdf' %>">PDF Guide</a></p>
```

Using these asset helpers ensures that your application correctly references assets in both development and production environments. The next topic will cover precompiling assets for production. 