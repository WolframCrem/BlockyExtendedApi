package handlers

import (
	"BlockyExtendedApi/database"
	"BlockyExtendedApi/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func GetStats(c *fiber.Ctx) error {

	// get blocked requests
	var err, totalCount = database.GetTotalRequested(1)
	if err != nil {
		return err
	}

	// get total requests
	err, blockedCount := database.GetTotalBlocked(1)
	if err != nil {
		return err
	}

	// get top clients
	err, topClients := database.GetTopClients()
	if err != nil {
		return err
	}

	// get gafam stats
	err, gafamStats := database.GetGafamStats()
	if err != nil {
		return err
	}

	var response = models.Stats{Total: totalCount, Blocked: blockedCount, TopClients: topClients, Gafam: gafamStats}
	responseJson, err := json.Marshal(response)
	if err != nil {
		return err
	}
	err = c.SendString(string(responseJson))
	if err != nil {
		return err
	}
	return nil
}
