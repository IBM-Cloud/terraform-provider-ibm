package ibm

import (
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

const (
	isNetworkACLName            = "name"
	isNetworkACLRules           = "rules"
	isNetworkACLSubnets         = "subnets"
	isNetworkACLRuleID          = "id"
	isNetworkACLRuleName        = "name"
	isNetworkACLRuleAction      = "action"
	isNetworkACLRuleIPVersion   = "ip_version"
	isNetworkACLRuleSource      = "source"
	isNetworkACLRuleDestination = "destination"
	isNetworkACLRuleDirection   = "direction"
	isNetworkACLRuleProtocol    = "protocol"
	isNetworkACLRuleICMP        = "icmp"
	isNetworkACLRuleICMPCode    = "code"
	isNetworkACLRuleICMPType    = "type"
	isNetworkACLRuleTCP         = "tcp"
	isNetworkACLRuleUDP         = "udp"
	isNetworkACLRulePortMax     = "port_max"
	isNetworkACLRulePortMin     = "port_min"
)

func resourceIBMISNetworkACL() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISNetworkACLCreate,
		Read:     resourceIBMISNetworkACLRead,
		Update:   resourceIBMISNetworkACLUpdate,
		Delete:   resourceIBMISNetworkACLDelete,
		Exists:   resourceIBMISNetworkACLExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			isNetworkACLName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
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
							Type:     schema.TypeString,
							Required: true,
							ForceNew: false,
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
									},
									isNetworkACLRulePortMin: {
										Type:     schema.TypeInt,
										Optional: true,
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
									},
									isNetworkACLRulePortMin: {
										Type:     schema.TypeInt,
										Optional: true,
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
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	nwaclC := network.NewNetworkAclClient(sess)

	var nwaclbody models.PostNetworkAclsParamsBody
	nwaclbody.Name = d.Get(isNetworkACLName).(string)

	//validate each rule before attempting to create the ACL
	rules := d.Get(isNetworkACLRules).([]interface{})
	err = validateInlineRules(rules)
	if err != nil {
		return err
	}

	reqbody, _ := json.Marshal(nwaclbody)
	log.Printf("[DEBUG] Creating Network ACL : %s", string(reqbody))

	//TODO : Tags

	nwacl, err := nwaclC.Create(&nwaclbody)
	if err != nil {
		log.Printf("[DEBUG] Network ACL creation failed with error : %s", isErrorToString(err))
		return err
	}

	d.SetId(nwacl.ID.String())
	nwaclid := nwacl.ID.String()

	//Remove default rules
	err = clearRules(nwaclC, nwaclid)
	if err != nil {
		return err
	}

	err = createInlineRules(nwaclC, nwaclid, rules)
	if err != nil {
		return err
	}

	return resourceIBMISNetworkACLRead(d, meta)

}

func resourceIBMISNetworkACLRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	nwaclC := network.NewNetworkAclClient(sess)
	log.Printf("[DEBUG] Looking up details for network ACL with id %s", d.Id())
	nwacl, err := nwaclC.Get(d.Id())
	if err != nil {
		return err
	}

	d.Set(isNetworkACLName, nwacl.Name)
	d.Set(isNetworkACLSubnets, len(nwacl.Subnets))

	log.Printf("[DEBUG] Looking up rules for network ACL with id %s", d.Id())
	rawrules, _, err := nwaclC.ListRules(d.Id(), "")
	if err != nil {
		return err
	}

	sortedrules := sortrules(rawrules)
	rules := make([]interface{}, 0)
	for rawrule := sortedrules.Front(); rawrule != nil; rawrule = rawrule.Next() {
		rulex := rawrule.Value.(*models.NetworkACLRule)
		rule := make(map[string]interface{})
		rule[isNetworkACLRuleID] = rulex.ID.String()
		rule[isNetworkACLRuleName] = rulex.Name
		rule[isNetworkACLRuleAction] = rulex.Action
		rule[isNetworkACLRuleIPVersion] = rulex.IPVersion
		rule[isNetworkACLRuleSource] = rulex.Source
		rule[isNetworkACLRuleDestination] = rulex.Destination
		if rulex.Direction == "inbound" {
			rule[isNetworkACLRuleDirection] = "ingress"
		} else {
			rule[isNetworkACLRuleDirection] = "egress"
		}

		if rulex.Protocol == "icmp" {
			rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
			rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
			icmp := make([]map[string]int, 1, 1)
			icmp[0] = map[string]int{
				isNetworkACLRuleICMPCode: checkNetworkACLNil(rulex.Code),
				isNetworkACLRuleICMPType: checkNetworkACLNil(rulex.Type),
			}
			rule[isNetworkACLRuleICMP] = icmp
		} else if rulex.Protocol == "tcp" {
			rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
			rule[isNetworkACLRuleUDP] = make([]map[string]int, 0, 0)
			tcp := make([]map[string]int, 1, 1)
			tcp[0] = map[string]int{
				isNetworkACLRulePortMax: checkNetworkACLNil(rulex.PortMax),
				isNetworkACLRulePortMin: checkNetworkACLNil(rulex.PortMin),
			}
			rule[isNetworkACLRuleTCP] = tcp
		} else if rulex.Protocol == "udp" {
			rule[isNetworkACLRuleICMP] = make([]map[string]int, 0, 0)
			rule[isNetworkACLRuleTCP] = make([]map[string]int, 0, 0)
			udp := make([]map[string]int, 1, 1)
			udp[0] = map[string]int{
				isNetworkACLRulePortMax: checkNetworkACLNil(rulex.PortMax),
				isNetworkACLRulePortMin: checkNetworkACLNil(rulex.PortMin),
			}
			rule[isNetworkACLRuleUDP] = udp
		}
		rules = append(rules, rule)
	}

	d.Set(isNetworkACLRules, rules)

	return nil
}

func resourceIBMISNetworkACLUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	nwaclC := network.NewNetworkAclClient(sess)
	nwaclid := d.Id()
	rules := d.Get(isNetworkACLRules).([]interface{})

	if d.HasChange(isNetworkACLName) {
		name := d.Get(isNetworkACLName).(string)
		_, err := nwaclC.Update(nwaclid, name)
		if err != nil {
			return err
		}
	}

	if d.HasChange(isNetworkACLRules) {

		err := validateInlineRules(rules)
		if err != nil {
			return err
		}

		//Delete all existing rules
		clearRules(nwaclC, nwaclid)

		//Create the rules as per the def
		err = createInlineRules(nwaclC, d.Id(), rules)
		if err != nil {
			return err
		}
	}

	return resourceIBMISNetworkACLRead(d, meta)
}

func resourceIBMISNetworkACLDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	nwaclC := network.NewNetworkAclClient(sess)
	log.Printf("Deleting the network ACL with id %s", d.Id())
	err = nwaclC.Delete(d.Id())
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMISNetworkACLExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	nwaclC := network.NewNetworkAclClient(sess)

	_, err = nwaclC.Get(d.Id())
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}

func sortrules(rules []*models.NetworkACLRule) *list.List {
	sortedrules := list.New()
	for _, rule := range rules {
		if rule.Before == nil {
			sortedrules.PushBack(rule)
		} else {
			inserted := false
			for e := sortedrules.Front(); e != nil; e = e.Next() {
				rulex := e.Value.(*models.NetworkACLRule)
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

func clearRules(nwaclC *network.NetworkAclClient, nwaclid string) error {
	rawrules, _, err := nwaclC.ListRules(nwaclid, "")
	if err != nil {
		return err
	}

	for _, rule := range rawrules {
		err := nwaclC.DeleteRule(nwaclid, rule.ID.String())
		if err != nil {
			return err
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
		if (direction != "ingress") && (direction != "egress") {
			return fmt.Errorf("Invalid direction. valid values are ingress|egress")
		}

		icmp := len(rulex[isNetworkACLRuleICMP].([]interface{})) > 0
		tcp := len(rulex[isNetworkACLRuleTCP].([]interface{})) > 0
		udp := len(rulex[isNetworkACLRuleUDP].([]interface{})) > 0

		if (icmp && tcp) || (icmp && udp) || (tcp && udp) {
			return fmt.Errorf("Only one of icmp|tcp|udp can be defined per rule")
		}

	}
	return nil
}

func createInlineRules(nwaclC *network.NetworkAclClient, nwaclid string, rules []interface{}) error {
	before := ""

	for i := len(rules) - 1; i >= 0; i-- {
		rulex := rules[i].(map[string]interface{})

		name := rulex[isNetworkACLRuleName].(string)
		source := rulex[isNetworkACLRuleSource].(string)
		destination := rulex[isNetworkACLRuleDestination].(string)
		action := rulex[isNetworkACLRuleAction].(string)
		direction := rulex[isNetworkACLRuleDirection].(string)
		icmp := rulex[isNetworkACLRuleICMP].([]interface{})
		tcp := rulex[isNetworkACLRuleTCP].([]interface{})
		udp := rulex[isNetworkACLRuleUDP].([]interface{})
		icmptype := -1
		icmpcode := -1
		minport := -1
		maxport := -1
		protocol := "all"

		if len(icmp) > 0 {
			protocol = "icmp"
			icmpval := icmp[0].(map[string]interface{})
			if val, ok := icmpval[isNetworkACLRuleICMPType]; ok {
				icmptype = val.(int)
			}
			if val, ok := icmpval[isNetworkACLRuleICMPCode]; ok {
				icmpcode = val.(int)
			}
		} else if len(tcp) > 0 {
			protocol = "tcp"
			tcpval := tcp[0].(map[string]interface{})
			if val, ok := tcpval[isNetworkACLRulePortMin]; ok {
				minport = val.(int)
			}
			if val, ok := tcpval[isNetworkACLRulePortMax]; ok {
				maxport = val.(int)
			}
		} else if len(udp) > 0 {
			protocol = "udp"
			udpval := udp[0].(map[string]interface{})
			if val, ok := udpval[isNetworkACLRulePortMin]; ok {
				minport = val.(int)
			}
			if val, ok := udpval[isNetworkACLRulePortMax]; ok {
				maxport = val.(int)
			}
		}

		rule, err := nwaclC.AddRule(nwaclid, name, source, destination, direction, action, protocol,
			int64(icmptype), int64(icmpcode), int64(minport), int64(maxport), before)
		if err != nil {
			return err
		}

		before = rule.ID.String()
	}
	return nil
}
