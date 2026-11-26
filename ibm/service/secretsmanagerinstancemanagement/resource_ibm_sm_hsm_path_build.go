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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func ResourceIbmSmHsmPathBuild() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmHsmPathBuildCreate,
		ReadContext:   resourceIbmSmHsmPathBuildRead,
		DeleteContext: resourceIbmSmHsmPathBuildDelete,
		Importer:      &schema.ResourceImporter{},

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
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_path_build", "instance_id"),
				Description:  "Secrets Manager instance ID.",
			},
			"destination_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_path_build", "destination_id"),
				Description:  "Destination ID.",
			},
			"path_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_path_build", "path_id"),
				Description:  "Path ID.",
			},
			"provisioning_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Always `provisioning` on a 202 response.",
			},
		},
	}
}

func ResourceIbmSmHsmPathBuildValidator() *validate.ResourceValidator {
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sm_hsm_path_build", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSmHsmPathBuildCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path_build", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerInstanceManagementClient, d)
	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region"), "ibm_sm_hsm_path_build", "create")
		return tfErr.GetDiag()
	}

	createPathBuildOptions := &secretsmanagerinstancemanagementv2.CreatePathBuildOptions{}

	createPathBuildOptions.SetInstanceID(d.Get("instance_id").(string))
	createPathBuildOptions.SetDestinationID(d.Get("destination_id").(string))
	createPathBuildOptions.SetPathID(d.Get("path_id").(string))

	_, _, err = secretsManagerInstanceManagementClient.CreatePathBuildWithContext(context, createPathBuildOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreatePathBuildWithContext failed: %s", err.Error()), "ibm_sm_hsm_path_build", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *createPathBuildOptions.InstanceID, *createPathBuildOptions.DestinationID, *createPathBuildOptions.PathID))

	return resourceIbmSmHsmPathBuildRead(context, d, meta)
}

func resourceIbmSmHsmPathBuildRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path_build", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path_build", "read", "sep-id-parts").GetDiag()
	}

	getInstanceHsmDestinationPathOptions.SetInstanceID(parts[0])
	getInstanceHsmDestinationPathOptions.SetDestinationID(parts[1])
	getInstanceHsmDestinationPathOptions.SetPathID(parts[2])

	path, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationPathWithContext(context, getInstanceHsmDestinationPathOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationPathWithContext failed: %s", err.Error()), "ibm_sm_hsm_path_build", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("provisioning_state", path.ProvisioningState); err != nil {
		err = fmt.Errorf("error setting provisioning_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path_build", "read", "set-provisioning_state").GetDiag()
	}
	return nil
}

func resourceIbmSmHsmPathBuildDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}
