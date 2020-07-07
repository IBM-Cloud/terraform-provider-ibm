package ibm

import (
	"container/list"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isNetworkACLName              = "name"
	isNetworkACLRules             = "rules"
	isNetworkACLSubnets           = "subnets"
	isNetworkACLRuleID            = "id"
	isNetworkACLRuleName          = "name"
	isNetworkACLRuleAction        = "action"
	isNetworkACLRuleIPVersion     = "ip_version"
	isNetworkACLRuleSource        = "source"
	isNetworkACLRuleDestination   = "destination"
	isNetworkACLRuleDirection     = "direction"
	isNetworkACLRuleProtocol      = "protocol"
	isNetworkACLRuleICMP          = "icmp"
	isNetworkACLRuleICMPCode      = "code"
	isNetworkACLRuleICMPType      = "type"
	isNetworkACLRuleTCP           = "tcp"
	isNetworkACLRuleUDP           = "udp"
	isNetworkACLRulePortMax       = "port_max"
	isNetworkACLRulePortMin       = "port_min"
	isNetworkACLRuleSourcePortMax = "source_port_max"
	isNetworkACLRuleSourcePortMin = "source_port_min"
	isNetworkACLVPC               = "vpc"
	isNetworkACLResourceGroup     = "resource_group"
)

