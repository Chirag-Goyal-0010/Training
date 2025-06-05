# Migration Syntax

This document provides common syntax examples for writing Rails migrations.

## Creating a Table

Use the `create_table` method within the `change` method.

```ruby
class CreateArticles < ActiveRecord::Migration[7.1]
  def change
    create_table :articles do |t|
      t.string :title
      t.text :body
      t.references :author, null: false, foreign_key: true

      t.timestamps
    end
  end
end
```

Common column types include `string`, `text`, `integer`, `float`, `decimal`, `datetime`, `time`, `date`, `binary`, `boolean`.

## Adding a Column

Use the `add_column` method.

```ruby
class AddStatusToArticles < ActiveRecord::Migration[7.1]
  def change
    add_column :articles, :status, :string, default: "draft"
  end
end
```

## Removing a Column

Use the `remove_column` method.

```ruby
class RemoveStatusFromArticles < ActiveRecord::Migration[7.1]
  def change
    remove_column :articles, :status
  end
end
```

## Changing a Column

Use the `change_column` method.

```ruby
class ChangeArticleBodyToText < ActiveRecord::Migration[7.1]
  def change
    change_column :articles, :body, :string, limit: 500
  end
end
```

Alternatively, use `change_column_null` or `change_column_default`.

## Renaming a Column

Use the `rename_column` method.

```ruby
class RenameArticleAuthorId < ActiveRecord::Migration[7.1]
  def change
    rename_column :articles, :author_id, :user_id
  end
end
```

## Adding an Index

Use the `add_index` method.

```ruby
class AddIndexToArticleTitle < ActiveRecord::Migration[7.1]
  def change
    add_index :articles, :title
  end
end
```

For unique indexes:

```ruby
class AddUniqueIndexToArticleSlug < ActiveRecord::Migration[7.1]
  def change
    add_index :articles, :slug, unique: true
  end
end
```

## Removing an Index

Use the `remove_index` method.

```ruby
class RemoveIndexFromArticleTitle < ActiveRecord::Migration[7.1]
  def change
    remove_index :articles, :title
  end
end
```

## Creating a Join Table (for Many-to-Many)

Use `create_join_table`.

```ruby
class CreateArticlesTagsJoinTable < ActiveRecord::Migration[7.1]
  def change
    create_join_table :articles, :tags
  end
end
```

This covers the basic syntax for common migration tasks. You can find more advanced options in the Rails Guides documentation on Active Record Migrations.

# Defining Models Syntax

This document provides common syntax examples for defining Rails models and their components.

## Basic Model Definition

```ruby
# app/models/article.rb
class Article < ApplicationRecord
  # Associations, validations, callbacks, and custom methods go here
end
```

## Defining Associations

Use methods like `belongs_to`, `has_one`, `has_many`, and `has_and_belongs_to_many`.

```ruby
class User < ApplicationRecord
  has_many :articles # A user can have many articles
  has_one :profile   # A user can have one profile
end

class Article < ApplicationRecord
  belongs_to :user # An article belongs to a user
  has_many :comments # An article can have many comments
  has_and_belongs_to_many :tags # An article can have many tags, and a tag can belong to many articles
end

class Profile < ApplicationRecord
  belongs_to :user # A profile belongs to a user
end

class Comment < ApplicationRecord
  belongs_to :article # A comment belongs to an article
end

class Tag < ApplicationRecord
  has_and_belongs_to_many :articles # A tag can belong to many articles
end
```

## Defining Validations

Use validation helper methods within the model class.

```ruby
class Article < ApplicationRecord
  validates :title, presence: true, length: { minimum: 5 }
  validates :body, presence: true
  validates :user_id, presence: true
  validates :slug, uniqueness: true, allow_nil: true
end
```

Common validations include `presence`, `uniqueness`, `length`, `numericality`, `format`, `inclusion`, `exclusion`.

## Defining Callbacks

Use callback methods to execute code at specific points in the model's lifecycle (e.g., `before_save`, `after_create`, `before_destroy`).

