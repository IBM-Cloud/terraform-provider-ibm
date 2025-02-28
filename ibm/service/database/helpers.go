// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import "time"

/*  TODO Move other helper functions here */

func isMoreThan24Hours(milliseconds int64) bool {
	twentyFourHours := 24 * time.Hour
	duration := time.Duration(milliseconds) * time.Millisecond

	return duration > twentyFourHours
}

func millisecondsToISOTimestamp(milliseconds int64) string {
	t := time.UnixMilli(milliseconds)

	// Format to ISO 8601 Example: 2024-02-27T12:32:03Z
	isoTimestamp := t.UTC().Format(time.RFC3339)

	return isoTimestamp
}
