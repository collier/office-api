package espn

// LeagueMember contains the results of the getDashboardData API request
type LeagueMember struct {
	IsLeagueManager bool   `json:"isLeagueManager"`
	LastName        string `json:"lastName"`
	InviteID        int    `json:"inviteId"`
	IsLeagueCreator bool   `json:"isLeagueCreator"`
	UserProfileID   int    `json:"userProfileId"`
	UserName        string `json:"userName"`
	DisplayName     string `json:"displayName,omitempty"`
	FirstName       string `json:"firstName"`
}
