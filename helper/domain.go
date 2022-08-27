package helper

import (
	"fmt"
	"strings"
)

func GetDomain(domain string) string {
	split := strings.Split(domain, ".")
	return fmt.Sprintf("%s.%s", split[len(split)-2], split[len(split)-1])
}
