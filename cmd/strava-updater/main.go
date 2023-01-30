package main

import (
	"encoding/json"
	"fmt"
	"os"

	strava "github.com/strava/go.strava"
)

func main() {
	strava_access_token := os.Getenv("STRAVA_ACCESS_TOKEN")
	client := strava.NewClient(strava_access_token)
	service := strava.NewCurrentAthleteService(client)
	athlete, err := service.Get().Do()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit((1))
	}
	athleteInfo, err := json.Marshal(athlete)
	fmt.Printf(string(athleteInfo))
}
