// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"context"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMCISAiSecuritySettings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMCISAiSecuritySettingsRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance CRN",
				Required:    true,
				ValidateFunc: validate.InvokeDataSourceValidator(
					ibmCISAiSecuritySettings,
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain ID",
				Required:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisAiSecurityEnabled: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether AI Security for Apps is enabled on the zone",
			},
		},
	}
}

func DataSourceIBMCISAiSecuritySettingsValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})
	iBMCISAiSecuritySettingsValidator := validate.ResourceValidator{
		ResourceName: ibmCISAiSecuritySettings,
		Schema:       validateSchema}
	return &iBMCISAiSecuritySettingsValidator
}

func dataSourceIBMCISAiSecuritySettingsRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisAiSecurityForAppsSession()
	if err != nil {
		return err
	}

	crn := d.Get(cisID).(string)
	domainID := d.Get(cisDomainID).(string)
	zoneID, _, _ := flex.ConvertTftoCisTwoVar(domainID)
	if zoneID == "" {
		zoneID = domainID
	}
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	opt := cisClient.NewGetAiSecuritySettingsOptions()
	result, response, err := cisClient.GetAiSecuritySettingsWithContext(context.Background(), opt)
	if err != nil {
		log.Printf("[ERROR] Get AI Security Settings failed: %v", response)
		return err
	}

	d.SetId(flex.ConvertCisToTfTwoVar(zoneID, crn))
	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	if result != nil && result.Result != nil && result.Result.Enabled != nil {
		d.Set(cisAiSecurityEnabled, *result.Result.Enabled)
	}
	return nil
}
