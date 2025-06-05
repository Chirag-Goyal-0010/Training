# Rails Controllers - Syntax Reference

## Basic Controller Syntax

```ruby
# Controller definition
class PostsController < ApplicationController
  # Actions
  def index
  end
  
  def show
  end
  
  def new
  end
  
  def create
  end
  
  def edit
  end
  
  def update
  end
  
  def destroy
  end
  
  # Private methods
  private
  
  def post_params
  end
end
```

## Controller Filter Syntax

```ruby
# Before filters
before_action :method_name
before_action :method_name, only: [:action1, :action2]
before_action :method_name, except: [:action1, :action2]
before_action :method_name, if: :condition?
before_action :method_name, unless: :condition?

# After filters
after_action :method_name
after_action :method_name, only: [:action1, :action2]
after_action :method_name, except: [:action1, :action2]

# Around filters
around_action :method_name
around_action :method_name, only: [:action1, :action2]

# Skip filters
skip_before_action :method_name
skip_before_action :method_name, only: [:action1, :action2]
skip_after_action :method_name
skip_after_action :method_name, only: [:action1, :action2]
```

## Controller Parameter Syntax

```ruby
# Strong parameters
def post_params
  params.require(:post).permit(:title, :content)
end

# Nested parameters
def post_params
  params.require(:post).permit(
    :title, :content,
    comments_attributes: [:content, :user_id]
  )
end

# Array parameters
def post_params
  params.require(:post).permit(
    :title, :content,
    tag_ids: []
  )
end

# Hash parameters
def post_params
  params.require(:post).permit(
    :title, :content,
    metadata: [:category, :status]
  )
end
```

## Controller Response Syntax

```ruby
# Render responses
render :action
render :action, status: :ok
render :action, layout: false
render :action, layout: 'admin'

# Redirect responses
redirect_to @post
redirect_to @post, notice: 'Success'
redirect_to @post, alert: 'Error'
redirect_to :back
redirect_to root_path

# JSON responses
render json: @post
render json: @post, status: :ok
render json: { error: 'Not found' }, status: :not_found

# XML responses
render xml: @post
render xml: @post, status: :ok

# File responses
send_file path
send_file path, filename: 'name'
send_file path, type: 'application/pdf'
send_file path, disposition: 'attachment'

# Stream responses
response.headers['Content-Type'] = 'text/event-stream'
self.response_body = Enumerator.new do |yielder|
  yielder << "data: #{content}\n\n"
end
```

## Controller Testing Syntax

```ruby
# Functional tests
test "should get index" do
  get posts_url
  assert_response :success
end

test "should create post" do
  assert_difference('Post.count') do
    post posts_url, params: { post: { title: 'Test' } }
  end
  assert_redirected_to post_url(Post.last)
end

# Integration tests
test "should create post and redirect" do
  post posts_url, params: { post: { title: 'Test' } }
  assert_redirected_to post_url(Post.last)
  follow_redirect!
  assert_response :success
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
```

## Controller Security Syntax

```ruby
# Authentication
before_action :authenticate_user!
before_action :authenticate_user!, except: [:index, :show]

# Authorization
before_action :authorize_post, only: [:edit, :update, :destroy]

# CSRF protection
protect_from_forgery with: :exception
protect_from_forgery with: :null_session

# Parameter filtering
before_action :filter_params

# Session handling
session[:key] = value
session.delete(:key)
reset_session
```

## Controller Performance Syntax

```ruby
# Action caching
caches_action :index, :show
caches_action :index, expires_in: 1.hour
caches_action :show, cache_path: -> { post_path(@post) }

# Fragment caching
fresh_when(@post)
fresh_when(@post, etag: @post, last_modified: @post.updated_at)

# Russian doll caching
fresh_when(@post, etag: [@post, @post.comments])

# ETag support
if stale?(@post)
  respond_to do |format|
    format.html
    format.json { render json: @post }
  end
end

# Conditional GET
if stale?(last_modified: @post.updated_at)
  respond_to do |format|
    format.html
    format.json { render json: @post }
  end
end
```

## Controller Helper Syntax

```ruby
# Flash messages
flash[:notice] = 'Success'
flash[:alert] = 'Error'
flash.now[:notice] = 'Success'

# Headers
response.headers['Content-Type'] = 'application/json'
response.headers['Last-Modified'] = Time.current.httpdate

# Cookies
cookies[:key] = value
cookies[:key] = { value: value, expires: 1.hour }
cookies.delete(:key)

# Session
session[:key] = value
session.delete(:key)
reset_session
```

## Controller Concern Syntax

```ruby
# Include concern
include Searchable
include Authenticatable

# Concern with options
include Searchable, only: [:index, :search]
include Authenticatable, except: [:index, :show]

# Concern with configuration
include Searchable do |config|
  config.search_fields = [:title, :content]
end
```

# Controller Syntax

This document provides a quick reference for common syntax used in Rails controllers.

## Defining a Controller

Controllers are Ruby classes that inherit from `ApplicationController`.

```ruby
class ArticlesController < ApplicationController
  # Actions go here
end
```

## Defining Actions

Actions are public methods within a controller.

```ruby
class ArticlesController < ApplicationController
  def index
    # Code to handle the index action (e.g., list all articles)
  end

  def show
    # Code to handle the show action (e.g., display a specific article)
  end

  def new
    # Code to handle the new action (e.g., prepare a new article for creation)
  end

  def create
    # Code to handle the create action (e.g., save a new article to the database)
  end

  def edit
    # Code to handle the edit action (e.g., prepare an existing article for update)
  end

  def update
    # Code to handle the update action (e.g., save updates to an existing article)
  end

  def destroy
    # Code to handle the destroy action (e.g., delete an article)
  end
end
```

## Accessing Parameters

Parameters sent with requests are available in the `params` hash.

```ruby
# Accessing a specific parameter
article_id = params[:id]

# Permitting strong parameters for mass assignment
def article_params
  params.require(:article).permit(:title, :body, :author_id)
end
```

## Rendering Views

By default, Rails renders a view template that matches the action name (e.g., `index.html.erb` for the `index` action). You can explicitly render a different view or content.

```ruby
# Render a specific view template
render "articles/show"

# Render a different template in the same controller's view directory
render :edit

# Render text
render plain: "Hello, World!"

# Render JSON
render json: @article

# Render nothing with a status code
head :no_content
```

## Redirecting

You can redirect the user to a different URL or action.

```ruby
# Redirect to a specific URL
redirect_to "http://www.example.com"

# Redirect to a specific action within the same controller
redirect_to action: :index

# Redirect to a specific controller and action
redirect_to controller: "articles", action: "show", id: @article.id

# Redirect to a model instance (Rails infers the show path)
redirect_to @article
```

## Using Before/After Callbacks

Callbacks allow you to run methods before, after, or around controller actions.

```ruby
class ArticlesController < ApplicationController
  before_action :set_article, only: [:show, :edit, :update, :destroy]
  after_action :log_article_access, only: [:show]

  # ... actions ...

  private

  def set_article
    @article = Article.find(params[:id])
  end

  def log_article_access
    Rails.logger.info "Article #{@article.id} was accessed."
  end
end
```

This covers some of the fundamental syntax for Rails controllers. Let me know when you're ready to move on to the next topic! 