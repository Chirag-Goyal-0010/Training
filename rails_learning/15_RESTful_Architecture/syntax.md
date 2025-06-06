# RESTful Architecture Syntax Examples

## Routes Configuration

### Basic Resource Routes
```ruby
# config/routes.rb
Rails.application.routes.draw do
  resources :articles
end
```

This creates the following routes:
```
GET    /articles          # index
GET    /articles/new      # new
POST   /articles          # create
GET    /articles/:id      # show
GET    /articles/:id/edit # edit
PATCH  /articles/:id      # update
PUT    /articles/:id      # update
DELETE /articles/:id      # destroy
```

### Nested Resources
```ruby
# config/routes.rb
Rails.application.routes.draw do
  resources :articles do
    resources :comments
  end
end
```

### Custom Member and Collection Routes
```ruby
# config/routes.rb
Rails.application.routes.draw do
  resources :articles do
    member do
      post :publish
      post :unpublish
    end
    
    collection do
      get :search
      get :archived
    end
  end
end
```

## Controller Implementation

### Basic REST Controller
```ruby
# app/controllers/articles_controller.rb
class ArticlesController < ApplicationController
  def index
    @articles = Article.all
  end

  def show
    @article = Article.find(params[:id])
  end

  def new
    @article = Article.new
  end

  def create
    @article = Article.new(article_params)
    
    if @article.save
      redirect_to @article, notice: 'Article was successfully created.'
    else
      render :new
    end
  end

  def edit
    @article = Article.find(params[:id])
  end

  def update
    @article = Article.find(params[:id])
    
    if @article.update(article_params)
      redirect_to @article, notice: 'Article was successfully updated.'
    else
      render :edit
    end
  end

  def destroy
    @article = Article.find(params[:id])
    @article.destroy
    
    redirect_to articles_url, notice: 'Article was successfully deleted.'
  end

  private

  def article_params
    params.require(:article).permit(:title, :content)
  end
end
```

### Custom Action Implementation
```ruby
# app/controllers/articles_controller.rb
class ArticlesController < ApplicationController
  # ... standard REST actions ...

  def publish
    @article = Article.find(params[:id])
    @article.update(published: true)
    redirect_to @article, notice: 'Article was published.'
  end

  def search
    @articles = Article.where('title LIKE ?', "%#{params[:query]}%")
    render :index
  end
end
```

## View Templates

### Index View
```erb
<%# app/views/articles/index.html.erb %>
<h1>Articles</h1>

<%= link_to 'New Article', new_article_path %>

<table>
  <thead>
    <tr>
      <th>Title</th>
      <th>Content</th>
      <th>Actions</th>
    </tr>
  </thead>
  <tbody>
    <% @articles.each do |article| %>
      <tr>
        <td><%= article.title %></td>
        <td><%= article.content %></td>
        <td>
          <%= link_to 'Show', article %>
          <%= link_to 'Edit', edit_article_path(article) %>
          <%= link_to 'Delete', article, method: :delete, data: { confirm: 'Are you sure?' } %>
        </td>
      </tr>
    <% end %>
  </tbody>
</table>
```

### Form Partial
```erb
<%# app/views/articles/_form.html.erb %>
<%= form_with(model: article, local: true) do |form| %>
  <% if article.errors.any? %>
    <div class="error-messages">
      <h2><%= pluralize(article.errors.count, "error") %> prohibited this article from being saved:</h2>
      <ul>
        <% article.errors.full_messages.each do |message| %>
          <li><%= message %></li>
        <% end %>
      </ul>
    </div>
  <% end %>

  <div class="field">
    <%= form.label :title %>
    <%= form.text_field :title %>
  </div>

  <div class="field">
    <%= form.label :content %>
    <%= form.text_area :content %>
  </div>

  <div class="actions">
    <%= form.submit %>
  </div>
<% end %>
```

## API Implementation

### JSON API Controller
```ruby
# app/controllers/api/v1/articles_controller.rb
module Api
  module V1
    class ArticlesController < ApplicationController
      def index
        @articles = Article.all
        render json: @articles
      end

      def show
        @article = Article.find(params[:id])
        render json: @article
      end

      def create
        @article = Article.new(article_params)
        
        if @article.save
          render json: @article, status: :created
        else
          render json: @article.errors, status: :unprocessable_entity
        end
      end

      def update
        @article = Article.find(params[:id])
        
        if @article.update(article_params)
          render json: @article
        else
          render json: @article.errors, status: :unprocessable_entity
        end
      end

      def destroy
        @article = Article.find(params[:id])
        @article.destroy
        head :no_content
      end

      private

      def article_params
        params.require(:article).permit(:title, :content)
      end
    end
  end
end
```

### API Routes
```ruby
# config/routes.rb
Rails.application.routes.draw do
  namespace :api do
    namespace :v1 do
      resources :articles
    end
  end
end
```

## Error Handling

### Controller Error Handling
```ruby
# app/controllers/application_controller.rb
class ApplicationController < ActionController::Base
  rescue_from ActiveRecord::RecordNotFound, with: :not_found
  rescue_from ActionController::ParameterMissing, with: :bad_request

  private

  def not_found
    render json: { error: 'Resource not found' }, status: :not_found
  end

  def bad_request
    render json: { error: 'Bad request' }, status: :bad_request
  end
end
```

## Testing REST Endpoints

### Controller Tests
```ruby
# test/controllers/articles_controller_test.rb
require 'test_helper'

class ArticlesControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get articles_url
    assert_response :success
  end

  test "should create article" do
    assert_difference('Article.count') do
      post articles_url, params: { article: { title: 'Test', content: 'Content' } }
    end
    assert_redirected_to article_url(Article.last)
  end

  test "should update article" do
    article = articles(:one)
    patch article_url(article), params: { article: { title: 'Updated' } }
    assert_redirected_to article_url(article)
  end

  test "should destroy article" do
    article = articles(:one)
    assert_difference('Article.count', -1) do
      delete article_url(article)
    end
    assert_redirected_to articles_url
  end
end
```

## API Documentation

### API Documentation with RDoc
```ruby
# app/controllers/api/v1/articles_controller.rb
module Api
  module V1
    class ArticlesController < ApplicationController
      # @api {get} /api/v1/articles List all articles
      # @apiName GetArticles
      # @apiGroup Articles
      # @apiVersion 1.0.0
      #
      # @apiSuccess {Object[]} articles List of articles
      # @apiSuccess {Number} articles.id Article ID
      # @apiSuccess {String} articles.title Article title
      # @apiSuccess {String} articles.content Article content
      #
      # @apiSuccessExample {json} Success-Response:
      #     HTTP/1.1 200 OK
      #     [
      #       {
      #         "id": 1,
      #         "title": "First Article",
      #         "content": "Content"
      #       }
      #     ]
      def index
        @articles = Article.all
        render json: @articles
      end
    end
  end
end
``` 