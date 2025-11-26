// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerNetworkInterfaceAvailable   = "available"
	isBareMetalServerNetworkInterfaceDeleting    = "deleting"
	isBareMetalServerNetworkInterfacePending     = "pending"
	isBareMetalServerNetworkInterfacePCIPending  = "pci_pending"
	isBareMetalServerNetworkInterfaceVlanPending = "vlan_pending"
	isBareMetalServerNetworkInterfaceDeleted     = "deleted"
	isBareMetalServerNetworkInterfaceFailed      = "failed"
	isBareMetalServerHardStop                    = "hard_stop"
)

func ResourceIBMIsBareMetalServerNetworkInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISBareMetalServerNetworkInterfaceCreate,
		ReadContext:   resourceIBMISBareMetalServerNetworkInterfaceRead,
		UpdateContext: resourceIBMISBareMetalServerNetworkInterfaceUpdate,
		DeleteContext: resourceIBMISBareMetalServerNetworkInterfaceDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isBareMetalServerID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Bare metal server identifier",
			},
			isBareMetalServerNicID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The bare metal server network interface identifier",
			},
			isBareMetalServerNicAllowIPSpoofing: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.",
			},
			isBareMetalServerHardStop: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Only used for PCI network interfaces, whether to hard/immediately stop server",
			},
			isBareMetalServerNicEnableInfraNAT: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "If true, the VPC infrastructure performs any needed NAT operations. If false, the packet is passed unmodified to/from the network interface, allowing the workload to perform any needed NAT operations.",
			},
			isBareMetalServerNicFloatingIPs: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The floating IPs associated with this network interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerNicIpAddress: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique IP address",
						},
						isBareMetalServerNicFloatingIPId: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique IP identifier",
						},
					},
				},
			},
			isBareMetalServerNicHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this network interface",
			},
			isBareMetalServerNicInterfaceType: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_bare_metal_server_network_interface", isBareMetalServerNicInterfaceType),
				Description:  "The network interface type: [ pci, vlan, hipersocket ]",
			},
			isBareMetalServerNicMacAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The MAC address of the interface. If absent, the value is not known.",
			},
			isBareMetalServerNicName: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The user-defined name for this network interface",
			},
			isBareMetalServerNicPortSpeed: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The network interface port speed in Mbps",
			},
			isBareMetalServerNicPrimaryIP: {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "title: IPv4, The IP address. ",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerNicIpAddress: {
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"primary_ip.0.reserved_ip"},
							Description:   "The globally unique IP address",
						},
						isBareMetalServerNicIpHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this reserved IP",
						},
						isBareMetalServerNicIpID: {
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"primary_ip.0.address", "primary_ip.0.auto_delete", "primary_ip.0.name"},
							Description:   "The unique identifier for this reserved IP",
						},
						isBareMetalServerNicIpName: {
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"primary_ip.0.reserved_ip"},
							Description:   "The unique user-defined name for this reserved IP",
						},
						isBareMetalServerNicResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type: [ subnet_reserved_ip ]",
						},
						isBareMetalServerNicIpAutoDelete: {
							Type:          schema.TypeBool,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"primary_ip.0.reserved_ip"},
							Description:   "If set to true, this reserved IP will be automatically deleted when the target is deleted or when the reserved IP is unbound.",
						},
					},
				},
			},
			isBareMetalServerNicResourceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type : [ subnet_reserved_ip ]",
			},

			isBareMetalServerNicSecurityGroups: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "Collection of security groups ids",
			},

			isBareMetalServerNicStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the network interface : [ available, deleting, failed, pending ]",
			},

			isBareMetalServerNicSubnet: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the associated subnet",
			},

			isBareMetalServerNicType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this bare metal server network interface : [ primary, secondary ]",
			},

			isBareMetalServerNicAllowedVlans: {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isBareMetalServerNicAllowInterfaceToFloat, isBareMetalServerNicVlan},
				Elem:          &schema.Schema{Type: schema.TypeInt},
				Set:           schema.HashInt,
				Description:   "Indicates what VLAN IDs (for VLAN type only) can use this physical (PCI type) interface. A given VLAN can only be in the allowed_vlans array for one PCI type adapter per bare metal server.",
			},

			isBareMetalServerNicAllowInterfaceToFloat: {
				Type:             schema.TypeBool,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				ConflictsWith:    []string{isBareMetalServerNicAllowedVlans},
				Description:      "Indicates if the interface can float to any other server within the same resource_group. The interface will float automatically if the network detects a GARP or RARP on another bare metal server in the resource group. Applies only to vlan type interfaces.",
			},

			isBareMetalServerNicVlan: {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isBareMetalServerNicAllowedVlans},
				Description:   "Indicates the 802.1Q VLAN ID tag that must be used for all traffic on this interface",
			},
		},
	}
}

func ResourceIBMIsBareMetalServerNetworkInterfaceValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 1)
	interface_types := "pci, hipersocket, vlan"
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isBareMetalServerNicInterfaceType,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			Default:                    "pci",
			AllowedValues:              interface_types})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isBareMetalServerName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISBareMetalServerNicResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_bare_metal_server_network_interface", Schema: validateSchema}
	return &ibmISBareMetalServerNicResourceValidator
}

