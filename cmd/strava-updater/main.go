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
var newActivityService = strava.NewActivitiesService(client)

func main() {
	activityId := _getLatestRunningActivityId()

	// returns a slice of LapEffortSummary objects
	laps, _ := newActivityService.ListLaps(activityId).Do()
	lapJson, _ := json.Marshal((laps))
	fmt.Printf(string(lapJson))
}

func _getLatestRunningActivityId() int64 {
	runs := []*strava.ActivitySummary{}
	activities, _ := service.ListActivities().Page(1).Do()

	for _, activity := range activities {
		if activity.Type == "Run" {
			runs = append(runs, activity)
		}
	}

	return runs[0].Id
}
