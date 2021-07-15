package espn

// ScheduleItemMatchup contains the data for a given period's head-to-head
// matchup
type ScheduleItemMatchup struct {
	MatchupTypeID  int       `json:"matchupTypeId"`
	AwayTeamScores []float64 `json:"awayTeamScores"`
	AwayTeam       struct {
		WaiverRank int `json:"waiverRank"`
		Division   struct {
			DivisionName string `json:"divisionName"`
			DivisionID   int    `json:"divisionId"`
			Size         int    `json:"size"`
		} `json:"division"`
		TeamAbbrev   string `json:"teamAbbrev"`
		TeamNickname string `json:"teamNickname"`
		LogoURL      string `json:"logoUrl"`
		TeamLocation string `json:"teamLocation"`
		TeamID       int    `json:"teamId"`
		LogoType     string `json:"logoType"`
	} `json:"awayTeam"`
	AwayTeamAdjustment float64   `json:"awayTeamAdjustment"`
	AwayTeamID         int       `json:"awayTeamId"`
	IsBye              bool      `json:"isBye"`
	HomeTeamID         int       `json:"homeTeamId"`
	HomeTeamAdjustment float64   `json:"homeTeamAdjustment"`
	HomeTeamScores     []float64 `json:"homeTeamScores"`
	HomeTeamBonus      float64   `json:"homeTeamBonus"`
	Outcome            int       `json:"outcome"`
	HomeTeam           struct {
		WaiverRank int `json:"waiverRank"`
		Division   struct {
			DivisionName string `json:"divisionName"`
			DivisionID   int    `json:"divisionId"`
			Size         int    `json:"size"`
		} `json:"division"`
		TeamAbbrev   string `json:"teamAbbrev"`
		TeamNickname string `json:"teamNickname"`
		LogoURL      string `json:"logoUrl"`
		TeamLocation string `json:"teamLocation"`
		TeamID       int    `json:"teamId"`
		LogoType     string `json:"logoType"`
	} `json:"homeTeam"`
}
