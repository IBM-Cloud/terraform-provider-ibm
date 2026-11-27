// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func ResourceIbmSmHsmDestinationApply() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmHsmDestinationApplyCreate,
		ReadContext:   resourceIbmSmHsmDestinationApplyRead,
		DeleteContext: resourceIbmSmHsmDestinationApplyDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The region of the Secrets Manager instance.",
			},
			"instance_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_destination_apply", "instance_id"),
				Description:  "Secrets Manager instance ID.",
			},
			"destination_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_destination_apply", "destination_id"),
				Description:  "Destination ID.",
			},
			"health_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Evaluated destination health state.",
			},
			"health_as_of": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the destination health snapshot.",
			},
		},
	}
}

func ResourceIbmSmHsmDestinationApplyValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "instance_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
		validate.ValidateSchema{
			Identifier:                 "destination_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sm_hsm_destination_apply", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSmHsmDestinationApplyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination_apply", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerInstanceManagementClient, d)
	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region"), "ibm_sm_hsm_destination_apply", "create")
		return tfErr.GetDiag()
	}

	applyHsmDestinationOptions := &secretsmanagerinstancemanagementv2.ApplyHsmDestinationOptions{}
	applyHsmDestinationOptions.SetInstanceID(d.Get("instance_id").(string))
	applyHsmDestinationOptions.SetDestinationID(d.Get("destination_id").(string))

	_, _, err = secretsManagerInstanceManagementClient.ApplyHsmDestinationWithContext(context, applyHsmDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ApplyHsmDestinationWithContext failed: %s", err.Error()), "ibm_sm_hsm_destination_apply", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *applyHsmDestinationOptions.InstanceID, *applyHsmDestinationOptions.DestinationID))

	_, err = waitForHsmDestinationApply(context, d, meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("waitForHsmDestinationApply failed: %s", err.Error()), "ibm_sm_hsm_destination_apply", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIbmSmHsmDestinationApplyRead(context, d, meta)
}

func waitForHsmDestinationApply(ctx context.Context, d *schema.ResourceData, meta interface{}) (interface{}, error) {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		return false, err
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return false, err
	}

	getHsmDestinationHealthSnapshotOptions := &secretsmanagerinstancemanagementv2.GetHsmDestinationHealthSnapshotOptions{}
	getHsmDestinationHealthSnapshotOptions.SetInstanceID(parts[0])
	getHsmDestinationHealthSnapshotOptions.SetDestinationID(parts[1])

	stateConf := &resource.StateChangeConf{
		Pending: []string{
			secretsmanagerinstancemanagementv2.DestinationHealthSnapshot_HealthState_InProgress,
			secretsmanagerinstancemanagementv2.DestinationHealthSnapshot_HealthState_Evaluating,
		},
		Target: []string{
			secretsmanagerinstancemanagementv2.DestinationHealthSnapshot_HealthState_Healthy,
			secretsmanagerinstancemanagementv2.DestinationHealthSnapshot_HealthState_Degraded,
			secretsmanagerinstancemanagementv2.DestinationHealthSnapshot_HealthState_Unhealthy,
		},
		Refresh: func() (interface{}, string, error) {
			snapshot, response, err := secretsManagerInstanceManagementClient.GetHsmDestinationHealthSnapshotWithContext(ctx, getHsmDestinationHealthSnapshotOptions)
			if err != nil {
				if response != nil && response.StatusCode == 404 {
					return nil, "", fmt.Errorf("destination health %s not found: %s", d.Id(), err)
				}
				return nil, "", err
			}
			if snapshot == nil || snapshot.HealthState == nil {
				return nil, "", fmt.Errorf("destination %s returned no health state", d.Id())
			}
			state := *snapshot.HealthState
			if state == secretsmanagerinstancemanagementv2.DestinationHealthSnapshot_HealthState_Unknown ||
				state == secretsmanagerinstancemanagementv2.DestinationHealthSnapshot_HealthState_NotConfigured {
				return snapshot, state, fmt.Errorf("destination %s health evaluation failed with state: %s", d.Id(), state)
			}
			return snapshot, state, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      15 * time.Second,
		MinTimeout: 15 * time.Second,
	}

	return stateConf.WaitForStateContext(ctx)
}

func resourceIbmSmHsmDestinationApplyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination_apply", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination_apply", "read", "sep-id-parts").GetDiag()
	}

	getHsmDestinationHealthSnapshotOptions := &secretsmanagerinstancemanagementv2.GetHsmDestinationHealthSnapshotOptions{}
	getHsmDestinationHealthSnapshotOptions.SetInstanceID(parts[0])
	getHsmDestinationHealthSnapshotOptions.SetDestinationID(parts[1])

	snapshot, response, err := secretsManagerInstanceManagementClient.GetHsmDestinationHealthSnapshotWithContext(context, getHsmDestinationHealthSnapshotOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetHsmDestinationHealthSnapshotWithContext failed: %s", err.Error()), "ibm_sm_hsm_destination_apply", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if snapshot.HealthState != nil {
		if err = d.Set("health_state", *snapshot.HealthState); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination_apply", "read", "set-health_state").GetDiag()
		}
	}
	if snapshot.HealthAsOf != nil {
		if err = d.Set("health_as_of", snapshot.HealthAsOf.String()); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination_apply", "read", "set-health_as_of").GetDiag()
		}
	}
	return nil
}

func resourceIbmSmHsmDestinationApplyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}
