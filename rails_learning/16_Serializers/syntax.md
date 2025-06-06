# Serializers Syntax Examples

## Basic Serialization

### Using as_json
```ruby
# app/models/user.rb
class User < ApplicationRecord
  def as_json(options = {})
    {
      id: id,
      name: name,
      email: email,
      created_at: created_at
    }
  end
end

# Usage in controller
def show
  @user = User.find(params[:id])
  render json: @user
end
```

### Using to_json with Options
```ruby
# app/models/user.rb
class User < ApplicationRecord
  def as_json(options = {})
    super(options.merge(
      only: [:id, :name, :email],
      methods: [:full_name],
      include: { posts: { only: [:id, :title] } }
    ))
  end

  def full_name
    "#{first_name} #{last_name}"
  end
end
```

## ActiveModel::Serializers

### Basic Serializer
```ruby
# app/serializers/user_serializer.rb
class UserSerializer < ActiveModel::Serializer
  attributes :id, :name, :email, :created_at
  
  def name
    "#{object.first_name} #{object.last_name}"
  end
end
```

### Serializer with Associations
```ruby
# app/serializers/user_serializer.rb
class UserSerializer < ActiveModel::Serializer
  attributes :id, :name, :email
  
  has_many :posts
  belongs_to :company
  
  def name
    "#{object.first_name} #{object.last_name}"
  end
end

# app/serializers/post_serializer.rb
class PostSerializer < ActiveModel::Serializer
  attributes :id, :title, :content
  
  belongs_to :user
end
```

### Conditional Attributes
```ruby
# app/serializers/user_serializer.rb
class UserSerializer < ActiveModel::Serializer
  attributes :id, :name, :email
  
  attribute :admin, if: :admin?
  attribute :last_login, if: :show_last_login?
  
  def admin?
    object.admin?
  end
  
  def show_last_login?
    scope.admin?
  end
end
```

## Fast JSON API

### Basic Serializer
```ruby
# app/serializers/user_serializer.rb
class UserSerializer
  include FastJsonapi::ObjectSerializer
  
  attributes :name, :email
  
  attribute :full_name do |object|
    "#{object.first_name} #{object.last_name}"
  end
end
```

### Serializer with Associations
```ruby
# app/serializers/user_serializer.rb
class UserSerializer
  include FastJsonapi::ObjectSerializer
  
  attributes :name, :email
  
  has_many :posts
  belongs_to :company
  
  attribute :post_count do |object|
    object.posts.count
  end
end
```

### Caching
```ruby
# app/serializers/user_serializer.rb
class UserSerializer
  include FastJsonapi::ObjectSerializer
  
  attributes :name, :email
  
  cache_options enabled: true, cache_length: 12.hours
  
  attribute :full_name do |object|
    "#{object.first_name} #{object.last_name}"
  end
end
```

## Jbuilder

### Basic Template
```ruby
# app/views/api/v1/users/show.json.jbuilder
json.user do
  json.id @user.id
  json.name @user.name
  json.email @user.email
  json.created_at @user.created_at
end
```

### Template with Associations
```ruby
# app/views/api/v1/users/show.json.jbuilder
json.user do
  json.id @user.id
  json.name @user.name
  json.email @user.email
  
  json.posts @user.posts do |post|
    json.id post.id
    json.title post.title
    json.content post.content
  end
  
  json.company do
    json.id @user.company.id
    json.name @user.company.name
  end
end
```

### Partial Templates
```ruby
# app/views/api/v1/users/_user.json.jbuilder
json.id user.id
json.name user.name
json.email user.email

# app/views/api/v1/users/index.json.jbuilder
json.users @users do |user|
  json.partial! 'api/v1/users/user', user: user
end
```

## Custom Serialization

### Custom Serializer Class
```ruby
# app/serializers/custom_user_serializer.rb
class CustomUserSerializer
  def initialize(user)
    @user = user
  end
  
  def as_json
    {
      id: @user.id,
      name: @user.name,
      email: @user.email,
      posts: @user.posts.map { |post| CustomPostSerializer.new(post).as_json }
    }
  end
end

# Usage in controller
def show
  @user = User.find(params[:id])
  render json: CustomUserSerializer.new(@user).as_json
end
```

### Versioned Serializers
```ruby
# app/serializers/v1/user_serializer.rb
module V1
  class UserSerializer < ActiveModel::Serializer
    attributes :id, :name, :email
  end
end

# app/serializers/v2/user_serializer.rb
module V2
  class UserSerializer < ActiveModel::Serializer
    attributes :id, :name, :email, :phone
    
    has_many :posts
  end
end

# Usage in controller
def show
  @user = User.find(params[:id])
  render json: @user, serializer: "V#{params[:version]}/UserSerializer".constantize
end
```

## Error Handling

### Serializer Error Handling
```ruby
# app/serializers/user_serializer.rb
class UserSerializer < ActiveModel::Serializer
  attributes :id, :name, :email
  
  def name
    object.name.presence || 'Anonymous'
  rescue StandardError => e
    Rails.logger.error("Error serializing user name: #{e.message}")
    'Unknown'
  end
end
```

### API Error Response
```ruby
# app/controllers/api/v1/base_controller.rb
class Api::V1::BaseController < ApplicationController
  rescue_from ActiveRecord::RecordNotFound do |e|
    render json: {
      error: 'Not Found',
      message: e.message
    }, status: :not_found
  end
  
  rescue_from ActiveRecord::RecordInvalid do |e|
    render json: {
      error: 'Validation Error',
      messages: e.record.errors.full_messages
    }, status: :unprocessable_entity
  end
end
```

## Testing Serializers

### Serializer Tests
```ruby
# test/serializers/user_serializer_test.rb
require 'test_helper'

class UserSerializerTest < ActiveSupport::TestCase
  test "serializes user attributes" do
    user = users(:one)
    serializer = UserSerializer.new(user)
    
    assert_equal user.id, serializer.as_json[:id]
    assert_equal user.name, serializer.as_json[:name]
    assert_equal user.email, serializer.as_json[:email]
  end
  
  test "includes associated posts" do
    user = users(:one)
    serializer = UserSerializer.new(user)
    
    assert_includes serializer.as_json.keys, :posts
    assert_equal user.posts.count, serializer.as_json[:posts].length
  end
end
```

### Integration Tests
```ruby
# test/controllers/api/v1/users_controller_test.rb
require 'test_helper'

class Api::V1::UsersControllerTest < ActionDispatch::IntegrationTest
  test "returns user with correct attributes" do
    user = users(:one)
    get api_v1_user_url(user), as: :json
    
    assert_response :success
    json_response = JSON.parse(response.body)
    
    assert_equal user.id, json_response['id']
    assert_equal user.name, json_response['name']
    assert_equal user.email, json_response['email']
  end
end 