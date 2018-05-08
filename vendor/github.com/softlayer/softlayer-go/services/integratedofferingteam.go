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

// This class represents an Integrated Offering Team region.
type IntegratedOfferingTeam_Region struct {
	Session *session.Session
	Options sl.Options
}

// GetIntegratedOfferingTeamRegionService returns an instance of the IntegratedOfferingTeam_Region SoftLayer service
func GetIntegratedOfferingTeamRegionService(sess *session.Session) IntegratedOfferingTeam_Region {
	return IntegratedOfferingTeam_Region{Session: sess}
}

func (r IntegratedOfferingTeam_Region) Id(id int) IntegratedOfferingTeam_Region {
	r.Options.Id = &id
	return r
}

func (r IntegratedOfferingTeam_Region) Mask(mask string) IntegratedOfferingTeam_Region {
	if !strings.HasPrefix(mask, "mask[") && (strings.Contains(mask, "[") || strings.Contains(mask, ",")) {
		mask = fmt.Sprintf("mask[%s]", mask)
	}

	r.Options.Mask = mask
	return r
}

func (r IntegratedOfferingTeam_Region) Filter(filter string) IntegratedOfferingTeam_Region {
	r.Options.Filter = filter
	return r
}

func (r IntegratedOfferingTeam_Region) Limit(limit int) IntegratedOfferingTeam_Region {
	r.Options.Limit = &limit
	return r
}

func (r IntegratedOfferingTeam_Region) Offset(offset int) IntegratedOfferingTeam_Region {
	r.Options.Offset = &offset
	return r
}