func resourceIBMISNetworkACL() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISNetworkACLCreate,
		Read:     resourceIBMISNetworkACLRead,
		Update:   resourceIBMISNetworkACLUpdate,
		Delete:   resourceIBMISNetworkACLDelete,
		Exists:   resourceIBMISNetworkACLExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			isNetworkACLName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateISName,
				Description:  "Network ACL name",
			},
			isNetworkACLVPC: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Network ACL VPC name",
			},
			isNetworkACLResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "Resource group ID for the network ACL",
			},
			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},
			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
			isNetworkACLRules: {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isNetworkACLRuleID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isNetworkACLRuleName: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: false,
						},
						isNetworkACLRuleAction: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: false,
						},
						isNetworkACLRuleIPVersion: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isNetworkACLRuleSource: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: false,
						},
						isNetworkACLRuleDestination: {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: false,
						},
						isNetworkACLRuleDirection: {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     false,
							Description:  "Direction of traffic to enforce, either inbound or outbound",
							ValidateFunc: validateIsNetworkAclRuleDirection,
						},
						isNetworkACLSubnets: {
							Type:     schema.TypeInt,
							Computed: true,
						},
						isNetworkACLRuleICMP: {
							Type:     schema.TypeList,
							MinItems: 0,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isNetworkACLRuleICMPCode: {
										Type:     schema.TypeInt,
										Optional: true,
									},
									isNetworkACLRuleICMPType: {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						isNetworkACLRuleTCP: {
							Type:     schema.TypeList,
							MinItems: 0,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isNetworkACLRulePortMax: {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  65535,
									},
									isNetworkACLRulePortMin: {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  1,
									},
									isNetworkACLRuleSourcePortMax: {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  65535,
									},
									isNetworkACLRuleSourcePortMin: {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  1,
									},
								},
							},
						},

						isNetworkACLRuleUDP: {
							Type:     schema.TypeList,
							MinItems: 0,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isNetworkACLRulePortMax: {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  65535,
									},
									isNetworkACLRulePortMin: {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  1,
									},
									isNetworkACLRuleSourcePortMax: {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  65535,
									},
									isNetworkACLRuleSourcePortMin: {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  1,
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

func resourceIBMISNetworkACLCreate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	name := d.Get(isNetworkACLName).(string)

	if userDetails.generation == 1 {
		err := classicNwaclCreate(d, meta, name)
		if err != nil {
			return err
		}
	} else {
		err := nwaclCreate(d, meta, name)
		if err != nil {
			return err
		}
	}
	return resourceIBMISNetworkACLRead(d, meta)

}

func classicNwaclCreate(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	nwaclTemplate := &vpcclassicv1.NetworkACLPrototype{
		Name: &name,
	}

	var rules []interface{}
	if rls, ok := d.GetOk(isNetworkACLRules); ok {
		rules = rls.([]interface{})
	}
	err = validateInlineRules(rules)
	if err != nil {
		return err
	}

	options := &vpcclassicv1.CreateNetworkAclOptions{
		NetworkACLPrototype: nwaclTemplate,
	}

	nwacl, response, err := sess.CreateNetworkAcl(options)
	if err != nil {
		return fmt.Errorf("[DEBUG]Error while creating Network ACL err %s\n%s", err, response)
	}
	d.SetId(*nwacl.ID)
	log.Printf("[INFO] Network ACL : %s", *nwacl.ID)
	nwaclid := *nwacl.ID

	//Remove default rules
	err = classicClearRules(sess, nwaclid)
	if err != nil {
		return err
	}

	err = classicCreateInlineRules(sess, nwaclid, rules)
	if err != nil {
		return err
	}
	return nil
}

func nwaclCreate(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	var vpc, rg string
	if vpcID, ok := d.GetOk(isNetworkACLVPC); ok {
		vpc = vpcID.(string)
	} else {
		return fmt.Errorf("Required parameter vpc is not set")
	}

	nwaclTemplate := &vpcv1.NetworkACLPrototype{
		Name: &name,
		Vpc: &vpcv1.VPCIdentity{
			ID: &vpc,
		},
	}

	if grp, ok := d.GetOk(isNetworkACLResourceGroup); ok {
		rg = grp.(string)
		nwaclTemplate.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	// validate each rule before attempting to create the ACL
	var rules []interface{}
	if rls, ok := d.GetOk(isNetworkACLRules); ok {
		rules = rls.([]interface{})
	}
	err = validateInlineRules(rules)
	if err != nil {
		return err
	}

	options := &vpcv1.CreateNetworkAclOptions{
		NetworkACLPrototype: nwaclTemplate,
	}

	nwacl, response, err := sess.CreateNetworkAcl(options)
	if err != nil {
		return fmt.Errorf("[DEBUG]Error while creating Network ACL err %s\n%s", err, response)
	}
	d.SetId(*nwacl.ID)
	log.Printf("[INFO] Network ACL : %s", *nwacl.ID)
	nwaclid := *nwacl.ID

	//Remove default rules
	err = clearRules(sess, nwaclid)
	if err != nil {
		return err
	}

	err = createInlineRules(sess, nwaclid, rules)
	if err != nil {
		return err
	}
	return nil
}

func resourceIBMISNetworkACLRead(d *schema.ResourceData, meta interface{}) error {

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicNwaclGet(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := nwaclGet(d, meta, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicNwaclGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getNetworkAclOptions := &vpcclassicv1.GetNetworkAclOptions{
		ID: &id,
	}
	nwacl, response, err := sess.GetNetworkAcl(getNetworkAclOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting Network ACL(%s) : %s\n%s", id, err, response)
	}
	d.Set(isNetworkACLName, *nwacl.Name)
	d.Set(isNetworkACLSubnets, len(nwacl.Subnets))

	rules := make([]interface{}, 0)
	if len(nwacl.Rules) > 0 {
		for _, rulex := range nwacl.Rules {
			log.Println("[DEBUG] Type of the Rule", reflect.TypeOf(rulex))
			rule := make(map[string]interface{})
			switch reflect.TypeOf(rulex).String() {
			case "*vpcclassicv1.NetworkACLRuleItemNetworkACLRuleProtocolICMP":
				{
					rulex := rulex.(*vpcclassicv1.NetworkACLRuleItemNetworkACLRuleProtocolICMP)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IpVersion
					rule[isNetworkACLRuleSource] = *rulex.Source
					rule[isNetworkACLRuleDestination] = *rulex.Destination
					rule[isNetworkACLRuleDirection] = *rulex.Direction
					rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
					icmp := make([]map[string]int, 1, 1)
					icmp[0] = map[string]int{
						isNetworkACLRuleICMPCode: checkNetworkACLNil(rulex.Code),
						isNetworkACLRuleICMPType: checkNetworkACLNil(rulex.Type),
					}
					rule[isNetworkACLRuleICMP] = icmp
				}
			case "*vpcclassicv1.NetworkACLRuleItemNetworkACLRuleProtocolTCPUDP":
				{
					rulex := rulex.(*vpcclassicv1.NetworkACLRuleItemNetworkACLRuleProtocolTCPUDP)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IpVersion
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
						tcp[0][isNetworkACLRulePortMax] = checkNetworkACLNil(rulex.PortMax)
						tcp[0][isNetworkACLRulePortMin] = checkNetworkACLNil(rulex.PortMin)
						rule[isNetworkACLRuleTCP] = tcp
					} else if *rulex.Protocol == "udp" {
						rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
						rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
						udp := make([]map[string]int, 1, 1)
						udp[0] = map[string]int{
							isNetworkACLRuleSourcePortMax: checkNetworkACLNil(rulex.SourcePortMax),
							isNetworkACLRuleSourcePortMin: checkNetworkACLNil(rulex.SourcePortMin),
						}
						udp[0][isNetworkACLRulePortMax] = checkNetworkACLNil(rulex.PortMax)
						udp[0][isNetworkACLRulePortMin] = checkNetworkACLNil(rulex.PortMin)
						rule[isNetworkACLRuleUDP] = udp
					}
				}
			case "*vpcclassicv1.NetworkACLRuleItemNetworkACLRuleProtocolAll":
				{
					rulex := rulex.(*vpcclassicv1.NetworkACLRuleItemNetworkACLRuleProtocolAll)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IpVersion
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
	d.Set(isNetworkACLRules, rules)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc/network/acl")
	d.Set(ResourceName, *nwacl.Name)
	// d.Set(ResourceCRN, *nwacl.Crn)
	return nil
}

func nwaclGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getNetworkAclOptions := &vpcv1.GetNetworkAclOptions{
		ID: &id,
	}
	nwacl, response, err := sess.GetNetworkAcl(getNetworkAclOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting Network ACL(%s) : %s\n%s", id, err, response)
	}
	d.Set(isNetworkACLName, *nwacl.Name)
	d.Set(isNetworkACLVPC, *nwacl.Vpc.ID)
	if nwacl.ResourceGroup != nil {
		d.Set(isNetworkACLResourceGroup, *nwacl.ResourceGroup.ID)
		d.Set(ResourceGroupName, *nwacl.ResourceGroup.Name)
	}
	d.Set(isNetworkACLSubnets, len(nwacl.Subnets))

	rules := make([]interface{}, 0)
	if len(nwacl.Rules) > 0 {
		for _, rulex := range nwacl.Rules {
			log.Println("[DEBUG] Type of the Rule", reflect.TypeOf(rulex))
			rule := make(map[string]interface{})
			switch reflect.TypeOf(rulex).String() {
			case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolICMP":
				{
					rulex := rulex.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolICMP)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IpVersion
					rule[isNetworkACLRuleSource] = *rulex.Source
					rule[isNetworkACLRuleDestination] = *rulex.Destination
					rule[isNetworkACLRuleDirection] = *rulex.Direction
					rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
					rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
					icmp := make([]map[string]int, 1, 1)
					icmp[0] = map[string]int{
						isNetworkACLRuleICMPCode: checkNetworkACLNil(rulex.Code),
						isNetworkACLRuleICMPType: checkNetworkACLNil(rulex.Type),
					}
					rule[isNetworkACLRuleICMP] = icmp
				}
			case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTCPUDP":
				{
					rulex := rulex.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTCPUDP)
					rule[isNetworkACLRuleID] = *rulex.ID
					rule[isNetworkACLRuleName] = *rulex.Name
					rule[isNetworkACLRuleAction] = *rulex.Action
					rule[isNetworkACLRuleIPVersion] = *rulex.IpVersion
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
					rule[isNetworkACLRuleIPVersion] = *rulex.IpVersion
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
	d.Set(isNetworkACLRules, rules)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc-ext/network/acl")
	d.Set(ResourceName, *nwacl.Name)
	// d.Set(ResourceCRN, *nwacl.Crn)
	return nil
}

func resourceIBMISNetworkACLUpdate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()

	name := ""
	hasChanged := false

	if d.HasChange(isNetworkACLName) {
		name = d.Get(isNetworkACLName).(string)
		hasChanged = true
	}

	if userDetails.generation == 1 {
		err := classicNwaclUpdate(d, meta, id, name, hasChanged)
		if err != nil {
			return err
		}
	} else {
		err := nwaclUpdate(d, meta, id, name, hasChanged)
		if err != nil {
			return err
		}
	}
	return resourceIBMISNetworkACLRead(d, meta)
}

func classicNwaclUpdate(d *schema.ResourceData, meta interface{}, id, name string, hasChanged bool) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	rules := d.Get(isNetworkACLRules).([]interface{})
	if hasChanged {
		updateNetworkAclOptions := &vpcclassicv1.UpdateNetworkAclOptions{
			ID:   &id,
			Name: &name,
		}
		_, response, err := sess.UpdateNetworkAcl(updateNetworkAclOptions)
		if err != nil {
			return fmt.Errorf("Error Updating Network ACL(%s) : %s\n%s", id, err, response)
		}
	}
	if d.HasChange(isNetworkACLRules) {
		err := validateInlineRules(rules)
		if err != nil {
			return err
		}
		//Delete all existing rules
		err = classicClearRules(sess, id)
		if err != nil {
			return err
		}
		//Create the rules as per the def
		err = classicCreateInlineRules(sess, id, rules)
		if err != nil {
			return err
		}
	}
	return nil
}

func nwaclUpdate(d *schema.ResourceData, meta interface{}, id, name string, hasChanged bool) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	rules := d.Get(isNetworkACLRules).([]interface{})
	if hasChanged {
		updateNetworkAclOptions := &vpcv1.UpdateNetworkAclOptions{
			ID:   &id,
			Name: &name,
		}
		_, response, err := sess.UpdateNetworkAcl(updateNetworkAclOptions)
		if err != nil {
			return fmt.Errorf("Error Updating Network ACL(%s) : %s\n%s", id, err, response)
		}
	}
	if d.HasChange(isNetworkACLRules) {
		err := validateInlineRules(rules)
		if err != nil {
			return err
		}
		//Delete all existing rules
		err = clearRules(sess, id)
		if err != nil {
			return err
		}
		//Create the rules as per the def
		err = createInlineRules(sess, id, rules)
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceIBMISNetworkACLDelete(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicNwaclDelete(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := nwaclDelete(d, meta, id)
		if err != nil {
			return err
		}
	}

	d.SetId("")
	return nil
}

func classicNwaclDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	getNetworkAclOptions := &vpcclassicv1.GetNetworkAclOptions{
		ID: &id,
	}
	_, response, err := sess.GetNetworkAcl(getNetworkAclOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting Network ACL (%s): %s\n%s", id, err, response)
	}

	deleteNetworkAclOptions := &vpcclassicv1.DeleteNetworkAclOptions{
		ID: &id,
	}
	response, err = sess.DeleteNetworkAcl(deleteNetworkAclOptions)
	if err != nil {
		return fmt.Errorf("Error Deleting Network ACL : %s\n%s", err, response)
	}
	d.SetId("")
	return nil
}

func nwaclDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	getNetworkAclOptions := &vpcv1.GetNetworkAclOptions{
		ID: &id,
	}
	_, response, err := sess.GetNetworkAcl(getNetworkAclOptions)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error Getting Network ACL (%s): %s\n%s", id, err, response)
	}

	deleteNetworkAclOptions := &vpcv1.DeleteNetworkAclOptions{
		ID: &id,
	}
	response, err = sess.DeleteNetworkAcl(deleteNetworkAclOptions)
	if err != nil {
		return fmt.Errorf("Error Deleting Network ACL : %s\n%s", err, response)
	}
	d.SetId("")
	return nil
}

func resourceIBMISNetworkACLExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		exists, err := classicNwaclExists(d, meta, id)
		return exists, err
	} else {
		exists, err := nwaclExists(d, meta, id)
		return exists, err
	}
}

func classicNwaclExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return false, err
	}
	getNetworkAclOptions := &vpcclassicv1.GetNetworkAclOptions{
		ID: &id,
	}
	_, response, err := sess.GetNetworkAcl(getNetworkAclOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting Network ACL: %s\n%s", err, response)
	}
	return true, nil
}

func nwaclExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	getNetworkAclOptions := &vpcv1.GetNetworkAclOptions{
		ID: &id,
	}
	_, response, err := sess.GetNetworkAcl(getNetworkAclOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting Network ACL: %s\n%s", err, response)
	}
	return true, nil
}

func sortclassicrules(rules []*vpcclassicv1.NetworkACLRuleItem) *list.List {
	sortedrules := list.New()
	for _, rule := range rules {
		if rule.Before == nil {
			sortedrules.PushBack(rule)
		} else {
			inserted := false
			for e := sortedrules.Front(); e != nil; e = e.Next() {
				rulex := e.Value.(*vpcclassicv1.NetworkACLRuleItem)
				if rulex.ID == rule.Before.ID {
					sortedrules.InsertAfter(rule, e)
					inserted = true
					break
				}
			}
			// if we didnt find before yet, just put it at the head of the list
			if !inserted {
				sortedrules.PushFront(rule)
			}
		}
	}
	return sortedrules
}

func checkNetworkACLNil(ptr *int64) int {
	if ptr == nil {
		return 0
	}
	return int(*ptr)
}

func classicClearRules(nwaclC *vpcclassicv1.VpcClassicV1, nwaclid string) error {
	listNetworkAclRulesOptions := &vpcclassicv1.ListNetworkAclRulesOptions{
		NetworkAclID: &nwaclid,
	}
	rawrules, response, err := nwaclC.ListNetworkAclRules(listNetworkAclRulesOptions)
	if err != nil {
		return fmt.Errorf("Error Listing network ACL rules : %s\n%s", err, response)
	}

	for _, rule := range rawrules.Rules {
		deleteNetworkAclRuleOptions := &vpcclassicv1.DeleteNetworkAclRuleOptions{
			NetworkAclID: &nwaclid,
		}
		switch reflect.TypeOf(rule).String() {
		case "*vpcclassicv1.NetworkACLRuleItemNetworkACLRuleProtocolICMP":
			rule := rule.(*vpcclassicv1.NetworkACLRuleItemNetworkACLRuleProtocolICMP)
			deleteNetworkAclRuleOptions.ID = rule.ID
		case "*vpcclassicv1.NetworkACLRuleItemNetworkACLRuleProtocolTCPUDP":
			rule := rule.(*vpcclassicv1.NetworkACLRuleItemNetworkACLRuleProtocolTCPUDP)
			deleteNetworkAclRuleOptions.ID = rule.ID
		case "*vpcclassicv1.NetworkACLRuleItemNetworkACLRuleProtocolAll":
			rule := rule.(*vpcclassicv1.NetworkACLRuleItemNetworkACLRuleProtocolAll)
			deleteNetworkAclRuleOptions.ID = rule.ID
		}

		response, err := nwaclC.DeleteNetworkAclRule(deleteNetworkAclRuleOptions)
		if err != nil {
			return fmt.Errorf("Error Deleting network ACL rule : %s\n%s", err, response)
		}
	}
	return nil
}

