# Introduction to Background Jobs Theory

Background jobs are tasks that are performed outside the normal request-response cycle of a web application. They are crucial for handling operations that would otherwise block the main application thread, leading to slow responses or timeouts for users.

## Why Use Background Jobs?

Web requests should be fast and responsive. If a user action triggers a long-running task, performing that task synchronously during the request will make the user wait. This leads to a poor user experience and can even cause the request to time out. Background jobs allow you to offload these tasks to be processed asynchronously, freeing up the web server to handle other requests.

Common use cases for background jobs include:

*   **Sending emails:** Especially bulk emails or emails that require external API calls (e.g., `deliver_later` uses background jobs).
*   **Processing images or videos:** Resizing, encoding, or applying filters.
*   **Generating reports:** Creating complex reports that might involve querying large datasets.
*   **Importing/Exporting data:** Handling large file uploads or downloads.
*   **Making API calls to external services:** Interacting with third-party APIs that might be slow or unreliable.
*   **Performing scheduled tasks:** Running jobs at specific times or intervals (e.g., daily cleanup, weekly report generation).

By moving these tasks to the background, you ensure your web application remains fast and scalable.

## Active Job Overview

Active Job is a framework introduced in Rails 4.2 for declaring jobs and making them run on a variety of queuing backends. It provides a common API for creating and enqueuing jobs, abstracting away the specifics of the underlying queuing system.

Instead of writing code directly for Sidekiq, Delayed Job, Resque, etc., you write your jobs using Active Job's API. You then configure Active Job to use your preferred queuing backend.

### Key Concepts of Active Job:

*   **Jobs:** Ruby classes that inherit from `ApplicationJob` (or `ActiveJob::Base`). Each job class defines a `perform` method, which contains the code that should be executed in the background.
*   **Queues:** Jobs are placed into queues for processing. You can specify the queue a job belongs to.
*   **Adapters:** Active Job uses adapters to connect to different queuing backends. You configure the adapter in your Rails application.
*   **Workers:** Separate processes that run in the background, picking up jobs from the queues and executing their `perform` methods.

Active Job allows you to switch queuing backends with minimal changes to your job code. This provides flexibility and prevents vendor lock-in.

```ruby
# app/jobs/guest_cleanup_job.rb

class GuestCleanupJob < ApplicationJob
  queue_as :low_priority # Assign the job to a queue

  # The perform method contains the logic to run in the background
  def perform(*guests)
    # Do something later
    # For example, clean up guest accounts
    guests.each { |guest| guest.destroy }
  end
end
```

# Generating Jobs Theory

Rails provides a convenient generator to create new job classes, making it easy to get started with Active Job. The generator creates the basic job file structure, including the `perform` method where you'll put the code that runs in the background.

## Using the Job Generator

Use the `rails generate job` command followed by the name of your job. By convention, job names are singular and end with `Job`.

```bash
rails generate job ProcessImageData
# This will create app/jobs/process_image_data_job.rb
# and test/jobs/process_image_data_job_test.rb

rails generate job SendWelcomeEmail
# This will create app/jobs/send_welcome_email_job.rb
# and test/jobs/send_welcome_email_job_test.rb
```

## Structure of a Generated Job File

A newly generated job file is placed in the `app/jobs/` directory and inherits from `ApplicationJob` (or `ActiveJob::Base`). It includes a placeholder `perform` method.

```ruby
# app/jobs/process_image_data_job.rb

class ProcessImageDataJob < ApplicationJob
  queue_as :default # You can specify a different queue here or later

  # The perform method is where you put the code that should run in the background.
  # Arguments passed to #perform_later are received as arguments to this method.
  def perform(*args)
    # Do something later
    # For example: resize an image, analyze data, send an email

    # The *args allows the method to accept any number of arguments.
    # You'll typically unpack these arguments or expect specific types.

    # Example: Assuming the first argument is an Image object ID
    # image_id = args[0]
    # image = Image.find(image_id)
    # image.resize! # Call a method that does the work

  end
end
```

*   **`queue_as :default`:** This line sets the default queue for this job. You can change `:default` to another name (e.g., `:high_priority`, `:low_priority`, `:mailers`) to organize your jobs into different queues. This is useful when you have workers configured to process specific queues.
*   **`perform(*args)`:** This is the core method where your background task logic goes. When you enqueue a job using `YourJob.perform_later(arg1, arg2, ...)`, the arguments `arg1`, `arg2`, etc., are passed to this `perform` method. The `*args` syntax is a Ruby splat operator that collects all arguments into an array. You can also define specific arguments if you know what to expect, e.g., `def perform(image_id)`.

The generator also creates a corresponding test file in `test/jobs/` to help you write tests for your job's logic. The next topic will cover how to enqueue these jobs to be processed in the background.

# Enqueuing Jobs Theory

Enqueuing a job means adding it to a queue so that it can be processed by a background worker. Active Job provides simple methods to add your defined jobs to the queue.

## The `perform_later` Method

The primary method for enqueuing a job is `perform_later`. You call this method on your job class, passing any arguments that the job's `perform` method needs.

