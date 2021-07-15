package sqldb

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GetEventsCtrl gin controller which returns all future events
func GetEventsCtrl(c *gin.Context) {
	events, err := GetFutureCompanyEvents()
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// AddEventCtrl gin controller which adds a new event
func AddEventCtrl(c *gin.Context) {
	var form CompanyEvent
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := AddCompanyEvent(&form)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, form)
}

// UpdateEventCtrl gin controller which updates an existing event
func UpdateEventCtrl(c *gin.Context) {
	var form CompanyEvent
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := UpdateCompanyEvent(&form)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, form)
}

// DeleteEventCtrl gin controller which deletes an existing event
func DeleteEventCtrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	err = DeleteCompanyEventByID(id)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, true)
}
