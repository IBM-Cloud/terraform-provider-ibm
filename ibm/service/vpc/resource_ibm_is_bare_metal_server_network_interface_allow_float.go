// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isFloatedBareMetalServerID = "floating_bare_metal_server"
)

func ResourceIBMIsBareMetalServerNetworkInterfaceAllowFloat() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISBareMetalServerNetworkInterfaceAllowFloatCreate,
		ReadContext:   resourceIBMISBareMetalServerNetworkInterfaceAllowFloatRead,
		UpdateContext: resourceIBMISBareMetalServerNetworkInterfaceAllowFloatUpdate,
		DeleteContext: resourceIBMISBareMetalServerNetworkInterfaceAllowFloatDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isBareMetalServerID: {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "Bare metal server identifier",
			},

			isFloatedBareMetalServerID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Bare metal server identifier of the server to which nic is floating to",
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The network interface type: [ pci, vlan ]",
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
						isBareMetalServerNicIpAutoDelete: {
							Type:          schema.TypeBool,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"primary_ip.0.reserved_ip"},
							Description:   "Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.",
						},
						isBareMetalServerNicIpName: {
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"primary_ip.0.reserved_ip"},
							Description:   "The user-defined name for this reserved IP. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the subnet the reserved IP resides in. ",
						},
						isBareMetalServerNicIpID: {
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"primary_ip.0.address", "primary_ip.0.auto_delete", "primary_ip.0.name"},
							Description:   "Identifies a reserved IP by a unique property.",
						},
						isBareMetalServerNicResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type",
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
			isBareMetalServerNicVlan: {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Indicates the 802.1Q VLAN ID tag that must be used for all traffic on this interface",
			},
			isBareMetalServerNicAllowInterfaceToFloat: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if the interface can float to any other server within the same resource_group. The interface will float automatically if the network detects a GARP or RARP on another bare metal server in the resource group. Applies only to vlan type interfaces.",
			},
		},
	}
}

func resourceIBMISBareMetalServerNetworkInterfaceAllowFloatCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	bareMetalServerId := ""
	if bmsId, ok := d.GetOk(isBareMetalServerID); ok {
		bareMetalServerId = bmsId.(string)
	}

	err := createVlanTypeNetworkInterfaceAllowFloat(context, d, meta, bareMetalServerId)
	if err != nil {
		return err
	}
	return nil
}

func createVlanTypeNetworkInterfaceAllowFloat(context context.Context, d *schema.ResourceData, meta interface{}, bareMetalServerId string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.CreateBareMetalServerNetworkInterfaceOptions{}
	interfaceType := "vlan"
	nicOptions := &vpcv1.BareMetalServerNetworkInterfacePrototypeBareMetalServerNetworkInterfaceByVlanPrototype{}
	allowInterfaceToFloat := true
	nicOptions.AllowInterfaceToFloat = &allowInterfaceToFloat
	if vlan, ok := d.GetOk(isBareMetalServerNicVlan); ok {
		vlanInt := int64(vlan.(int))
		nicOptions.Vlan = &vlanInt
	}

	if name, ok := d.GetOk(isBareMetalServerNicName); ok {
		nameStr := name.(string)
		nicOptions.Name = &nameStr
	}
	nicOptions.InterfaceType = &interfaceType

	if aisOk, ok := d.GetOkExists(isBareMetalServerNicAllowIPSpoofing); ok {
		allowIPSpoofing := aisOk.(bool)
		nicOptions.AllowIPSpoofing = &allowIPSpoofing
	}

	if ein, ok := d.GetOkExists(isBareMetalServerNicEnableInfraNAT); ok {
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
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.Set(isFloatedBareMetalServerID, bareMetalServerId)
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
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "create", "sep-id-parts").GetDiag()
	}

	log.Printf("[INFO] Bare Metal Server Network Interface : %s", d.Id())
	nicAfterWait, err := isWaitForBareMetalServerNetworkInterfaceAvailable(sess, bareMetalServerId, nicId, d.Timeout(schema.TimeoutCreate), d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerNetworkInterfaceAvailable failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	diagErr := bareMetalServerNICAllowFloatGet(context, d, meta, sess, nicAfterWait, bareMetalServerId)
	if diagErr != nil {
		return diagErr
	}

	return nil
}

