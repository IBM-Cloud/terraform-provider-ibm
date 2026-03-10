// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isNetworkACLRuleHref = "href"
)

func DataSourceIBMISNetworkACLRule() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISNetworkACLRuleRead,

		Schema: map[string]*schema.Schema{
			isNwACLID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Network ACL id",
			},
			isNwACLRuleId: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Network ACL rule id",
			},
			isNwACLRuleBefore: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The rule that this rule is immediately before. If absent, this is the last rule.",
			},
			isNetworkACLRuleName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_network_acl", isNetworkACLRuleName),
				Description:  "The user-defined name for this rule",
			},
			isNetworkACLRuleProtocol: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the network protocol",
			},
			isNetworkACLRuleHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this network ACL rule",
			},
			isNetworkACLRuleAction: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Whether to allow or deny matching traffic.",
			},
			isNetworkACLRuleIPVersion: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP version for this rule.",
			},
			isNetworkACLRuleSource: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The source IP address or CIDR block.",
			},
			isNetworkACLRuleDestination: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The destination IP address or CIDR block.",
			},
			isNetworkACLRuleDirection: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Whether the traffic to be matched is inbound or outbound.",
			},
			isNetworkACLRuleICMP: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The protocol ICMP",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isNetworkACLRuleICMPCode: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ICMP traffic code to allow. Valid values from 0 to 255.",
						},
						isNetworkACLRuleICMPType: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ICMP traffic type to allow. Valid values from 0 to 254.",
						},
					},
				},
			},

			isNetworkACLRuleTCP: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "TCP protocol",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isNetworkACLRulePortMax: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The highest port in the range of ports to be matched",
						},
						isNetworkACLRulePortMin: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The lowest port in the range of ports to be matched",
						},
						isNetworkACLRuleSourcePortMax: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The highest port in the range of ports to be matched",
						},
						isNetworkACLRuleSourcePortMin: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The lowest port in the range of ports to be matched",
						},
					},
				},
			},

			isNetworkACLRuleUDP: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "UDP protocol",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isNetworkACLRulePortMax: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The highest port in the range of ports to be matched",
						},
						isNetworkACLRulePortMin: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The lowest port in the range of ports to be matched",
						},
						isNetworkACLRuleSourcePortMax: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The highest port in the range of ports to be matched",
						},
						isNetworkACLRuleSourcePortMin: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The lowest port in the range of ports to be matched",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMISNetworkACLRuleRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	nwACLID := d.Get(isNwACLID).(string)
	name := d.Get(isNetworkACLRuleName).(string)
	err := nawaclRuleDataGet(context, d, meta, name, nwACLID)
	if err != nil {
		return err
	}

	return nil
}

