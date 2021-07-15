package dcmetrohero

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TripInfo object from the dcmetrohero api
type TripInfo struct {
	FromStationName             string                    `json:"fromStationName"`
	FromStationCode             string                    `json:"fromStationCode"`
	ToStationName               string                    `json:"toStationName"`
	ToStationCode               string                    `json:"toStationCode"`
	TripStationCodes            []string                  `json:"tripStationCodes"`
	LineCodes                   []string                  `json:"lineCodes"`
	ExpectedRideTime            float64                   `json:"expectedRideTime"`
	PredictedRideTime           float64                   `json:"predictedRideTime"`
	TimeUntilNextTrain          float64                   `json:"timeUntilNextTrain"`
	TimeSinceLastTrain          float64                   `json:"timeSinceLastTrain"`
	FromStationTrainStatuses    []TrainPrediction         `json:"fromStationTrainStatuses"`
	MetroAlerts                 []MetroAlert              `json:"metroAlerts"`
	MetroAlertKeywords          []string                  `json:"metroAlertKeywords"`
	Tweets                      []Tweet                   `json:"tweets"`
	TweetKeywords               []string                  `json:"tweetKeywords"`
	FromStationElevatorOutages  []ElevatorEscalatorOutage `json:"fromStationElevatorOutages"`
	FromStationEscalatorOutages []ElevatorEscalatorOutage `json:"fromStationEscalatorOutages"`
	ToStationElevatorOutages    []ElevatorEscalatorOutage `json:"toStationElevatorOutages"`
	ToStationEscalatorOutages   []ElevatorEscalatorOutage `json:"toStationEscalatorOutages"`
	Date                        string                    `json:"date"`
}

// GetTripInfo makes an HTTP request to the dcmetrohero api, and returns the
// result in a TripInfo struct
func GetTripInfo(toCode string, fromCode string, token string) (*TripInfo, error) {
	url := fmt.Sprintf("https://dcmetrohero.com/api/v1/metrorail/trips/%s/%s", toCode, fromCode)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apiKey", token)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var trip TripInfo
	err = json.Unmarshal(body, &trip)
	if err != nil {
		return nil, err
	}
	return &trip, nil
}