func clearRules(nwaclC *vpcv1.VpcV1, nwaclid string) error {
	listNetworkAclRulesOptions := &vpcv1.ListNetworkAclRulesOptions{
		NetworkAclID: &nwaclid,
	}
	rawrules, response, err := nwaclC.ListNetworkAclRules(listNetworkAclRulesOptions)
	if err != nil {
		return fmt.Errorf("Error Listing network ACL rules : %s\n%s", err, response)
	}

	for _, rule := range rawrules.Rules {
		deleteNetworkAclRuleOptions := &vpcv1.DeleteNetworkAclRuleOptions{
			NetworkAclID: &nwaclid,
		}
		switch reflect.TypeOf(rule).String() {
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolICMP":
			rule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolICMP)
			deleteNetworkAclRuleOptions.ID = rule.ID
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTCPUDP":
			rule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolTCPUDP)
			deleteNetworkAclRuleOptions.ID = rule.ID
		case "*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolAll":
			rule := rule.(*vpcv1.NetworkACLRuleItemNetworkACLRuleProtocolAll)
			deleteNetworkAclRuleOptions.ID = rule.ID
		}

		response, err := nwaclC.DeleteNetworkAclRule(deleteNetworkAclRuleOptions)
		if err != nil {
			return fmt.Errorf("Error Deleting network ACL rule : %s\n%s", err, response)
		}
	}
	return nil
}

