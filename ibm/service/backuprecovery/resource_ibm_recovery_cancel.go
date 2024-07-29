// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmRecoveryCancel() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmRecoveryCancelCreate,
		ReadContext:   resourceIbmRecoveryCancelRead,
		UpdateContext: resourceIbmRecoveryCancelUpdate,
		DeleteContext: resourceIbmRecoveryCancelDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"recovery_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique ID.",
			},
		},
	}
}

func resourceIbmRecoveryCancelCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	cancelRecoveryByIdOptions := &backuprecoveryv1.CancelRecoveryByIdOptions{}
	cancelRecoveryByIdOptions.SetID(d.Get("recovery_id").(string))
	cancelRecoveryResponse, err := backupRecoveryClient.CancelRecoveryByIDWithContext(context, cancelRecoveryByIdOptions)
	if err != nil {
		log.Printf("[DEBUG] CancelRecoveryByIDWithContext failed %s\n%s", err, cancelRecoveryResponse)
		return diag.FromErr(fmt.Errorf("CancelRecoveryByIDWithContext failed %s\n%s", err, cancelRecoveryResponse))
	}

	d.SetId(resourceIbmRecoveryCancelID(d))
	return nil
}

func resourceIbmRecoveryCancelID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmRecoveryCancelRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}

func resourceIbmRecoveryCancelDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIbmRecoveryCancelUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
