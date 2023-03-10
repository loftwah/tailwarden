# Tailwarden challenge AWS Lambda API

This is a sample Go project that connects to an AWS account and exposes an API to retrieve information about Lambda functions.

## Requirements

To run this project, you'll need:

* Go 1.20+
* An AWS account with an IAM user with appropriate permissions to interact with the Lambda API.

## Setting up AWS account and IAM user

To set up an AWS account, go to the [AWS website](https://aws.amazon.com/) and create an account. You'll need to provide a valid email address, credit card information, and phone number to verify your identity.

Once you've created an AWS account, you'll need to create an IAM user with appropriate permissions to interact with the Lambda API. To do this, follow these steps:

1. Go to the [IAM console](https://console.aws.amazon.com/iam/home).
2. Click on "Users" in the left navigation menu.
3. Click on "Add user" button.
4. Enter a user name and select "Programmatic access" as the access type.
5. Click on "Next: Permissions" button.
6. Click on "Attach existing policies directly" and search for "AWSLambdaFullAccess" policy.
7. Select the policy and click on "Next: Tags" button.
8. Add any tags (optional) and click on "Next: Review" button.
9. Review the user details and click on "Create user" button.
10. Take note of the Access key ID and Secret access key for the user.

## Configuring AWS credentials

To configure your AWS credentials, set the following environment variables:

```bash
export AWS_ACCESS_KEY_ID=<your-access-key-id>
export AWS_SECRET_ACCESS_KEY=<your-secret-access-key>
```

## Installation

To install this project, clone the repository and install dependencies:

```bash
git clone https://github.com/loftwah/tailwarden.git
cd tailwarden
go mod download
```

## Running the server

To start the server, run:

```bash
go run main.go
```

This will start the server on `http://localhost:8080`. You can test the server by visiting `http://localhost:8080/` in your web browser or using a tool like `curl`:

```bash
curl http://localhost:8080/
```

## API endpoints

The following API endpoints are available:

### `GET /functions`

Returns a list of all Lambda functions across regions.

#### `GET /functions` Request parameters

None.

#### `GET /functions` Example response

```json
[
  {
    "FunctionName": "my-function",
    "Runtime": "go1.x",
    "Description": "",
    "MemorySize": 128,
    "Timeout": 3,
    "LastModified": "2021-03-10T09:12:23.260+0000",
    "CodeSize": 1250285,
    "Handler": "main",
    "Role": "arn:aws:iam::123456789012:role/service-role/my-role",
    "Environment": {
      "Variables": {
        "MY_VAR": "my-value"
      }
    },
    "Tags": {
      "my-tag-key": "my-tag-value"
    }
  }
]
```

### `GET /functions/search`

Searches for Lambda functions using the following query params:

* `runtime`: filter functions with a given runtime environment.
* `tagKey`: filter functions with a given tag key.
* `tagValue`: filter functions with a given tag value.
* `region`: filter functions running in a given AWS region.

#### `GET /functions/search` Request parameters

* `runtime` (optional): The runtime environment to filter by.
* `tagKey` (optional): The tag key to filter by.
* `tagValue` (optional): The tag value to filter by.
* `region` (optional): The AWS region to filter by.

#### `GET /functions/search` Example request

```bash
curl 'http://localhost:8080/functions/search?runtime=go1.x&tagKey=my-tag-key&tagValue=my-tag-value&region=ap-southeast-2'
```

#### Example response

```json
[
  {
    "FunctionName": "my-function",
    "Runtime": "go1.x",
    "Description": "",
    "MemorySize": 128,
    "Timeout": 3,
    "LastModified": "2021-03-10T09:12:23.260+0000",
    "CodeSize": 1250285,
    "Handler": "main",
    "Role": "arn:aws:iam::123456789012:role/service-role/my-role",
    "Environment": {
      "Variables": {
        "MY_VAR": "my-value"
      }
    },
    "Tags": {
      "my-tag-key": "my-tag-value"
    }
  }
]
```

## CI/CD pipeline

To set up a CI/CD pipeline, you can use a service like CircleCI or GitHub Actions. Here's an example CircleCI configuration file:

```yaml
version: 2.1
jobs:
  build:
    docker:
      - image: golang:1.20
    steps:
      - checkout
      - run: go mod download
      - run: go test ./...
      - run: go build -o aws-lambda-api
  deploy:
    docker:
      - image: circleci/aws-cli:latest
    steps:
      - run:
          name: Deploy to AWS
          command: |
            aws lambda update-function-code \
              --function-name my-function \
              --zip-file fileb://aws-lambda-api \
              --region ap-southeast-2
workflows:
  version: 2
  build-and-deploy:
    jobs:
      - build:
          filters:
            branches:
              only: master
      - deploy:
          requires:
            - build
          filters:
            branches:
              only: master
```

This configuration file defines two jobs: `build` and `deploy`. The `build` job builds the project, runs tests, and generates an executable binary. The `deploy` job deploys the binary to an existing Lambda function in the `ap-southeast-2` region.

The `build-and-deploy` workflow runs the `build` job on the `master` branch, followed by the `deploy` job if the `build` job succeeds. You'll need to configure CircleCI to use your AWS credentials when deploying to AWS.

## Hosting

To host the API somewhere, you can use a service like AWS Lambda, Amazon EC2/EKS, Heroku, Google App Engine, or any other platform that supports running Go applications. If you choose to use AWS Lambda, you can use the `aws-lambda-go` library to package and deploy your application.

## Conclusion

This project demonstrates how to connect to an AWS account and retrieve information about Lambda functions using Go. With a few modifications, you can adapt this project to work with other AWS services and APIs.
