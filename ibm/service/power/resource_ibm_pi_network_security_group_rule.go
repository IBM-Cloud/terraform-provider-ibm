// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"time"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceIBMPINetworkSecurityGroupMember() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPINetworkSecurityGroupMemberCreate,
		ReadContext:   resourceIBMPINetworkSecurityGroupMemberRead,
		DeleteContext: resourceIBMPINetworkSecurityGroupMemberDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_Action: {
				Description:  "The action to take if the rule matches network traffic.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"allow", "deny"}),
			},
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_DestinationPorts: {
				Computed:    true,
				Description: "The list of destination port.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Arg_EndingIPAddress: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The end of the port range, if applicable. If values are not present then all ports are in the range.",
						},
						Arg_StartingIPAddress: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The start of the port range, if applicable. If values are not present then all ports are in the range.",
						},
					},
				},
				Optional: true,
				Type:     schema.TypeList,
			},
			Arg_NetworkSecurityGroupID: {
				Description: "The unique identifier of the network security group.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
			Arg_NetworkSecurityGroupMemberID: {
				Description: "The network security group member id to remove.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			Arg_Protocol: {
				Description: "The protocol of the network traffic.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Arg_ICMPTypes: {
							Description: "If icmp type, the list of ICMP packet types (by numbers) affected by ICMP rules and if not present then all types are matched.",
							Elem:        &schema.Schema{Type: schema.TypeFloat},
							Required:    true,
							Type:        schema.TypeList,
						},
						Arg_TCPFlags: {
							Description: "If tcp type is chosen, the list of TCP flags and if not present then all flags are matched.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									Attr_Flag: {
										Computed:    true,
										Description: "TCP flag.",
										Type:        schema.TypeString,
									},
								},
							},
							Required: true,
							Type:     schema.TypeList,
						},
						Arg_Type: {
							Description:  "The protocol of the network traffic.",
							Required:     true,
							Type:         schema.TypeString,
							ValidateFunc: validate.ValidateAllowedStringValues([]string{"all", "icmp", "tcp", "udp"}),
						},
					},
				},
				Required: true,
				Type:     schema.TypeList,
			},
			Arg_Remote: {
				Description: "The protocol of the network traffic.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Arg_ID: {
							Description: "The id of the network security group rule to be removed.",
							Required:    true,
							Type:        schema.TypeString,
						},
						Arg_Type: {
							Description:  "The type of remote group (MAC addresses, IP addresses, CIDRs, external CIDRs) that are the originators of rule's network traffic to match.",
							Required:     true,
							Type:         schema.TypeString,
							ValidateFunc: validate.ValidateAllowedStringValues([]string{"default-network-address-group", "network-address-group", "network-security-group"}),
						},
					},
				},
				Required: true,
				Type:     schema.TypeSet,
			},
			Arg_RuleName: {
				Description: "The unique name of the network security group rule to be added.",
				Required:    true,
				Type:        schema.TypeString,
			},
			Arg_SourcePorts: {
				Computed:    true,
				Description: "Source port ranges.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Arg_EndingIPAddress: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Ending ip address",
						},
						Arg_StartingIPAddress: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Starting ip address",
						},
					},
				},
				Optional: true,
				Type:     schema.TypeList,
			},

			// Attributes
			Attr_CRN: {
				Computed:    true,
				Description: "The network security group's crn.",
				Type:        schema.TypeString,
			},
			Attr_Members: {
				Computed:    true,
				Description: "The list of IPv4 addresses and, or network interfaces in the network security group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_ID: {
							Computed:    true,
							Description: "The ID of the member in a network security group.",
							Type:        schema.TypeString,
						},
						Attr_MacAddress: {
							Computed:    true,
							Description: "The mac address of a network interface included if the type is network-interface.",
							Type:        schema.TypeString,
						},
						Attr_Target: {
							Computed:    true,
							Description: "If ipv4-address type, then IPv4 address or if network-interface type, then network interface ID.",
							Type:        schema.TypeString,
						},
						Attr_Type: {
							Computed:    true,
							Description: "The type of member.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
			Attr_Name: {
				Computed:    true,
				Description: "The name of the network security group.",
				Type:        schema.TypeString,
			},
			Attr_Rules: {
				Computed:    true,
				Description: "The list of rules in the network security group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_Action: {
							Computed:    true,
							Description: "The action to take if the rule matches network traffic.",
							Type:        schema.TypeString,
						},
						Attr_DestinationPort: {
							Computed:    true,
							Description: "The list of destination port.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									Attr_Maximum: {
										Computed:    true,
										Description: "The end of the port range, if applicable. If values are not present then all ports are in the range.",
										Type:        schema.TypeFloat,
									},
									Attr_Minimum: {
										Computed:    true,
										Description: "The start of the port range, if applicable. If values are not present then all ports are in the range.",
										Type:        schema.TypeFloat,
									},
								},
							},
							Type: schema.TypeList,
						},
						Attr_ID: {
							Computed:    true,
							Description: "The ID of the rule in a network security group.",
							Type:        schema.TypeString,
						},
						Attr_Name: {
							Computed:    true,
							Description: "The unique name of the network security group rule.",
							Type:        schema.TypeString,
						},
						Attr_Protocol: {
							Computed:    true,
							Description: "The list of protocol.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									Attr_ICMPTypes: {
										Computed:    true,
										Description: "If icmp type, the list of ICMP packet types (by numbers) affected by ICMP rules and if not present then all types are matched.",
										Elem:        &schema.Schema{Type: schema.TypeFloat},
										Type:        schema.TypeList,
									},
									Attr_TCPFlags: {
										Computed:    true,
										Description: "If tcp type, the list of TCP flags and if not present then all flags are matched.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												Attr_Flag: {
													Computed:    true,
													Description: "TCP flag.",
													Type:        schema.TypeString,
												},
											},
										},
										Type: schema.TypeList,
									},
									Attr_Type: {
										Computed:    true,
										Description: "The protocol of the network traffic.",
										Type:        schema.TypeString,
									},
								},
							},
							Type: schema.TypeList,
						},
						Attr_Remote: {
							Computed:    true,
							Description: "List of remote.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									Attr_ID: {
										Computed:    true,
										Description: "The ID of the remote Network Address Group or network security group the rules apply to. Not required for default-network-address-group.",
										Type:        schema.TypeString,
									},
									Attr_Type: {
										Computed:    true,
										Description: "The type of remote group the rules apply to.",
										Type:        schema.TypeString,
									},
								},
							},
							Type: schema.TypeList,
						},
						Attr_SourcePort: {
							Computed:    true,
							Description: "ist of source port",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									Attr_Maximum: {
										Computed:    true,
										Description: "The end of the port range, if applicable, If values are not present then all ports are in the range.",
										Type:        schema.TypeFloat,
									},
									Attr_Minimum: {
										Computed:    true,
										Description: "The start of the port range, if applicable. If values are not present then all ports are in the range.",
										Type:        schema.TypeFloat,
									},
								},
							},
							Type: schema.TypeList,
						},
					},
				},
				Type: schema.TypeList,
			},
			Attr_UserTags: {
				Computed:    true,
				Description: "The user tags associated with this resource.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeList,
			},
		},
	}
}

