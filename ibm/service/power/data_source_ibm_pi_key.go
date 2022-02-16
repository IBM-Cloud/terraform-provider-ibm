// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

const (
	// Arguments
	PIKeyName         = "pi_key_name"
	PIKeyCreationDate = "pi_creation_date"
	PIKeyRSA          = "pi_ssh_key"

	// Attributes
	Keys            = "keys"
	KeyID           = "key_id"
	KeyRSA          = "sshkey"
	KeyCreationDate = "creation_date"
	KeyName         = "name"

	// Attributes need to fix
	KeySSHKey = "ssh_key"
)

func DataSourceIBMPIKey() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPIKeyRead,
		Schema: map[string]*schema.Schema{

			PIKeyName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "SSHKey Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			//Computed Attributes
			KeyCreationDate: {
				Type:     schema.TypeString,
				Computed: true,
			},
			KeyRSA: {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
		},
	}
}

func dataSourceIBMPIKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)

	sshkeyC := instance.NewIBMPIKeyClient(ctx, sess, cloudInstanceID)
	sshkeydata, err := sshkeyC.Get(d.Get(PIKeyName).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*sshkeydata.Name)
	d.Set(KeyCreationDate, sshkeydata.CreationDate.String())
	d.Set(KeyRSA, sshkeydata.SSHKey)

	return nil
}
