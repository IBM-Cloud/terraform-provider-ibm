// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerNicEnableInfraNAT        = "enable_infrastructure_nat"
	isBareMetalServerNicFloatingIPs           = "floating_ips"
	isBareMetalServerNicFloatingIPId          = "id"
	isBareMetalServerNicIpAddress             = "address"
	isBareMetalServerNicIpCRN                 = "crn"
	isBareMetalServerNicIpHref                = "href"
	isBareMetalServerNicIpID                  = "reserved_ip"
	isBareMetalServerNicIpName                = "name"
	isBareMetalServerNicIpAutoDelete          = "auto_delete"
	isBareMetalServerNicHref                  = "href"
	isBareMetalServerNicID                    = "network_interface"
	isBareMetalServerNicInterfaceType         = "interface_type"
	isBareMetalServerNicReservedIps           = "ips"
	isBareMetalServerNicMacAddress            = "mac_address"
	isBareMetalServerNicPrimaryIP             = "primary_ip"
	isBareMetalServerNicResourceType          = "resource_type"
	isBareMetalServerNicStatus                = "status"
	isBareMetalServerNicType                  = "type"
	isBareMetalServerNicAllowedVlans          = "allowed_vlans"
	isBareMetalServerNicAllowInterfaceToFloat = "allow_interface_to_float"
	isBareMetalServerNicVlan                  = "vlan"
)

func DataSourceIBMIsBareMetalServerNetworkInterface() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISBareMetalServerNetworkInterfaceRead,

		Schema: map[string]*schema.Schema{
			isBareMetalServerID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The bare metal server identifier",
			},

			isBareMetalServerNicID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The bare metal server network interface identifier",
			},
			//network interface properties

			isBareMetalServerNicAllowIPSpoofing: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.",
			},
			isBareMetalServerNicEnableInfraNAT: {
				Type:        schema.TypeBool,
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
							Deprecated:  "This field is deprecated - replaced by id",
							Description: "The unique identifier for this floating IP",
						},
						isBareMetalServerNicFloatingIPId: {
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

			isBareMetalServerNicMacAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The MAC address of the interface. If absent, the value is not known.",
			},
			isBareMetalServerNicName: {
				Type:        schema.TypeString,
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
				Computed:    true,
				Description: "IPv4, The IP address. ",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerNicIpAddress: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique IP address",
						},
						isBareMetalServerNicIpHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this reserved IP",
						},
						isBareMetalServerNicIpName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this reserved IP. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the subnet the reserved IP resides in. ",
						},
						isBareMetalServerNicIpID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identifies a reserved IP by a unique property.",
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
				Computed:    true,
				Description: "The id of the associated subnet",
			},

			isBareMetalServerNicType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this bare metal server network interface : [ primary, secondary ]",
			},

			isBareMetalServerNicAllowedVlans: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Set:         schema.HashInt,
				Description: "Indicates what VLAN IDs (for VLAN type only) can use this physical (PCI type) interface. A given VLAN can only be in the allowed_vlans array for one PCI type adapter per bare metal server.",
			},

			isBareMetalServerNicAllowInterfaceToFloat: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if the interface can float to any other server within the same resource_group. The interface will float automatically if the network detects a GARP or RARP on another bare metal server in the resource group. Applies only to vlan type interfaces.",
			},

			isBareMetalServerNicVlan: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Indicates the 802.1Q VLAN ID tag that must be used for all traffic on this interface",
			},
		},
	}
}