func validateInlineRules(rules []interface{}) error {
	for _, rule := range rules {
		rulex := rule.(map[string]interface{})
		action := rulex[isNetworkACLRuleAction].(string)
		if (action != "allow") && (action != "deny") {
			return fmt.Errorf("Invalid action. valid values are allow|deny")
		}

		direction := rulex[isNetworkACLRuleDirection].(string)
		direction = strings.ToLower(direction)

		icmp := len(rulex[isNetworkACLRuleICMP].([]interface{})) > 0
		tcp := len(rulex[isNetworkACLRuleTCP].([]interface{})) > 0
		udp := len(rulex[isNetworkACLRuleUDP].([]interface{})) > 0

		if (icmp && tcp) || (icmp && udp) || (tcp && udp) {
			return fmt.Errorf("Only one of icmp|tcp|udp can be defined per rule")
		}

	}
	return nil
}

func classicCreateInlineRules(nwaclC *vpcclassicv1.VpcClassicV1, nwaclid string, rules []interface{}) error {
	before := ""

	for i := 0; i <= len(rules)-1; i++ {
		rulex := rules[i].(map[string]interface{})

		name := rulex[isNetworkACLRuleName].(string)
		source := rulex[isNetworkACLRuleSource].(string)
		destination := rulex[isNetworkACLRuleDestination].(string)
		action := rulex[isNetworkACLRuleAction].(string)
		direction := rulex[isNetworkACLRuleDirection].(string)
		icmp := rulex[isNetworkACLRuleICMP].([]interface{})
		tcp := rulex[isNetworkACLRuleTCP].([]interface{})
		udp := rulex[isNetworkACLRuleUDP].([]interface{})
		icmptype := int64(-1)
		icmpcode := int64(-1)
		minport := int64(-1)
		maxport := int64(-1)
		sourceminport := int64(-1)
		sourcemaxport := int64(-1)
		protocol := "all"

		ruleTemplate := &vpcclassicv1.NetworkACLRulePrototype{
			Action:      &action,
			Destination: &destination,
			Direction:   &direction,
			Source:      &source,
			Name:        &name,
		}

		if before != "" {
			ruleTemplate.Before = &vpcclassicv1.NetworkACLRulePrototypeBefore{
				ID: &before,
			}
		}

		if len(icmp) > 0 {
			protocol = "icmp"
			ruleTemplate.Protocol = &protocol
			if icmp[0] != nil {
				icmpval := icmp[0].(map[string]interface{})
				if val, ok := icmpval[isNetworkACLRuleICMPType]; ok {
					icmptype = int64(val.(int))
					ruleTemplate.Type = &icmptype
				}
				if val, ok := icmpval[isNetworkACLRuleICMPCode]; ok {
					icmpcode = int64(val.(int))
					ruleTemplate.Code = &icmpcode
				}
			}
		} else if len(tcp) > 0 {
			protocol = "tcp"
			ruleTemplate.Protocol = &protocol
			tcpval := tcp[0].(map[string]interface{})
			if val, ok := tcpval[isNetworkACLRulePortMin]; ok {
				minport = int64(val.(int))
				ruleTemplate.PortMin = &minport
			}
			if val, ok := tcpval[isNetworkACLRulePortMax]; ok {
				maxport = int64(val.(int))
				ruleTemplate.PortMax = &maxport
			}
			if val, ok := tcpval[isNetworkACLRuleSourcePortMin]; ok {
				sourceminport = int64(val.(int))
				ruleTemplate.SourcePortMin = &sourceminport
			}
			if val, ok := tcpval[isNetworkACLRuleSourcePortMax]; ok {
				sourcemaxport = int64(val.(int))
				ruleTemplate.SourcePortMax = &sourcemaxport
			}
		} else if len(udp) > 0 {
			protocol = "udp"
			ruleTemplate.Protocol = &protocol
			udpval := udp[0].(map[string]interface{})
			if val, ok := udpval[isNetworkACLRulePortMin]; ok {
				minport = int64(val.(int))
				ruleTemplate.PortMin = &minport
			}
			if val, ok := udpval[isNetworkACLRulePortMax]; ok {
				maxport = int64(val.(int))
				ruleTemplate.PortMax = &maxport
			}
			if val, ok := udpval[isNetworkACLRuleSourcePortMin]; ok {
				sourceminport = int64(val.(int))
				ruleTemplate.SourcePortMin = &sourceminport
			}
			if val, ok := udpval[isNetworkACLRuleSourcePortMax]; ok {
				sourcemaxport = int64(val.(int))
				ruleTemplate.SourcePortMax = &sourcemaxport
			}
		}
		if protocol == "all" {
			ruleTemplate.Protocol = &protocol
		}

		createNetworkAclRuleOptions := &vpcclassicv1.CreateNetworkAclRuleOptions{
			NetworkAclID:            &nwaclid,
			NetworkACLRulePrototype: ruleTemplate,
		}
		_, response, err := nwaclC.CreateNetworkAclRule(createNetworkAclRuleOptions)
		if err != nil {
			return fmt.Errorf("Error Creating network ACL rule : %s\n%s", err, response)
		}
	}
	return nil
}

