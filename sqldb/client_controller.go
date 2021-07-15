package sqldb

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetClientsCtrl gin controller which returns all clients
func GetClientsCtrl(c *gin.Context) {
	clients, err := GetAllClients()
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}
