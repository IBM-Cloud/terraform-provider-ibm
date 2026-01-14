// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package logsrouter

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
)

func ResourceIBMLogsRouterTarget() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMLogsRouterTargetCreate,
		ReadContext:   resourceIBMLogsRouterTargetRead,
		UpdateContext: resourceIBMLogsRouterTargetUpdate,
		DeleteContext: resourceIBMLogsRouterTargetDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_router_target", "name"),
				Description:  "The name of the target resource.",
			},
			"destination_crn": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_router_target", "destination_crn"),
				Description:  "Cloud Resource Name (CRN) of the destination resource. Ensure you have a service authorization between IBM Cloud Logs Routing and your Cloud resource. See [service-to-service authorization](https://cloud.ibm.com/docs/logs-router?topic=logs-router-target-monitoring&interface=ui#target-monitoring-ui) for details.",
			},
			"region": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_router_target", "region"),
				Description:  "Include this optional field if you used it to create a target in a different region other than the one you are connected.",
			},
			"managed_by": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_router_target", "managed_by"),
				Description:  "Present when the target is enterprise-managed (`managed_by: enterprise`). For account-managed targets this field is omitted.",
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
				Description: "The status of the write attempt to the target with the provided endpoint parameters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status such as failed or success.",
						},
						"last_failure": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The timestamp of the failure.",
						},
						"reason_for_last_failure": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
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
		},
	}
}

func ResourceIBMLogsRouterTargetValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9 \-._:]+$`,
			MinValueLength:             1,
			MaxValueLength:             1000,
		},
		validate.ValidateSchema{
			Identifier:                 "destination_crn",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9 \-._:\/]+$`,
			MinValueLength:             3,
			MaxValueLength:             1000,
		},
		validate.ValidateSchema{
			Identifier:                 "region",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9 \-._:]+$`,
			MinValueLength:             3,
			MaxValueLength:             1000,
		},
		validate.ValidateSchema{
			Identifier:                 "managed_by",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "account, enterprise",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_logs_router_target", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMLogsRouterTargetCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createTargetOptions := &logsrouterv3.CreateTargetOptions{}

	createTargetOptions.SetName(d.Get("name").(string))
	createTargetOptions.SetDestinationCRN(d.Get("destination_crn").(string))
	if _, ok := d.GetOk("region"); ok {
		createTargetOptions.SetRegion(d.Get("region").(string))
	}
	if _, ok := d.GetOk("managed_by"); ok {
		createTargetOptions.SetManagedBy(d.Get("managed_by").(string))
	}

	target, _, err := logsRouterClient.CreateTargetWithContext(context, createTargetOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateTargetWithContext failed: %s", err.Error()), "ibm_logs_router_target", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*target.ID)

	return resourceIBMLogsRouterTargetRead(context, d, meta)
}

func resourceIBMLogsRouterTargetRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getTargetOptions := &logsrouterv3.GetTargetOptions{}

	getTargetOptions.SetID(d.Id())

	target, response, err := logsRouterClient.GetTargetWithContext(context, getTargetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetTargetWithContext failed: %s", err.Error()), "ibm_logs_router_target", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("name", target.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "read", "set-name").GetDiag()
	}
	if err = d.Set("destination_crn", target.DestinationCRN); err != nil {
		err = fmt.Errorf("Error setting destination_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "read", "set-destination_crn").GetDiag()
	}
	if !core.IsNil(target.Region) {
		if err = d.Set("region", target.Region); err != nil {
			err = fmt.Errorf("Error setting region: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "read", "set-region").GetDiag()
		}
	}
	if !core.IsNil(target.ManagedBy) {
		if err = d.Set("managed_by", target.ManagedBy); err != nil {
			err = fmt.Errorf("Error setting managed_by: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "read", "set-managed_by").GetDiag()
		}
	}
	if err = d.Set("crn", target.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "read", "set-crn").GetDiag()
	}
	if err = d.Set("target_type", target.TargetType); err != nil {
		err = fmt.Errorf("Error setting target_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "read", "set-target_type").GetDiag()
	}
	writeStatusMap, err := ResourceIBMLogsRouterTargetWriteStatusToMap(target.WriteStatus)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "read", "write_status-to-map").GetDiag()
	}
	if err = d.Set("write_status", []map[string]interface{}{writeStatusMap}); err != nil {
		err = fmt.Errorf("Error setting write_status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "read", "set-write_status").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(target.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("updated_at", flex.DateTimeToString(target.UpdatedAt)); err != nil {
		err = fmt.Errorf("Error setting updated_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "read", "set-updated_at").GetDiag()
	}

	return nil
}

func resourceIBMLogsRouterTargetUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateTargetOptions := &logsrouterv3.UpdateTargetOptions{}

	updateTargetOptions.SetID(d.Id())

	hasChange := false

	if d.HasChange("name") {
		updateTargetOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("destination_crn") {
		updateTargetOptions.SetDestinationCRN(d.Get("destination_crn").(string))
		hasChange = true
	}

	if hasChange {
		_, _, err = logsRouterClient.UpdateTargetWithContext(context, updateTargetOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateTargetWithContext failed: %s", err.Error()), "ibm_logs_router_target", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMLogsRouterTargetRead(context, d, meta)
}

func resourceIBMLogsRouterTargetDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_target", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteTargetOptions := &logsrouterv3.DeleteTargetOptions{}

	deleteTargetOptions.SetID(d.Id())

	_, err = logsRouterClient.DeleteTargetWithContext(context, deleteTargetOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteTargetWithContext failed: %s", err.Error()), "ibm_logs_router_target", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMLogsRouterTargetWriteStatusToMap(model *logsrouterv3.WriteStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["status"] = *model.Status
	if model.LastFailure != nil {
		modelMap["last_failure"] = model.LastFailure.String()
	}
	if model.ReasonForLastFailure != nil {
		modelMap["reason_for_last_failure"] = *model.ReasonForLastFailure
	}
	return modelMap, nil
}
