// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/scc-go-sdk/v4/addonmanagerv1"
)

func ResourceIBMSccAddonNetworkInsightsConfiguration() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMSccAddonNetworkInsightsConfigurationCreate,
		ReadContext:   ResourceIBMSccAddonNetworkInsightsConfigurationRead,
		UpdateContext: ResourceIBMSccAddonNetworkInsightsConfigurationUpdate,
		DeleteContext: ResourceIBMSccAddonNetworkInsightsConfigurationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"region_id": &schema.Schema{
				DiffSuppressFunc: flex.ApplyOnce,
				ForceNew:         true,
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Region id for example - us.",
			},
			"addon": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Region id for example - us.",
			},
			"status": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_scc_addon_network_insights_configuration", "status"),
				Description:  "Enable or Disable.",
			},
		},
	}
}

func ResourceIBMSccAddonNetworkInsightsConfigurationValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 1)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "status",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "disable, enable",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_scc_addon_network_insights_configuration", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMSccAddonNetworkInsightsConfigurationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	addonManagerClient, err := meta.(conns.ClientSession).AddonManagerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	accountID := d.Get("account_id").(string)
	log.Println(fmt.Sprintf("[DEBUG] using specified AccountID %s", accountID))
	if accountID == "" {
		accountID = userDetails.UserAccount
		log.Println(fmt.Sprintf("[DEBUG] AccountID not spedified, using %s", accountID))
	}
	addonManagerClient.AccountID = &accountID

	updateNetworkInsightStatusV2Options := &addonmanagerv1.UpdateNetworkInsightStatusV2Options{}

	updateNetworkInsightStatusV2Options.SetRegionID(d.Get("region_id").(string))
	updateNetworkInsightStatusV2Options.SetStatus(d.Get("status").(string))

	response, err := addonManagerClient.UpdateNetworkInsightStatusV2WithContext(context, updateNetworkInsightStatusV2Options)
	if err != nil {
		log.Printf("[DEBUG] UpdateNetworkInsightStatusV2WithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("UpdateNetworkInsightStatusV2WithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *addonManagerClient.AccountID, "network-insights"))

	return ResourceIBMSccAddonNetworkInsightsConfigurationRead(context, d, meta)
}

func ResourceIBMSccAddonNetworkInsightsConfigurationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	addonManagerClient, err := meta.(conns.ClientSession).AddonManagerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	addonManagerClient.AccountID = &parts[0]

	getNetworkInsightStatusV2Options := &addonmanagerv1.GetNetworkInsightStatusV2Options{}

	niEnableAddOn, response, err := addonManagerClient.GetNetworkInsightStatusV2WithContext(context, getNetworkInsightStatusV2Options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetNetworkInsightStatusV2WithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetNetworkInsightStatusV2WithContext failed %s\n%s", err, response))
	}

	if err = d.Set("addon", niEnableAddOn.Addon); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting region_id: %s", err))
	}
	if err = d.Set("status", niEnableAddOn.Status); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
	}

	return nil
}

func ResourceIBMSccAddonNetworkInsightsConfigurationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	addonManagerClient, err := meta.(conns.ClientSession).AddonManagerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	accountID := d.Get("account_id").(string)
	log.Println(fmt.Sprintf("[DEBUG] using specified AccountID %s", accountID))
	if accountID == "" {
		accountID = userDetails.UserAccount
		log.Println(fmt.Sprintf("[DEBUG] AccountID not spedified, using %s", accountID))
	}
	addonManagerClient.AccountID = &accountID

	updateNetworkInsightStatusV2Options := &addonmanagerv1.UpdateNetworkInsightStatusV2Options{}

	hasChange := false

	if d.HasChange("region_id") || d.HasChange("status") {
		updateNetworkInsightStatusV2Options.SetRegionID(d.Get("region_id").(string))
		updateNetworkInsightStatusV2Options.SetStatus(d.Get("status").(string))
		hasChange = true
	}

	if hasChange {
		response, err := addonManagerClient.UpdateNetworkInsightStatusV2WithContext(context, updateNetworkInsightStatusV2Options)
		if err != nil {
			log.Printf("[DEBUG] UpdateNetworkInsightStatusV2WithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateNetworkInsightStatusV2WithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMSccAddonNetworkInsightsConfigurationRead(context, d, meta)
}

func ResourceIBMSccAddonNetworkInsightsConfigurationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Println(fmt.Sprintf("[DEBUG] configuration deletion is not supported"))
	return nil
}
