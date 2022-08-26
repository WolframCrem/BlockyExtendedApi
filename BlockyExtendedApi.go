package main

import (
	"BlockyExtendedApi/config"
	"BlockyExtendedApi/gafam"
	"fmt"
)

func main() {
	// load gafam information
	fmt.Println(fmt.Sprintf("Loaded %d different gafam companies", len(gafam.Gafam)))
	// load config
	config.LoadConfigData()
}
