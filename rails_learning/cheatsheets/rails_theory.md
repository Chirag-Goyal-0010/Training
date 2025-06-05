# Ruby on Rails Theory Guide

## 🔹 1. Core Ruby

### Variables and Data Types
- Local Variables: Scoped to the current method or block
- Instance Variables: Scoped to the current object instance
- Class Variables: Shared across all instances of a class
- Global Variables: Accessible throughout the application
- Constants: Values that shouldn't change
- Data Types: String, Integer, Float, Boolean, Array, Hash, Symbol

### Methods and Blocks
- Method Definition: Reusable code blocks that can take parameters
- Default Parameters: Values used when no argument is provided
- Blocks: Anonymous code blocks passed to methods
- Lambdas: Anonymous functions that can be stored in variables

### Classes and Objects
- Classes: Blueprints for creating objects
- Objects: Instances of classes
- Instance Methods: Methods available to instances
- Class Methods: Methods available to the class itself
- Accessors: Methods for getting and setting instance variables

### Modules and Mixins
- Modules: Collections of methods and constants
- Mixins: Including module functionality in classes
- Namespacing: Organizing code and preventing naming conflicts
- Module Methods: Methods available to the module itself

### Inheritance and Polymorphism
- Inheritance: Creating new classes from existing ones
- Superclass: The parent class being inherited from
- Subclass: The class inheriting from a parent
- Method Overriding: Redefining methods in subclasses
- Polymorphism: Using different objects through a common interface

### Exception Handling
- Begin/Rescue: Handling errors gracefully
- Exception Types: Different categories of errors
- Ensure: Code that always runs
- Custom Exceptions: Creating your own error types

### Ruby Gems and Bundler
- Gems: Ruby packages and libraries
- Gemfile: List of project dependencies
- Bundler: Dependency management tool
- Version Constraints: Specifying gem versions

## 🔹 2. Introduction to Rails

