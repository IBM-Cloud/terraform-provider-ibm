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

const (
	ibmCISAiSecuritySettings = "ibm_cis_ai_security_settings"
	cisAiSecurityEnabled     = "enabled"
)

func ResourceIBMCISAiSecuritySettings() *schema.Resource {
	return &schema.Resource{
		Create:   ResourceIBMCISAiSecuritySettingsUpdate,
		Read:     ResourceIBMCISAiSecuritySettingsRead,
		Update:   ResourceIBMCISAiSecuritySettingsUpdate,
		Delete:   ResourceIBMCISAiSecuritySettingsDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CIS Instance CRN",
				ValidateFunc: validate.InvokeValidator(ibmCISAiSecuritySettings,
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Associated CIS domain ID",
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisAiSecurityEnabled: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether AI Security for Apps is enabled on the zone",
			},
		},
	}
}

func ResourceIBMCISAiSecuritySettingsValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})
	ibmCISAiSecuritySettingsValidator := validate.ResourceValidator{ResourceName: ibmCISAiSecuritySettings, Schema: validateSchema}
	return &ibmCISAiSecuritySettingsValidator
}

func ResourceIBMCISAiSecuritySettingsUpdate(d *schema.ResourceData, meta interface{}) error {
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

	if d.HasChange(cisAiSecurityEnabled) {
		enabled := d.Get(cisAiSecurityEnabled).(bool)
		opt := cisClient.NewReplaceZoneAiSecuritySettingsOptions()
		opt.SetEnabled(enabled)
		_, response, err := cisClient.ReplaceZoneAiSecuritySettingsWithContext(context.Background(), opt)
		if err != nil {
			log.Printf("[ERROR] Update AI Security Settings failed: %v", response)
			return err
		}
	}

	d.SetId(flex.ConvertCisToTfTwoVar(zoneID, crn))
	return ResourceIBMCISAiSecuritySettingsRead(d, meta)
}

func ResourceIBMCISAiSecuritySettingsRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisAiSecurityForAppsSession()
	if err != nil {
		return err
	}
	zoneID, crn, err := flex.ConvertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	opt := cisClient.NewGetAiSecuritySettingsOptions()
	result, response, err := cisClient.GetAiSecuritySettingsWithContext(context.Background(), opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			log.Printf("[WARN] AI Security Settings for zone %s not found, removing from state", zoneID)
			d.SetId("")
			return nil
		}
		log.Printf("[ERROR] Get AI Security Settings failed: %v", response)
		return err
	}

	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	if result != nil && result.Result != nil && result.Result.Enabled != nil {
		d.Set(cisAiSecurityEnabled, *result.Result.Enabled)
	}
	return nil
}

func ResourceIBMCISAiSecuritySettingsDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisAiSecurityForAppsSession()
	if err != nil {
		return err
	}
	zoneID, crn, err := flex.ConvertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	opt := cisClient.NewReplaceZoneAiSecuritySettingsOptions()
	opt.SetEnabled(false)
	_, response, err := cisClient.ReplaceZoneAiSecuritySettingsWithContext(context.Background(), opt)
	if err != nil {
		log.Printf("[ERROR] Disabling AI Security Settings on delete failed: %v", response)
		return err
	}

	d.SetId("")
	return nil
}
