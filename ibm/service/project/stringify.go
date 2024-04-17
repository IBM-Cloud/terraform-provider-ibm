// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project

import (
	"encoding/json"
)

func stringify(v interface{}) (result *string) {
	if s, ok := v.(string); ok {
		result = &s
	} else {
		bytes, err := json.Marshal(v)
		if err == nil {
			s := string(bytes)
			result = &s
		} else {
			s := err.Error()
			result = &s
		}
	}
	return
}
