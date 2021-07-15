package main

import (
	"log"
	"net/http"
	"os"

	"github.com/collier/office-api/adventofcode"
	"github.com/collier/office-api/cafection"
	"github.com/collier/office-api/config"
	"github.com/collier/office-api/darksky"
	"github.com/collier/office-api/dcmetrohero"
	"github.com/collier/office-api/espn"
	"github.com/collier/office-api/files"
	"github.com/collier/office-api/sqldb"
	"github.com/collier/office-api/wiki"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup error log file for production
	mode := gin.Mode()
	if mode == gin.ReleaseMode {
		logf, err := os.OpenFile("./office-api.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer logf.Close()
		log.SetOutput(logf)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Load configuration file
	err := config.InitConfig()
	if err != nil {
		log.Fatalf("ERROR %v", err)
	}

	// Initialize database
	err = sqldb.InitDB(config.DBUser, config.DBPass, config.DBAddr, config.DBSchema)
	if err != nil {
		log.Fatalf("ERROR %v", err)
	}

	startRestAPI()
}

func startRestAPI() {
	// Initialize and start Gin web server
	r := gin.Default()

	// Define API routes
	api := r.Group("api")
	{
		// Advent of Code routes
		api.GET("/aoc-leaderboard", adventofcode.GetLeaderboardCtrl)

		// Cafection routes
		api.GET("/drink-stats", cafection.GetDrinkStats)

		// Darksky routes
		api.GET("/weather", darksky.GetWeatherCtrl)

		// dcmetrohero routes
		api.GET("/metro-stats", dcmetrohero.GetMetroStatsCtrl)

		// ESPN routes
		api.GET("/fantasy-football/league", espn.GetFantasyFootballLeagueCtrl)
		api.GET("/fantasy-football/matchups", espn.GetFantasyFootballMatchupsCtrl)
		api.GET("/march-madness", espn.GetMarchMadnessChallengeCtrl)

		// Local file system routes
		api.POST("/pet-image", files.UploadPetImageCtrl)
		api.GET("/icons", files.GetIconsCtrl)

		// Counterpoint Wiki routes
		api.GET("/staff-stats", wiki.GetStaffStatsCtrl)
		api.GET("/staff", wiki.GetStaffCtrl)

		// Database routes
		api.GET("/clients", sqldb.GetClientsCtrl)
		api.GET("/competitions", sqldb.GetCompetitionsCtrl)
		api.POST("/competitions", sqldb.AddCompetitionCtrl)
		api.PUT("/competitions/:id", sqldb.UpdateCompetitionCtrl)
		api.DELETE("/competitions/:id", sqldb.DeleteCompetitionCtrl)
		api.GET("/events", sqldb.GetEventsCtrl)
		api.POST("/events", sqldb.AddEventCtrl)
		api.PUT("/events/:id", sqldb.UpdateEventCtrl)
		api.DELETE("/events/:id", sqldb.DeleteEventCtrl)
		api.GET("/pet-of-the-month", sqldb.GetPetOfTheMonthCtrl)
		api.GET("/pets-of-the-month", sqldb.GetAllPetsOfTheMonthCtrl)
		api.POST("/pets-of-the-month", sqldb.AddPetOfTheMonthCtrl)
		api.PUT("/pets-of-the-month/:id", sqldb.UpdatePetOfTheMonthCtrl)
		api.DELETE("/pets-of-the-month/:id", sqldb.DeletePetOfTheMonthCtrl)
	}

	r.Static("img", "./web/img")
	r.Static("static", "./web/static")
	r.StaticFile("manifest.json", "./web/manifest.json")
	r.StaticFile("service-worker.js", "./web/service-worker.js")
	r.StaticFile("asset-manifest.json", "./web/asset-manifest.json")
	r.LoadHTMLFiles("web/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":" + config.APIPort)
}
