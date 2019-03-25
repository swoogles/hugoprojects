package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	//"github.com/swoogles/hugoprojects/lambdas/weather"
	"github.com/swoogles/hugoprojects/lambdas/weather"
	"net/http"
	"os"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var darkSkyToken = os.Getenv("DARK_SKY_TOKEN")

    req2, _ := http.NewRequest("GET", "https://api.github.com/repos/swoogles/Physics/commits", nil)
    req2.SetBasicAuth("swoogles", "01fd15407121c063f75f086f21440095e753c869")

    var weatherForecast = weather.GetBasicForecast(darkSkyToken, weather.GpsCoordinates{38.8697, -106.9878})
	/*
                for i := 0; i < len(commitList2); i++ {
                    finalText += commitList2[i].Sha + " " + commitList2[i].Commit.Message
                }
	*/
	out, _ := json.MarshalIndent(weatherForecast, "", "  ")
	var weatherText string = "\n\n\nThis is live weather data: \n" + string(out)
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       weatherText,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
