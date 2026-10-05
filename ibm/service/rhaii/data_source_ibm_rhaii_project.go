// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"maps"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/resourcecontroller"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceIBMRhaiiProject reads an existing Red Hat AI Inference (RHAII) project
// by name or by ID (guid).
func DataSourceIBMRhaiiProject() *schema.Resource {
	// Clone to avoid mutating the shared resource-controller schema map.
	rhaiiSchema := maps.Clone(resourcecontroller.DataSourceIBMResourceInstance().Schema)

	// The service is always "instructlab". The lookup filters on it, so it is
	// not user configurable.
	rhaiiSchema["service"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The service type of the instance. Always `instructlab`.",
	}

	identifier := rhaiiSchema["identifier"]
	rhaiiSchema["identifier"] = &schema.Schema{
		Type:          schema.TypeString,
		Optional:      true,
		ExactlyOneOf:  identifier.ExactlyOneOf,
		ConflictsWith: []string{"resource_group_id", "name", "location"},
		Description:   "The ID (guid) of the project.",
	}

	maps.Copy(rhaiiSchema, rhaiiProjectSchema())

	return &schema.Resource{
		Read:   dataSourceIBMRhaiiProjectRead,
		Schema: rhaiiSchema,
	}
}

func dataSourceIBMRhaiiProjectRead(d *schema.ResourceData, meta interface{}) error {
	// Restrict a lookup by name to RHAII instances only.
	if _, ok := d.GetOk("name"); ok {
		if err := d.Set("service", rhaiiServiceName); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting service: %s", err)
		}
	}
	if err := resourcecontroller.DataSourceIBMResourceInstanceRead(d, meta); err != nil {
		return err
	}
	return setRhaiiProjectAttributes(d)
}
