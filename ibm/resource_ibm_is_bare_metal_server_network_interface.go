// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerNetworkInterfaceAvailable = "available"
	isBareMetalServerNetworkInterfaceDeleting  = "deleting"
	isBareMetalServerNetworkInterfacePending   = "pending"
	isBareMetalServerNetworkInterfaceDeleted   = "deleted"
	isBareMetalServerNetworkInterfaceFailed    = "failed"
)

func resourceIBMisBareMetalServerNetworkInterface() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISBareMetalServerNetworkInterfaceCreate,
		Read:     resourceIBMISBareMetalServerNetworkInterfaceRead,
		Update:   resourceIBMISBareMetalServerNetworkInterfaceUpdate,
		Delete:   resourceIBMISBareMetalServerNetworkInterfaceDelete,
		Exists:   resourceIBMISBareMetalServerNetworkInterfaceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{

			isBareMetalServerID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Bare metal server identifier",
			},
			isBareMetalServerNicAllowIPSpoofing: {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.",
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

						isBareMetalServerNicIpCRN: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this floating IP",
						},
						isBareMetalServerNicIpHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this floating IP",
						},
						isBareMetalServerNicIpID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this floating IP",
						},
						isBareMetalServerNicIpName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this floating IP",
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The network interface type: [ pci, vlan ]",
			},
			isBareMetalServerNicReservedIps: {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The reserved IPs bound to this network interface.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerNicIpAddress: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The globally unique IP address",
						},
						isBareMetalServerNicIpHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this reserved IP",
						},
						isBareMetalServerNicIpID: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The unique identifier for this reserved IP",
						},
						isBareMetalServerNicIpName: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The unique user-defined name for this reserved IP",
						},
						isBareMetalServerNicResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type: [ subnet_reserved_ip ]",
						},
						isBareMetalServerNicIpAutoDelete: {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "If set to true, this reserved IP will be automatically deleted when the target is deleted or when the reserved IP is unbound.",
						},
					},
				},
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
				Description: "title: IPv4, The IP address. ",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerNicIpAddress: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The globally unique IP address",
						},
						isBareMetalServerNicIpHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this reserved IP",
						},
						isBareMetalServerNicIpID: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The unique identifier for this reserved IP",
						},
						isBareMetalServerNicIpName: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The unique user-defined name for this reserved IP",
						},
						isBareMetalServerNicResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type: [ subnet_reserved_ip ]",
						},
						isBareMetalServerNicIpAutoDelete: {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "If set to true, this reserved IP will be automatically deleted when the target is deleted or when the reserved IP is unbound.",
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
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isBareMetalServerNicAllowedVlans},
				Description:   "Indicates if the interface can float to any other server within the same resource_group. The interface will float automatically if the network detects a GARP or RARP on another bare metal server in the resource group. Applies only to vlan type interfaces.",
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

func resourceIBMISBareMetalServerNetworkInterfaceValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 isBareMetalServerName,
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "tag",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISBareMetalServerResourceValidator := ResourceValidator{ResourceName: "ibm_is_bare_metal_server", Schema: validateSchema}
	return &ibmISBareMetalServerResourceValidator
}

