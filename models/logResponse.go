package models

type Log struct {
	ClientName   string
	ResponseType string
	Question     string
	Duration     int
}

type LogResponse struct {
	Logs          []Log
	LastTimestamp string
}
