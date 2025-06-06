# Active Storage Syntax Examples

## Basic Setup

### Installation
```ruby
# Terminal
rails active_storage:install
rails db:migrate
```

### Configuration
```ruby
# config/storage.yml
local:
  service: Disk
  root: <%= Rails.root.join("storage") %>

amazon:
  service: S3
  access_key_id: <%= Rails.application.credentials.dig(:aws, :access_key_id) %>
  secret_access_key: <%= Rails.application.credentials.dig(:aws, :secret_access_key) %>
  region: us-east-1
  bucket: your-bucket-name

# config/environments/development.rb
config.active_storage.service = :local

# config/environments/production.rb
config.active_storage.service = :amazon
```

## Model Attachments

### Single File Attachment
```ruby
# app/models/user.rb
class User < ApplicationRecord
  has_one_attached :avatar
  
  validates :avatar, content_type: ['image/png', 'image/jpeg'],
                    size: { less_than: 5.megabytes }
end

# Usage
user = User.find(1)
user.avatar.attach(io: File.open("path/to/avatar.jpg"), filename: "avatar.jpg")
user.avatar.attached? # => true
user.avatar.purge # Remove attachment
```

### Multiple File Attachments
```ruby
# app/models/article.rb
class Article < ApplicationRecord
  has_many_attached :images
  
  validates :images, content_type: ['image/png', 'image/jpeg'],
                    size: { less_than: 5.megabytes }
end

# Usage
article = Article.find(1)
article.images.attach(
  io: File.open("path/to/image1.jpg"),
  filename: "image1.jpg"
)
article.images.attach(
  io: File.open("path/to/image2.jpg"),
  filename: "image2.jpg"
)
```

## Image Processing

### Image Variants
```ruby
# app/models/user.rb
class User < ApplicationRecord
  has_one_attached :avatar do |attachable|
    attachable.variant :thumb, resize_to_limit: [100, 100]
    attachable.variant :medium, resize_to_limit: [300, 300]
    attachable.variant :large, resize_to_limit: [800, 800]
  end
end

# Usage in view
<%= image_tag user.avatar.variant(:thumb) %>
<%= image_tag user.avatar.variant(:medium) %>
<%= image_tag user.avatar.variant(:large) %>
```

### Image Transformations
```ruby
# app/models/product.rb
class Product < ApplicationRecord
  has_one_attached :image do |attachable|
    attachable.variant :thumb, resize_to_fill: [100, 100]
    attachable.variant :medium, resize_to_fit: [300, 300]
    attachable.variant :large, resize_and_pad: [800, 800, background: :white]
  end
end

# Usage
product.image.variant(
  resize_to_limit: [100, 100],
  strip: true,
  quality: 80
)
```

## File Uploads

### Form Setup
```erb
<%# app/views/users/_form.html.erb %>
<%= form_with(model: @user, local: true) do |form| %>
  <div class="field">
    <%= form.label :avatar %>
    <%= form.file_field :avatar, direct_upload: true %>
  </div>

  <div class="field">
    <%= form.label :documents %>
    <%= form.file_field :documents, multiple: true, direct_upload: true %>
  </div>

  <div class="actions">
    <%= form.submit %>
  </div>
<% end %>
```

### Controller Handling
```ruby
# app/controllers/users_controller.rb
class UsersController < ApplicationController
  def create
    @user = User.new(user_params)
    
    if @user.save
      redirect_to @user, notice: 'User was successfully created.'
    else
      render :new
    end
  end

  private

  def user_params
    params.require(:user).permit(:name, :email, :avatar, documents: [])
  end
end
```

## Direct Uploads

### JavaScript Setup
```javascript
// app/javascript/controllers/direct_upload_controller.js
import { Controller } from "stimulus"

export default class extends Controller {
  static targets = ["input", "progress"]

  connect() {
    this.inputTarget.addEventListener("change", this.upload.bind(this))
  }

  upload(event) {
    const file = event.target.files[0]
    const upload = new DirectUpload(file, this.url, this)

    upload.create((error, blob) => {
      if (error) {
        console.error(error)
      } else {
        this.createHiddenBlobInput(blob)
      }
    })
  }

  directUploadWillStoreFileWithXHR(request) {
    request.upload.addEventListener("progress", event => {
      const progress = event.loaded / event.total * 100
      this.progressTarget.style.width = `${progress}%`
    })
  }

  createHiddenBlobInput(blob) {
    const hiddenField = document.createElement("input")
    hiddenField.type = "hidden"
    hiddenField.name = this.inputTarget.name
    hiddenField.value = blob.signed_id
    this.inputTarget.parentNode.appendChild(hiddenField)
  }

  get url() {
    return this.data.get("url")
  }
}
```

