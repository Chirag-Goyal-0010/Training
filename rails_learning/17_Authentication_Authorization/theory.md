# Authentication & Authorization Theory

## Authentication vs Authorization

### Authentication
- Verifies user identity
- Confirms "who you are"
- Login process
- Session management
- Password handling

### Authorization
- Controls access to resources
- Determines "what you can do"
- Permission management
- Role-based access
- Resource-level control

## Session Management

### Sessions
- Server-side storage
- Session ID in cookie
- Secure session handling
- Session expiration
- Session hijacking prevention

### Cookies
- Client-side storage
- Secure cookie options
- HTTP-only cookies
- Same-site policy
- Cookie encryption

## Authentication Methods

### 1. Devise
- Complete authentication solution
- User registration
- Password management
- Email confirmation
- Account locking
- OAuth integration

### 2. JWT (JSON Web Tokens)
- Stateless authentication
- Token-based security
- Token structure
- Token validation
- Token expiration
- Refresh tokens

### 3. OAuth
- Third-party authentication
- OAuth providers
- OAuth flow
- Access tokens
- Refresh tokens
- Security considerations

## Authorization Methods

### 1. Pundit
- Policy-based authorization
- Resource-level control
- Policy objects
- Scope objects
- Error handling
- Testing policies

### 2. CanCanCan
- Ability-based authorization
- Role management
- Resource authorization
- Action authorization
- Nested resources
- Testing abilities

## Security Best Practices

### 1. Password Security
- Secure password storage
- Password hashing
- Salt usage
- Password complexity
- Password reset
- Account recovery

### 2. Session Security
- Secure session storage
- Session timeout
- Session fixation prevention
- CSRF protection
- XSS prevention
- Secure cookies

### 3. Token Security
- Secure token storage
- Token expiration
- Token rotation
- Token validation
- Token revocation
- Refresh token security

## Implementation Patterns

### 1. User Model
```ruby
class User < ApplicationRecord
  # Authentication
  has_secure_password
  has_many :sessions
  
  # Authorization
  has_many :roles
  has_many :permissions, through: :roles
  
  # Methods
  def admin?
    roles.include?('admin')
  end
  
  def can?(action, resource)
    permissions.exists?(action: action, resource: resource)
  end
end
```

### 2. Session Management
```ruby
class Session < ApplicationRecord
  belongs_to :user
  
  # Session security
  before_create :generate_token
  before_create :set_expiration
  
  private
  
  def generate_token
    self.token = SecureRandom.hex(32)
  end
  
  def set_expiration
    self.expires_at = 24.hours.from_now
  end
end
```

### 3. Policy Objects
```ruby
class ArticlePolicy < ApplicationPolicy
  def index?
    true
  end
  
  def show?
    true
  end
  
  def create?
    user.present?
  end
  
  def update?
    user.present? && (record.user == user || user.admin?)
  end
  
  def destroy?
    user.present? && (record.user == user || user.admin?)
  end
end
```

## Testing Authentication

### 1. Controller Tests
```ruby
class ArticlesControllerTest < ActionDispatch::IntegrationTest
  setup do
    @user = users(:one)
    @article = articles(:one)
  end
  
  test "should get index when authenticated" do
    sign_in @user
    get articles_url
    assert_response :success
  end
  
  test "should not get index when not authenticated" do
    get articles_url
    assert_redirected_to login_url
  end
end
```

### 2. Policy Tests
```ruby
class ArticlePolicyTest < ActiveSupport::TestCase
  setup do
    @user = users(:one)
    @article = articles(:one)
    @policy = ArticlePolicy.new(@user, @article)
  end
  
  test "allows admin to update any article" do
    @user.update(role: 'admin')
    assert @policy.update?
  end
  
  test "allows author to update own article" do
    @article.update(user: @user)
    assert @policy.update?
  end
end
```

## Error Handling

### 1. Authentication Errors
- Invalid credentials
- Account locked
- Session expired
- Token invalid
- OAuth errors

### 2. Authorization Errors
- Insufficient permissions
- Resource not found
- Action not allowed
- Role restrictions
- Policy violations

## Security Headers

### 1. HTTP Headers
- X-Frame-Options
- X-XSS-Protection
- X-Content-Type-Options
- Content-Security-Policy
- Strict-Transport-Security

### 2. Cookie Headers
- Secure flag
- HttpOnly flag
- SameSite attribute
- Domain attribute
- Path attribute

## Monitoring and Logging

### 1. Authentication Logs
- Login attempts
- Failed logins
- Password resets
- Account locks
- Session creation

### 2. Authorization Logs
- Permission checks
- Policy violations
- Role changes
- Access attempts
- Resource access

## Best Practices

### 1. Authentication
- Use secure password storage
- Implement rate limiting
- Use HTTPS
- Implement MFA
- Regular security audits

### 2. Authorization
- Principle of least privilege
- Regular permission reviews
- Audit logging
- Role-based access
- Resource-level control

### 3. Session Management
- Secure session storage
- Session timeout
- Session rotation
- Secure cookies
- CSRF protection

### 4. Token Management
- Secure token storage
- Token expiration
- Token rotation
- Token validation
- Refresh token security 