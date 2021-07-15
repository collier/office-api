package files

import (
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	// Used to check dimensions of uploaded jpegs
	_ "image/jpeg"
	// Used to check dimensions of uploaded pngs
	_ "image/png"
)

// UploadPetImageCtrl gin controller for uploading new images into the pet
// folder
func UploadPetImageCtrl(c *gin.Context) {
	_, header, _ := c.Request.FormFile("petImage")
	file, _ := c.FormFile("petImage")
	fileName := header.Filename
	reader, err := file.Open()
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	img, _, err := image.DecodeConfig(reader)
	if img.Height != img.Width {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Image must be square",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid file uploaded",
		})
		return
	}
	c.SaveUploadedFile(file, "./web/img/pets/"+fileName)
	c.JSON(http.StatusOK, true)
}

// GetIconsCtrl gin controller which returns an array containing all of the icons
// located inside of the icons directory within the web folder
func GetIconsCtrl(c *gin.Context) {
	var files []string
	root := "./web/img/icons"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			rootDir := filepath.FromSlash("web/")
			files = append(files, filepath.ToSlash(strings.TrimPrefix(path, rootDir)))
		}
		return nil
	})
	if err != nil {
		log.Printf("ERROR %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, files)
}