func resourceIBMISBareMetalServerNetworkInterfaceCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	bareMetalServerId := ""
	if bmsId, ok := d.GetOk(isBareMetalServerID); ok {
		bareMetalServerId = bmsId.(string)
	}

	if allowedVlansOk, ok := d.GetOk(isBareMetalServerNicAllowedVlans); ok {
		sess, err := vpcClient(meta)
		if err != nil {
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "create", "initialize-client")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options := &vpcv1.CreateBareMetalServerNetworkInterfaceOptions{}
		interfaceType := "pci"
		// to create pci, server needs to be in stopped state

		getbmsoptions := &vpcv1.GetBareMetalServerOptions{
			ID: &bareMetalServerId,
		}

		bms, _, err := sess.GetBareMetalServerWithContext(context, getbmsoptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		// failed, pending, restarting, running, starting, stopped, stopping, maintenance
		if *bms.Status == "failed" {
			err = fmt.Errorf("Error cannot attach network interface to a failed bare metal server")
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		} else if *bms.Status == "running" {
			log.Printf("[DEBUG] Stopping bare metal server (%s) to create a PCI network interface", bareMetalServerId)
			stopType := "hard"
			if _, ok := d.GetOk(isBareMetalServerHardStop); ok && !d.Get(isBareMetalServerHardStop).(bool) {
				stopType = "soft"
			}
			createstopaction := &vpcv1.StopBareMetalServerOptions{
				ID:   &bareMetalServerId,
				Type: &stopType,
			}
			res, err := sess.StopBareMetalServerWithContext(context, createstopaction)
			if err != nil || res.StatusCode != 204 {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("StopBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			_, err = isWaitForBareMetalServerStoppedForNIC(sess, bareMetalServerId, d.Timeout(schema.TimeoutCreate), d)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerStoppedForNIC failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		} else if *bms.Status != "stopped" {
			err = fmt.Errorf("[ERROR] Error bare metal server in %s state, please try after some time", *bms.Status)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		nicOptions := &vpcv1.BareMetalServerNetworkInterfacePrototypeBareMetalServerNetworkInterfaceByPciPrototype{}
		allowedVlansList := allowedVlansOk.(*schema.Set).List()

		allowedVlans := make([]int64, 0, len(allowedVlansList))
		for _, k := range allowedVlansList {
			allowedVlans = append(allowedVlans, int64(k.(int)))
		}
		nicOptions.AllowedVlans = allowedVlans

		if name, ok := d.GetOk(isBareMetalServerNicName); ok {
			nameStr := name.(string)
			nicOptions.Name = &nameStr
		}
		nicOptions.InterfaceType = &interfaceType

		if ais, ok := d.GetOk(isBareMetalServerNicAllowIPSpoofing); ok {
			allowIPSpoofing := ais.(bool)
			if allowIPSpoofing {
				nicOptions.AllowIPSpoofing = &allowIPSpoofing
			}
		}
		if ein, ok := d.GetOk(isBareMetalServerNicEnableInfraNAT); ok {
			enableInfrastructureNat := ein.(bool)
			nicOptions.EnableInfrastructureNat = &enableInfrastructureNat
		}
		if subnetOk, ok := d.GetOk(isBareMetalServerNicSubnet); ok {
			subnet := subnetOk.(string)
			nicOptions.Subnet = &vpcv1.SubnetIdentity{
				ID: &subnet,
			}
		}

		if primaryIpIntf, ok := d.GetOk(isBareMetalServerNicPrimaryIP); ok && len(primaryIpIntf.([]interface{})) > 0 {
			primaryIp := primaryIpIntf.([]interface{})[0].(map[string]interface{})

			reservedIpIdOk, ok := primaryIp[isBareMetalServerNicIpID]
			if ok && reservedIpIdOk.(string) != "" {
				ipid := reservedIpIdOk.(string)
				nicOptions.PrimaryIP = &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
					ID: &ipid,
				}
			} else {
				primaryip := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{}

				reservedIpAddressOk, okAdd := primaryIp[isBareMetalServerNicIpAddress]
				if okAdd && reservedIpAddressOk.(string) != "" {
					reservedIpAddress := reservedIpAddressOk.(string)
					primaryip.Address = &reservedIpAddress
				}
				reservedIpNameOk, okName := primaryIp[isBareMetalServerNicIpName]
				if okName && reservedIpNameOk.(string) != "" {
					reservedIpName := reservedIpNameOk.(string)
					primaryip.Name = &reservedIpName
				}
				reservedIpAutoOk, okAuto := primaryIp[isBareMetalServerNicIpAutoDelete]
				if okAuto {
					reservedIpAuto := reservedIpAutoOk.(bool)
					primaryip.AutoDelete = &reservedIpAuto
				}
				if okAdd || okName || okAuto {
					nicOptions.PrimaryIP = primaryip
				}
			}

		}

		sGroups := d.Get(isBareMetalServerNicSecurityGroups).(*schema.Set).List()
		var sGroupList []vpcv1.SecurityGroupIdentityIntf
		// Add new allowed_subnets
		for _, sGroup := range sGroups {
			sGroupStr := sGroup.(string)
			sgModel := &vpcv1.SecurityGroupIdentity{
				ID: &sGroupStr,
			}
			sGroupList = append(sGroupList, sgModel)
		}
		nicOptions.SecurityGroups = sGroupList
		options.BareMetalServerID = &bareMetalServerId
		options.BareMetalServerNetworkInterfacePrototype = nicOptions
		nic, _, err := sess.CreateBareMetalServerNetworkInterfaceWithContext(context, options)
		if err != nil || nic == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		switch reflect.TypeOf(nic).String() {
		case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
			{
				nicIntf := nic.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
				d.SetId(MakeTerraformNICID(bareMetalServerId, *nicIntf.ID))
			}
		case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
			{
				nicIntf := nic.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
				d.SetId(MakeTerraformNICID(bareMetalServerId, *nicIntf.ID))
			}
		}

		_, nicId, err := ParseNICTerraformID(d.Id())
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "create", "sep-id-parts").GetDiag()
		}

		log.Printf("[INFO] Bare Metal Server Network Interface : %s", d.Id())
		nicAfterWait, err := isWaitForBareMetalServerNetworkInterfaceAvailable(sess, bareMetalServerId, nicId, d.Timeout(schema.TimeoutCreate), d)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerNetworkInterfaceAvailable failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		diagErr := bareMetalServerNICGet(context, d, meta, sess, nicAfterWait, bareMetalServerId)
		if diagErr != nil {
			return diagErr
		}

		// restarting the server after PCI creation
		log.Printf("[DEBUG] Starting bare metal server (%s) to create a PCI network interface", bareMetalServerId)

		createstartaction := &vpcv1.StartBareMetalServerOptions{
			ID: &bareMetalServerId,
		}
		res, err := sess.StartBareMetalServerWithContext(context, createstartaction)
		if err != nil || res.StatusCode != 204 {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("StartBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForBareMetalServerAvailableForNIC(sess, bareMetalServerId, d.Timeout(schema.TimeoutCreate), d)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerAvailableForNIC failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

	} else if interfaceTypeOk, ok := d.GetOk(isBareMetalServerNicInterfaceType); ok {
		interfaceType := interfaceTypeOk.(string)
		diagErr := createHiperSocketTypeNetworkInterface(context, d, meta, interfaceType, bareMetalServerId)
		if diagErr != nil {
			return diagErr
		}
	} else {
		diagErr := createVlanTypeNetworkInterface(context, d, meta, bareMetalServerId)
		if diagErr != nil {
			return diagErr
		}
	}
	return nil
}

func createVlanTypeNetworkInterface(context context.Context, d *schema.ResourceData, meta interface{}, bareMetalServerId string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.CreateBareMetalServerNetworkInterfaceOptions{}
	interfaceType := "vlan"
	nicOptions := &vpcv1.BareMetalServerNetworkInterfacePrototypeBareMetalServerNetworkInterfaceByVlanPrototype{}
	if aitf, ok := d.GetOk(isBareMetalServerNicAllowInterfaceToFloat); ok {
		allowInterfaceToFloat := aitf.(bool)
		nicOptions.AllowInterfaceToFloat = &allowInterfaceToFloat
	}
	if vlan, ok := d.GetOk(isBareMetalServerNicVlan); ok {
		vlanInt := int64(vlan.(int))
		nicOptions.Vlan = &vlanInt
	}

	if name, ok := d.GetOk(isBareMetalServerNicName); ok {
		nameStr := name.(string)
		nicOptions.Name = &nameStr
	}
	nicOptions.InterfaceType = &interfaceType

	if ais, ok := d.GetOk(isBareMetalServerNicAllowIPSpoofing); ok {
		allowIPSpoofing := ais.(bool)
		if allowIPSpoofing {
			nicOptions.AllowIPSpoofing = &allowIPSpoofing
		}
	}
	if ein, ok := d.GetOk(isBareMetalServerNicEnableInfraNAT); ok {
		enableInfrastructureNat := ein.(bool)
		nicOptions.EnableInfrastructureNat = &enableInfrastructureNat
	}
	if subnetOk, ok := d.GetOk(isBareMetalServerNicSubnet); ok {
		subnet := subnetOk.(string)
		nicOptions.Subnet = &vpcv1.SubnetIdentity{
			ID: &subnet,
		}
	}

	if primaryIpIntf, ok := d.GetOk(isBareMetalServerNicPrimaryIP); ok && len(primaryIpIntf.([]interface{})) > 0 {
		primaryIp := primaryIpIntf.([]interface{})[0].(map[string]interface{})

		reservedIpIdOk, ok := primaryIp[isBareMetalServerNicIpID]
		if ok && reservedIpIdOk.(string) != "" {
			ipid := reservedIpIdOk.(string)
			nicOptions.PrimaryIP = &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
				ID: &ipid,
			}
		} else {
			primaryip := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{}
			reservedIpAddressOk, okAdd := primaryIp[isBareMetalServerNicIpAddress]
			if okAdd && reservedIpAddressOk.(string) != "" {
				reservedIpAddress := reservedIpAddressOk.(string)
				primaryip.Address = &reservedIpAddress
			}

			reservedIpNameOk, okName := primaryIp[isBareMetalServerNicIpName]
			if okName && reservedIpNameOk.(string) != "" {
				reservedIpName := reservedIpNameOk.(string)
				primaryip.Name = &reservedIpName
			}
			reservedIpAutoOk, okAuto := primaryIp[isBareMetalServerNicIpAutoDelete]
			if okAuto {
				reservedIpAuto := reservedIpAutoOk.(bool)
				primaryip.AutoDelete = &reservedIpAuto
			}
			if okAdd || okName || okAuto {
				nicOptions.PrimaryIP = primaryip
			}
		}

	}

	sGroups := d.Get(isBareMetalServerNicSecurityGroups).(*schema.Set).List()
	var sGroupList []vpcv1.SecurityGroupIdentityIntf
	// Add new allowed_subnets
	for _, sGroup := range sGroups {
		sGroupStr := sGroup.(string)
		sgModel := &vpcv1.SecurityGroupIdentity{
			ID: &sGroupStr,
		}
		sGroupList = append(sGroupList, sgModel)
	}
	nicOptions.SecurityGroups = sGroupList
	options.BareMetalServerID = &bareMetalServerId
	options.BareMetalServerNetworkInterfacePrototype = nicOptions
	nic, _, err := sess.CreateBareMetalServerNetworkInterfaceWithContext(context, options)
	if err != nil || nic == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	switch reflect.TypeOf(nic).String() {
	case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
		{
			nicIntf := nic.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
			d.SetId(MakeTerraformNICID(bareMetalServerId, *nicIntf.ID))
		}
	case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
		{
			nicIntf := nic.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
			d.SetId(MakeTerraformNICID(bareMetalServerId, *nicIntf.ID))
		}
	}

	_, nicId, err := ParseNICTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "create", "sep-id-parts").GetDiag()
	}
	log.Printf("[INFO] Bare Metal Server Network Interface : %s", d.Id())
	nicAfterWait, err := isWaitForBareMetalServerNetworkInterfaceAvailable(sess, bareMetalServerId, nicId, d.Timeout(schema.TimeoutCreate), d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerNetworkInterfaceAvailable failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	diagErr := bareMetalServerNICGet(context, d, meta, sess, nicAfterWait, bareMetalServerId)
	if diagErr != nil {
		return diagErr
	}
	return nil
}

