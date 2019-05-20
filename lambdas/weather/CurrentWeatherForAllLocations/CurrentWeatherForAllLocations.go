package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	//"github.com/swoogles/hugoprojects/lambdas/weather"
	"github.com/swoogles/hugoprojects/lambdas/weather"
	"os"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var darkSkyToken = os.Getenv("DARK_SKY_TOKEN")

	mountainCoordinates := []weather.GpsCoordinates{
		{38.8697, -106.9878, "Crested Butte"},
	}
	var weatherForecast = weather.GetBasicForecast(darkSkyToken, mountainCoordinates[0], "Crested Butte")
	out, _ := json.MarshalIndent(weatherForecast, "", "  ")
	weatherText := string(out)
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       weatherText,
	}, nil
}

func main() {
	lambda.Start(handler)
}
