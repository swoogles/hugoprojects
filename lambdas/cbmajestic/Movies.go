package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"os"
	"time"
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
}

func serializeMovie(movie Movie) string {
	out, err := json.MarshalIndent(movie, "", "  ")
	if err != nil {
		panic (err)
	}
	return string(out)
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("Movie things!")
	var openMoviesDbToken = os.Getenv("OPEN_MOVIES_DB_TOKEN")
	fmt.Println("movieToken:" + openMoviesDbToken)

	guardiansOfTheGalaxy2Url := "http://www.omdbapi.com/?i=tt3896198&apikey=" + openMoviesDbToken
	movieRequest, _ := http.NewRequest(
		"GET",
		guardiansOfTheGalaxy2Url,
		nil)
	resp3, error := myClient.Do(movieRequest)
	if (error != nil) {
		fmt.Println(error)
		fmt.Fprintf(os.Stderr, "error: %v\n", error)
		os.Exit(1)
	}

	defer resp3.Body.Close()
	var movie Movie
	json.NewDecoder(resp3.Body).Decode(&movie)
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       serializeMovie(movie),
	}, nil
}

func main() {
	lambda.Start(handler)
}
