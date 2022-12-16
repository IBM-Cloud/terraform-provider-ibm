// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metricsrouter

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
)

func ResourceIBMMetricsRouterTarget() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMMetricsRouterTargetCreate,
		ReadContext:   resourceIBMMetricsRouterTargetRead,
		UpdateContext: resourceIBMMetricsRouterTargetUpdate,
		DeleteContext: resourceIBMMetricsRouterTargetDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_metrics_router_target", "name"),
				Description:  "The name of the target. The name must be 1000 characters or less, and cannot include any special characters other than `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names.",
			},
			"destination_crn": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_metrics_router_target", "destination_crn"),
				Description:  "The CRN of a destination service instance or resource.",
			},
			"region": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_metrics_router_target", "region"),
				Description:  "Include this optional field if you want to create a target in a different region other than the one you are connected.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the target resource.",
			},
			"target_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the target.",
			},
			"write_status": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The status of the write attempt to the target with the provided destination info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The status such as failed or success.",
						},
						"last_failure": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The timestamp of the failure.",
						},
						"reason_for_last_failure": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Detailed description of the cause of the failure.",
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp of the target creation time.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp of the target last updated time.",
			},
			"api_version": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The API version of the target.",
			},
		},
	}
}

func ResourceIBMMetricsRouterTargetValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9 -._:]+$`,
			MinValueLength:             1,
			MaxValueLength:             1000,
		},
		validate.ValidateSchema{
			Identifier:                 "destination_crn",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9 -._:\/]+$`,
			MinValueLength:             3,
			MaxValueLength:             1000,
		},
		validate.ValidateSchema{
			Identifier:                 "region",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9 -._:]+$`,
			MinValueLength:             3,
			MaxValueLength:             1000,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_metrics_router_target", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMMetricsRouterTargetCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	createTargetOptions := &metricsrouterv3.CreateTargetOptions{}

	createTargetOptions.SetName(d.Get("name").(string))
	createTargetOptions.SetDestinationCRN(d.Get("destination_crn").(string))
	if _, ok := d.GetOk("region"); ok {
		createTargetOptions.SetRegion(d.Get("region").(string))
	}

	target, response, err := metricsRouterClient.CreateTargetWithContext(context, createTargetOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateTargetWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateTargetWithContext failed %s\n%s", err, response))
	}

	d.SetId(*target.ID)

	return resourceIBMMetricsRouterTargetRead(context, d, meta)
}

func resourceIBMMetricsRouterTargetRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	getTargetOptions := &metricsrouterv3.GetTargetOptions{}

	getTargetOptions.SetID(d.Id())

	target, response, err := metricsRouterClient.GetTargetWithContext(context, getTargetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetTargetWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetTargetWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("name", target.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("destination_crn", target.DestinationCRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting destination_crn: %s", err))
	}
	if _, exists := d.GetOk("region"); exists {
		if target.Region != nil && len(*target.Region) > 0 {
			d.Set("region", *target.Region)
			if err = d.Set("region", *target.Region); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting region: %s", err))
			}
		}
	}
	if err = d.Set("crn", target.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("target_type", target.TargetType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting target_type: %s", err))
	}
	writeStatusMap, err := resourceIBMMetricsRouterTargetWriteStatusToMap(target.WriteStatus)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("write_status", []map[string]interface{}{writeStatusMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting write_status: %s", err))
	}
	if err = d.Set("created_at", flex.DateTimeToString(target.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("updated_at", flex.DateTimeToString(target.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("api_version", flex.IntValue(target.APIVersion)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting api_version: %s", err))
	}

	return nil
}

func resourceIBMMetricsRouterTargetUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceTargetOptions := &metricsrouterv3.ReplaceTargetOptions{}

	replaceTargetOptions.SetID(d.Id())

	hasChange := false

	if d.HasChange("name") || d.HasChange("destination_crn") || d.HasChange("region") {
		replaceTargetOptions.SetName(d.Get("name").(string))
		replaceTargetOptions.SetDestinationCRN(d.Get("destination_crn").(string))
		hasChange = true
	}

	if hasChange {
		_, response, err := metricsRouterClient.ReplaceTargetWithContext(context, replaceTargetOptions)
		if err != nil {
			log.Printf("[DEBUG] ReplaceTargetWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("ReplaceTargetWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIBMMetricsRouterTargetRead(context, d, meta)
}

func resourceIBMMetricsRouterTargetDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTargetOptions := &metricsrouterv3.DeleteTargetOptions{}

	deleteTargetOptions.SetID(d.Id())

	_, response, err := metricsRouterClient.DeleteTargetWithContext(context, deleteTargetOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteTargetWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteTargetWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIBMMetricsRouterTargetWriteStatusToMap(model *metricsrouterv3.WriteStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["status"] = model.Status
	if model.LastFailure != nil {
		modelMap["last_failure"] = model.LastFailure.String()
	}
	if model.ReasonForLastFailure != nil {
		modelMap["reason_for_last_failure"] = model.ReasonForLastFailure
	}
	return modelMap, nil
}
