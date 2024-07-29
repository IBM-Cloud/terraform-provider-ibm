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

func ResourceIbmRecoveryTearDown() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmRecoveryTearDownCreate,
		ReadContext:   resourceIbmRecoveryTearDownRead,
		UpdateContext: resourceIbmRecoveryTearDownUpdate,
		DeleteContext: resourceIbmRecoveryTearDownDelete,
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

func resourceIbmRecoveryTearDownCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	tearDownRecoveryByIdOptions := &backuprecoveryv1.TearDownRecoveryByIdOptions{}
	tearDownRecoveryByIdOptions.SetID(d.Get("recovery_id").(string))
	tearDownRecoveryResponse, err := backupRecoveryClient.TearDownRecoveryByIDWithContext(context, tearDownRecoveryByIdOptions)
	if err != nil {
		log.Printf("[DEBUG] tearDownRecoveryByIDWithContext failed %s\n%s", err, tearDownRecoveryResponse)
		return diag.FromErr(fmt.Errorf("tearDownRecoveryByIDWithContext failed %s\n%s", err, tearDownRecoveryResponse))
	}

	d.SetId(resourceIbmRecoveryTearDownID(d))
	return nil
}

func resourceIbmRecoveryTearDownID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmRecoveryTearDownRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}

func resourceIbmRecoveryTearDownDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIbmRecoveryTearDownUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
