package database

import (
	"BlockyExtendedApi/enums"
	"BlockyExtendedApi/gafam"
	"BlockyExtendedApi/models"
	"database/sql"
)

func GetTotalRequested(days int) (error, int) {
	stmt, err := DB.Prepare("SELECT COUNT(*) AS count_blocked_queries FROM log_entries WHERE request_ts > now() - INTERVAL ? DAY;")
	defer stmt.Close()
	if err != nil {
		return err, 0
	} else {
		var countBlockedQueries int
		err := stmt.QueryRow(days).Scan(&countBlockedQueries)
		if err != nil {
			return err, 0
		}
		return nil, countBlockedQueries
	}
}

func GetTotalBlocked(days int) (error, int) {

	stmt, err := DB.Prepare("SELECT COUNT(*) AS count_blocked_queries FROM log_entries WHERE response_type = 'BLOCKED' AND request_ts > now() - INTERVAL ? DAY;")
	defer stmt.Close()
	if err != nil {
		return err, 0
	} else {
		var countBlockedQueries int
		err := stmt.QueryRow(days).Scan(&countBlockedQueries)
		if err != nil {
			return err, 0
		}
		return nil, countBlockedQueries
	}
}
func GetGafamStats() (error, []gafam.GafamStats) {
	result, err := DB.Query("SELECT question_name FROM log_entries WHERE request_ts > now() - INTERVAL 1 DAY")
	if err != nil {
		return err, nil
	}
	googleCount := 0
	AmazonCount := 0
	FacebookCount := 0
	AppleCount := 0
	MicrosoftCount := 0
	OtherCount := 0
	for result.Next() {
		var question = ""
		err := result.Scan(&question)
		if err != nil {
			return err, nil
		}
		var gafamName = gafam.GetGafamName(question)
		switch gafamName {
		case enums.Google:
			googleCount++
		case enums.Amazon:
			AmazonCount++
		case enums.Facebook:
			FacebookCount++
		case enums.Apple:
			AppleCount++
		case enums.Microsoft:
			MicrosoftCount++
		case enums.Other:
			OtherCount++
		}
	}
	var stats []gafam.GafamStats
	// map the gafam stats
	stats = append(stats, gafam.GafamStats{Name: "Google", Count: googleCount})
	stats = append(stats, gafam.GafamStats{Name: "Amazon", Count: AmazonCount})
	stats = append(stats, gafam.GafamStats{Name: "Facebook", Count: FacebookCount})
	stats = append(stats, gafam.GafamStats{Name: "Apple", Count: AppleCount})
	stats = append(stats, gafam.GafamStats{Name: "Microsoft", Count: MicrosoftCount})
	stats = append(stats, gafam.GafamStats{Name: "Other", Count: OtherCount})
	return nil, stats
}

func GetTopClients() (error, []models.Client) {
	var clients []models.Client
	result, err := DB.Query("SELECT client_name, COUNT(*) AS count FROM log_entries WHERE request_ts > now() - INTERVAL 1 DAY GROUP BY client_name ORDER BY count DESC ")
	if err != nil {
		return err, nil
	}
	for result.Next() {
		var name = ""
		var count = 0
		err := result.Scan(&name, &count)
		if err != nil {
			return err, nil
		}
		clients = append(clients, models.Client{Name: name, Count: count})
	}
	return nil, clients
}

func GetRecentLogs(timestamp string) (error, models.LogResponse) {
	var query = ""
	// set query to one with no timestamp if none is received
	if timestamp == "" {
		query = "SELECT client_name, response_type, question_name, duration_ms, request_ts FROM log_entries ORDER BY request_ts DESC LIMIT 50;"
	} else {
		// query = "SELECT client_name, response_type, question_name, duration_ms, request_ts FROM log_entries WHERE request_ts <= ? ORDER BY request_ts DESC LIMIT 50;"
		query = "SELECT client_name, response_type, question_name, duration_ms, request_ts FROM log_entries WHERE request_ts < ? ORDER BY request_ts DESC LIMIT 50;"
	}
	stmt, err := DB.Prepare(query)
	var result *sql.Rows
	if timestamp == "" {
		result, err = stmt.Query()
		if err != nil {
			return err, models.LogResponse{}
		}
	} else {
		result, err = stmt.Query(timestamp)
		if err != nil {
			return err, models.LogResponse{}
		}
	}

	var response = models.LogResponse{}
	var latestTimestamp = ""
	for result.Next() {
		var clientName = ""
		var responseType = ""
		var question = ""
		var duration = ""
		var RequestTs = ""
		err := result.Scan(&clientName, &responseType, &question, &duration, &RequestTs)
		if err != nil {
			return err, models.LogResponse{}
		}
		response.Logs = append(response.Logs, models.Log{ClientName: clientName, Duration: 0, ResponseType: responseType, Question: question})
		latestTimestamp = RequestTs
	}

	defer stmt.Close()
	if err != nil {
		return err, models.LogResponse{}
	}
	response.LastTimestamp = latestTimestamp
	return nil, response
}