func resourceIBMISBareMetalServerNetworkInterfaceCreate(d *schema.ResourceData, meta interface{}) error {

	bareMetalServerId := ""
	if bmsId, ok := d.GetOk(isBareMetalServerID); ok {
		bareMetalServerId = bmsId.(string)
	}
	if allowedVlansOk, ok := d.GetOk(isBareMetalServerNicAllowedVlans); ok {
		sess, err := vpcClient(meta)
		if err != nil {
			return err
		}
		options := &vpcv1.CreateBareMetalServerNetworkInterfaceOptions{}
		interfaceType := "pci"
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
			nicOptions.AllowIPSpoofing = &allowIPSpoofing
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

		if ipsIntf, ok := d.GetOk(isBareMetalServerNicReservedIps); ok {
			ips := ipsIntf.([]interface{})
			var intfs []vpcv1.NetworkInterfaceIPPrototypeIntf
			for _, resource := range ips {
				ip := resource.(map[string]interface{})
				if reservedIpId, ok := ip[isBareMetalServerNicIpID]; ok {
					reservedIpIdStr := reservedIpId.(string)
					nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
						ID: &reservedIpIdStr,
					}
					intfs = append(intfs, nicReservedIP)
				} else {
					reservedIpAddress := ip[isBareMetalServerNicIpAddress].(string)
					reservedIpAutoDelete := ip[isBareMetalServerNicIpAutoDelete].(bool)
					reservedIpName := ip[isBareMetalServerNicIpName].(string)
					nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
						Address:    &reservedIpAddress,
						AutoDelete: &reservedIpAutoDelete,
						Name:       &reservedIpName,
					}
					intfs = append(intfs, nicReservedIP)
				}
			}
			nicOptions.Ips = intfs
		}

		if primaryIpIntf, ok := d.GetOk(isBareMetalServerNicPrimaryIP); ok && len(primaryIpIntf.([]interface{})) > 0 {
			primaryIp := primaryIpIntf.([]interface{})[0].(map[string]interface{})
			if reservedIpId, ok := primaryIp[isBareMetalServerNicIpID]; ok {
				reservedIpIdStr := reservedIpId.(string)
				nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
					ID: &reservedIpIdStr,
				}
				nicOptions.PrimaryIP = nicReservedIP
			} else {
				reservedIpAddress := primaryIp[isBareMetalServerNicIpAddress].(string)
				reservedIpAutoDelete := primaryIp[isBareMetalServerNicIpAutoDelete].(bool)
				reservedIpName := primaryIp[isBareMetalServerNicIpName].(string)
				nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
					Address:    &reservedIpAddress,
					AutoDelete: &reservedIpAutoDelete,
					Name:       &reservedIpName,
				}
				nicOptions.PrimaryIP = nicReservedIP
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
		nic, response, err := sess.CreateBareMetalServerNetworkInterface(options)
		if err != nil || nic == nil {
			return fmt.Errorf("[DEBUG] Create bare metal server (%s) network interface err %s\n%s", bareMetalServerId, err, response)
		}
		err = bareMetalServerNICGet(d, meta, nic, bareMetalServerId)
		if err != nil {
			return err
		}
		_, nicId, err := parseNICTerraformID(d.Id())
		if err != nil {
			return err
		}
		log.Printf("[INFO] Bare Metal Server Network Interface : %s", d.Id())
		_, err = isWaitForBareMetalServerNetworkInterfaceAvailable(sess, bareMetalServerId, nicId, d.Timeout(schema.TimeoutCreate), d)
		if err != nil {
			return err
		}
	} else {
		createVlanTypeNetworkInterface(d, meta, bareMetalServerId)
	}

	return nil
}

func createVlanTypeNetworkInterface(d *schema.ResourceData, meta interface{}, bareMetalServerId string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
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
		nicOptions.AllowIPSpoofing = &allowIPSpoofing
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

	if ipsIntf, ok := d.GetOk(isBareMetalServerNicReservedIps); ok {
		ips := ipsIntf.([]interface{})
		var intfs []vpcv1.NetworkInterfaceIPPrototypeIntf
		for _, resource := range ips {
			ip := resource.(map[string]interface{})
			if reservedIpId, ok := ip[isBareMetalServerNicIpID]; ok {
				reservedIpIdStr := reservedIpId.(string)
				nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
					ID: &reservedIpIdStr,
				}
				intfs = append(intfs, nicReservedIP)
			} else {
				reservedIpAddress := ip[isBareMetalServerNicIpAddress].(string)
				reservedIpAutoDelete := ip[isBareMetalServerNicIpAutoDelete].(bool)
				reservedIpName := ip[isBareMetalServerNicIpName].(string)
				nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
					Address:    &reservedIpAddress,
					AutoDelete: &reservedIpAutoDelete,
					Name:       &reservedIpName,
				}
				intfs = append(intfs, nicReservedIP)
			}
		}
		nicOptions.Ips = intfs
	}

	if primaryIpIntf, ok := d.GetOk(isBareMetalServerNicPrimaryIP); ok && len(primaryIpIntf.([]interface{})) > 0 {
		primaryIp := primaryIpIntf.([]interface{})[0].(map[string]interface{})
		if reservedIpId, ok := primaryIp[isBareMetalServerNicIpID]; ok {
			reservedIpIdStr := reservedIpId.(string)
			nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPIdentity{
				ID: &reservedIpIdStr,
			}
			nicOptions.PrimaryIP = nicReservedIP
		} else {
			reservedIpAddress := primaryIp[isBareMetalServerNicIpAddress].(string)
			reservedIpAutoDelete := primaryIp[isBareMetalServerNicIpAutoDelete].(bool)
			reservedIpName := primaryIp[isBareMetalServerNicIpName].(string)
			nicReservedIP := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
				Address:    &reservedIpAddress,
				AutoDelete: &reservedIpAutoDelete,
				Name:       &reservedIpName,
			}
			nicOptions.PrimaryIP = nicReservedIP
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
	nic, response, err := sess.CreateBareMetalServerNetworkInterface(options)
	if err != nil || nic == nil {
		return fmt.Errorf("[DEBUG] Create bare metal server (%s) network interface err %s\n%s", bareMetalServerId, err, response)
	}
	err = bareMetalServerNICGet(d, meta, nic, bareMetalServerId)
	if err != nil {
		return err
	}
	_, nicId, err := parseNICTerraformID(d.Id())
	if err != nil {
		return err
	}
	log.Printf("[INFO] Bare Metal Server Network Interface : %s", d.Id())
	_, err = isWaitForBareMetalServerNetworkInterfaceAvailable(sess, bareMetalServerId, nicId, d.Timeout(schema.TimeoutCreate), d)
	if err != nil {
		return err
	}
	return nil
}

