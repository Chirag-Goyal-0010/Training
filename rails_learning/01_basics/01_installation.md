# Rails Installation and Setup

## Prerequisites
- Ruby (version 3.0.0 or higher)
- Node.js and Yarn
- SQLite3
- Git

## Installation Steps

1. Install Ruby
```bash
# Using rbenv
rbenv install 3.2.2
rbenv global 3.2.2

# Using RVM
rvm install 3.2.2
rvm use 3.2.2
```

2. Install Rails
```bash
gem install rails
```

3. Verify Installation
```bash
rails --version
```

## Creating Your First Rails Application

```bash
# Create a new Rails application
rails new my_first_app

# Navigate to the application directory
cd my_first_app

# Start the Rails server
rails server
```

Visit http://localhost:3000 to see your application running.

## Project Structure

```
my_first_app/
├── app/                    # Contains controllers, models, views, etc.
├── bin/                    # Contains the rails script
├── config/                 # Configuration files
├── db/                     # Database files
├── lib/                    # Library modules
├── log/                    # Log files
├── public/                 # Static files
├── test/                   # Test files
├── tmp/                    # Temporary files
├── vendor/                 # Third-party code
├── Gemfile                 # Ruby dependencies
└── README.md              # Project documentation
```

## Next Steps
- Learn about the MVC architecture
- Understand Rails routing
- Create your first controller and view 