```ruby
class Article < ApplicationRecord
  before_save :set_published_at
  after_destroy :log_deletion

  private

  def set_published_at
    if status == "published" && published_at.nil?
      self.published_at = Time.current
    end
  end

  def log_deletion
    Rails.logger.info "Article #{title} (ID: #{id}) was deleted."
  end
end
```

## Defining Custom Methods

Add your own methods to encapsulate business logic related to the model.

```ruby
class Article < ApplicationRecord
  def published?
    status == "published"
  end

  def word_count
    body.split.count if body.present?
  end
end
```

This covers basic syntax for defining Rails models, associations, validations, and callbacks. The next step would be to look at CRUD operations and how to perform them using Active Record.

# CRUD Operations Syntax

This document provides common syntax examples for performing CRUD (Create, Read, Update, Delete) operations using Active Record.

## Create

```ruby
# 1. Using new and save
article = Article.new(title: "New Article Title", body: "Content goes here")
if article.save
  puts "Article created successfully!"
else
  puts "Failed to create article: #{article.errors.full_messages.join(', ')}"
end

# 2. Using create
article = Article.create(title: "Another Article", body: "More content")
if article.persisted?
  puts "Article created successfully with create!"
else
  puts "Failed to create article with create: #{article.errors.full_messages.join(', ')}"
end

# 3. Using create! (raises exception on failure)
begin
  article = Article.create!(title: "Guaranteed Article", body: "Content.")
  puts "Article created successfully with create!!"
rescue ActiveRecord::RecordInvalid => e
  puts "Failed to create article with create!: #{e.message}"
end
```

## Read

```ruby
# 1. Get all articles
all_articles = Article.all
puts "Total articles: #{all_articles.count}"

# 2. Find an article by ID
article = Article.find(1) # Raises ActiveRecord::RecordNotFound if not found
puts "Found article: #{article.title}"

# 3. Find the first article matching a condition
article = Article.find_by(title: "New Article Title")
if article
  puts "Found article by title: #{article.title}"
else
  puts "Article not found by title."
end

# 4. Get all articles matching a condition
published_articles = Article.where(status: "published")
puts "Published articles count: #{published_articles.count}"

# 5. Chaining queries
recent_published_articles = Article.where(status: "published").order(created_at: :desc).limit(5)
puts "Recent published articles count: #{recent_published_articles.count}"

# 6. Checking for existence
exists = Article.exists?(title: "New Article Title")
puts "Article with title 'New Article Title' exists: #{exists}"
```

## Update

```ruby
# Assume an article with ID 1 exists
article = Article.find(1)

# 1. Modify attributes and save
article.body = "Updated content for the article."
if article.save
  puts "Article updated successfully with save!"
else
  puts "Failed to update article with save: #{article.errors.full_messages.join(', ')}"
end

# 2. Using update
if article.update(title: "Revised Article Title", status: "published")
  puts "Article updated successfully with update!"
else
  puts "Failed to update article with update: #{article.errors.full_messages.join(', ')}"
end

# 3. Using update! (raises exception on failure)
begin
  article.update!(body: "Final content.")
  puts "Article updated successfully with update!!"
rescue ActiveRecord::RecordInvalid => e
  puts "Failed to update article with update!: #{e.message}"
end

# 4. Update multiple records
Article.where(status: "draft").update_all(status: "archived", updated_at: Time.current)
puts "Updated draft articles to archived."
```

## Delete

```ruby
# Assume an article with ID 1 exists
article = Article.find_by(id: 1)

if article
  # 1. Using destroy (runs callbacks)
  article.destroy
  puts "Article destroyed successfully (callbacks ran)."

  # Re-create for the next example (for demonstration)
  article = Article.create!(title: "To Be Deleted", body: "Ephemeral content.")

  # 2. Using delete (does NOT run callbacks)
  article.delete
  puts "Article deleted successfully (no callbacks ran)."
end

# 3. Destroy multiple records (runs callbacks for each)
# Assuming you have some archived articles
Article.where(status: "archived").destroy_all
puts "Destroyed all archived articles (callbacks ran)."

# 4. Delete multiple records (does NOT run callbacks)
# Assuming you have some other articles to delete
Article.where("created_at < ?", 1.year.ago).delete_all
puts "Deleted old articles (no callbacks ran)."
```

