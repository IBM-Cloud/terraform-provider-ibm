// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import "time"

/*  TODO Move other helper functions here */
type TimeoutHelper struct {
	Now time.Time
}

func (t *TimeoutHelper) isMoreThan24Hours(duration time.Duration) bool {
	return duration > 24*time.Hour
}

func (t *TimeoutHelper) durationToISO8601(duration time.Duration) string {
	return t.Now.Add(duration).Format(time.RFC3339) // TODO Should it be UTC??
}

func (t *TimeoutHelper) calculateExpirationDatetime(timeoutDuration time.Duration) string {
	if t.isMoreThan24Hours(timeoutDuration) {
		return t.durationToISO8601(24 * time.Hour)
	}

	return t.durationToISO8601(timeoutDuration)
}
