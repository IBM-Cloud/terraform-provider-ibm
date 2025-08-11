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

func ResourceIBMAppConfigIntegrationKms() *schema.Resource {
	return &schema.Resource{
		Read:     resourceIntegrationKmsRead,
		Create:   resourceIntegrationKmsCreate,
		Update:   resourceIntegrationKmsUpdate,
		Delete:   resourceIntegrationKmsDelete,
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
				Description: "Integration ID for KMS integration",
			},
			"kms_instance_crn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "CRN of KMS instance",
			},
			"kms_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Endpoint of KMS instance",
			},
			"root_key_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of root key which will be used for encryption",
			},
			"integration_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Integration type [will be KMS always]",
			},
			"kms_schema_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of key protect instance being used",
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

func resourceIntegrationKmsCreate(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)
	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", fmt.Sprintf("%s", err))
	}
	options := &appconfigurationv1.CreateIntegrationOptions{}
	options.SetIntegrationType("KMS")
	options.SetIntegrationID(d.Get("integration_id").(string))
	options.Metadata = &appconfigurationv1.CreateIntegrationMetadataCreateKmsIntegrationMetadata{
		KmsInstanceCrn: core.StringPtr(d.Get("kms_instance_crn").(string)),
		KmsEndpoint:    core.StringPtr(d.Get("kms_endpoint").(string)),
		RootKeyID:      core.StringPtr(d.Get("root_key_id").(string)),
	}

	_, response, err := appconfigClient.CreateIntegration(options)

	if err != nil {
		return flex.FmtErrorf("[ERROR] Create KMS integration failed %s\n%s", err, response)
	}
	d.SetId(fmt.Sprintf("%s/%s", guid, *options.IntegrationID))

	return resourceIntegrationKmsRead(d, meta)
}

func resourceIntegrationKmsUpdate(d *schema.ResourceData, meta interface{}) error {
	return flex.FmtErrorf("[ERROR] Update KMS is not yet implemented")
}

func resourceIntegrationKmsRead(d *schema.ResourceData, meta interface{}) error {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return nil
	}
	appconfigClient, err := getAppConfigClient(meta, parts[0])
	if err != nil {
		return flex.FmtErrorf("%s", fmt.Sprintf("%s", err))
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
	metadata := result.Metadata.(*appconfigurationv1.IntegrationMetadata)
	if metadata.KmsSchemeType != nil {
		if err = d.Set("kms_schema_type", *metadata.KmsSchemeType); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting kms schema type: %s", err)
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

func resourceIntegrationKmsDelete(d *schema.ResourceData, meta interface{}) error {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return nil
	}

	appconfigClient, err := getAppConfigClient(meta, parts[0])
	if err != nil {
		return flex.FmtErrorf("%s", fmt.Sprintf("%s", err))
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
