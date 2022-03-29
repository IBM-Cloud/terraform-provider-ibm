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

func ResourceIBMSccAddonActivityInsightsConfiguration() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMSccAddonActivityInsightsConfigurationCreate,
		ReadContext:   ResourceIBMSccAddonActivityInsightsConfigurationRead,
		UpdateContext: ResourceIBMSccAddonActivityInsightsConfigurationUpdate,
		DeleteContext: ResourceIBMSccAddonActivityInsightsConfigurationDelete,
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
				ValidateFunc: validate.InvokeValidator("ibm_scc_addon_activity_insights_configuration", "status"),
				Description:  "Enable or Disable.",
			},
		},
	}
}

func ResourceIBMSccAddonActivityInsightsConfigurationValidator() *validate.ResourceValidator {
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_scc_addon_activity_insights_configuration", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMSccAddonActivityInsightsConfigurationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	updateActivityInsightStatusV2Options := &addonmanagerv1.UpdateActivityInsightStatusV2Options{}

	updateActivityInsightStatusV2Options.SetRegionID(d.Get("region_id").(string))
	updateActivityInsightStatusV2Options.SetStatus(d.Get("status").(string))

	response, err := addonManagerClient.UpdateActivityInsightStatusV2WithContext(context, updateActivityInsightStatusV2Options)
	if err != nil {
		log.Printf("[DEBUG] UpdateActivityInsightStatusV2WithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("UpdateActivityInsightStatusV2WithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *addonManagerClient.AccountID, "activity-insights"))

	return ResourceIBMSccAddonActivityInsightsConfigurationRead(context, d, meta)
}

func ResourceIBMSccAddonActivityInsightsConfigurationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	addonManagerClient, err := meta.(conns.ClientSession).AddonManagerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	addonManagerClient.AccountID = &parts[0]

	getActivityInsightStatusV2Options := &addonmanagerv1.GetActivityInsightStatusV2Options{}

	aiEnableAddOn, response, err := addonManagerClient.GetActivityInsightStatusV2WithContext(context, getActivityInsightStatusV2Options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetActivityInsightStatusV2WithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetActivityInsightStatusV2WithContext failed %s\n%s", err, response))
	}

	if err = d.Set("addon", aiEnableAddOn.Addon); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting region_id: %s", err))
	}
	if err = d.Set("status", aiEnableAddOn.Status); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
	}

	return nil
}

func ResourceIBMSccAddonActivityInsightsConfigurationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	updateActivityInsightStatusV2Options := &addonmanagerv1.UpdateActivityInsightStatusV2Options{}

	hasChange := false

	if d.HasChange("region_id") || d.HasChange("status") {
		updateActivityInsightStatusV2Options.SetRegionID(d.Get("region_id").(string))
		updateActivityInsightStatusV2Options.SetStatus(d.Get("status").(string))
		hasChange = true
	}

	if hasChange {
		response, err := addonManagerClient.UpdateActivityInsightStatusV2WithContext(context, updateActivityInsightStatusV2Options)
		if err != nil {
			log.Printf("[DEBUG] UpdateActivityInsightStatusV2WithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateActivityInsightStatusV2WithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMSccAddonActivityInsightsConfigurationRead(context, d, meta)
}

func ResourceIBMSccAddonActivityInsightsConfigurationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Println(fmt.Sprintf("[DEBUG] configuration deletion is not supported"))
	return nil
}
