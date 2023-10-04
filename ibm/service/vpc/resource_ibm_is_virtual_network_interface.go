// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMIsVirtualNetworkInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsVirtualNetworkInterfaceCreate,
		ReadContext:   resourceIBMIsVirtualNetworkInterfaceRead,
		UpdateContext: resourceIBMIsVirtualNetworkInterfaceUpdate,
		DeleteContext: resourceIBMIsVirtualNetworkInterfaceDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"allow_ip_spoofing": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface.",
			},
			"auto_delete": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Indicates whether this virtual network interface will be automatically deleted when`target` is deleted.",
			},
			"enable_infrastructure_nat": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the network interface,  allowing the workload to perform any needed NAT operations.- `allow_ip_spoofing` must be `false`.- If the virtual network interface is attached:  - The target `resource_type` must be `bare_metal_server_network_attachment`.  - The target `interface_type` must not be `hipersocket`.",
			},
			"ips": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The reserved IPs bound to this virtual network interface.May be empty when `lifecycle_state` is `pending`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.",
						},
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"auto_delete": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this reserved IP.",
						},
						"reserved_ip": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The unique identifier for this reserved IP.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The name for this reserved IP. The name is unique across all reserved IPs in a subnet.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_virtual_network_interface", "name"),
				Description:  "The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.",
			},
			"primary_ip": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The reserved IP for this virtual network interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.",
						},
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this reserved IP.",
						},
						"reserved_ip": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The unique identifier for this reserved IP.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The name for this reserved IP. The name is unique across all reserved IPs in a subnet.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The resource group id for this virtual network interface.",
			},
			"security_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The security groups for this virtual network interface.",
			},
			"subnet": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The associated subnet id.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the virtual network interface was created.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this virtual network interface.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this virtual network interface.",
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the virtual network interface.",
			},
			"mac_address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The MAC address of the interface. Absent when the interface is not attached to a target.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"target": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The target of this virtual network interface.If absent, this virtual network interface is not attached to a target.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this share mount target.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this share mount target.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this share mount target. The name is unique across all mount targets for the file share.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"vpc": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The VPC this virtual network interface resides in.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this VPC.",
						},
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this VPC.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this VPC.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this VPC. The name is unique across all VPCs in the region.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The zone name this virtual network interface resides in.",
			},
		},
	}
}

