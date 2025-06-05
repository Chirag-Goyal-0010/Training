# Database Queries in Ruby

## Basic Database Setup
```ruby
require 'sqlite3'
require 'active_record'

# Configure ActiveRecord
ActiveRecord::Base.establish_connection(
  adapter: 'sqlite3',
  database: 'example.db'
)

# Define models
class User < ActiveRecord::Base
  has_many :posts
end

class Post < ActiveRecord::Base
  belongs_to :user
end
```

## Basic Queries
```ruby
# Find all records
users = User.all

# Find by ID
user = User.find(1)

# Find by conditions
users = User.where(age: 25)
users = User.where("age > ?", 25)

# Find first/last
first_user = User.first
last_user = User.last

# Find or create
user = User.find_or_create_by(name: "John")
```

## Advanced Queries
```ruby
# Joins
users_with_posts = User.joins(:posts)

# Includes (eager loading)
users = User.includes(:posts)

# Group and count
post_counts = Post.group(:user_id).count

# Order
users = User.order(age: :desc)
users = User.order("age DESC, name ASC")

# Limit and offset
users = User.limit(10).offset(20)
```

## Scopes
```ruby
class User < ActiveRecord::Base
  scope :active, -> { where(active: true) }
  scope :adults, -> { where("age >= ?", 18) }
  scope :recent, -> { order(created_at: :desc) }
  
  # Dynamic scope
  scope :age_between, ->(min, max) { where(age: min..max) }
end

# Usage
active_adults = User.active.adults
recent_users = User.recent.limit(5)
```

## Aggregations
```ruby
# Count
total_users = User.count
active_users = User.where(active: true).count

# Average
avg_age = User.average(:age)

# Sum
total_posts = Post.sum(:views)

# Maximum/Minimum
max_age = User.maximum(:age)
min_age = User.minimum(:age)
```

## Complex Queries
```ruby
# Subqueries
popular_posts = Post.where("views > ?", Post.average(:views))

# Having
user_post_counts = User.joins(:posts)
                      .group(:id)
                      .having("COUNT(posts.id) > ?", 5)

# Union
all_users = User.where(active: true)
               .union(User.where("age > ?", 18))
```

## Transactions
```ruby
ActiveRecord::Base.transaction do
  user = User.create!(name: "John")
  user.posts.create!(title: "First Post")
end
```

## Batch Processing
```ruby
# Find in batches
User.find_in_batches(batch_size: 1000) do |users|
  users.each do |user|
    # Process each user
  end
end

# Find each
User.find_each do |user|
  # Process each user
end
```

## Query Optimization
```ruby
# Select specific columns
users = User.select(:id, :name)

# Use indexes
add_index :users, :email, unique: true
add_index :posts, [:user_id, :created_at]

# Use counter cache
belongs_to :user, counter_cache: true
```

## Raw SQL
```ruby
# Execute raw SQL
results = ActiveRecord::Base.connection.execute("
  SELECT users.*, COUNT(posts.id) as post_count
  FROM users
  LEFT JOIN posts ON users.id = posts.user_id
  GROUP BY users.id
")

# Find by SQL
users = User.find_by_sql("
  SELECT * FROM users
  WHERE age > 18
  ORDER BY created_at DESC
")
```

## Best Practices
1. Use appropriate indexes
2. Avoid N+1 queries
3. Use eager loading when needed
4. Write efficient scopes
5. Use transactions for data integrity
6. Batch process large datasets
7. Monitor query performance
8. Use appropriate data types
9. Validate data before saving
10. Handle errors gracefully 