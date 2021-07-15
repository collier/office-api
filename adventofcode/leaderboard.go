package adventofcode

import (
	"encoding/json"
	"strconv"
)

// FlexInt type which represents value that returns from JSON API which is
// sometimes an int and sometimes a string
type FlexInt int

// UnmarshalJSON custom JSON unmarshaller method for the FlexInt type
func (fi *FlexInt) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(fi))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*fi = FlexInt(i)
	return nil
}

// Star contains data about a star earned during Advent of Code
type Star struct {
	GetStarTs FlexInt `json:"get_star_ts"`
}

// Member a participant within a leaderboard in Advent of Code
type Member struct {
	LocalScore         int                         `json:"local_score"`
	Name               string                      `json:"name"`
	Stars              int                         `json:"stars"`
	LastStarTs         FlexInt                     `json:"last_star_ts"`
	ID                 string                      `json:"id"`
	GlobalScore        int                         `json:"global_score"`
	CompletionDayLevel map[string]map[string]*Star `json:"completion_day_level"`
	CompletionDays     map[string]int              `json:"completion_days"`
}

// Leaderboard contains all data about an Advent of Code leaderboard
type Leaderboard struct {
	OwnerID string             `json:"owner_id"`
	Event   string             `json:"event"`
	Members map[string]*Member `json:"members"`
}
