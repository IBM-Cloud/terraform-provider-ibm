package database

import "time"

func isMoreThan24Hours(milliseconds int64) bool {
	twentyFourHours := 24 * time.Hour
	duration := time.Duration(milliseconds) * time.Millisecond

	return duration > twentyFourHours
}

func millisecondsToISOTimestamp(milliseconds int64) string{
		// Convert to time.Time
		t := time.UnixMilli(milliseconds)
	
		// Format to ISO 8601
		isoTimestamp := t.UTC().Format(time.RFC3339)
	
		return isoTimestamp // Output: 2024-02-27T12:32:03Z
}