```ruby
# Assuming you have a job defined like this:
# app/jobs/process_image_data_job.rb
# class ProcessImageDataJob < ApplicationJob; def perform(image_id); ... end; end

# To enqueue the job from a controller, model, or service:
image = Image.find(params[:id])
ProcessImageDataJob.perform_later(image.id)

# You can pass multiple arguments:
ProcessImageDataJob.perform_later(image.id, current_user.id, 'medium')
```

When `perform_later` is called, Active Job serializes the job details (job class name, method name, arguments) and sends them to the configured queue adapter. The adapter then adds this job information to the queue backend (e.g., a database table, Redis, etc.).

## Enqueuing Jobs with a Delay or Specific Time

Active Job allows you to schedule jobs to run at a future time using the `set` method chain.

*   **Run in the future:** Use `wait` with a duration.
    ```ruby
    SendFollowUpEmailJob.set(wait: 1.hour).perform_later(user.id)
    # This job will be enqueued but won't be processed by a worker until at least 1 hour from now.
    ```

*   **Run at a specific time:** Use `wait_until` with a `Time` object.
    ```ruby
    report_date = 1.week.from_now.end_of_day
    GenerateWeeklyReportJob.set(wait_until: report_date).perform_later('weekly', Date.today)
    # This job will be enqueued and processed after the specified time.
    ```

The `set` method must be chained before `perform_later`.

## Passing GlobalID Objects

Active Job has built-in support for GlobalID. If you pass an Active Record object to `perform_later`, Active Job will automatically serialize it into a GlobalID and deserialize it back into an Active Record object when the job is performed. This is generally preferred over passing just the ID, as it ensures the record exists when the job runs.

```ruby
# Assuming you have a user object
user = User.find(params[:id])

# Pass the entire user object (Active Job handles serialization/deserialization)
SendWelcomeEmailJob.perform_later(user)
```

```ruby
# In your job's perform method
class SendWelcomeEmailJob < ApplicationJob
  def perform(user)
    # 'user' here is the deserialized User object
    UserMailer.with(user: user).welcome_email.deliver_now # Or deliver_later if mailer uses async
  end
end
```

Enqueuing jobs is typically done from parts of your application that respond to user input or events (like controllers or service objects) or from scheduled tasks. The next topic will cover how to set up a queue adapter and run the background workers that actually process these enqueued jobs.

# Processing Jobs Theory

Once jobs are enqueued using `perform_later`, they sit in a queue waiting to be processed. To execute these jobs, you need two things:

1.  A configured **queue adapter** that tells Active Job which queuing backend to use.
2.  Running **background workers** that connect to the queue backend, pull jobs off the queue, and execute their `perform` methods.

## Configuring the Queue Adapter

The queue adapter is configured in your environment files (`config/environments/*.rb`) using `config.active_job.queue_adapter`.

*   **Development/Test:** Rails defaults to the `:async` adapter in development and test. This adapter runs jobs in a simple in-process thread pool, which is convenient for development and testing but not suitable for production.

    ```ruby
    # config/environments/development.rb
    Rails.application.configure do
      # ...
      config.active_job.queue_adapter = :async # Default in development
      # ...
    end

    # config/environments/test.rb
    Rails.application.configure do
      # ...
      config.active_job.queue_adapter = :test # Default in test
      # ...
    end
    ```
    The `:test` adapter does not process jobs automatically; instead, it stores them in `ActiveJob::Base.queue_adapter.enqueued_jobs` and `performed_jobs` for inspection in tests.

*   **Production:** In production, you should configure a robust, standalone queuing backend like Sidekiq, Delayed Job, Resque, or others. You'll add the corresponding gem to your `Gemfile` and configure the adapter.

    ```ruby
    # Gemfile
    # ...
    gem 'sidekiq' # For Sidekiq
    # gem 'delayed_job_active_record' # For Delayed Job
    # ...
    ```

    ```ruby
    # config/environments/production.rb
    Rails.application.configure do
      # ...
      config.active_job.queue_adapter = :sidekiq # Example for Sidekiq
      # config.active_job.queue_adapter = :delayed_job # Example for Delayed Job
      # ...
    end
    ```

## Running Background Workers

Once the adapter is configured and jobs are enqueued, you need to start the background worker processes. How you do this depends on the chosen queue backend.

*   **`:async` adapter (Development/Test):** Workers run automatically within the same Rails process when jobs are enqueued with `perform_later`. You don't need to start a separate worker process.

*   **Sidekiq:** Requires a separate Sidekiq process. You start it from your application's root directory:

    ```bash
    bundle exec sidekiq
    ```
    Sidekiq connects to Redis (its required backend) and starts processing jobs from the queues.

*   **Delayed Job (Active Record):** Requires a separate worker process. You start it:

    ```bash
    bundle exec rake jobs:work
    ```
    Delayed Job processes jobs stored in a database table.

*   **Other Backends:** Consult the documentation for your chosen queuing system.

Workers are typically long-running processes deployed alongside your web application servers. They continuously monitor the queue for new jobs. You can often configure the number of worker processes, the queues they listen to, and other settings depending on the backend.

Properly setting up your queue adapter and running workers is essential for your background jobs to be processed. The next topic will cover how Active Job handles passing arguments to your jobs, including serialization.
