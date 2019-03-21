package main

import (
	// "github.com/aws/aws-lambda-go/events"
	// "github.com/aws/aws-lambda-go/lambda"
  	"encoding/json"
  	"fmt"
  	"time"
  	"net/http"
    // "net/url"
)

type SmallCommit struct {
    Message string
}

type CommitList struct {
    Commits []GitHubCommit
}

type GitHubCommit struct {
    Sha string
    Commit SmallCommit
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
    fmt.Println("hello world")
        client := &http.Client{}

        req2, _ := http.NewRequest("GET", "https://api.github.com/repos/swoogles/Physics/commits", nil)
        req2.SetBasicAuth("swoogles", "84966102c7b7c430d40e6c0ae372916d83d51506")

        resp2, _ := client.Do(req2)
            defer resp2.Body.Close()

            var commitList2 []GitHubCommit
    json.NewDecoder(resp2.Body).Decode(&commitList2)
                for i := 0; i < len(commitList2); i++ {
                    fmt.Println(commitList2[i].Sha + " " + commitList2[i].Commit.Message)
                }
    fmt.Println("goodbye world")
}
