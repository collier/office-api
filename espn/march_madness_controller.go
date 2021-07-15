package espn

import (
	"log"
	"net/http"

	"github.com/collier/office-api/config"
	"github.com/gin-gonic/gin"
)

// GetMarchMadnessChallengeCtrl gin controller which calls the ESPN fantasy
// March Madness API to get all group challenge information
func GetMarchMadnessChallengeCtrl(c *gin.Context) {
	bc, err := GetBracketChallenge(config.ESPNMarchMadnessGroupID, config.ESPNMarchMadnessYear)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bc)
}
