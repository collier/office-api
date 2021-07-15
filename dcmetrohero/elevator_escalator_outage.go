package dcmetrohero

// ElevatorEscalatorOutage object from the dcmetrohero api
type ElevatorEscalatorOutage struct {
	StationCode         string `json:"stationCode"`
	StationName         string `json:"stationName"`
	LocationDescription string `json:"locationDescription"`
	SymptomDescription  string `json:"symptomDescription"`
	UnitName            string `json:"unitName"`
	UnitType            string `jso/n:"unitType"`
	OutOfServiceDate    string `json:"outOfServiceDate"`
	UpdatedDate         string `json:"updatedDate"`
}
