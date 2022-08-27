package gafam

import (
	"BlockyExtendedApi/enums"
	"strings"
)

func GetGafamName(domain string) enums.GafamType {
	for _, gafam := range LoadedGafam {
		for _, gafamDomain := range gafam.Domains {
			if strings.Contains(domain, gafamDomain) {
				switch gafam.Name {
				case "google":
					return enums.Google
				case "amazon":
					return enums.Amazon
				case "facebook":
					return enums.Facebook
				case "apple":
					return enums.Apple
				case "microsoft":
					return enums.Microsoft
				}
			}
		}
	}
	return enums.Other
}
