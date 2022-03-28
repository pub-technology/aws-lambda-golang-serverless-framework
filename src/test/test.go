package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//Returning response with AWS Lambda Proxy Response
	return events.APIGatewayProxyResponse{Body: string("Hello world"), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
