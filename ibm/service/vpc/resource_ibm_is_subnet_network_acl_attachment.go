// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isSubnetID     = "subnet"
	isNetworkACLID = "network_acl"
)

func ResourceIBMISSubnetNetworkACLAttachment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISSubnetNetworkACLAttachmentCreate,
		ReadContext:   resourceIBMISSubnetNetworkACLAttachmentRead,
		UpdateContext: resourceIBMISSubnetNetworkACLAttachmentUpdate,
		DeleteContext: resourceIBMISSubnetNetworkACLAttachmentDelete,
		Exists:        resourceIBMISSubnetNetworkACLAttachmentExists,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isSubnetID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet identifier",
			},

			isNetworkACLID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of network ACL",
			},

			isNetworkACLName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Network ACL name",
			},

			isNetworkACLCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn for this Network ACL",
			},

			isNetworkACLVPC: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Network ACL VPC",
			},

			isNetworkACLResourceGroup: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource group ID for the network ACL",
			},

			isNetworkACLRules: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isNetworkACLRuleID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this Network ACL rule",
						},
						isNetworkACLRuleName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this rule",
						},
						isNetworkACLRuleAction: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether to allow or deny matching traffic",
						},
						isNetworkACLRuleIPVersion: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP version for this rule",
						},
						isNetworkACLRuleSource: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source CIDR block",
						},
						isNetworkACLRuleDestination: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination CIDR block",
						},
						isNetworkACLRuleDirection: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Direction of traffic to enforce, either inbound or outbound",
						},
						isNetworkACLRuleICMP: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isNetworkACLRuleICMPCode: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The ICMP traffic code to allow",
									},
									isNetworkACLRuleICMPType: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The ICMP traffic type to allow",
									},
								},
							},
						},

						isNetworkACLRuleTCP: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isNetworkACLRulePortMax: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The inclusive upper bound of TCP destination port range",
									},
									isNetworkACLRulePortMin: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The inclusive lower bound of TCP destination port range",
									},
									isNetworkACLRuleSourcePortMax: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The inclusive upper bound of TCP source port range",
									},
									isNetworkACLRuleSourcePortMin: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The inclusive lower bound of TCP source port range",
									},
								},
							},
						},

						isNetworkACLRuleUDP: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isNetworkACLRulePortMax: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The inclusive upper bound of UDP destination port range",
									},
									isNetworkACLRulePortMin: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The inclusive lower bound of UDP destination port range",
									},
									isNetworkACLRuleSourcePortMax: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The inclusive upper bound of UDP source port range",
									},
									isNetworkACLRuleSourcePortMin: {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The inclusive lower bound of UDP source port range",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceIBMISSubnetNetworkACLAttachmentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_network_acl_attachment", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	subnet := d.Get(isSubnetID).(string)
	networkACL := d.Get(isNetworkACLID).(string)

	// Construct an instance of the NetworkACLIdentityByID model
	networkACLIdentityModel := new(vpcv1.NetworkACLIdentityByID)
	networkACLIdentityModel.ID = &networkACL

	// Construct an instance of the ReplaceSubnetNetworkACLOptions model
	replaceSubnetNetworkACLOptionsModel := new(vpcv1.ReplaceSubnetNetworkACLOptions)
	replaceSubnetNetworkACLOptionsModel.ID = &subnet
	replaceSubnetNetworkACLOptionsModel.NetworkACLIdentity = networkACLIdentityModel
	resultACL, _, err := sess.ReplaceSubnetNetworkACLWithContext(context, replaceSubnetNetworkACLOptionsModel)

	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceSubnetNetworkACLWithContext failed: %s", err.Error()), "ibm_is_subnet_network_acl_attachment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(subnet)
	log.Printf("[INFO] Network ACL : %s", *resultACL.ID)
	log.Printf("[INFO] Subnet ID : %s", subnet)

	return resourceIBMISSubnetNetworkACLAttachmentRead(context, d, meta)
}

