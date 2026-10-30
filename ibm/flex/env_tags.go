// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package flex

import (
	"os"
	"sort"
	"strings"
)

// EnvTagsVar holds comma-separated tags that are attached to every resource,
// on top of the configured tags.
const EnvTagsVar = "IC_ENV_TAGS"

// EnvTags returns the IC_ENV_TAGS tags, trimmed, or nil when the variable is not set.
func EnvTags() []string {
	v := os.Getenv(EnvTagsVar)
	if v == "" {
		return nil
	}
	var tags []string
	for _, t := range strings.Split(v, ",") {
		if t = strings.TrimSpace(t); t != "" {
			tags = append(tags, t)
		}
	}
	return tags
}

// IsEnvTag reports whether a tag is one of the IC_ENV_TAGS tags. The
// comparison ignores case, as the Global Tagging service lower-cases tags.
func IsEnvTag(tag string) bool {
	for _, envTag := range EnvTags() {
		if strings.EqualFold(tag, envTag) {
			return true
		}
	}
	return false
}

// TagsWithoutEnvTags returns the tags without the IC_ENV_TAGS tags.
func TagsWithoutEnvTags(tags []string) []string {
	if len(EnvTags()) == 0 {
		return tags
	}
	out := make([]string, 0, len(tags))
	for _, t := range tags {
		if !IsEnvTag(t) {
			out = append(out, t)
		}
	}
	return out
}

// ReadGlobalTagsUsingCRN returns the sorted tags of a CRN. For user tags
// (tagType "" or "user") the IC_ENV_TAGS tags are left out, so a Plugin
// Framework resource can store the result in the state without a diff.
func ReadGlobalTagsUsingCRN(meta interface{}, resourceCRN, resourceType, tagType string) ([]string, error) {
	set, err := GetGlobalTagsUsingCRN(meta, resourceCRN, resourceType, tagType)
	if err != nil {
		return nil, err
	}
	tags := make([]string, 0)
	if set != nil {
		for _, t := range set.List() {
			tags = append(tags, t.(string))
		}
	}
	if strings.TrimSpace(tagType) == "" || tagType == "user" {
		tags = TagsWithoutEnvTags(tags)
	}
	sort.Strings(tags)
	return tags, nil
}
