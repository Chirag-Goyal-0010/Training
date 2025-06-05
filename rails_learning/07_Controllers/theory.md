# Rails Controllers - Code Examples

## Basic Controller Structure

```ruby
# app/controllers/posts_controller.rb
class PostsController < ApplicationController
  # Controller callbacks
  before_action :set_post, only: [:show, :edit, :update, :destroy]
  before_action :authenticate_user!, except: [:index, :show]
  
  # Instance variables
  def index
    @posts = Post.all
  end
  
  def show
    # @post is set by before_action
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
  
  def edit
    # @post is set by before_action
  end
  
  def update
    if @post.update(post_params)
      redirect_to @post, notice: 'Post was successfully updated.'
    else
      render :edit
    end
  end
  
  def destroy
    @post.destroy
    redirect_to posts_url, notice: 'Post was successfully deleted.'
  end
  
  private
  
  def set_post
    @post = Post.find(params[:id])
  end
  
  def post_params
    params.require(:post).permit(:title, :content, :user_id)
  end
end
```

## Controller Actions

```ruby
# app/controllers/posts_controller.rb
class PostsController < ApplicationController
  # Index action with pagination
  def index
    @posts = Post.page(params[:page]).per(10)
  end
  
  # Show action with includes
  def show
    @post = Post.includes(:comments, :user).find(params[:id])
  end
  
  # New action with default values
  def new
    @post = Post.new(user: current_user)
  end
  
  # Create action with transaction
  def create
    @post = Post.new(post_params)
    
    Post.transaction do
      @post.save!
      @post.create_activity(:create, owner: current_user)
    end
    
    redirect_to @post, notice: 'Post was successfully created.'
  rescue ActiveRecord::RecordInvalid
    render :new
  end
  
  # Edit action with authorization
  def edit
    @post = Post.find(params[:id])
    authorize @post
  end
  
  # Update action with versioning
  def update
    @post = Post.find(params[:id])
    authorize @post
    
    if @post.update(post_params)
      @post.create_activity(:update, owner: current_user)
      redirect_to @post, notice: 'Post was successfully updated.'
    else
      render :edit
    end
  end
  
  # Destroy action with soft delete
  def destroy
    @post = Post.find(params[:id])
    authorize @post
    
    @post.update(deleted_at: Time.current)
    redirect_to posts_url, notice: 'Post was successfully deleted.'
  end
  
  # Custom action with search
  def search
    @posts = Post.search(params[:q])
    render :index
  end
end
```

## Controller Filters

```ruby
# app/controllers/posts_controller.rb
class PostsController < ApplicationController
  # Basic filters
  before_action :set_post, only: [:show, :edit, :update, :destroy]
  after_action :log_action, only: [:create, :update, :destroy]
  
  # Filter with options
  before_action :authenticate_user!, except: [:index, :show]
  before_action :set_locale, if: :user_signed_in?
  
  # Filter with conditions
  before_action :check_permissions, only: [:edit, :update, :destroy] do
    @post.user == current_user
  end
  
  # Around filter
  around_action :wrap_in_transaction, only: [:create, :update]
  
  # Skip filter
  skip_before_action :authenticate_user!, only: [:index, :show]
  
  private
  
  def set_post
    @post = Post.find(params[:id])
  end
  
  def log_action
    Rails.logger.info("Action #{action_name} performed on Post #{@post.id}")
  end
  
  def set_locale
    I18n.locale = current_user.locale
  end
  
  def check_permissions
    unless @post.user == current_user
      redirect_to posts_url, alert: 'You are not authorized to perform this action.'
    end
  end
  
  def wrap_in_transaction
    ActiveRecord::Base.transaction do
      yield
    end
  end
end
```

## Controller Parameters

```ruby
# app/controllers/posts_controller.rb
class PostsController < ApplicationController
  # Strong parameters
  def post_params
    params.require(:post).permit(:title, :content, :user_id)
  end
  
  # Nested parameters
  def post_with_comments_params
    params.require(:post).permit(
      :title, :content,
      comments_attributes: [:content, :user_id]
    )
  end
  
  # Array parameters
  def post_with_tags_params
    params.require(:post).permit(
      :title, :content,
      tag_ids: []
    )
  end
  
  # Hash parameters
  def post_with_metadata_params
    params.require(:post).permit(
      :title, :content,
      metadata: [:category, :status, :priority]
    )
  end
  
  # Parameter validation
  def validate_post_params
    if params[:post][:title].blank?
      flash.now[:alert] = "Title can't be blank"
      render :new
      return
    end
    
    if params[:post][:content].length < 10
      flash.now[:alert] = "Content is too short"
      render :new
      return
    end
  end
end
```