func nawaclRuleDataGet(context context.Context, d *schema.ResourceData, meta interface{}, name, nwACLID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_network_acl_rule", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	start := ""
	allrecs := []vpcv1.NetworkACLRuleItemIntf{}
	for {
		listNetworkACLRulesOptions := &vpcv1.ListNetworkACLRulesOptions{
			NetworkACLID: &nwACLID,
		}
		if start != "" {
			listNetworkACLRulesOptions.Start = &start
		}

		ruleList, _, err := sess.ListNetworkACLRulesWithContext(context, listNetworkACLRulesOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListNetworkACLRulesWithContext failed: %s", err.Error()), "(Data) ibm_is_network_acl_rule", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		start = flex.GetNext(ruleList.Next)

		allrecs = append(allrecs, ruleList.Rules...)
		if start == "" {
			break
		}
	}

	for _, rule := range allrecs {
		switch reflect.TypeOf(rule).String() {
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmp":
			{
				networkACLRule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmp)
				if *networkACLRule.Name == name {
					d.SetId(makeTerraformACLRuleID(nwACLID, *networkACLRule.ID))
					d.Set(isNwACLRuleId, *networkACLRule.ID)
					if networkACLRule.Before != nil {
						if err = d.Set("before", *networkACLRule.Before.ID); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting before: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-before").GetDiag()
						}
					}
					if err = d.Set("href", networkACLRule.Href); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-href").GetDiag()
					}
					if err = d.Set("protocol", networkACLRule.Protocol); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting protocol: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-protocol").GetDiag()
					}
					if err = d.Set("name", networkACLRule.Name); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-name").GetDiag()
					}
					if err = d.Set("action", networkACLRule.Action); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-action").GetDiag()
					}
					if err = d.Set("ip_version", networkACLRule.IPVersion); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting ip_version: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-ip_version").GetDiag()
					}
					if err = d.Set("source", networkACLRule.Source); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-source").GetDiag()
					}
					if err = d.Set("destination", networkACLRule.Destination); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting destination: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-destination").GetDiag()
					}
					if err = d.Set("direction", networkACLRule.Direction); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting direction: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-direction").GetDiag()
					}
					d.Set(isNetworkACLRuleTCP, make([]map[string]int, 0, 0))
					d.Set(isNetworkACLRuleUDP, make([]map[string]int, 0, 0))
					icmp := make([]map[string]int, 1, 1)
					if networkACLRule.Code != nil && networkACLRule.Type != nil {
						icmp[0] = map[string]int{
							isNetworkACLRuleICMPCode: int(*networkACLRule.Code),
							isNetworkACLRuleICMPType: int(*networkACLRule.Type),
						}
					}
					if err = d.Set(isNetworkACLRuleICMP, icmp); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting icmp: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-icmp").GetDiag()
					}
					break
				}
			}
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp":
			{
				networkACLRule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp)
				if *networkACLRule.Name == name {
					d.SetId(makeTerraformACLRuleID(nwACLID, *networkACLRule.ID))
					d.Set(isNwACLRuleId, *networkACLRule.ID)
					if networkACLRule.Before != nil {
						if err = d.Set("before", *networkACLRule.Before.ID); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting before: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-before").GetDiag()
						}
					}
					if err = d.Set("href", networkACLRule.Href); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-href").GetDiag()
					}
					if err = d.Set("protocol", networkACLRule.Protocol); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting protocol: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-protocol").GetDiag()
					}
					if err = d.Set("name", networkACLRule.Name); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-name").GetDiag()
					}
					if err = d.Set("action", networkACLRule.Action); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-action").GetDiag()
					}
					if err = d.Set("ip_version", networkACLRule.IPVersion); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting ip_version: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-ip_version").GetDiag()
					}
					if err = d.Set("source", networkACLRule.Source); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-source").GetDiag()
					}
					if err = d.Set("destination", networkACLRule.Destination); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting destination: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-destination").GetDiag()
					}
					if err = d.Set("direction", networkACLRule.Direction); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting direction: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-direction").GetDiag()
					}
					if *networkACLRule.Protocol == "tcp" {
						d.Set(isNetworkACLRuleICMP, make([]map[string]int, 0, 0))
						d.Set(isNetworkACLRuleUDP, make([]map[string]int, 0, 0))
						tcp := make([]map[string]int, 1, 1)
						tcp[0] = map[string]int{
							isNetworkACLRuleSourcePortMax: checkNetworkACLNil(networkACLRule.SourcePortMax),
							isNetworkACLRuleSourcePortMin: checkNetworkACLNil(networkACLRule.SourcePortMin),
						}
						tcp[0][isNetworkACLRulePortMax] = checkNetworkACLNil(networkACLRule.DestinationPortMax)
						tcp[0][isNetworkACLRulePortMin] = checkNetworkACLNil(networkACLRule.DestinationPortMin)
						if err = d.Set(isNetworkACLRuleTCP, tcp); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tcp: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-tcp").GetDiag()
						}
					} else if *networkACLRule.Protocol == "udp" {
						d.Set(isNetworkACLRuleICMP, make([]map[string]int, 0, 0))
						d.Set(isNetworkACLRuleTCP, make([]map[string]int, 0, 0))
						udp := make([]map[string]int, 1, 1)
						udp[0] = map[string]int{
							isNetworkACLRuleSourcePortMax: checkNetworkACLNil(networkACLRule.SourcePortMax),
							isNetworkACLRuleSourcePortMin: checkNetworkACLNil(networkACLRule.SourcePortMin),
						}
						udp[0][isNetworkACLRulePortMax] = checkNetworkACLNil(networkACLRule.DestinationPortMax)
						udp[0][isNetworkACLRulePortMin] = checkNetworkACLNil(networkACLRule.DestinationPortMin)
						if err = d.Set(isNetworkACLRuleUDP, udp); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting udp: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-udp").GetDiag()
						}
						break
					}
				}
			}
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolAny":
			{
				networkACLRule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolAny)
				if *networkACLRule.Name == name {
					d.SetId(makeTerraformACLRuleID(nwACLID, *networkACLRule.ID))
					d.Set(isNwACLRuleId, *networkACLRule.ID)
					if networkACLRule.Before != nil {
						if err = d.Set("before", *networkACLRule.Before.ID); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting before: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-before").GetDiag()
						}
					}
					if err = d.Set("href", networkACLRule.Href); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-href").GetDiag()
					}
					if err = d.Set("protocol", networkACLRule.Protocol); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting protocol: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-protocol").GetDiag()
					}
					if err = d.Set("name", networkACLRule.Name); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-name").GetDiag()
					}
					if err = d.Set("action", networkACLRule.Action); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-action").GetDiag()
					}
					if err = d.Set("ip_version", networkACLRule.IPVersion); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting ip_version: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-ip_version").GetDiag()
					}
					if err = d.Set("source", networkACLRule.Source); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-source").GetDiag()
					}
					if err = d.Set("destination", networkACLRule.Destination); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting destination: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-destination").GetDiag()
					}
					if err = d.Set("direction", networkACLRule.Direction); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting direction: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-direction").GetDiag()
					}
					d.Set(isNetworkACLRuleICMP, make([]map[string]int, 0, 0))
					d.Set(isNetworkACLRuleTCP, make([]map[string]int, 0, 0))
					d.Set(isNetworkACLRuleUDP, make([]map[string]int, 0, 0))
					break
				}
			}
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIndividual":
			{
				networkACLRule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIndividual)
				if *networkACLRule.Name == name {
					d.SetId(makeTerraformACLRuleID(nwACLID, *networkACLRule.ID))
					d.Set(isNwACLRuleId, *networkACLRule.ID)
					if networkACLRule.Before != nil {
						if err = d.Set("before", *networkACLRule.Before.ID); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting before: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-before").GetDiag()
						}
					}
					if err = d.Set("href", networkACLRule.Href); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-href").GetDiag()
					}
					if err = d.Set("protocol", networkACLRule.Protocol); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting protocol: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-protocol").GetDiag()
					}
					if err = d.Set("name", networkACLRule.Name); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-name").GetDiag()
					}
					if err = d.Set("action", networkACLRule.Action); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-action").GetDiag()
					}
					if err = d.Set("ip_version", networkACLRule.IPVersion); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting ip_version: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-ip_version").GetDiag()
					}
					if err = d.Set("source", networkACLRule.Source); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-source").GetDiag()
					}
					if err = d.Set("destination", networkACLRule.Destination); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting destination: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-destination").GetDiag()
					}
					if err = d.Set("direction", networkACLRule.Direction); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting direction: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-direction").GetDiag()
					}
					d.Set(isNetworkACLRuleICMP, make([]map[string]int, 0, 0))
					d.Set(isNetworkACLRuleTCP, make([]map[string]int, 0, 0))
					d.Set(isNetworkACLRuleUDP, make([]map[string]int, 0, 0))
					break
				}
			}
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmptcpudp":
			{
				networkACLRule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmptcpudp)
				if *networkACLRule.Name == name {
					d.SetId(makeTerraformACLRuleID(nwACLID, *networkACLRule.ID))
					d.Set(isNwACLRuleId, *networkACLRule.ID)
					if networkACLRule.Before != nil {
						if err = d.Set("before", *networkACLRule.Before.ID); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting before: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-before").GetDiag()
						}
					}
					if err = d.Set("href", networkACLRule.Href); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-href").GetDiag()
					}
					if err = d.Set("protocol", networkACLRule.Protocol); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting protocol: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-protocol").GetDiag()
					}
					if err = d.Set("name", networkACLRule.Name); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-name").GetDiag()
					}
					if err = d.Set("action", networkACLRule.Action); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-action").GetDiag()
					}
					if err = d.Set("ip_version", networkACLRule.IPVersion); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting ip_version: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-ip_version").GetDiag()
					}
					if err = d.Set("source", networkACLRule.Source); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-source").GetDiag()
					}
					if err = d.Set("destination", networkACLRule.Destination); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting destination: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-destination").GetDiag()
					}
					if err = d.Set("direction", networkACLRule.Direction); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting direction: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-direction").GetDiag()
					}
					d.Set(isNetworkACLRuleICMP, make([]map[string]int, 0, 0))
					d.Set(isNetworkACLRuleTCP, make([]map[string]int, 0, 0))
					d.Set(isNetworkACLRuleUDP, make([]map[string]int, 0, 0))
					break
				}
			}
		case "*vpcv1.NetworkACLRuleItem":
			{
				networkACLRule := rule.(*vpcv1.NetworkACLRuleItem)
				if *networkACLRule.Name == name {
					d.SetId(makeTerraformACLRuleID(nwACLID, *networkACLRule.ID))
					d.Set(isNwACLRuleId, *networkACLRule.ID)
					if networkACLRule.Before != nil {
						if err = d.Set("before", *networkACLRule.Before.ID); err != nil {
							return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting before: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-before").GetDiag()
						}
					}
					if err = d.Set("href", networkACLRule.Href); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-href").GetDiag()
					}
					if err = d.Set("protocol", networkACLRule.Protocol); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting protocol: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-protocol").GetDiag()
					}
					if err = d.Set("name", networkACLRule.Name); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-name").GetDiag()
					}
					if err = d.Set("action", networkACLRule.Action); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-action").GetDiag()
					}
					if err = d.Set("ip_version", networkACLRule.IPVersion); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting ip_version: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-ip_version").GetDiag()
					}
					if err = d.Set("source", networkACLRule.Source); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting source: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-source").GetDiag()
					}
					if err = d.Set("destination", networkACLRule.Destination); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting destination: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-destination").GetDiag()
					}
					if err = d.Set("direction", networkACLRule.Direction); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting direction: %s", err), "(Data) ibm_is_network_acl_rule", "read", "set-direction").GetDiag()
					}
					d.Set(isNetworkACLRuleICMP, make([]map[string]int, 0, 0))
					d.Set(isNetworkACLRuleTCP, make([]map[string]int, 0, 0))
					d.Set(isNetworkACLRuleUDP, make([]map[string]int, 0, 0))
					break
				}
			}
		}
	}
	return nil
}
