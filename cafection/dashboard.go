package cafection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type getDashboardDataReq struct {
	SerialNumber string `json:"serialNumber"`
}

type getDashboardDataRes struct {
	Content          string `json:"content"`
	Success          bool   `json:"success"`
	ShowErrorMessage bool   `json:"showErrorMessage"`
	ErrorMessage     string `json:"errorMessage"`
	Exception        string `json:"exception"`
}

// DashboardData contains the results of the getDashboardData API request
type DashboardData struct {
	GaugeList []struct {
		GaugeCurrentPercent int    `json:"gaugeCurrentPercent"`
		GaugeTreshold       int    `json:"gaugeTreshold"`
		GaugeName           string `json:"gaugeName"`
	} `json:"gaugeList"`
	MaintenanceComponentList []struct {
		Name                   string `json:"name"`
		TresholdQantityRequire int    `json:"tresholdQantityRequire"`
		CurrentCounter         int    `json:"currentCounter"`
		LifeCounter            int    `json:"lifeCounter"`
	} `json:"maintenanceComponentList"`
	MachineState           string        `json:"machineState"`
	ErrorLevel             string        `json:"errorLevel"`
	MachineTemperature     int           `json:"machineTemperature"`
	MainControllerFirmware string        `json:"mainControllerFirmware"`
	ThermoFirmware         string        `json:"thermoFirmware"`
	CurrentDateTime        string        `json:"currentDateTime"`
	WasteBinCounter        int           `json:"wasteBinCounter"`
	WasteBinMaxCounter     int           `json:"wasteBinMaxCounter"`
	RinseCounter           int           `json:"rinseCounter"`
	RinseMaxCounter        int           `json:"rinseMaxCounter"`
	IsFreeVendMode         bool          `json:"isFreeVendMode"`
	CurrentFreeVendUser    int           `json:"currentFreeVendUser"`
	LastTimeEnterSetupMode string        `json:"lastTimeEnterSetupMode"`
	SalesCounter           int           `json:"salesCounter"`
	SalesCash              float64       `json:"salesCash"`
	LastUpTimeDate         string        `json:"lastUpTimeDate"`
	FamilyNames            interface{}   `json:"familyNames"`
	Recipies               interface{}   `json:"recipies"`
	ActualWarnings         []interface{} `json:"actualWarnings"`
}

// GetDashboardData takes the serial number of a coffee machine, and an
// authentication token, and makes an HTTP request to get the dashboard data
// associated with the target coffee machine. It returns the result in the form
// of the DashboardData.
func GetDashboardData(serialNum string, token string) (*DashboardData, error) {
	url := "https://mngtool.cafection.com/mng3/api/Brewer/getDashboardData"
	params := getDashboardDataReq{
		SerialNumber: serialNum,
	}
	reqJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", token))
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var dashRes getDashboardDataRes
	err = json.Unmarshal(body, &dashRes)
	if err != nil {
		return nil, err
	}
	var dash DashboardData
	err = json.Unmarshal([]byte(dashRes.Content), &dash)
	return &dash, nil
}