func resourceIBMISBareMetalServerNetworkInterfaceAllowFloatRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	bareMetalServerId, nicID, err := ParseNICTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "sep-id-parts").GetDiag()
	}
	d.Set(isFloatedBareMetalServerID, bareMetalServerId)

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
		BareMetalServerID: &bareMetalServerId,
		ID:                &nicID,
	}
	var nicIntf vpcv1.BareMetalServerNetworkInterfaceIntf
	// try to fetch original nic
	nicIntf, response, err := sess.GetBareMetalServerNetworkInterfaceWithContext(context, options)
	if (err != nil || nicIntf == nil) && response != nil {
		//if original nic is not present, try fetching nic without server id
		nicIntf, response, err = findNicsWithoutBMS(context, d, sess, nicID)
		// response here can be either nil or not nil and if it returns 404 means nic is deleted
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		// if response returns an error
		if err != nil || nicIntf == nil {
			if response != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("findNicsWithoutBMS failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			} else {
				d.SetId("")
				return nil
				// return diag.FromErr(fmt.Errorf("[ERROR] Error getting Bare Metal Server2 (%s) network interface (%s): %s", bareMetalServerId, nicID, err))
			} // else is returning that the nic is not found anywhere
		}
	}
	diagErr := bareMetalServerNICAllowFloatGet(context, d, meta, sess, nicIntf, bareMetalServerId)
	if diagErr != nil {
		return diagErr
	}
	return nil
}

func findNicsWithoutBMS(context context.Context, d *schema.ResourceData, sess *vpcv1.VpcV1, nicId string) (result vpcv1.BareMetalServerNetworkInterfaceIntf, response *core.DetailedResponse, err error) {
	// listing all servers
	start := ""
	allrecs := []vpcv1.BareMetalServer{}
	for {
		listBareMetalServersOptions := &vpcv1.ListBareMetalServersOptions{}
		if start != "" {
			listBareMetalServersOptions.Start = &start
		}
		availableServers, _, err := sess.ListBareMetalServersWithContext(context, listBareMetalServersOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListBareMetalServersWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return nil, nil, tfErr
		}
		start = flex.GetNext(availableServers.Next)
		allrecs = append(allrecs, availableServers.BareMetalServers...)
		if start == "" {
			break
		}
	}
	// finding nic id each server
	for _, server := range allrecs {
		nics := server.NetworkInterfaces
		for _, nic := range nics {
			if *nic.ID == nicId {
				options := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
					BareMetalServerID: server.ID,
					ID:                &nicId,
				}
				//return response of the server nic matches
				d.Set(isFloatedBareMetalServerID, *server.ID)
				return sess.GetBareMetalServerNetworkInterfaceWithContext(context, options)
			}
		}
	}
	err = fmt.Errorf("ListBareMetalServersWithContext failed: Network interface not found")
	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerNetworkInterfaceFloatingIPWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "read")
	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	// if not found return nil response and error
	return nil, nil, tfErr
}