func resourceIBMISBareMetalServerNetworkInterfaceRead(d *schema.ResourceData, meta interface{}) error {
	bareMetalServerId, nicID, err := parseNICTerraformID(d.Id())
	if err != nil {
		return err
	}

	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
		BareMetalServerID: &bareMetalServerId,
		ID:                &nicID,
	}

	nicIntf, response, err := sess.GetBareMetalServerNetworkInterface(options)
	if err != nil || nicIntf == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting Bare Metal Server (%s) network interface (%s): %s\n%s", bareMetalServerId, nicID, err, response)
	}
	err = bareMetalServerNICGet(d, meta, nicIntf, bareMetalServerId)
	if err != nil {
		return err
	}
	return nil
}

func bareMetalServerNICGet(d *schema.ResourceData, meta interface{}, nicIntf interface{}, bareMetalServerId string) error {

	switch reflect.TypeOf(nicIntf).String() {
	case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
		{

			nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
			d.SetId(makeTerraformNICID(bareMetalServerId, *nic.ID))
			d.Set(isBareMetalServerNicAllowIPSpoofing, *nic.AllowIPSpoofing)
			d.Set(isBareMetalServerNicEnableInfraNAT, *nic.EnableInfrastructureNat)

			floatingIPList := make([]map[string]interface{}, 0)
			if nic.FloatingIps != nil {
				for _, ip := range nic.FloatingIps {
					currentIP := map[string]interface{}{
						isBareMetalServerNicIpHref:    *ip.Href,
						isBareMetalServerNicIpID:      *ip.ID,
						isBareMetalServerNicIpCRN:     *ip.CRN,
						isBareMetalServerNicIpName:    *ip.Name,
						isBareMetalServerNicIpAddress: *ip.Address,
					}
					floatingIPList = append(floatingIPList, currentIP)
				}
			}
			d.Set(isBareMetalServerNicFloatingIPs, floatingIPList)

			d.Set(isBareMetalServerNicHref, *nic.Href)
			d.Set(isBareMetalServerNicID, *nic.ID)
			d.Set(isBareMetalServerNicInterfaceType, *nic.InterfaceType)

			ipsList := make([]map[string]interface{}, 0)
			if nic.Ips != nil {
				for _, ip := range nic.Ips {
					currentIP := map[string]interface{}{
						isBareMetalServerNicIpHref:       *ip.Href,
						isBareMetalServerNicIpID:         *ip.ID,
						isBareMetalServerNicResourceType: *ip.ResourceType,
						isBareMetalServerNicIpName:       *ip.Name,
						isBareMetalServerNicIpAddress:    *ip.Address,
					}
					ipsList = append(ipsList, currentIP)
				}
			}
			d.Set(isBareMetalServerNicReservedIps, ipsList)

			d.Set(isBareMetalServerNicMacAddress, *nic.MacAddress)
			d.Set(isBareMetalServerNicName, *nic.Name)
			if nic.PortSpeed != nil {
				d.Set(isBareMetalServerNicPortSpeed, *nic.PortSpeed)
			}
			primaryIpList := make([]map[string]interface{}, 0)
			if nic.PrimaryIP != nil {
				currentIP := map[string]interface{}{
					isBareMetalServerNicIpHref:       *nic.PrimaryIP.Href,
					isBareMetalServerNicIpID:         *nic.PrimaryIP.ID,
					isBareMetalServerNicResourceType: *nic.PrimaryIP.ResourceType,
					isBareMetalServerNicIpName:       *nic.PrimaryIP.Name,
					isBareMetalServerNicIpAddress:    *nic.PrimaryIP.Address,
				}
				primaryIpList = append(primaryIpList, currentIP)
			}
			d.Set(isBareMetalServerNicPrimaryIP, primaryIpList)

			d.Set(isBareMetalServerNicResourceType, *nic.ResourceType)

			if nic.SecurityGroups != nil && len(nic.SecurityGroups) != 0 {
				secgrpList := []string{}
				for i := 0; i < len(nic.SecurityGroups); i++ {
					secgrpList = append(secgrpList, string(*(nic.SecurityGroups[i].ID)))
				}
				d.Set(isBareMetalServerNicSecurityGroups, newStringSet(schema.HashString, secgrpList))
			}

			d.Set(isBareMetalServerNicStatus, *nic.Status)
			d.Set(isBareMetalServerNicSubnet, *nic.Subnet.ID)
			d.Set(isBareMetalServerNicType, *nic.Type)
			if nic.AllowedVlans != nil {
				var out = make([]interface{}, len(nic.AllowedVlans), len(nic.AllowedVlans))
				for i, v := range nic.AllowedVlans {
					out[i] = int(v)
				}
				d.Set(isBareMetalServerNicAllowedVlans, schema.NewSet(schema.HashInt, out))
			}
		}
	case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
		{
			nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
			d.SetId(makeTerraformNICID(bareMetalServerId, *nic.ID))
			d.Set(isBareMetalServerNicAllowIPSpoofing, *nic.AllowIPSpoofing)
			d.Set(isBareMetalServerNicEnableInfraNAT, *nic.EnableInfrastructureNat)

			floatingIPList := make([]map[string]interface{}, 0)
			if nic.FloatingIps != nil {
				for _, ip := range nic.FloatingIps {
					currentIP := map[string]interface{}{
						isBareMetalServerNicIpHref:    *ip.Href,
						isBareMetalServerNicIpID:      *ip.ID,
						isBareMetalServerNicIpCRN:     *ip.CRN,
						isBareMetalServerNicIpName:    *ip.Name,
						isBareMetalServerNicIpAddress: *ip.Address,
					}
					floatingIPList = append(floatingIPList, currentIP)
				}
			}
			d.Set(isBareMetalServerNicFloatingIPs, floatingIPList)

			d.Set(isBareMetalServerNicHref, *nic.Href)
			d.Set(isBareMetalServerNicID, *nic.ID)
			d.Set(isBareMetalServerNicInterfaceType, *nic.InterfaceType)

			ipsList := make([]map[string]interface{}, 0)
			if nic.Ips != nil {
				for _, ip := range nic.Ips {
					currentIP := map[string]interface{}{
						isBareMetalServerNicIpHref:       *ip.Href,
						isBareMetalServerNicIpID:         *ip.ID,
						isBareMetalServerNicResourceType: *ip.ResourceType,
						isBareMetalServerNicIpName:       *ip.Name,
						isBareMetalServerNicIpAddress:    *ip.Address,
					}
					ipsList = append(ipsList, currentIP)
				}
			}
			d.Set(isBareMetalServerNicReservedIps, ipsList)

			d.Set(isBareMetalServerNicMacAddress, *nic.MacAddress)
			d.Set(isBareMetalServerNicName, *nic.Name)
			d.Set(isBareMetalServerNicPortSpeed, *nic.PortSpeed)

			primaryIpList := make([]map[string]interface{}, 0)
			if nic.PrimaryIP != nil {
				currentIP := map[string]interface{}{
					isBareMetalServerNicIpHref:       *nic.PrimaryIP.Href,
					isBareMetalServerNicIpID:         *nic.PrimaryIP.ID,
					isBareMetalServerNicResourceType: *nic.PrimaryIP.ResourceType,
					isBareMetalServerNicIpName:       *nic.PrimaryIP.Name,
					isBareMetalServerNicIpAddress:    *nic.PrimaryIP.Address,
				}
				primaryIpList = append(primaryIpList, currentIP)
			}
			d.Set(isBareMetalServerNicPrimaryIP, primaryIpList)

			d.Set(isBareMetalServerNicResourceType, *nic.ResourceType)

			if len(nic.SecurityGroups) != 0 {
				secgrpList := []string{}
				for i := 0; i < len(nic.SecurityGroups); i++ {
					secgrpList = append(secgrpList, string(*(nic.SecurityGroups[i].ID)))
				}
				d.Set(isBareMetalServerNicSecurityGroups, newStringSet(schema.HashString, secgrpList))
			}

			d.Set(isBareMetalServerNicStatus, *nic.Status)
			d.Set(isBareMetalServerNicSubnet, *nic.Subnet.ID)
			d.Set(isBareMetalServerNicType, *nic.Type)
			d.Set(isBareMetalServerNicAllowInterfaceToFloat, *nic.AllowInterfaceToFloat)
			d.Set(isBareMetalServerNicVlan, *nic.Vlan)
		}
	}
	return nil
}

