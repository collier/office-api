package dcmetrohero

// MetroAlert object from the dcmetrohero api
type MetroAlert struct {
	Description  string   `json:"description"`
	StationCodes []string `json:"stationCodes"`
	LineCodes    []string `json:"lineCodes"`
	Keywords     []string `json:"keywords"`
	Date         string   `json:"date"`
}
