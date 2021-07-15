package darksky

import (
	"github.com/collier/office-api/config"

	forecast "github.com/mlbright/forecast/v2"
)

// GetHQWeatherForecast calls the darksky API, and returns weather information
// specific to the current counterpoint headquarters for a given time.
func GetHQWeatherForecast(time string) (*forecast.Forecast, error) {
	key := config.DarkSkyAPIKey
	lat := config.OfficeLatitude
	lon := config.OfficeLongitude

	f, err := forecast.Get(key, lat, lon, time, forecast.US, forecast.English)
	if err != nil {
		return nil, err
	}
	return f, nil
}