These examples demonstrate the basic syntax for performing CRUD operations using Active Record. Understanding these methods is fundamental to interacting with your database in Rails.

# Associations Syntax

This document provides syntax examples for defining and using Active Record associations.

## `belongs_to`

Used in the child model to indicate a one-to-one or one-to-many relationship where the child contains the foreign key.

```ruby
# app/models/article.rb
class Article < ApplicationRecord
  belongs_to :user # Assumes a user_id column in the articles table

  # Optional: Add validation for presence
  validates :user, presence: true
end

# Example usage:
article = Article.find(1)
user = article.user # Accessing the associated user object
```

## `has_one`

Used in the parent model to indicate a one-to-one relationship where the child model contains the foreign key.

```ruby
# app/models/user.rb
class User < ApplicationRecord
  has_one :profile # Assumes a user_id column in the profiles table

  # Optional: Dependent option to destroy the profile when the user is destroyed
  has_one :profile, dependent: :destroy
end

# Example usage:
user = User.find(1)
profile = user.profile # Accessing the associated profile object
```

## `has_many`

Used in the parent model to indicate a one-to-many relationship where the child model contains the foreign key.

```ruby
# app/models/user.rb
class User < ApplicationRecord
  has_many :articles # Assumes a user_id column in the articles table

  # Optional: Dependent option to destroy articles when the user is destroyed
  has_many :articles, dependent: :destroy
end

# Example usage:
user = User.find(1)
articles = user.articles # Accessing the collection of associated articles

# Creating a new associated object:
new_article = user.articles.create(title: "New Article", body: "Content.")
```

## `has_and_belongs_to_many` (HABTM)

Used to set up a many-to-many connection with another model. Requires a separate join table (e.g., `articles_tags`) with foreign keys for both models (e.g., `article_id` and `tag_id`). No model is needed for the join table itself.

```ruby
# app/models/article.rb
class Article < ApplicationRecord
  has_and_belongs_to_many :tags # Assumes an articles_tags join table
end

# app/models/tag.rb
class Tag < ApplicationRecord
  has_and_belongs_to_many :articles # Assumes an articles_tags join table
end

# Example usage:
article = Article.find(1)
tags = article.tags # Accessing the collection of associated tags

tag = Tag.find(1)
articles_with_tag = tag.articles # Accessing the collection of associated articles

# Adding/removing associations:
article.tags << tag # Add a tag to an article
article.tags.delete(tag) # Remove a tag from an article
```

## `has_many :through`

Used to set up a many-to-many relationship through an intermediate model. This is generally preferred over `has_and_belongs_to_many` as it allows for adding attributes and validations to the join table.

```ruby
# app/models/assembly.rb
class Assembly < ApplicationRecord
  has_many :assemblies_parts # Assumes an assemblies_parts join table model
  has_many :parts, through: :assemblies_parts # Through the AssembliesPart model
end

# app/models/part.rb
class Part < ApplicationRecord
  has_many :assemblies_parts # Assumes an assemblies_parts join table model
  has_many :assemblies, through: :assemblies_parts # Through the AssembliesPart model
end

# app/models/assemblies_part.rb
class AssembliesPart < ApplicationRecord
  belongs_to :assembly
  belongs_to :part
  # You can add attributes/validations here, e.g., quantity
  validates :quantity, presence: true
end
```

Understanding associations is crucial for modeling relationships in your database and working efficiently with related data in Rails. The next topic will cover Validations.

# Validations Syntax

This document provides common syntax examples for defining validations in Rails models.

## Basic Syntax

The `validates` helper is the most common way to define validations.

```ruby
class Article < ApplicationRecord
  validates :title, presence: true
  validates :body, length: { minimum: 10 }
end
```

You can apply multiple validations to a single attribute:

```ruby
class User < ApplicationRecord
  validates :email, presence: true, uniqueness: true, format: { with: URI::MailTo::EMAIL_REGEXP }
end
```

## Common Validation Helpers with Examples

### `presence`

Ensures the attribute is not nil or empty.

```ruby
validates :name, presence: true
validates :address, presence: true, message: "must be provided"
```

### `uniqueness`

Checks if the attribute value is unique within the table.

