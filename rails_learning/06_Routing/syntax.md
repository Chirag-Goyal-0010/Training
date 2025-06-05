# Rails Routing - Syntax Reference

## Basic Route Syntax

```ruby
# HTTP Verb Routes
get 'path', to: 'controller#action'
post 'path', to: 'controller#action'
put 'path', to: 'controller#action'
patch 'path', to: 'controller#action'
delete 'path', to: 'controller#action'

# Root Route
root 'controller#action'

# Named Routes
get 'path', to: 'controller#action', as: :route_name

# Route with Parameters
get 'path/:id', to: 'controller#action'
get 'path/:id/:other_id', to: 'controller#action'

# Route with Defaults
get 'path(/:id)', to: 'controller#action', defaults: { id: 1 }

# Route with Constraints
get 'path/:id', to: 'controller#action', constraints: { id: /\d+/ }
```

## Resource Route Syntax

```ruby
# Basic Resource
resources :resource_name

# Resource with Options
resources :resource_name, only: [:index, :show]
resources :resource_name, except: [:destroy]
resources :resource_name, shallow: true

# Resource with Additional Routes
resources :resource_name do
  member do
    get 'action'
    post 'action'
    delete 'action'
  end
  
  collection do
    get 'action'
    post 'action'
  end
end

# Single Resource
resource :resource_name

# Resource with Path
resources :resource_name, path: 'custom_path'
```

## Nested Route Syntax

```ruby
# Basic Nesting
resources :parent do
  resources :child
end

# Shallow Nesting
resources :parent, shallow: true do
  resources :child
end

# Multiple Nesting
resources :grandparent do
  resources :parent do
    resources :child
  end
end

# Nested with Options
resources :parent do
  resources :child, only: [:index, :create]
end
```

## Namespace Syntax

```ruby
# Basic Namespace
namespace :name do
  resources :resource_name
end

# Namespace with Module
namespace :name, module: 'module_name' do
  resources :resource_name
end

# Namespace with Path
namespace :name, path: 'custom_path' do
  resources :resource_name
end

# Nested Namespace
namespace :parent do
  namespace :child do
    resources :resource_name
  end
end
```

## Scope Syntax

```ruby
# Basic Scope
scope :name do
  resources :resource_name
end

# Scope with Module
scope module: 'module_name' do
  resources :resource_name
end

# Scope with Path
scope path: 'custom_path' do
  resources :resource_name
end

# Scope with Constraints
scope constraints: { subdomain: 'api' } do
  resources :resource_name
end
```

## Constraint Syntax

```ruby
# Format Constraint
resources :resource_name, constraints: { format: 'json' }

# Parameter Constraint
resources :resource_name, constraints: { id: /\d+/ }

# Custom Constraint Class
class CustomConstraint
  def self.matches?(request)
    # constraint logic
  end
end

constraints CustomConstraint do
  resources :resource_name
end

# Lambda Constraint
constraints lambda { |req| req.remote_ip == "127.0.0.1" } do
  resources :resource_name
end
```

## Concern Syntax

```ruby
# Basic Concern
concern :name do
  resources :resource_name
end

# Concern with Options
concern :name do
  resources :resource_name, only: [:index, :create]
end

# Using Concerns
resources :resource_name, concerns: :name
resources :resource_name, concerns: [:name1, :name2]

# Nested Concern
concern :name do
  resources :resource_name do
    resources :nested_resource
  end
end
```

## Route Helper Syntax

```ruby
# Path Helpers
resource_name_path
resource_name_path(id)
new_resource_name_path
edit_resource_name_path(id)

# URL Helpers
resource_name_url
resource_name_url(id)
new_resource_name_url
edit_resource_name_url(id)

# Nested Path Helpers
parent_child_path(parent_id, child_id)
new_parent_child_path(parent_id)

# Namespaced Path Helpers
namespace_resource_name_path
namespace_resource_name_path(id)
```

## Route Testing Syntax

```ruby
# Basic Route Test
assert_routing "/path", controller: "controller", action: "action"

# Route with Parameters
assert_routing "/path/1", controller: "controller", action: "action", id: "1"

# Route with Method
assert_routing({ method: "post", path: "/path" },
               { controller: "controller", action: "action" })

# Route with Format
assert_routing "/path.json", controller: "controller", action: "action", format: "json"
```

## Route Security Syntax

```ruby
# Authentication Routes
devise_for :users

# Protected Routes
authenticate :user do
  resources :resource_name
end

# Role-based Routes
authenticate :user, ->(user) { user.admin? } do
  resources :resource_name
end

# Rate-limited Routes
constraints Rack::Attack do
  resources :resource_name
end
```

## Route Performance Syntax

```ruby
# Route Caching
get 'path', to: 'controller#action', cache: true

# Route Optimization
resources :resource_name, only: [:index, :show]

# Route Constraints for Performance
resources :resource_name, constraints: { format: 'json' }

# Route Grouping for Performance
scope module: 'module_name' do
  resources :resource_name
end
``` 