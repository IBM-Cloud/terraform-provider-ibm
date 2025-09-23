// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package appconfiguration

import (
	"fmt"

	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func DataSourceIBMAppConfigIntegrationKms() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmAppConfigIntegrationKmsRead,

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
			"kms_instance_crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of KMS instance",
			},
			"kms_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint of KMS instance",
			},
			"root_key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of root key used for encryption",
			},
			"key_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of root key",
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

func dataSourceIbmAppConfigIntegrationKmsRead(d *schema.ResourceData, meta interface{}) error {
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
	if integrationType != "KMS" {
		return flex.FmtErrorf("Integration is not of type KMS")
	}

	metadata := result.Metadata.(*appconfigurationv1.IntegrationMetadata)

	error = d.Set("integration_type", "KMS")
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

	error = d.Set("kms_instance_crn", *metadata.KmsInstanceCrn)
	if error != nil {
		return flex.FmtErrorf("Error while setting kms_crn %s", error)
	}

	error = d.Set("kms_endpoint", *metadata.KmsEndpoint)
	if error != nil {
		return flex.FmtErrorf("Error while setting kms_endpoint %s", error)
	}

	error = d.Set("root_key_id", *metadata.RootKeyID)
	if error != nil {
		return flex.FmtErrorf("Error while setting kms_endpoint %s", error)
	}

	error = d.Set("key_status", *metadata.KeyStatus)
	if error != nil {
		return flex.FmtErrorf("Error while setting kms_status %s", error)
	}

	error = d.Set("kms_schema_type", *metadata.KmsSchemeType)
	if error != nil {
		return flex.FmtErrorf("Error while setting kms_schema_type %s", error)
	}

	return nil
}