func resourceIBMISBareMetalServerNetworkInterfaceUpdate(d *schema.ResourceData, meta interface{}) error {

	bareMetalServerId, nicId, err := parseNICTerraformID(d.Id())
	if err != nil {
		return err
	}
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	options := &vpcv1.UpdateBareMetalServerNetworkInterfaceOptions{
		BareMetalServerID: &bareMetalServerId,
		ID:                &nicId,
	}
	nicPatchModel := &vpcv1.BareMetalServerNetworkInterfacePatch{}
	flag := false
	if d.HasChange(isBareMetalServerNicAllowIPSpoofing) {
		flag = true
		aisBool := false
		if ais, ok := d.GetOk(isBareMetalServerNicEnableInfraNAT); ok {
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
			return fmt.Errorf("Error calling asPatch for BareMetalServerNetworkInterfacePatch %s", err)
		}
		options.BareMetalServerNetworkInterfacePatch = nicPatchModelAsPatch

		nicIntf, response, err := sess.UpdateBareMetalServerNetworkInterface(options)
		if err != nil {
			return fmt.Errorf("Error updating Bare Metal Server: %s\n%s", err, response)
		}
		return bareMetalServerNICGet(d, meta, nicIntf, bareMetalServerId)
	}

	return nil
}

func resourceIBMISBareMetalServerNetworkInterfaceDelete(d *schema.ResourceData, meta interface{}) error {
	bareMetalServerId, nicId, err := parseNICTerraformID(d.Id())
	if err != nil {
		return err
	}

	err = bareMetalServerNetworkInterfaceDelete(d, meta, bareMetalServerId, nicId)
	if err != nil {
		return err
	}

	return nil
}

