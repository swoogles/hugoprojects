package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
    "fmt"
  	"encoding/json"
  	"time"
  	"net/http"
  	"os"
	"github.com/swoogles/hugoprojects/lambdas/weather"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
	fmt.Println(r)
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(&target)
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var darkSkyToken = os.Getenv("DARK_SKY_TOKEN")
    client := &http.Client{}

    req2, _ := http.NewRequest("GET", "https://api.github.com/repos/swoogles/Physics/commits", nil)
    req2.SetBasicAuth("swoogles", "01fd15407121c063f75f086f21440095e753c869")

    resp2, _ := client.Do(req2)
    defer resp2.Body.Close()

	var weatherForecast weather.ForeCast
    getJson("https://api.darksky.net/forecast/" + darkSkyToken + "/37.8267,-122.4233", weatherForecast)


    var responseHead string = fmt.Sprintf("%b", resp2)
    var finalText string = responseHead + "\nThis is live data: \n"
            var commitList2 []GitHubCommit
    json.NewDecoder(resp2.Body).Decode(&commitList2)
                for i := 0; i < len(commitList2); i++ {
                    finalText += commitList2[i].Sha + " " + commitList2[i].Commit.Message
                }
	var weatherText string = "This is live data: \n" + weatherForecast.Timezone
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       finalText + weatherText,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
