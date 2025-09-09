// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package appconfiguration

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func DataSourceIBMAppConfigIntegrationEn() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmAppConfigIntegrationEnRead,

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"integration_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of integration",
			},
			"integration_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Integration type [will be EVENT_NOTIFICATIONS always]",
			},
			"en_instance_crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of EN instance",
			},
			"en_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint of EN instance",
			},
			"en_source_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "EN source id which appears in instance",
			},
			"created_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of the environment.",
			},
			"updated_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last modified time of the environment data.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Integration URL",
			},
		},
	}
}

func dataSourceIbmAppConfigIntegrationEnRead(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("getAppConfigClient failed %s", err)
	}

	integrationId := d.Get("integration_id").(string)

	options := &appconfigurationv1.GetIntegrationOptions{
		IntegrationID: core.StringPtr(integrationId),
	}

	result, response, error := appconfigClient.GetIntegration(options)

	if error != nil {
		return flex.FmtErrorf("Get Integration failed %s\n%s", error, response)
	}
	d.SetId(fmt.Sprintf("%s/%s", guid, integrationId))

	integrationType := *result.IntegrationType
	if integrationType != "EVENT_NOTIFICATIONS" {
		return flex.FmtErrorf("Integration is not of type Event Notification")
	}

	metadata := result.Metadata.(*appconfigurationv1.IntegrationMetadata)

	error = d.Set("integration_type", "EVENT_NOTIFICATIONS")
	if error != nil {
		return flex.FmtErrorf("Error while setting integration_type %s", error)
	}

	error = d.Set("created_time", result.CreatedTime.String())
	if error != nil {
		return flex.FmtErrorf("Error while setting created_time %s", error)
	}

	error = d.Set("updated_time", result.UpdatedTime.String())
	if error != nil {
		return flex.FmtErrorf("Error while setting updated_time %s", error)
	}

	error = d.Set("href", *result.Href)
	if error != nil {
		return flex.FmtErrorf("Error while setting href %s", error)
	}

	error = d.Set("en_instance_crn", *metadata.EventNotificationsInstanceCrn)
	if error != nil {
		return flex.FmtErrorf("Error while setting en_crn %s", error)
	}

	error = d.Set("en_endpoint", *metadata.EventNotificationsEndpoint)
	if error != nil {
		return flex.FmtErrorf("Error while setting en_endpoint %s", error)
	}

	error = d.Set("en_source_id", *metadata.EventNotificationsSourceID)
	if error != nil {
		return flex.FmtErrorf("Error while setting en_source_name %s", error)
	}

	return nil
}
