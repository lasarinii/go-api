package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var PORT string = ":8080"
var factsURL string = "https://uselessfacts.jsph.pl/api/v2/facts/random?lang=en"
var productsURL string = "https://dummyjson.com/products"

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       int     `json:"price"`
	Rating      float32 `json:"rating"`
	Brand       string  `json:"brand"`
	Category    string  `json:"category"`
}
type ProductList struct {
	Products []Product `json:"products"`
}

type Fact struct {
	Id         string `json:"id"`
	Text       string `json:"text"`
	Source     string `json:"source"`
	Source_url string `json:"source_url"`
	Language   string `json:"language"`
	Permalink  string `json:"permalink"`
}

func getProducts() (productsStr string) {
	resp, reqErr := http.Get(productsURL)
	if reqErr != nil {
		log.Fatalln(reqErr)
	}

	defer resp.Body.Close()

	bodyBytes, convertErr := ioutil.ReadAll(resp.Body)
	if convertErr != nil {
		log.Fatalln(convertErr)
	}

	var products ProductList

	unmarshalErr := json.Unmarshal(bodyBytes, &products)
	if unmarshalErr != nil {
		log.Fatalln(unmarshalErr)
	}

	productsBytes, marshalErr := json.Marshal(products.Products)
	if marshalErr != nil {
		log.Fatalln(marshalErr)
	}

	productsStr = string(productsBytes)

	return
}

func getFacts() (factStr string) {
	resp, reqErr := http.Get(factsURL)
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

func factsReq(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, getFacts())
}

func productsReq(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, getProducts())
}

func main() {
	http.HandleFunc("/facts", factsReq)
	http.HandleFunc("/products", productsReq)

	fmt.Printf("Running on port %v", PORT)

	http.ListenAndServe(PORT, nil)
}
