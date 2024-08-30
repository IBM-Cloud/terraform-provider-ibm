// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
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
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmConnectionRegistrationToken() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmConnectionRegistrationTokenCreate,
		ReadContext:   resourceIbmConnectionRegistrationTokenRead,
		DeleteContext: resourceIbmConnectionRegistrationTokenDelete,
		UpdateContext: resourceIbmConnectionRegistrationTokenUpdate,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"connection_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the ID of the connection, connectors belonging to which are to be fetched.",
			},
			"registration_token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIbmConnectionRegistrationTokenCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_connection-registration-token", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	generateDataSourceConnectionRegistrationTokenOptions := &backuprecoveryv1.GenerateDataSourceConnectionRegistrationTokenOptions{}

	generateDataSourceConnectionRegistrationTokenOptions.SetConnectionID(d.Get("connection_id").(string))

	connectionRegistrationTokenString, _, err := backupRecoveryClient.GenerateDataSourceConnectionRegistrationTokenWithContext(context, generateDataSourceConnectionRegistrationTokenOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GenerateDataSourceConnectionRegistrationTokenWithContext failed: %s", err.Error()), "ibm_connection-registration-token", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(resourceIbmConnectionRegistrationTokenID(d))

	if !core.IsNil(connectionRegistrationTokenString) {
		if err = d.Set("registration_token", connectionRegistrationTokenString); err != nil {
			err = fmt.Errorf("Error setting registration_token: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_connection-registration-token", "read", "set-registration_token").GetDiag()
		}
	}

	return resourceIbmConnectionRegistrationTokenRead(context, d, meta)
}

func resourceIbmConnectionRegistrationTokenRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}
func resourceIbmConnectionRegistrationTokenID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmConnectionRegistrationTokenDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.

	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Not Supported",
		Detail:   "Delete operation is not supported for this resource. The resource will be removed from the terraform file but will continue to exist in the backend.",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func resourceIbmConnectionRegistrationTokenUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Update Not Supported",
		Detail:   "Update operation is not supported for this resource. No changes will be applied.",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}