func bareMetalServerNICAllowFloatGet(context context.Context, d *schema.ResourceData, meta interface{}, sess *vpcv1.VpcV1, nicIntf interface{}, bareMetalServerId string) diag.Diagnostics {
	var err error
	switch reflect.TypeOf(nicIntf).String() {
	case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
		{
			nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
			if err = d.Set(isBareMetalServerNicAllowIPSpoofing, *nic.AllowIPSpoofing); err != nil {
				err = fmt.Errorf("Error setting allow_ip_spoofing: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-allow_ip_spoofing").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicEnableInfraNAT, *nic.EnableInfrastructureNat); err != nil {
				err = fmt.Errorf("Error setting enable_infrastructure_nat: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-enable_infrastructure_nat").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicStatus, *nic.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-status").GetDiag()
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
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-floating_ips").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicHref, nic.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-href").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicID, *nic.ID); err != nil {
				err = fmt.Errorf("Error setting network_interface: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-network_interface").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicInterfaceType, *nic.InterfaceType); err != nil {
				err = fmt.Errorf("Error setting interface_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-interface_type").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicMacAddress, *nic.MacAddress); err != nil {
				err = fmt.Errorf("Error setting mac_address: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-mac_address").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicName, *nic.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-name").GetDiag()
			}

			if nic.PortSpeed != nil {
				if err = d.Set(isBareMetalServerNicPortSpeed, nic.PortSpeed); err != nil {
					err = fmt.Errorf("Error setting port_speed: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-port_speed").GetDiag()
				}
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
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			}
			currentIP[isBareMetalServerNicIpAutoDelete] = bmsRip.AutoDelete

			primaryIpList = append(primaryIpList, currentIP)

			if err = d.Set(isBareMetalServerNicPrimaryIP, primaryIpList); err != nil {
				err = fmt.Errorf("Error setting primary_ip: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-primary_ip").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicResourceType, nic.ResourceType); err != nil {
				err = fmt.Errorf("Error setting resource_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-resource_type").GetDiag()
			}
			if nic.SecurityGroups != nil && len(nic.SecurityGroups) != 0 {
				secgrpList := []string{}
				for i := 0; i < len(nic.SecurityGroups); i++ {
					secgrpList = append(secgrpList, string(*(nic.SecurityGroups[i].ID)))
				}
				if err = d.Set(isBareMetalServerNicSecurityGroups, flex.NewStringSet(schema.HashString, secgrpList)); err != nil {
					err = fmt.Errorf("Error setting security_groups: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-security_groups").GetDiag()
				}
			}

			if err = d.Set(isBareMetalServerNicStatus, *nic.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-status").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicSubnet, *nic.Subnet.ID); err != nil {
				err = fmt.Errorf("Error setting subnet: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-subnet").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicType, *nic.Type); err != nil {
				err = fmt.Errorf("Error setting type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-type").GetDiag()
			}
		}
	case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
		{
			nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
			d.SetId(MakeTerraformNICID(bareMetalServerId, *nic.ID))
			if err = d.Set(isBareMetalServerNicAllowIPSpoofing, *nic.AllowIPSpoofing); err != nil {
				err = fmt.Errorf("Error setting allow_ip_spoofing: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-allow_ip_spoofing").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicEnableInfraNAT, *nic.EnableInfrastructureNat); err != nil {
				err = fmt.Errorf("Error setting enable_infrastructure_nat: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-enable_infrastructure_nat").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicStatus, *nic.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-status").GetDiag()
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
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-floating_ips").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicHref, nic.Href); err != nil {
				err = fmt.Errorf("Error setting href: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-href").GetDiag()
			}

			if err = d.Set(isBareMetalServerNicID, *nic.ID); err != nil {
				err = fmt.Errorf("Error setting network_interface: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-network_interface").GetDiag()
			}

			if err = d.Set(isBareMetalServerNicInterfaceType, *nic.InterfaceType); err != nil {
				err = fmt.Errorf("Error setting interface_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-interface_type").GetDiag()
			}

			if err = d.Set(isBareMetalServerNicMacAddress, *nic.MacAddress); err != nil {
				err = fmt.Errorf("Error setting mac_address: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-mac_address").GetDiag()
			}

			if err = d.Set(isBareMetalServerNicName, *nic.Name); err != nil {
				err = fmt.Errorf("Error setting name: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-name").GetDiag()
			}

			if err = d.Set(isBareMetalServerNicPortSpeed, nic.PortSpeed); err != nil {
				err = fmt.Errorf("Error setting port_speed: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-port_speed").GetDiag()
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
			primaryIpList = append(primaryIpList, currentIP)
			if err = d.Set(isBareMetalServerNicPrimaryIP, primaryIpList); err != nil {
				err = fmt.Errorf("Error setting primary_ip: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-primary_ip").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicResourceType, nic.ResourceType); err != nil {
				err = fmt.Errorf("Error setting resource_type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-resource_type").GetDiag()
			}
			if len(nic.SecurityGroups) != 0 {
				secgrpList := []string{}
				for i := 0; i < len(nic.SecurityGroups); i++ {
					secgrpList = append(secgrpList, string(*(nic.SecurityGroups[i].ID)))
				}
				if err = d.Set(isBareMetalServerNicSecurityGroups, flex.NewStringSet(schema.HashString, secgrpList)); err != nil {
					err = fmt.Errorf("Error setting security_groups: %s", err)
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-security_groups").GetDiag()
				}
			}
			if err = d.Set(isBareMetalServerNicStatus, *nic.Status); err != nil {
				err = fmt.Errorf("Error setting status: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-status").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicSubnet, *nic.Subnet.ID); err != nil {
				err = fmt.Errorf("Error setting subnet: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-subnet").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicType, *nic.Type); err != nil {
				err = fmt.Errorf("Error setting type: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-type").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicAllowInterfaceToFloat, *nic.AllowInterfaceToFloat); err != nil {
				err = fmt.Errorf("Error setting allow_interface_to_float: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-allow_interface_to_float").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicVlan, *nic.Vlan); err != nil {
				err = fmt.Errorf("Error setting vlan: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "read", "set-vlan").GetDiag()
			}
		}
	}
	return nil
}

func resourceIBMISBareMetalServerNetworkInterfaceAllowFloatUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	bareMetalServerId, nicId, err := ParseNICTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "update", "sep-id-parts").GetDiag()
	}
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "update", "initialize-client")
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
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSecurityGroupTargetBindingWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForBareMetalServerAvailableForNIC(sess, bareMetalServerId, d.Timeout(schema.TimeoutUpdate), d)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerAvailableForNIC failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "update")
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
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSecurityGroupTargetBindingWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForBareMetalServerAvailableForNIC(sess, bareMetalServerId, d.Timeout(schema.TimeoutUpdate), d)
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerAvailableForNIC failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "update")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		}
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
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("reservedIpPath.AsPatch() failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateripoptions.ReservedIPPatch = reservedIpPathAsPatch
		_, _, err = sess.UpdateSubnetReservedIPWithContext(context, updateripoptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
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
		if ais, ok := d.GetOk(isBareMetalServerNicAllowIPSpoofing); ok {
			aisBool = ais.(bool)
		}
		nicPatchModel.AllowIPSpoofing = &aisBool
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
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("nicPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options.BareMetalServerNetworkInterfacePatch = nicPatchModelAsPatch

		nicIntf, _, err := sess.UpdateBareMetalServerNetworkInterfaceWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		return bareMetalServerNICAllowFloatGet(context, d, meta, sess, nicIntf, bareMetalServerId)
	}

	return nil
}

func resourceIBMISBareMetalServerNetworkInterfaceAllowFloatDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	bareMetalServerId, nicId, err := ParseNICTerraformID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "delete", "sep-id-parts").GetDiag()
	}

	diagErr := bareMetalServerNetworkInterfaceAllowFloatDelete(context, d, meta, bareMetalServerId, nicId)
	if diagErr != nil {
		return diagErr
	}

	return nil
}

func bareMetalServerNetworkInterfaceAllowFloatDelete(context context.Context, d *schema.ResourceData, meta interface{}, bareMetalServerId, nicId string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_network_interface_allow_float", "delete", "initialize-client")
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
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "delete")
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
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "delete")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			// failed, pending, restarting, running, starting, stopped, stopping, maintenance
			if *bms.Status == "failed" {
				err = fmt.Errorf("Error cannot detach network interface from a failed bare metal server")
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "delete")
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
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("StopBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "delete")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err = isWaitForBareMetalServerStoppedForNIC(sess, bareMetalServerId, d.Timeout(schema.TimeoutCreate), d)
				if err != nil || res.StatusCode != 204 {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerStoppedForNIC failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "delete")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			} else if *bms.Status != "stopped" {
				err = fmt.Errorf("Error bare metal server in %s state, please try after some time", *bms.Status)
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "delete")
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
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForBareMetalServerNetworkInterfaceDeleted(sess, bareMetalServerId, nicId, nicType, nicIntf, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerNetworkInterfaceDeleted failed: %s", err.Error()), "ibm_is_bare_metal_server_network_interface_allow_float", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}
