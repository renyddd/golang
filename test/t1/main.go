package main

import (
	"context"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {
	rootCtx := context.Background()

	ctx, cancel := context.WithTimeout(rootCtx, 2*time.Millisecond)
	defer cancel()

	Query(ctx)
}

func Query(c context.Context) {
	url := "https://googlefdfs.com"

	client := resty.New().SetRetryCount(53)
	resp, err := client.R().SetContext(c).Get(url)

	if err != nil {
		log.Fatalln("get", resp)
	}



	log.Println(resp.RawResponse.Status)
}