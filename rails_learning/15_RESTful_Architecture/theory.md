# RESTful Architecture Theory

## What is REST?

REST (Representational State Transfer) is an architectural style for designing networked applications. It's based on a set of principles that describe how networked resources are defined and addressed.

## REST Principles

1. **Client-Server Architecture**
   - Separation of concerns
   - Independent evolution of client and server

2. **Stateless**
   - Each request contains all information needed
   - No client context stored on server

3. **Cacheable**
   - Responses must define themselves as cacheable or not
   - Improves efficiency and scalability

4. **Uniform Interface**
   - Resource identification
   - Resource manipulation through representations
   - Self-descriptive messages
   - HATEOAS (Hypermedia as the Engine of Application State)

5. **Layered System**
   - Client can't tell if connected directly to server
   - Intermediary servers can improve scalability

## RESTful Resources in Rails

### The 7 Default REST Actions

1. **index**
   - Lists all resources
   - GET request
   - Example: List all articles

2. **show**
   - Displays a specific resource
   - GET request
   - Example: Show a specific article

3. **new**
   - Displays form to create new resource
   - GET request
   - Example: Show form to create new article

4. **create**
   - Creates a new resource
   - POST request
   - Example: Save new article

5. **edit**
   - Displays form to edit resource
   - GET request
   - Example: Show form to edit article

6. **update**
   - Updates an existing resource
   - PATCH/PUT request
   - Example: Save changes to article

7. **destroy**
   - Removes a resource
   - DELETE request
   - Example: Delete an article

## HTTP Verbs and Status Codes

### HTTP Verbs
- **GET**: Retrieve a resource
- **POST**: Create a new resource
- **PUT**: Update an entire resource
- **PATCH**: Partially update a resource
- **DELETE**: Remove a resource

### Common Status Codes
- **200 OK**: Request succeeded
- **201 Created**: Resource created
- **204 No Content**: Request succeeded, no response body
- **400 Bad Request**: Invalid request
- **401 Unauthorized**: Authentication required
- **403 Forbidden**: Server refuses action
- **404 Not Found**: Resource not found
- **422 Unprocessable Entity**: Validation failed
- **500 Internal Server Error**: Server error

## CRUD Operations with ActiveRecord

### Create
```ruby
# Create a new record
Article.create(title: "New Article", content: "Content")

# Create with validation
article = Article.new(title: "New Article")
article.save
```

### Read
```ruby
# Find by ID
Article.find(1)

# Find by conditions
Article.where(published: true)

# Find first matching record
Article.find_by(title: "My Article")
```

### Update
```ruby
# Update attributes
article.update(title: "Updated Title")

# Update single attribute
article.update_attribute(:title, "New Title")

# Update multiple attributes
article.update_attributes(title: "New Title", content: "New Content")
```

### Delete
```ruby
# Delete a record
article.destroy

# Delete without callbacks
article.delete
```

## RESTful Routes in Rails

### Basic Resource
```ruby
resources :articles
```

### Nested Resources
```ruby
resources :articles do
  resources :comments
end
```

### Custom Actions
```ruby
resources :articles do
  member do
    post :publish
  end
  collection do
    get :search
  end
end
```

## Best Practices

1. **Use Appropriate HTTP Verbs**
   - GET for reading
   - POST for creating
   - PUT/PATCH for updating
   - DELETE for deleting

2. **Return Proper Status Codes**
   - Use status codes to indicate success/failure
   - Include appropriate error messages

3. **Keep URLs Clean and Meaningful**
   - Use resource names in URLs
   - Avoid unnecessary nesting
   - Use query parameters for filtering

4. **Version Your API**
   - Include version in URL or header
   - Plan for future changes

5. **Handle Errors Gracefully**
   - Return meaningful error messages
   - Use appropriate status codes
   - Include error details when helpful

6. **Use Proper Response Formats**
   - JSON for APIs
   - HTML for web interfaces
   - XML when required

7. **Implement Proper Authentication**
   - Use tokens or sessions
   - Secure sensitive endpoints
   - Implement rate limiting

8. **Document Your API**
   - Document all endpoints
   - Include request/response examples
   - Document authentication requirements

9. **Test Your Endpoints**
   - Write integration tests
   - Test error cases
   - Test authentication

10. **Monitor Performance**
    - Track response times
    - Monitor error rates
    - Set up alerts 