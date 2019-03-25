package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Rating struct {
	Source string
	Value string
}

type Movie struct {
	Title string
	Rated string
	Runtime string
	Genre string
	Director string
	Plot string
	Poster string
	Ratings []Rating
}


func getLordOfTheRingsData() {
	/*
req3, _ := http.NewRequest(
"GET",
"https://api.darksky.net/forecast/" +
darkSkyToken +
"/" + stringOf(coordinates), nil)
resp3, error := myClient.Do(req3)
if (error != nil) {
fmt.Fprintf(os.Stderr, "error: %v\n", error)
os.Exit(1)
}

defer resp3.Body.Close()
var weatherForecast ForeCast
json.NewDecoder(resp3.Body).Decode(&weatherForecast)
return weatherForecast;
	*/
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Movie content!",
	}, nil
}

func main() {
	fmt.Println("Movie things!")
	lambda.Start(handler)
}
