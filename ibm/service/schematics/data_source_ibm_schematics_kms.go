// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func DataSourceIBMSchematicsKMS() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMSchematicsKMSRead,

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The location to integrate kms instance. For example, location can be `US` and `EU`.",
			},
			"encryption_scheme": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The encryption scheme values. Allowable values: `byok`, `kyok`.",
			},
			"resource_group": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The kms instance resource group to integrate.",
			},
			"primary_crk": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The primary kms instance details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The primary kms instance name.",
						},
						"kms_private_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The primary kms instance private endpoint.",
						},
						"key_crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of the primary root key.",
						},
					},
				},
			},
			"secondary_crk": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The secondary kms instance details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The secondary kms instance name.",
						},
						"kms_private_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The secondary kms instance private endpoint.",
						},
						"key_crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of the secondary key.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMSchematicsKMSRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIBMSchematicsKMSRead schematicsClient initialization failed: %s", err.Error()), "ibm_schematics_kms", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Get location from datasource or provider region
	var region string
	if r, ok := d.GetOk("location"); ok {
		region = r.(string)
	} else {
		// Get region from provider
		sess, err := meta.(conns.ClientSession).BluemixSession()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIBMSchematicsKMSRead failed to get session: %s", err.Error()), "ibm_schematics_kms", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		region = sess.Config.Region
	}

	// Update the service URL based on region if needed
	if region != "" {
		schematicsURL, updatedURL, _ := SchematicsEndpointURL(region, meta)
		if updatedURL {
			schematicsClient.Service.Options.URL = schematicsURL
		}
	}

	// Map region to KMS location (US or EU)
	kmsLocation := mapRegionToKMSLocation(region)

	getKmsSettingsOptions := &schematicsv1.GetKmsSettingsOptions{}
	getKmsSettingsOptions.SetLocation(kmsLocation)

	kmsSettings, response, err := schematicsClient.GetKmsSettingsWithContext(context, getKmsSettingsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIBMSchematicsKMSRead GetKmsSettingsWithContext failed with error: %s and response:\n%s", err, response), "ibm_schematics_kms", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s", kmsLocation))

	if err = d.Set("location", kmsSettings.Location); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIBMSchematicsKMSRead failed with error: %s", err), "ibm_schematics_kms", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("encryption_scheme", kmsSettings.EncryptionScheme); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIBMSchematicsKMSRead failed with error: %s", err), "ibm_schematics_kms", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("resource_group", kmsSettings.ResourceGroup); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIBMSchematicsKMSRead failed with error: %s", err), "ibm_schematics_kms", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	primaryCrk := []map[string]interface{}{}
	if kmsSettings.PrimaryCrk != nil {
		modelMap, err := dataSourceIBMSchematicsKMSPrimaryCrkToMap(kmsSettings.PrimaryCrk)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIBMSchematicsKMSRead failed: %s", err.Error()), "ibm_schematics_kms", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		primaryCrk = append(primaryCrk, modelMap)
	}
	if err = d.Set("primary_crk", primaryCrk); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIBMSchematicsKMSRead failed with error: %s", err), "ibm_schematics_kms", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	secondaryCrk := []map[string]interface{}{}
	if kmsSettings.SecondaryCrk != nil {
		modelMap, err := dataSourceIBMSchematicsKMSSecondaryCrkToMap(kmsSettings.SecondaryCrk)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIBMSchematicsKMSRead failed: %s", err.Error()), "ibm_schematics_kms", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		secondaryCrk = append(secondaryCrk, modelMap)
	}
	if err = d.Set("secondary_crk", secondaryCrk); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("dataSourceIBMSchematicsKMSRead failed with error: %s", err), "ibm_schematics_kms", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return nil
}

func dataSourceIBMSchematicsKMSPrimaryCrkToMap(model *schematicsv1.KMSSettingsPrimaryCrk) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.KmsName != nil {
		modelMap["kms_name"] = *model.KmsName
	}
	if model.KmsPrivateEndpoint != nil {
		modelMap["kms_private_endpoint"] = *model.KmsPrivateEndpoint
	}
	if model.KeyCrn != nil {
		modelMap["key_crn"] = *model.KeyCrn
	}
	return modelMap, nil
}

func dataSourceIBMSchematicsKMSSecondaryCrkToMap(model *schematicsv1.KMSSettingsSecondaryCrk) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.KmsName != nil {
		modelMap["kms_name"] = *model.KmsName
	}
	if model.KmsPrivateEndpoint != nil {
		modelMap["kms_private_endpoint"] = *model.KmsPrivateEndpoint
	}
	if model.KeyCrn != nil {
		modelMap["key_crn"] = *model.KeyCrn
	}
	return modelMap, nil
}

// mapRegionToKMSLocation maps IBM Cloud regions to KMS location (US or EU)
func mapRegionToKMSLocation(region string) string {
	switch region {
	case "us-south", "us-east":
		return "US"
	case "eu-de", "eu-gb":
		return "EU"
	default:
		// For other regions or if already US/EU, return as-is
		return region
	}
}