### What is Ruby on Rails?
- Web application framework
- MVC architecture
- Convention over Configuration
- DRY (Don't Repeat Yourself) principle

### MVC Architecture
- Model: Data and business logic
- View: User interface
- Controller: Request handling and coordination

### Convention over Configuration
- File naming conventions
- Directory structure
- Database table naming
- RESTful routing

### Rails Directory Structure
- app/: Application code
- config/: Configuration files
- db/: Database files
- lib/: Library modules
- test/: Test files

### Common Rails Commands
- rails new: Create new application
- rails server: Start development server
- rails console: Interactive Ruby console
- rails generate: Create new components
- rails db:migrate: Run database migrations

## 🔹 3. Rails Routing

### Basic Routes
- HTTP Verb Routes: GET, POST, PUT, PATCH, DELETE
- Named Routes: Custom route names
- Route Constraints: Limiting route matches
- Nested Routes: Hierarchical resource relationships

### RESTful Routes
- Resource Routes: Standard CRUD operations
- REST Actions: Index, Show, New, Create, Edit, Update, Destroy
- HTTP Verbs: Matching HTTP methods to actions
- Route Helpers: Generated path and URL helpers

### Custom Routes
- HTTP Verb Routes: Custom endpoints
- Named Routes: Custom route names
- Route Constraints: Limiting route matches
- Namespace Routes: Grouping related routes

### Nested Resources
- Parent-Child Relationships
- Shallow Nesting
- Deep Nesting
- Resource Options

## 🔹 4. Controllers

### Controller Actions
- Index: List resources
- Show: Display single resource
- New: Form for new resource
- Create: Save new resource
- Edit: Form for existing resource
- Update: Save changes to resource
- Destroy: Remove resource

### Filters
- Before Action: Run before specific actions
- After Action: Run after specific actions
- Around Action: Run before and after actions
- Skip Action: Skip specific filters

### Params Handling
- Strong Parameters: Whitelist allowed parameters
- Nested Parameters: Handle complex form data
- Parameter Validation: Ensure required parameters
- Parameter Transformation: Modify parameters

## 🔹 5. Models

### ActiveRecord Models
- Model Definition: Class inheriting from ApplicationRecord
- Associations: Relationships between models
- Validations: Data validation rules
- Callbacks: Hooks into model lifecycle
- Scopes: Reusable queries

### Migrations
- Migration Creation: Database schema changes
- Schema Definition: Table and column definitions
- Index Creation: Performance optimization
- Data Migration: Moving or transforming data

### Model Relationships
- One-to-One: Single related record
- One-to-Many: Multiple related records
- Many-to-Many: Multiple records on both sides
- Polymorphic: Flexible relationship types

### Validations
- Presence: Required fields
- Uniqueness: Unique values
- Format: Pattern matching
- Length: String size limits
- Numericality: Number validation
- Custom: User-defined rules

## 🔹 6. Views

### ERB Templates
- Embedded Ruby: Mixing Ruby and HTML
- Output Tags: <%= %> for output
- Logic Tags: <% %> for logic
- Helper Methods: View-specific methods

### Layouts and Partials
- Application Layout: Main template
- Content For: Yield specific content
- Partials: Reusable view components
- Local Variables: Passing data to partials

### Helpers
- View Helpers: Reusable view methods
- Form Helpers: Form generation
- URL Helpers: Path generation
- Asset Helpers: Asset management

### View Rendering
- Render Options: Different rendering methods
- Format Handling: HTML, JSON, XML
- Status Codes: HTTP response codes
- Content Types: Response content types

### Flash Messages
- Notice: Success messages
- Alert: Error messages
- Flash Types: Different message categories
- Flash Now: Current request only

## 🔹 7. RESTful Architecture

### REST Actions
- Index: List resources
- Show: Single resource
- New: Creation form
- Create: Save new
- Edit: Update form
- Update: Save changes
- Destroy: Remove resource

### HTTP Verbs and Status Codes
- GET: Retrieve data
- POST: Create data
- PUT/PATCH: Update data
- DELETE: Remove data
- Status Codes: Response status

### CRUD Operations
- Create: New records
- Read: Retrieve records
- Update: Modify records
- Delete: Remove records

## 🔹 8. Forms and User Input

### Form Helpers
- form_with: Modern form helper
- form_for: Model-based forms
- form_tag: Generic forms
- Field Helpers: Input types

### Strong Parameters
- Parameter Whitelisting
- Nested Parameters
- Array Parameters
- Custom Parameters

### CSRF Protection
- Cross-Site Request Forgery
- Token Generation
- Token Verification
- Exception Handling

## 🔹 9. ActiveRecord Queries

### Query Methods
- Find: Retrieve by ID
- Find By: Retrieve by attribute
- Where: Filter records
- Order: Sort records
- Limit: Restrict results
- Offset: Skip records

### Scopes
- Named Scopes: Reusable queries
- Scope Chaining: Combining scopes
- Default Scopes: Always applied
- Scope Parameters: Dynamic queries

### Callbacks
- Before Callbacks: Pre-operation
- After Callbacks: Post-operation
- Around Callbacks: Wrap operation
- Conditional Callbacks: When to run

## 🔹 10. Serializers

### JSON Serialization
- Basic Serialization: to_json
- Active Model Serializers
- Jbuilder
- Custom Serialization

### API Serialization
- Controller Serialization
- Custom Methods
- Association Handling
- Format Options

## 🔹 11. Validations

### Model Validations
- Presence Validation
- Uniqueness Validation
- Format Validation
- Length Validation
- Numericality Validation
- Custom Validation

### Custom Validators
- Validator Classes
- Validation Methods
- Error Messages
- Validation Context

### Database Constraints
- Not Null Constraints
- Unique Constraints
- Foreign Key Constraints
- Check Constraints

## 🔹 12. REST APIs

### API Controller
- Resource Actions
- Response Format
- Status Codes
- Error Handling

### API Versioning
- URL Versioning
- Header Versioning
- Content Type Versioning
- Custom Versioning

### API Authentication
- Token Authentication
- Basic Authentication
- OAuth
- JWT

## 🔹 13. Authentication & Authorization

### Devise Setup
- User Model
- Authentication Views
- Configuration
- Customization

### JWT Authentication
- Token Generation
- Token Validation
- Token Refresh
- Security Measures

### Authorization with Pundit
- Policy Objects
- Authorization Rules
- Scope Objects
- Error Handling

## 🔹 14. Mailers and Background Jobs

### ActionMailer
- Mailer Classes
- Email Templates
- Delivery Methods
- Configuration

### Background Jobs
- Job Classes
- Queue Management
- Error Handling
- Job Scheduling

## 🔹 15. Testing

### RSpec Setup
- Configuration
- Directory Structure
- Helper Methods
- Custom Matchers

### Model Specs
- Validation Testing
- Association Testing
- Method Testing
- Callback Testing

### Controller Specs
- Action Testing
- Parameter Testing
- Response Testing
- Authentication Testing

### System Specs
- Feature Testing
- JavaScript Testing
- Browser Testing
- Screenshot Testing

## 🔹 16. Asset Pipeline

### Asset Management
- JavaScript Organization
- Stylesheet Organization
- Image Management
- Asset Precompilation

### Webpacker
- JavaScript Modules
- Asset Bundling
- Development Server
- Production Build

## 🔹 17. Active Storage

### File Uploads
- Attachment Types
- Storage Services
- File Validation
- Direct Upload

### Image Processing
- Image Variants
- Processing Options
- Storage Options
- Delivery Methods

## 🔹 18. Real-Time Features

### ActionCable
- Channel Setup
- Connection Handling
- Broadcasting
- Subscription Management

## 🔹 19. Security

### Security Measures
- CSRF Protection
- XSS Prevention
- SQL Injection Prevention
- Secure Headers

## 🔹 20. Caching

### Cache Methods
- Fragment Caching
- Russian Doll Caching
- Low-Level Caching
- Cache Keys

## 🔹 21. Internationalization

### I18n Setup
- Locale Files
- Translation Keys
- Pluralization
- Date/Time Formatting

## 🔹 22. Performance Optimization

### Query Optimization
- Eager Loading
- Counter Cache
- Database Indexes
- Query Analysis

### Background Processing
- Job Queues
- Worker Processes
- Error Handling
- Monitoring

## 🔹 23. Deployment

### Heroku Deployment
- Application Setup
- Database Setup
- Environment Variables
- Add-ons

### Capistrano Deployment
- Server Configuration
- Deployment Tasks
- Environment Setup
- Rollback Procedures

## 🔹 24. Monitoring & Debugging

### Logging
- Log Levels
- Log Format
- Log Rotation
- Custom Logging

### Debugging
- Console Debugging
- Debugger Usage
- Log Analysis
- Error Tracking

### Exception Tracking
- Error Monitoring
- Alert Configuration
- Error Grouping
- Performance Monitoring 