func dataSourceIBMISBareMetalServerNetworkInterfaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	bareMetalServerID := d.Get(isBareMetalServerID).(string)
	bareMetalServerNicID := d.Get(isBareMetalServerNicID).(string)
	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_network_interface", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
		BareMetalServerID: &bareMetalServerID,
		ID:                &bareMetalServerNicID,
	}

	nicIntf, _, err := sess.GetBareMetalServerNetworkInterfaceWithContext(context, options)
	if err != nil || nicIntf == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "(Data) ibm_is_bare_metal_server_network_interface", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	switch reflect.TypeOf(nicIntf).String() {
	case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
		{
			nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
			d.SetId(*nic.ID)
			if err = d.Set(isBareMetalServerNicAllowIPSpoofing, *nic.AllowIPSpoofing); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allow_ip_spoofing: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-allow_ip_spoofing").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicEnableInfraNAT, *nic.EnableInfrastructureNat); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting enable_infrastructure_nat: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-enable_infrastructure_nat").GetDiag()
			}
			floatingIPList := make([]map[string]interface{}, 0)
			if nic.FloatingIps != nil {
				for _, ip := range nic.FloatingIps {
					currentIP := map[string]interface{}{
						isBareMetalServerNicIpID:         *ip.ID,
						isBareMetalServerNicFloatingIPId: *ip.ID,
						isBareMetalServerNicIpAddress:    *ip.Address,
					}
					floatingIPList = append(floatingIPList, currentIP)
				}
			}
			if err = d.Set(isBareMetalServerNicFloatingIPs, floatingIPList); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting floating_ips: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-floating_ips").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicHref, *nic.Href); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-href").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicID, *nic.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_interface: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-network_interface").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicInterfaceType, *nic.InterfaceType); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting interface_type: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-interface_type").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicMacAddress, *nic.MacAddress); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting mac_address: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-mac_address").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicName, *nic.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-name").GetDiag()
			}
			if nic.PortSpeed != nil {
				if err = d.Set(isBareMetalServerNicPortSpeed, *nic.PortSpeed); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting port_speed: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-port_speed").GetDiag()
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
			primaryIpList = append(primaryIpList, currentIP)

			if err = d.Set(isBareMetalServerNicPrimaryIP, primaryIpList); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_ip: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-primary_ip").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicResourceType, *nic.ResourceType); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-resource_type").GetDiag()
			}
			if nic.SecurityGroups != nil && len(nic.SecurityGroups) != 0 {
				secgrpList := []string{}
				for i := 0; i < len(nic.SecurityGroups); i++ {
					secgrpList = append(secgrpList, string(*(nic.SecurityGroups[i].ID)))
				}
				if err = d.Set(isBareMetalServerNicSecurityGroups, flex.NewStringSet(schema.HashString, secgrpList)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting security_groups: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-security_groups").GetDiag()
				}
			}
			if err = d.Set(isBareMetalServerNicStatus, *nic.Status); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-status").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicSubnet, *nic.Subnet.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting subnet: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-subnet").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicType, *nic.Type); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-type").GetDiag()
			}
			if nic.AllowedVlans != nil {
				var out = make([]interface{}, len(nic.AllowedVlans), len(nic.AllowedVlans))
				for i, v := range nic.AllowedVlans {
					out[i] = int(v)
				}
				if err = d.Set(isBareMetalServerNicAllowedVlans, schema.NewSet(schema.HashInt, out)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_vlans: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-allowed_vlans").GetDiag()
				}
			}
		}
	case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
		{
			nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
			d.SetId(*nic.ID)

			if err = d.Set(isBareMetalServerNicAllowIPSpoofing, *nic.AllowIPSpoofing); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allow_ip_spoofing: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-allow_ip_spoofing").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicEnableInfraNAT, *nic.EnableInfrastructureNat); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting enable_infrastructure_nat: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-enable_infrastructure_nat").GetDiag()
			}
			floatingIPList := make([]map[string]interface{}, 0)
			if nic.FloatingIps != nil {
				for _, ip := range nic.FloatingIps {
					currentIP := map[string]interface{}{
						isBareMetalServerNicIpID:         *ip.ID,
						isBareMetalServerNicFloatingIPId: *ip.ID,
						isBareMetalServerNicIpAddress:    *ip.Address,
					}
					floatingIPList = append(floatingIPList, currentIP)
				}
			}

			if err = d.Set(isBareMetalServerNicFloatingIPs, floatingIPList); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting floating_ips: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-floating_ips").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicHref, *nic.Href); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-href").GetDiag()
			}

			if err = d.Set(isBareMetalServerNicID, *nic.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_interface: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-network_interface").GetDiag()
			}

			if err = d.Set(isBareMetalServerNicInterfaceType, *nic.InterfaceType); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting interface_type: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-interface_type").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicMacAddress, *nic.MacAddress); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting mac_address: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-mac_address").GetDiag()
			}

			if err = d.Set(isBareMetalServerNicName, *nic.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-name").GetDiag()
			}
			if nic.PortSpeed != nil {
				if err = d.Set(isBareMetalServerNicPortSpeed, *nic.PortSpeed); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting port_speed: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-port_speed").GetDiag()
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
				primaryIpList = append(primaryIpList, currentIP)

				if err = d.Set(isBareMetalServerNicPrimaryIP, primaryIpList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_ip: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-primary_ip").GetDiag()
				}
			}

			if err = d.Set(isBareMetalServerNicResourceType, *nic.ResourceType); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-resource_type").GetDiag()
			}
			if len(nic.SecurityGroups) != 0 {
				secgrpList := []string{}
				for i := 0; i < len(nic.SecurityGroups); i++ {
					secgrpList = append(secgrpList, string(*(nic.SecurityGroups[i].ID)))
				}
				if err = d.Set(isBareMetalServerNicSecurityGroups, flex.NewStringSet(schema.HashString, secgrpList)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting security_groups: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-security_groups").GetDiag()
				}
			}

			if err = d.Set(isBareMetalServerNicStatus, *nic.Status); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-status").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicSubnet, *nic.Subnet.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting subnet: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-subnet").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicType, *nic.Type); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-type").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicAllowInterfaceToFloat, *nic.AllowInterfaceToFloat); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allow_interface_to_float: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-allow_interface_to_float").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicVlan, *nic.Vlan); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vlan: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-vlan").GetDiag()
			}
		}
	case "*vpcv1.BareMetalServerNetworkInterfaceByHiperSocket":
		{
			nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByHiperSocket)
			d.SetId(*nic.ID)

			if err = d.Set(isBareMetalServerNicAllowIPSpoofing, *nic.AllowIPSpoofing); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allow_ip_spoofing: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-allow_ip_spoofing").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicEnableInfraNAT, *nic.EnableInfrastructureNat); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting enable_infrastructure_nat: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-enable_infrastructure_nat").GetDiag()
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
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting floating_ips: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-floating_ips").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicHref, *nic.Href); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-href").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicID, *nic.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting network_interface: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-network_interface").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicInterfaceType, *nic.InterfaceType); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting interface_type: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-interface_type").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicMacAddress, *nic.MacAddress); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting mac_address: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-mac_address").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicName, *nic.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-name").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicPortSpeed, *nic.PortSpeed); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting port_speed: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-port_speed").GetDiag()
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
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting primary_ip: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-primary_ip").GetDiag()
			}

			if err = d.Set(isBareMetalServerNicResourceType, *nic.ResourceType); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-resource_type").GetDiag()
			}
			if len(nic.SecurityGroups) != 0 {
				secgrpList := []string{}
				for i := 0; i < len(nic.SecurityGroups); i++ {
					secgrpList = append(secgrpList, string(*(nic.SecurityGroups[i].ID)))
				}
				if err = d.Set(isBareMetalServerNicSecurityGroups, flex.NewStringSet(schema.HashString, secgrpList)); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting security_groups: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-security_groups").GetDiag()
				}
			}

			if err = d.Set(isBareMetalServerNicStatus, *nic.Status); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-status").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicSubnet, *nic.Subnet.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting subnet: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-subnet").GetDiag()
			}
			if err = d.Set(isBareMetalServerNicType, *nic.Type); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_is_bare_metal_server_network_interface", "read", "set-type").GetDiag()
			}
		}
	}
	return nil
}
