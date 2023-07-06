package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Word struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func main() {

	args := os.Args
	if len(args) < 2 {
		os.Exit(1)
	}

	fmt.Printf("URl Entered %v \n", args[1])

	givenUrl := args[1]

	if _, err := url.Parse(givenUrl); err != nil {
		log.Fatalf("Invalid URL - %v", err)
	}

	response, err := http.Get(givenUrl)

	if err != nil {
		log.Fatal("Failed to get the response - %v", err)
	}

	defer response.Body.Close()

	fmt.Printf("Response before parsing %v \n", response.Body)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatalf("Failed to parse the response %v", err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("HTTP STatus code - %v \n ", response.StatusCode)
		os.Exit(1)
	}

	var display Word

	err = json.Unmarshal(body, &display)

	if err != nil {
		fmt.Printf("Json unmarshal failed - %v", err)
	}

	fmt.Printf("Json Parsed \n Page : %v \n Words : %v \n", display.Page, strings.Join(display.Words, ","))

}
