// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMIsVirtualNetworkInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsVirtualNetworkInterfaceCreate,
		ReadContext:   resourceIBMIsVirtualNetworkInterfaceRead,
		UpdateContext: resourceIBMIsVirtualNetworkInterfaceUpdate,
		DeleteContext: resourceIBMIsVirtualNetworkInterfaceDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceTagsCustomizeDiff(diff)
				},
			),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),

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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_volume", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "UserTags for the vni instance",
			},
			"access_tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_volume", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Access management tags for the vni instance",
			},

			"ips": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      hashIpsList,
				// DiffSuppressFunc: suppressIPsVNI,
				Description: "The reserved IPs bound to this virtual network interface.May be empty when `lifecycle_state` is `pending`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type: schema.TypeString,
							// Optional:    true,
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
							Type:     schema.TypeBool,
							Computed: true,
							// Default:          true,
							// DiffSuppressFunc: flex.ApplyOnce,
							Description: "Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this reserved IP.",
						},
						"reserved_ip": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							// Computed:    true,
							Description: "The unique identifier for this reserved IP.",
						},
						"name": &schema.Schema{
							Type: schema.TypeString,
							// Optional:    true,
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
			"protocol_state_filtering_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_virtual_network_interface", "protocol_state_filtering_mode"),
				Description:  "The protocol state filtering mode used for this virtual network interface.",
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
						"auto_delete": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.",
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
				ForceNew:    true,
				Description: "The resource group id for this virtual network interface.",
			},
			"security_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The security groups for this virtual network interface.",
			},
			"subnet": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
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
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "protocol_state_filtering_mode",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "auto, enabled, disabled",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_virtual_network_interface", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsVirtualNetworkInterfaceCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
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
		for _, v := range d.Get("ips").(*schema.Set).List() {
			value := v.(map[string]interface{})
			ipsItem, err := resourceIBMIsVirtualNetworkInterfaceMapToVirtualNetworkInterfaceIPsReservedIPPrototype(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "create", "parse-ips").GetDiag()
			}
			ips = append(ips, ipsItem)
		}
		createVirtualNetworkInterfaceOptions.SetIps(ips)
	}
	if _, ok := d.GetOk("name"); ok {
		createVirtualNetworkInterfaceOptions.SetName(d.Get("name").(string))
	}
	if psFilteringIntf, ok := d.GetOk("protocol_state_filtering_mode"); ok {
		createVirtualNetworkInterfaceOptions.SetProtocolStateFilteringMode(psFilteringIntf.(string))
	}
	if _, ok := d.GetOk("primary_ip"); ok {
		autodelete := true
		if autodeleteOk, ok := d.GetOkExists("primary_ip.0.auto_delete"); ok {
			autodelete = autodeleteOk.(bool)
		}
		primaryIPModel, err := resourceIBMIsVirtualNetworkInterfaceMapToVirtualNetworkInterfacePrimaryIPReservedIPPrototype(d.Get("primary_ip.0").(map[string]interface{}), autodelete)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "create", "parse-primary_ip").GetDiag()
		}
		createVirtualNetworkInterfaceOptions.SetPrimaryIP(primaryIPModel)
	}
	if rgOk, ok := d.GetOk("resource_group"); ok {
		rg := rgOk.(string)
		createVirtualNetworkInterfaceOptions.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
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
	// log.Printf("[INFO] vnip2 request map is %s", output(createVirtualNetworkInterfaceOptions))

	virtualNetworkInterface, _, err := sess.CreateVirtualNetworkInterfaceWithContext(context, createVirtualNetworkInterfaceOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateVirtualNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_virtual_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*virtualNetworkInterface.ID)
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk("tags"); ok || v != "" {
		oldList, newList := d.GetChange("tags")
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *virtualNetworkInterface.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource vni (%s) tags: %s", d.Id(), err)
		}
	}
	if _, ok := d.GetOk("access_tags"); ok {
		oldList, newList := d.GetChange("access_tags")
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *virtualNetworkInterface.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource vni (%s) access tags: %s", d.Id(), err)
		}
	}
	return resourceIBMIsVirtualNetworkInterfaceRead(context, d, meta)
}

func resourceIBMIsVirtualNetworkInterfaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getVirtualNetworkInterfaceOptions := &vpcv1.GetVirtualNetworkInterfaceOptions{}

	getVirtualNetworkInterfaceOptions.SetID(d.Id())

	virtualNetworkInterface, response, err := sess.GetVirtualNetworkInterfaceWithContext(context, getVirtualNetworkInterfaceOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVirtualNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_virtual_network_interface", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(virtualNetworkInterface.AllowIPSpoofing) {
		if err = d.Set("allow_ip_spoofing", virtualNetworkInterface.AllowIPSpoofing); err != nil {
			err = fmt.Errorf("Error setting allow_ip_spoofing: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-allow_ip_spoofing").GetDiag()
		}
	}
	if !core.IsNil(virtualNetworkInterface.AutoDelete) {
		if err = d.Set("auto_delete", virtualNetworkInterface.AutoDelete); err != nil {
			err = fmt.Errorf("Error setting auto_delete: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-auto_delete").GetDiag()
		}
	}
	if !core.IsNil(virtualNetworkInterface.EnableInfrastructureNat) {
		if err = d.Set("enable_infrastructure_nat", virtualNetworkInterface.EnableInfrastructureNat); err != nil {
			err = fmt.Errorf("Error setting enable_infrastructure_nat: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-enable_infrastructure_nat").GetDiag()
		}
	}
	if !core.IsNil(virtualNetworkInterface.Ips) {
		ips := []map[string]interface{}{}
		for _, ipsItem := range virtualNetworkInterface.Ips {
			if *virtualNetworkInterface.PrimaryIP.ID != *ipsItem.ID {
				ipsItemMap, err := resourceIBMIsVirtualNetworkInterfaceReservedIPReferenceToMap(&ipsItem, false)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "ips-to-map").GetDiag()
				}
				ips = append(ips, ipsItemMap)
			}
		}
		if err = d.Set("ips", ips); err != nil {
			err = fmt.Errorf("Error setting ips: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-ips").GetDiag()
		}
	}
	if !core.IsNil(virtualNetworkInterface.Name) {
		if err = d.Set("name", virtualNetworkInterface.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(virtualNetworkInterface.ProtocolStateFilteringMode) {
		if err = d.Set("protocol_state_filtering_mode", virtualNetworkInterface.ProtocolStateFilteringMode); err != nil {
			err = fmt.Errorf("Error setting protocol_state_filtering_mode: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-protocol_state_filtering_mode").GetDiag()
		}
	}
	if !core.IsNil(virtualNetworkInterface.PrimaryIP) {
		autodelete := d.Get("primary_ip.0.auto_delete").(bool)
		primaryIPMap, err := resourceIBMIsVirtualNetworkInterfaceReservedIPReferenceToMap(virtualNetworkInterface.PrimaryIP, autodelete)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "primary_ip-to-map").GetDiag()
		}
		if err = d.Set("primary_ip", []map[string]interface{}{primaryIPMap}); err != nil {
			err = fmt.Errorf("Error setting primary_ip: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-primary_ip").GetDiag()
		}
	}
	if !core.IsNil(virtualNetworkInterface.ResourceGroup) {
		if err = d.Set("resource_group", virtualNetworkInterface.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-resource_group").GetDiag()
		}
	}
	if !core.IsNil(virtualNetworkInterface.SecurityGroups) {
		securityGroups := make([]string, 0)
		for _, securityGroupsItem := range virtualNetworkInterface.SecurityGroups {
			if securityGroupsItem.ID != nil {
				securityGroups = append(securityGroups, *securityGroupsItem.ID)
			}
		}
		if err = d.Set("security_groups", securityGroups); err != nil {
			err = fmt.Errorf("Error setting security_groups: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-security_groups").GetDiag()
		}
	}
	if !core.IsNil(virtualNetworkInterface.Subnet) {
		if err = d.Set("subnet", virtualNetworkInterface.Subnet.ID); err != nil {
			err = fmt.Errorf("Error setting subnet: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-subnet").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(virtualNetworkInterface.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("crn", virtualNetworkInterface.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-crn").GetDiag()
	}
	if err = d.Set("href", virtualNetworkInterface.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-href").GetDiag()
	}
	if err = d.Set("lifecycle_state", virtualNetworkInterface.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-lifecycle_state").GetDiag()
	}
	if !core.IsNil(virtualNetworkInterface.MacAddress) {
		if err = d.Set("mac_address", virtualNetworkInterface.MacAddress); err != nil {
			err = fmt.Errorf("Error setting mac_address: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-mac_address").GetDiag()
		}
	}
	if err = d.Set("resource_type", virtualNetworkInterface.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-resource_type").GetDiag()
	}
	if !core.IsNil(virtualNetworkInterface.Target) {
		targetMap, err := resourceIBMIsVirtualNetworkInterfaceVirtualNetworkInterfaceTargetToMap(virtualNetworkInterface.Target)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "target-to-map").GetDiag()
		}
		if err = d.Set("target", []map[string]interface{}{targetMap}); err != nil {
			err = fmt.Errorf("Error setting target: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-target").GetDiag()
		}
	} else {
		d.Set("target", nil)
	}
	vpcMap, err := resourceIBMIsVirtualNetworkInterfaceVPCReferenceToMap(virtualNetworkInterface.VPC)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "vpc-to-map").GetDiag()
	}
	if err = d.Set("vpc", []map[string]interface{}{vpcMap}); err != nil {
		err = fmt.Errorf("Error setting vpc: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-vpc").GetDiag()
	}

	if virtualNetworkInterface.Zone != nil {
		if err = d.Set("zone", *virtualNetworkInterface.Zone.Name); err != nil {
			err = fmt.Errorf("Error setting zone: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-zone").GetDiag()
		}
	}

	tags, err := flex.GetGlobalTagsUsingCRN(meta, *virtualNetworkInterface.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource vni (%s) tags: %s", d.Id(), err)
	}

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *virtualNetworkInterface.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource vni (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set("tags", tags); err != nil {
		err = fmt.Errorf("Error setting tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-tags").GetDiag()
	}
	if err = d.Set("access_tags", accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "read", "set-access_tags").GetDiag()
	}

	return nil
}

