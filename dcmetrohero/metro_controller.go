package dcmetrohero

import (
	"log"
	"net/http"

	"github.com/collier/office-api/config"

	"github.com/gin-gonic/gin"
)

// GetMetroStatsCtrl gin controller which calls the dcmetrohero API, and returns
// metro trip information specific to the current counterpoint headquarters
func GetMetroStatsCtrl(c *gin.Context) {
	sysMet, err := GetSystemMetrics(config.DCMetroHeroAPIKey)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	greensboro, err := GetTripInfo("N03", "C01", config.DCMetroHeroAPIKey)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tysons, err := GetTripInfo("N02", "C01", config.DCMetroHeroAPIKey)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"greensboro":        greensboro,
		"tysons":            tysons,
		"silverLineMetrics": sysMet.LineMetricsByLine.SV.DirectionMetricsByDirection.Num1,
	})
}
