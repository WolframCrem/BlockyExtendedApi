package main

import (
	"BlockyExtendedApi/config"
	"BlockyExtendedApi/database"
	"BlockyExtendedApi/gafam"
	"BlockyExtendedApi/handlers"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// load gafam information
	fmt.Println(fmt.Sprintf("Loaded %d different gafam companies", len(gafam.LoadedGafam)))
	// print the port number which will call the func to load the config
	fmt.Println(fmt.Sprintf("Running HTTP Server on: 0.0.0.0:%d", config.LoadedConfig.Port))
	// connect to the database
	database.ConnectToDatabase()
	// start API server
	app := fiber.New()

	app.Get("/v1/stats/", handlers.GetStats)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.LoadedConfig.Port)))
}
