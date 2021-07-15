package sqldb

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GetPetOfTheMonthCtrl gin controller which returns the current month's pet of the
// month
func GetPetOfTheMonthCtrl(c *gin.Context) {
	potm, err := GetPetOfTheMonth()
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, potm)
}

// GetAllPetsOfTheMonthCtrl gin controller which returns the 12 latestet entries for pet
// of the month
func GetAllPetsOfTheMonthCtrl(c *gin.Context) {
	potm, err := GetAllPetsOfTheMonth()
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, potm)
}

// AddPetOfTheMonthCtrl gin controller which adds a new pet of the month
func AddPetOfTheMonthCtrl(c *gin.Context) {
	var form PetOfTheMonth
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := AddPetOfTheMonth(&form)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, form)
}

// UpdatePetOfTheMonthCtrl gin controller which updates an existing pet of the month
func UpdatePetOfTheMonthCtrl(c *gin.Context) {
	var form PetOfTheMonth
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := UpdatePetOfTheMonth(&form)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, form)
}

// DeletePetOfTheMonthCtrl gin controller which deletes an existing pet of the month
func DeletePetOfTheMonthCtrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	err = DeletePetOfTheMonthByID(id)
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, true)
}
