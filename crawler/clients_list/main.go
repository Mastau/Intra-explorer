package main

import (
	"log"

	"crawler/api42"
)

func main() {
	token, err := api42.GetAccessToken()
	if err != nil {
		log.Fatalf("Error fetching access token: %v", err)
	}

	log.Println("Access Token:", token)
}
