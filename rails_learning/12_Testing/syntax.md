# Unit Testing (Models) Syntax

This document provides syntax examples for writing unit tests for your Rails models using the default MiniTest framework.

Model tests are typically located in `test/models/` and inherit from `ActiveSupport::TestCase`.

## Basic Model Test Structure

```ruby
# test/models/article_test.rb

require "test_helper"

class ArticleTest < ActiveSupport::TestCase
  # Test methods start with 'test_'

  test "should be valid" do
    article = Article.new(title: "Test Article", body: "This is a test article body.")
    assert article.valid?
  end

  test "should require a title" do
    article = Article.new(body: "Some body")
    assert_not article.valid?
    assert_includes article.errors[:title], "can't be blank"
  end

  # Add more test methods here
end
```

## Testing Validations

Test each validation rule defined in your model.

```ruby
# test/models/user_test.rb

require "test_helper"

class UserTest < ActiveSupport::TestCase
  test "email should be unique" do
    User.create!(email: "test@example.com", password: "password")
    duplicate_user = User.new(email: "test@example.com", password: "another_password")
    assert_not duplicate_user.valid?
    assert_includes duplicate_user.errors[:email], "has already been taken"
  end

  test "password should have a minimum length" do
    user = User.new(email: "test2@example.com", password: "short")
    assert_not user.valid?
    assert_includes user.errors[:password], "is too short (minimum is 6 characters)"
  end
end
```

## Testing Custom Methods

Test any custom methods you've added to your model.\n
```ruby
# test/models/article_test.rb (continued)

  test "word_count returns correct count" do
    article = Article.new(body: "This body has five words.")
    assert_equal 5, article.word_count

    article_no_body = Article.new(body: nil)
    assert_nil article_no_body.word_count
  end

  test "published? returns true for published articles" do
    published_article = Article.new(status: "published")
    assert published_article.published?

    draft_article = Article.new(status: "draft")
    assert_not draft_article.published?
  end
```

## Testing Associations

Test that associations are correctly defined.

```ruby
# test/models/article_test.rb (continued)

  test "should belong to a user" do
    article = Article.new(title: "Belongs Test", body: "Body", user: users(:one)) # Using a fixture
    assert article.valid?
    assert_instance_of User, article.user
  end

# test/models/user_test.rb (continued)

  test "should have many articles" do
    user = users(:one) # Using a fixture
    assert_instance_of Article, user.articles.first # Check if the collection contains Article instances
    assert_equal 2, user.articles.count # Assuming fixture user has 2 articles
  end
```

## Using Fixtures

Fixtures are a way to populate your test database with sample data. They are YAML files located in `test/fixtures/`.

```yaml
# test/fixtures/users.yml

one:
  name: John Doe
  email: john@example.com

two:
  name: Jane Smith
  email: jane@example.com
```

In your tests, you can access these fixtures as methods with the same name as the fixture file (singularized) and the fixture name:

```ruby
# test/models/article_test.rb

require "test_helper"

class ArticleTest < ActiveSupport::TestCase
  # Load fixtures (implicitly done if fixtures are in test/fixtures)
  # fixtures :users

  test "article with user fixture is valid" do
    article = Article.new(title: "Fixture Test", body: "Body", user: users(:one))
    assert article.valid?
  end
end
```
You can also use Factories (like FactoryBot) for creating test data, which offer more flexibility than fixtures, but fixtures are the default in Rails.

Unit tests are crucial for verifying the core logic of your models in isolation. The next topic will cover Integration Testing (Controllers). 

# Fixtures and Factories Syntax

This document provides syntax examples for defining and using test data with Rails fixtures and FactoryBot factories.

## Fixtures Syntax

Fixtures are defined in YAML files in the `test/fixtures/` directory. The filename (pluralized) should match the model name.