## Controller Responses

```ruby
# app/controllers/posts_controller.rb
class PostsController < ApplicationController
  # Render responses
  def index
    @posts = Post.all
    render :index
  end
  
  def show
    @post = Post.find(params[:id])
    render :show, status: :ok
  end
  
  # Redirect responses
  def create
    @post = Post.new(post_params)
    if @post.save
      redirect_to @post, notice: 'Post was successfully created.'
    else
      redirect_to new_post_path, alert: 'Failed to create post.'
    end
  end
  
  # JSON responses
  def api_index
    @posts = Post.all
    render json: @posts, status: :ok
  end
  
  def api_show
    @post = Post.find(params[:id])
    render json: @post, status: :ok
  end
  
  # XML responses
  def xml_index
    @posts = Post.all
    render xml: @posts, status: :ok
  end
  
  # File responses
  def download
    @post = Post.find(params[:id])
    send_file @post.attachment.path,
              filename: @post.attachment_file_name,
              type: @post.attachment_content_type
  end
  
  # Stream responses
  def stream
    @post = Post.find(params[:id])
    response.headers['Content-Type'] = 'text/event-stream'
    response.headers['Last-Modified'] = @post.updated_at.httpdate
    
    self.response_body = Enumerator.new do |yielder|
      yielder << "data: #{@post.content}\n\n"
    end
  end
end
```

## Controller Testing

```ruby
# test/controllers/posts_controller_test.rb
require "test_helper"

class PostsControllerTest < ActionDispatch::IntegrationTest
  # Functional tests
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
      post posts_url, params: { post: { title: 'Test', content: 'Content' } }
    end
    
    assert_redirected_to post_url(Post.last)
  end
  
  # Integration tests
  test "should create post and redirect to show" do
    post posts_url, params: { post: { title: 'Test', content: 'Content' } }
    assert_redirected_to post_url(Post.last)
    
    follow_redirect!
    assert_response :success
    assert_select 'h1', 'Test'
  end
  
  # Request tests
  test "should require authentication" do
    get new_post_url
    assert_redirected_to login_url
  end
  
  # Response tests
  test "should return json" do
    get post_url(@post), as: :json
    assert_response :success
    assert_equal 'application/json', @response.content_type
  end
  
  # Parameter tests
  test "should not create post with invalid params" do
    assert_no_difference('Post.count') do
      post posts_url, params: { post: { title: '' } }
    end
    
    assert_response :unprocessable_entity
  end
end
```

## Controller Security

```ruby
# app/controllers/posts_controller.rb
class PostsController < ApplicationController
  # Authentication
  before_action :authenticate_user!
  
  # Authorization
  before_action :authorize_post, only: [:edit, :update, :destroy]
  
  # CSRF protection
  protect_from_forgery with: :exception
  
  # Parameter filtering
  before_action :filter_params
  
  # Session handling
  before_action :set_session_variables
  
  private
  
  def authorize_post
    @post = Post.find(params[:id])
    unless @post.user == current_user
      redirect_to posts_url, alert: 'You are not authorized to perform this action.'
    end
  end
  
  def filter_params
    params[:post].delete(:user_id) unless current_user.admin?
  end
  
  def set_session_variables
    session[:last_visited_post] = params[:id] if params[:id]
  end
end
```

## Controller Performance

```ruby
# app/controllers/posts_controller.rb
class PostsController < ApplicationController
  # Action caching
  caches_action :index, :show
  
  # Fragment caching
  def index
    @posts = Post.all
    fresh_when(@posts)
  end
  
  # Russian doll caching
  def show
    @post = Post.includes(:comments).find(params[:id])
    fresh_when(@post)
  end
  
  # ETag support
  def index
    @posts = Post.all
    if stale?(@posts)
      respond_to do |format|
        format.html
        format.json { render json: @posts }
      end
    end
  end
  
  # Conditional GET
  def show
    @post = Post.find(params[:id])
    if stale?(last_modified: @post.updated_at)
      respond_to do |format|
        format.html
        format.json { render json: @post }
      end
    end
  end
end
``` 