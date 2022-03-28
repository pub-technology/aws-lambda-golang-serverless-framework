package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/kobee-tech-stack/aws-lambda-golang-serverless-framework/services"
	"testing"
)

func TestLambdaLocally(t *testing.T) {
	// fake a request and call a handler
	request := events.APIGatewayV2HTTPRequest{
		Body: "test",
		PathParameters: map[string]string{
			"id": "1",
		},
	}
	tableName := "global-product-development"
	dynamodb := services.NewDynamoDBStore(context.TODO(), tableName)
	domain := services.NewProductsDomain(dynamodb)
	handler := services.NewAPIGatewayV2Handler(domain)

	handler.GetHandler(context.TODO(), request)
}
