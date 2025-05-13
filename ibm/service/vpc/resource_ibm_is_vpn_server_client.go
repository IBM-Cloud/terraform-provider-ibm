// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func ResourceIBMIsVPNServerClient() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsVPNServerClientDisconnect,
		ReadContext:   resourceIBMIsVPNServerClientDisconnect,
		UpdateContext: resourceIBMIsVPNServerClientDisconnect,
		DeleteContext: resourceIBMIsVPNServerClientDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"vpn_server": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPN server identifier.",
			},
			"vpn_client": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPN Client identifier.",
			},
			"delete": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "The delete to use for this VPN client to be deleted or not, when false, client is disconneted and when set to true client is deleted.",
			},
			"status_code": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "status code of the result.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "description of the result.",
			},
		},
	}
}

func resourceIBMIsVPNServerClientDisconnect(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_server_client", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getVPNServerClientOptions := &vpcv1.GetVPNServerClientOptions{}

	getVPNServerClientOptions.SetVPNServerID(d.Get("vpn_server").(string))
	getVPNServerClientOptions.SetID(d.Get("vpn_client").(string))

	_, response, err := vpcClient.GetVPNServerClientWithContext(context, getVPNServerClientOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNServerClientWithContext failed: %s", err.Error()), "ibm_is_vpn_server_client", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var flag bool
	if v, ok := d.GetOk("delete"); ok {
		flag = v.(bool)
	}

	if flag == false {

		disconnectVPNServerRouteOptions := &vpcv1.DisconnectVPNClientOptions{}
		disconnectVPNServerRouteOptions.SetVPNServerID(d.Get("vpn_server").(string))
		disconnectVPNServerRouteOptions.SetID(d.Get("vpn_client").(string))

		_, err := vpcClient.DisconnectVPNClientWithContext(context, disconnectVPNServerRouteOptions)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "[ERROR] DisconnectVPNClientWithContext failed", "ibm_is_vpn_server_client", "disconnect").GetDiag()

		}

		if err = d.Set("status_code", response.StatusCode); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "[ERROR] Error setting status_code", "ibm_is_vpn_server_client", "disconnect").GetDiag()
		}

		if err = d.Set("description", "The VPN client disconnection request was accepted."); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "[ERROR] Error setting description", "ibm_is_vpn_server_client", "disconnect").GetDiag()
		}

		d.SetId(fmt.Sprintf("%s/%s/%v", d.Get("vpn_server").(string), d.Get("vpn_client").(string), response.StatusCode))

	} else if flag == true {

		deleteVPNServerClientOptions := &vpcv1.DeleteVPNServerClientOptions{}
		deleteVPNServerClientOptions.SetVPNServerID(d.Get("vpn_server").(string))
		deleteVPNServerClientOptions.SetID(d.Get("vpn_client").(string))

		response, err := vpcClient.DeleteVPNServerClientWithContext(context, deleteVPNServerClientOptions)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "[ERROR] DeleteVPNServerClientWithContext failed", "ibm_is_vpn_server_client", "delete-client").GetDiag()

		}

		if err = d.Set("status_code", response.StatusCode); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "[ERROR] Error setting status_code", "ibm_is_vpn_server_client", "status-code").GetDiag()

		}

		if err = d.Set("description", "The VPN client disconnection request was accepted."); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "[ERROR] DisconnectVPNClientWithContext failed", "ibm_is_vpn_server_client", "description").GetDiag()

		}

		d.SetId(fmt.Sprintf("%s/%s", d.Get("vpn_server").(string), d.Get("vpn_client").(string)))
	}

	if err = d.Set("delete", d.Get("delete")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "[ERROR] Error setting delete", "ibm_is_vpn_server_client", "delete").GetDiag()

	}
	return nil
}

func resourceIBMIsVPNServerClientDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_server_client", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_server_client", "delete", "sep-id-parts").GetDiag()

	}
	if len(parts) != 2 {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "[ERROR] Incorrect ID. ID should be a combination of vpnServer/vpnClient", "ibm_is_vpn_server_client", "delete").GetDiag()

	}
	vpnServer := parts[0]
	vpnClient := parts[1]

	getVPNServerClientOptions := &vpcv1.GetVPNServerClientOptions{}

	getVPNServerClientOptions.SetVPNServerID(vpnServer)
	getVPNServerClientOptions.SetID(vpnClient)

	_, response, err := vpcClient.GetVPNServerClientWithContext(context, getVPNServerClientOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "[ERROR] GetVPNServerClientWithContext failed", "ibm_is_vpn_server_client", "delete").GetDiag()

	}

	deleteVPNServerClientOptions := &vpcv1.DeleteVPNServerClientOptions{}
	deleteVPNServerClientOptions.SetVPNServerID(d.Get("vpn_server").(string))
	deleteVPNServerClientOptions.SetID(d.Get("vpn_client").(string))

	response, err = vpcClient.DeleteVPNServerClientWithContext(context, deleteVPNServerClientOptions)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "[ERROR] DeleteVPNServerClientWithContext failed", "ibm_is_vpn_server_client", "delete").GetDiag()

	}

	d.SetId("")
	return nil
}
