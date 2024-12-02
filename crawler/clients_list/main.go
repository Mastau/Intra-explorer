package main

import (
	"crawler/api42"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func FilterProfiles(data []byte) ([]api42.UserProfile, error) {
	var responses []api42.UserProfileResponse
	if err := json.Unmarshal(data, &responses); err != nil {
		return nil, err
	}

	if len(responses) == 0 {
		return nil, fmt.Errorf("empty body")
	}

	var profiles []api42.UserProfile
	for _, resp := range responses {
		if resp.Active && !resp.Staff {
			profiles = append(profiles, api42.UserProfile{
				Id:        resp.Id,
				Login:     resp.Login,
				Image:     resp.ProfileImage.Link,
				FirstName: resp.FirstName,
				LastName:  resp.LastName,
				PoolYear:  resp.PoolYear,
			})
		}
	}

	return profiles, nil
}

func requestAllActifProfil(token string) ([]api42.UserProfile, error) {
	var allProfiles []api42.UserProfile
	for i := 2; ; i++ {
		resp, err := api42.RequestPageUser(token, i)
		if err != nil {
			return allProfiles, err
		}

		profileResponse, err := FilterProfiles(resp.Body())

		if err != nil {
			return allProfiles, err
		}

		log.Println("page :", i)
		log.Printf("profileResponse: %v\n", profileResponse)

		allProfiles = append(allProfiles, profileResponse...)

		// need to change this shit
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go output")
		return
	}

	token, err := api42.GetAccessToken()
	if err != nil {
		log.Fatalf("Error fetching access token: %v", err)
	}
	log.Println("Access Token:", token)

	allProfiles, err := requestAllActifProfil(token)
	if err != nil {
		log.Printf("err: %v\n", err)
	}

	finalJSON, err := json.MarshalIndent(allProfiles, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling final JSON: %v\n", err)
		return
	}

	outputFilePath := os.Args[1]
	err = os.WriteFile(outputFilePath, finalJSON, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON to file: %v\n", err)
		return
	}

	fmt.Printf("Aggregated Profiles JSON saved to %s\n", outputFilePath)

}
