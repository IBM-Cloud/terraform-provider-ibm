// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMAppConfigIntegrationEn() *schema.Resource {
	return &schema.Resource{
		Read:     resourceIntegrationEnRead,
		Create:   resourceIntegrationEnCreate,
		Update:   resourceIntegrationEnUpdate,
		Delete:   resourceIntegrationEnDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"integration_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Integration ID for EN integration",
			},
			"en_instance_crn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "CRN of EN instance",
			},
			"en_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Endpoint of EN instance",
			},
			"en_source_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "EN source name which appears in instance",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "EN integration description",
			},
			"integration_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Integration type [will be EVENT_NOTIFICATIONS always]",
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

func resourceIntegrationEnCreate(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)
	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf(fmt.Sprintf("%s", err))
	}
	options := &appconfigurationv1.CreateIntegrationOptions{}
	options.SetIntegrationType("EVENT_NOTIFICATIONS")
	options.SetIntegrationID(d.Get("integration_id").(string))

	if _, ok := GetFieldExists(d, "description"); ok {
		options.Metadata = &appconfigurationv1.CreateIntegrationMetadataCreateEnIntegrationMetadata{
			EventNotificationsInstanceCrn:       core.StringPtr(d.Get("en_instance_crn").(string)),
			EventNotificationsEndpoint:          core.StringPtr(d.Get("en_endpoint").(string)),
			EventNotificationsSourceName:        core.StringPtr(d.Get("en_source_name").(string)),
			EventNotificationsSourceDescription: core.StringPtr(d.Get("description").(string)),
		}
	} else {
		options.Metadata = &appconfigurationv1.CreateIntegrationMetadataCreateEnIntegrationMetadata{
			EventNotificationsInstanceCrn: core.StringPtr(d.Get("en_instance_crn").(string)),
			EventNotificationsEndpoint:    core.StringPtr(d.Get("en_endpoint").(string)),
			EventNotificationsSourceName:  core.StringPtr(d.Get("en_source_name").(string)),
		}
	}

	_, response, err := appconfigClient.CreateIntegration(options)

	if err != nil {
		return flex.FmtErrorf("[ERROR] Create EN integration failed %s\n%s", err, response)
	}
	d.SetId(fmt.Sprintf("%s/%s", guid, *options.IntegrationID))

	return resourceIntegrationEnRead(d, meta)
}

func resourceIntegrationEnUpdate(d *schema.ResourceData, meta interface{}) error {
	return flex.FmtErrorf("[ERROR] Update EN Integration is not yet implemented")
}

func resourceIntegrationEnRead(d *schema.ResourceData, meta interface{}) error {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return nil
	}
	appconfigClient, err := getAppConfigClient(meta, parts[0])
	if err != nil {
		return flex.FmtErrorf(fmt.Sprintf("%s", err))
	}

	options := &appconfigurationv1.GetIntegrationOptions{}
	options.SetIntegrationID(parts[1])

	result, response, err := appconfigClient.GetIntegration(options)

	if err != nil {
		return flex.FmtErrorf("[ERROR] GetIntegration failed %s\n%s", err, response)
	}

	d.Set("guid", parts[0])
	if result.IntegrationType != nil {
		if err = d.Set("integration_type", *result.IntegrationType); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting integration type: %s", err)
		}
	}
	if result.CreatedTime != nil {
		if err = d.Set("created_time", result.CreatedTime.String()); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting created_time: %s", err)
		}
	}
	if result.UpdatedTime != nil {
		if err = d.Set("updated_time", result.UpdatedTime.String()); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting updated_time: %s", err)
		}
	}
	if result.Href != nil {
		if err = d.Set("href", *result.Href); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting href: %s", err)
		}
	}
	return nil
}

func resourceIntegrationEnDelete(d *schema.ResourceData, meta interface{}) error {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return nil
	}

	appconfigClient, err := getAppConfigClient(meta, parts[0])
	if err != nil {
		return flex.FmtErrorf(fmt.Sprintf("%s", err))
	}

	options := &appconfigurationv1.DeleteIntegrationOptions{}
	options.SetIntegrationID(parts[1])

	response, err := appconfigClient.DeleteIntegration(options)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return flex.FmtErrorf("[ERROR] Delete Integration failed %s\n%s", err, response)
	}
	d.SetId("")
	return nil
}
