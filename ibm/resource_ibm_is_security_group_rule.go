package ibm

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

const (
	isSecurityGroupRuleCode             = "code"
	isSecurityGroupRuleDirection        = "direction"
	isSecurityGroupRuleIPVersion        = "ip_version"
	isSecurityGroupRuleIPVersionDefault = "ipv4"
	isSecurityGroupRulePortMax          = "port_max"
	isSecurityGroupRulePortMin          = "port_min"
	isSecurityGroupRuleProtocolICMP     = "icmp"
	isSecurityGroupRuleProtocolTCP      = "tcp"
	isSecurityGroupRuleProtocolUDP      = "udp"
	isSecurityGroupRuleProtocol         = "protocol"
	isSecurityGroupRuleRemote           = "remote"
	isSecurityGroupRuleType             = "type"
	isSecurityGroupID                   = "group"
	isSecurityGroupRuleID               = "rule_id"
)

func resourceIBMISSecurityGroupRule() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMISSecurityGroupRuleCreate,
		Read:     resourceIBMISSecurityGroupRuleRead,
		Update:   resourceIBMISSecurityGroupRuleUpdate,
		Delete:   resourceIBMISSecurityGroupRuleDelete,
		Exists:   resourceIBMISSecurityGroupRuleExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			isSecurityGroupID: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Security group id",
				ForceNew:     true,
				ValidateFunc: validateSecurityGroupId,
			},

			isSecurityGroupRuleID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Rule id",
			},

			isSecurityGroupRuleDirection: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Direction of traffic to enforce, either ingress or egress",
				ValidateFunc: validateSecurityRuleDirection,
			},

			isSecurityGroupRuleIPVersion: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "IP version: ipv4 or ipv6",
				Default:      isSecurityGroupRuleIPVersionDefault,
				ValidateFunc: validateIPVersion,
			},

			isSecurityGroupRuleRemote: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Security group id: an IP address, a CIDR block, or a single security group identifier",
				ValidateFunc: validateSecurityGroupRemote,
			},

			isSecurityGroupRuleProtocolICMP: {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				MinItems:      1,
				ConflictsWith: []string{isSecurityGroupRuleProtocolTCP, isSecurityGroupRuleProtocolUDP},
				Description:   "protocol=icmp",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isSecurityGroupRuleType: {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateICMPType,
						},
						isSecurityGroupRuleCode: {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateICMPCode,
						},
					},
				},
			},

			isSecurityGroupRuleProtocolTCP: {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				MinItems:      1,
				Description:   "protocol=tcp",
				ConflictsWith: []string{isSecurityGroupRuleProtocolUDP, isSecurityGroupRuleProtocolICMP},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isSecurityGroupRulePortMin: {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateISSecurityRulePort,
						},
						isSecurityGroupRulePortMax: {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateISSecurityRulePort,
						},
					},
				},
			},

			isSecurityGroupRuleProtocolUDP: {
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				MinItems:      1,
				Description:   "protocol=udp",
				ConflictsWith: []string{isSecurityGroupRuleProtocolTCP, isSecurityGroupRuleProtocolICMP},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isSecurityGroupRulePortMin: {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateISSecurityRulePort,
						},
						isSecurityGroupRulePortMax: {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateISSecurityRulePort,
						},
					},
				},
			},
		},
	}
}

func resourceIBMISSecurityGroupRuleCreate(d *schema.ResourceData, meta interface{}) error {
	parsed, err := parseIBMISSecurityGroupRuleDictionary(d, "create")
	if err != nil {
		return err
	}
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgC := network.NewSecurityGroupClient(sess)

	rule, err := sgC.AddRule(parsed.secgrpID, parsed.direction, parsed.ipversion, parsed.protocol,
		parsed.remoteAddress, parsed.remoteCIDR, parsed.remoteSecGrpID,
		parsed.icmpType, parsed.icmpCode, parsed.portMin, parsed.portMax)
	if err != nil {
		return err
	}
	d.Set(isSecurityGroupRuleID, rule.ID.String())
	tfID := makeTerraformRuleID(parsed.secgrpID, rule.ID.String())
	d.SetId(tfID)
	err = resourceIBMISSecurityGroupRuleRead(d, meta)
	return err
}

func resourceIBMISSecurityGroupRuleRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgC := network.NewSecurityGroupClient(sess)
	secgrpID, ruleID, err := parseISTerraformID(d.Id())
	if err != nil {
		return err
	}

	rule, err := sgC.GetRule(secgrpID, ruleID)
	if err != nil {
		return err
	}
	tfID := makeTerraformRuleID(secgrpID, rule.ID.String())

	d.Set(isSecurityGroupID, secgrpID)
	d.Set(isSecurityGroupRuleID, rule.ID.String())
	if rule.Direction == "inbound" {
		d.Set(isSecurityGroupRuleDirection, "ingress")
	} else {
		d.Set(isSecurityGroupRuleDirection, "egress")
	}

	d.Set(isSecurityGroupRuleIPVersion, rule.IPVersion)
	d.Set(isSecurityGroupRuleProtocol, rule.Protocol)
	protocol := "all"
	if rule.Protocol != nil {
		protocol = *rule.Protocol
	}

	if protocol == isSecurityGroupRuleProtocolTCP || protocol == isSecurityGroupRuleProtocolUDP {

		tcpProtocol := map[string]interface{}{}

		if rule.PortMin != nil {
			tcpProtocol["port_min"] = *rule.PortMin
		}
		if rule.PortMax != nil {
			tcpProtocol["port_max"] = *rule.PortMax
		}
		protocolList := make([]map[string]interface{}, 0)
		protocolList = append(protocolList, tcpProtocol)
		if protocol == isSecurityGroupRuleProtocolTCP {
			d.Set(isSecurityGroupRuleProtocolTCP, protocolList)
		} else {
			d.Set(isSecurityGroupRuleProtocolUDP, protocolList)
		}
	}
	if protocol == isSecurityGroupRuleProtocolICMP {
		icmpProtocol := map[string]interface{}{}

		if rule.Type != nil {
			icmpProtocol["type"] = *rule.Type
		}
		if rule.Code != nil {
			icmpProtocol["code"] = *rule.Code
		}
		protocolList := make([]map[string]interface{}, 0)
		protocolList = append(protocolList, icmpProtocol)
		d.Set(isSecurityGroupRuleProtocolICMP, protocolList)
	}
	remote, err := extractRuleRemote(rule.Remote)
	if err != nil {
		return err
	}
	d.Set(isSecurityGroupRuleRemote, remote)
	d.SetId(tfID)
	return nil
}

func resourceIBMISSecurityGroupRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgC := network.NewSecurityGroupClient(sess)
	parsed, err := parseIBMISSecurityGroupRuleDictionary(d, "update")
	if err != nil {
		return err
	}

	_, err = sgC.UpdateRule(parsed.secgrpID, parsed.ruleID, parsed.direction,
		parsed.ipversion, parsed.protocol, parsed.remoteAddress,
		parsed.remoteCIDR, parsed.remoteSecGrpID, parsed.icmpType,
		parsed.icmpCode, parsed.portMin, parsed.portMax)
	if err != nil {
		return err
	}
	return resourceIBMISSecurityGroupRuleRead(d, meta)
}

func resourceIBMISSecurityGroupRuleDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgC := network.NewSecurityGroupClient(sess)

	secgrpID, ruleID, err := parseISTerraformID(d.Id())
	if err != nil {
		return err
	}
	err = sgC.DeleteRule(secgrpID, ruleID)
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return nil
			}
		}
		return err
	}
	return err
}

func resourceIBMISSecurityGroupRuleExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	sgC := network.NewSecurityGroupClient(sess)

	secgrpID, ruleID, err := parseISTerraformID(d.Id())
	if err != nil {
		return false, err
	}
	_, err = sgC.GetRule(secgrpID, ruleID)
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

func parseISTerraformID(s string) (string, string, error) {
	segments := strings.Split(s, ".")
	if len(segments) != 2 {
		return "", "", fmt.Errorf("invalid terraform Id %s (incorrect number of segments)", s)
	}
	if segments[0] == "" || segments[1] == "" {
		return "", "", fmt.Errorf("invalid terraform Id %s (one or more empty segments)", s)
	}
	return segments[0], segments[1], nil
}

type parsedIBMISSecurityGroupRuleDictionary struct {
	// After parsing, unused string fields are set to
	// "" and unused int64 fields will be set to -1.
	// This ("" for unused strings and -1 for unused int64s)
	// is expected by our riaas API client.
	secgrpID       string
	ruleID         string
	direction      string
	ipversion      string
	remote         string
	remoteAddress  string
	remoteCIDR     string
	remoteSecGrpID string
	protocol       string
	icmpType       int64
	icmpCode       int64
	portMin        int64
	portMax        int64
}

