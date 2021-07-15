package cafection

import (
	"log"
	"net/http"

	"github.com/collier/office-api/config"

	"github.com/gin-gonic/gin"
)

// GetDrinkStats gin controller which calls the cafection API, and returns
// information about the number of cups of coffe made by the machine.
func GetDrinkStats(c *gin.Context) {
	authRes, err := Auth(config.CoffeeUser, config.CoffeePass)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dash, err := GetDashboardData(config.CoffeeSerial, authRes.Token)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	compList := dash.MaintenanceComponentList
	var totalCount = 0
	for i := range compList {
		if compList[i].Name == "brewer" {
			totalCount = compList[i].LifeCounter
		}
	}
	// TODO: Fix the hardcoded laCroix drink count, probably refactor the
	// slide.
	c.JSON(http.StatusOK, gin.H{
		"coffeeToday":   dash.SalesCounter,
		"coffeeAllTime": totalCount,
		"snackCounter": gin.H{
			"id":               1,
			"laCroixCansDrunk": "2064",
		},
	})
}