func resourceIBMPINetworkSecurityGroupMemberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	action := d.Get(Arg_Action).(string)
	name := d.Get(Arg_Name).(string)

	nsgClient := instance.NewIBMIPINetworkSecurityGroupClient(ctx, sess, cloudInstanceID)

	networkSecurityGroupAddRule := models.NetworkSecurityGroupAddRule{
		Action: &action,
		Name:   &name,
	}

	// Add protocol
	protocol := d.Get(Arg_Protocol).([]interface{})
	networkSecurityGroupRuleProtocol := models.NetworkSecurityGroupRuleProtocol{
		Type: protocolFlag,
	}
	if protocolFlag == "icmp" {
		networkSecurityGroupRuleProtocol.IcmpTypes = icmpTypesFlag
	} else if protocolFlag == "tcp" {
		networkSecurityGroupRuleProtocolTCPFlagArray := []*models.NetworkSecurityGroupRuleProtocolTCPFlag{}
		for _, tcp := range tcpFlag {
			networkSecurityGroupRuleProtocolTCPFlag := models.NetworkSecurityGroupRuleProtocolTCPFlag{}
			networkSecurityGroupRuleProtocolTCPFlag.Flag = tcp
			networkSecurityGroupRuleProtocolTCPFlagArray = append(networkSecurityGroupRuleProtocolTCPFlagArray, &networkSecurityGroupRuleProtocolTCPFlag)
		}
		networkSecurityGroupRuleProtocol.TCPFlags = networkSecurityGroupRuleProtocolTCPFlagArray
	}
	networkSecurityGroupAddRule.Protocol = &networkSecurityGroupRuleProtocol

	// Add remote
	networkSecurityGroupRuleRemote := models.NetworkSecurityGroupRuleRemote{
		ID:   remoteGroupIDFlag,
		Type: remoteTypeFlag,
	}
	networkSecurityGroupAddRule.Remote = &networkSecurityGroupRuleRemote

	// Optional Fields
	if destinationPortsFlag != "" {
		destinationPortRanges, err := convertPortRange(destinationPortsFlag, cmd.UI)
		if err != nil {
			return err
		} else {
			networkSecurityGroupAddRule.DestinationPorts = destinationPortRanges
		}
	}
	if sourcePortsFlag != "" {
		sourcePortRanges, err := convertPortRange(sourcePortsFlag, cmd.UI)
		if err != nil {
			return err
		} else {
			networkSecurityGroupAddRule.SourcePorts = sourcePortRanges
		}
	}

	body := &models.NetworkSecurityGroupCreate{
		Name: &name,
	}
	if v, ok := d.GetOk(Arg_UserTags); ok {
		var userTags []string
		for _, v := range v.([]interface{}) {
			userTag := v.(string)
			userTags = append(userTags, userTag)
		}
		body.UserTags = userTags
	}

	networkSecurityGroup, err := nsgClient.Create(body)
	if err != nil {
		return diag.FromErr(err)
	}
	nsgID := *networkSecurityGroup.ID
	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, nsgID))

	return resourceIBMPINetworkSecurityGroupMemberRead(ctx, d, meta)
}

