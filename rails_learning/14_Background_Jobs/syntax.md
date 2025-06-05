# Processing Jobs Syntax

This document provides syntax examples for configuring your Active Job queue adapter and commands for running background workers.

## Configuring the Queue Adapter

Set `config.active_job.queue_adapter` in your environment files (`config/environments/*.rb`).

### Development Environment (`config/environments/development.rb`)

The default `:async` adapter is usually sufficient for development.

```ruby
# config/environments/development.rb

Rails.application.configure do
  # ... existing development configurations ...

  # Use the in-line async adapter for development.
  # This runs jobs in a thread pool in the same process.
  config.active_job.queue_adapter = :async

  # ... rest of development configurations ...
end
```

### Test Environment (`config/environments/test.rb`)

The default `:test` adapter is used for testing, capturing jobs without processing.

```ruby
# config/environments/test.rb

Rails.application.configure do
  # ... existing test configurations ...

  # The test adapter collects enqueued and performed jobs to verify in tests.
  config.active_job.queue_adapter = :test

  # ... rest of test configurations ...
end
```

### Production Environment (`config/environments/production.rb`)

Configure a robust production-ready adapter here. Examples for Sidekiq and Delayed Job:

```ruby
# config/environments/production.rb

Rails.application.configure do
  # ... existing production configurations ...

  # Example using Sidekiq (requires 'sidekiq' gem)
  config.active_job.queue_adapter = :sidekiq

  # Example using Delayed Job (requires 'delayed_job_active_record' or similar gem)
  # config.active_job.queue_adapter = :delayed_job

  # ... other adapter configurations (specific to your chosen backend) ...

  # ... rest of production configurations ...
end
```

*(Remember to add the necessary gem for your chosen adapter to your Gemfile and run `bundle install`.)*

## Running Background Workers

How you start workers depends on your chosen adapter.

### For Sidekiq:

Start the Sidekiq process from your Rails application's root directory. Sidekiq requires a running Redis server.

```bash
bundle exec sidekiq
```

You can specify queues and other options:

```bash
bundle exec sidekiq -q critical,high,default,low -c 10
```
*(This starts Sidekiq listening to specified queues with 10 concurrency.)*

### For Delayed Job (Active Record):

Run the rake task from your Rails application's root directory.

```bash
bundle exec rake jobs:work
```

You can run multiple workers or specify queues:

```bash
QUEUE=my_queue bundle exec rake jobs:work
RAILS_MAX_THREADS=5 bundle exec rake jobs:work # For multi-threaded processing
```

### For other adapters:

Consult the documentation for your specific queueing backend (e.g., Resque, Faktory, etc.) for instructions on starting workers.

Running background workers is a separate concern from running your web server processes (like Puma or Unicorn). You typically need both running to handle web requests and process background jobs concurrently. The next topic will cover how to pass arguments to your jobs and how Rails handles their serialization. 