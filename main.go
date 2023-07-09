package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Page struct {
	Name string `json:"page"`
}

type Word struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

type Occurences struct {
	WordOccurence map[string]int `json:"words"`
}

type Response interface {
	GetResponse() string
}

func (w Word) GetResponse() string {
	return fmt.Sprintf("Json Parsed for Words : %v \n", strings.Join(w.Words, ","))
}

func (o Occurences) GetResponse() string {

	out := []string{}
	for word, count := range o.WordOccurence {
		out = append(out, fmt.Sprintf("%s : %v", word, count))
	}
	return fmt.Sprintf("%v .\n", strings.Join(out, ",\n"))
}

func main() {

	var (
		requestURL string
		password   string
		parsedURl  *url.URL
		err        error
	)

	flag.StringVar(&requestURL, "url", "", "URL to access")
	flag.StringVar(&password, "Password", "", "Enter a password to access ")

	flag.Parse()

	if parsedURl, err = url.ParseRequestURI(requestURL); err != nil {
		fmt.Printf("Validation Error : URL is not valid: %s\n", err)
		flag.Usage()
		os.Exit(1)
	}

	chapter := "introduction"
	//chapter = "interface"

	if chapter == "introduction" {
		// args := os.Args
		// if len(args) < 2 {
		// 	os.Exit(1)
		// }

		// fmt.Printf("URl Entered %v \n", args[1])

		// if _, err := url.Parse(args[1]); err != nil {
		// 	log.Fatalf("Invalid URL - %v", err)
		// }

		// response, err := doRequest(args[1])
		response, err := doRequest(parsedURl.String())

		if err != nil {
			log.Fatalf("Failed to get the response - %v", err)
		}

		if response == nil {
			log.Fatalf("Failed to get the response ( Nil )- %v", err)
		}

		fmt.Printf("Response : \n%s", response.GetResponse())

	} else if chapter == "interface" {
		MyReader()
	}
}

func doRequest(requestURL string) (Response, error) {

	response, err := http.Get(requestURL)

	if err != nil {
		log.Fatalf("Failed to get the response - %v", err)
	}

	defer response.Body.Close()

	//fmt.Printf("Response before parsing %v \n", response.Body)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatalf("Failed to parse the response %v", err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("HTTP STatus code - %v \n ", response.StatusCode)
		os.Exit(1)
	}

	var page Page

	err = json.Unmarshal(body, &page)

	if err != nil {
		fmt.Printf("Json unmarshal failed - %v", err)
	}

	switch page.Name {
	case "occurrence":
		var occurrence Occurences
		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			fmt.Printf("Json unmarshal failed - %v", err)
			return nil, err
		}
		fmt.Printf("Json Parsed for Occurences \n")
		return occurrence, nil

	case "words":
		var words Word
		err = json.Unmarshal(body, &words)
		if err != nil {
			fmt.Printf("Json unmarshal failed - %v", err)
			return nil, err
		}
		return words, nil

	}

	return nil, nil

}
