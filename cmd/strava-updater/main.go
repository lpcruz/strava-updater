package main

import (
	"fmt"
	"os"

	strava "github.com/strava/go.strava"
)

var token = os.Getenv("STRAVA_ACCESS_TOKEN")
var client = strava.NewClient(token)
var service = strava.NewCurrentAthleteService(client)
var newActivityService = strava.NewActivitiesService(client)

func main() {
	times := _getLapsForRun()
	fmt.Println(times)
}

func _getLapsForRun() []string {
	laps := []string{}
	activityId := _getLatestRunningActivityId()
	runs, _ := newActivityService.ListLaps(activityId).Do()

	for _, lap := range runs {
		laps = append(laps, _secondsToMinutes(lap.ElapsedTime))
	}
	return laps
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

func _secondsToMinutes(inSeconds int) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str := fmt.Sprint(minutes, ":", seconds)
	return str
}
