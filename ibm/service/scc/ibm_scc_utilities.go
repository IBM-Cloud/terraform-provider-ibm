// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"strings"

	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// AddSchemaData will add the data instance_id and region to the resource 
func AddSchemaData(resource *schema.Resource) *schema.Resource {
	resource.Schema["instance_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "The ID of the Security and Compliance Center instance.",
	}
	resource.Schema["region"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "The region code of the instance",
	}
	return resource
}

// getRegionData will check if the field region is defined
func getRegionData(client securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3, d *schema.ResourceData) string {
	_, ok := d.GetOk("region")
	if ok {
		return d.Get("region").(string)
	} else {
		url := client.Service.GetServiceURL()
		return strings.Split(url, ".")[1]
	}
}
