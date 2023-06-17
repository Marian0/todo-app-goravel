package helpers

import (
	"database/sql"
	"time"
)

func FormatTimeToISO(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z07:00")
}

func FormatNullTimeToISO(nullTime sql.NullTime) string {
	if nullTime.Valid {
		return FormatTimeToISO(nullTime.Time)
	}
	return ""
}
