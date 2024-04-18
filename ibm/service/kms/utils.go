package kms

import (
	"time"

	"github.com/go-openapi/strfmt"
)

func DateTimeToRFC3339(dt *strfmt.DateTime) (s string) {
	if dt != nil {
		s = time.Time(*dt).Format(time.RFC3339)
	}
	return
}