### View Setup
```erb
<%# app/views/users/_form.html.erb %>
<div data-controller="direct-upload" data-direct-upload-url="<%= rails_direct_uploads_url %>">
  <%= form_with(model: @user, local: true) do |form| %>
    <div class="field">
      <%= form.label :avatar %>
      <%= form.file_field :avatar,
          data: { 
            direct_upload_target: "input"
          },
          direct_upload: true %>
      <div class="progress">
        <div class="progress-bar" data-direct-upload-target="progress"></div>
      </div>
    </div>
  <% end %>
</div>
```

## File Validation

### Model Validation
```ruby
# app/models/document.rb
class Document < ApplicationRecord
  has_one_attached :file
  
  validates :file, presence: true,
                  content_type: ['application/pdf', 'application/msword'],
                  size: { less_than: 10.megabytes }
                  
  validate :acceptable_file
  
  private
  
  def acceptable_file
    return unless file.attached?
    
    unless file.blob.byte_size <= 10.megabytes
      errors.add(:file, "is too big (maximum is 10MB)")
    end
    
    acceptable_types = ["application/pdf", "application/msword"]
    unless acceptable_types.include?(file.blob.content_type)
      errors.add(:file, "must be a PDF or Word document")
    end
  end
end
```

## Image Processing with MiniMagick

### Configuration
```ruby
# config/initializers/image_processing.rb
require "image_processing/mini_magick"

ImageProcessing::MiniMagick.configure do |config|
  config.strip = true
  config.quality = 80
  config.sampling_factor = "4:2:0"
end
```

### Usage
```ruby
# app/models/product.rb
class Product < ApplicationRecord
  has_one_attached :image do |attachable|
    attachable.variant :thumb, resize_to_fill: [100, 100]
    attachable.variant :medium, resize_to_fit: [300, 300]
    attachable.variant :large, resize_to_limit: [800, 800]
  end
end

# Usage in view
<%= image_tag product.image.variant(
  resize_to_limit: [100, 100],
  strip: true,
  quality: 80
) %>
```

## Testing

### Model Tests
```ruby
# test/models/user_test.rb
require "test_helper"

class UserTest < ActiveSupport::TestCase
  test "validates avatar content type" do
    user = User.new(name: "John")
    user.avatar.attach(
      io: File.open(Rails.root.join("test", "fixtures", "files", "test.txt")),
      filename: "test.txt"
    )
    
    assert_not user.valid?
    assert_includes user.errors[:avatar], "is not a valid image type"
  end
  
  test "validates avatar size" do
    user = User.new(name: "John")
    user.avatar.attach(
      io: File.open(Rails.root.join("test", "fixtures", "files", "large_image.jpg")),
      filename: "large_image.jpg"
    )
    
    assert_not user.valid?
    assert_includes user.errors[:avatar], "is too big"
  end
end
```

### Controller Tests
```ruby
# test/controllers/users_controller_test.rb
require "test_helper"

class UsersControllerTest < ActionDispatch::IntegrationTest
  test "should create user with avatar" do
    assert_difference("User.count") do
      post users_url, params: {
        user: {
          name: "John",
          email: "john@example.com",
          avatar: fixture_file_upload("test/fixtures/files/avatar.jpg", "image/jpeg")
        }
      }
    end
    
    assert_redirected_to user_url(User.last)
    assert User.last.avatar.attached?
  end
end
```

## Error Handling

### Controller Error Handling
```ruby
# app/controllers/application_controller.rb
class ApplicationController < ActionController::Base
  rescue_from ActiveStorage::FileNotFoundError do |exception|
    render json: { error: "File not found" }, status: :not_found
  end
  
  rescue_from ActiveStorage::IntegrityError do |exception|
    render json: { error: "Invalid file" }, status: :unprocessable_entity
  end
end
```

### Model Error Handling
```ruby
# app/models/user.rb
class User < ApplicationRecord
  has_one_attached :avatar
  
  def avatar_url
    return nil unless avatar.attached?
    
    begin
      Rails.application.routes.url_helpers.url_for(avatar)
    rescue ActiveStorage::FileNotFoundError
      nil
    end
  end
end
``` 