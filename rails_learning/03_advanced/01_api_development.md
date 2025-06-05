# Rails API Development

## Building a RESTful API

### API Configuration
```ruby
# config/application.rb
module YourApp
  class Application < Rails::Application
    # ... other config ...
    
    # API configuration
    config.api_only = true
    config.middleware.use ActionDispatch::Flash
    config.middleware.use ActionDispatch::Cookies
    config.middleware.use ActionDispatch::Session::CookieStore
  end
end
```

### API Controller Example
```ruby
# app/controllers/api/v1/posts_controller.rb
module Api
  module V1
    class PostsController < ApplicationController
      before_action :set_post, only: [:show, :update, :destroy]
      
      def index
        @posts = Post.published.recent
        render json: @posts, status: :ok
      end
      
      def show
        render json: @post, status: :ok
      end
      
      def create
        @post = Post.new(post_params)
        if @post.save
          render json: @post, status: :created
        else
          render json: { errors: @post.errors }, status: :unprocessable_entity
        end
      end
      
      def update
        if @post.update(post_params)
          render json: @post, status: :ok
        else
          render json: { errors: @post.errors }, status: :unprocessable_entity
        end
      end
      
      def destroy
        @post.destroy
        head :no_content
      end
      
      private
      
      def set_post
        @post = Post.find(params[:id])
      end
      
      def post_params
        params.require(:post).permit(:title, :content, :published)
      end
    end
  end
end
```

### API Serializer
```ruby
# app/serializers/post_serializer.rb
class PostSerializer < ActiveModel::Serializer
  attributes :id, :title, :content, :published, :created_at
  belongs_to :user
  has_many :comments
  
  def comments
    object.comments.recent
  end
end
```

### API Routes
```ruby
# config/routes.rb
Rails.application.routes.draw do
  namespace :api do
    namespace :v1 do
      resources :posts do
        resources :comments, only: [:index, :create]
        resources :likes, only: [:create, :destroy]
      end
      resources :users, only: [:show, :create]
    end
  end
end
```

## Practice Exercise: Blog API
Create a RESTful API for the blog platform with the following features:
1. CRUD operations for posts
2. Authentication using JWT
3. Rate limiting
4. API versioning
5. Documentation using Swagger/OpenAPI

## Solution Steps

1. Add necessary gems to Gemfile:
```ruby
# Gemfile
gem 'jwt'
gem 'rack-cors'
gem 'rack-attack'
gem 'rswag-api'
gem 'rswag-specs'
gem 'rswag-ui'
```

2. Configure CORS:
```ruby
# config/initializers/cors.rb
Rails.application.config.middleware.insert_before 0, Rack::Cors do
  allow do
    origins '*'
    resource '*',
      headers: :any,
      methods: [:get, :post, :put, :patch, :delete, :options, :head]
  end
end
```

3. Create JWT authentication:
```ruby
# app/controllers/api/v1/authentication_controller.rb
module Api
  module V1
    class AuthenticationController < ApplicationController
      def authenticate
        user = User.find_by(email: params[:email])
        if user&.authenticate(params[:password])
          token = jwt_encode(user_id: user.id)
          render json: { token: token }, status: :ok
        else
          render json: { error: 'Invalid credentials' }, status: :unauthorized
        end
      end
    end
  end
end

# app/controllers/application_controller.rb
class ApplicationController < ActionController::API
  include ActionController::HttpAuthentication::Token::ControllerMethods
  
  before_action :authenticate_request
  
  private
  
  def authenticate_request
    header = request.headers['Authorization']
    token = header.split(' ').last if header
    begin
      @decoded = jwt_decode(token)
      @current_user = User.find(@decoded[:user_id])
    rescue ActiveRecord::RecordNotFound, JWT::DecodeError
      render json: { error: 'Unauthorized' }, status: :unauthorized
    end
  end
end
```

4. Implement rate limiting:
```ruby
# config/initializers/rack_attack.rb
class Rack::Attack
  throttle('req/ip', limit: 300, period: 5.minutes) do |req|
    req.ip
  end
  
  throttle('logins/ip', limit: 5, period: 20.seconds) do |req|
    if req.path == '/api/v1/authenticate' && req.post?
      req.ip
    end
  end
end
```

5. Create API documentation:
```ruby
# app/controllers/api/v1/posts_controller.rb
module Api
  module V1
    class PostsController < ApplicationController
      swagger_controller :posts, "Post Management"
      
      swagger_api :index do
        summary "Fetches all posts"
        response :ok, "Success"
        response :unauthorized, "Unauthorized"
      end
      
      swagger_api :show do
        summary "Fetches a single post"
        param :path, :id, :integer, :required, "Post ID"
        response :ok, "Success"
        response :not_found, "Not Found"
      end
      
      # ... other actions ...
    end
  end
end
```

6. Test the API:
```ruby
# spec/requests/api/v1/posts_spec.rb
require 'rails_helper'

RSpec.describe "Api::V1::Posts", type: :request do
  let(:user) { create(:user) }
  let(:token) { jwt_encode(user_id: user.id) }
  let(:headers) { { 'Authorization' => "Bearer #{token}" } }
  
  describe "GET /api/v1/posts" do
    it "returns all posts" do
      get '/api/v1/posts', headers: headers
      expect(response).to have_http_status(:ok)
    end
  end
  
  describe "POST /api/v1/posts" do
    it "creates a new post" do
      post_params = { post: { title: "Test", content: "Content" } }
      post '/api/v1/posts', params: post_params, headers: headers
      expect(response).to have_http_status(:created)
    end
  end
end
```

## Testing Your Understanding
1. How does JWT authentication work in this API?
2. What are the benefits of API versioning?
3. How does rate limiting protect your API?
4. What additional security measures would you implement? 