package dcmetrohero

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// SystemMetrics object from the dcmetrohero api
type SystemMetrics struct {
	LineMetricsByLine struct {
		RD LineMetrics `json:"RD"`
		OR LineMetrics `json:"OR"`
		SV LineMetrics `json:"SV"`
		BL LineMetrics `json:"BL"`
		YL LineMetrics `json:"YL"`
		GR LineMetrics `json:"GR"`
	} `json:"lineMetricsByLine"`
	Date string `json:"date"`
}

// GetSystemMetrics makes an HTTP request to the dcmetrohero api, and returns the
// result in a TripInfo struct
func GetSystemMetrics(token string) (*SystemMetrics, error) {
	url := "https://dcmetrohero.com/api/v1/metrorail/metrics"
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
	var sysMet SystemMetrics
	err = json.Unmarshal(body, &sysMet)
	if err != nil {
		return nil, err
	}
	return &sysMet, nil
}