func inferRemoteSecurityGroup(s string) (address, cidr, id string, err error) {
	if isSecurityGroupAddress(s) {
		address = s
		return
	}
	if isSecurityGroupCIDR(s) {
		cidr = s
		return
	}
	if isSecurityGroupIdentityByID(s) {
		id = s
		return
	}
	err = fmt.Errorf("%s is not an acceptable %s ", s, isSecurityGroupRuleRemote)
	return
}

func parseIBMISSecurityGroupRuleDictionary(d *schema.ResourceData, tag string) (*parsedIBMISSecurityGroupRuleDictionary, error) {
	parsed := &parsedIBMISSecurityGroupRuleDictionary{}
	var err error
	parsed.icmpType = -1
	parsed.icmpCode = -1
	parsed.portMin = -1
	parsed.portMax = -1
	parsed.secgrpID, parsed.ruleID, err = parseISTerraformID(d.Id())
	if err != nil {
		parsed.secgrpID = d.Get(isSecurityGroupID).(string)
	}
	parsed.direction = d.Get(isSecurityGroupRuleDirection).(string)
	parsed.ipversion = d.Get(isSecurityGroupRuleIPVersion).(string)
	parsed.remote = d.Get(isSecurityGroupRuleRemote).(string)
	parsed.remoteAddress, parsed.remoteCIDR, parsed.remoteSecGrpID, err = inferRemoteSecurityGroup(parsed.remote)
	if err != nil {
		return nil, err
	}
	parsed.protocol = "all"

	if icmpInterface, ok := d.GetOk("icmp"); ok {
		haveType := false
		if icmpInterface.([]interface{})[0] == nil {
			return nil, fmt.Errorf("Internal error. icmp interface is nil")
		}
		icmp := icmpInterface.([]interface{})[0].(map[string]interface{})
		if value, ok := icmp["type"]; ok {
			parsed.icmpType = int64(value.(int))
			haveType = true
		}
		if value, ok := icmp["code"]; ok {
			if !haveType {
				return nil, fmt.Errorf("icmp code requires icmp type")
			}
			parsed.icmpCode = int64(value.(int))
		}
		parsed.protocol = "icmp"
	}
	for _, prot := range []string{"tcp", "udp"} {
		if tcpInterface, ok := d.GetOk(prot); ok {
			haveMin := false
			haveMax := false
			if tcpInterface.([]interface{})[0] == nil {
				return nil, fmt.Errorf("Internal error. %q interface is nil", prot)
			}
			ports := tcpInterface.([]interface{})[0].(map[string]interface{})
			if value, ok := ports["port_min"]; ok {
				parsed.portMin = int64(value.(int))
				haveMin = true
			}
			if value, ok := ports["port_max"]; ok {
				parsed.portMax = int64(value.(int))
				haveMax = true
			}

			// If only min or max is set, ensure that both min and max are set to the same value
			if haveMin && !haveMax {
				parsed.portMax = parsed.portMin
			}
			if haveMax && !haveMin {
				parsed.portMin = parsed.portMax
			}
			parsed.protocol = prot
		}
	}
	//	log.Printf("[DEBUG] parse tag=%s\n\t%v  \n\t%v  \n\t%v  \n\t%v  \n\t%v \n\t%v \n\t%v \n\t%v  \n\t%v  \n\t%v  \n\t%v  \n\t%v ",
	//		tag, parsed.secgrpID, parsed.ruleID, parsed.direction, parsed.ipversion, parsed.protocol, parsed.remoteAddress,
	//		parsed.remoteCIDR, parsed.remoteSecGrpID, parsed.icmpType, parsed.icmpCode, parsed.portMin, parsed.portMax)
	return parsed, nil
}

func makeTerraformRuleID(id1, id2 string) string {
	// Include both group and rule id to create a unique Terraform id.  As a bonus,
	// we can extract the group id as needed for API calls such as READ.
	return id1 + "." + id2
}

func extractRuleRemote(remote *models.SecurityGroupRuleRemote) (string, error) {
	if remote == nil {
		return "", fmt.Errorf("security group remote is nil")
	}
	if remote.Address != "" {
		return remote.Address, nil
	}
	if remote.CidrBlock != "" {
		return remote.CidrBlock, nil
	}
	if remote.ID.String() != "" {
		return remote.ID.String(), nil
	}
	return "", fmt.Errorf("security group remote is not set.")
}
