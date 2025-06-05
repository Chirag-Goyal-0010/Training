# app/models/post.rb
class Post < ApplicationRecord
  # Associations
  belongs_to :user
  belongs_to :category, optional: true
  has_many :comments, dependent: :destroy
  has_many :likes, dependent: :destroy
  has_many :liking_users, through: :likes, source: :user
  
  # Validations
  validates :title, presence: true, length: { minimum: 5, maximum: 100 }
  validates :content, presence: true, length: { minimum: 50 }
  validates :slug, presence: true, uniqueness: true
  
  # Callbacks
  before_validation :generate_slug, on: :create
  after_create :notify_followers
  
  # Scopes
  scope :published, -> { where(published: true) }
  scope :recent, -> { order(created_at: :desc) }
  scope :popular, -> { joins(:likes).group('posts.id').order('COUNT(likes.id) DESC') }
  scope :by_category, ->(category) { where(category: category) }
  
  # Instance methods
  def excerpt
    truncate(content, length: 200)
  end
  
  def reading_time
    (content.split.size / 200.0).ceil
  end
  
  def liked_by?(user)
    liking_users.include?(user)
  end
  
  # Class methods
  def self.search(query)
    where("title ILIKE ? OR content ILIKE ?", "%#{query}%", "%#{query}%")
  end
  
  private
  
  def generate_slug
    self.slug = title.parameterize
  end
  
  def notify_followers
    user.followers.each do |follower|
      Notification.create(
        user: follower,
        notifiable: self,
        message: "#{user.name} published a new post: #{title}"
      )
    end
  end
end

# app/models/user.rb
class User < ApplicationRecord
  # Include devise modules
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :validatable,
         :confirmable, :lockable, :trackable
  
  # Associations
  has_many :posts, dependent: :destroy
  has_many :comments, dependent: :destroy
  has_many :likes, dependent: :destroy
  
  has_many :active_relationships, class_name: "Relationship",
                                foreign_key: "follower_id",
                                dependent: :destroy
  has_many :passive_relationships, class_name: "Relationship",
                                 foreign_key: "followed_id",
                                 dependent: :destroy
  has_many :following, through: :active_relationships, source: :followed
  has_many :followers, through: :passive_relationships, source: :follower
  
  # Validations
  validates :username, presence: true, uniqueness: true,
                      length: { minimum: 3, maximum: 20 },
                      format: { with: /\A[a-zA-Z0-9_]+\z/ }
  validates :email, presence: true, uniqueness: true,
                   format: { with: URI::MailTo::EMAIL_REGEXP }
  
  # Callbacks
  before_save :downcase_email
  after_create :send_welcome_email
  
  # Instance methods
  def feed
    following_ids = "SELECT followed_id FROM relationships WHERE follower_id = :user_id"
    Post.where("user_id IN (#{following_ids}) OR user_id = :user_id", user_id: id)
        .published.recent
  end
  
  def follow(other_user)
    following << other_user unless self == other_user
  end
  
  def unfollow(other_user)
    following.delete(other_user)
  end
  
  def following?(other_user)
    following.include?(other_user)
  end
  
  private
  
  def downcase_email
    self.email = email.downcase
  end
  
  def send_welcome_email
    UserMailer.welcome_email(self).deliver_later
  end
end

# app/models/comment.rb
class Comment < ApplicationRecord
  # Associations
  belongs_to :user
  belongs_to :post
  has_many :likes, as: :likeable, dependent: :destroy
  
  # Validations
  validates :content, presence: true, length: { minimum: 2, maximum: 1000 }
  
  # Scopes
  scope :recent, -> { order(created_at: :desc) }
  
  # Callbacks
  after_create :notify_post_author
  
  private
  
  def notify_post_author
    return if user == post.user
    
    Notification.create(
      user: post.user,
      notifiable: self,
      message: "#{user.name} commented on your post: #{post.title}"
    )
  end
end 