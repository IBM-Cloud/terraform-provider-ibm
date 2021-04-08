// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.ibm.com/ibmcloud/kubernetesservice-go-sdk/kubernetesserviceapiv1"
)

func dataSourceIBMSatelliteLocation() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMSatelliteLocationRead,

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A unique name for the new Satellite location",
			},
			"managed_from": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IBM Cloud metro from which the Satellite location is managed",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A description of the new Satellite location",
			},
			"logging_account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account ID for IBM Log Analysis with LogDNA log forwarding",
			},
			"zones": {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The names of at least three high availability zones to use for the location",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Location CRN",
			},
		},
	}
}

func dataSourceIBMSatelliteLocationRead(d *schema.ResourceData, meta interface{}) error {
	location := d.Get("location").(string)

	satClient, err := meta.(ClientSession).SatelliteClientSession()
	if err != nil {
		return err
	}

	getSatLocOptions := &kubernetesserviceapiv1.GetSatelliteLocationOptions{
		Controller: &location,
	}

	instance, resp, err := satClient.GetSatelliteLocation(getSatLocOptions)
	if err != nil || instance == nil {
		return fmt.Errorf("Error retrieving IBM Cloud satellite location %s : %s\n%s", name, err, resp)

	}

	d.SetId(*instance.ID)
	d.Set("location", location)
	d.Set("description", *instance.Description)
	d.Set("zones", instance.WorkerZones)
	d.Set("managed_from", *instance.Datacenter)
	d.Set("crn", *instance.Crn)

	return nil
}
