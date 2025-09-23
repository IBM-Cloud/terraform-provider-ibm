// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isSgName          = "name"
	isSgRules         = "rules"
	isSgRuleID        = "rule_id"
	isSgRuleDirection = "direction"
	isSgRuleIPVersion = "ip_version"
	isSgRuleRemote    = "remote"
	isSgRuleLocal     = "local"
	isSgRuleType      = "type"
	isSgRuleCode      = "code"
	isSgRulePortMax   = "port_max"
	isSgRulePortMin   = "port_min"
	isSgRuleProtocol  = "protocol"
	isSgVPC           = "vpc"
	isSgVPCName       = "vpc_name"
	isSgTags          = "tags"
	isSgCRN           = "crn"
)

func DataSourceIBMISSecurityGroup() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceIBMISSecurityGroupRuleRead,

		Schema: map[string]*schema.Schema{

			isSgName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Security group name",
			},

			isSecurityGroupVPC: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Security group's vpc id",
			},
			isSgVPCName: {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{isSecurityGroupVPC},
				Description:   "Security group's vpc name",
			},
			isSecurityGroupResourceGroup: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Security group's resource group id",
			},

			isSgRules: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Security Rules",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						isSgRuleID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rule id",
						},

						isSgRuleDirection: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Direction of traffic to enforce, either inbound or outbound",
						},

						isSgRuleIPVersion: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP version: ipv4",
						},

						isSgRuleRemote: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Security group id: an IP address, a CIDR block, or a single security group identifier",
						},

						"local": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The local IP address or range of local IP addresses to which this rule will allow inbound traffic (or from which, for outbound traffic). A CIDR block of 0.0.0.0/0 allows traffic to all local IP addresses (or from all local IP addresses, for outbound rules).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.",
									},
									"cidr_block": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CIDR block. This property may add support for IPv6 CIDR blocks in the future. When processing a value in this property, verify that the CIDR block is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected CIDR block format was encountered.",
									},
								},
							},
						},

						isSgRuleType: {
							Type:     schema.TypeInt,
							Computed: true,
						},

						isSgRuleCode: {
							Type:     schema.TypeInt,
							Computed: true,
						},

						isSgRulePortMin: {
							Type:     schema.TypeInt,
							Computed: true,
						},

						isSgRulePortMax: {
							Type:     schema.TypeInt,
							Computed: true,
						},

						isSgRuleProtocol: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},

			isSgTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags",
			},

			isSecurityGroupAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},

			isSgCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},
		},
	}
}

func dataSourceIBMISSecurityGroupRuleRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sgName := d.Get(isSgName).(string)
	vpcId := ""
	vpcName := ""
	rgId := ""
	if vpcIdOk, ok := d.GetOk(isSgVPC); ok {
		vpcId = vpcIdOk.(string)
	}
	if rgIdOk, ok := d.GetOk(isSecurityGroupResourceGroup); ok {
		rgId = rgIdOk.(string)
	}
	if vpcNameOk, ok := d.GetOk(isSgVPCName); ok {
		vpcName = vpcNameOk.(string)
	}
	err := securityGroupGet(context, d, meta, sgName, vpcId, vpcName, rgId)
	if err != nil {
		return err
	}
	return nil
}

