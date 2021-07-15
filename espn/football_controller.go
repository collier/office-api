package espn

import (
	"log"
	"net/http"

	"github.com/collier/office-api/config"

	"github.com/gin-gonic/gin"
)

// GetFantasyFootballLeagueCtrl gin controller which calls the ESPN fantasy
// football API to get all league information
func GetFantasyFootballLeagueCtrl(c *gin.Context) {
	league, err := GetFantasyLeague(config.ESPNFantasyLeagueID, config.ESPNFantasySeasonID)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, league)
}

// GetFantasyFootballMatchupsCtrl gin controller which calls the ESPN fantasy
// football API to get current week's matchups
func GetFantasyFootballMatchupsCtrl(c *gin.Context) {
	m, err := GetCurrentMatchups(config.ESPNFantasyLeagueID, config.ESPNFantasySeasonID)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}
