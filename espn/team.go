package espn

// Team contains the results of the getDashboardData API request
type Team struct {
	DivisionStanding int     `json:"divisionStanding"`
	Percentile       float64 `json:"percentile"`
	OverallStanding  int     `json:"overallStanding"`
	TeamAbbrev       string  `json:"teamAbbrev"`
	WaiverRank       int     `json:"waiverRank"`
	TeamNickname     string  `json:"teamNickname"`
	LogoURL          string  `json:"logoUrl"`
	LogoType         string  `json:"logoType"`
	TeamID           int     `json:"teamId"`
	TeamTransactions struct {
		MoveToIR                  int     `json:"moveToIR"`
		MiscTeamCharges           int     `json:"miscTeamCharges"`
		AcquisitionBudgetSpent    float64 `json:"acquisitionBudgetSpent"`
		AmountPaid                float64 `json:"amountPaid"`
		Drops                     int     `json:"drops"`
		MoveToActive              int     `json:"moveToActive"`
		Trades                    int     `json:"trades"`
		OverallAcquisitionTotal   float64 `json:"overallAcquisitionTotal"`
		OffseasonAcquisitionTotal float64 `json:"offseasonAcquisitionTotal"`
	} `json:"teamTransactions"`
	ScheduleItems []struct {
		Matchups        []ScheduleItemMatchup `json:"matchups"`
		MatchupPeriodID int                   `json:"matchupPeriodId"`
	} `json:"scheduleItems"`
	Record struct {
		AwayPercentage     float64 `json:"awayPercentage"`
		DivisionStanding   int     `json:"divisionStanding"`
		OverallStanding    int     `json:"overallStanding"`
		DivisionLosses     int     `json:"divisionLosses"`
		HomePercentage     float64 `json:"homePercentage"`
		AwayTies           int     `json:"awayTies"`
		DivisionWins       int     `json:"divisionWins"`
		StreakType         int     `json:"streakType"`
		OverallTies        int     `json:"overallTies"`
		HomeTies           int     `json:"homeTies"`
		HomeWins           int     `json:"homeWins"`
		DivisionTies       int     `json:"divisionTies"`
		OverallPercentage  float64 `json:"overallPercentage"`
		OverallWins        int     `json:"overallWins"`
		OverallLosses      int     `json:"overallLosses"`
		StreakLength       int     `json:"streakLength"`
		PointsAgainst      float64 `json:"pointsAgainst"`
		AwayWins           int     `json:"awayWins"`
		DivisionPercentage float64 `json:"divisionPercentage"`
		HomeLosses         int     `json:"homeLosses"`
		PointsFor          float64 `json:"pointsFor"`
		AwayLosses         int     `json:"awayLosses"`
	} `json:"record"`
	Rank     int `json:"rank"`
	Division struct {
		DivisionName string `json:"divisionName"`
		DivisionID   int    `json:"divisionId"`
		Size         int    `json:"size"`
	} `json:"division"`
	TeamLocation string `json:"teamLocation"`
	Owners       []struct {
		LastName      string `json:"lastName"`
		PrimaryOwner  bool   `json:"primaryOwner"`
		LeagueManager bool   `json:"leagueManager"`
		Joined        bool   `json:"joined"`
		InviteID      int    `json:"inviteId"`
		OwnerID       int    `json:"ownerId"`
		UserProfileID int    `json:"userProfileId"`
		PhotoURL      string `json:"photoUrl"`
		FirstName     string `json:"firstName"`
	} `json:"owners"`
}
