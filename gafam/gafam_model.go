package gafam

type GafamData struct {
	Name    string   `json:"name"`
	Domains []string `json:"domains"`
}
type GafamStats struct {
	Name  string
	Count int
}
