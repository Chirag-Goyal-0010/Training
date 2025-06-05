# Rails Model Syntax

## Basic Model Definition

```ruby
# app/models/post.rb
class Post < ApplicationRecord
  # Table name (optional)
  self.table_name = 'blog_posts'
  
  # Primary key (optional)
  self.primary_key = 'post_id'
  
  # Timestamps (optional)
  self.record_timestamps = false
end
```

## Model Associations

```ruby
# One-to-One
belongs_to :user
has_one :profile
has_one :address, as: :addressable

# One-to-Many
has_many :comments
has_many :posts, dependent: :destroy
has_many :posts, through: :categories

# Many-to-Many
has_and_belongs_to_many :tags
has_many :tags, through: :post_tags

# Polymorphic
belongs_to :commentable, polymorphic: true
has_many :comments, as: :commentable

# Association Options
belongs_to :user, optional: true
has_many :comments, dependent: :destroy
has_one :profile, dependent: :nullify
has_many :posts, through: :categories, source: :articles
```

## Model Validations

```ruby
# Presence
validates :title, presence: true
validates :email, presence: true, allow_nil: true

# Uniqueness
validates :email, uniqueness: true
validates :username, uniqueness: { case_sensitive: false }
validates :slug, uniqueness: { scope: :category_id }

# Length
validates :title, length: { minimum: 3 }
validates :content, length: { maximum: 1000 }
validates :password, length: { in: 6..20 }
validates :code, length: { is: 6 }

# Format
validates :email, format: { with: URI::MailTo::EMAIL_REGEXP }
validates :phone, format: { with: /\A\+?\d{10,15}\z/ }

# Numericality
validates :age, numericality: true
validates :points, numericality: { only_integer: true }
validates :price, numericality: { greater_than: 0 }
validates :quantity, numericality: { less_than: 100 }

# Inclusion/Exclusion
validates :status, inclusion: { in: %w[pending active archived] }
validates :role, exclusion: { in: %w[admin superadmin] }

# Custom
validate :custom_validation_method
validates_with CustomValidator
```

## Model Callbacks

```ruby
# Before callbacks
before_validation :normalize_title
before_save :set_published_at
before_create :generate_slug
before_update :update_timestamp
before_destroy :check_dependencies

# After callbacks
after_validation :log_validation_errors
after_save :update_cache
after_create :send_notification
after_update :update_search_index
after_destroy :cleanup_files

# Around callbacks
around_save :log_changes
around_create :wrap_in_transaction
around_destroy :archive_record

# Callback options
before_validation :normalize_title, on: :create
after_save :update_cache, if: :published?
before_destroy :check_dependencies, unless: :admin?
```

## Model Scopes

```ruby
# Basic scopes
scope :active, -> { where(active: true) }
scope :recent, -> { order(created_at: :desc) }
scope :popular, -> { where('views_count > ?', 100) }

# Scopes with parameters
scope :created_after, ->(date) { where('created_at > ?', date) }
scope :with_tag, ->(tag) { joins(:tags).where(tags: { name: tag }) }

# Scopes with conditions
scope :published, -> { where(published: true) }
scope :draft, -> { where(published: false) }

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
```

## Model Methods

```ruby
# Class methods
def self.search(query)
  where('title ILIKE ?', "%#{query}%")
end

def self.by_author(author)
  joins(:user).where(users: { username: author })
end

# Instance methods
def full_name
  "#{first_name} #{last_name}"
end

def publish!
  update(published: true, published_at: Time.current)
end

# Private methods
private

def generate_slug
  self.slug = title.parameterize
end

def notify_subscribers
  PostNotificationJob.perform_later(self)
end
```

## Model Attributes

```ruby
# Attribute accessors
attr_accessor :password_confirmation
attr_reader :calculated_field
attr_writer :internal_field

# Attribute protection
attr_protected :admin
attr_readonly :slug

# Attribute serialization
serialize :preferences, JSON
serialize :tags, Array

# Attribute encryption
encrypts :secret_content
encrypts :api_key, deterministic: true
```

## Model Queries

```ruby
# Basic queries
Post.find(1)
Post.find_by(title: 'Hello')
Post.where(published: true)
Post.order(created_at: :desc)

# Complex queries
Post.joins(:user).where(users: { role: 'admin' })
Post.includes(:comments).where('comments.created_at > ?', 1.week.ago)
Post.group(:category_id).having('COUNT(*) > 1')

# Chaining
Post.published.recent.limit(10)
Post.with_comments.where('comments.count > 5')

# Aggregations
Post.count
Post.average(:views_count)
Post.sum(:likes_count)
Post.maximum(:created_at)
```

## Model Transactions

```ruby
# Basic transaction
Post.transaction do
  post.save!
  post.comments.create!(content: 'Great post!')
end

# Nested transaction
Post.transaction do
  post.save!
  Comment.transaction do
    comment.save!
  end
end

# Transaction with rollback
Post.transaction do
  post.save!
  raise ActiveRecord::Rollback if post.invalid?
end
```

## Model Callbacks

```ruby
# Callback registration
before_validation :normalize_title
after_save :update_cache
around_create :wrap_in_transaction

# Callback options
before_validation :normalize_title, on: :create
after_save :update_cache, if: :published?
before_destroy :check_dependencies, unless: :admin?

# Callback methods
private

def normalize_title
  self.title = title.strip.titleize if title.present?
end

def update_cache
  Rails.cache.write("post_#{id}", self)
end

def wrap_in_transaction
  ActiveRecord::Base.transaction do
    yield
  end
end
``` 