func resourceIBMPINetworkSecurityGroupMemberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID, nsgID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	nsgClient := instance.NewIBMIPINetworkSecurityGroupClient(ctx, sess, cloudInstanceID)
	networkSecurityGroup, err := nsgClient.Get(nsgID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set(Attr_CRN, networkSecurityGroup.Crn)

	if networkSecurityGroup.Members != nil {
		members := []map[string]interface{}{}
		for _, mbr := range networkSecurityGroup.Members {
			mbrMap := networkSecurityGroupMemberToMap(mbr)
			members = append(members, mbrMap)
		}
		d.Set(Attr_Members, members)
	}
	d.Set(Arg_Name, networkSecurityGroup.Name)
	d.Set(Attr_NetworkSecurityGroupID, networkSecurityGroup.ID)
	if networkSecurityGroup.Rules != nil {
		rules := []map[string]interface{}{}
		for _, rule := range networkSecurityGroup.Rules {
			ruleMap := networkSecurityGroupRuleToMap(rule)
			rules = append(rules, ruleMap)
		}
		d.Set(Attr_Rules, rules)
	}

	return nil
}

func resourceIBMPINetworkSecurityGroupMemberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func isWaitForIBMPINetworkSecurityGroupMemberAdd(ctx context.Context, client *instance.IBMPINetworkSecurityGroupClient, id, memberID string, timeout time.Duration) (interface{}, error) {

	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Pending},
		Target:     []string{State_Available},
		Refresh:    isIBMPINetworkSecurityGroupMemberAddRefreshFunc(client, id, memberID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Minute,
	}

	return stateConf.WaitForStateContext(ctx)
}
func isIBMPINetworkSecurityGroupMemberAddRefreshFunc(client *instance.IBMPINetworkSecurityGroupClient, id, memberID string) retry.StateRefreshFunc {

	return func() (interface{}, string, error) {
		networkSecurityGroup, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if networkSecurityGroup.Members != nil {
			for _, mbr := range networkSecurityGroup.Members {
				if *mbr.ID == memberID {
					return networkSecurityGroup, State_Available, nil
				}

			}
		}
		return networkSecurityGroup, State_Pending, nil
	}
}