func resourceIBMIsVirtualNetworkInterfaceUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	id := d.Id()

	if d.HasChange("tags") {
		oldList, newList := d.GetChange("tags")
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get("crn").(string), "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource vni (%s) tags: %s", d.Id(), err)
		}
	}

	if d.HasChange("access_tags") {
		oldList, newList := d.GetChange("access_tags")
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get("crn").(string), "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource vni (%s) access tags: %s", d.Id(), err)
		}
	}
	updateVirtualNetworkInterfaceOptions := &vpcv1.UpdateVirtualNetworkInterfaceOptions{}

	updateVirtualNetworkInterfaceOptions.SetID(id)

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

	if d.HasChange("protocol_state_filtering_mode") {
		pStateFilteringMode := d.Get("protocol_state_filtering_mode").(string)
		patchVals.ProtocolStateFilteringMode = &pStateFilteringMode
		hasChange = true
	}

	if d.HasChange("ips") {
		oldips, newips := d.GetChange("ips")
		os := oldips.(*schema.Set)
		ns := newips.(*schema.Set)
		var oldset, newset *schema.Set

		var out = make([]interface{}, ns.Len(), ns.Len())
		for i, nA := range ns.List() {
			newPack := nA.(map[string]interface{})
			out[i] = newPack["reserved_ip"].(string)
		}
		newset = schema.NewSet(schema.HashString, out)

		out = make([]interface{}, os.Len(), os.Len())
		for i, oA := range os.List() {
			oldPack := oA.(map[string]interface{})
			out[i] = oldPack["reserved_ip"].(string)
		}
		oldset = schema.NewSet(schema.HashString, out)

		remove := flex.ExpandStringList(oldset.Difference(newset).List())
		add := flex.ExpandStringList(newset.Difference(oldset).List())

		// log.Printf("[INFO] vnip2 during patch old set is %s", output(os))
		// log.Printf("[INFO] vnip2 during patch new set is %s", output(ns))

		// for _, nA := range ns.List() {
		// 	newPack := nA.(map[string]interface{})
		// 	for _, oA := range os.List() {
		// 		oldPack := oA.(map[string]interface{})
		// 		if strings.Compare(newPack["address"].(string), oldPack["address"].(string)) == 0 {
		// 			reserved_ip := oldPack["reserved_ip"].(string)
		// 			subnetId := d.Get("subnet").(string)
		// 			newName := newPack["name"].(string)
		// 			newAutoDelete := newPack["auto_delete"].(bool)

		// 			oldName := oldPack["name"].(string)
		// 			oldAutoDelete := oldPack["auto_delete"].(bool)

		// 			if newName != oldName || newAutoDelete != oldAutoDelete {

		// 				updatereservedipoptions := &vpcv1.UpdateSubnetReservedIPOptions{
		// 					SubnetID: &subnetId,
		// 					ID:       &reserved_ip,
		// 				}

		// 				reservedIpPatchModel := &vpcv1.ReservedIPPatch{}
		// 				if strings.Compare(newName, oldName) != 0 {
		// 					reservedIpPatchModel.Name = &newName
		// 				}

		// 				if newAutoDelete != oldAutoDelete {
		// 					reservedIpPatchModel.AutoDelete = &newAutoDelete
		// 				}

		// 				reservedIpPatch, err := reservedIpPatchModel.AsPatch()
		// 				if err != nil {
		// 					return diag.FromErr(fmt.Errorf("[ERROR] Error calling asPatch for ReservedIPPatch: %s", err))
		// 				}
		// 				updatereservedipoptions.ReservedIPPatch = reservedIpPatch
		// 				log.Printf("[INFO] vnip2 updatereservedipoptions %s", output(updatereservedipoptions))
		// 				_, response, err := sess.UpdateSubnetReservedIP(updatereservedipoptions)
		// 				if err != nil {
		// 					return diag.FromErr(fmt.Errorf("[ERROR] Error while updating reserved ip(%s) of vni(%s) \n%s: %q", reserved_ip, d.Id(), err, response))
		// 				}
		// 				ns.Remove(nA)
		// 				os.Remove(oA)
		// 			}
		// 		}
		// 	}
		// }
		// remove := os.Difference(ns).List()
		// log.Printf("[INFO] vnip2 remove map %s", output(remove))
		// if remove != nil && len(remove) > 0 {
		// 	subnetId := d.Get("subnet").(string)
		// 	for _, ipItem := range remove {
		// 		value := ipItem.(map[string]interface{})
		// 		if value["reserved_ip"] != nil && value["reserved_ip"].(string) != "" {
		// 			reservedipid := value["reserved_ip"].(string)
		// 			removeVirtualNetworkInterfaceIPOptions := &vpcv1.RemoveVirtualNetworkInterfaceIPOptions{}
		// 			removeVirtualNetworkInterfaceIPOptions.SetVirtualNetworkInterfaceID(id)
		// 			removeVirtualNetworkInterfaceIPOptions.SetID(reservedipid)
		// 			response, err := sess.RemoveVirtualNetworkInterfaceIPWithContext(context, removeVirtualNetworkInterfaceIPOptions)
		// 			if err != nil {
		// 				log.Printf("[DEBUG] RemoveVirtualNetworkInterfaceIPWithContext failed in VirtualNetworkInterface patch %s\n%s", err, response)
		// 				return diag.FromErr(fmt.Errorf("RemoveVirtualNetworkInterfaceIPWithContext failed in VirtualNetworkInterface patch %s\n%s", err, response))
		// 			}
		// 		}
		// 		if value["address"] != nil && value["address"].(string) != "" {
		// 			reservedipid := value["reserved_ip"].(string)
		// 			removeSubnetReservedIPOptions := &vpcv1.DeleteSubnetReservedIPOptions{}
		// 			removeSubnetReservedIPOptions.SetSubnetID(subnetId)
		// 			removeSubnetReservedIPOptions.SetID(reservedipid)
		// 			response, err := sess.DeleteSubnetReservedIPWithContext(context, removeSubnetReservedIPOptions)
		// 			if err != nil {
		// 				log.Printf("[DEBUG] DeleteSubnetReservedIPWithContext failed in VirtualNetworkInterface patch %s\n%s", err, response)
		// 				return diag.FromErr(fmt.Errorf("DeleteSubnetReservedIPWithContext failed in VirtualNetworkInterface patch %s\n%s", err, response))
		// 			}
		// 		}
		// 	}
		// }
		// add := ns.Difference(os).List()
		// log.Printf("[INFO] vnip2 add map %s", output(add))

		if add != nil && len(add) > 0 {
			for _, ipItem := range add {
				if ipItem != "" {

					addVirtualNetworkInterfaceIPOptions := &vpcv1.AddVirtualNetworkInterfaceIPOptions{}
					addVirtualNetworkInterfaceIPOptions.SetVirtualNetworkInterfaceID(id)
					addVirtualNetworkInterfaceIPOptions.SetID(ipItem)
					_, _, err := sess.AddVirtualNetworkInterfaceIPWithContext(context, addVirtualNetworkInterfaceIPOptions)
					if err != nil {
						tfErr := flex.TerraformErrorf(err, fmt.Sprintf("AddVirtualNetworkInterfaceIPWithContext failed: %s", err.Error()), "ibm_is_virtual_network_interface", "update")
						log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
						return tfErr.GetDiag()
					}
				}
			}
		}
		if remove != nil && len(remove) > 0 {
			for _, ipItem := range remove {
				if ipItem != "" {

					removeVirtualNetworkInterfaceIPOptions := &vpcv1.RemoveVirtualNetworkInterfaceIPOptions{}
					removeVirtualNetworkInterfaceIPOptions.SetVirtualNetworkInterfaceID(id)
					removeVirtualNetworkInterfaceIPOptions.SetID(ipItem)
					_, err := sess.RemoveVirtualNetworkInterfaceIPWithContext(context, removeVirtualNetworkInterfaceIPOptions)
					if err != nil {
						tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RemoveVirtualNetworkInterfaceIPWithContext failed: %s", err.Error()), "ibm_is_virtual_network_interface", "update")
						log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
						return tfErr.GetDiag()
					}
				}
			}
		}

	}
	if !d.IsNewResource() && d.HasChange("primary_ip") {
		subnetId := d.Get("subnet").(string)
		ripId := d.Get("primary_ip.0.reserved_ip").(string)
		updateripoptions := &vpcv1.UpdateSubnetReservedIPOptions{
			SubnetID: &subnetId,
			ID:       &ripId,
		}
		reservedIpPath := &vpcv1.ReservedIPPatch{}
		if d.HasChange("primary_ip.0.name") {
			name := d.Get("primary_ip.0.name").(string)
			reservedIpPath.Name = &name
		}
		if d.HasChange("primary_ip.0.auto_delete") {
			auto := d.Get("primary_ip.0.auto_delete").(bool)
			reservedIpPath.AutoDelete = &auto
		}
		reservedIpPathAsPatch, err := reservedIpPath.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("reservedIpPath.AsPatch() failed: %s", err.Error()), "ibm_is_virtual_network_interface", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateripoptions.ReservedIPPatch = reservedIpPathAsPatch
		_, _, err = sess.UpdateSubnetReservedIPWithContext(context, updateripoptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_virtual_network_interface", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	if d.HasChange("security_groups") && !d.IsNewResource() {
		ovs, nvs := d.GetChange("security_groups")
		vniId := d.Id()
		ov := ovs.(*schema.Set)
		nv := nvs.(*schema.Set)
		remove := flex.ExpandStringList(ov.Difference(nv).List())
		add := flex.ExpandStringList(nv.Difference(ov).List())
		if len(add) > 0 {
			for i := range add {
				createsgnicoptions := &vpcv1.CreateSecurityGroupTargetBindingOptions{
					SecurityGroupID: &add[i],
					ID:              &vniId,
				}
				_, _, err := sess.CreateSecurityGroupTargetBindingWithContext(context, createsgnicoptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSecurityGroupTargetBindingWithContext failed: %s", err.Error()), "ibm_is_virtual_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForVirtualNetworkInterfaceAvailable(sess, vniId, d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVirtualNetworkInterfaceAvailable failed: %s", err.Error()), "ibm_is_virtual_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}

		}
		if len(remove) > 0 {
			for i := range remove {
				deletesgnicoptions := &vpcv1.DeleteSecurityGroupTargetBindingOptions{
					SecurityGroupID: &remove[i],
					ID:              &vniId,
				}
				_, err := sess.DeleteSecurityGroupTargetBindingWithContext(context, deletesgnicoptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSecurityGroupTargetBindingWithContext failed: %s", err.Error()), "ibm_is_virtual_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForVirtualNetworkInterfaceAvailable(sess, vniId, d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVirtualNetworkInterfaceAvailable failed: %s", err.Error()), "ibm_is_virtual_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}
	}

	if hasChange {
		updateVirtualNetworkInterfaceOptions.VirtualNetworkInterfacePatch, _ = patchVals.AsPatch()
		_, _, err := sess.UpdateVirtualNetworkInterfaceWithContext(context, updateVirtualNetworkInterfaceOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVirtualNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_virtual_network_interface", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsVirtualNetworkInterfaceRead(context, d, meta)
}

func resourceIBMIsVirtualNetworkInterfaceDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_network_interface", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteVirtualNetworkInterfacesOptions := &vpcv1.DeleteVirtualNetworkInterfacesOptions{}

	deleteVirtualNetworkInterfacesOptions.SetID(d.Id())

	vni, _, err := sess.DeleteVirtualNetworkInterfacesWithContext(context, deleteVirtualNetworkInterfacesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteVirtualNetworkInterfacesWithContext failed: %s", err.Error()), "ibm_is_virtual_network_interface", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForVirtualNetworkInterfaceDeleted(sess, d.Id(), d.Timeout(schema.TimeoutDelete), vni)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVirtualNetworkInterfaceDeleted failed: %s", err.Error()), "ibm_is_virtual_network_interface", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func resourceIBMIsVirtualNetworkInterfaceMapToVirtualNetworkInterfaceIPsReservedIPPrototype(modelMap map[string]interface{}) (vpcv1.VirtualNetworkInterfaceIPPrototypeIntf, error) {
	model := &vpcv1.VirtualNetworkInterfaceIPPrototype{}
	if modelMap["reserved_ip"] != nil && modelMap["reserved_ip"].(string) != "" {
		model.ID = core.StringPtr(modelMap["reserved_ip"].(string))
	}
	// if modelMap["href"] != nil && modelMap["href"].(string) != "" {
	// 	model.Href = core.StringPtr(modelMap["href"].(string))
	// }
	// if modelMap["address"] != nil && modelMap["address"].(string) != "" {
	// 	model.Address = core.StringPtr(modelMap["address"].(string))
	// }
	// if modelMap["auto_delete"] != nil {
	// 	model.AutoDelete = core.BoolPtr(modelMap["auto_delete"].(bool))
	// }
	// if modelMap["name"] != nil && modelMap["name"].(string) != "" {
	// 	model.Name = core.StringPtr(modelMap["name"].(string))
	// }
	return model, nil
}

func resourceIBMIsVirtualNetworkInterfaceMapToVirtualNetworkInterfacePrimaryIPReservedIPPrototype(modelMap map[string]interface{}, autodelete bool) (vpcv1.VirtualNetworkInterfacePrimaryIPPrototypeIntf, error) {
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
	model.AutoDelete = core.BoolPtr(autodelete)
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	return model, nil
}

func resourceIBMIsVirtualNetworkInterfaceReservedIPReferenceToMap(model *vpcv1.ReservedIPReference, autodelete bool) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = model.Address
	modelMap["auto_delete"] = autodelete
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

func resourceIBMIsVirtualNetworkInterfaceReservedIPReferenceDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
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

func resourceIBMIsVirtualNetworkInterfaceSecurityGroupReferenceDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
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

func resourceIBMIsVirtualNetworkInterfaceShareMountTargetReferenceDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
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

func isWaitForVirtualNetworkInterfaceAvailable(client *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VirtualNetworkInterface (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"", "pending"},
		Target:     []string{"done", "failed", "stable"},
		Refresh:    isVirtualNetworkInterfaceRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVirtualNetworkInterfaceRefreshFunc(client *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vnigetoptions := &vpcv1.GetVirtualNetworkInterfaceOptions{
			ID: &id,
		}
		vni, response, err := client.GetVirtualNetworkInterface(vnigetoptions)
		if err != nil {
			return nil, "failed", fmt.Errorf("[ERROR] Error getting vni: %s\n%s", err, response)
		}
		if *vni.LifecycleState == "failed" || *vni.LifecycleState == "suspended" {
			return vni, *vni.LifecycleState, fmt.Errorf("[ERROR] Error VirtualNetworkInterface in : %s state", *vni.LifecycleState)
		}
		return vni, *vni.LifecycleState, nil
	}
}
func isWaitForVirtualNetworkInterfaceDeleted(client *vpcv1.VpcV1, id string, timeout time.Duration, vni *vpcv1.VirtualNetworkInterface) (interface{}, error) {
	log.Printf("Waiting for VirtualNetworkInterface (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"", "pending", "deleting", "updating", "waiting"},
		Target:     []string{"done", "failed", "stable", "suspended"},
		Refresh:    isVirtualNetworkInterfaceDeleteRefreshFunc(client, vni, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVirtualNetworkInterfaceDeleteRefreshFunc(client *vpcv1.VpcV1, vnir *vpcv1.VirtualNetworkInterface, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vnigetoptions := &vpcv1.GetVirtualNetworkInterfaceOptions{
			ID: &id,
		}
		vni, response, err := client.GetVirtualNetworkInterface(vnigetoptions)
		if err != nil {
			if response.StatusCode == 404 {
				return vnir, "done", nil
			}
			return vni, "failed", fmt.Errorf("[ERROR] Error getting vni: %s\n%s", err, response)
		}
		if *vni.LifecycleState == "failed" || *vni.LifecycleState == "suspended" {
			return vni, *vni.LifecycleState, fmt.Errorf("[ERROR] Error VirtualNetworkInterface in : %s state", *vni.LifecycleState)
		}
		if *vni.LifecycleState == "stable" {
			return vni, *vni.LifecycleState, fmt.Errorf("[ERROR] Error VirtualNetworkInterface in : %s state", *vni.LifecycleState)
		}
		return vni, *vni.LifecycleState, nil
	}
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

func resourceIBMIsVirtualNetworkInterfaceVPCReferenceDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = model.MoreInfo
	return modelMap, nil
}

func hashIpsList(v interface{}) int {
	var buf bytes.Buffer
	a := v.(map[string]interface{})
	// buf.WriteString(fmt.Sprintf("%s-", a["address"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", a["reserved_ip"].(string)))
	return conns.String(buf.String())
}

// func suppressIPsVNI(k, old, new string, d *schema.ResourceData) bool {
// 	oldips, newips := d.GetChange("ips")
// 	os := oldips.(*schema.Set)
// 	ns := newips.(*schema.Set)
// 	if os.Len() == ns.Len() {
// 		for _, nA := range ns.List() {
// 			newPack := nA.(map[string]interface{})
// 			for _, oA := range os.List() {
// 				oldPack := oA.(map[string]interface{})
// 				if strings.Compare(newPack["name"].(string), oldPack["address"].(string)) == 0 {
// 				}
// 			}
// 		}
// 		return true
// 	} else {
// 		return false
// 	}
// }
