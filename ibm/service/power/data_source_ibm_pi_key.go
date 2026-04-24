// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPIKey() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIKeyRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},

			Arg_KeyName: {
				AtLeastOneOf:  []string{Arg_SSHKeyID, Arg_KeyName},
				ConflictsWith: []string{Arg_SSHKeyID},
				Deprecated:    "The pi_key_name field is deprecated. Please use pi_ssh_key_id instead",
				Description:   "The name of the SSH key.",
				Optional:      true,
				Type:          schema.TypeString,
			},
			Arg_SSHKeyID: {
				AtLeastOneOf:  []string{Arg_SSHKeyID, Arg_KeyName},
				ConflictsWith: []string{Arg_KeyName},
				Description:   "The ID of the SSH key.",
				Optional:      true,
				Type:          schema.TypeString,
			},

			// Attributes
			Attr_CreationDate: {
				Computed:    true,
				Description: "Date of SSH Key creation.",
				Type:        schema.TypeString,
			},
			Attr_Description: {
				Computed:    true,
				Description: "Description of the ssh key.",
				Type:        schema.TypeString,
			},
			Attr_KeyName: {
				Computed:    true,
				Description: "Name of SSH key.",
				Type:        schema.TypeString,
			},
			Attr_PrimaryWorkspace: {
				Computed:    true,
				Description: "Indicates if the current workspace owns the ssh key or not.",
				Type:        schema.TypeBool,
			},
			Attr_SSHKey: {
				Computed:    true,
				Description: "SSH RSA key.",
				Sensitive:   true,
				Type:        schema.TypeString,
			},
			Attr_SSHKeyID: {
				Computed:    true,
				Description: "Unique ID of SSH key.",
				Type:        schema.TypeString,
			},
			Attr_Visibility: {
				Computed:    true,
				Description: "Visibility of the ssh key.",
				Type:        schema.TypeString,
			},
		},
	}
}

func dataSourceIBMPIKeyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "(Data) ibm_pi_key", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	var sshKeyID string
	if v, ok := d.GetOk(Arg_SSHKeyID); ok {
		sshKeyID = v.(string)
	} else if v, ok := d.GetOk(Arg_KeyName); ok {
		sshKeyID = v.(string)
	}

	sshkeyC := instance.NewIBMPISSHKeyClient(ctx, sess, cloudInstanceID)
	sshkeydata, err := sshkeyC.Get(sshKeyID)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Get failed: %s", err.Error()), "(Data) ibm_pi_key", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*sshkeydata.Name)
	d.Set(Attr_CreationDate, sshkeydata.CreationDate.String())
	d.Set(Attr_Description, sshkeydata.Description)
	d.Set(Attr_KeyName, sshkeydata.Name)
	d.Set(Attr_PrimaryWorkspace, sshkeydata.PrimaryWorkspace)
	d.Set(Attr_SSHKey, sshkeydata.SSHKey)
	d.Set(Attr_SSHKeyID, sshkeydata.ID)
	d.Set(Attr_Visibility, sshkeydata.Visibility)

	return nil
}
