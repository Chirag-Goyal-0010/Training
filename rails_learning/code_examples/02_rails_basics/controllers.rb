# app/controllers/posts_controller.rb

class PostsController < ApplicationController
  # Filters
  before_action :set_post, only: [:show, :edit, :update, :destroy]
  before_action :authenticate_user!, except: [:index, :show]
  before_action :authorize_user!, only: [:edit, :update, :destroy]
  
  # Index action
  def index
    @posts = Post.published.recent.page(params[:page])
    respond_to do |format|
      format.html
      format.json { render json: @posts }
    end
  end
  
  # Show action
  def show
    @comments = @post.comments.recent
    @comment = Comment.new
  end
  
  # New action
  def new
    @post = Post.new
  end
  
  # Create action
  def create
    @post = current_user.posts.build(post_params)
    
    if @post.save
      redirect_to @post, notice: 'Post was successfully created.'
    else
      render :new, status: :unprocessable_entity
    end
  end
  
  # Edit action
  def edit
  end
  
  # Update action
  def update
    if @post.update(post_params)
      redirect_to @post, notice: 'Post was successfully updated.'
    else
      render :edit, status: :unprocessable_entity
    end
  end
  
  # Destroy action
  def destroy
    @post.destroy
    redirect_to posts_path, notice: 'Post was successfully deleted.'
  end
  
  # Custom action
  def search
    @posts = Post.search(params[:query])
    render :index
  end
  
  # Action with different formats
  def export
    @posts = Post.all
    
    respond_to do |format|
      format.csv { send_data @posts.to_csv }
      format.pdf { render pdf: "posts" }
    end
  end
  
  private
  
  # Strong parameters
  def post_params
    params.require(:post).permit(:title, :content, :published, :category_id)
  end
  
  # Before action methods
  def set_post
    @post = Post.find(params[:id])
  rescue ActiveRecord::RecordNotFound
    redirect_to posts_path, alert: 'Post not found.'
  end
  
  def authorize_user!
    unless @post.user == current_user
      redirect_to posts_path, alert: 'Not authorized.'
    end
  end
end

# app/controllers/application_controller.rb
class ApplicationController < ActionController::Base
  # Common before actions
  before_action :set_locale
  before_action :set_timezone
  
  # Common methods
  def current_user
    @current_user ||= User.find(session[:user_id]) if session[:user_id]
  end
  
  def authenticate_user!
    unless current_user
      redirect_to login_path, alert: 'Please log in.'
    end
  end
  
  private
  
  def set_locale
    I18n.locale = params[:locale] || I18n.default_locale
  end
  
  def set_timezone
    Time.zone = current_user&.timezone || 'UTC'
  end
  
  # Error handling
  rescue_from ActiveRecord::RecordNotFound do |exception|
    redirect_to root_path, alert: 'Record not found.'
  end
  
  rescue_from ActionController::ParameterMissing do |exception|
    redirect_to root_path, alert: 'Required parameters missing.'
  end
end 