func createHiperSocketTypeNetworkInterface(context context.Context, d *schema.ResourceData, meta interface{}, interfaceType, bareMetalServerId string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.CreateBareMetalServerNetworkInterfaceOptions{}
	nicOptions := &vpcv1.BareMetalServerNetworkInterfacePrototypeBareMetalServerNetworkInterfaceByHiperSocketPrototype{}

	if name, ok := d.GetOk(isBareMetalServerNicName); ok {
		nameStr := name.(string)
		nicOptions.Name = &nameStr
	}
	nicOptions.InterfaceType = &interfaceType

	if ais, ok := d.GetOk(isBareMetalServerNicAllowIPSpoofing); ok {
		allowIPSpoofing := ais.(bool)
		if allowIPSpoofing {
			nicOptions.AllowIPSpoofing = &allowIPSpoofing
		}
	}
	if ein, ok := d.GetOk(isBareMetalServerNicEnableInfraNAT); ok {
		enableInfrastructureNat := ein.(bool)
		nicOptions.EnableInfrastructureNat = &enableInfrastructureNat
	}
	if subnetOk, ok := d.GetOk(isBareMetalServerNicSubnet); ok {
		subnet := subnetOk.(string)
		nicOptions.Subnet = &vpcv1.SubnetIdentity{
			ID: &subnet,
		}
	}

	if primaryIpIntf, ok := d.GetOk(isBareMetalServerNicPrimaryIP); ok && len(primaryIpIntf.([]interface{})) > 0 {
		primaryIp := primaryIpIntf.([]interface{})[0].(map[string]interface{})

		reservedIpIdOk, ok := primaryIp[isBareMetalServerNicIpID]
		if ok && reservedIpIdOk.(string) != "" {
			ipid := reservedIpIdOk.(string)
			nicOptions.PrimaryIP = &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
				ID: &ipid,
			}
		} else {
			primaryip := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{}
			reservedIpAddressOk, okAdd := primaryIp[isBareMetalServerNicIpAddress]
			if okAdd && reservedIpAddressOk.(string) != "" {
				reservedIpAddress := reservedIpAddressOk.(string)
				primaryip.Address = &reservedIpAddress
			}
			reservedIpNameOk, okName := primaryIp[isBareMetalServerNicIpName]
			if okName && reservedIpNameOk.(string) != "" {
				reservedIpName := reservedIpNameOk.(string)
				primaryip.Name = &reservedIpName
			}
			reservedIpAutoOk, okAuto := primaryIp[isBareMetalServerNicIpAutoDelete]
			if okAuto {
				reservedIpAuto := reservedIpAutoOk.(bool)
				primaryip.AutoDelete = &reservedIpAuto
			}
			if okAdd || okName || okAuto {
				nicOptions.PrimaryIP = primaryip
			}

			// reservedIpAddressOk, ok := primaryIp[isBareMetalServerNicIpAddress]
			// if ok && reservedIpAddressOk.(string) != "" {
			// 	reservedIpAddress := reservedIpAddressOk.(string)
			// 	nicOptions.PrimaryIpv4Address = &reservedIpAddress
			// }
		}

	}

	sGroups := d.Get(isBareMetalServerNicSecurityGroups).(*schema.Set).List()
	var sGroupList []vpcv1.SecurityGroupIdentityIntf
	// Add new allowed_subnets
	for _, sGroup := range sGroups {
		sGroupStr := sGroup.(string)
		sgModel := &vpcv1.SecurityGroupIdentity{
			ID: &sGroupStr,
		}
		sGroupList = append(sGroupList, sgModel)
	}
	nicOptions.SecurityGroups = sGroupList
	options.BareMetalServerID = &bareMetalServerId
	options.BareMetalServerNetworkInterfacePrototype = nicOptions
	nic, _, err := sess.CreateBareMetalServerNetworkInterfaceWithContext(context, options)
	if err != nil || nic == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	diagErr := bareMetalServerNICGet(context, d, meta, sess, nic, bareMetalServerId)
	if diagErr != nil {
		return diagErr
	}
	_, nicId, err := ParseNICTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "create", "sep-id-parts").GetDiag()
	}
	log.Printf("[INFO] Bare Metal Server Network Interface : %s", d.Id())
	_, err = isWaitForBareMetalServerNetworkInterfaceAvailable(sess, bareMetalServerId, nicId, d.Timeout(schema.TimeoutCreate), d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerNetworkInterfaceAvailable failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	return nil
}