```yaml
# test/fixtures/users.yml

one:
  name: John Doe
  email: john@example.com
  password_digest: <%= BCrypt::Password.create('password') %> # Using ERB to generate dynamic data

two:
  name: Jane Smith
  email: jane@example.com
  password_digest: <%= BCrypt::Password.create('secure_password') %>

# Example with an association
three:
  name: Bob Johnson
  email: bob@example.com
  password_digest: <%= BCrypt::Password.create('another_password') %>
```

```yaml
# test/fixtures/articles.yml

article_one:
  title: My First Fixture Article
  body: This is the body for article one.
  user: one # Refers to the user fixture named 'one' in users.yml
  created_at: <%= Time.current %>
  status: published

article_two:
  title: Second Fixture Article
  body: This is the body for article two.
  user: one # Both articles belong to the same user fixture
  created_at: <%= 1.day.ago %>
  status: draft

article_three:
  title: Another User's Article
  body: Body content.
  user: three # Belongs to the user fixture named 'three'
  created_at: <%= Time.current %>
  status: published
```

In your tests, you can access these fixtures by their filename (singularized) and fixture name:

```ruby
# test/models/article_test.rb

require "test_helper"

class ArticleTest < ActiveSupport::TestCase
  # Fixtures are loaded by default in ActiveSupport::TestCase

  test "should find fixtures" do
    assert_equal "My First Fixture Article", articles(:article_one).title
    assert_equal users(:one).email, articles(:article_one).user.email
  end

  test "should count fixtures" do
    assert_equal 3, Article.count # Assumes only these 3 fixtures exist initially
  end
end
```

## FactoryBot Syntax

Factories are typically defined in files in `test/factories/`.

First, add the `factory_bot_rails` gem to your `Gemfile` and run `bundle install`.

```ruby
# Gemfile

group :development, :test do
  gem 'factory_bot_rails'
  gem 'faker', '~> 2.0' # Optional: for generating fake data
end
```

Then, define your factories:

```ruby
# test/factories/users.rb

FactoryBot.define do
  factory :user do
    # Use sequence for unique attributes like email
    sequence(:email) { |n| "user_#{n}@example.com" }
    name { Faker::Name.name } # Using the Faker gem
    password { "password" }

    # Define traits for variations
    trait :admin do
      admin { true }
    end
  end
end
```

```ruby
# test/factories/articles.rb

FactoryBot.define do
  factory :article do
    title { Faker::Lorem.sentence }
    body { Faker::Lorem.paragraph }
    # Define associations
    association :user # Creates an associated user using the :user factory

    # Define traits for status
    trait :published do
      status { "published" }
    end

    trait :draft do
      status { "draft" }
    end
  end
end
```

In your tests, use FactoryBot methods to create data:

```ruby
# test/models/article_test.rb (using FactoryBot)

require "test_helper"
# require "factory_bot_rails" # Not always needed if loaded globally

class ArticleTest < ActiveSupport::TestCase
  # Optional: disable fixtures if using only factories
  # self.use_fixtures = false

  test "factory creates a valid article" do
    article = FactoryBot.create(:article) # Creates and saves to DB
    assert article.valid?
    assert_instance_of Article, article
    assert_not_nil article.user # Ensure associated user was created
  end

  test "build creates an article object but doesn't save" do
    article = FactoryBot.build(:article) # Creates in memory
    assert_not article.persisted?
  end

  test "published trait creates a published article" do
    published_article = FactoryBot.create(:article, :published)
    assert published_article.published?
  end

  test "creating multiple articles with same user" do
    user = FactoryBot.create(:user)
    article1 = FactoryBot.create(:article, user: user)
    article2 = FactoryBot.create(:article, user: user)

    assert_equal user, article1.user
    assert_equal user, article2.user
    assert_equal 2, user.articles.count
  end
end
```

Both fixtures and factories have their place. Fixtures are good for a baseline dataset, while factories offer more flexibility for specific test scenarios. The final topic in this section covers how to run your tests. 