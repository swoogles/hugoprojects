package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMovieStuff(t *testing.T) {

	var decodedMovie Movie
	error := json.Unmarshal(json.RawMessage(LordOfTheRingsData()), &decodedMovie)

	if error != nil {
		panic(error)
	}
	fmt.Println(decodedMovie.Title)

	out, err := json.MarshalIndent(decodedMovie, "", "  ")
	if err != nil {
		panic (err)
	}

	fmt.Println(string(out))
}
