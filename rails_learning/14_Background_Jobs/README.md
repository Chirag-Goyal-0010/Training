# 14 - Background Jobs

This section covers how to perform tasks asynchronously in your Rails application using Background Jobs and Active Job.

## Topics:

1.  **Introduction to Background Jobs:** Why use them and an overview of Active Job.
2.  **Generating Jobs:** Creating job classes using the Rails generator.
3.  **Enqueuing Jobs:** Sending jobs to the queue using methods like `perform_later`.
4.  **Processing Jobs:** Setting up a queue adapter and running workers.
5.  **Job Arguments and Serialization:** Passing data to jobs and how Rails handles serialization.
6.  **Retries and Error Handling:** Implementing strategies for failed jobs.
7.  **Testing Jobs:** Writing tests for your Active Job classes.

Each topic will have a `theory.md` for explanation, `syntax.md` for code examples, and an `examples` subdirectory for practical demonstrations. 