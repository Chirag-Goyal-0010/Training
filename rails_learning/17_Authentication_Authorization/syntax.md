# Authentication & Authorization Syntax Examples

## Devise Implementation

### Installation and Setup
```ruby
# Gemfile
gem 'devise'

# Terminal
rails generate devise:install
rails generate devise User
rails db:migrate
```

### User Model
```ruby
# app/models/user.rb
class User < ApplicationRecord
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :validatable,
         :confirmable, :lockable, :timeoutable,
         :trackable, :omniauthable
         
  has_many :articles
  has_many :comments
  
  def admin?
    role == 'admin'
  end
end
```

### Devise Configuration
```ruby
# config/initializers/devise.rb
Devise.setup do |config|
  config.mailer_sender = 'please-change-me-at-config-initializers-devise@example.com'
  config.case_insensitive_keys = [:email]
  config.strip_whitespace_keys = [:email]
  config.skip_session_storage = [:http_auth]
  config.stretches = Rails.env.test? ? 1 : 12
  config.reconfirmable = true
  config.expire_all_remember_me_on_sign_out = true
  config.password_length = 6..128
  config.email_regexp = /\A[^@\s]+@[^@\s]+\z/
  config.reset_password_within = 6.hours
  config.sign_out_via = :delete
end
```

## JWT Implementation

### JWT Service
```ruby
# app/services/jwt_service.rb
class JwtService
  def self.encode(payload)
    JWT.encode(payload, Rails.application.credentials.secret_key_base)
  end

  def self.decode(token)
    JWT.decode(token, Rails.application.credentials.secret_key_base)[0]
  rescue JWT::DecodeError
    nil
  end
end
```

### JWT Authentication
```ruby
# app/controllers/api/v1/base_controller.rb
module Api
  module V1
    class BaseController < ApplicationController
      before_action :authenticate_user_from_token!
      
      private
      
      def authenticate_user_from_token!
        token = request.headers['Authorization']&.split(' ')&.last
        return render_unauthorized unless token
        
        payload = JwtService.decode(token)
        return render_unauthorized unless payload
        
        @current_user = User.find_by(id: payload['user_id'])
        render_unauthorized unless @current_user
      end
      
      def render_unauthorized
        render json: { error: 'Unauthorized' }, status: :unauthorized
      end
    end
  end
end
```

## OAuth Implementation

### OmniAuth Setup
```ruby
# Gemfile
gem 'omniauth'
gem 'omniauth-google-oauth2'
gem 'omniauth-facebook'

# config/initializers/omniauth.rb
Rails.application.config.middleware.use OmniAuth::Builder do
  provider :google_oauth2, ENV['GOOGLE_CLIENT_ID'], ENV['GOOGLE_CLIENT_SECRET']
  provider :facebook, ENV['FACEBOOK_APP_ID'], ENV['FACEBOOK_APP_SECRET']
end
```

### OAuth Controller
```ruby
# app/controllers/omniauth_controller.rb
class OmniauthController < ApplicationController
  def google_oauth2
    handle_auth("Google")
  end

  def facebook
    handle_auth("Facebook")
  end

  private

  def handle_auth(kind)
    @user = User.from_omniauth(request.env["omniauth.auth"])
    if @user.persisted?
      sign_in_and_redirect @user, event: :authentication
    else
      session["devise.auth_data"] = request.env["omniauth.auth"].except(:extra)
      redirect_to new_user_registration_url, alert: @user.errors.full_messages.join("\n")
    end
  end
end
```

## Pundit Implementation

### Policy Setup
```ruby
# app/controllers/application_controller.rb
class ApplicationController < ActionController::Base
  include Pundit::Authorization
  
  after_action :verify_authorized, except: :index
  after_action :verify_policy_scoped, only: :index
  
  rescue_from Pundit::NotAuthorizedError, with: :user_not_authorized
  
  private
  
  def user_not_authorized
    flash[:alert] = "You are not authorized to perform this action."
    redirect_to(request.referrer || root_path)
  end
end
```

### Article Policy
```ruby
# app/policies/article_policy.rb
class ArticlePolicy < ApplicationPolicy
  class Scope < Scope
    def resolve
      if user.admin?
        scope.all
      else
        scope.where(published: true)
      end
    end
  end

  def index?
    true
  end

  def show?
    record.published? || user.admin? || record.user == user
  end

  def create?
    user.present?
  end

  def update?
    user.admin? || record.user == user
  end

  def destroy?
    user.admin? || record.user == user
  end
end
```