func securityGroupGet(context context.Context, d *schema.ResourceData, meta interface{}, name, vpcId, vpcName, rgId string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_security_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Support for pagination
	start := ""
	allrecs := []vpcv1.SecurityGroup{}

	listSgOptions := &vpcv1.ListSecurityGroupsOptions{}
	if vpcId != "" {
		listSgOptions.VPCID = &vpcId
	}
	if vpcName != "" {
		listSgOptions.VPCName = &vpcName
	}
	if rgId != "" {
		listSgOptions.ResourceGroupID = &rgId
	}
	for {
		if start != "" {
			listSgOptions.Start = &start
		}
		sgs, _, err := sess.ListSecurityGroupsWithContext(context, listSgOptions)
		if err != nil || sgs == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListSecurityGroupsWithContext failed: %s", err.Error()), "(Data) ibm_is_security_group", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		if *sgs.TotalCount == int64(0) {
			break
		}
		start = flex.GetNext(sgs.Next)
		allrecs = append(allrecs, sgs.SecurityGroups...)

		if start == "" {
			break
		}

	}

	for _, securityGroup := range allrecs {
		if *securityGroup.Name == name {

			if err = d.Set("name", securityGroup.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_security_group", "read", "set-name").GetDiag()
			}
			if err = d.Set("vpc", securityGroup.VPC.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc: %s", err), "(Data) ibm_is_security_group", "read", "set-vpc").GetDiag()
			}
			if err = d.Set("vpc_name", securityGroup.VPC.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc_name: %s", err), "(Data) ibm_is_security_group", "read", "set-vpc_name").GetDiag()
			}
			if err = d.Set("resource_group", securityGroup.ResourceGroup.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_security_group", "read", "set-resource_group").GetDiag()
			}
			if err = d.Set("crn", securityGroup.CRN); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_security_group", "read", "set-crn").GetDiag()
			}
			tags, err := flex.GetGlobalTagsUsingCRN(meta, *securityGroup.CRN, "", isUserTagType)
			if err != nil {
				log.Printf(
					"An error occured during reading of security group (%s) tags : %s", *securityGroup.ID, err)
			}
			if err = d.Set(isSgTags, tags); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_is_security_group", "read", "set-tags").GetDiag()
			}
			accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *securityGroup.CRN, "", isAccessTagType)
			if err != nil {
				log.Printf(
					"Error on get of security group (%s) access tags: %s", d.Id(), err)
			}
			if err = d.Set(isSecurityGroupAccessTags, accesstags); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_security_group", "read", "set-access_tags").GetDiag()
			}
			rules := make([]map[string]interface{}, 0)
			for _, sgrule := range securityGroup.Rules {
				switch reflect.TypeOf(sgrule).String() {
				case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp":
					{
						rule := sgrule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp)
						r := make(map[string]interface{})
						if rule.Code != nil {
							r[isSgRuleCode] = int(*rule.Code)
						}
						if rule.Type != nil {
							r[isSgRuleType] = int(*rule.Type)
						}
						r[isSgRuleDirection] = *rule.Direction
						r[isSgRuleIPVersion] = *rule.IPVersion
						if rule.Protocol != nil {
							r[isSgRuleProtocol] = *rule.Protocol
						}
						r[isSgRuleID] = *rule.ID
						remote, ok := rule.Remote.(*vpcv1.SecurityGroupRuleRemote)
						if ok {
							if remote != nil && reflect.ValueOf(remote).IsNil() == false {
								if remote.ID != nil {
									r[isSgRuleRemote] = remote.ID
								} else if remote.Address != nil {
									r[isSgRuleRemote] = remote.Address
								} else if remote.CIDRBlock != nil {
									r[isSgRuleRemote] = remote.CIDRBlock
								}
							}
						}
						local, ok := rule.Local.(*vpcv1.SecurityGroupRuleLocal)
						if ok {
							if local != nil && !reflect.ValueOf(local).IsNil() {
								localList := []map[string]interface{}{}
								localMap := dataSourceSecurityGroupRuleLocalToMap(local)
								localList = append(localList, localMap)
								r["local"] = localList
							}
						}
						rules = append(rules, r)
					}

				case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll":
					{
						rule := sgrule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll)
						r := make(map[string]interface{})
						r[isSgRuleDirection] = *rule.Direction
						r[isSgRuleIPVersion] = *rule.IPVersion
						if rule.Protocol != nil {
							r[isSgRuleProtocol] = *rule.Protocol
						}
						r[isSgRuleID] = *rule.ID
						remote, ok := rule.Remote.(*vpcv1.SecurityGroupRuleRemote)
						if ok {
							if remote != nil && reflect.ValueOf(remote).IsNil() == false {
								if remote.ID != nil {
									r[isSgRuleRemote] = remote.ID
								} else if remote.Address != nil {
									r[isSgRuleRemote] = remote.Address
								} else if remote.CIDRBlock != nil {
									r[isSgRuleRemote] = remote.CIDRBlock
								}
							}
						}
						local, ok := rule.Local.(*vpcv1.SecurityGroupRuleLocal)
						if ok {
							if local != nil && !reflect.ValueOf(local).IsNil() {
								localList := []map[string]interface{}{}
								localMap := dataSourceSecurityGroupRuleLocalToMap(local)
								localList = append(localList, localMap)
								r["local"] = localList
							}
						}
						rules = append(rules, r)
					}

				case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp":
					{
						rule := sgrule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp)
						r := make(map[string]interface{})
						if rule.PortMin != nil {
							r[isSgRulePortMin] = int(*rule.PortMin)
						}
						if rule.PortMax != nil {
							r[isSgRulePortMax] = int(*rule.PortMax)
						}
						r[isSgRuleDirection] = *rule.Direction
						r[isSgRuleIPVersion] = *rule.IPVersion
						if rule.Protocol != nil {
							r[isSgRuleProtocol] = *rule.Protocol
						}
						remote, ok := rule.Remote.(*vpcv1.SecurityGroupRuleRemote)
						if ok {
							if remote != nil && reflect.ValueOf(remote).IsNil() == false {
								if remote.ID != nil {
									r[isSgRuleRemote] = remote.ID
								} else if remote.Address != nil {
									r[isSgRuleRemote] = remote.Address
								} else if remote.CIDRBlock != nil {
									r[isSgRuleRemote] = remote.CIDRBlock
								}
							}
						}
						local, ok := rule.Local.(*vpcv1.SecurityGroupRuleLocal)
						if ok {
							if local != nil && !reflect.ValueOf(local).IsNil() {
								localList := []map[string]interface{}{}
								localMap := dataSourceSecurityGroupRuleLocalToMap(local)
								localList = append(localList, localMap)
								r["local"] = localList
							}
						}
						rules = append(rules, r)
					}
				}
			}
			if err = d.Set(isSgRules, rules); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting rules: %s", err), "(Data) ibm_is_security_group", "read", "set-rules").GetDiag()
			}
			d.SetId(*securityGroup.ID)

			if securityGroup.ResourceGroup != nil {
				if securityGroup.ResourceGroup.Name != nil {
					if err = d.Set(flex.ResourceGroupName, securityGroup.ResourceGroup.Name); err != nil {
						return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group_name: %s", err), "(Data) ibm_is_security_group", "read", "set-resource_group_name").GetDiag()
					}
				}
			}

			controller, err := flex.GetBaseController(meta)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "(Data) ibm_is_security_group", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			if err = d.Set(flex.ResourceControllerURL, controller+"/vpc/network/securityGroups"); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_controller_url: %s", err), "(Data) ibm_is_security_group", "read", "set-resource_controller_url").GetDiag()
			}
			if securityGroup.Name != nil {
				if err = d.Set(flex.ResourceName, securityGroup.Name); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_name: %s", err), "(Data) ibm_is_security_group", "read", "set-resource_name").GetDiag()
				}
			}

			if securityGroup.CRN != nil {
				if err = d.Set(flex.ResourceCRN, securityGroup.CRN); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_crn: %s", err), "(Data) ibm_is_security_group", "read", "set-resource_crn").GetDiag()
				}
			}
			return nil
		}
	}
	err = fmt.Errorf("[ERROR] No Security Group found with name %s", name)
	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListSecurityGroupsWithContext failed: %s", err.Error()), "(Data) ibm_is_security_group", "read")
	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	return tfErr.GetDiag()

}
