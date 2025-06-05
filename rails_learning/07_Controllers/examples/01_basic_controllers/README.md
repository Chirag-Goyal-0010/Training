# Basic Controllers Example

This example demonstrates fundamental Rails controller concepts using a blog post system.

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

### 1. Application Controller
```ruby
# app/controllers/application_controller.rb
class ApplicationController < ActionController::Base
  protect_from_forgery with: :exception
  
  before_action :set_locale
  before_action :authenticate_user!
  
  private
  
  def set_locale
    I18n.locale = params[:locale] || I18n.default_locale
  end
  
  def authenticate_user!
    unless current_user
      flash[:alert] = 'Please sign in to continue'
      redirect_to login_path
    end
  end
  
  def current_user
    @current_user ||= User.find_by(id: session[:user_id])
  end
  helper_method :current_user
end
```

### 2. Posts Controller
```ruby
# app/controllers/posts_controller.rb
class PostsController < ApplicationController
  before_action :set_post, only: [:show, :edit, :update, :destroy]
  before_action :authorize_post, only: [:edit, :update, :destroy]
  
  def index
    @posts = Post.includes(:user, :tags)
                 .order(created_at: :desc)
                 .page(params[:page])
  end
  
  def show
    @comments = @post.comments.includes(:user)
                    .order(created_at: :desc)
  end
  
  def new
    @post = Post.new
  end
  
  def create
    @post = current_user.posts.build(post_params)
    
    if @post.save
      redirect_to @post, notice: 'Post was successfully created'
    else
      render :new
    end
  end
  
  def edit
  end
  
  def update
    if @post.update(post_params)
      redirect_to @post, notice: 'Post was successfully updated'
    else
      render :edit
    end
  end
  
  def destroy
    @post.destroy
    redirect_to posts_path, notice: 'Post was successfully deleted'
  end
  
  def search
    @posts = Post.search(params[:q])
                 .includes(:user, :tags)
                 .page(params[:page])
    render :index
  end
  
  private
  
  def set_post
    @post = Post.find(params[:id])
  end
  
  def authorize_post
    unless @post.user == current_user
      flash[:alert] = 'You are not authorized to perform this action'
      redirect_to posts_path
    end
  end
  
  def post_params
    params.require(:post).permit(
      :title, :content,
      tag_ids: []
    )
  end
end
```

### 3. Comments Controller
```ruby
# app/controllers/comments_controller.rb
class CommentsController < ApplicationController
  before_action :set_post
  before_action :set_comment, only: [:edit, :update, :destroy]
  before_action :authorize_comment, only: [:edit, :update, :destroy]
  
  def create
    @comment = @post.comments.build(comment_params)
    @comment.user = current_user
    
    if @comment.save
      redirect_to @post, notice: 'Comment was successfully added'
    else
      redirect_to @post, alert: 'Error adding comment'
    end
  end
  
  def edit
  end
  
  def update
    if @comment.update(comment_params)
      redirect_to @post, notice: 'Comment was successfully updated'
    else
      render :edit
    end
  end
  
  def destroy
    @comment.destroy
    redirect_to @post, notice: 'Comment was successfully deleted'
  end
  
  private
  
  def set_post
    @post = Post.find(params[:post_id])
  end
  
  def set_comment
    @comment = @post.comments.find(params[:id])
  end
  
  def authorize_comment
    unless @comment.user == current_user
      flash[:alert] = 'You are not authorized to perform this action'
      redirect_to @post
    end
  end
  
  def comment_params
    params.require(:comment).permit(:content)
  end
end
```

### 4. Tags Controller
```ruby
# app/controllers/tags_controller.rb
class TagsController < ApplicationController
  def index
    @tags = Tag.includes(:posts)
               .order(name: :asc)
  end
  
  def show
    @tag = Tag.find(params[:id])
    @posts = @tag.posts.includes(:user)
                 .order(created_at: :desc)
                 .page(params[:page])
  end
end
```

