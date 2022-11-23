package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kobee-tech-stack/aws-lambda-golang-serverless-framework/services"
	"os"
)

func main() {
	tableName, ok := os.LookupEnv("GLOBAL_DYNAMODB_TABLE")
	if !ok {
		panic("Need GLOBAL_DYNAMODB_TABLE environment variable")
	}

	dynamodb := services.NewDynamoDBStore(context.TODO(), tableName)
	domain := services.NewProductsDomain(dynamodb)
	handler := services.NewAPIGatewayV2Handler(domain)
	lambda.Start(handler.DeleteHandler)
}
