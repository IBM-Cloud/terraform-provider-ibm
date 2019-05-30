package utils

import (
	"net/url"

	"github.ibm.com/riaas/rias-api/riaas/models"
)

// GetNext ...
func GetNext(next *models.Next) string {
	if next == nil {
		return ""
	}

	u, err := url.Parse(next.Href)
	if err != nil {
		return ""
	}

	q := u.Query()
	return q.Get("start")
}

// GetPageLink ...
func GetPageLink(pageLink *models.PageLink) string {
	if pageLink == nil {
		return ""
	}

	u, err := url.Parse(pageLink.Href)
	if err != nil {
		return ""
	}

	q := u.Query()
	return q.Get("start")
}
