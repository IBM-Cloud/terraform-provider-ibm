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

func ResourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDeny() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDenyCreate,
		ReadContext:   resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDenyRead,
		UpdateContext: resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDenyUpdate,
		DeleteContext: resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDenyDelete,
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
				Description: "Indicates whether this will become the access policy for any pending and future endpoint gateway bindings from the same account.",
			},
		},
	}
}

func resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDenyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	denyPrivatePathServiceGatewayEndpointGatewayBindingOptions := &vpcv1.DenyPrivatePathServiceGatewayEndpointGatewayBindingOptions{}

	denyPrivatePathServiceGatewayEndpointGatewayBindingOptions.SetPrivatePathServiceGatewayID(d.Get("private_path_service_gateway").(string))
	denyPrivatePathServiceGatewayEndpointGatewayBindingOptions.SetID(d.Get("endpoint_gateway_binding").(string))
	if setAccountPolicyIntf, ok := d.GetOkExists("set_account_policy"); ok {
		setAccountPolicy := setAccountPolicyIntf.(bool)
		denyPrivatePathServiceGatewayEndpointGatewayBindingOptions.SetSetAccountPolicy(setAccountPolicy)
	}

	response, err := vpcClient.DenyPrivatePathServiceGatewayEndpointGatewayBindingWithContext(context, denyPrivatePathServiceGatewayEndpointGatewayBindingOptions)
	if err != nil {
		log.Printf("[DEBUG] DenyPrivatePathServiceGatewayEndpointGatewayBindingWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DenyPrivatePathServiceGatewayEndpointGatewayBindingWithContext failed %s\n%s", err, response))
	}

	d.SetId(*denyPrivatePathServiceGatewayEndpointGatewayBindingOptions.PrivatePathServiceGatewayID)

	return resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDenyRead(context, d, meta)
}

func resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDenyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDenyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDenyRead(context, d, meta)
}

func resourceIBMIsPrivatePathServiceGatewayEndpointGatewayBindingDenyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	d.SetId("")

	return nil
}
