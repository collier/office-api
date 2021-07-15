package sqldb

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GetCompetitionsCtrl gin controller which returns all competitions
func GetCompetitionsCtrl(c *gin.Context) {
	comps, err := GetCompetitions()
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comps)
}

// AddCompetitionCtrl gin controller which adds a new competition
func AddCompetitionCtrl(c *gin.Context) {
	var form Competition
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := AddCompetition(&form)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, form)
}

// UpdateCompetitionCtrl gin controller which updates an existing competition
func UpdateCompetitionCtrl(c *gin.Context) {
	var form Competition
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := UpdateCompetition(&form)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, form)
}

// DeleteCompetitionCtrl gin controller which deletes an existing competition
func DeleteCompetitionCtrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = DeleteCompetitionByID(id)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, true)
}
