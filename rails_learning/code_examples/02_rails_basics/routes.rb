# config/routes.rb

Rails.application.routes.draw do
  # Basic routes
  get 'welcome/index'
  get 'about', to: 'pages#about'
  get 'contact', to: 'pages#contact'
  
  # RESTful routes
  resources :posts
  
  # RESTful routes with only specific actions
  resources :comments, only: [:create, :destroy]
  
  # RESTful routes with except
  resources :users, except: [:destroy]
  
  # Nested routes
  resources :posts do
    resources :comments
    resources :likes, only: [:create, :destroy]
  end
  
  # Namespace routes
  namespace :admin do
    resources :posts
    resources :users
  end
  
  # Scope routes
  scope :api do
    resources :posts
  end
  
  # Custom member routes
  resources :posts do
    member do
      post 'publish'
      post 'unpublish'
    end
  end
  
  # Custom collection routes
  resources :posts do
    collection do
      get 'search'
      get 'popular'
    end
  end
  
  # Root route
  root 'welcome#index'
  
  # Custom constraints
  constraints(lambda { |req| req.format == :json }) do
    resources :api_posts
  end
  
  # Route with parameters
  get 'posts/:id/comments/:comment_id', to: 'comments#show'
  
  # Route with format
  get 'posts.:format', to: 'posts#index'
  
  # Route with default format
  get 'posts', to: 'posts#index', defaults: { format: 'json' }
  
  # Route with as: option for custom path helper
  get 'login', to: 'sessions#new', as: :login
  
  # Route with to: option for custom controller action
  get 'dashboard', to: 'users#dashboard'
  
  # Route with constraints
  get 'users/:id', to: 'users#show', constraints: { id: /[0-9]+/ }
  
  # Route with redirect
  get 'old_path', to: redirect('new_path')
  
  # Route with subdomain
  constraints(lambda { |req| req.subdomain.present? }) do
    get '/', to: 'subdomains#show'
  end
end 