package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
  "fmt"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	const name, age = "Kim", 22
	fmt.Println(name, "is", age, "years old.")
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, World!!",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
