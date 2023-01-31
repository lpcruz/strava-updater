package main

import (
	"encoding/json"
	"fmt"
	"os"

	strava "github.com/strava/go.strava"
)

var token = os.Getenv("STRAVA_ACCESS_TOKEN")
var client = strava.NewClient(token)
var service = strava.NewCurrentAthleteService(client)

func main() {
	activities := _getLatestRunningActivity()
	fmt.Printf(activities)
}

func _getLatestRunningActivity() string {
	runs := []*strava.ActivitySummary{}
	activities, _ := service.ListActivities().Page(1).Do()

	for _, activity := range activities {
		if activity.Type == "Run" {
			runs = append(runs, activity)
		}
	}

	activitiesJson, err := json.Marshal(runs[0])
	if err != nil {
		fmt.Println("Error", err)
		os.Exit((1))
	}
	return string(activitiesJson)
}
