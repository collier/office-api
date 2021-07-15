package darksky

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/collier/office-api/config"

	forecast "github.com/mlbright/forecast/v2"
	"github.com/nlopes/slack"
)

func sendSlackAlert(dp forecast.DataPoint, nextDay time.Time) error {
	url := config.SlackAppInclementWeatherWebhookURL
	dateStr := "Tomorrow"
	if nextDay.Weekday() == 5 {
		dateStr = nextDay.Format("Monday, Jan 2nd")
	}
	msgTmpl := "There is inclement weather in the forecast for *%s*. Remember to bring home your laptop to work from home. See the forecast below:"
	notfMsg := fmt.Sprintf(msgTmpl, dateStr)
	fields := make([]slack.AttachmentField, 0)
	// Add chance of snow or sleet as a field on the slack message when
	// applicable
	if dp.PrecipType == "snow" || dp.PrecipType == "sleet" {
		val := math.Round(dp.PrecipProbability * 100)
		percent := fmt.Sprintf("%v%%", val)
		fields = append(fields, slack.AttachmentField{
			Title: "Chance of " + dp.PrecipType,
			Value: percent,
			Short: true,
		})
	}
	// Add wind speed as a field on the slack message when applicable
	if dp.WindSpeed > 50 {
		val := math.Round(dp.WindSpeed)
		windSpeed := fmt.Sprintf("%v mph", val)
		fields = append(fields, slack.AttachmentField{
			Title: "Wind Speed",
			Value: windSpeed,
			Short: true,
		})
	}
	att := slack.Attachment{
		Title:  "Forecast",
		Text:   dp.Summary,
		Fields: fields,
	}
	wm := slack.WebhookMessage{
		Text:        notfMsg,
		Attachments: []slack.Attachment{att},
	}
	return slack.PostWebhook(url, &wm)
}

// InclementWeatherAlertJob gets the weather for the next day from the darksky
// API, and if inclement weather is in the forecast, send out appropriate alerts
func InclementWeatherAlertJob() {
	now := time.Now()
	daysToAdd := 1
	// Get Monday's forecast when it's Friday.
	if now.Weekday() == 5 {
		daysToAdd = 3
	}
	nextDay := now.AddDate(0, 0, daysToAdd)
	nextDayStamp := fmt.Sprintf("%v", nextDay.Unix())
	timeStr := nextDay.Format("01/02/2006")
	f, err := GetHQWeatherForecast(nextDayStamp)
	if err != nil {
		log.Printf("ERROR %v", err)
	}
	if len(f.Daily.Data) > 0 {
		dp := f.Daily.Data[0]
		fmt.Println("Precipetation Type: " + dp.PrecipType)
		if dp.PrecipType == "snow" || dp.PrecipType == "sleet" || dp.WindSpeed > 50 {
			err := sendSlackAlert(dp, nextDay)
			if err != nil {
				log.Printf("ERROR %v", err)
			}
		}
	} else {
		log.Printf("WARN No daily forecast data available for %s", timeStr)
	}
}
