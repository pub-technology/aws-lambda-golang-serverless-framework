package main

import (
	"context"
	"os"

	"github.com/kobee-tech-stack/aws-lambda-golang-serverless-framework/services"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	tableName, ok := os.LookupEnv("GLOBAL_DYNAMODB_TABLE")
	if !ok {
		panic("Need GLOBAL_DYNAMODB_TABLE environment variable")
	}

	dynamodb := services.NewDynamoDBStore(context.TODO(), tableName)
	domain := services.NewProductsDomain(dynamodb)
	handler := services.NewAPIGatewayV2Handler(domain)
	lambda.Start(handler.AllHandler)
}
