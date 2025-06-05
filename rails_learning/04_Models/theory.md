# Rails Models - Code Examples

## Basic Model Structure

```ruby
# app/models/post.rb
class Post < ApplicationRecord
  # Associations
  belongs_to :user
  has_many :comments, dependent: :destroy
  has_many :tags, through: :post_tags
  
  # Validations
  validates :title, presence: true, length: { minimum: 3 }
  validates :content, presence: true
  validates :slug, uniqueness: true
  
  # Callbacks
  before_validation :generate_slug
  after_create :notify_subscribers
  
  # Scopes
  scope :published, -> { where(published: true) }
  scope :recent, -> { order(created_at: :desc) }
  scope :popular, -> { where('views_count > ?', 100) }
  
  # Instance methods
  def excerpt
    content.truncate(100)
  end
  
  def publish!
    update(published: true, published_at: Time.current)
  end
  
  private
  
  def generate_slug
    self.slug = title.parameterize if title_changed?
  end
  
  def notify_subscribers
    PostNotificationJob.perform_later(self)
  end
end
```

## Model Associations

```ruby
# app/models/user.rb
class User < ApplicationRecord
  has_many :posts, dependent: :destroy
  has_many :comments, dependent: :destroy
  has_many :likes, dependent: :destroy
  has_many :liked_posts, through: :likes, source: :post
  
  has_one :profile, dependent: :destroy
  has_one :address, as: :addressable
  
  belongs_to :role
  belongs_to :organization, optional: true
end

# app/models/comment.rb
class Comment < ApplicationRecord
  belongs_to :user
  belongs_to :post
  belongs_to :parent, class_name: 'Comment', optional: true
  has_many :replies, class_name: 'Comment', foreign_key: :parent_id
end

# app/models/tag.rb
class Tag < ApplicationRecord
  has_many :post_tags
  has_many :posts, through: :post_tags
end
```

## Model Validations

```ruby
# app/models/user.rb
class User < ApplicationRecord
  # Presence validations
  validates :email, presence: true
  validates :username, presence: true
  
  # Uniqueness validations
  validates :email, uniqueness: { case_sensitive: false }
  validates :username, uniqueness: true
  
  # Format validations
  validates :email, format: { with: URI::MailTo::EMAIL_REGEXP }
  validates :phone, format: { with: /\A\+?\d{10,15}\z/ }, allow_blank: true
  
  # Length validations
  validates :password, length: { minimum: 8 }, if: :password_required?
  validates :bio, length: { maximum: 500 }
  
  # Numericality validations
  validates :age, numericality: { greater_than: 0, less_than: 120 }
  validates :points, numericality: { only_integer: true }
  
  # Custom validations
  validate :password_complexity
  validate :username_format
  
  private
  
  def password_complexity
    return if password.blank?
    
    unless password.match?(/\A(?=.*[A-Z])(?=.*[a-z])(?=.*\d)/)
      errors.add(:password, 'must include uppercase, lowercase, and number')
    end
  end
  
  def username_format
    return if username.blank?
    
    unless username.match?(/\A[a-zA-Z0-9_]+\z/)
      errors.add(:username, 'can only contain letters, numbers, and underscores')
    end
  end
end
```

## Model Callbacks

```ruby
# app/models/post.rb
class Post < ApplicationRecord
  # Before callbacks
  before_validation :normalize_title
  before_save :set_published_at
  before_create :generate_slug
  
  # After callbacks
  after_create :notify_subscribers
  after_update :update_search_index
  after_destroy :cleanup_attachments
  
  # Around callbacks
  around_save :log_changes
  around_destroy :archive_post
  
  private
  
  def normalize_title
    self.title = title.strip.titleize if title.present?
  end
  
  def set_published_at
    self.published_at = Time.current if published_changed? && published?
  end
  
  def generate_slug
    self.slug = title.parameterize
  end
  
  def notify_subscribers
    PostNotificationJob.perform_later(self)
  end
  
  def update_search_index
    SearchIndexer.perform_later(self)
  end
  
  def cleanup_attachments
    AttachmentCleanupJob.perform_later(self)
  end
  
  def log_changes
    old_attributes = attributes.dup
    yield
    ChangeLogger.log(self, old_attributes)
  end
  
  def archive_post
    archived_post = attributes.merge(archived_at: Time.current)
    yield
    ArchivedPost.create!(archived_post)
  end
end
```

## Model Scopes