```ruby
validates :email, uniqueness: true

# Case-insensitive uniqueness
validates :username, uniqueness: { case_sensitive: false }

# Scoped uniqueness (e.g., a product name must be unique within a category)
validates :name, uniqueness: { scope: :category_id }
```

### `length`

Validates the length of a string or array.

```ruby
validates :bio, length: { maximum: 500 }
validates :password, length: { minimum: 8, maximum: 20 }
validates :tags, length: { is: 3 } # Exactly 3 tags
validates :description, length: { in: 50..500 } # Length within a range
validates :comment, length: { minimum: 10, too_short: "must have at least %{count} words" } # Custom error message
```

### `numericality`

Validates that an attribute is a number.

```ruby
validates :price, numericality: true
validates :quantity, numericality: { only_integer: true, greater_than: 0 }
validates :rating, numericality: { less_than_or_equal_to: 5.0 }
```

### `inclusion` and `exclusion`

Validates that an attribute's value is or is not in a given set.

```ruby
validates :size, inclusion: { in: %w(small medium large), message: "%{value} is not a valid size" }
validates :status, exclusion: { in: %w(pending archived), message: "%{value} is reserved." }
```

### `format`

Validates the format using a regular expression.

```ruby
validates :legacy_code, format: { with: /\A[a-zA-Z]+\z/, message: "only allows letters" }
```

### `संबंध` (Association) Validations

Presence validation for `belongs_to` associations is often included automatically or can be added explicitly.

```ruby
class Article < ApplicationRecord
  belongs_to :user
  validates :user, presence: true # Explicit presence validation for the association
end
```

## Custom Validations

Define a method and use the `validate` helper.

```ruby
class Product < ApplicationRecord
  validate :expiration_date_cannot_be_in_the_past

  def expiration_date_cannot_be_in_the_past
    if expiration_date.present? && expiration_date < Date.current
      errors.add(:expiration_date, "can't be in the past")
    end
  end
end
```

Validations are essential for data integrity. Use these examples to enforce rules on your model attributes. The next topic in Active Record is Callbacks.

# Callbacks Syntax

This document provides common syntax examples for defining and using Active Record callbacks.

## Basic Callback Definition

Callbacks are defined using class methods corresponding to the callback name.

```ruby
class Article < ApplicationRecord
  before_save :normalize_title
  after_create :send_new_article_notification
  before_destroy :ensure_not_referenced

  private

  def normalize_title
    self.title = title.strip.downcase.capitalize if title.present?
  end

  def send_new_article_notification
    # Code to send an email or notification after creating the article
    puts "Notification sent for new article: #{title}"
  end

  def ensure_not_referenced
    if comments.any?
      errors.add(:base, "Cannot delete article with comments.")
      throw :abort # Halt the destruction process
    end
  end
end
```

## Using a Block for a Callback

You can define a simple callback using a block directly.

```ruby
class Comment < ApplicationRecord
  after_save do |comment|
    Rails.logger.info "Comment on article #{comment.article_id} saved."
  end
end
```

## Using a Proc or Lambda

Callbacks can also be defined using Procs or lambdas.

```ruby
class User < ApplicationRecord
  before_validation Proc.new { |user| user.email = user.email.downcase if user.email.present? }
end
```

## Conditional Callbacks

Use `:if`, `:unless`, or `:on` options to make callbacks conditional.

### `:if` and `:unless`

Run the callback only if the condition is true (`:if`) or false (`:unless`). The condition can be a symbol pointing to a method, a Proc, or a string to be evaluated.

```ruby
class Article < ApplicationRecord
  before_save :set_published_at, if: :published_status_changed?
  after_commit :send_notification, if: Proc.new { |article| article.status == "published" }
  before_destroy :archive_article, unless: :admin?

  def published_status_changed?
    status_changed? && status == "published"
  end

  def admin?
    # Method to check if the current user is an admin
    false # Placeholder
  end
end
```

### `:on`

Specify when the callback should run (e.g., `:create`, `:update`, `:save`, `:destroy`, `:commit`).

```ruby
class Product < ApplicationRecord
  after_commit :update_inventory, on: [:create, :destroy]
  after_rollback :log_rollback_error, on: :update
end
```

