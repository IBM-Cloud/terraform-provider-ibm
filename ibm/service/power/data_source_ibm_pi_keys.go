// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"log"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

// Attributes and Arguments defined in data_source_ibm_pi_key.go
func DataSourceIBMPIKeys() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIKeysRead,
		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "PI cloud instance ID",
				ValidateFunc: validation.NoZeroValues,
			},
			// Computed Attributes
			Keys: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						KeyName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User defined name for the SSH key",
						},
						KeySSHKey: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "SSH RSA key",
						},
						KeyCreationDate: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date of SSH key creation",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPIKeysRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)

	client := st.NewIBMPIKeyClient(ctx, sess, cloudInstanceID)
	sshKeys, err := client.GetAll()
	if err != nil {
		log.Printf("[ERROR] get all keys failed %v", err)
		return diag.FromErr(err)
	}

	result := make([]map[string]interface{}, 0, len(sshKeys.SSHKeys))
	for _, sshKey := range sshKeys.SSHKeys {
		key := map[string]interface{}{
			KeyName:         sshKey.Name,
			KeySSHKey:       sshKey.SSHKey,
			KeyCreationDate: sshKey.CreationDate.String(),
		}
		result = append(result, key)
	}

	var genID, _ = uuid.GenerateUUID()
	d.SetId(genID)
	d.Set(Keys, result)

	return nil
}
