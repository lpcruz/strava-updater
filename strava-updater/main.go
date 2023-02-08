package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	strava "github.com/strava/go.strava"
)

var token = os.Getenv("STRAVA_ACCESS_TOKEN")
var client = strava.NewClient(token)
var service = strava.NewCurrentAthleteService(client)
var newActivityService = strava.NewActivitiesService(client)
var appName = "https://github.com/lpcruz/strava-updater"

func main() {
	activityId := _getLatestRunningActivityId()
	laps := _getLapsForRun(activityId)
	averageHeartRate := _getHeartRate(activityId)
	description := "❤️ avg HR: " + strconv.FormatFloat(math.Round(averageHeartRate), 'g', 5, 64) + " bpm\n⏱️ splits: " + strings.Join(laps, ", ") + " - via " + appName
	newActivityService.Update(activityId).Description(description).Do()
	fmt.Printf(description)
}

func _getLapsForRun(activityId int64) []string {
	laps := []string{}
	runs, _ := newActivityService.ListLaps(activityId).Do()

	for _, lap := range runs {
		laps = append(laps, _secondsToMinutes(lap.ElapsedTime))
	}
	return laps
}

func _getHeartRate(activityId int64) float64 {
	activity, _ := newActivityService.Get(activityId).IncludeAllEfforts().Do()
	return activity.AverageHeartrate
}

func _getLatestRunningActivityId() int64 {
	runs := []*strava.ActivitySummary{}
	activities, err := service.ListActivities().Page(1).Do()

	if err != nil {
		fmt.Println("Something went wrong", err)
	}

	for _, activity := range activities {
		if activity.Type == "Run" || activity.Type == "VirtualRun" {
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
