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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func ResourceIBMIsPrivatePathServiceGatewayOperations() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsPrivatePathServiceGatewayOperationsCreate,
		ReadContext:   resourceIBMIsPrivatePathServiceGatewayOperationsRead,
		UpdateContext: resourceIBMIsPrivatePathServiceGatewayOperationsUpdate,
		DeleteContext: resourceIBMIsPrivatePathServiceGatewayOperationsDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"private_path_service_gateway": {
				Type:     schema.TypeString,
				Required: true,
				// ForceNew:    true,
				Description: "The private path service gateway identifier.",
			},
			"published": {
				Type:     schema.TypeBool,
				Required: true,
				// ForceNew:    true,
				Description: "Publish or unpublish PPSG.",
			},
		},
	}
}

func resourceIBMIsPrivatePathServiceGatewayOperationsCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_private_path_service_gateway_operations", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	ppsgId := d.Get("private_path_service_gateway").(string)
	publish := d.Get("published").(bool)
	if publish {
		publishPrivatePathServiceGatewayOptions := &vpcv1.PublishPrivatePathServiceGatewayOptions{}

		publishPrivatePathServiceGatewayOptions.SetPrivatePathServiceGatewayID(ppsgId)

		response, err := vpcClient.PublishPrivatePathServiceGatewayWithContext(context, publishPrivatePathServiceGatewayOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PublishPrivatePathServiceGatewayWithContext failed: %s\n%s", err.Error(), response), "ibm_is_private_path_service_gateway_operations", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

	} else {
		unpublishPrivatePathServiceGatewayOptions := &vpcv1.UnpublishPrivatePathServiceGatewayOptions{}

		unpublishPrivatePathServiceGatewayOptions.SetPrivatePathServiceGatewayID(ppsgId)

		response, err := vpcClient.UnpublishPrivatePathServiceGatewayWithContext(context, unpublishPrivatePathServiceGatewayOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UnpublishPrivatePathServiceGatewayWithContext failed: %s\n%s", err.Error(), response), "ibm_is_private_path_service_gateway_operations", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

	}

	d.SetId(ppsgId)

	return nil
}

func resourceIBMIsPrivatePathServiceGatewayOperationsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIBMIsPrivatePathServiceGatewayOperationsUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		return diag.FromErr(err)
	}
	ppsgId := d.Get("private_path_service_gateway").(string)
	publish := d.Get("published").(bool)
	if publish {
		publishPrivatePathServiceGatewayOptions := &vpcv1.PublishPrivatePathServiceGatewayOptions{}

		publishPrivatePathServiceGatewayOptions.SetPrivatePathServiceGatewayID(ppsgId)

		response, err := vpcClient.PublishPrivatePathServiceGatewayWithContext(context, publishPrivatePathServiceGatewayOptions)
		if err != nil {
			resetPublishedSchemaValue(context, d)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PublishPrivatePathServiceGatewayWithContext failed: %s\n%s", err.Error(), response), "ibm_is_private_path_service_gateway_operations", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

	} else {
		unpublishPrivatePathServiceGatewayOptions := &vpcv1.UnpublishPrivatePathServiceGatewayOptions{}

		unpublishPrivatePathServiceGatewayOptions.SetPrivatePathServiceGatewayID(ppsgId)

		response, err := vpcClient.UnpublishPrivatePathServiceGatewayWithContext(context, unpublishPrivatePathServiceGatewayOptions)
		if err != nil {
			resetPublishedSchemaValue(context, d)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UnpublishPrivatePathServiceGatewayWithContext failed: %s\n%s", err.Error(), response), "ibm_is_private_path_service_gateway_operations", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

	}
	return nil
}
func resetPublishedSchemaValue(context context.Context, d *schema.ResourceData) {
	if d.HasChange("published") {
		oldIntf, newIntf := d.GetChange("published")
		if oldIntf.(bool) != newIntf.(bool) {
			d.Set("published", oldIntf.(bool))
		}
	}
}
func resourceIBMIsPrivatePathServiceGatewayOperationsDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	d.SetId("")

	return nil
}
