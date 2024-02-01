// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func ResourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingPermit() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingPermitCreate,
		ReadContext:   resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingPermitRead,
		UpdateContext: resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingPermitUpdate,
		DeleteContext: resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingPermitDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"private_path_service_gateway": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The private path service gateway identifier.",
			},
			"endpoint_gateway_binding": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The private path service gateway identifier.",
			},
			"set_account_policy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Indicates whether this will become the access policy for any pending and future endpoint gateway bindings from the same account..",
			},
		},
	}
}

func resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingPermitCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	permitPrivatePathServiceGatewayEndpointGatewayBindingOptions := &vpcv1.PermitPrivatePathServiceGatewayEndpointGatewayBindingOptions{}

	permitPrivatePathServiceGatewayEndpointGatewayBindingOptions.SetPrivatePathServiceGatewayID(d.Get("private_path_service_gateway").(string))
	permitPrivatePathServiceGatewayEndpointGatewayBindingOptions.SetID(d.Get("endpoint_gateway_binding").(string))
	if setAccountPolicyIntf, ok := d.GetOkExists("set_account_policy"); ok {
		setAccountPolicy := setAccountPolicyIntf.(bool)
		permitPrivatePathServiceGatewayEndpointGatewayBindingOptions.SetSetAccountPolicy(setAccountPolicy)
	}

	response, err := vpcClient.PermitPrivatePathServiceGatewayEndpointGatewayBindingWithContext(context, permitPrivatePathServiceGatewayEndpointGatewayBindingOptions)
	if err != nil {
		log.Printf("[DEBUG] PermitPrivatePathServiceGatewayEndpointGatewayBindingWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PermitPrivatePathServiceGatewayEndpointGatewayBindingWithContext failed %s\n%s", err, response))
	}

	d.SetId(*permitPrivatePathServiceGatewayEndpointGatewayBindingOptions.PrivatePathServiceGatewayID)

	return resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingPermitRead(context, d, meta)
}

func resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingPermitRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingPermitUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingPermitRead(context, d, meta)
}

func resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingPermitDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	d.SetId("")

	return nil
}
