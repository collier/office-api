package adventofcode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetLeaderboard makes an HTTP request to the Advent of Code API, and gets the
// leaderboard information.
func GetLeaderboard(leaderboardID string, year string, session string) (*Leaderboard, error) {
	urlRaw := "https://adventofcode.com/%s/leaderboard/private/view/%s.json"
	url := fmt.Sprintf(urlRaw, year, leaderboardID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	c := http.Cookie{
		Name:  "session",
		Value: session,
	}
	req.AddCookie(&c)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var lb Leaderboard
	err = json.Unmarshal(body, &lb)
	if err != nil {
		return nil, err
	}
	return &lb, nil
}
