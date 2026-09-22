// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement

import (
	"context"
	"fmt"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

const AdminTokenResourceName = "ibm_sm_admin_token"
const validityInterval = time.Hour                // Token validity interval
const minimumExpirationInterval = time.Minute * 5 // How long before expiration the token should be refreshed

func ResourceIbmSmAdminToken() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmAdminTokenCreate,
		ReadContext:   resourceIbmSmAdminTokenRead,
		DeleteContext: resourceIbmSmAdminTokenDelete,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The service instance ID.",
				ForceNew:    true,
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date that the admin token was created. The date format follows RFC 3339.",
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The admin token.",
				Sensitive:   true,
			},
		},
	}
}

func generateSmAdminToken(context context.Context, d *schema.ResourceData, meta interface{}, operation string) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), AdminTokenResourceName, operation, "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createTokenOptions := &secretsmanagerinstancemanagementv2.CreateVaultAdmintokenOptions{}
	createTokenOptions.SetID(d.Get("instance_id").(string))
	tokenResponse, _, err := secretsManagerInstanceManagementClient.CreateVaultAdmintokenWithContext(context, createTokenOptions)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), AdminTokenResourceName, operation, "create-token")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.Set("token", tokenResponse.Token)
	d.Set("created_at", time.Now().Format(time.RFC3339))
	return nil
}

func resourceIbmSmAdminTokenUniqueId(d *schema.ResourceData) string {
	instanceId := d.Get("instance_id").(string)
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 16)
	return fmt.Sprintf("%s:%s", instanceId, timestamp)
}

func resourceIbmSmAdminTokenCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	err := generateSmAdminToken(context, d, meta, "create")
	if err == nil {
		d.SetId(resourceIbmSmAdminTokenUniqueId(d))
	}
	return err
}

func resourceIbmSmAdminTokenRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	creationTime, err := time.Parse(time.RFC3339, d.Get("created_at").(string))
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), AdminTokenResourceName, "read", "parse-creation-time")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	isFrseh := creationTime.Add(validityInterval - minimumExpirationInterval).After(time.Now())
	if !isFrseh {
		return generateSmAdminToken(context, d, meta, "read")
	}
	return nil
}

func resourceIbmSmAdminTokenDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
