# MVC Architecture in Rails

## Model Example
```ruby
# app/models/post.rb
class Post < ApplicationRecord
  validates :title, presence: true
  validates :content, length: { minimum: 10 }
  
  belongs_to :user
  has_many :comments
  
  scope :published, -> { where(published: true) }
  scope :recent, -> { order(created_at: :desc) }
end
```

## Controller Example
```ruby
# app/controllers/posts_controller.rb
class PostsController < ApplicationController
  def index
    @posts = Post.published.recent
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

  private

  def post_params
    params.require(:post).permit(:title, :content, :published)
  end
end
```

## View Example
```erb
<%# app/views/posts/index.html.erb %>
<h1>Posts</h1>

<% @posts.each do |post| %>
  <article>
    <h2><%= post.title %></h2>
    <p><%= post.content %></p>
    <p>By: <%= post.user.name %></p>
    <%= link_to 'Read more', post_path(post) %>
  </article>
<% end %>

<%= link_to 'New Post', new_post_path %>
```

## Practice Exercise
Create a simple blog application with the following features:
1. Create a Post model with title and content
2. Add a User model that has many posts
3. Create views to list all posts and show individual posts
4. Add the ability to create new posts
5. Add basic styling to make it look good

## Solution Steps
1. Generate the models:
```bash
rails generate model User name:string email:string
rails generate model Post title:string content:text user:references
```

2. Set up the associations in models:
```ruby
# app/models/user.rb
class User < ApplicationRecord
  has_many :posts
end

# app/models/post.rb
class Post < ApplicationRecord
  belongs_to :user
  validates :title, presence: true
  validates :content, presence: true
end
```

3. Create the controller:
```bash
rails generate controller Posts index show new create
```

4. Set up the routes:
```ruby
# config/routes.rb
Rails.application.routes.draw do
  resources :posts
  root 'posts#index'
end
```

5. Create the views as shown in the examples above.

6. Add some basic styling:
```css
/* app/assets/stylesheets/posts.css */
article {
  margin: 20px 0;
  padding: 20px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

h1 {
  color: #333;
}

h2 {
  color: #666;
}
```

## Testing Your Understanding
1. What happens when you create a new post?
2. How does the MVC pattern help organize this code?
3. What would you add to make this a more complete blog application? 