# Action Cable Syntax

## 1. Channel Creation

### Basic Channel
```ruby
# app/channels/chat_channel.rb
class ChatChannel < ApplicationCable::Channel
  def subscribed
    stream_from "chat_#{params[:room]}"
  end

  def unsubscribed
    # Any cleanup needed
  end

  def receive(data)
    ActionCable.server.broadcast("chat_#{params[:room]}", data)
  end
end
```

### Model Channel
```ruby
# app/channels/notification_channel.rb
class NotificationChannel < ApplicationCable::Channel
  def subscribed
    stream_for current_user
  end

  def unsubscribed
    stop_all_streams
  end
end
```

## 2. Broadcasting

### Direct Broadcasting
```ruby
# From anywhere in your application
ActionCable.server.broadcast("chat_room_1", {
  message: "Hello!",
  user: current_user.name
})
```

### Model Broadcasting
```ruby
# app/models/message.rb
class Message < ApplicationRecord
  after_create_commit -> {
    broadcast_append_to "chat_#{room_id}",
      target: "messages",
      partial: "messages/message",
      locals: { message: self }
  }
end
```

## 3. JavaScript Integration

### Basic Setup
```javascript
// app/javascript/channels/consumer.js
import consumer from "./consumer"

// Subscribe to a channel
const subscription = consumer.subscriptions.create("ChatChannel", {
  room: "room_1"
}, {
  connected() {
    console.log("Connected to chat channel")
  },
  
  disconnected() {
    console.log("Disconnected from chat channel")
  },
  
  received(data) {
    console.log("Received:", data)
  }
})

// Send a message
subscription.send({ message: "Hello!" })
```

### Channel with Parameters
```javascript
// app/javascript/channels/notification_channel.js
import consumer from "./consumer"

consumer.subscriptions.create({ 
  channel: "NotificationChannel",
  user_id: currentUserId 
}, {
  received(data) {
    this.appendNotification(data)
  },
  
  appendNotification(data) {
    const element = document.getElementById("notifications")
    element.insertAdjacentHTML("beforeend", data.html)
  }
})
```

## 4. Connection Authentication

### Basic Authentication
```ruby
# app/channels/application_cable/connection.rb
module ApplicationCable
  class Connection < ActionCable::Connection::Base
    identified_by :current_user

    def connect
      self.current_user = find_verified_user
    end

    private
      def find_verified_user
        if verified_user = User.find_by(id: cookies.signed[:user_id])
          verified_user
        else
          reject_unauthorized_connection
        end
      end
  end
end
```

### Token-based Authentication
```ruby
# app/channels/application_cable/connection.rb
module ApplicationCable
  class Connection < ActionCable::Connection::Base
    identified_by :current_user

    def connect
      self.current_user = find_verified_user
    end

    private
      def find_verified_user
        if token = request.params[:token]
          User.find_by(authentication_token: token)
        else
          reject_unauthorized_connection
        end
      end
  end
end
```

## 5. Channel Authorization

### Basic Authorization
```ruby
# app/channels/chat_channel.rb
class ChatChannel < ApplicationCable::Channel
  def subscribed
    if can_access_room?(params[:room])
      stream_from "chat_#{params[:room]}"
    else
      reject
    end
  end

  private
    def can_access_room?(room_id)
      current_user.rooms.exists?(room_id)
    end
end
```

### Role-based Authorization
```ruby
# app/channels/admin_channel.rb
class AdminChannel < ApplicationCable::Channel
  def subscribed
    if current_user.admin?
      stream_from "admin_channel"
    else
      reject
    end
  end
end
```

## 6. Testing

### Channel Tests
```ruby
# test/channels/chat_channel_test.rb
require "test_helper"

class ChatChannelTest < ActionCable::Channel::TestCase
  def setup
    @user = users(:one)
    @connection = ApplicationCable::Connection.new(
      server,
      "warden" => OpenStruct.new(user: @user)
    )
  end

  test "subscribes and streams for room" do
    subscribe room: "room_1"
    assert subscription.confirmed?
    assert_has_stream "chat_room_1"
  end

  test "broadcasts message to room" do
    subscribe room: "room_1"
    message = { content: "Hello!" }
    
    assert_broadcast_on("chat_room_1", message) do
      perform :receive, message
    end
  end
end
```

### Connection Tests
```ruby
# test/channels/application_cable/connection_test.rb
require "test_helper"

class ApplicationCable::ConnectionTest < ActionCable::Connection::TestCase
  def setup
    @user = users(:one)
  end

  test "connects with valid user" do
    cookies.signed[:user_id] = @user.id
    connect
    assert_equal @user, connection.current_user
  end

  test "rejects connection without user" do
    assert_reject_connection { connect }
  end
end
```

## 7. Production Configuration

### Redis Adapter
```ruby
# config/cable.yml
production:
  adapter: redis
  url: redis://localhost:6379/1
  channel_prefix: your_app_production
```

### Custom Adapter
```ruby
# config/initializers/action_cable.rb
ActionCable.server.config.cable = {
  adapter: :custom,
  custom_option: "value"
}
```

## 8. Error Handling

### Channel Error Handling
```ruby
# app/channels/chat_channel.rb
class ChatChannel < ApplicationCable::Channel
  def subscribed
    stream_from "chat_#{params[:room]}"
  rescue StandardError => e
    reject_subscription
    logger.error "ChatChannel Error: #{e.message}"
  end
end
```

### JavaScript Error Handling
```javascript
// app/javascript/channels/chat_channel.js
consumer.subscriptions.create("ChatChannel", {
  connected() {
    console.log("Connected")
  },
  
  rejected() {
    console.error("Connection rejected")
  },
  
  received(data) {
    try {
      this.handleMessage(data)
    } catch (error) {
      console.error("Error handling message:", error)
    }
  }
})
``` 