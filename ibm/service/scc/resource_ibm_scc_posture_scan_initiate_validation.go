// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/scc-go-sdk/v4/posturemanagementv2"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMSccPostureScanInitiateValidation() *schema.Resource {
	return &schema.Resource{
		CreateContext:      resourceIBMSccPostureScanInitiateValidation,
		ReadContext:        resourceIBMSccPostureScanInitiateRead,
		DeleteContext:      resourceIBMSccPostureScanInitiateDelete,
		Importer:           &schema.ResourceImporter{},
		DeprecationMessage: "**Removal Notification** Resource Removal: Resource ibm_scc_posture_scan_initiate_validation is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).",
		Schema: map[string]*schema.Schema{
			"scope_id": {
				Type:         schema.TypeString,
				Description:  "The unique ID of the scope.",
				ForceNew:     true,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_scc_posture_scan_initiate_validation", "scope_id"),
			},
			"profile_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "The unique ID of the profile.",
				ValidateFunc: validate.InvokeValidator("ibm_scc_posture_scan_initiate_validation", "profile_id"),
			},
			"group_profile_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  "The ID of the profile group.",
				ValidateFunc: validate.InvokeValidator("ibm_scc_posture_scan_initiate_validation", "group_profile_id"),
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  "The name of a scheduled scan.",
				ValidateFunc: validate.InvokeValidator("ibm_scc_posture_scan_initiate_validation", "name"),
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  "The description of a scheduled scan.",
				ValidateFunc: validate.InvokeValidator("ibm_scc_posture_scan_initiate_validation", "description"),
			},
			"frequency": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The frequency at which a scan is run specified in milliseconds.",
			},
			"no_of_occurrences": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The number of times that a scan should be run.",
			},
			"end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The date on which a scan should stop running specified in UTC.",
			},
			"result": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the profile group.",
			},
		},
	}
}

func ResourceIBMSccPostureScanInitiateValidationValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "scope_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9]*$`,
			MinValueLength:             1,
			MaxValueLength:             20,
		},
		validate.ValidateSchema{
			Identifier:                 "profile_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9]*$`,
			MinValueLength:             1,
			MaxValueLength:             20,
		},
		validate.ValidateSchema{
			Identifier:                 "group_profile_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[0-9]*$`,
			MinValueLength:             1,
			MaxValueLength:             20,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9-\.,_\s]*$`,
			MinValueLength:             1,
			MaxValueLength:             32,
		},
		validate.ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9-._,\s]*$`,
			MinValueLength:             1,
			MaxValueLength:             255,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_scc_posture_scan_initiate_validation", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMSccPostureScanInitiateValidation(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	postureManagementClient, err := meta.(conns.ClientSession).PostureManagementV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createValidationOptions := &posturemanagementv2.CreateValidationOptions{}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting userDetails %s", err))
	}
	createValidationOptions.SetAccountID(userDetails.UserAccount)

	createValidationOptions.SetScopeID(d.Get("scope_id").(string))
	createValidationOptions.SetProfileID(d.Get("profile_id").(string))

	if _, ok := d.GetOk("group_profile_id"); ok {
		createValidationOptions.SetGroupProfileID(d.Get("group_profile_id").(string))
	}

	if _, ok := d.GetOk("name"); ok {
		createValidationOptions.SetName(d.Get("name").(string))
	}

	if _, ok := d.GetOk("description"); ok {
		createValidationOptions.SetDescription(d.Get("description").(string))
	}

	if frequency, ok := d.GetOk("frequency"); ok {
		createValidationOptions.SetFrequency(int64(frequency.(int)))
	}

	if no_of_occurrences, ok := d.GetOk("no_of_occurrences"); ok {
		createValidationOptions.SetNoOfOccurrences(int64(no_of_occurrences.(int)))
	}

	if end_time, ok := d.GetOk("end_time"); ok {
		createValidationOptions.SetEndTime(end_time.(*strfmt.DateTime))
	}

	result, response, err := postureManagementClient.CreateValidationWithContext(context, createValidationOptions)
	if result == nil || err != nil {
		log.Printf("[DEBUG] CreateValidationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateValidationWithContext failed %s\n%s", err, response))
	}

	if *result.Result {
		correlationId := strings.Split(*result.Message, "= ")[1]
		d.SetId(correlationId)
		d.Set("result", fmt.Sprintf("%v", *result.Result))
		return nil
	}

	return diag.FromErr(fmt.Errorf("CreateValidationWithContext failed %s\n%s", err, *result.Message))
}

func resourceIBMSccPostureScanInitiateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceIBMSccPostureScanInitiateDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