```ruby
# app/models/post.rb
class Post < ApplicationRecord
  # Basic scopes
  scope :published, -> { where(published: true) }
  scope :draft, -> { where(published: false) }
  scope :recent, -> { order(created_at: :desc) }
  
  # Scopes with parameters
  scope :created_after, ->(date) { where('created_at > ?', date) }
  scope :with_tag, ->(tag) { joins(:tags).where(tags: { name: tag }) }
  
  # Scopes with conditions
  scope :popular, -> { where('views_count > ?', 100) }
  scope :trending, -> { where('created_at > ?', 1.week.ago) }
  
  # Scopes with joins
  scope :with_comments, -> { joins(:comments).distinct }
  scope :with_user, -> { includes(:user) }
  
  # Scopes with group
  scope :by_month, -> { group("DATE_TRUNC('month', created_at)") }
  scope :by_category, -> { group(:category_id).count }
  
  # Scopes with having
  scope :with_multiple_comments, -> { 
    joins(:comments)
      .group('posts.id')
      .having('COUNT(comments.id) > 1')
  }
end
```

## Model Methods

```ruby
# app/models/post.rb
class Post < ApplicationRecord
  # Class methods
  def self.search(query)
    where('title ILIKE ? OR content ILIKE ?', "%#{query}%", "%#{query}%")
  end
  
  def self.by_author(author)
    joins(:user).where(users: { username: author })
  end
  
  # Instance methods
  def publish!
    update(published: true, published_at: Time.current)
  end
  
  def unpublish!
    update(published: false, published_at: nil)
  end
  
  def increment_view_count
    increment!(:views_count)
  end
  
  def add_comment(user, content)
    comments.create(user: user, content: content)
  end
  
  # Private methods
  private
  
  def generate_slug
    self.slug = title.parameterize
  end
  
  def notify_subscribers
    PostNotificationJob.perform_later(self)
  end
end
```

## Model Testing

```ruby
# test/models/post_test.rb
require 'test_helper'

class PostTest < ActiveSupport::TestCase
  setup do
    @user = users(:one)
    @post = posts(:one)
  end
  
  test "should be valid" do
    assert @post.valid?
  end
  
  test "title should be present" do
    @post.title = nil
    assert_not @post.valid?
  end
  
  test "title should be at least 3 characters" do
    @post.title = "ab"
    assert_not @post.valid?
  end
  
  test "should generate slug from title" do
    @post.title = "New Post Title"
    @post.valid?
    assert_equal "new-post-title", @post.slug
  end
  
  test "should increment view count" do
    assert_difference '@post.views_count', 1 do
      @post.increment_view_count
    end
  end
  
  test "should add comment" do
    assert_difference '@post.comments.count', 1 do
      @post.add_comment(@user, "Great post!")
    end
  end
end
```

## Model Security

```ruby
# app/models/user.rb
class User < ApplicationRecord
  # Attribute protection
  attr_protected :admin
  
  # Attribute accessors
  attr_accessor :password_confirmation
  
  # Secure password
  has_secure_password
  
  # Token authentication
  before_create :generate_authentication_token
  
  # Role-based access
  enum role: { user: 0, moderator: 1, admin: 2 }
  
  private
  
  def generate_authentication_token
    loop do
      self.authentication_token = SecureRandom.hex(20)
      break unless self.class.exists?(authentication_token: authentication_token)
    end
  end
end

# app/models/post.rb
class Post < ApplicationRecord
  # Mass assignment protection
  attr_readonly :slug
  
  # Attribute encryption
  encrypts :secret_content
  
  # Audit logging
  audited
  
  # Paper trail
  has_paper_trail
end
```

## Model Performance

```ruby
# app/models/post.rb
class Post < ApplicationRecord
  # Eager loading
  scope :with_comments, -> { includes(:comments) }
  scope :with_user_and_comments, -> { includes(:user, :comments) }
  
  # Counter cache
  belongs_to :user, counter_cache: true
  
  # Touch option
  belongs_to :category, touch: true
  
  # Caching
  def cached_comments_count
    Rails.cache.fetch([self, 'comments_count']) do
      comments.count
    end
  end
  
  # Batch processing
  def self.process_in_batches
    find_each(batch_size: 1000) do |post|
      post.process
    end
  end
end

# app/models/comment.rb
class Comment < ApplicationRecord
  # Counter cache
  belongs_to :post, counter_cache: true
  
  # Touch option
  belongs_to :user, touch: true
end
``` 