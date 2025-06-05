# Basic Model Example

This example demonstrates fundamental Rails model concepts using a blog post system.

## Setup

1. Create a new Rails application:
```bash
rails new blog_example
cd blog_example
```

2. Generate the necessary models:
```bash
rails generate model User email:string username:string password_digest:string
rails generate model Post title:string content:text user:references published:boolean
rails generate model Comment content:text user:references post:references
rails generate model Tag name:string
rails generate model PostTag post:references tag:references
rails db:migrate
```

## Code Structure

### 1. User Model

```ruby
# app/models/user.rb
class User < ApplicationRecord
  # Associations
  has_many :posts, dependent: :destroy
  has_many :comments, dependent: :destroy
  
  # Validations
  validates :email, presence: true, 
                   uniqueness: { case_sensitive: false },
                   format: { with: URI::MailTo::EMAIL_REGEXP }
  validates :username, presence: true,
                      uniqueness: true,
                      length: { minimum: 3, maximum: 20 },
                      format: { with: /\A[a-zA-Z0-9_]+\z/ }
  validates :password, length: { minimum: 6 }, if: :password_required?
  
  # Callbacks
  before_validation :normalize_email
  before_save :normalize_username
  
  # Secure password
  has_secure_password
  
  private
  
  def normalize_email
    self.email = email.downcase.strip if email.present?
  end
  
  def normalize_username
    self.username = username.downcase.strip if username.present?
  end
  
  def password_required?
    new_record? || password.present?
  end
end
```

### 2. Post Model

```ruby
# app/models/post.rb
class Post < ApplicationRecord
  # Associations
  belongs_to :user
  has_many :comments, dependent: :destroy
  has_many :post_tags, dependent: :destroy
  has_many :tags, through: :post_tags
  
  # Validations
  validates :title, presence: true,
                   length: { minimum: 3, maximum: 100 }
  validates :content, presence: true,
                     length: { minimum: 10 }
  
  # Callbacks
  before_validation :generate_slug
  after_create :notify_subscribers
  
  # Scopes
  scope :published, -> { where(published: true) }
  scope :draft, -> { where(published: false) }
  scope :recent, -> { order(created_at: :desc) }
  scope :popular, -> { where('views_count > ?', 100) }
  
  # Instance methods
  def publish!
    update(published: true, published_at: Time.current)
  end
  
  def unpublish!
    update(published: false, published_at: nil)
  end
  
  def excerpt
    content.truncate(100)
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

### 3. Comment Model

```ruby
# app/models/comment.rb
class Comment < ApplicationRecord
  # Associations
  belongs_to :user
  belongs_to :post
  belongs_to :parent, class_name: 'Comment', optional: true
  has_many :replies, class_name: 'Comment', foreign_key: :parent_id
  
  # Validations
  validates :content, presence: true,
                     length: { minimum: 2, maximum: 500 }
  
  # Callbacks
  after_create :notify_post_author
  
  # Scopes
  scope :recent, -> { order(created_at: :desc) }
  scope :top_level, -> { where(parent_id: nil) }
  
  private
  
  def notify_post_author
    CommentNotificationJob.perform_later(self)
  end
end
```

### 4. Tag Model

```ruby
# app/models/tag.rb
class Tag < ApplicationRecord
  # Associations
  has_many :post_tags, dependent: :destroy
  has_many :posts, through: :post_tags
  
  # Validations
  validates :name, presence: true,
                  uniqueness: { case_sensitive: false },
                  length: { minimum: 2, maximum: 20 }
  
  # Callbacks
  before_validation :normalize_name
  
  # Scopes
  scope :popular, -> { 
    joins(:posts)
      .group('tags.id')
      .order('COUNT(posts.id) DESC')
  }
  
  private
  
  def normalize_name
    self.name = name.downcase.strip if name.present?
  end
end
```

### 5. PostTag Model

```ruby
# app/models/post_tag.rb
class PostTag < ApplicationRecord
  # Associations
  belongs_to :post
  belongs_to :tag
  
  # Validations
  validates :post_id, uniqueness: { scope: :tag_id }
end
```

## Testing

```ruby
# test/models/user_test.rb
require 'test_helper'

class UserTest < ActiveSupport::TestCase
  setup do
    @user = users(:one)
  end
  
  test "should be valid" do
    assert @user.valid?
  end
  
  test "email should be present" do
    @user.email = nil
    assert_not @user.valid?
  end
  
  test "email should be unique" do
    duplicate_user = @user.dup
    duplicate_user.email = @user.email.upcase
    assert_not duplicate_user.valid?
  end
  
  test "username should be present" do
    @user.username = nil
    assert_not @user.valid?
  end
  
  test "password should be present (on create)" do
    @user = User.new
    @user.password = nil
    assert_not @user.valid?
  end
end

# test/models/post_test.rb
require 'test_helper'

class PostTest < ActiveSupport::TestCase
  setup do
    @post = posts(:one)
  end
  
  test "should be valid" do
    assert @post.valid?
  end
  
  test "title should be present" do
    @post.title = nil
    assert_not @post.valid?
  end
  
  test "content should be present" do
    @post.content = nil
    assert_not @post.valid?
  end
  
  test "should generate slug from title" do
    @post.title = "New Post Title"
    @post.valid?
    assert_equal "new-post-title", @post.slug
  end
  
  test "should publish post" do
    @post.publish!
    assert @post.published?
    assert_not_nil @post.published_at
  end
end
```

## Key Features Demonstrated

1. **Model Associations**: One-to-many, many-to-many relationships
2. **Validations**: Presence, uniqueness, format, length
3. **Callbacks**: Before/after validation, create, save
4. **Scopes**: Common query patterns
5. **Instance Methods**: Business logic
6. **Testing**: Model validations and methods

## Next Steps

1. Add more features:
   - User authentication
   - Post categories
   - Post search
   - Post analytics
   - User roles and permissions

2. Enhance models:
   - Add more validations
   - Implement soft deletes
   - Add model concerns
   - Add model caching

3. Improve testing:
   - Add more test cases
   - Add integration tests
   - Add performance tests 