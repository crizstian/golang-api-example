package main

import (
	"easycast/src/api"
	"easycast/src/db"
	"easycast/src/models"
	"easycast/src/routes"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// static content
	e.Static("/v1/public", "public")

	// main route will render the documentation
	e.GET("/", api.IndexAPI)

	e.GET("/robots.txt", api.RobotsAPI)
	e.GET("/favicon.ico", api.FaviconAPI)

	// route for health check
	e.GET("/ping", api.PingAPI)

	// group for versioning this case version 1
	a := e.Group("/v1")

	// all routes are defined in the routes/index.go
	routes.ApiRoutes(a)

	// init the database
	// startDB()

	// start server
	e.Logger.Fatal(e.Start(":8000"))
}

// internal functions starts with lowercase
func startDB() {
	Db := db.New()
	defer Db.Close()
	models.AutoMigrate()
}
