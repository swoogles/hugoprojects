package main

import (
	// "github.com/aws/aws-lambda-go/events"
	// "github.com/aws/aws-lambda-go/lambda"
  	"encoding/json"
  	"fmt"
  	"log"
  	"time"
  	"net/http"
    // "net/url"
)

type CommitList struct {
    Commits []Commit
}

type Commit struct {
    Sha string
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
	// Make the handler available for Remote Procedure Call by AWS Lambda
		// phone := "14158586273"
    	// QueryEscape escapes the phone string so
    	// it can be safely placed inside a URL query
    	// safePhone := url.QueryEscape(phone)

    	// url := fmt.Sprintf("/orgs/swoogles/repos", safePhone)
    	// url := fmt.Sprintf("https://api.github.com/repos/swoogles/Physics/commits")
    	url := fmt.Sprintf("https://api.github.com/repos/swoogles/Physics/commits/8b415d488aec897b7d8fda662e0ef1b99b4c34a7")

    	// Build the request
    	req, err := http.NewRequest("GET", url, nil)
    	if err != nil {
            fmt.Println(err)
    		log.Fatal("NewRequest: ", err)
    		return
    		}
            fmt.Println(req)
        	// For control over HTTP client headers,
        	// redirect policy, and other settings,
        	// create a Client
        	// A Client is an HTTP client
        	client := &http.Client{}

        	// Send the request via a client
        	// Do sends an HTTP request and
        	// returns an HTTP response
        	resp, err := client.Do(req)
        	if err != nil {
        		log.Fatal("Do: ", err)
        		return
        	}

        	// Callers should close resp.Body
        	// when done reading from it
        	// Defer the closing of the body
            defer resp.Body.Close()
        	fmt.Println(json.NewDecoder(resp.Body))
        	// var record Numverify

            foo1 := new(Commit) // or &Foo{}
            getJson("https://api.github.com/repos/swoogles/Physics/commits/8b415d488aec897b7d8fda662e0ef1b99b4c34a7", foo1)
            println(foo1.Sha)

            var commitList []Commit
            // commitList := [2] Commit{ Commit {}, Commit {} }// or &Foo{}
            getJson("https://api.github.com/repos/swoogles/Physics/commits", commitList)
            fmt.Println(commitList)


        	// Use json.Decode for reading streams of JSON data
        	// if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
        		// log.Println(err)
        // }
    // fmt.Println(record)
        req2, err := http.NewRequest("GET", "https://api.github.com/repos/swoogles/Physics/commits", nil)
        req2.SetBasicAuth("swoogles", "84966102c7b7c430d40e6c0ae372916d83d51506")

        // req2.Header.Add("Authorization","Basic " + basicAuth("swoogles","84966102c7b7c430d40e6c0ae372916d83d51506"))

        resp2, err2 := client.Do(req2)
            defer resp2.Body.Close()


            fmt.Println(commitList)
            var commitList2 []Commit
    json.NewDecoder(resp2.Body).Decode(&commitList2)
            fmt.Println(resp2.Body)
            fmt.Println(err2)
            fmt.Println(commitList2)
    fmt.Println("goodbye world")
}
