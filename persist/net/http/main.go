package main

import (
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"log"
)

func main() {
	client := resty.New().SetRetryCount(544)

	resp, err := client.R().Get("https://go00ogle.com")
	if err != nil {
		log.Fatalln(err)
	}

	text, err := ioutil.ReadAll(resp.RawResponse.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf(string(text))
}
