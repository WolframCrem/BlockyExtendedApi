package database

import (
	"database/sql"
	"time"
)

type LogEntries struct {
	RequestTs     time.Time      `gorm:"primary_key;column:request_ts;type:datetime;"`
	ClientIP      sql.NullString `gorm:"column:client_ip;type:text;size:4294967295;"`
	ClientName    sql.NullString `gorm:"column:client_name;type:varchar;size:191;"`
	DurationMs    sql.NullInt64  `gorm:"column:duration_ms;type:bigint;"`
	Reason        sql.NullString `gorm:"column:reason;type:text;size:4294967295;"`
	ResponseType  sql.NullString `gorm:"column:response_type;type:varchar;size:191;"`
	QuestionType  sql.NullString `gorm:"column:question_type;type:text;size:4294967295;"`
	QuestionName  sql.NullString `gorm:"column:question_name;type:text;size:4294967295;"`
	EffectiveTldp sql.NullString `gorm:"column:effective_tldp;type:text;size:4294967295;"`
	Answer        sql.NullString `gorm:"column:answer;type:text;size:4294967295;"`
	ResponseCode  sql.NullString `gorm:"column:response_code;type:text;size:4294967295;"`
}