func bareMetalServerNetworkInterfaceDelete(d *schema.ResourceData, meta interface{}, bareMetalServerId, nicId string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	getBmsNicOptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
		BareMetalServerID: &bareMetalServerId,
		ID:                &nicId,
	}
	_, response, err := sess.GetBareMetalServerNetworkInterface(getBmsNicOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		return fmt.Errorf("Error getting Bare Metal Server (%s) network interface(%s) : %s\n%s", bareMetalServerId, nicId, err, response)
	}

	options := &vpcv1.DeleteBareMetalServerNetworkInterfaceOptions{
		BareMetalServerID: &bareMetalServerId,
		ID:                &nicId,
	}
	response, err = sess.DeleteBareMetalServerNetworkInterface(options)
	if err != nil {
		return fmt.Errorf("Error Deleting Bare Metal Server (%s) network interface (%s) : %s\n%s", bareMetalServerId, nicId, err, response)
	}
	_, err = isWaitForBareMetalServerNetworkInterfaceDeleted(sess, bareMetalServerId, nicId, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func isWaitForBareMetalServerNetworkInterfaceDeleted(bmsC *vpcv1.VpcV1, bareMetalServerId, nicId string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for (%s) / (%s) to be deleted.", bareMetalServerId, nicId)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{isBareMetalServerNetworkInterfaceAvailable, isBareMetalServerNetworkInterfaceDeleting, isBareMetalServerNetworkInterfacePending},
		Target:     []string{isBareMetalServerNetworkInterfaceDeleted, isBareMetalServerNetworkInterfaceFailed, ""},
		Refresh:    isBareMetalServerNetworkInterfaceDeleteRefreshFunc(bmsC, bareMetalServerId, nicId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isBareMetalServerNetworkInterfaceDeleteRefreshFunc(bmsC *vpcv1.VpcV1, bareMetalServerId, nicId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getBmsNicOptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
			BareMetalServerID: &bareMetalServerId,
			ID:                &nicId,
		}
		bmsNic, response, err := bmsC.GetBareMetalServerNetworkInterface(getBmsNicOptions)

		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return bmsNic, isBareMetalServerNetworkInterfaceDeleted, nil
			}
			return bmsNic, isBareMetalServerNetworkInterfaceFailed, fmt.Errorf("Error getting Bare Metal Server(%s) Network Interface (%s): %s\n%s", bareMetalServerId, nicId, err, response)
		}
		return bmsNic, isBareMetalServerNetworkInterfaceDeleting, err
	}
}