func resourceIBMISBareMetalServerNetworkInterfaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	bareMetalServerId, nicID, err := ParseNICTerraformID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	sess, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	options := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
		BareMetalServerID: &bareMetalServerId,
		ID:                &nicID,
	}

	nicIntf, response, err := sess.GetBareMetalServerNetworkInterfaceWithContext(context, options)
	if err != nil || nicIntf == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting Bare Metal Server (%s) network interface (%s): %s\n%s", bareMetalServerId, nicID, err, response))
	}
	diagErr := bareMetalServerNICGet(context, d, meta, sess, nicIntf, bareMetalServerId)
	if diagErr != nil {
		return diagErr
	}
	return nil
}

func bareMetalServerNICGet(context context.Context, d *schema.ResourceData, meta interface{}, sess *vpcv1.VpcV1, nicIntf interface{}, bareMetalServerId string) diag.Diagnostics {
	var err error
	switch reflect.TypeOf(nicIntf).String() {
	case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
		{
			nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
			if err = d.Set("allow_ip_spoofing", nic.AllowIPSpoofing); err != nil {
				err = fmt.Errorf("Error setting allow_ip_spoofing: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-allow_ip_spoofing").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicID, *nic.ID); err != nil {
				err = fmt.Errorf("Error setting network_interface: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-network_interface").GetDiag()
			}
			if err = d.Set("enable_infrastructure_nat", nic.EnableInfrastructureNat); err != nil {
				err = fmt.Errorf("Error setting enable_infrastructure_nat: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-enable_infrastructure_nat").GetDiag()
			}
			if err = d.Set("status", nic.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-status").GetDiag()
			}
			floatingIPList := make([]map[string]interface{}, 0)
			if nic.FloatingIps != nil {
				for _, ip := range nic.FloatingIps {
					currentIP := map[string]interface{}{
						isBareMetalServerNicFloatingIPId: *ip.ID,
						isBareMetalServerNicIpAddress:    *ip.Address,
					}
					floatingIPList = append(floatingIPList, currentIP)
				}
			}
			if err = d.Set(isBareMetalServerNicFloatingIPs, floatingIPList); err != nil {
				err = fmt.Errorf("Error setting floating_ips: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-floating_ips").GetDiag()
			}
			if err = d.Set("href", nic.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-href").GetDiag()
			}
			if err = d.Set("interface_type", nic.InterfaceType); err != nil {
				err = fmt.Errorf("Error setting interface_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-interface_type").GetDiag()
			}
			if err = d.Set("mac_address", nic.MacAddress); err != nil {
				err = fmt.Errorf("Error setting mac_address: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-mac_address").GetDiag()
			}
			if err = d.Set("name", nic.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-name").GetDiag()
			}
			if nic.PortSpeed != nil {
				if err = d.Set("port_speed", flex.IntValue(nic.PortSpeed)); err != nil {
					err = fmt.Errorf("Error setting port_speed: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-port_speed").GetDiag()
				}
			}
			if nic.PrimaryIP != nil {
				primaryIpList := make([]map[string]interface{}, 0)
				currentIP := map[string]interface{}{}
				if nic.PrimaryIP.Href != nil {
					currentIP[isBareMetalServerNicIpAddress] = *nic.PrimaryIP.Address
				}
				if nic.PrimaryIP.Href != nil {
					currentIP[isBareMetalServerNicIpHref] = *nic.PrimaryIP.Href
				}
				if nic.PrimaryIP.Name != nil {
					currentIP[isBareMetalServerNicIpName] = *nic.PrimaryIP.Name
				}
				if nic.PrimaryIP.ID != nil {
					currentIP[isBareMetalServerNicIpID] = *nic.PrimaryIP.ID
				}
				if nic.PrimaryIP.ResourceType != nil {
					currentIP[isBareMetalServerNicResourceType] = *nic.PrimaryIP.ResourceType
				}
				getripoptions := &vpcv1.GetSubnetReservedIPOptions{
					SubnetID: nic.Subnet.ID,
					ID:       nic.PrimaryIP.ID,
				}
				bmsRip, _, err := sess.GetSubnetReservedIPWithContext(context, getripoptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "read")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				currentIP[isBareMetalServerNicIpAutoDelete] = bmsRip.AutoDelete
				primaryIpList = append(primaryIpList, currentIP)
				if err = d.Set("primary_ip", primaryIpList); err != nil {
					err = fmt.Errorf("Error setting primary_ip: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-primary_ip").GetDiag()
				}
			}
			if err = d.Set("resource_type", nic.ResourceType); err != nil {
				err = fmt.Errorf("Error setting resource_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-resource_type").GetDiag()
			}
			if nic.SecurityGroups != nil && len(nic.SecurityGroups) != 0 {
				secgrpList := []string{}
				for i := 0; i < len(nic.SecurityGroups); i++ {
					secgrpList = append(secgrpList, string(*(nic.SecurityGroups[i].ID)))
				}

				if err = d.Set(isBareMetalServerNicSecurityGroups, flex.NewStringSet(schema.HashString, secgrpList)); err != nil {
					err = fmt.Errorf("Error setting security_groups: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-security_groups").GetDiag()
				}
			}

			if err = d.Set("status", nic.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-status").GetDiag()
			}
			if err = d.Set("subnet", *nic.Subnet.ID); err != nil {
				err = fmt.Errorf("Error setting subnet: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-subnet").GetDiag()
			}
			if err = d.Set("type", nic.Type); err != nil {
				err = fmt.Errorf("Error setting type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-type").GetDiag()
			}
			if nic.AllowedVlans != nil {
				var out = make([]interface{}, len(nic.AllowedVlans), len(nic.AllowedVlans))
				for i, v := range nic.AllowedVlans {
					out[i] = int(v)
				}
				if err = d.Set("allowed_vlans", schema.NewSet(schema.HashInt, out)); err != nil {
					err = fmt.Errorf("Error setting allowed_vlans: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-allowed_vlans").GetDiag()
				}
			}
		}
	case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
		{
			nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
			d.Set(isBareMetalServerNicAllowIPSpoofing, *nic.AllowIPSpoofing)
			if err = d.Set("allow_ip_spoofing", nic.AllowIPSpoofing); err != nil {
				err = fmt.Errorf("Error setting allow_ip_spoofing: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-allow_ip_spoofing").GetDiag()
			}
			if err = d.Set("enable_infrastructure_nat", nic.EnableInfrastructureNat); err != nil {
				err = fmt.Errorf("Error setting enable_infrastructure_nat: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-enable_infrastructure_nat").GetDiag()
			}
			if err = d.Set("status", nic.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-status").GetDiag()
			}
			floatingIPList := make([]map[string]interface{}, 0)
			if nic.FloatingIps != nil {
				for _, ip := range nic.FloatingIps {
					currentIP := map[string]interface{}{
						isBareMetalServerNicFloatingIPId: *ip.ID,
						isBareMetalServerNicIpAddress:    *ip.Address,
					}
					floatingIPList = append(floatingIPList, currentIP)
				}
			}
			if err = d.Set(isBareMetalServerNicFloatingIPs, floatingIPList); err != nil {
				err = fmt.Errorf("Error setting floating_ips: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-floating_ips").GetDiag()
			}

			if err = d.Set("href", nic.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-href").GetDiag()
			}

			if err = d.Set(isBareMetalServerNicID, *nic.ID); err != nil {
				err = fmt.Errorf("Error setting network_interface: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-network_interface").GetDiag()
			}
			if err = d.Set("interface_type", nic.InterfaceType); err != nil {
				err = fmt.Errorf("Error setting interface_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-interface_type").GetDiag()
			}

			if err = d.Set("mac_address", nic.MacAddress); err != nil {
				err = fmt.Errorf("Error setting mac_address: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-mac_address").GetDiag()
			}
			if err = d.Set("name", nic.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-name").GetDiag()
			}
			if err = d.Set("port_speed", flex.IntValue(nic.PortSpeed)); err != nil {
				err = fmt.Errorf("Error setting port_speed: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-port_speed").GetDiag()
			}

			primaryIpList := make([]map[string]interface{}, 0)
			currentIP := map[string]interface{}{}
			if nic.PrimaryIP.Href != nil {
				currentIP[isBareMetalServerNicIpAddress] = *nic.PrimaryIP.Address
			}
			if nic.PrimaryIP.Href != nil {
				currentIP[isBareMetalServerNicIpHref] = *nic.PrimaryIP.Href
			}
			if nic.PrimaryIP.Name != nil {
				currentIP[isBareMetalServerNicIpName] = *nic.PrimaryIP.Name
			}
			if nic.PrimaryIP.ID != nil {
				currentIP[isBareMetalServerNicIpID] = *nic.PrimaryIP.ID
			}
			if nic.PrimaryIP.ResourceType != nil {
				currentIP[isBareMetalServerNicResourceType] = *nic.PrimaryIP.ResourceType
			}
			getripoptions := &vpcv1.GetSubnetReservedIPOptions{
				SubnetID: nic.Subnet.ID,
				ID:       nic.PrimaryIP.ID,
			}
			bmsRip, _, err := sess.GetSubnetReservedIPWithContext(context, getripoptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			currentIP[isBareMetalServerNicIpAutoDelete] = bmsRip.AutoDelete

			primaryIpList = append(primaryIpList, currentIP)
			if err = d.Set("primary_ip", primaryIpList); err != nil {
				err = fmt.Errorf("Error setting primary_ip: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-primary_ip").GetDiag()
			}

			if err = d.Set("resource_type", nic.ResourceType); err != nil {
				err = fmt.Errorf("Error setting resource_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-resource_type").GetDiag()
			}

			if len(nic.SecurityGroups) != 0 {
				secgrpList := []string{}
				for i := 0; i < len(nic.SecurityGroups); i++ {
					secgrpList = append(secgrpList, string(*(nic.SecurityGroups[i].ID)))
				}
				if err = d.Set(isBareMetalServerNicSecurityGroups, flex.NewStringSet(schema.HashString, secgrpList)); err != nil {
					err = fmt.Errorf("Error setting security_groups: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-security_groups").GetDiag()
				}
			}

			if err = d.Set("status", nic.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-status").GetDiag()
			}
			if err = d.Set("subnet", *nic.Subnet.ID); err != nil {
				err = fmt.Errorf("Error setting subnet: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-subnet").GetDiag()
			}
			if err = d.Set("type", nic.Type); err != nil {
				err = fmt.Errorf("Error setting type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-type").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicAllowInterfaceToFloat, *nic.AllowInterfaceToFloat); err != nil {
				err = fmt.Errorf("Error setting allow_interface_to_float: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-allow_interface_to_float").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicVlan, *nic.Vlan); err != nil {
				err = fmt.Errorf("Error setting vlan: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-vlan").GetDiag()
			}
		}
	case "*vpcv1.BareMetalServerNetworkInterfaceByHiperSocket":
		{
			nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByHiperSocket)
			d.SetId(MakeTerraformNICID(bareMetalServerId, *nic.ID))
			if err = d.Set("allow_ip_spoofing", nic.AllowIPSpoofing); err != nil {
				err = fmt.Errorf("Error setting allow_ip_spoofing: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-allow_ip_spoofing").GetDiag()
			}
			if err = d.Set("enable_infrastructure_nat", nic.EnableInfrastructureNat); err != nil {
				err = fmt.Errorf("Error setting enable_infrastructure_nat: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-enable_infrastructure_nat").GetDiag()
			}
			if err = d.Set("status", nic.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-status").GetDiag()
			}
			floatingIPList := make([]map[string]interface{}, 0)
			if nic.FloatingIps != nil {
				for _, ip := range nic.FloatingIps {
					currentIP := map[string]interface{}{
						isBareMetalServerNicIpID:      *ip.ID,
						isBareMetalServerNicIpAddress: *ip.Address,
					}
					floatingIPList = append(floatingIPList, currentIP)
				}
			}
			if err = d.Set(isBareMetalServerNicFloatingIPs, floatingIPList); err != nil {
				err = fmt.Errorf("Error setting floating_ips: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-floating_ips").GetDiag()
			}

			if err = d.Set("href", nic.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-href").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicID, *nic.ID); err != nil {
				err = fmt.Errorf("Error setting network_interface: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-network_interface").GetDiag()
			}
			if err = d.Set("interface_type", nic.InterfaceType); err != nil {
				err = fmt.Errorf("Error setting interface_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-interface_type").GetDiag()
			}

			if err = d.Set("mac_address", nic.MacAddress); err != nil {
				err = fmt.Errorf("Error setting mac_address: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-mac_address").GetDiag()
			}
			if err = d.Set("name", nic.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-name").GetDiag()
			}
			if err = d.Set("port_speed", flex.IntValue(nic.PortSpeed)); err != nil {
				err = fmt.Errorf("Error setting port_speed: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-port_speed").GetDiag()
			}

			primaryIpList := make([]map[string]interface{}, 0)
			// currentIP := map[string]interface{}{

			// 	isBareMetalServerNicIpAddress: *nic.PrimaryIpv4Address,
			// }
			currentIP := map[string]interface{}{}
			if nic.PrimaryIP.Href != nil {
				currentIP[isBareMetalServerNicIpAddress] = *nic.PrimaryIP.Address
			}
			if nic.PrimaryIP.Href != nil {
				currentIP[isBareMetalServerNicIpHref] = *nic.PrimaryIP.Href
			}
			if nic.PrimaryIP.Name != nil {
				currentIP[isBareMetalServerNicIpName] = *nic.PrimaryIP.Name
			}
			if nic.PrimaryIP.ID != nil {
				currentIP[isBareMetalServerNicIpID] = *nic.PrimaryIP.ID
			}
			if nic.PrimaryIP.ResourceType != nil {
				currentIP[isBareMetalServerNicResourceType] = *nic.PrimaryIP.ResourceType
			}
			getripoptions := &vpcv1.GetSubnetReservedIPOptions{
				SubnetID: nic.Subnet.ID,
				ID:       nic.PrimaryIP.ID,
			}
			bmsRip, _, err := sess.GetSubnetReservedIPWithContext(context, getripoptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			currentIP[isBareMetalServerNicIpAutoDelete] = bmsRip.AutoDelete

			primaryIpList = append(primaryIpList, currentIP)
			// primaryIpList = append(primaryIpList, currentIP)
			if err = d.Set("primary_ip", primaryIpList); err != nil {
				err = fmt.Errorf("Error setting primary_ip: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-primary_ip").GetDiag()
			}

			if err = d.Set("resource_type", nic.ResourceType); err != nil {
				err = fmt.Errorf("Error setting resource_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-resource_type").GetDiag()
			}
			if len(nic.SecurityGroups) != 0 {
				secgrpList := []string{}
				for i := 0; i < len(nic.SecurityGroups); i++ {
					secgrpList = append(secgrpList, string(*(nic.SecurityGroups[i].ID)))
				}
				if err = d.Set(isBareMetalServerNicSecurityGroups, flex.NewStringSet(schema.HashString, secgrpList)); err != nil {
					err = fmt.Errorf("Error setting security_groups: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-security_groups").GetDiag()
				}
			}

			if err = d.Set("status", nic.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-status").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicSubnet, *nic.Subnet.ID); err != nil {
				err = fmt.Errorf("Error setting subnet: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-subnet").GetDiag()
			}
			if err = d.Set("type", nic.Type); err != nil {
				err = fmt.Errorf("Error setting type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "read", "set-type").GetDiag()
			}
		}
	}
	return nil
}

func resourceIBMISBareMetalServerNetworkInterfaceUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	bareMetalServerId, nicId, err := ParseNICTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "update", "sep-id-parts").GetDiag()
	}
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if d.HasChange("security_groups") && !d.IsNewResource() {
		ovs, nvs := d.GetChange("security_groups")
		ov := ovs.(*schema.Set)
		nv := nvs.(*schema.Set)
		remove := flex.ExpandStringList(ov.Difference(nv).List())
		add := flex.ExpandStringList(nv.Difference(ov).List())
		if len(add) > 0 {
			for i := range add {
				createsgnicoptions := &vpcv1.CreateSecurityGroupTargetBindingOptions{
					SecurityGroupID: &add[i],
					ID:              &nicId,
				}
				_, _, err := sess.CreateSecurityGroupTargetBindingWithContext(context, createsgnicoptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSecurityGroupTargetBindingWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForBareMetalServerAvailableForNIC(sess, bareMetalServerId, d.Timeout(schema.TimeoutUpdate), d)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerAvailableForNIC failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}

		}
		if len(remove) > 0 {
			for i := range remove {
				deletesgnicoptions := &vpcv1.DeleteSecurityGroupTargetBindingOptions{
					SecurityGroupID: &remove[i],
					ID:              &nicId,
				}
				_, err := sess.DeleteSecurityGroupTargetBindingWithContext(context, deletesgnicoptions)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSecurityGroupTargetBindingWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForBareMetalServerAvailableForNIC(sess, bareMetalServerId, d.Timeout(schema.TimeoutUpdate), d)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerAvailableForNIC failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}
	}

	options := &vpcv1.UpdateBareMetalServerNetworkInterfaceOptions{
		BareMetalServerID: &bareMetalServerId,
		ID:                &nicId,
	}

	if d.HasChange("primary_ip.0.name") || d.HasChange("primary_ip.0.auto_delete") {
		subnetId := d.Get(isBareMetalServerNicSubnet).(string)
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
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("reservedIpPath.AsPatch() failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateripoptions.ReservedIPPatch = reservedIpPathAsPatch
		_, _, err = sess.UpdateSubnetReservedIPWithContext(context, updateripoptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	nicPatchModel := &vpcv1.BareMetalServerNetworkInterfacePatch{}
	flag := false
	if d.HasChange(isBareMetalServerNicAllowIPSpoofing) {
		flag = true
		aisBool := false
		if ais, ok := d.GetOk(isBareMetalServerNicAllowIPSpoofing); ok {
			aisBool = ais.(bool)
		}
		nicPatchModel.AllowIPSpoofing = &aisBool
	}
	if d.HasChange(isBareMetalServerNicAllowedVlans) {
		flag = true

		if allowedVlansOk, ok := d.GetOk(isBareMetalServerNicAllowedVlans); ok {
			allowedVlansList := allowedVlansOk.(*schema.Set).List()
			allowedVlans := make([]int64, 0, len(allowedVlansList))
			for _, k := range allowedVlansList {
				allowedVlans = append(allowedVlans, int64(k.(int)))
			}
			nicPatchModel.AllowedVlans = allowedVlans
		}
	}
	if d.HasChange(isBareMetalServerNicEnableInfraNAT) {
		flag = true
		einBool := false
		if ein, ok := d.GetOk(isBareMetalServerNicEnableInfraNAT); ok {
			einBool = ein.(bool)
		}
		nicPatchModel.EnableInfrastructureNat = &einBool
	}
	if d.HasChange(isBareMetalServerNicName) {
		flag = true
		nameStr := ""
		if name, ok := d.GetOk(isBareMetalServerNicName); ok {
			nameStr = name.(string)
		}
		nicPatchModel.Name = &nameStr
	}

	if flag {
		nicPatchModelAsPatch, err := nicPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("nicPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options.BareMetalServerNetworkInterfacePatch = nicPatchModelAsPatch

		nicIntf, _, err := sess.UpdateBareMetalServerNetworkInterfaceWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		return bareMetalServerNICGet(context, d, meta, sess, nicIntf, bareMetalServerId)
	}

	return nil
}

func resourceIBMISBareMetalServerNetworkInterfaceDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	bareMetalServerId, nicId, err := ParseNICTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "delete", "sep-id-parts").GetDiag()
	}

	diagErr := bareMetalServerNetworkInterfaceDelete(context, d, meta, bareMetalServerId, nicId)
	if diagErr != nil {
		return diagErr
	}

	return nil
}

func bareMetalServerNetworkInterfaceDelete(context context.Context, d *schema.ResourceData, meta interface{}, bareMetalServerId, nicId string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getBmsNicOptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
		BareMetalServerID: &bareMetalServerId,
		ID:                &nicId,
	}
	nicIntf, response, err := sess.GetBareMetalServerNetworkInterfaceWithContext(context, getBmsNicOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	nicType := ""
	switch reflect.TypeOf(nicIntf).String() {
	case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
		{
			nicType = "pci"
			log.Printf("[DEBUG] PCI type network interface needs the server in stopped state")
			log.Printf("[DEBUG] Stopping the bare metal server %s", bareMetalServerId)
			// to delete pci, server needs to be in stopped state

			getbmsoptions := &vpcv1.GetBareMetalServerOptions{
				ID: &bareMetalServerId,
			}

			bms, _, err := sess.GetBareMetalServerWithContext(context, getbmsoptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "delete")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			// failed, pending, restarting, running, starting, stopped, stopping, maintenance
			if *bms.Status == "failed" {
				err := fmt.Errorf("Error cannot detach network interface from a failed bare metal server")
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "delete")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			} else if *bms.Status == "running" {
				log.Printf("[DEBUG] Stopping bare metal server (%s) to create a PCI network interface", bareMetalServerId)
				stopType := "soft"
				if d.Get(isBareMetalServerHardStop).(bool) {
					stopType = "hard"
				}
				createstopaction := &vpcv1.StopBareMetalServerOptions{
					ID:   &bareMetalServerId,
					Type: &stopType,
				}
				res, err := sess.StopBareMetalServerWithContext(context, createstopaction)
				if err != nil || res.StatusCode != 204 {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("StopBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "delete")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForBareMetalServerStoppedForNIC(sess, bareMetalServerId, d.Timeout(schema.TimeoutDelete), d)
				if err != nil || res.StatusCode != 204 {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerStoppedForNIC failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "delete")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			} else if *bms.Status != "stopped" {
				err = fmt.Errorf("Error bare metal server in %s state, please try after some time", *bms.Status)
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "delete")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		}
	case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
		{
			nicType = "vlan"
		}
	}

	options := &vpcv1.DeleteBareMetalServerNetworkInterfaceOptions{
		BareMetalServerID: &bareMetalServerId,
		ID:                &nicId,
	}
	response, err = sess.DeleteBareMetalServerNetworkInterfaceWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForBareMetalServerNetworkInterfaceDeleted(sess, bareMetalServerId, nicId, nicType, nicIntf, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerNetworkInterfaceDeleted failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if nicType == "pci" {
		// restarting the server after PCI deletion
		log.Printf("[DEBUG] Starting bare metal server (%s) after deleting the PCI network interface", bareMetalServerId)
		createstartaction := &vpcv1.StartBareMetalServerOptions{
			ID: &bareMetalServerId,
		}
		res, err := sess.StartBareMetalServerWithContext(context, createstartaction)
		if err != nil || res.StatusCode != 204 {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("StartBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		_, err = isWaitForBareMetalServerAvailableForNIC(sess, bareMetalServerId, d.Timeout(schema.TimeoutCreate), d)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerAvailableForNIC failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	d.SetId("")
	return nil
}

func isWaitForBareMetalServerNetworkInterfaceDeleted(bmsC *vpcv1.VpcV1, bareMetalServerId, nicId, nicType string, nicIntf vpcv1.BareMetalServerNetworkInterfaceIntf, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for (%s) / (%s) to be deleted.", bareMetalServerId, nicId)
	stateConf := &retry.StateChangeConf{
		Pending:    []string{isBareMetalServerNetworkInterfaceAvailable, isBareMetalServerNetworkInterfaceDeleting, isBareMetalServerNetworkInterfacePending},
		Target:     []string{isBareMetalServerNetworkInterfaceDeleted, isBareMetalServerNetworkInterfaceVlanPending, isBareMetalServerNetworkInterfaceFailed, isBareMetalServerNetworkInterfacePCIPending, ""},
		Refresh:    isBareMetalServerNetworkInterfaceDeleteRefreshFunc(bmsC, bareMetalServerId, nicId, nicType, nicIntf),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isBareMetalServerNetworkInterfaceDeleteRefreshFunc(bmsC *vpcv1.VpcV1, bareMetalServerId, nicId, nicType string, nicIntf vpcv1.BareMetalServerNetworkInterfaceIntf) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getBmsNicOptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
			BareMetalServerID: &bareMetalServerId,
			ID:                &nicId,
		}
		bmsNic, response, err := bmsC.GetBareMetalServerNetworkInterface(getBmsNicOptions)
		if bmsNic != nil && nicType == "vlan" {
			getBmsOptions := &vpcv1.GetBareMetalServerOptions{
				ID: &bareMetalServerId,
			}
			bms, response, err := bmsC.GetBareMetalServer(getBmsOptions)
			if err != nil {
				return bmsNic, isBareMetalServerNetworkInterfaceFailed, fmt.Errorf("[ERROR] Error getting Bare Metal Server(%s) : %s\n%s", bareMetalServerId, err, response)
			}
			if *bms.Status == "stopped" {
				return bmsNic, isBareMetalServerNetworkInterfaceVlanPending, fmt.Errorf("[ERROR] Error deleting Bare Metal Server(%s) Network Interface (%s), server in stopped state ", bareMetalServerId, nicId)
			}
		}
		if bmsNic != nil && nicType == "pci" {
			getBmsOptions := &vpcv1.GetBareMetalServerOptions{
				ID: &bareMetalServerId,
			}
			bms, response, err := bmsC.GetBareMetalServer(getBmsOptions)
			if err != nil {
				return bmsNic, isBareMetalServerNetworkInterfaceFailed, fmt.Errorf("[ERROR] Error getting Bare Metal Server(%s) : %s\n%s", bareMetalServerId, err, response)
			}
			if *bms.Status == "stopped" {
				return bmsNic, isBareMetalServerNetworkInterfacePCIPending, nil
			}
		}
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return nicIntf, isBareMetalServerNetworkInterfaceDeleted, nil
			}
			return bmsNic, isBareMetalServerNetworkInterfaceFailed, fmt.Errorf("[ERROR] Error getting Bare Metal Server(%s) Network Interface (%s): %s\n%s", bareMetalServerId, nicId, err, response)
		}
		return bmsNic, isBareMetalServerNetworkInterfaceDeleting, err
	}
}

func isWaitForBareMetalServerNetworkInterfaceAvailable(client *vpcv1.VpcV1, bareMetalServerId, nicId string, timeout time.Duration, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for Bare Metal Server (%s) Network Interface (%s) to be available.", bareMetalServerId, nicId)
	stateConf := &retry.StateChangeConf{
		Pending:    []string{isBareMetalServerNetworkInterfacePending},
		Target:     []string{isBareMetalServerNetworkInterfaceAvailable, isBareMetalServerNetworkInterfacePCIPending, isBareMetalServerNetworkInterfaceFailed},
		Refresh:    isBareMetalServerNetworkInterfaceRefreshFunc(client, bareMetalServerId, nicId, d),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isBareMetalServerNetworkInterfaceRefreshFunc(client *vpcv1.VpcV1, bareMetalServerId, nicId string, d *schema.ResourceData) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getBmsNicOptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
			BareMetalServerID: &bareMetalServerId,
			ID:                &nicId,
		}
		bmsNic, response, err := client.GetBareMetalServerNetworkInterface(getBmsNicOptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error getting Bare Metal Server (%s) Network Interface (%s) : %s\n%s", bareMetalServerId, nicId, err, response)
		}
		status := ""
		pcipending := false
		switch reflect.TypeOf(bmsNic).String() {
		case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
			{
				nic := bmsNic.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
				status = *nic.Status
				d.Set(isBareMetalServerNicStatus, *nic.Status)
				getBmsOptions := &vpcv1.GetBareMetalServerOptions{
					ID: &bareMetalServerId,
				}
				bms, response, err := client.GetBareMetalServer(getBmsOptions)
				if err != nil {
					return nil, "", fmt.Errorf("[ERROR] Error getting Bare Metal Server (%s)  : %s\n%s", bareMetalServerId, err, response)
				}
				if *bms.Status == "stopped" {
					pcipending = true
				}

			}
		case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
			{
				nic := bmsNic.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
				status = *nic.Status
				d.Set(isBareMetalServerNicStatus, *nic.Status)
				if *nic.PrimaryIP.Address == "0.0.0.0" {
					return bmsNic, "pending", nil
				}
			}
		}

		if status == "available" || status == "failed" {
			return bmsNic, status, nil
		}
		if pcipending {
			return bmsNic, isBareMetalServerNetworkInterfacePCIPending, nil
		}
		return bmsNic, "pending", nil
	}
}

func MakeTerraformNICID(id1, id2 string) string {
	// Include both bare metal sever id and network interface id to create a unique Terraform id.  As a bonus,
	// we can extract the bare metal sever id as needed for API calls such as READ.
	return fmt.Sprintf("%s/%s", id1, id2)
}

func ParseNICTerraformID(s string) (string, string, error) {
	segments := strings.Split(s, "/")
	if len(segments) != 2 {
		return "", "", fmt.Errorf("invalid terraform Id %s (incorrect number of segments)", s)
	}
	if segments[0] == "" || segments[1] == "" {
		return "", "", fmt.Errorf("invalid terraform Id %s (one or more empty segments)", s)
	}
	return segments[0], segments[1], nil
}

func isWaitForBareMetalServerAvailableForNIC(client *vpcv1.VpcV1, id string, timeout time.Duration, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for Bare Metal Server (%s) to be available.", id)
	stateConf := &retry.StateChangeConf{
		Pending:    []string{isBareMetalServerStatusPending, isBareMetalServerActionStatusStarting, "running"},
		Target:     []string{isBareMetalServerStatusRunning, isBareMetalServerStatusFailed},
		Refresh:    isBareMetalServerForNICRefreshFunc(client, id, d),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isBareMetalServerForNICRefreshFunc(client *vpcv1.VpcV1, id string, d *schema.ResourceData) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		bmsgetoptions := &vpcv1.GetBareMetalServerOptions{
			ID: &id,
		}
		bms, response, err := client.GetBareMetalServer(bmsgetoptions)
		if err != nil {
			return nil, "failed", fmt.Errorf("[ERROR] Error getting Bare Metal Server: %s\n%s", err, response)
		}

		if *bms.Status == "running" || *bms.Status == "failed" {
			return bms, *bms.Status, nil
		}
		return bms, "pending", nil
	}
}

func isWaitForBareMetalServerStoppedForNIC(client *vpcv1.VpcV1, id string, timeout time.Duration, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for Bare Metal Server (%s) to be stopped.", id)
	stateConf := &retry.StateChangeConf{
		Pending:    []string{isBareMetalServerStatusPending, isBareMetalServerActionStatusStarting},
		Target:     []string{isBareMetalServerActionStatusStopped},
		Refresh:    isBareMetalServerForNICStoppedRefreshFunc(client, id, d),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isBareMetalServerForNICStoppedRefreshFunc(client *vpcv1.VpcV1, id string, d *schema.ResourceData) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		bmsgetoptions := &vpcv1.GetBareMetalServerOptions{
			ID: &id,
		}
		bms, response, err := client.GetBareMetalServer(bmsgetoptions)
		if err != nil {
			return nil, "failed", fmt.Errorf("[ERROR] Error getting Bare Metal Server: %s\n%s", err, response)
		}
		if *bms.Status == "stopped" || *bms.Status == "failed" {
			// let know the isRestartStartAction() to stop
			if *bms.Status == "failed" {
				return bms, *bms.Status, fmt.Errorf("[ERROR] Error bare metal server in failed state")
			}
			return bms, "stopped", nil

		}
		return bms, isBareMetalServerStatusPending, nil
	}
}
