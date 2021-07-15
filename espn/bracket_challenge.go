package espn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// GroupMember contains the data for a single user's entry into the ESPN
// Bracket Challenege
type GroupMember struct {
	Percentile          float64 `json:"percentile"`
	Tied                bool    `json:"isTied"`
	Bracket             string  `json:"bracket"`
	MaxPoints           int     `json:"maxPoints"`
	Points              int     `json:"points"`
	DisplayName         string  `json:"displayName"`
	Rank                int     `json:"rank"`
	BracketName         string  `json:"bracketName"`
	ID                  int     `json:"id"`
	UserName            string  `json:"userName"`
	TeamChampionName    string  `json:"teamChampionName"`
	TeamChampionLogoURL string  `json:"teamChampionLogoUrl"`
}

type bracketChallenge struct {
	ID    int64 `json:"p"`
	Group struct {
		Members []struct {
			Percentile  float64 `json:"pct"`
			Tied        bool    `json:"tied"`
			Bracket     string  `json:"ps"`
			MaxPoints   int     `json:"max"`
			Points      int     `json:"p"`
			DisplayName string  `json:"n_d"`
			Rank        int     `json:"r"`
			BracketName string  `json:"n_e"`
			ID          int     `json:"id"`
			UserName    string  `json:"n_m"`
		} `json:"e"`
	} `json:"g"`
}

var teams = [...]string{"Duke", "North Dakota St", "VCU", "UCF",
	"Mississippi St", "Liberty", "Virginia Tech", "Saint Louis", "Maryland",
	"Belmont", "LSU", "Yale", "Louisville", "Minnesota", "Michigan State",
	"Bradley", "Gonzaga", "F. Dickinson", "Syracuse", "Baylor", "Marquette",
	"Murray State", "Florida St", "Vermont", "Buffalo", "Arizona State",
	"Texas Tech", "N Kentucky", "Nevada", "Florida", "Michigan", "Montana",
	"Virginia", "Gardner-Webb", "Ole Miss", "Oklahoma", "Wisconsin",
	"Oregon", "Kansas State", "UC Irvine", "Villanova", "Saint Mary's",
	"Purdue", "Old Dominion", "Cincinnati", "Iowa", "Tennessee", "Colgate",
	"North Carolina", "Iona", "Utah State", "Washington", "Auburn",
	"New Mexico St", "Kansas", "Northeastern", "Iowa State", "Ohio State",
	"Houston", "Georgia State", "Wofford", "Seton Hall", "Kentucky",
	"Abil Christian"}

var eids = [...]string{"150", "2449", "2670", "2116", "344", "2335", "259",
	"139", "120", "2057", "99", "43", "97", "135", "127", "71", "2250",
	"161", "183", "239", "269", "93", "52", "261", "2084", "9", "2641",
	"94", "2440", "57", "130", "149", "258", "2241", "145", "201", "275",
	"2483", "2306", "300", "222", "2608", "2509", "295", "2132", "2294",
	"2633", "2142", "153", "314", "328", "264", "2", "166", "2305", "111",
	"66", "194", "248", "2247", "2747", "2550", "96", "2000"}

// GetBracketChallenge makes an HTTP request to the ESPN March Madness group
// bracket challenge, and returns the list of members in the group.
func GetBracketChallenge(groupID string, year string) ([]*GroupMember, error) {
	urlRaw := "http://fantasy.espncdn.com/tournament-challenge-bracket/%s/en/api/group?groupID=%s&sort=-1&start=0&length=50&periodPoints=true"
	url := fmt.Sprintf(urlRaw, year, groupID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var bc bracketChallenge
	err = json.Unmarshal(body, &bc)
	if err != nil {
		return nil, err
	}
	members := make([]*GroupMember, 0)
	for _, m := range bc.Group.Members {
		b := strings.Split(m.Bracket, "|")
		rawID := b[len(b)-1]
		id, err := strconv.Atoi(rawID)
		champName := teams[id-1]
		eid := eids[id-1]
		logoURLTmpl := "http://a.espncdn.com/combiner/i?img=/i/teamlogos/ncaa/500/%s.png&w=80&h=80&scale=crop"
		logoURL := fmt.Sprintf(logoURLTmpl, eid)
		if err != nil {
			return nil, err
		}
		gm := GroupMember{
			m.Percentile,
			m.Tied,
			m.Bracket,
			m.MaxPoints,
			m.Points,
			m.DisplayName,
			m.Rank,
			m.BracketName,
			m.ID,
			m.UserName,
			champName,
			logoURL,
		}
		members = append(members, &gm)
	}
	return members, nil
}
