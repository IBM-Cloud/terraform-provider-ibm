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

func ResourceIBMIsPrivatePathServiceGatewayPublish() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsPrivatePathServiceGatewayPublishCreate,
		ReadContext:   resourceIBMIsPrivatePathServiceGatewayPublishRead,
		UpdateContext: resourceIBMIsPrivatePathServiceGatewayPublishUpdate,
		DeleteContext: resourceIBMIsPrivatePathServiceGatewayPublishDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"private_path_service_gateway": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The private path service gateway identifier.",
			},
		},
	}
}

func resourceIBMIsPrivatePathServiceGatewayPublishCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	publishPrivatePathServiceGatewayOptions := &vpcv1.PublishPrivatePathServiceGatewayOptions{}

	publishPrivatePathServiceGatewayOptions.SetPrivatePathServiceGatewayID(d.Get("private_path_service_gateway").(string))

	response, err := vpcClient.PublishPrivatePathServiceGatewayWithContext(context, publishPrivatePathServiceGatewayOptions)
	if err != nil {
		log.Printf("[DEBUG] PublishPrivatePathServiceGatewayWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PublishPrivatePathServiceGatewayWithContext failed %s\n%s", err, response))
	}

	d.SetId(*publishPrivatePathServiceGatewayOptions.PrivatePathServiceGatewayID)

	return resourceIBMIsPrivatePathServiceGatewayPublishRead(context, d, meta)
}

func resourceIBMIsPrivatePathServiceGatewayPublishRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIBMIsPrivatePathServiceGatewayPublishUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIBMIsPrivatePathServiceGatewayPublishRead(context, d, meta)
}

func resourceIBMIsPrivatePathServiceGatewayPublishDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	d.SetId("")

	return nil
}