func resourceIBMISBareMetalServerNetworkInterfaceExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	bareMetalServerId, nicId, err := parseNICTerraformID(d.Id())
	if err != nil {
		return false, err
	}

	exists, err := bareMetalServerNetworkInterfaceExists(d, meta, bareMetalServerId, nicId)
	return exists, err

}

func bareMetalServerNetworkInterfaceExists(d *schema.ResourceData, meta interface{}, bareMetalServerId, nicId string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	getBmsNicOptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
		BareMetalServerID: &bareMetalServerId,
		ID:                &nicId,
	}
	_, response, err := sess.GetBareMetalServerNetworkInterface(getBmsNicOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting Bare Metal Server Network Interface : %s\n%s", err, response)
	}
	return true, nil
}

func isWaitForBareMetalServerNetworkInterfaceAvailable(client *vpcv1.VpcV1, bareMetalServerId, nicId string, timeout time.Duration, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for Bare Metal Server (%s) Network Interface (%s) to be available.", bareMetalServerId, nicId)
	communicator := make(chan interface{})
	stateConf := &resource.StateChangeConf{
		Pending:    []string{isBareMetalServerNetworkInterfacePending},
		Target:     []string{isBareMetalServerNetworkInterfaceAvailable, isBareMetalServerNetworkInterfaceFailed},
		Refresh:    isBareMetalServerNetworkInterfaceRefreshFunc(client, bareMetalServerId, nicId, d, communicator),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isBareMetalServerNetworkInterfaceRefreshFunc(client *vpcv1.VpcV1, bareMetalServerId, nicId string, d *schema.ResourceData, communicator chan interface{}) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getBmsNicOptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
			BareMetalServerID: &bareMetalServerId,
			ID:                &nicId,
		}
		bmsNic, response, err := client.GetBareMetalServerNetworkInterface(getBmsNicOptions)
		if err != nil {
			return nil, "", fmt.Errorf("Error getting Bare Metal Server (%s) Network Interface (%s) : %s\n%s", bareMetalServerId, nicId, err, response)
		}
		status := ""
		switch reflect.TypeOf(bmsNic).String() {
		case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
			{
				nic := bmsNic.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
				status = *nic.Status
				d.Set(isBareMetalServerNicStatus, *nic.Status)
			}
		case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
			{
				nic := bmsNic.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
				status = *nic.Status
				d.Set(isBareMetalServerNicStatus, *nic.Status)
			}
		}

		select {
		case data := <-communicator:
			return nil, "", data.(error)
		default:
			fmt.Println("no message sent")
		}

		if status == "available" || status == "failed" {
			close(communicator)
			return bmsNic, status, nil

		}
		return bmsNic, "pending", nil
	}
}

func makeTerraformNICID(id1, id2 string) string {
	// Include both bare metal sever id and network interface id to create a unique Terraform id.  As a bonus,
	// we can extract the bare metal sever id as needed for API calls such as READ.
	return fmt.Sprintf("%s/%s", id1, id2)
}

func parseNICTerraformID(s string) (string, string, error) {
	segments := strings.Split(s, "/")
	if len(segments) != 2 {
		return "", "", fmt.Errorf("invalid terraform Id %s (incorrect number of segments)", s)
	}
	if segments[0] == "" || segments[1] == "" {
		return "", "", fmt.Errorf("invalid terraform Id %s (one or more empty segments)", s)
	}
	return segments[0], segments[1], nil
}
