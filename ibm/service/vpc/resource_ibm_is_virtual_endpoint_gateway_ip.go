// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isVirtualEndpointGatewayID                   = "gateway"
	isVirtualEndpointGatewayIPID                 = "reserved_ip"
	isVirtualEndpointGatewayIPName               = "name"
	isVirtualEndpointGatewayIPAddress            = "address"
	isVirtualEndpointGatewayIPResourceType       = "resource_type"
	isVirtualEndpointGatewayIPAutoDelete         = "auto_delete"
	isVirtualEndpointGatewayIPCreatedAt          = "created_at"
	isVirtualEndpointGatewayIPTarget             = "target"
	isVirtualEndpointGatewayIPTargetID           = "id"
	isVirtualEndpointGatewayIPTargetName         = "name"
	isVirtualEndpointGatewayIPTargetResourceType = "resource_type"
)

func ResourceIBMISEndpointGatewayIP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMisVirtualEndpointGatewayIPCreate,
		ReadContext:   resourceIBMisVirtualEndpointGatewayIPRead,
		DeleteContext: resourceIBMisVirtualEndpointGatewayIPDelete,
		Exists:        resourceIBMisVirtualEndpointGatewayIPExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			isVirtualEndpointGatewayID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Endpoint gateway ID",
			},
			isVirtualEndpointGatewayIPID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Endpoint gateway IP id",
			},
			isVirtualEndpointGatewayIPName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint gateway IP name",
			},
			isVirtualEndpointGatewayIPResourceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint gateway IP resource type",
			},
			isVirtualEndpointGatewayIPCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint gateway IP created date and time",
			},
			isVirtualEndpointGatewayIPAutoDelete: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Endpoint gateway IP auto delete",
			},
			isVirtualEndpointGatewayIPAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Endpoint gateway IP address",
			},
			isVirtualEndpointGatewayIPTarget: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Endpoint gateway detail",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVirtualEndpointGatewayIPTargetID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPs target id",
						},
						isVirtualEndpointGatewayIPTargetName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IPs target name",
						},
						isVirtualEndpointGatewayIPTargetResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Endpoint gateway resource type",
						},
					},
				},
			},
		},
	}
}

func resourceIBMisVirtualEndpointGatewayIPCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_ip", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	gatewayID := d.Get(isVirtualEndpointGatewayID).(string)
	ipID := d.Get(isVirtualEndpointGatewayIPID).(string)
	opt := sess.NewAddEndpointGatewayIPOptions(gatewayID, ipID)
	_, response, err := sess.AddEndpointGatewayIP(opt)
	if err != nil {
		log.Printf("Add Endpoint Gateway failed: %v", response)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_ip", "create", "id").GetDiag()
	}
	d.SetId(fmt.Sprintf("%s/%s", gatewayID, ipID))
	return resourceIBMisVirtualEndpointGatewayIPRead(context, d, meta)
}

func resourceIBMisVirtualEndpointGatewayIPRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_ip", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_ip", "read", "sep-id-parts").GetDiag()
	}
	gatewayID := parts[0]
	ipID := parts[1]
	opt := sess.NewGetEndpointGatewayIPOptions(gatewayID, ipID)
	result, response, err := sess.GetEndpointGatewayIP(opt)
	if err != nil {
		log.Printf("Get Endpoint Gateway IP failed: %v", response)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_ip", "read", "get-endpoint-gateway-ip").GetDiag()
	}
	d.Set(isVirtualEndpointGatewayIPID, result.ID)
	d.Set(isVirtualEndpointGatewayIPName, result.Name)
	d.Set(isVirtualEndpointGatewayIPAddress, result.Address)
	d.Set(isVirtualEndpointGatewayIPCreatedAt, (result.CreatedAt).String())
	d.Set(isVirtualEndpointGatewayIPResourceType, result.ResourceType)
	d.Set(isVirtualEndpointGatewayIPAutoDelete, result.AutoDelete)
	d.Set(isVirtualEndpointGatewayIPTarget,
		flattenEndpointGatewayIPTarget(result.Target.(*vpcv1.ReservedIPTarget)))
	return nil
}

func resourceIBMisVirtualEndpointGatewayIPDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_ip", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_ip", "delete", "sep-id-parts").GetDiag()
	}
	gatewayID := parts[0]
	ipID := parts[1]
	opt := sess.NewRemoveEndpointGatewayIPOptions(gatewayID, ipID)
	response, err := sess.RemoveEndpointGatewayIP(opt)
	if err != nil && response.StatusCode != 404 {
		log.Printf("Remove Endpoint Gateway IP failed: %v", response)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_virtual_endpoint_gateway_ip", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	d.SetId("")
	return nil
}

func resourceIBMisVirtualEndpointGatewayIPExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return false, err
	}
	if len(parts) != 2 {
		return false, fmt.Errorf("[ERROR] Incorrect ID %s: ID should be a combination of gatewayID/ipID", d.Id())
	}
	gatewayID := parts[0]
	ipID := parts[1]
	opt := sess.NewGetEndpointGatewayIPOptions(gatewayID, ipID)
	_, response, err := sess.GetEndpointGatewayIP(opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			log.Printf("Endpoint Gateway IP does not exist.")
			return false, nil
		}
		log.Printf("Error : %s", response)
		return false, err
	}
	return true, nil
}

func flattenEndpointGatewayIPTarget(target *vpcv1.ReservedIPTarget) interface{} {
	targetSlice := []interface{}{}
	targetOutput := map[string]string{}
	if target == nil {
		return targetOutput
	}
	if target.ID != nil {
		targetOutput[isVirtualEndpointGatewayIPTargetID] = *target.ID
	}
	if target.Name != nil {
		targetOutput[isVirtualEndpointGatewayIPTargetName] = *target.Name
	}
	if target.ResourceType != nil {
		targetOutput[isVirtualEndpointGatewayIPTargetResourceType] = *target.ResourceType
	}
	targetSlice = append(targetSlice, targetOutput)
	return targetSlice
}
