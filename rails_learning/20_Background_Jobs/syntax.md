# Background Jobs Syntax

## 1. Basic Job Creation

### Simple Job
```ruby
# app/jobs/email_job.rb
class EmailJob < ApplicationJob
  queue_as :default

  def perform(user_id)
    user = User.find(user_id)
    UserMailer.welcome_email(user).deliver_now
  end
end
```

### Job with Parameters
```ruby
# app/jobs/notification_job.rb
class NotificationJob < ApplicationJob
  queue_as :notifications

  def perform(user_id, message, options = {})
    user = User.find(user_id)
    NotificationService.send(user, message, options)
  end
end
```

## 2. Job Enqueuing

### Basic Enqueuing
```ruby
# Enqueue immediately
EmailJob.perform_later(user.id)

# Enqueue with delay
EmailJob.set(wait: 1.hour).perform_later(user.id)

# Enqueue at specific time
EmailJob.set(wait_until: Date.tomorrow.noon).perform_later(user.id)
```

### Queue Selection
```ruby
# app/jobs/priority_job.rb
class PriorityJob < ApplicationJob
  queue_as :high_priority

  def perform(data)
    # Process high priority task
  end
end

# Enqueue to specific queue
PriorityJob.set(queue: :urgent).perform_later(data)
```

## 3. Job Scheduling

### Recurring Jobs
```ruby
# config/initializers/sidekiq.rb
Sidekiq::Cron::Job.create(
  name: 'Daily report - every day at 6am',
  cron: '0 6 * * *',
  class: 'DailyReportJob'
)
```

### Delayed Jobs
```ruby
# app/jobs/delayed_job.rb
class DelayedJob < ApplicationJob
  def perform
    # Job logic here
  end
end

# Schedule with delay
DelayedJob.set(wait: 1.hour).perform_later
```

## 4. Job Chaining

### Basic Chaining
```ruby
# app/jobs/chain_job.rb
class ChainJob < ApplicationJob
  def perform
    # First job
    FirstJob.perform_later
    # Second job
    SecondJob.perform_later
  end
end
```

### Dependent Jobs
```ruby
# app/jobs/dependent_job.rb
class DependentJob < ApplicationJob
  def perform
    # First job
    FirstJob.perform_later
    # Second job depends on first
    SecondJob.set(wait: 5.minutes).perform_later
  end
end
```

## 5. Error Handling

### Retry Logic
```ruby
# app/jobs/retry_job.rb
class RetryJob < ApplicationJob
  retry_on StandardError, wait: :exponentially_longer, attempts: 3

  def perform
    # Job logic here
  end
end
```

### Error Callbacks
```ruby
# app/jobs/callback_job.rb
class CallbackJob < ApplicationJob
  def perform
    # Job logic here
  end

  def on_error(execution)
    # Handle error
    ErrorNotifier.notify(execution.error)
  end
end
```

## 6. Job Testing

### Unit Tests
```ruby
# test/jobs/email_job_test.rb
require "test_helper"

class EmailJobTest < ActiveJob::TestCase
  test "sends welcome email" do
    user = users(:one)
    
    assert_enqueued_with(job: EmailJob, args: [user.id]) do
      EmailJob.perform_later(user.id)
    end
  end
end
```

### Integration Tests
```ruby
# test/integration/job_integration_test.rb
require "test_helper"

class JobIntegrationTest < ActionDispatch::IntegrationTest
  test "job processing flow" do
    user = users(:one)
    
    assert_enqueued_jobs 1 do
      EmailJob.perform_later(user.id)
    end
    
    perform_enqueued_jobs
    
    assert_performed_jobs 1
  end
end
```

## 7. Job Monitoring

### Sidekiq Monitoring
```ruby
# config/initializers/sidekiq.rb
Sidekiq.configure_server do |config|
  config.server_middleware do |chain|
    chain.add Sidekiq::Middleware::Server::RetryJobs, max_retries: 3
  end
end
```

### Custom Monitoring
```ruby
# app/jobs/monitored_job.rb
class MonitoredJob < ApplicationJob
  def perform
    start_time = Time.current
    
    # Job logic here
    
    end_time = Time.current
    duration = end_time - start_time
    
    JobMetrics.record_duration(self.class.name, duration)
  end
end
```

## 8. Job Configuration

### Active Job Configuration
```ruby
# config/application.rb
module YourApp
  class Application < Rails::Application
    config.active_job.queue_adapter = :sidekiq
    config.active_job.queue_name_prefix = "your_app_#{Rails.env}"
  end
end
```

### Sidekiq Configuration
```ruby
# config/initializers/sidekiq.rb
Sidekiq.configure_server do |config|
  config.redis = { url: ENV['REDIS_URL'] }
  
  config.server_middleware do |chain|
    chain.add Sidekiq::Middleware::Server::RetryJobs, max_retries: 3
  end
end

Sidekiq.configure_client do |config|
  config.redis = { url: ENV['REDIS_URL'] }
end
```

## 9. Job Cleanup

### Job Cleanup
```ruby
# app/jobs/cleanup_job.rb
class CleanupJob < ApplicationJob
  def perform
    # Cleanup old jobs
    Sidekiq::RetrySet.new.clear
    Sidekiq::DeadSet.new.clear
  end
end
```

### Data Cleanup
```ruby
# app/jobs/data_cleanup_job.rb
class DataCleanupJob < ApplicationJob
  def perform
    # Cleanup old data
    OldData.where('created_at < ?', 30.days.ago).delete_all
  end
end
```
