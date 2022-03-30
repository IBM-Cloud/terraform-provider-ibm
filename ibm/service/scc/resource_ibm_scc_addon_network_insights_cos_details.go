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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v4/addonmanagerv1"
)

func ResourceIBMSccAddonNetworkInsightsCosDetails() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMSccAddonNetworkInsightsCosDetailsCreate,
		ReadContext:   ResourceIBMSccAddonNetworkInsightsCosDetailsRead,
		UpdateContext: ResourceIBMSccAddonNetworkInsightsCosDetailsUpdate,
		DeleteContext: ResourceIBMSccAddonNetworkInsightsCosDetailsDelete,
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
				Description:      "Region for example - us-south, eu-gb.",
			},
			"cos_details": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cos_instance": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"bucket_name": &schema.Schema{
							ForceNew: true,
							Type:     schema.TypeString,
							Required: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Insights type.",
						},
						"cos_bucket_url": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "cos bucket url.",
						},
					},
				},
			},
		},
	}
}

func ResourceIBMSccAddonNetworkInsightsCosDetailsCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	addNetworkInsightsCosDetailsV2Options := &addonmanagerv1.AddNetworkInsightsCosDetailsV2Options{}

	addNetworkInsightsCosDetailsV2Options.SetRegionID(d.Get("region_id").(string))
	var cosDetails []addonmanagerv1.CosDetails
	for _, e := range d.Get("cos_details").([]interface{}) {
		value := e.(map[string]interface{})
		cosDetailsItem, err := ResourceIBMSccAddonNetworkInsightsCosDetailsMapToNiCosDetailsV2InputCosDetailsItem(value)
		if err != nil {
			return diag.FromErr(err)
		}
		cosDetails = append(cosDetails, *cosDetailsItem)
	}
	addNetworkInsightsCosDetailsV2Options.SetCosDetails(cosDetails)

	niCosDetailsV2Output, response, err := addonManagerClient.AddNetworkInsightsCosDetailsV2WithContext(context, addNetworkInsightsCosDetailsV2Options)
	if err != nil {
		log.Printf("[DEBUG] AddNetworkInsightsCosDetailsV2WithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("AddNetworkInsightsCosDetailsV2WithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *addonManagerClient.AccountID, *niCosDetailsV2Output.CosDetails[0].ID))

	return ResourceIBMSccAddonNetworkInsightsCosDetailsRead(context, d, meta)
}

func ResourceIBMSccAddonNetworkInsightsCosDetailsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	addonManagerClient, err := meta.(conns.ClientSession).AddonManagerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	addonManagerClient.AccountID = &parts[0]

	getNetworkInsightsCosDetailsV2Options := &addonmanagerv1.GetNetworkInsightsCosDetailsV2Options{}

	niCosDetailsV2Output, response, err := addonManagerClient.GetNetworkInsightsCosDetailsV2WithContext(context, getNetworkInsightsCosDetailsV2Options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetNetworkInsightsCosDetailsV2WithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetNetworkInsightsCosDetailsV2WithContext failed %s\n%s", err, response))
	}

	for _, cosDetail := range niCosDetailsV2Output.CosDetails {
		if *cosDetail.ID == parts[1] {
			if err = d.Set("region_id", d.Get("region_id")); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting region_id: %s", err))
			}

			cosDetailsItemMap, err := ResourceIBMSccAddonNetworkInsightsCosDetailsNiCosDetailsV2OutputCosDetailsItemToMap(&cosDetail)
			if err != nil {
				return diag.FromErr(err)
			}
			cosDetails := []map[string]interface{}{
				cosDetailsItemMap,
			}

			if err = d.Set("cos_details", cosDetails); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting cos_details: %s", err))
			}

			return nil
		}
	}

	d.SetId("")

	return nil
}

func ResourceIBMSccAddonNetworkInsightsCosDetailsUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	addNetworkInsightsCosDetailsV2Options := &addonmanagerv1.AddNetworkInsightsCosDetailsV2Options{}

	hasChange := false

	if d.HasChange("region_id") || d.HasChange("cos_details") {
		addNetworkInsightsCosDetailsV2Options.SetRegionID(d.Get("region_id").(string))
		var cosDetails []addonmanagerv1.CosDetails
		for _, e := range d.Get("cos_details").([]interface{}) {
			value := e.(map[string]interface{})
			cosDetailsItem, err := ResourceIBMSccAddonNetworkInsightsCosDetailsMapToNiCosDetailsV2InputCosDetailsItem(value)
			if err != nil {
				return diag.FromErr(err)
			}
			cosDetails = append(cosDetails, *cosDetailsItem)
		}
		addNetworkInsightsCosDetailsV2Options.SetCosDetails(cosDetails)
		hasChange = true
	}

	if hasChange {
		_, response, err := addonManagerClient.AddNetworkInsightsCosDetailsV2WithContext(context, addNetworkInsightsCosDetailsV2Options)
		if err != nil {
			log.Printf("[DEBUG] AddNetworkInsightsCosDetailsV2WithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("AddNetworkInsightsCosDetailsV2WithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMSccAddonNetworkInsightsCosDetailsRead(context, d, meta)
}

func ResourceIBMSccAddonNetworkInsightsCosDetailsDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	addonManagerClient, err := meta.(conns.ClientSession).AddonManagerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	addonManagerClient.AccountID = &parts[0]

	deleteNetworkInsightsCosDetailsV2Options := &addonmanagerv1.DeleteNetworkInsightsCosDetailsV2Options{
		Ids: []string{
			parts[1],
		},
	}

	response, err := addonManagerClient.DeleteNetworkInsightsCosDetailsV2WithContext(context, deleteNetworkInsightsCosDetailsV2Options)
	if err != nil {
		log.Printf("[DEBUG] DeleteNetworkInsightsCosDetailsV2WithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteNetworkInsightsCosDetailsV2WithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIBMSccAddonNetworkInsightsCosDetailsMapToNiCosDetailsV2InputCosDetailsItem(modelMap map[string]interface{}) (*addonmanagerv1.CosDetails, error) {
	model := &addonmanagerv1.CosDetails{}
	model.CosInstance = core.StringPtr(modelMap["cos_instance"].(string))
	model.BucketName = core.StringPtr(modelMap["bucket_name"].(string))
	model.Description = core.StringPtr(modelMap["description"].(string))
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.CosBucketURL = core.StringPtr(modelMap["cos_bucket_url"].(string))
	return model, nil
}

func ResourceIBMSccAddonNetworkInsightsCosDetailsNiCosDetailsV2OutputCosDetailsItemToMap(model *addonmanagerv1.CosDetailsWithID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID
	modelMap["cos_instance"] = model.CosInstance
	modelMap["bucket_name"] = model.BucketName
	modelMap["description"] = model.Description
	modelMap["type"] = model.Type
	modelMap["cos_bucket_url"] = model.CosBucketURL
	return modelMap, nil
}