func resourceIBMISSubnetNetworkACLAttachmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_network_acl_attachment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getSubnetNetworkACLOptionsModel := &vpcv1.GetSubnetNetworkACLOptions{
		ID: &id,
	}
	nwacl, response, err := sess.GetSubnetNetworkACLWithContext(context, getSubnetNetworkACLOptionsModel)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetNetworkACLWithContext failed: %s", err.Error()), "ibm_is_subnet_network_acl_attachment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(isNetworkACLName, nwacl.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_network_acl_attachment", "read", "set-name").GetDiag()
	}
	if err = d.Set(isNetworkACLCRN, nwacl.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_network_acl_attachment", "read", "set-crn").GetDiag()
	}
	if !core.IsNil(nwacl.VPC) {
		if err = d.Set(isNetworkACLVPC, *nwacl.VPC.ID); err != nil {
			err = fmt.Errorf("Error setting vpc: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_network_acl_attachment", "read", "set-vpc").GetDiag()
		}
	}
	if err = d.Set(isNetworkACLID, *nwacl.ID); err != nil {
		err = fmt.Errorf("Error setting network_acl: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_network_acl_attachment", "read", "set-network_acl").GetDiag()
	}
	if nwacl.ResourceGroup != nil {
		if err = d.Set(isNetworkACLResourceGroup, *nwacl.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_network_acl_attachment", "read", "set-resource_group").GetDiag()
		}
	}
	rules := make([]interface{}, 0)
	if len(nwacl.Rules) > 0 {
		for _, rulex := range nwacl.Rules {
			log.Println("[DEBUG] Type of the Rule", reflect.TypeOf(rulex))
			rule := make(map[string]interface{})
			switch reflect.TypeOf(rulex).String() {
			case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmp":
				{
					rulex := rulex.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolIcmp)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IPVersion
					rule[isNetworkACLRuleSource] = *rulex.Source
					rule[isNetworkACLRuleDestination] = *rulex.Destination
					rule[isNetworkACLRuleDirection] = *rulex.Direction
					rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
					icmp := make([]map[string]int, 1, 1)
					if rulex.Code != nil && rulex.Type != nil {
						icmp[0] = map[string]int{
							isNetworkACLRuleICMPCode: int(*rulex.Code),
							isNetworkACLRuleICMPType: int(*rulex.Code),
						}
					}
					rule[isNetworkACLRuleICMP] = icmp
				}
			case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp":
				{
					rulex := rulex.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTcpudp)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IPVersion
					rule[isNetworkACLRuleSource] = *rulex.Source
					rule[isNetworkACLRuleDestination] = *rulex.Destination
					rule[isNetworkACLRuleDirection] = *rulex.Direction
					if *rulex.Protocol == "tcp" {
						rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
						rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
						tcp := make([]map[string]int, 1, 1)
						tcp[0] = map[string]int{
							isNetworkACLRuleSourcePortMax: checkNetworkACLNil(rulex.SourcePortMax),
							isNetworkACLRuleSourcePortMin: checkNetworkACLNil(rulex.SourcePortMin),
						}
						tcp[0][isNetworkACLRulePortMax] = checkNetworkACLNil(rulex.DestinationPortMax)
						tcp[0][isNetworkACLRulePortMin] = checkNetworkACLNil(rulex.DestinationPortMin)
						rule[isNetworkACLRuleTCP] = tcp
					} else if *rulex.Protocol == "udp" {
						rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
						rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
						udp := make([]map[string]int, 1, 1)
						udp[0] = map[string]int{
							isNetworkACLRuleSourcePortMax: checkNetworkACLNil(rulex.SourcePortMax),
							isNetworkACLRuleSourcePortMin: checkNetworkACLNil(rulex.SourcePortMin),
						}
						udp[0][isNetworkACLRulePortMax] = checkNetworkACLNil(rulex.DestinationPortMax)
						udp[0][isNetworkACLRulePortMin] = checkNetworkACLNil(rulex.DestinationPortMin)
						rule[isNetworkACLRuleUDP] = udp
					}
				}
			case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolAll":
				{
					rulex := rulex.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolAll)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IPVersion
					rule[isNetworkACLRuleSource] = *rulex.Source
					rule[isNetworkACLRuleDestination] = *rulex.Destination
					rule[isNetworkACLRuleDirection] = *rulex.Direction
					rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
				}
			}
			rules = append(rules, rule)
		}
	}
	if err = d.Set(isNetworkACLRules, rules); err != nil {
		err = fmt.Errorf("Error setting rules: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_network_acl_attachment", "read", "set-rules").GetDiag()
	}

	return nil
}

