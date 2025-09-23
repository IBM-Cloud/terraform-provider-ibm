// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceIBMPIKey() *schema.Resource {
	return &schema.Resource{
		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return customizeNameAndSSHKeyPIKeyDiff(diff)
			},
		),
		CreateContext: resourceIBMPIKeyCreate,
		ReadContext:   resourceIBMPIKeyRead,
		UpdateContext: resourceIBMPIKeyUpdate,
		DeleteContext: resourceIBMPIKeyDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_Description: {
				Description: "Description of the ssh key.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			Arg_KeyName: {
				Description:  "User defined name for the SSH key.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_SSHKey: {
				Description:  "SSH RSA key.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_Visibility: {
				Default:      Workspace,
				Description:  "Visibility of the ssh key. Valid values are: [\"account\", \"workspace\"].",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{Account, Workspace}, false),
			},
			// Attributes
			Attr_CreationDate: {
				Computed:    true,
				Description: "Date of SSH Key creation.",
				Type:        schema.TypeString,
			},
			Attr_Name: {
				Computed:    true,
				Deprecated:  "This field is deprecated and will be removed in a future release. Use pi_key_name instead.",
				Description: "User defined name for the SSH key.",
				Type:        schema.TypeString,
			},
			Attr_Key: {
				Computed:    true,
				Deprecated:  "This field is deprecated and will be removed in a future release. Use pi_ssh_key instead.",
				Description: "SSH RSA key.",
				Type:        schema.TypeString,
			},
			Attr_PrimaryWorkspace: {
				Computed:    true,
				Description: "Indicates if the current workspace owns the ssh key or not.",
				Type:        schema.TypeBool,
			},
			Attr_SSHKeyID: {
				Computed:    true,
				Description: "Unique ID of SSH key.",
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceIBMPIKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// session
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	// arguments
	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	name := d.Get(Arg_KeyName).(string)
	sshkey := d.Get(Arg_SSHKey).(string)
	visibility := d.Get(Arg_Visibility).(string)

	// create key
	client := instance.NewIBMPISSHKeyClient(ctx, sess, cloudInstanceID)
	body := &models.CreateWorkspaceSSHKey{
		Name:       &name,
		SSHKey:     &sshkey,
		Visibility: &visibility,
	}

	if v, ok := d.GetOk(Arg_Description); ok {
		description := v.(string)
		body.Description = description
	}

	sshResponse, err := client.Create(body)
	if err != nil {
		log.Printf("[DEBUG]  err %s", err)
		return diag.FromErr(err)
	}

	log.Printf("Printing the sshkey %+v", *sshResponse)
	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *sshResponse.ID))
	return resourceIBMPIKeyRead(ctx, d, meta)
}

func resourceIBMPIKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// session
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	// arguments
	cloudInstanceID, key, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// get key
	sshkeyC := instance.NewIBMPISSHKeyClient(ctx, sess, cloudInstanceID)
	sshkeydata, err := sshkeyC.Get(key)
	if err != nil {
		return diag.FromErr(err)
	}

	// Arguments
	d.Set(Arg_CloudInstanceID, cloudInstanceID)
	d.Set(Arg_Description, sshkeydata.Description)
	d.Set(Arg_KeyName, sshkeydata.Name)
	d.Set(Arg_SSHKey, sshkeydata.SSHKey)
	d.Set(Arg_Visibility, sshkeydata.Visibility)

	// Attributes
	d.Set(Attr_CreationDate, sshkeydata.CreationDate.String())
	d.Set(Attr_Key, sshkeydata.SSHKey)
	d.Set(Attr_Name, sshkeydata.Name)
	d.Set(Attr_PrimaryWorkspace, sshkeydata.PrimaryWorkspace)
	d.Set(Attr_SSHKeyID, sshkeydata.ID)

	return nil
}

func resourceIBMPIKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// session
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	// arguments
	cloudInstanceID, key, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	client := instance.NewIBMPISSHKeyClient(ctx, sess, cloudInstanceID)
	updateBody := &models.UpdateWorkspaceSSHKey{}

	if d.HasChange(Arg_Description) {
		newDescription := d.Get(Arg_Description).(string)
		updateBody.Description = &newDescription
	}

	if d.HasChange(Arg_KeyName) {
		newKeyName := d.Get(Arg_KeyName).(string)
		updateBody.Name = &newKeyName
	}

	if d.HasChange(Arg_SSHKey) {
		newSSHKey := d.Get(Arg_SSHKey).(string)
		updateBody.SSHKey = &newSSHKey
	}

	if d.HasChange(Arg_Visibility) {
		newVisibility := d.Get(Arg_Visibility).(string)
		updateBody.Visibility = &newVisibility
	}

	_, err = client.Update(key, updateBody)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIBMPIKeyRead(ctx, d, meta)
}

func resourceIBMPIKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// session
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	// arguments
	cloudInstanceID, key, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// delete key
	sshkeyC := instance.NewIBMPISSHKeyClient(ctx, sess, cloudInstanceID)
	err = sshkeyC.Delete(key)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

func customizeNameAndSSHKeyPIKeyDiff(diff *schema.ResourceDiff) error {
	if diff.Id() != "" && diff.HasChange(Arg_KeyName) {
		diff.SetNewComputed(Attr_Name)
	}
	if diff.Id() != "" && diff.HasChange(Arg_SSHKey) {
		diff.SetNewComputed(Attr_Key)
	}
	return nil
}