### Controller Usage
```ruby
# app/controllers/articles_controller.rb
class ArticlesController < ApplicationController
  def index
    @articles = policy_scope(Article)
  end

  def show
    @article = Article.find(params[:id])
    authorize @article
  end

  def new
    @article = Article.new
    authorize @article
  end

  def create
    @article = Article.new(article_params)
    authorize @article
    
    if @article.save
      redirect_to @article, notice: 'Article was successfully created.'
    else
      render :new
    end
  end

  def edit
    @article = Article.find(params[:id])
    authorize @article
  end

  def update
    @article = Article.find(params[:id])
    authorize @article
    
    if @article.update(article_params)
      redirect_to @article, notice: 'Article was successfully updated.'
    else
      render :edit
    end
  end

  def destroy
    @article = Article.find(params[:id])
    authorize @article
    
    @article.destroy
    redirect_to articles_url, notice: 'Article was successfully deleted.'
  end

  private

  def article_params
    params.require(:article).permit(:title, :content, :published)
  end
end
```

## CanCanCan Implementation

### Ability Setup
```ruby
# app/models/ability.rb
class Ability
  include CanCan::Ability

  def initialize(user)
    user ||= User.new # guest user

    if user.admin?
      can :manage, :all
    else
      can :read, Article, published: true
      can :read, Article, user_id: user.id
      can :create, Article
      can :update, Article, user_id: user.id
      can :destroy, Article, user_id: user.id
    end
  end
end
```

### Controller Usage
```ruby
# app/controllers/articles_controller.rb
class ArticlesController < ApplicationController
  load_and_authorize_resource

  def index
    @articles = @articles.accessible_by(current_ability)
  end

  def show
    # @article is already loaded and authorized
  end

  def new
    # @article is already loaded and authorized
  end

  def create
    @article.user = current_user
    if @article.save
      redirect_to @article, notice: 'Article was successfully created.'
    else
      render :new
    end
  end

  def edit
    # @article is already loaded and authorized
  end

  def update
    if @article.update(article_params)
      redirect_to @article, notice: 'Article was successfully updated.'
    else
      render :edit
    end
  end

  def destroy
    @article.destroy
    redirect_to articles_url, notice: 'Article was successfully deleted.'
  end

  private

  def article_params
    params.require(:article).permit(:title, :content, :published)
  end
end
```

## Testing Authentication

### Controller Tests
```ruby
# test/controllers/articles_controller_test.rb
require 'test_helper'

class ArticlesControllerTest < ActionDispatch::IntegrationTest
  setup do
    @user = users(:one)
    @article = articles(:one)
  end

  test "should get index when authenticated" do
    sign_in @user
    get articles_url
    assert_response :success
  end

  test "should not get index when not authenticated" do
    get articles_url
    assert_redirected_to new_user_session_url
  end

  test "should create article when authenticated" do
    sign_in @user
    assert_difference('Article.count') do
      post articles_url, params: { article: { title: 'Test', content: 'Content' } }
    end
    assert_redirected_to article_url(Article.last)
  end
end
```

### Policy Tests
```ruby
# test/policies/article_policy_test.rb
require 'test_helper'

class ArticlePolicyTest < ActiveSupport::TestCase
  setup do
    @user = users(:one)
    @admin = users(:admin)
    @article = articles(:one)
  end

  test "admin can manage all articles" do
    policy = ArticlePolicy.new(@admin, @article)
    assert policy.update?
    assert policy.destroy?
  end

  test "user can manage own articles" do
    @article.update(user: @user)
    policy = ArticlePolicy.new(@user, @article)
    assert policy.update?
    assert policy.destroy?
  end

  test "user cannot manage other users articles" do
    policy = ArticlePolicy.new(@user, @article)
    refute policy.update?
    refute policy.destroy?
  end
end
```

## Security Headers

### Application Configuration
```ruby
# config/application.rb
module YourApp
  class Application < Rails::Application
    config.action_dispatch.default_headers = {
      'X-Frame-Options' => 'SAMEORIGIN',
      'X-XSS-Protection' => '1; mode=block',
      'X-Content-Type-Options' => 'nosniff',
      'X-Download-Options' => 'noopen',
      'X-Permitted-Cross-Domain-Policies' => 'none',
      'Referrer-Policy' => 'strict-origin-when-cross-origin'
    }
  end
end
```

### Cookie Configuration
```ruby
# config/initializers/session_store.rb
Rails.application.config.session_store :cookie_store,
  key: '_your_app_session',
  secure: Rails.env.production?,
  httponly: true,
  same_site: :lax
``` 