Callbacks are powerful tools for adding logic to your model's life cycle. Use them carefully to keep your application's flow clear. The next topic in Active Record is Querying the Database.

# Querying the Database Syntax

This document provides common syntax examples for querying the database using Active Record.

## Retrieving Single Objects

```ruby
# Find by primary key (id)
article = Article.find(1) # Raises ActiveRecord::RecordNotFound if not found

# Find the first record matching conditions
article = Article.find_by(title: "My Article")
article = Article.find_by(status: "published", user_id: 5)

# Get the first record (useful with order/where)
first_article = Article.order(:created_at).first

# Get the last record (useful with order/where)
last_article = Article.order(:created_at).last
```

## Retrieving Collections of Objects (Relations)

```ruby
# Get all records
all_articles = Article.all

# Get records matching conditions using `where` (hash syntax)
published_articles = Article.where(status: "published")
draft_articles = Article.where(status: ["draft", "pending"])
articles_by_user = Article.where(user_id: 1)

# Get records matching conditions using `where` (string syntax - use with caution against SQL injection)
articles_with_views = Article.where("views > 100")
articles_from_last_week = Article.where("created_at >= ?", 1.week.ago)

# Get records matching conditions using `where` (array syntax)
articles_by_user_and_status = Article.where("user_id = ? AND status = ?", 1, "published")
```

## Ordering Results

```ruby
# Order by a single column
articles_asc = Article.order(:title) # ASC is default
articles_desc = Article.order(created_at: :desc)

# Order by multiple columns
ordered_articles = Article.order(status: :asc, created_at: :desc)
```

## Limiting and Offsetting Results

```ruby
# Limit the number of results
first_ten = Article.limit(10)

# Offset the results (for pagination)
second_page = Article.limit(10).offset(10)
```

## Selecting Specific Columns

```ruby
# Select only specific columns
article_titles = Article.select(:id, :title)
article_titles.each do |article|
  puts article.title # Accessing title is fine
  # puts article.body # This would likely raise an error as body wasn't selected
end
```

## Using `joins` for SQL JOINs

```ruby
# Inner join with the users table and filter by user country
articles_from_usa_users = Article.joins(:user).where(users: { country: "USA" })

# Left outer join
articles_with_users = Article.left_joins(:user)
```

## Using `includes` for Eager Loading

Avoids N+1 queries when accessing associated data.

```ruby
# Eager load the associated user for each article
articles_with_users = Article.includes(:user)
articles_with_users.each do |article|
  puts article.user.name # Doesn't run a new query for each article
end

# Eager load multiple associations
articles_with_user_and_comments = Article.includes(:user, :comments)

# Eager load nested associations
articles_with_user_profile = Article.includes(user: :profile)
```

## Querying Through Associations

```ruby
user = User.find(1)

# Get articles belonging to a specific user
user_articles = user.articles

# Query associated articles
published_user_articles = user.articles.where(status: "published")

# Find users with more than 10 articles
users_with_many_articles = User.joins(:articles).group(:user_id).having("COUNT(articles.id) > 10")
```

These examples cover common database querying patterns in Active Record. Experiment with chaining these methods to build complex queries. The final topics in this section will cover Scopes and Advanced Concepts.

# Scopes Syntax

This document provides common syntax examples for defining and using Active Record scopes.

## Defining Scopes

Use the `scope` class method within your model.

```ruby
class Article < ApplicationRecord
  # Basic scope for published articles
  scope :published, -> { where(status: "published") }

  # Scope with an argument for articles by a specific user
  scope :by_user, ->(user_id) { where(user_id: user_id) }

  # Scope with multiple arguments
  scope :created_between, ->(start_date, end_date) { where(created_at: start_date..end_date) }

  # Using a method for a more complex scope
  scope :popular, -> { where("views > ?", 100).order(views: :desc) }

  # Using a method that returns a relation
  def self.recent(limit = 5)
    order(created_at: :desc).limit(limit)
  end
end
```

Lambdas (`-> {}`) are the preferred way to define scopes, especially if the query involves arguments or needs to be evaluated at the time of calling.

## Using Scopes

Scopes are called on the model class or on an existing `ActiveRecord::Relation`.

