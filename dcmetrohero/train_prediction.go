package dcmetrohero

// TrainPrediction object from the dcmetrohero api
type TrainPrediction struct {
	TrainID                        string  `json:"trainId"`
	Car                            string  `json:"Car"`
	Destination                    string  `json:"Destination"`
	DestinationCode                string  `json:"DestinationCode"`
	DestinationName                string  `json:"DestinationName"`
	Group                          string  `json:"Group"`
	Line                           string  `json:"Line"`
	LocationCode                   string  `json:"LocationCode"`
	LocationName                   string  `json:"LocationName"`
	Min                            string  `json:"Min"`
	MinutesAway                    float64 `json:"minutesAway"`
	DirectionNumber                int     `json:"directionNumber"`
	IsScheduled                    bool    `json:"isScheduled"`
	MaxMinutesAway                 float64 `json:"maxMinutesAway"`
	NumPositiveTags                int     `json:"numPositiveTags"`
	NumNegativeTags                int     `json:"numNegativeTags"`
	TrackNumber                    int     `json:"trackNumber"`
	TrackCircuitID                 int     `json:"trackCircuitId"`
	CurrentStationCode             string  `json:"currentStationCode"`
	CurrentStationName             string  `json:"currentStationName"`
	PreviousStationCode            string  `json:"PreviousStationCode"`
	PreviousStationName            string  `json:"previousStationName"`
	ShouldRenderOnLeft             bool    `json:"ShouldRenderOnLeft"`
	SecondsSinceLastMoved          int     `json:"secondsSinceLastMoved"`
	IsCurrentlyHoldingOrSlow       bool    `json:"isCurrentlyHoldingOrSlow"`
	DelayedCount                   int     `json:"delayedCount"`
	SecondsOffSchedule             int     `json:"secondsOffSchedule"`
	IsNotOnRevenueTrack            bool    `json:"isNotOnRevenueTrack"`
	IsKeyedDown                    bool    `json:"isKeyedDown"`
	WasKeyedDown                   bool    `json:"wasKeyedDown"`
	ParentMin                      string  `json:"parentMin"`
	RawTrackCircuitID              int     `json:"rawTrackCircuitId"`
	DistanceFromNextStation        int     `json:"distanceFromNextStation"`
	TripID                         string  `json:"tripId"`
	DestinationStationAbbreviation string  `json:"destinationStationAbbreviation"`
	EstimatedMinutesAway           float64 `json:"estimatedMinutesAway"`
	ObservedDate                   string  `json:"observedDate"`
}
