/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package services

import (
	"fmt"
	"strings"

	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/sl"
)

// no documentation yet
type BluePages_Search struct {
	Session *session.Session
	Options sl.Options
}

// GetBluePagesSearchService returns an instance of the BluePages_Search SoftLayer service
func GetBluePagesSearchService(sess *session.Session) BluePages_Search {
	return BluePages_Search{Session: sess}
}

func (r BluePages_Search) Id(id int) BluePages_Search {
	r.Options.Id = &id
	return r
}

func (r BluePages_Search) Mask(mask string) BluePages_Search {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r BluePages_Search) Filter(filter string) BluePages_Search {
	r.Options.Filter = filter
	return r
}

func (r BluePages_Search) Limit(limit int) BluePages_Search {
	r.Options.Limit = &limit
	return r
}

func (r BluePages_Search) Offset(offset int) BluePages_Search {
	r.Options.Offset = &offset
	return r
}