func ResourceIBMIsVirtualNetworkInterfaceValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_virtual_network_interface", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsVirtualNetworkInterfaceCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	createVirtualNetworkInterfaceOptions := &vpcv1.CreateVirtualNetworkInterfaceOptions{}

	if _, ok := d.GetOkExists("allow_ip_spoofing"); ok {
		createVirtualNetworkInterfaceOptions.SetAllowIPSpoofing(d.Get("allow_ip_spoofing").(bool))
	}
	if _, ok := d.GetOkExists("auto_delete"); ok {
		createVirtualNetworkInterfaceOptions.SetAutoDelete(d.Get("auto_delete").(bool))
	}
	if _, ok := d.GetOkExists("enable_infrastructure_nat"); ok {
		createVirtualNetworkInterfaceOptions.SetEnableInfrastructureNat(d.Get("enable_infrastructure_nat").(bool))
	}
	if _, ok := d.GetOk("ips"); ok {
		var ips []vpcv1.VirtualNetworkInterfaceIPPrototypeIntf
		for _, v := range d.Get("ips").([]interface{}) {
			value := v.(map[string]interface{})
			ipsItem, err := resourceIBMIsVirtualNetworkInterfaceMapToVirtualNetworkInterfaceIPsReservedIPPrototype(value)
			if err != nil {
				return diag.FromErr(err)
			}
			ips = append(ips, ipsItem)
		}
		createVirtualNetworkInterfaceOptions.SetIps(ips)
	}
	if _, ok := d.GetOk("name"); ok {
		createVirtualNetworkInterfaceOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("primary_ip"); ok {
		primaryIPModel, err := resourceIBMIsVirtualNetworkInterfaceMapToVirtualNetworkInterfacePrimaryIPReservedIPPrototype(d.Get("primary_ip.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createVirtualNetworkInterfaceOptions.SetPrimaryIP(primaryIPModel)
	}
	if _, ok := d.GetOk("resource_group"); ok {
		resourceGroupModel, err := resourceIBMIsVirtualNetworkInterfaceMapToVirtualNetworkInterfacePrototypeResourceGroup(d.Get("resource_group.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createVirtualNetworkInterfaceOptions.SetResourceGroup(resourceGroupModel)
	}
	if _, ok := d.GetOk("security_groups"); ok {
		var securityGroups []vpcv1.SecurityGroupIdentityIntf
		sg := d.Get("security_groups").(*schema.Set)
		for _, v := range sg.List() {
			value := v.(string)
			securityGroupsItem := &vpcv1.SecurityGroupIdentity{
				ID: &value,
			}
			securityGroups = append(securityGroups, securityGroupsItem)
		}
		createVirtualNetworkInterfaceOptions.SetSecurityGroups(securityGroups)
	}
	if subnetOk, ok := d.GetOk("subnet"); ok {
		subnetid := subnetOk.(string)
		subnetModel := &vpcv1.SubnetIdentity{
			ID: &subnetid,
		}
		createVirtualNetworkInterfaceOptions.SetSubnet(subnetModel)
	}

	virtualNetworkInterface, response, err := sess.CreateVirtualNetworkInterfaceWithContext(context, createVirtualNetworkInterfaceOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateVirtualNetworkInterfaceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateVirtualNetworkInterfaceWithContext failed %s\n%s", err, response))
	}

	d.SetId(*virtualNetworkInterface.ID)

	return resourceIBMIsVirtualNetworkInterfaceRead(context, d, meta)
}

func resourceIBMIsVirtualNetworkInterfaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	getVirtualNetworkInterfaceOptions := &vpcv1.GetVirtualNetworkInterfaceOptions{}

	getVirtualNetworkInterfaceOptions.SetID(d.Id())

	virtualNetworkInterface, response, err := sess.GetVirtualNetworkInterfaceWithContext(context, getVirtualNetworkInterfaceOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetVirtualNetworkInterfaceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetVirtualNetworkInterfaceWithContext failed %s\n%s", err, response))
	}

	if !core.IsNil(virtualNetworkInterface.AllowIPSpoofing) {
		if err = d.Set("allow_ip_spoofing", virtualNetworkInterface.AllowIPSpoofing); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting allow_ip_spoofing: %s", err))
		}
	}
	if !core.IsNil(virtualNetworkInterface.AutoDelete) {
		if err = d.Set("auto_delete", virtualNetworkInterface.AutoDelete); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting auto_delete: %s", err))
		}
	}
	if !core.IsNil(virtualNetworkInterface.EnableInfrastructureNat) {
		if err = d.Set("enable_infrastructure_nat", virtualNetworkInterface.EnableInfrastructureNat); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting enable_infrastructure_nat: %s", err))
		}
	}
	if !core.IsNil(virtualNetworkInterface.Ips) {
		ips := []map[string]interface{}{}
		for _, ipsItem := range virtualNetworkInterface.Ips {
			ipsItemMap, err := resourceIBMIsVirtualNetworkInterfaceReservedIPReferenceToMap(&ipsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			ips = append(ips, ipsItemMap)
		}
		if err = d.Set("ips", ips); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting ips: %s", err))
		}
	}
	if !core.IsNil(virtualNetworkInterface.Name) {
		if err = d.Set("name", virtualNetworkInterface.Name); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
		}
	}
	if !core.IsNil(virtualNetworkInterface.PrimaryIP) {
		primaryIPMap, err := resourceIBMIsVirtualNetworkInterfaceReservedIPReferenceToMap(virtualNetworkInterface.PrimaryIP)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("primary_ip", []map[string]interface{}{primaryIPMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting primary_ip: %s", err))
		}
	}
	if !core.IsNil(virtualNetworkInterface.ResourceGroup) {
		d.Set("resource_group", virtualNetworkInterface.ResourceGroup.ID)
	}
	if !core.IsNil(virtualNetworkInterface.SecurityGroups) {
		securityGroups := make([]string, 0)
		for _, securityGroupsItem := range virtualNetworkInterface.SecurityGroups {
			if securityGroupsItem.ID != nil {
				securityGroups = append(securityGroups, *securityGroupsItem.ID)
			}
		}
		if err = d.Set("security_groups", securityGroups); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting security_groups for vni: %s", err))
		}
	}
	if !core.IsNil(virtualNetworkInterface.Subnet) {
		d.Set("subnet", virtualNetworkInterface.Subnet.ID)
	}
	if err = d.Set("created_at", flex.DateTimeToString(virtualNetworkInterface.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("crn", virtualNetworkInterface.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("href", virtualNetworkInterface.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}
	if err = d.Set("lifecycle_state", virtualNetworkInterface.LifecycleState); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting lifecycle_state: %s", err))
	}
	if !core.IsNil(virtualNetworkInterface.MacAddress) {
		if err = d.Set("mac_address", virtualNetworkInterface.MacAddress); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting mac_address: %s", err))
		}
	}
	if err = d.Set("resource_type", virtualNetworkInterface.ResourceType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_type: %s", err))
	}
	if !core.IsNil(virtualNetworkInterface.Target) {
		targetMap, err := resourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetToMap(virtualNetworkInterface.Target)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("target", []map[string]interface{}{targetMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting target: %s", err))
		}
	} else {
		d.Set("target", nil)
	}
	vpcMap, err := resourceIBMIsVirtualNetworkInterfaceVPCReferenceToMap(virtualNetworkInterface.VPC)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("vpc", []map[string]interface{}{vpcMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting vpc: %s", err))
	}

	if virtualNetworkInterface.Zone != nil {
		d.Set("zone", *virtualNetworkInterface.Zone.Name)
	}

	return nil
}

func resourceIBMIsVirtualNetworkInterfaceUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	updateVirtualNetworkInterfaceOptions := &vpcv1.UpdateVirtualNetworkInterfaceOptions{}

	updateVirtualNetworkInterfaceOptions.SetID(d.Id())

	hasChange := false

	patchVals := &vpcv1.VirtualNetworkInterfacePatch{}
	if d.HasChange("allow_ip_spoofing") {
		newAllowIPSpoofing := d.Get("allow_ip_spoofing").(bool)
		patchVals.AllowIPSpoofing = &newAllowIPSpoofing
		hasChange = true
	}
	if d.HasChange("auto_delete") {
		newAutoDelete := d.Get("auto_delete").(bool)
		patchVals.AutoDelete = &newAutoDelete
		hasChange = true
	}
	if d.HasChange("enable_infrastructure_nat") {
		newEnableInfrastructureNat := d.Get("enable_infrastructure_nat").(bool)
		patchVals.EnableInfrastructureNat = &newEnableInfrastructureNat
		hasChange = true
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}

	if hasChange {
		updateVirtualNetworkInterfaceOptions.VirtualNetworkInterfacePatch, _ = patchVals.AsPatch()
		_, response, err := sess.UpdateVirtualNetworkInterfaceWithContext(context, updateVirtualNetworkInterfaceOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateVirtualNetworkInterfaceWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateVirtualNetworkInterfaceWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIBMIsVirtualNetworkInterfaceRead(context, d, meta)
}

func resourceIBMIsVirtualNetworkInterfaceDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	deleteVirtualNetworkInterfacesOptions := &vpcv1.DeleteVirtualNetworkInterfacesOptions{}

	deleteVirtualNetworkInterfacesOptions.SetID(d.Id())

	response, err := sess.DeleteVirtualNetworkInterfacesWithContext(context, deleteVirtualNetworkInterfacesOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteVirtualNetworkInterfacesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteVirtualNetworkInterfacesWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIBMIsVirtualNetworkInterfaceMapToVirtualNetworkInterfaceIPsReservedIPPrototype(modelMap map[string]interface{}) (vpcv1.VirtualNetworkInterfaceIPPrototypeIntf, error) {
	model := &vpcv1.VirtualNetworkInterfaceIPPrototype{}
	if modelMap["reserved_ip"] != nil && modelMap["reserved_ip"].(string) != "" {
		model.ID = core.StringPtr(modelMap["reserved_ip"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	if modelMap["address"] != nil && modelMap["address"].(string) != "" {
		model.Address = core.StringPtr(modelMap["address"].(string))
	}
	if modelMap["auto_delete"] != nil {
		model.AutoDelete = core.BoolPtr(modelMap["auto_delete"].(bool))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	return model, nil
}

func resourceIBMIsVirtualNetworkInterfaceMapToVirtualNetworkInterfacePrimaryIPReservedIPPrototype(modelMap map[string]interface{}) (vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeIntf, error) {
	model := &vpcv1.VirtualNetworkInterfacePrimaryIPPrototype{}
	if modelMap["reserved_ip"] != nil && modelMap["reserved_ip"].(string) != "" {
		model.ID = core.StringPtr(modelMap["reserved_ip"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	if modelMap["address"] != nil && modelMap["address"].(string) != "" {
		model.Address = core.StringPtr(modelMap["address"].(string))
	}
	if modelMap["auto_delete"] != nil {
		model.AutoDelete = core.BoolPtr(modelMap["auto_delete"].(bool))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	return model, nil
}

func resourceIBMIsVirtualNetworkInterfaceMapToVirtualNetworkInterfacePrototypeResourceGroup(modelMap map[string]interface{}) (vpcv1.ResourceGroupIdentityIntf, error) {
	model := &vpcv1.ResourceGroupIdentity{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	return model, nil
}

func resourceIBMIsVirtualNetworkInterfaceReservedIPReferenceToMap(model *vpcv1.ReservedIPReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = model.Address
	if model.Deleted != nil {
		deletedMap, err := resourceIBMIsVirtualNetworkInterfaceReservedIPReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = model.Href
	modelMap["reserved_ip"] = model.ID
	modelMap["name"] = model.Name
	modelMap["resource_type"] = model.ResourceType
	return modelMap, nil
}

func resourceIBMIsVirtualNetworkInterfaceReservedIPReferenceDeletedToMap(model *vpcv1.ReservedIPReferenceDeleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = model.MoreInfo
	return modelMap, nil
}

func resourceIBMIsVirtualNetworkInterfaceSecurityGroupReferenceToMap(model *vpcv1.SecurityGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = model.CRN
	if model.Deleted != nil {
		deletedMap, err := resourceIBMIsVirtualNetworkInterfaceSecurityGroupReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = model.Href
	modelMap["security_group"] = model.ID
	modelMap["name"] = model.Name
	return modelMap, nil
}

func resourceIBMIsVirtualNetworkInterfaceSecurityGroupReferenceDeletedToMap(model *vpcv1.SecurityGroupReferenceDeleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = model.MoreInfo
	return modelMap, nil
}

func resourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetToMap(model vpcv1.VirtualNetworkInterfaceTargetIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.VirtualNetworkInterfaceTargetShareMountTargetReference); ok {
		return resourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetShareMountTargetReferenceToMap(model.(*vpcv1.VirtualNetworkInterfaceTargetShareMountTargetReference))
	} else if _, ok := model.(*vpcv1.VirtualNetworkInterfaceTargetInstanceNetworkAttachmentReferenceVirtualNetworkInterfaceContext); ok {
		return resourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetInstanceNetworkAttachmentReferenceVirtualNetworkInterfaceContextToMap(model.(*vpcv1.VirtualNetworkInterfaceTargetInstanceNetworkAttachmentReferenceVirtualNetworkInterfaceContext))
	} else if _, ok := model.(*vpcv1.VirtualNetworkInterfaceTargetBareMetalServerNetworkAttachmentReferenceVirtualNetworkInterfaceContext); ok {
		return resourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetBareMetalServerNetworkAttachmentReferenceVirtualNetworkInterfaceContextToMap(model.(*vpcv1.VirtualNetworkInterfaceTargetBareMetalServerNetworkAttachmentReferenceVirtualNetworkInterfaceContext))
	} else if _, ok := model.(*vpcv1.VirtualNetworkInterfaceTarget); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.VirtualNetworkInterfaceTarget)
		if model.Deleted != nil {
			deletedMap, err := resourceIBMIsVirtualNetworkInterfaceShareMountTargetReferenceDeletedToMap(model.Deleted)
			if err != nil {
				return modelMap, err
			}
			modelMap["deleted"] = []map[string]interface{}{deletedMap}
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		if model.ResourceType != nil {
			modelMap["resource_type"] = model.ResourceType
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.VirtualNetworkInterfaceTargetIntf subtype encountered")
	}
}

func resourceIBMIsVirtualNetworkInterfaceShareMountTargetReferenceDeletedToMap(model *vpcv1.ShareMountTargetReferenceDeleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = model.MoreInfo
	return modelMap, nil
}

func resourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetShareMountTargetReferenceToMap(model *vpcv1.VirtualNetworkInterfaceTargetShareMountTargetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := resourceIBMIsVirtualNetworkInterfaceShareMountTargetReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = model.Href
	modelMap["id"] = model.ID
	modelMap["name"] = model.Name
	modelMap["resource_type"] = model.ResourceType
	return modelMap, nil
}

func resourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetInstanceNetworkAttachmentReferenceVirtualNetworkInterfaceContextToMap(model *vpcv1.VirtualNetworkInterfaceTargetInstanceNetworkAttachmentReferenceVirtualNetworkInterfaceContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	// if model.Deleted != nil {
	// 	deletedMap, err := resourceIBMIsVirtualNetworkInterfaceInstanceNetworkAttachmentReferenceVirtualNetworkInterfaceContextDeletedToMap(model.Deleted)
	// 	if err != nil {
	// 		return modelMap, err
	// 	}
	// 	modelMap["deleted"] = []map[string]interface{}{deletedMap}
	// }
	modelMap["href"] = model.Href
	modelMap["id"] = model.ID
	modelMap["name"] = model.Name
	modelMap["resource_type"] = model.ResourceType
	return modelMap, nil
}

// func resourceIBMIsVirtualNetworkInterfaceInstanceNetworkAttachmentReferenceVirtualNetworkInterfaceContextDeletedToMap(model *vpcv1.InstanceNetworkAttachmentReferenceVirtualNetworkInterfaceContextDeleted) (map[string]interface{}, error) {
// 	modelMap := make(map[string]interface{})
// 	modelMap["more_info"] = model.MoreInfo
// 	return modelMap, nil
// }

func resourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetBareMetalServerNetworkAttachmentReferenceVirtualNetworkInterfaceContextToMap(model *vpcv1.VirtualNetworkInterfaceTargetBareMetalServerNetworkAttachmentReferenceVirtualNetworkInterfaceContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	// if model.Deleted != nil {
	// 	deletedMap, err := resourceIBMIsVirtualNetworkInterfaceBareMetalServerNetworkAttachmentReferenceVirtualNetworkInterfaceContextDeletedToMap(model.Deleted)
	// 	if err != nil {
	// 		return modelMap, err
	// 	}
	// 	modelMap["deleted"] = []map[string]interface{}{deletedMap}
	// }
	modelMap["href"] = model.Href
	modelMap["id"] = model.ID
	modelMap["name"] = model.Name
	modelMap["resource_type"] = model.ResourceType
	return modelMap, nil
}

// func resourceIBMIsVirtualNetworkInterfaceBareMetalServerNetworkAttachmentReferenceVirtualNetworkInterfaceContextDeletedToMap(model *vpcv1.BareMetalServerNetworkAttachmentReferenceVirtualNetworkInterfaceContextDeleted) (map[string]interface{}, error) {
// 	modelMap := make(map[string]interface{})
// 	modelMap["more_info"] = model.MoreInfo
// 	return modelMap, nil
// }

func resourceIBMIsVirtualNetworkInterfaceVPCReferenceToMap(model *vpcv1.VPCReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = model.CRN
	if model.Deleted != nil {
		deletedMap, err := resourceIBMIsVirtualNetworkInterfaceVPCReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = model.Href
	modelMap["id"] = model.ID
	modelMap["name"] = model.Name
	modelMap["resource_type"] = model.ResourceType
	return modelMap, nil
}

func resourceIBMIsVirtualNetworkInterfaceVPCReferenceDeletedToMap(model *vpcv1.VPCReferenceDeleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = model.MoreInfo
	return modelMap, nil
}
