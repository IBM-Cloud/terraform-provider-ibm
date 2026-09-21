// Copyright IBM Corp. 2021, 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudant

import (
	"log"
	"maps"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/resourcecontroller"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
)

// Cloudant specific resource instance additional schema
var cloudantSchema = map[string]*schema.Schema{
	"service": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The service type of the instance",
	},
	"version": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Vendor version.",
	},
	"features": {
		Type:        schema.TypeList,
		Computed:    true,
		Description: "List of enabled optional features.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"features_flags": {
		Type:        schema.TypeList,
		Computed:    true,
		Description: "List of feature flags.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"include_data_events": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "Include data event types in events sent to IBM Cloud Activity Tracker Event Routing for the IBM Cloudant instance. By default only emitted events are of \"management\" type.",
	},
	"capacity": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "A number of blocks of throughput units. A block consists of 100 reads/sec, 50 writes/sec, and 5 global queries/sec of provisioned throughput capacity.",
	},
	"throughput": {
		Type:        schema.TypeMap,
		Computed:    true,
		Description: "Schema for detailed information about throughput capacity with breakdown by specific throughput requests classes. This is only available for IBM Cloudant Gen 1.",
		Elem: &schema.Schema{
			Type: schema.TypeInt,
		},
	},
	"enable_cors": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "Boolean value to turn CORS on and off.",
	},
	"cors_config": {
		Type:        schema.TypeList,
		Computed:    true,
		Description: "Configuration for CORS.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allow_credentials": {
					Type:        schema.TypeBool,
					Computed:    true,
					Description: "Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true.",
				},
				"origins": {
					Type:        schema.TypeList,
					Computed:    true,
					Description: "An array of strings that contain allowed origin domains. You have to specify the full URL including the protocol. It is recommended that only the HTTPS protocol is used. Subdomains count as separate domains, so you have to specify all subdomains used.",
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	},
}

func DataSourceIBMCloudant() *schema.Resource {

	// Make a Cloudant resource instance schema from the resource instance schema and Cloudant's additional entries
	resourceInstanceSchema := resourcecontroller.DataSourceIBMResourceInstance().Schema

	// Clone the resource instance schema
	cloudantResourceInstanceSchema := maps.Clone(resourceInstanceSchema)
	// Add Cloudant additional entries
	for key, value := range cloudantSchema {
		cloudantResourceInstanceSchema[key] = value
	}

	return &schema.Resource{
		Read:   dataSourceIBMCloudantRead,
		Schema: cloudantResourceInstanceSchema,
	}
}

func dataSourceIBMCloudantRead(d *schema.ResourceData, meta interface{}) error {
	err := resourcecontroller.DataSourceIBMResourceInstanceRead(d, meta)
	if err != nil {
		return err
	}

	err = setCloudantResourceControllerURL(d, meta)
	if err != nil {
		return err
	}

	client, tfErr := GetCloudantClientFromResource(d, meta, "ibm_cloudant", "read")
	if tfErr != nil {
		return tfErr
	}

	err = setCloudantServerInformation(client, d)
	if err != nil {
		return err
	}

	resourceInstanceToCloudant(d)

	if !isCloudantGen2PlanFrom(d) {
		err = setCloudantActivityTrackerEvents(client, d)
		if err != nil {
			return err
		}

		err = setCloudantInstanceCapacity(client, d)
		if err != nil {
			return err
		}

		err = setCloudantInstanceCors(client, d)
		if err != nil {
			return err
		}
	}

	return nil
}

func setCloudantServerInformation(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	serverInformation, err := readCloudantServerInformation(client)
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error retrieving server information: %s", err)
	}

	if serverInformation.Vendor != nil && serverInformation.Vendor.Version != nil {
		d.Set("version", serverInformation.Vendor.Version)
	}

	if serverInformation.Features != nil {
		d.Set("features", serverInformation.Features)
	}

	if serverInformation.FeaturesFlags != nil {
		d.Set("features_flags", serverInformation.FeaturesFlags)
	}
	return nil
}

func readCloudantServerInformation(client *cloudantv1.CloudantV1) (*cloudantv1.ServerInformation, error) {
	opts := client.NewGetServerInformationOptions()

	serverInformation, response, err := client.GetServerInformation(opts)
	if err != nil {
		log.Printf("[DEBUG] Error retrieving server information: %s\n%s", err, response)
	}
	return serverInformation, err
}
