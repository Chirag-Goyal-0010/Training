# Active Record Associations and Validations

## Model Relationships Example
```ruby
# app/models/user.rb
class User < ApplicationRecord
  has_many :posts
  has_many :comments
  has_many :likes
  has_many :liked_posts, through: :likes, source: :post
  
  validates :email, presence: true, uniqueness: true, format: { with: URI::MailTo::EMAIL_REGEXP }
  validates :username, presence: true, uniqueness: true, length: { minimum: 3 }
end

# app/models/post.rb
class Post < ApplicationRecord
  belongs_to :user
  has_many :comments, dependent: :destroy
  has_many :likes, dependent: :destroy
  has_many :liking_users, through: :likes, source: :user
  
  validates :title, presence: true, length: { minimum: 5 }
  validates :content, presence: true, length: { minimum: 20 }
  
  scope :published, -> { where(published: true) }
  scope :recent, -> { order(created_at: :desc) }
  scope :popular, -> { joins(:likes).group('posts.id').order('COUNT(likes.id) DESC') }
end

# app/models/comment.rb
class Comment < ApplicationRecord
  belongs_to :user
  belongs_to :post
  
  validates :content, presence: true, length: { minimum: 5 }
  
  scope :recent, -> { order(created_at: :desc) }
end

# app/models/like.rb
class Like < ApplicationRecord
  belongs_to :user
  belongs_to :post
  
  validates :user_id, uniqueness: { scope: :post_id }
end
```

## Practice Exercise: Social Blog Platform
Create a social blog platform with the following features:
1. Users can create posts
2. Users can comment on posts
3. Users can like posts
4. Show popular posts based on likes
5. Show recent comments on posts

## Solution Steps

1. Generate the models:
```bash
rails generate model User username:string email:string
rails generate model Post title:string content:text published:boolean user:references
rails generate model Comment content:text user:references post:references
rails generate model Like user:references post:references
```

2. Set up the database:
```bash
rails db:migrate
```

3. Create the controllers:
```ruby
# app/controllers/posts_controller.rb
class PostsController < ApplicationController
  def index
    @posts = Post.published.recent
    @popular_posts = Post.published.popular.limit(5)
  end

  def show
    @post = Post.find(params[:id])
    @comments = @post.comments.recent
    @comment = Comment.new
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

  def post_params
    params.require(:post).permit(:title, :content, :published)
  end
end

# app/controllers/comments_controller.rb
class CommentsController < ApplicationController
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

# app/controllers/likes_controller.rb
class LikesController < ApplicationController
  def create
    @post = Post.find(params[:post_id])
    @like = @post.likes.build(user: current_user)
    
    if @like.save
      redirect_to @post, notice: 'Post liked!'
    else
      redirect_to @post, alert: 'Error liking post.'
    end
  end

  def destroy
    @like = current_user.likes.find(params[:id])
    @like.destroy
    redirect_to @like.post, notice: 'Post unliked.'
  end
end
```

4. Create the views:
```erb
<%# app/views/posts/index.html.erb %>
<h1>Recent Posts</h1>
<div class="posts">
  <% @posts.each do |post| %>
    <article>
      <h2><%= post.title %></h2>
      <p><%= post.content %></p>
      <p>By: <%= post.user.username %></p>
      <p>Likes: <%= post.likes.count %></p>
      <%= link_to 'Read more', post_path(post) %>
    </article>
  <% end %>
</div>

<h2>Popular Posts</h2>
<div class="popular-posts">
  <% @popular_posts.each do |post| %>
    <article>
      <h3><%= post.title %></h3>
      <p>Likes: <%= post.likes.count %></p>
      <%= link_to 'Read more', post_path(post) %>
    </article>
  <% end %>
</div>

<%# app/views/posts/show.html.erb %>
<article>
  <h1><%= @post.title %></h1>
  <p><%= @post.content %></p>
  <p>By: <%= @post.user.username %></p>
  
  <div class="likes">
    <%= @post.likes.count %> likes
    <% if current_user && !current_user.liked_posts.include?(@post) %>
      <%= button_to 'Like', post_likes_path(@post), method: :post %>
    <% else %>
      <%= button_to 'Unlike', like_path(@post.likes.find_by(user: current_user)), method: :delete %>
    <% end %>
  </div>
  
  <div class="comments">
    <h2>Comments</h2>
    <%= render @comments %>
    
    <%= form_with(model: [@post, @comment]) do |f| %>
      <%= f.text_area :content %>
      <%= f.submit 'Add Comment' %>
    <% end %>
  </div>
</article>
```

5. Add some styling:
```css
/* app/assets/stylesheets/posts.css */
.posts, .popular-posts {
  display: grid;
  gap: 20px;
  margin: 20px 0;
}

article {
  padding: 20px;
  border: 1px solid #ddd;
  border-radius: 5px;
  background: #fff;
}

.likes {
  margin: 10px 0;
}

.comments {
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid #eee;
}

.comment {
  padding: 10px;
  margin: 10px 0;
  background: #f9f9f9;
  border-radius: 3px;
}
```

## Testing Your Understanding
1. How do the model associations help in building the social features?
2. What are the benefits of using scopes in the Post model?
3. How does the Like model ensure a user can only like a post once?
4. What would you add to make this a more complete social platform? 