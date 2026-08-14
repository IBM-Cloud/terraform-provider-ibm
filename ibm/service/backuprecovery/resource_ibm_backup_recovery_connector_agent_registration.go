// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmBackupRecoveryConnectorAgentRegistration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBackupRecoveryConnectorAgentRegistrationCreate,
		ReadContext:   resourceIbmBackupRecoveryConnectorAgentRegistrationRead,
		DeleteContext: resourceIbmBackupRecoveryConnectorAgentRegistrationDelete,
		UpdateContext: resourceIbmBackupRecoveryConnectorAgentRegistrationUpdate,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"registration_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				ForceNew:    true,
				Description: "The JWT registration token. A single token can be used to register multiple connector agents in that tenant. By default, the token is valid for 24 hours.",
			},
			"connection_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the name to be associated with the connector agent. This must be unique within the tenant to which this connector agent is registered.",
			},
			"join_existing_connection": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Whether this agent is joining a connection that was already claimed by a previous registration (e.g. another agent in the same cluster for clustered sources). When true, the server adds this agent to the existing connection instead of rejecting the request as a duplicate. If the connection does not yet exist, a new one is created regardless of this flag.",
			},
			"registration_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Used to indicate if a duplicate registration was attempted.",
			},
		},
	}
}

func resourceIbmBackupRecoveryConnectorAgentRegistrationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_agent_registration", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	registerConnectorAgentOptions := &backuprecoveryv1.RegisterConnectorAgentOptions{}
	if _, ok := d.GetOk("registration_token"); ok {
		registerConnectorAgentOptions.SetRegistrationToken(d.Get("registration_token").(string))
	}
	if _, ok := d.GetOk("connection_name"); ok {
		registerConnectorAgentOptions.SetConnectionName(d.Get("connection_name").(string))
	}
	if _, ok := d.GetOk("join_existing_connection"); ok {
		registerConnectorAgentOptions.SetJoinExistingConnection(d.Get("join_existing_connection").(bool))
	}

	detailedResp, err := backupRecoveryClient.RegisterConnectorAgentWithContext(context, registerConnectorAgentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RegisterConnectorAgentWithContext failed: %s", err.Error()), "ibm_backup_recovery_connector_agent_registration", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	registrationStatus := detailedResp.Headers.Get("X-Registration-Status")

	d.SetId(resourceIbmBackupRecoveryConnectorAgentRegistrationID(d))

	if !core.IsNil(registrationStatus) {
		if err = d.Set("registration_status", registrationStatus); err != nil {
			err = fmt.Errorf("Error setting registration_status: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_connector_agent_registration", "read", "set-registration_status").GetDiag()
		}
	}

	return resourceIbmBackupRecoveryConnectorAgentRegistrationRead(context, d, meta)
}

func resourceIbmBackupRecoveryConnectorAgentRegistrationID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmBackupRecoveryConnectorAgentRegistrationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIbmBackupRecoveryConnectorAgentRegistrationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.

	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Not Supported",
		Detail:   "The resource definition will be only be removed from the terraform statefile. This resource cannot be deleted from the backend. ",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func resourceIbmBackupRecoveryConnectorAgentRegistrationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "update" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Resource update will only affect terraform state and not the actual backend resource",
		Detail:   "Update operation for this resource is not supported and will only affect the terraform statefile. No changes will be made to the backend resource.",
	}
	// d.SetId("")
	diags = append(diags, warning)
	return diags
}
