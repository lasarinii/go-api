package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var PORT string = ":8080"
var url string = "https://uselessfacts.jsph.pl/api/v2/facts/random?lang=en"

type Fact struct {
	Id         string `json:"id"`
	Text       string `json:"text"`
	Source     string `json:"source"`
	Source_url string `json:"source_url"`
	Language   string `json:"language"`
	Permalink  string `json:"permalink"`
}

func getFacts() (factStr string) {
	resp, reqErr := http.Get(url)
	if reqErr != nil {
		log.Fatalln(reqErr)
	}

	defer resp.Body.Close()

	bodyBytes, convertErr := ioutil.ReadAll(resp.Body)
	if convertErr != nil {
		log.Fatalln(convertErr)
	}

	var randFact Fact

	unmarshalErr := json.Unmarshal(bodyBytes, &randFact)
	if unmarshalErr != nil {
		log.Fatalln(unmarshalErr)
	}

	factBytes, marshalErr := json.Marshal(randFact)
	if marshalErr != nil {
		log.Fatalln(marshalErr)
	}

	factStr = string(factBytes)

	return
}

func handleReq(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, getFacts())
}

func main() {
	server := &http.Server{
		Addr:    PORT,
		Handler: http.HandlerFunc(handleReq),
	}

	fmt.Printf("Running on port %v", PORT)

	if serverErr := server.ListenAndServe(); serverErr != nil {
		log.Fatalln(serverErr)
	}
}