func createInlineRules(nwaclC *vpcv1.VpcV1, nwaclid string, rules []interface{}) error {
	before := ""

	for i := 0; i <= len(rules)-1; i++ {
		rulex := rules[i].(map[string]interface{})

		name := rulex[isNetworkACLRuleName].(string)
		source := rulex[isNetworkACLRuleSource].(string)
		destination := rulex[isNetworkACLRuleDestination].(string)
		action := rulex[isNetworkACLRuleAction].(string)
		direction := rulex[isNetworkACLRuleDirection].(string)
		icmp := rulex[isNetworkACLRuleICMP].([]interface{})
		tcp := rulex[isNetworkACLRuleTCP].([]interface{})
		udp := rulex[isNetworkACLRuleUDP].([]interface{})
		icmptype := int64(-1)
		icmpcode := int64(-1)
		minport := int64(-1)
		maxport := int64(-1)
		sourceminport := int64(-1)
		sourcemaxport := int64(-1)
		protocol := "all"

		ruleTemplate := &vpcv1.NetworkACLRulePrototype{
			Action:      &action,
			Destination: &destination,
			Direction:   &direction,
			Source:      &source,
			Name:        &name,
		}

		if before != "" {
			ruleTemplate.Before = &vpcv1.NetworkACLRulePrototypeBefore{
				ID: &before,
			}
		}

		if len(icmp) > 0 {
			protocol = "icmp"
			ruleTemplate.Protocol = &protocol
			if icmp[0] != nil {
				icmpval := icmp[0].(map[string]interface{})
				if val, ok := icmpval[isNetworkACLRuleICMPType]; ok {
					icmptype = int64(val.(int))
					ruleTemplate.Type = &icmptype
				}
				if val, ok := icmpval[isNetworkACLRuleICMPCode]; ok {
					icmpcode = int64(val.(int))
					ruleTemplate.Code = &icmpcode
				}
			}
		} else if len(tcp) > 0 {
			protocol = "tcp"
			ruleTemplate.Protocol = &protocol
			tcpval := tcp[0].(map[string]interface{})
			if val, ok := tcpval[isNetworkACLRulePortMin]; ok {
				minport = int64(val.(int))
				ruleTemplate.DestinationPortMin = &minport
			}
			if val, ok := tcpval[isNetworkACLRulePortMax]; ok {
				maxport = int64(val.(int))
				ruleTemplate.DestinationPortMax = &maxport
			}
			if val, ok := tcpval[isNetworkACLRuleSourcePortMin]; ok {
				sourceminport = int64(val.(int))
				ruleTemplate.SourcePortMin = &sourceminport
			}
			if val, ok := tcpval[isNetworkACLRuleSourcePortMax]; ok {
				sourcemaxport = int64(val.(int))
				ruleTemplate.SourcePortMax = &sourcemaxport
			}
		} else if len(udp) > 0 {
			protocol = "udp"
			ruleTemplate.Protocol = &protocol
			udpval := udp[0].(map[string]interface{})
			if val, ok := udpval[isNetworkACLRulePortMin]; ok {
				minport = int64(val.(int))
				ruleTemplate.DestinationPortMin = &minport
			}
			if val, ok := udpval[isNetworkACLRulePortMax]; ok {
				maxport = int64(val.(int))
				ruleTemplate.DestinationPortMax = &maxport
			}
			if val, ok := udpval[isNetworkACLRuleSourcePortMin]; ok {
				sourceminport = int64(val.(int))
				ruleTemplate.SourcePortMin = &sourceminport
			}
			if val, ok := udpval[isNetworkACLRuleSourcePortMax]; ok {
				sourcemaxport = int64(val.(int))
				ruleTemplate.SourcePortMax = &sourcemaxport
			}
		}
		if protocol == "all" {
			ruleTemplate.Protocol = &protocol
		}

		createNetworkAclRuleOptions := &vpcv1.CreateNetworkAclRuleOptions{
			NetworkAclID:            &nwaclid,
			NetworkACLRulePrototype: ruleTemplate,
		}
		_, response, err := nwaclC.CreateNetworkAclRule(createNetworkAclRuleOptions)
		if err != nil {
			return fmt.Errorf("Error Creating network ACL rule : %s\n%s", err, response)
		}
	}
	return nil
}
