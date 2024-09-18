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

func DataSourceIbmBaasConnectionRegistrationToken() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceIbmBaasConnectionRegistrationTokenRead,
		Schema: map[string]*schema.Schema{
			"connection_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the ID of the connection, connectors belonging to which are to be fetched.",
			},
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.",
			},
			"registration_token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIbmBaasConnectionRegistrationTokenRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_connection_registration_token", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	generateDataSourceConnectionRegistrationTokenOptions := &backuprecoveryv1.GenerateDataSourceConnectionRegistrationTokenOptions{}

	generateDataSourceConnectionRegistrationTokenOptions.SetConnectionID(d.Get("connection_id").(string))
	generateDataSourceConnectionRegistrationTokenOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))

	connectionRegistrationTokenString, _, err := backupRecoveryClient.GenerateDataSourceConnectionRegistrationTokenWithContext(context, generateDataSourceConnectionRegistrationTokenOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GenerateDataSourceConnectionRegistrationTokenWithContext failed: %s", err.Error()), "ibm_baas_connection_registration_token", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(resourceIbmBaasConnectionRegistrationTokenID(d))

	if !core.IsNil(connectionRegistrationTokenString) {
		if err = d.Set("registration_token", connectionRegistrationTokenString); err != nil {
			err = fmt.Errorf("Error setting registration_token: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_connection-registration-token", "read", "set-registration_token").GetDiag()
		}
	}

	return nil
}

func resourceIbmBaasConnectionRegistrationTokenID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
