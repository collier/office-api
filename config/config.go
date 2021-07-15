package config

import (
	"encoding/json"
	"io/ioutil"
)

// Conf is a pointer to the populated Config struct
var (
	CoffeeUser                         string
	CoffeePass                         string
	CoffeeSerial                       string
	WikiUser                           string
	WikiPass                           string
	DBSchema                           string
	DBUser                             string
	DBPass                             string
	DBAddr                             string
	APIPort                            string
	AdventOfCodeSession                string
	AdventOfCodeLeaderboardID          string
	DarkSkyAPIKey                      string
	OfficeLongitude                    string
	OfficeLatitude                     string
	DCMetroHeroAPIKey                  string
	ESPNFantasyLeagueID                string
	ESPNFantasySeasonID                string
	ESPNMarchMadnessYear               string
	ESPNMarchMadnessGroupID            string
	SlackAppInclementWeatherWebhookURL string
	CompanyWikiURL                     string
)

// InitConfig reads in the config.json file, and initializes the Conf property
// with the populated Config struct
func InitConfig() error {
	type config struct {
		CoffeeUser                         string `json:"coffeeUser"`
		CoffeePass                         string `json:"coffeePass"`
		CoffeeSerial                       string `json:"coffeeSerial"`
		WikiUser                           string `json:"wikiUser"`
		WikiPass                           string `json:"wikiPass"`
		DBSchema                           string `json:"dbSchema"`
		DBUser                             string `json:"dbUser"`
		DBPass                             string `json:"dbPass"`
		DBAddr                             string `json:"dbAddr"`
		APIPort                            string `json:"apiPort"`
		AdventOfCodeSession                string `json:"adventOfCodeSession"`
		AdventOfCodeLeaderboardID          string `json:"adventOfCodeLeaderboardId"`
		DarkSkyAPIKey                      string `json:"darkSkyApiKey"`
		OfficeLongitude                    string `json:"OfficeLongitude"`
		OfficeLatitude                     string `json:"OfficeLatitude"`
		DCMetroHeroAPIKey                  string `json:"dcMetroHeroApiKey"`
		ESPNFantasyLeagueID                string `json:"espnFantasyLeagueId"`
		ESPNFantasySeasonID                string `json:"espnFantasySeasonId"`
		ESPNMarchMadnessYear               string `json:"espnMarchMadnessYear"`
		ESPNMarchMadnessGroupID            string `json:"espnMarchMadnessGroupID"`
		SlackAppInclementWeatherWebhookURL string `json:"slackAppInclementWeatherWebhookUrl"`
		CompanyWikiURL                     string `json:"companyWikiUrl"`
	}
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return err
	}
	var c config
	err = json.Unmarshal(file, &c)
	if err != nil {
		return err
	}
	CoffeeUser = c.CoffeeUser
	CoffeePass = c.CoffeePass
	CoffeeSerial = c.CoffeeSerial
	WikiUser = c.WikiUser
	WikiPass = c.WikiPass
	DBSchema = c.DBSchema
	DBUser = c.DBUser
	DBPass = c.DBPass
	DBAddr = c.DBAddr
	APIPort = c.APIPort
	AdventOfCodeSession = c.AdventOfCodeSession
	AdventOfCodeLeaderboardID = c.AdventOfCodeLeaderboardID
	DarkSkyAPIKey = c.DarkSkyAPIKey
	OfficeLongitude = c.OfficeLongitude
	OfficeLatitude = c.OfficeLatitude
	DCMetroHeroAPIKey = c.DCMetroHeroAPIKey
	ESPNFantasyLeagueID = c.ESPNFantasyLeagueID
	ESPNFantasySeasonID = c.ESPNFantasySeasonID
	ESPNMarchMadnessYear = c.ESPNMarchMadnessYear
	ESPNMarchMadnessGroupID = c.ESPNMarchMadnessGroupID
	SlackAppInclementWeatherWebhookURL = c.SlackAppInclementWeatherWebhookURL
	CompanyWikiURL = c.CompanyWikiURL
	return nil
}
