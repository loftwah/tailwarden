package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

// LambdaService is a wrapper around the AWS Lambda API client.
type LambdaService struct {
	client *lambda.Client
}

// NewLambdaService creates a new LambdaService instance.
func NewLambdaService() (*LambdaService, error) {
	// Load AWS config from environment variables or shared config file
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	// Create new Lambda API client with loaded config
	client := lambda.NewFromConfig(cfg)

	return &LambdaService{client}, nil
}

// ListFunctions returns a list of all Lambda functions across regions.
func (s *LambdaService) ListFunctions() ([]lambda.FunctionConfiguration, error) {
	var functions []lambda.FunctionConfiguration

	// Retrieve functions from each region
	regions := s.client.Config.Region.Partition().Regions()
	for _, region := range regions {
		req := s.client.ListFunctionsRequest(&lambda.ListFunctionsInput{
			FunctionVersion: "ALL",
			MaxItems:        1000,
		})
		req.SetContext(context.Background())
		req.SetRegion(region)

		res, err := req.Send()
		if err != nil {
			return nil, err
		}

		functions = append(functions, res.Functions...)
	}

	return functions, nil
}

// SearchFunctions returns a list of Lambda functions matching the given search criteria.
func (s *LambdaService) SearchFunctions(runtime, tagKey, tagValue, region string) ([]lambda.FunctionConfiguration, error) {
	// Get functions list
	functions, err := s.ListFunctions()
	if err != nil {
		return nil, err
	}

	// Filter functions by search criteria
	var filtered []lambda.FunctionConfiguration
	for _, f := range functions {
		if runtime != "" && f.Runtime != runtime {
			continue
		}

		if tagKey != "" && tagValue != "" && (f.Tags[tagKey] == nil || *f.Tags[tagKey] != tagValue) {
			continue
		}

		if region != "" && *f.FunctionArn.Region != region {
			continue
		}

		filtered = append(filtered, f)
	}

	return filtered, nil
}
