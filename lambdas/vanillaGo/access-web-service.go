package main

import (
	// "github.com/aws/aws-lambda-go/events"
	// "github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	//"io/ioutil"
	// "net/url"
)

type SmallCommit struct {
	Message string
}

type CommitList struct {
	Commits []GitHubCommit
}

type GitHubCommit struct {
	Sha    string
	Commit SmallCommit
}

type ForeCast struct {
	Timezone string //`json:"timezone"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	client := &http.Client{}

	req2, _ := http.NewRequest("GET", "https://api.github.com/repos/swoogles/Physics/commits", nil)
	req2.SetBasicAuth("swoogles", "e67b98d1a2d17730e45e9e26182f44d77985597b")

	resp2, _ := client.Do(req2)
	defer resp2.Body.Close()

	var finalText = "This is live data: \n"
	var commitList2 []GitHubCommit
	json.NewDecoder(resp2.Body).Decode(&commitList2)
	for i := 0; i < len(commitList2); i++ {
		finalText += commitList2[i].Sha + "\n    " + commitList2[i].Commit.Message + "\n\n"
	}
	fmt.Println(finalText)
}
