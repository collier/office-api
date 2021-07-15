package espn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"
)

// League contains the results of the getDashboardData API request
type League struct {
	Leaguesettings struct {
		ID                                int              `json:"id"`
		Name                              string           `json:"name"`
		PlayoffHomeTeamBonus              float64          `json:"playoffHomeTeamBonus"`
		DraftStatusTypeID                 int              `json:"draftStatusTypeId"`
		PlayerAcquisitionType             int              `json:"playerAcquisitionType"`
		DraftOrderTypeID                  int              `json:"draftOrderTypeId"`
		LeagueStatusTypeID                int              `json:"leagueStatusTypeId"`
		HomeTeamBonus                     float64          `json:"homeTeamBonus"`
		RosterMoveLimit                   int              `json:"rosterMoveLimit"`
		DateDraftCompleted                time.Time        `json:"dateDraftCompleted"`
		FinalRegularSeasonMatchupPeriodID int              `json:"finalRegularSeasonMatchupPeriodId"`
		PlayoffSeedings                   []int            `json:"playoffSeedings"`
		WaiverProcessHour                 int              `json:"waiverProcessHour"`
		PlayoffTieRuleRawStatID           int              `json:"playoffTieRuleRawStatId"`
		MatchupPeriodTypeID               int              `json:"matchupPeriodTypeId"`
		RestrictionTypeID                 int              `json:"restrictionTypeId"`
		UsingUndroppableList              bool             `json:"usingUndroppableList"`
		DraftPickTradingEnabled           bool             `json:"draftPickTradingEnabled"`
		WaiverProcessDays                 int              `json:"waiverProcessDays"`
		WaiverHours                       int              `json:"waiverHours"`
		PlayoffMatchupLength              int              `json:"playoffMatchupLength"`
		VetoVotesRequired                 int              `json:"vetoVotesRequired"`
		Size                              int              `json:"size"`
		ScoringDecimalPlaces              int              `json:"scoringDecimalPlaces"`
		Teams                             map[string]*Team `json:"teams"`
		MinimumBidAmount                  float64          `json:"minimumBidAmount"`
		WaiverOrderSystemTypeID           int              `json:"waiverOrderSystemTypeId"`
		PlayoffSeedingTieRule             int              `json:"playoffSeedingTieRule"`
		TradeLimit                        int              `json:"tradeLimit"`
		AllowsTrades                      bool             `json:"allowsTrades"`
		Season                            int              `json:"season"`
		PlayoffTeamCount                  int              `json:"playoffTeamCount"`
		TimePerDraftSelection             int              `json:"timePerDraftSelection"`
		LineupLocktimeType                int              `json:"lineupLocktimeType"`
		FinalMatchupPeriodID              int              `json:"finalMatchupPeriodId"`
		RegularSeasonMatchupPeriodCount   int              `json:"regularSeasonMatchupPeriodCount"`
		AccessTypeID                      int              `json:"accessTypeId"`
		ScoringTypeID                     int              `json:"scoringTypeId"`
		FutureKeeperCount                 int              `json:"futureKeeperCount"`
		TieRule                           int              `json:"tieRule"`
		DefaultWaiverOrder                int              `json:"defaultWaiverOrder"`
		DefaultUniverseID                 int              `json:"defaultUniverseId"`
		PlayerAcquisitionBudget           float64          `json:"playerAcquisitionBudget"`
		DraftTypeID                       int              `json:"draftTypeId"`
		DraftAuctionBudget                float64          `json:"draftAuctionBudget"`
		TieRuleRawStatID                  int              `json:"tieRuleRawStatId"`
		LeagueSubTypeID                   int              `json:"leagueSubTypeId"`
		PremiumTypeID                     int              `json:"premiumTypeId"`
		LeagueFormatTypeID                int              `json:"leagueFormatTypeId"`
		TradeRevisionHours                int              `json:"tradeRevisionHours"`
		PlayoffSeedingTieRuleRawStatID    int              `json:"playoffSeedingTieRuleRawStatId"`
		DraftOrder                        []int            `json:"draftOrder"`
		IsViewable                        bool             `json:"isViewable"`
		UsesPlayerAcquisitionBudget       bool             `json:"usesPlayerAcquisitionBudget"`
		LeagueMembers                     []LeagueMember   `json:"leagueMembers"`
		FirstScoringPeriodID              int              `json:"firstScoringPeriodId"`
		RegularSeasonMatchupLength        int              `json:"regularSeasonMatchupLength"`
		PlayoffTieRule                    int              `json:"playoffTieRule"`
		AllowOutOfUniverseStatsAndTrades  bool             `json:"allowOutOfUniverseStatsAndTrades"`
		FirstMatchupPeriodID              int              `json:"firstMatchupPeriodId"`
		LeagueTypeID                      int              `json:"leagueTypeId"`
		FinalScoringPeriodID              int              `json:"finalScoringPeriodId"`
		SeasonAcquisitionLimit            int              `json:"seasonAcquisitionLimit"`
		RosterLocktimeType                int              `json:"rosterLocktimeType"`
	} `json:"leaguesettings"`
}

// Matchup convience struct which contains data for a Head-to-Head matchup for
// a given period in fantasyfootball
type Matchup struct {
	HomeTeamID int                 `json:"homeTeamId"`
	AwayTeamID int                 `json:"awayTeamId"`
	HomeTeam   *Team               `json:"homeTeam"`
	AwayTeam   *Team               `json:"awayTeam"`
	Matchup    ScheduleItemMatchup `json:"matchup"`
}