func resourceIBMISSubnetNetworkACLAttachmentUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_network_acl_attachment", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if d.HasChange(isNetworkACLID) {
		subnet := d.Get(isSubnetID).(string)
		networkACL := d.Get(isNetworkACLID).(string)

		// Construct an instance of the NetworkACLIdentityByID model
		networkACLIdentityModel := new(vpcv1.NetworkACLIdentityByID)
		networkACLIdentityModel.ID = &networkACL

		// Construct an instance of the ReplaceSubnetNetworkACLOptions model
		replaceSubnetNetworkACLOptionsModel := new(vpcv1.ReplaceSubnetNetworkACLOptions)
		replaceSubnetNetworkACLOptionsModel.ID = &subnet
		replaceSubnetNetworkACLOptionsModel.NetworkACLIdentity = networkACLIdentityModel
		resultACL, _, err := sess.ReplaceSubnetNetworkACLWithContext(context, replaceSubnetNetworkACLOptionsModel)

		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceSubnetNetworkACLWithContext failed: %s", err.Error()), "ibm_is_subnet_network_acl_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		log.Printf("[INFO] Updated subnet %s with Network ACL : %s", subnet, *resultACL.ID)

		d.SetId(subnet)
		return resourceIBMISSubnetNetworkACLAttachmentRead(context, d, meta)
	}

	return resourceIBMISSubnetNetworkACLAttachmentRead(context, d, meta)
}

func resourceIBMISSubnetNetworkACLAttachmentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_network_acl_attachment", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	// Set the subnet with VPC default network ACL
	getSubnetOptions := &vpcv1.GetSubnetOptions{
		ID: &id,
	}
	subnet, response, err := sess.GetSubnetWithContext(context, getSubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetWithContext failed: %s", err.Error()), "ibm_is_subnet_network_acl_attachment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	// Fetch VPC
	vpcID := *subnet.VPC.ID

	getvpcOptions := &vpcv1.GetVPCOptions{
		ID: &vpcID,
	}
	vpc, response, err := sess.GetVPCWithContext(context, getvpcOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCWithContext failed: %s", err.Error()), "ibm_is_subnet_network_acl_attachment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Fetch default network ACL
	if vpc.DefaultNetworkACL != nil {
		log.Printf("[DEBUG] vpc default network acl is not null :%s", *vpc.DefaultNetworkACL.ID)
		// Construct an instance of the NetworkACLIdentityByID model
		networkACLIdentityModel := new(vpcv1.NetworkACLIdentityByID)
		networkACLIdentityModel.ID = vpc.DefaultNetworkACL.ID

		// Construct an instance of the ReplaceSubnetNetworkACLOptions model
		replaceSubnetNetworkACLOptionsModel := new(vpcv1.ReplaceSubnetNetworkACLOptions)
		replaceSubnetNetworkACLOptionsModel.ID = &id
		replaceSubnetNetworkACLOptionsModel.NetworkACLIdentity = networkACLIdentityModel
		resultACL, _, err := sess.ReplaceSubnetNetworkACLWithContext(context, replaceSubnetNetworkACLOptionsModel)

		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceSubnetNetworkACLWithContext failed: %s", err.Error()), "ibm_is_subnet_network_acl_attachment", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		log.Printf("[INFO] Updated subnet %s with VPC default Network ACL : %s", id, *resultACL.ID)
	} else {
		log.Printf("[DEBUG] vpc default network acl is  null")
	}

	d.SetId("")
	return nil
}

func resourceIBMISSubnetNetworkACLAttachmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id := d.Id()
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_network_acl_attachment", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, err
	}
	getSubnetNetworkACLOptionsModel := &vpcv1.GetSubnetNetworkACLOptions{
		ID: &id,
	}
	_, response, err := sess.GetSubnetNetworkACL(getSubnetNetworkACLOptionsModel)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetNetworkACL failed: %s", err.Error()), "ibm_is_subnet_network_acl_attachment", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}
