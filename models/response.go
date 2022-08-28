package models

import "BlockyExtendedApi/gafam"

type Stats struct {
	Total      int
	Blocked    int
	TopClients []Client
	Gafam      []gafam.GafamStats
}
