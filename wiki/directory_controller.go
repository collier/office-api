package wiki

import (
	"log"
	"net/http"
	"time"

	"github.com/collier/office-api/config"

	"github.com/gin-gonic/gin"
)

// GetStaffStatsCtrl gin controller which calls the company wiki API, and returns
// stats which are calculated from the results.
func GetStaffStatsCtrl(c *gin.Context) {
	staff, err := GetActiveStaff(config.WikiUser, config.WikiPass)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	// Below line simulates day with a lot of anniversaries
	// now, _ := time.Parse("Jan 02", "Aug 08")
	// Below line simulates day with a lot of birthdays
	// now, _ := time.Parse("Jan 02", "Feb 14")
	today := now.Format("Jan 02")
	staffBDays := make([]StaffMember, 0)
	anniversaries := make([]StaffMember, 0)
	for i := range staff {
		if today == staff[i].Birthday {
			staffBDays = append(staffBDays, staff[i])
		}
		startDate := staff[i].StartDate.Format("Jan 02")
		if today == startDate {
			anniversaries = append(anniversaries, staff[i])
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"numberOfStaff":  len(staff),
		"longestTenure":  staff[3], // Don't count Bill, Steve, or Kevin
		"staffBirthdays": staffBDays,
		"anniversaries":  anniversaries,
	})
}

// GetStaffCtrl gin controller which calls the company wiki API, and returns list
// of users from counterpoint directory
func GetStaffCtrl(c *gin.Context) {
	staff, err := GetActiveStaff(config.WikiUser, config.WikiPass)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, staff)
}
