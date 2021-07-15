package adventofcode

import (
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/collier/office-api/config"

	"github.com/gin-gonic/gin"
)

// GetLeaderboardCtrl gin controller which calls the Advent of Code API, and
// returns leaderboard information.
func GetLeaderboardCtrl(c *gin.Context) {
	year := time.Now().Format("2006")
	// Used for testing, earlier years had more data
	// year := "2017"
	l, err := GetLeaderboard(config.AdventOfCodeLeaderboardID, year, config.AdventOfCodeSession)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Move all members with names to an array of members, and also transform
	// the timestamp Star structs into a map of string to number of stars
	// for that day.
	members := make([]*Member, 0)
	for _, m := range l.Members {
		if m.Name != "" {
			m.CompletionDays = make(map[string]int)
			for key, cdl := range m.CompletionDayLevel {
				m.CompletionDays[key] = len(cdl)
			}
			// People who are too good to use their actual names....
			if m.ID == "379342" {
				m.Name = "Ryan Zaki"
			}
			if m.ID == "388356" {
				m.Name = "Bill Markmann"
			}
			if m.ID == "386738" {
				m.Name = "Bhanvi Dubey"
			}
			m.CompletionDayLevel = nil
			members = append(members, m)
		}
	}
	// Sort the members in descending order by LocalScore
	sort.Slice(members, func(i, j int) bool {
		if members[i].LocalScore > members[j].LocalScore {
			return true
		}
		if members[i].LocalScore < members[j].LocalScore {
			return false
		}
		return members[i].Name < members[j].Name
	})
	c.JSON(http.StatusOK, members)
}
