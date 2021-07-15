package dcmetrohero

// LineMetrics object from the dcmetrohero api
type LineMetrics struct {
	LineCode                                string  `json:"lineCode"`
	ExpectedNumTrains                       int     `json:"expectedNumTrains"`
	NumTrains                               int     `json:"numTrains"`
	NumEightCarTrains                       int     `json:"numEightCarTrains"`
	NumDelayedTrains                        int     `json:"numDelayedTrains"`
	NumCars                                 int     `json:"numCars"`
	AverageTrainDelay                       int     `json:"averageTrainDelay"`
	MedianTrainDelay                        int     `json:"medianTrainDelay"`
	MinimumTrainDelay                       int     `json:"minimumTrainDelay"`
	MaximumTrainDelay                       int     `json:"maximumTrainDelay"`
	AverageMinimumHeadways                  float64 `json:"averageMinimumHeadways"`
	AverageTrainFrequency                   float64 `json:"averageTrainFrequency"`
	ExpectedTrainFrequency                  float64 `json:"expectedTrainFrequency"`
	AveragePlatformWaitTime                 float64 `json:"averagePlatformWaitTime"`
	ExpectedPlatformWaitTime                float64 `json:"expectedPlatformWaitTime"`
	TrainFrequencyStatus                    string  `json:"trainFrequencyStatus"`
	PlatformWaitTimeTrendStatus             string  `json:"platformWaitTimeTrendStatus"`
	AverageHeadwayAdherence                 float64 `json:"averageHeadwayAdherence"`
	AverageScheduleAdherence                float64 `json:"averageScheduleAdherence"`
	StandardDeviationTrainFrequency         float64 `json:"standardDeviationTrainFrequency"`
	ExpectedStandardDeviationTrainFrequency float64 `json:"expectedStandardDeviationTrainFrequency"`
	DirectionMetricsByDirection             struct {
		Num1 DirecitonMetrics `json:"1"`
		Num2 DirecitonMetrics `json:"2"`
	} `json:"directionMetricsByDirection"`
	Date string `json:"date"`
}
