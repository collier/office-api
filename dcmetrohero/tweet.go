package dcmetrohero

// Tweet object from the dcmetrohero api
type Tweet struct {
	TwitterID       int64    `json:"twitterId"`
	TwitterIDString string   `json:"twitterIdString"`
	UserID          int      `json:"userId"`
	Text            string   `json:"text"`
	StationCodes    []string `json:"stationCodes"`
	LineCodes       []string `json:"lineCodes"`
	Keywords        []string `json:"keywords"`
	URL             string   `json:"url"`
	Date            string   `json:"date"`
}