// GetFantasyLeague makes an HTTP request to the ESPN Fantasy Football api, and
// returns the result in a League struct
func GetFantasyLeague(leagueID string, seasonID string) (*League, error) {
	urlRaw := "http://games.espn.com/ffl/api/v2/leagueSettings?leagueId=%s&seasonId=%s"
	url := fmt.Sprintf(urlRaw, leagueID, seasonID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var league League
	err = json.Unmarshal(body, &league)
	if err != nil {
		return nil, err
	}
	return &league, nil
}

func getCurrentPeriod() int {
	now := time.Now()
	// now, _ := time.Parse("2006-01-02 15:04:05", "2018-09-27 00:01:00")
	sd, _ := time.Parse("2006-01-02 15:04:05", "2018-09-06 00:00:00")
	if now.After(sd) && now.Before(sd.Add(time.Hour*24*7)) {
		return 1
	} else if now.After(sd.Add(time.Hour*24*7*1)) && now.Before(sd.Add(time.Hour*24*7).Add(time.Hour*24*7*1)) {
		return 2
	} else if now.After(sd.Add(time.Hour*24*7*2)) && now.Before(sd.Add(time.Hour*24*7).Add(time.Hour*24*7*2)) {
		return 3
	} else if now.After(sd.Add(time.Hour*24*7*3)) && now.Before(sd.Add(time.Hour*24*7).Add(time.Hour*24*7*3)) {
		return 4
	} else if now.After(sd.Add(time.Hour*24*7*4)) && now.Before(sd.Add(time.Hour*24*7).Add(time.Hour*24*7*4)) {
		return 5
	} else if now.After(sd.Add(time.Hour*24*7*5)) && now.Before(sd.Add(time.Hour*24*7).Add(time.Hour*24*7*5)) {
		return 6
	} else if now.After(sd.Add(time.Hour*24*7*6)) && now.Before(sd.Add(time.Hour*24*7).Add(time.Hour*24*7*6)) {
		return 7
	} else if now.After(sd.Add(time.Hour*24*7*7)) && now.Before(sd.Add(time.Hour*24*7).Add(time.Hour*24*7*7)) {
		return 8
	} else if now.After(sd.Add(time.Hour*24*7*8)) && now.Before(sd.Add(time.Hour*24*7).Add(time.Hour*24*7*8)) {
		return 9
	} else if now.After(sd.Add(time.Hour*24*7*9)) && now.Before(sd.Add(time.Hour*24*7).Add(time.Hour*24*7*9)) {
		return 10
	} else if now.After(sd.Add(time.Hour*24*7*10)) && now.Before(sd.Add(time.Hour*24*7).Add(time.Hour*24*7*10)) {
		return 11
	} else if now.After(sd.Add(time.Hour*24*7*11)) && now.Before(sd.Add(time.Hour*24*7).Add(time.Hour*24*7*11)) {
		return 12
	} else if now.After(sd.Add(time.Hour*24*7*12)) && now.Before(sd.Add(time.Hour*24*7).Add(time.Hour*24*7*12)) {
		return 13
	}
	return 1
}

// GetCurrentMatchups makes an HTTP request to the ESPN Fantasy Football api, and
// returns all Matchups for the current week
func GetCurrentMatchups(leagueID string, seasonID string) ([]*Matchup, error) {
	league, err := GetFantasyLeague(leagueID, seasonID)
	if err != nil {
		return nil, err
	}
	periodID := getCurrentPeriod()
	mids := make(map[int]int)
	for _, team := range league.Leaguesettings.Teams {
		for _, si := range team.ScheduleItems {
			if periodID == si.MatchupPeriodID {
				mids[si.Matchups[0].HomeTeam.TeamID] = si.Matchups[0].AwayTeam.TeamID
			}
		}
	}
	matchups := make([]*Matchup, 0)
	for homeID, awayID := range mids {
		// Convert home and away IDs to strings
		homeIDStr := strconv.Itoa(homeID)
		awayIDStr := strconv.Itoa(awayID)
		homeTeam := league.Leaguesettings.Teams[homeIDStr]
		awayTeam := league.Leaguesettings.Teams[awayIDStr]
		// Get the corrosponding ScheduleItemMatchup struct for this
		// matchup from the home team's schedule items
		var sim ScheduleItemMatchup
		for index, si := range homeTeam.ScheduleItems {
			if periodID == si.MatchupPeriodID {
				sim = homeTeam.ScheduleItems[index].Matchups[0]
			}
		}
		// Remove large unused data structure
		if homeTeam != nil {
			homeTeam.ScheduleItems = nil
		}
		if awayTeam != nil {
			awayTeam.ScheduleItems = nil
		}
		mu := Matchup{
			HomeTeamID: homeID,
			AwayTeamID: awayID,
			HomeTeam:   homeTeam,
			AwayTeam:   awayTeam,
			Matchup:    sim,
		}
		matchups = append(matchups, &mu)
	}
	// Sort the matchups in ascending order by HomeTeamID
	sort.Slice(matchups, func(i, j int) bool {
		return matchups[i].HomeTeamID < matchups[j].HomeTeamID
	})
	return matchups, nil
}
