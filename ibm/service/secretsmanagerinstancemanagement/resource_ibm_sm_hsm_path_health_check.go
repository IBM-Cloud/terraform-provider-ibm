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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func ResourceIbmSmHsmPathHealthCheck() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmHsmPathHealthCheckCreate,
		ReadContext:   resourceIbmSmHsmPathHealthCheckRead,
		DeleteContext: resourceIbmSmHsmPathHealthCheckDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_path_health_check", "instance_id"),
				Description:  "Secrets Manager instance ID.",
			},
			"destination_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_path_health_check", "destination_id"),
				Description:  "Destination ID.",
			},
			"path_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_path_health_check", "path_id"),
				Description:  "Path ID.",
			},
			"health_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Latest health state of the path.",
			},
			"health_time_stamp": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last health evaluation.",
			},
			"provisioning_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current provisioning status of the path.",
			},
			"vpc_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current VPC endpoint gateway status.",
			},
		},
	}
}

func ResourceIbmSmHsmPathHealthCheckValidator() *validate.ResourceValidator {
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
		validate.ValidateSchema{
			Identifier:                 "path_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sm_hsm_path_health_check", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSmHsmPathHealthCheckCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path_health_check", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPathVpeStatusOptions := &secretsmanagerinstancemanagementv2.GetPathVpeStatusOptions{}
	getPathVpeStatusOptions.SetInstanceID(d.Get("instance_id").(string))
	getPathVpeStatusOptions.SetDestinationID(d.Get("destination_id").(string))
	getPathVpeStatusOptions.SetPathID(d.Get("path_id").(string))

	result, _, err := secretsManagerInstanceManagementClient.GetPathVpeStatusWithContext(context, getPathVpeStatusOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPathVpeStatusWithContext failed: %s", err.Error()), "ibm_sm_hsm_path_health_check", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *getPathVpeStatusOptions.InstanceID, *getPathVpeStatusOptions.DestinationID, *getPathVpeStatusOptions.PathID))

	if result.HealthState != nil {
		if err = d.Set("health_state", *result.HealthState); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path_health_check", "create", "set-health_state").GetDiag()
		}
	}
	if result.HealthTimeStamp != nil {
		if err = d.Set("health_time_stamp", result.HealthTimeStamp.String()); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path_health_check", "create", "set-health_time_stamp").GetDiag()
		}
	}
	if result.ProvisioningState != nil {
		if err = d.Set("provisioning_state", *result.ProvisioningState); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path_health_check", "create", "set-provisioning_state").GetDiag()
		}
	}

	return resourceIbmSmHsmPathHealthCheckRead(context, d, meta)
}

func resourceIbmSmHsmPathHealthCheckRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path_health_check", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path_health_check", "read", "sep-id-parts").GetDiag()
	}

	getInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions{}
	getInstanceHsmDestinationPathOptions.SetInstanceID(parts[0])
	getInstanceHsmDestinationPathOptions.SetDestinationID(parts[1])
	getInstanceHsmDestinationPathOptions.SetPathID(parts[2])

	path, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationPathWithContext(context, getInstanceHsmDestinationPathOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationPathWithContext failed: %s", err.Error()), "ibm_sm_hsm_path_health_check", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if path.ProvisioningState != nil {
		if err = d.Set("provisioning_state", *path.ProvisioningState); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path_health_check", "read", "set-provisioning_state").GetDiag()
		}
	}
	if path.HealthState != nil {
		if err = d.Set("health_state", *path.HealthState); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path_health_check", "read", "set-health_state").GetDiag()
		}
	}
	return nil
}

func resourceIbmSmHsmPathHealthCheckDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}
