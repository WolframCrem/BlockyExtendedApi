package main

import (
	"BlockyExtendedApi/config"
	"BlockyExtendedApi/gafam"
	"fmt"
)

func main() {
	// load gafam information
	fmt.Println(fmt.Sprintf("Loaded %d different gafam companies", len(gafam.LoadedGafam)))
	// print the port number which will call the func to load the config
	fmt.Println(fmt.Sprintf("Running HTTP Server on: 0.0.0.0:%d", config.LoadedConfig.Port))
}
