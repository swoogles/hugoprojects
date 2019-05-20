package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/swoogles/ColoradoWeatherMap/darksky"
	"os"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var darkSkyToken = os.Getenv("DARK_SKY_TOKEN")

	mountainCoordinates := []darksky.GpsCoordinates{
		{38.8697, -106.9878, "Crested Butte"},
	}
	var weatherForecast = darksky.GetBasicForecast(darkSkyToken, mountainCoordinates[0], "Crested Butte")
	out, _ := json.MarshalIndent(weatherForecast, "", "  ")
	weatherText := string(out)
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       weatherText,
	}, nil
}

func main() {
	fmt.Println("WHY?!")
	lambda.Start(handler)
}