func isWaitForIBMPINetworkSecurityGroupMemberRemove(ctx context.Context, client *instance.IBMPINetworkSecurityGroupClient, id, memberID string, timeout time.Duration) (interface{}, error) {

	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Pending},
		Target:     []string{State_Removed},
		Refresh:    isIBMPINetworkSecurityGroupMemberRemoveRefreshFunc(client, id, memberID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Minute,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isIBMPINetworkSecurityGroupMemberRemoveRefreshFunc(client *instance.IBMPINetworkSecurityGroupClient, id, memberID string) retry.StateRefreshFunc {

	return func() (interface{}, string, error) {
		networkSecurityGroup, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if networkSecurityGroup.Members != nil {
			isMember := false
			for _, mbr := range networkSecurityGroup.Members {
				if *mbr.ID == memberID {
					isMember = true
					return networkSecurityGroup, State_Pending, nil
				}
			}
			if !isMember {
				return networkSecurityGroup, State_Removed, nil
			}
		}
		return networkSecurityGroup, State_Pending, nil
	}
}

func networkSecurityGroupMemberToMap(mbr *models.NetworkSecurityGroupMember) map[string]interface{} {
	mbrMap := make(map[string]interface{})
	mbrMap[Attr_ID] = mbr.ID
	if mbr.MacAddress != "" {
		mbrMap[Attr_MacAddress] = mbr.MacAddress
	}
	mbrMap[Attr_Target] = mbr.Target
	mbrMap[Attr_Type] = mbr.Type
	return mbrMap
}

func networkSecurityGroupRuleToMap(rule *models.NetworkSecurityGroupRule) map[string]interface{} {
	ruleMap := make(map[string]interface{})
	ruleMap[Attr_Action] = rule.Action
	if rule.DestinationPort != nil {
		destinationPortMap := networkSecurityGroupRulePortToMap(rule.DestinationPort)
		ruleMap[Attr_DestinationPort] = []map[string]interface{}{destinationPortMap}
	}

	ruleMap[Attr_ID] = rule.ID
	ruleMap[Attr_Name] = rule.Name

	protocolMap := networkSecurityGroupRuleProtocolToMap(rule.Protocol)
	ruleMap[Attr_Protocol] = []map[string]interface{}{protocolMap}

	remoteMap := networkSecurityGroupRuleRemoteToMap(rule.Remote)

	ruleMap[Attr_Remote] = []map[string]interface{}{remoteMap}

	if rule.SourcePort != nil {
		sourcePortMap := networkSecurityGroupRulePortToMap(rule.SourcePort)
		ruleMap[Attr_SourcePort] = []map[string]interface{}{sourcePortMap}
	}

	return ruleMap
}

func networkSecurityGroupRulePortToMap(port *models.NetworkSecurityGroupRulePort) map[string]interface{} {
	portMap := make(map[string]interface{})
	portMap[Attr_Maximum] = port.Maximum
	portMap[Attr_Minimum] = port.Minimum
	return portMap
}

func networkSecurityGroupRuleProtocolToMap(protocol *models.NetworkSecurityGroupRuleProtocol) map[string]interface{} {
	protocolMap := make(map[string]interface{})
	if protocol.IcmpTypes != nil {
		protocolMap[Attr_ICMPTypes] = protocol.IcmpTypes
	}
	if protocol.TCPFlags != nil {
		tcpFlags := []map[string]interface{}{}
		for _, tcpFlagsItem := range protocol.TCPFlags {
			tcpFlagsItemMap := make(map[string]interface{})
			tcpFlagsItemMap[Attr_Flag] = tcpFlagsItem.Flag
			tcpFlags = append(tcpFlags, tcpFlagsItemMap)
		}
		protocolMap[Attr_TCPFlags] = tcpFlags
	}
	if protocol.Type != "" {
		protocolMap[Attr_Type] = protocol.Type
	}
	return protocolMap
}

func networkSecurityGroupRuleRemoteToMap(remote *models.NetworkSecurityGroupRuleRemote) map[string]interface{} {
	remoteMap := make(map[string]interface{})
	if remote.ID != "" {
		remoteMap[Attr_ID] = remote.ID
	}
	if remote.Type != "" {
		remoteMap[Attr_Type] = remote.Type
	}
	return remoteMap
}
