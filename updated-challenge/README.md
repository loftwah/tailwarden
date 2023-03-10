# Tailwarden Backend Engineer Coding Challenge

![Tailwarden](banner.jpg)

This is a coding challenge for the Backend Engineer position at Tailwarden. The challenge requires you to build a server that connects to an AWS account and exposes an API with two endpoints:

* Endpoint to list all AWS Lambda functions created in the AWS account across regions.

* Endpoint to search for Lambda functions using the following query parameters:

  * `runtime`: filter functions with a given runtime environment.
  * `tags`: filter functions with a given tag value or tag key.
  * `region`: filter functions running in a given AWS region.

You are free to use any modern language/framework, but we are mainly familiar with Golang, JavaScript, Python, Java, and Ruby. The final solution should be scalable, maintainable, and readable.

## Getting Started

To get started with the challenge, please follow the below steps:

1. Fork the challenge repository to your Github account.
2. Clone the forked repository to your local machine.
3. Create a new branch from the `main` branch.
4. Write the solution and commit the code changes to your branch.
5. Create a new pull request against the `main` branch.

## Requirements

### 1. List Lambda functions

The first endpoint should list all the Lambda functions created in the AWS account across regions. This endpoint should not take any query parameters.

### 2. Search Lambda functions

The second endpoint should allow searching for Lambda functions using the following query parameters:

* `runtime`: filter functions with a given runtime environment.
* `tags`: filter functions with a given tag value or tag key.
* `region`: filter functions running in a given AWS region.

This endpoint should return a list of Lambda functions matching the query parameters.

## CI/CD Pipeline

You need to write a CI/CD pipeline to test, build, and deploy the API. You can use any CI/CD tools like CircleCI, GitHub Actions, etc.

## Submission

After you complete the challenge, please create a new repo on your favorite git platform (GitHub, Gitlab, etc) and share the repository URL with us. The repo should contain a README file with instructions on how to run the project.

## Review

After submitting the challenge, we will review your solution and pay special attention to:

* Coding skills
* Code organization (modularity, dependencies between modules, naming, etc)
* Overall code quality (edge cases, usage of tools, performance, best practices)

Good luck!
