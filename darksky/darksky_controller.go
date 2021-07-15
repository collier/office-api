package darksky

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetWeatherCtrl gin controller which calls the darksky API, and returns
// weather information specific to the current counterpoint headquarters
func GetWeatherCtrl(c *gin.Context) {
	f, err := GetHQWeatherForecast("now")
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, f)
}