### 5. Users Controller
```ruby
# app/controllers/users_controller.rb
class UsersController < ApplicationController
  before_action :set_user, only: [:show, :edit, :update]
  before_action :authorize_user, only: [:edit, :update]
  
  def show
    @posts = @user.posts.includes(:tags)
                 .order(created_at: :desc)
                 .page(params[:page])
  end
  
  def edit
  end
  
  def update
    if @user.update(user_params)
      redirect_to @user, notice: 'Profile was successfully updated'
    else
      render :edit
    end
  end
  
  private
  
  def set_user
    @user = User.find(params[:id])
  end
  
  def authorize_user
    unless @user == current_user
      flash[:alert] = 'You are not authorized to perform this action'
      redirect_to @user
    end
  end
  
  def user_params
    params.require(:user).permit(:name, :email)
  end
end
```

## Testing

### 1. Posts Controller Test
```ruby
# test/controllers/posts_controller_test.rb
require 'test_helper'

class PostsControllerTest < ActionDispatch::IntegrationTest
  setup do
    @user = users(:one)
    @post = posts(:one)
    sign_in @user
  end
  
  test "should get index" do
    get posts_url
    assert_response :success
  end
  
  test "should get new" do
    get new_post_url
    assert_response :success
  end
  
  test "should create post" do
    assert_difference('Post.count') do
      post posts_url, params: {
        post: {
          title: 'Test Post',
          content: 'Test Content',
          tag_ids: [tags(:one).id]
        }
      }
    end
    
    assert_redirected_to post_url(Post.last)
  end
  
  test "should show post" do
    get post_url(@post)
    assert_response :success
  end
  
  test "should get edit" do
    get edit_post_url(@post)
    assert_response :success
  end
  
  test "should update post" do
    patch post_url(@post), params: {
      post: {
        title: 'Updated Title',
        content: 'Updated Content'
      }
    }
    assert_redirected_to post_url(@post)
  end
  
  test "should destroy post" do
    assert_difference('Post.count', -1) do
      delete post_url(@post)
    end
    
    assert_redirected_to posts_url
  end
end
```

### 2. Comments Controller Test
```ruby
# test/controllers/comments_controller_test.rb
require 'test_helper'

class CommentsControllerTest < ActionDispatch::IntegrationTest
  setup do
    @user = users(:one)
    @post = posts(:one)
    @comment = comments(:one)
    sign_in @user
  end
  
  test "should create comment" do
    assert_difference('Comment.count') do
      post post_comments_url(@post), params: {
        comment: { content: 'Test Comment' }
      }
    end
    
    assert_redirected_to post_url(@post)
  end
  
  test "should update comment" do
    patch post_comment_url(@post, @comment), params: {
      comment: { content: 'Updated Comment' }
    }
    assert_redirected_to post_url(@post)
  end
  
  test "should destroy comment" do
    assert_difference('Comment.count', -1) do
      delete post_comment_url(@post, @comment)
    end
    
    assert_redirected_to post_url(@post)
  end
end
```

## Key Features Demonstrated

1. **Basic Controller Structure**
   - Controller inheritance
   - Action methods
   - Private methods
   - Instance variables

2. **Controller Filters**
   - Before actions
   - Authorization checks
   - Resource loading

3. **Controller Parameters**
   - Strong parameters
   - Nested parameters
   - Array parameters

4. **Controller Responses**
   - Redirects
   - Renders
   - Flash messages

5. **Controller Testing**
   - Functional tests
   - Integration tests
   - Authentication tests

## Next Steps

1. Add more controller features:
   - API endpoints
   - File uploads
   - Background jobs
   - Caching

2. Enhance security:
   - Role-based authorization
   - Rate limiting
   - API authentication
   - Request validation

3. Improve performance:
   - Action caching
   - Fragment caching
   - Database optimization
   - Background processing

4. Add more tests:
   - Request specs
   - Response specs
   - Security specs
   - Performance specs 