```ruby
# Get all published articles
published_articles = Article.published

# Get published articles by a specific user
published_user_articles = Article.published.by_user(1)

# Get popular articles created this month
this_month_popular_articles = Article.popular.created_between(Date.current.beginning_of_month, Date.current.end_of_month)

# Using the method-based scope
recent_articles = Article.recent(10)
```

## Chaining Scopes

Scopes return `ActiveRecord::Relation` objects, allowing you to chain multiple scopes together to build more complex queries.

```ruby
# Get recent published articles by a specific user
query = Article.published.by_user(1).recent(5)

# The query is executed when you access the data:
query.each do |article|
  puts article.title
end
```

## Default Scopes

You can define a default scope that will be applied to all queries on the model unless explicitly unscoped.

```ruby
class Article < ApplicationRecord
  default_scope { order(created_at: :desc) }

  # To remove the default scope for a specific query
  Article.unscoped.all # Get all articles without the default ordering
  Article.unscoped.where(status: "draft")
end
```

Use default scopes with caution as they can sometimes lead to unexpected behavior.

Scopes help keep your model code clean and your queries readable by encapsulating common query patterns. The next topic in Active Record is Advanced Concepts.

# Advanced Active Record Concepts Syntax

This document provides syntax examples for some advanced Active Record concepts.

## Transactions Syntax

Using the `transaction` method with a block is the most common and recommended way to handle transactions.

```ruby
# Example: Transferring money between two accounts
begin
  ActiveRecord::Base.transaction do
    account1 = Account.find(1)
    account2 = Account.find(2)

    account1.withdraw!(50) # Assume withdraw! raises an error if balance is insufficient
    account2.deposit!(50)

    # Both operations succeed, transaction commits
  end
rescue ActiveRecord::RecordInvalid => e
  puts "Transaction failed: #{e.message}. Changes rolled back."
  # Handle insufficient balance or other validation errors
rescue => e
  puts "An unexpected error occurred: #{e.message}. Transaction rolled back."
  # Handle other potential errors during the transaction
end

# You can also use nested transactions, though often discouraged
ActiveRecord::Base.transaction do
  # Outer transaction
  ActiveRecord::Base.transaction do
    # Inner transaction (often behaves like savepoints)
  end
end
```

## Locking Syntax

### Optimistic Locking

Requires a `lock_version` integer column in your table.

```ruby
# app/models/product.rb
class Product < ApplicationRecord
  # No explicit code needed here, just the lock_version column in the database
end

# Example of handling StaleObjectError
begin
  product = Product.find(params[:id])
  product.update!(product_params) # If another process updated it, this raises ActiveRecord::StaleObjectError
rescue ActiveRecord::StaleObjectError
  # Handle the conflict, e.g., reload the record, show an error message, ask user to reapply changes
  redirect_to edit_product_url(product), alert: "Another user has updated this product. Please review and resubmit your changes."
end
```

### Pessimistic Locking

Use the `lock` method within a transaction.

```ruby
# Lock a single record for exclusive access
ActiveRecord::Base.transaction do
  product = Product.lock.find(1)
  # Perform operations on the locked product
  product.stock -= 1
  product.save!
end # Lock is released when transaction commits or rolls back

# Lock multiple records
ActiveRecord::Base.transaction do
  low_stock_products = Product.where("stock < ?", 10).lock
  low_stock_products.each do |product|
    # Update stock, etc.
    product.restock!(50)
  end
end
```

## Explaining Queries Syntax

Use the `explain` method on an `ActiveRecord::Relation`.

```ruby
# Explain a simple query
puts Article.where(status: "published").explain

# Explain a query with includes and ordering
puts Article.includes(:user).where(users: { country: "Canada" }).order(created_at: :desc).explain

# Explain a query with a join and group by
puts User.joins(:articles).group(:user_id).having("COUNT(articles.id) > 10").explain
```
The output will be specific to your database (e.g., PostgreSQL, MySQL, SQLite) and show the query plan, including how indexes are used and the cost of operations.

These examples provide a starting point for using advanced Active Record features. Always refer to the official Rails documentation for the most up-to-date and in-depth information. 