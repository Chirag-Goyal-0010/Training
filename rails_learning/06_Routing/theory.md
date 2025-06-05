# Rails Routing - Code Examples

## Basic Route Configuration

```ruby
# config/routes.rb
Rails.application.routes.draw do
  # Basic routes
  get 'welcome/index'
  get 'about', to: 'pages#about'
  get 'contact', to: 'pages#contact'
  
  # Root route
  root 'welcome#index'
  
  # Custom routes
  get 'posts/:id/preview', to: 'posts#preview'
  get 'posts/:id/publish', to: 'posts#publish'
  get 'posts/:id/unpublish', to: 'posts#unpublish'
  
  # Route with constraints
  get 'users/:id', to: 'users#show', constraints: { id: /\d+/ }
  
  # Route with defaults
  get 'posts(/:year(/:month))', to: 'posts#index', 
      defaults: { year: Time.current.year, month: Time.current.month }
end
```

## RESTful Resource Routes

```ruby
# config/routes.rb
Rails.application.routes.draw do
  # Basic resources
  resources :posts
  resources :comments
  resources :users
  
  # Resources with options
  resources :posts, only: [:index, :show, :new, :create]
  resources :posts, except: [:destroy]
  resources :posts, shallow: true
  
  # Resources with additional routes
  resources :posts do
    member do
      get 'preview'
      post 'publish'
      delete 'unpublish'
    end
    
    collection do
      get 'search'
      get 'popular'
      get 'recent'
    end
  end
end
```

## Nested Routes

```ruby
# config/routes.rb
Rails.application.routes.draw do
  # Nested resources
  resources :posts do
    resources :comments
  end
  
  # Shallow nesting
  resources :posts, shallow: true do
    resources :comments
  end
  
  # Multiple nesting
  resources :users do
    resources :posts do
      resources :comments
    end
  end
  
  # Nested resources with options
  resources :posts do
    resources :comments, only: [:create, :destroy]
    resources :likes, only: [:create, :destroy]
    resources :tags, only: [:index]
  end
end
```

## Namespace Routes

```ruby
# config/routes.rb
Rails.application.routes.draw do
  # Basic namespace
  namespace :admin do
    resources :posts
    resources :users
    resources :comments
  end
  
  # Namespace with module
  namespace :api, module: 'api' do
    namespace :v1 do
      resources :posts
      resources :users
    end
    
    namespace :v2 do
      resources :posts
      resources :users
    end
  end
  
  # Namespace with path
  namespace :admin, path: 'administrator' do
    resources :posts
    resources :users
  end
end
```

## Scope Routes

```ruby
# config/routes.rb
Rails.application.routes.draw do
  # Basic scope
  scope :admin do
    resources :posts
    resources :users
  end
  
  # Scope with module
  scope module: 'admin' do
    resources :posts
    resources :users
  end
  
  # Scope with path
  scope path: '/admin' do
    resources :posts
    resources :users
  end
  
  # Scope with constraints
  scope constraints: { subdomain: 'api' } do
    resources :posts
    resources :users
  end
end
```

## Route Constraints

```ruby
# config/routes.rb
Rails.application.routes.draw do
  # Format constraints
  resources :posts, constraints: { format: 'json' }
  
  # Parameter constraints
  resources :users, constraints: { id: /\d+/ }
  
  # Custom constraints
  class SubdomainConstraint
    def self.matches?(request)
      request.subdomain.present? && request.subdomain != 'www'
    end
  end
  
  constraints SubdomainConstraint do
    resources :posts
    resources :users
  end
  
  # Multiple constraints
  resources :posts, constraints: {
    id: /\d+/,
    format: 'json',
    subdomain: 'api'
  }
end
```

## Route Concerns

```ruby
# config/routes.rb
Rails.application.routes.draw do
  # Basic concern
  concern :commentable do
    resources :comments
  end
  
  resources :posts, concerns: :commentable
  resources :articles, concerns: :commentable
  
  # Concern with options
  concern :taggable do
    resources :tags, only: [:index, :create, :destroy]
  end
  
  resources :posts, concerns: [:commentable, :taggable]
  
  # Nested concern
  concern :commentable do
    resources :comments do
      resources :likes
    end
  end
end
```

## Route Testing

```ruby
# test/routes/posts_test.rb
require "test_helper"

class PostsRoutesTest < ActionDispatch::IntegrationTest
  test "routes to index" do
    assert_routing "/posts", controller: "posts", action: "index"
  end
  
  test "routes to show" do
    assert_routing "/posts/1", controller: "posts", action: "show", id: "1"
  end
  
  test "routes to new" do
    assert_routing "/posts/new", controller: "posts", action: "new"
  end
  
  test "routes to create" do
    assert_routing({ method: "post", path: "/posts" },
                  { controller: "posts", action: "create" })
  end
  
  test "routes to edit" do
    assert_routing "/posts/1/edit",
                  controller: "posts", action: "edit", id: "1"
  end
  
  test "routes to update" do
    assert_routing({ method: "patch", path: "/posts/1" },
                  { controller: "posts", action: "update", id: "1" })
  end
  
  test "routes to destroy" do
    assert_routing({ method: "delete", path: "/posts/1" },
                  { controller: "posts", action: "destroy", id: "1" })
  end
end
```

## Route Security

```ruby
# config/routes.rb
Rails.application.routes.draw do
  # Authentication routes
  devise_for :users
  
  # Protected routes
  authenticate :user do
    resources :posts
    resources :comments
  end
  
  # Role-based routes
  authenticate :user, ->(user) { user.admin? } do
    namespace :admin do
      resources :posts
      resources :users
    end
  end
  
  # Rate-limited routes
  constraints Rack::Attack do
    resources :api do
      resources :posts
      resources :users
    end
  end
  
  # IP-restricted routes
  constraints lambda { |req| req.remote_ip == "127.0.0.1" } do
    namespace :admin do
      resources :posts
      resources :users
    end
  end
end
```

## Route Performance

```ruby
# config/routes.rb
Rails.application.routes.draw do
  # Route caching
  get 'posts/:id', to: 'posts#show', cache: true
  
  # Route optimization
  resources :posts, only: [:index, :show] do
    resources :comments, only: [:index, :create]
  end
  
  # Route monitoring
  resources :posts do
    resources :comments
  end
  
  # Route constraints for performance
  resources :posts, constraints: { format: 'json' }
  resources :users, constraints: { format: 'json' }
  
  # Route grouping for performance
  scope module: 'api' do
    resources :posts
    resources :users
  end
end
``` 