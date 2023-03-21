// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const ()

func ResourceIBMISVPCDefaultSecurityGroup() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMISVPCDefaultSecurityGroupCreate,
		Read:     resourceIBMISVPCDefaultSecurityGroupRead,
		Update:   resourceIBMISVPCDefaultSecurityGroupUpdate,
		Delete:   resourceIBMISVPCDefaultSecurityGroupDelete,
		Exists:   resourceIBMISVPCDefaultSecurityGroupExists,
		Importer: &schema.ResourceImporter{},

		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceTagsCustomizeDiff(diff)
				},
			),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isSecurityGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Security group name",
			},
			isVPCDefaultSecurityGroup: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Security group id",
			},

			isSecurityGroupVPC: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Security group's resource group id",
			},

			isSecurityGroupTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_vpc_default_security_group", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags",
			},

			isSecurityGroupAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_vpc_default_security_group", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},

			isSecurityGroupCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			isSecurityGroupRules: {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "Security Group Rules",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isSecurityGroupRuleDirection: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Direction of traffic to enforce, either inbound or outbound",
						},

						isSecurityGroupRuleIPVersion: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "IP version: ipv4",
						},

						isSecurityGroupRuleRemote: {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Security group id: an IP address, a CIDR block, or a single security group identifier",
						},

						isSecurityGroupRuleType: {
							Optional: true,
							Type:     schema.TypeInt,
							Computed: true,
						},

						isSecurityGroupRuleCode: {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						isSecurityGroupRulePortMin: {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						isSecurityGroupRulePortMax: {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						isSecurityGroupRuleProtocol: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			isSecurityGroupResourceGroup: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Resource Group ID",
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
		},
	}
}

func ResourceIBMISVPCDefaultSecurityGroupValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isSecurityGroupName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "accesstag",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISSecurityGroupResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_vpc_default_security_group", Schema: validateSchema}
	return &ibmISSecurityGroupResourceValidator
}

func resourceIBMISVPCDefaultSecurityGroupCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	sgId := d.Get(isVPCDefaultSecurityGroup).(string)

	getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
		ID: &sgId,
	}

	sg, response, err := sess.GetSecurityGroup(getSecurityGroupOptions)
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting default Security Group of the vpc %s\n%s", err, response)
	}
	d.SetId(*sg.ID)

	err = cleanExistingDefaultSecurityGroupRules(d, sess, sg)
	if err != nil {
		return err
	}
	err = addNewDefaultSecurityGroupRules(d, sess, sg)
	if err != nil {
		return err
	}
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isSecurityGroupTags); ok || v != "" {
		oldList, newList := d.GetChange(isSecurityGroupTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *sg.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error while creating Security Group tags : %s\n%s", *sg.ID, err)
		}
	}
	if _, ok := d.GetOk(isSecurityGroupAccessTags); ok {
		oldList, newList := d.GetChange(isSecurityGroupAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *sg.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of Security Group (%s) access tags: %s", d.Id(), err)
		}
	}
	return resourceIBMISSecurityGroupRead(d, meta)
}

func cleanExistingDefaultSecurityGroupRules(d *schema.ResourceData, sess *vpcv1.VpcV1, sg *vpcv1.SecurityGroup) error {
	id := d.Id()

	for _, ruleIntf := range sg.Rules {
		rule := ruleIntf.(*vpcv1.SecurityGroupRule)
		removeSgRuleOptions := &vpcv1.DeleteSecurityGroupRuleOptions{
			SecurityGroupID: &id,
			ID:              rule.ID,
		}
		res, err := sess.DeleteSecurityGroupRule(removeSgRuleOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error while removing default security group rule %s/%s", res, err)
		}
	}
	return nil
}
func addNewDefaultSecurityGroupRules(d *schema.ResourceData, sess *vpcv1.VpcV1, sg *vpcv1.SecurityGroup) error {
	id := d.Id()

	var rules []interface{}
	if rulesIntf, ok := d.GetOk(isSecurityGroupRules); ok {
		rules = rulesIntf.([]interface{})
		for _, rule := range rules {
			rulex := rule.(map[string]interface{})
			direction := rulex[isSecurityGroupRuleDirection].(string)
			direction = strings.ToLower(direction)
			remote := rulex[isSecurityGroupRuleRemote].(string)
			remote = strings.ToLower(direction)
			ruletype := rulex[isSecurityGroupRuleType].(string)
			ruletype = strings.ToLower(direction)
			code := rulex[isSecurityGroupRuleCode].(string)
			code = strings.ToLower(direction)
			portmin := rulex[isSecurityGroupRulePortMin].(string)
			portmin = strings.ToLower(direction)
			portmax := rulex[isSecurityGroupRulePortMax].(string)
			portmax = strings.ToLower(direction)
			protocol := rulex[isSecurityGroupRuleProtocol].(string)
			protocol = strings.ToLower(direction)

			isSecurityGroupRuleKey := "security_group_rule_key_" + id
			conns.IbmMutexKV.Lock(isSecurityGroupRuleKey)
			defer conns.IbmMutexKV.Unlock(isSecurityGroupRuleKey)

			parsed := &parsedIBMISSecurityGroupRuleDictionary{}
			sgTemplate := &vpcv1.SecurityGroupRulePrototype{}
			sgTemplateUpdate := &vpcv1.UpdateSecurityGroupRuleOptions{}
			var err error
			parsed.icmpType = -1
			parsed.icmpCode = -1
			parsed.portMin = -1
			parsed.portMax = -1

			parsed.secgrpID, parsed.ruleID, err = parseISTerraformID(d.Id())
			if err != nil {
				parsed.secgrpID = d.Get(isSecurityGroupID).(string)
			} else {
				sgTemplateUpdate.SecurityGroupID = &parsed.secgrpID
				sgTemplateUpdate.ID = &parsed.ruleID
			}

			securityGroupRulePatchModel := &vpcv1.SecurityGroupRulePatch{}

			parsed.direction = d.Get(isSecurityGroupRuleDirection).(string)
			sgTemplate.Direction = &parsed.direction
			securityGroupRulePatchModel.Direction = &parsed.direction

			if version, ok := d.GetOk(isSecurityGroupRuleIPVersion); ok {
				parsed.ipversion = version.(string)
				sgTemplate.IPVersion = &parsed.ipversion
				securityGroupRulePatchModel.IPVersion = &parsed.ipversion
			} else {
				parsed.ipversion = "IPv4"
				sgTemplate.IPVersion = &parsed.ipversion
				securityGroupRulePatchModel.IPVersion = &parsed.ipversion
			}

			parsed.remote = ""
			if pr, ok := d.GetOk(isSecurityGroupRuleRemote); ok {
				parsed.remote = pr.(string)
			}
			parsed.remoteAddress = ""
			parsed.remoteCIDR = ""
			parsed.remoteSecGrpID = ""
			err = nil
			if parsed.remote != "" {
				parsed.remoteAddress, parsed.remoteCIDR, parsed.remoteSecGrpID, err = inferRemoteSecurityGroup(parsed.remote)
				remoteTemplate := &vpcv1.SecurityGroupRuleRemotePrototype{}
				remoteTemplateUpdate := &vpcv1.SecurityGroupRuleRemotePatch{}
				if parsed.remoteAddress != "" {
					remoteTemplate.Address = &parsed.remoteAddress
					remoteTemplateUpdate.Address = &parsed.remoteAddress
				} else if parsed.remoteCIDR != "" {
					remoteTemplate.CIDRBlock = &parsed.remoteCIDR
					remoteTemplateUpdate.CIDRBlock = &parsed.remoteCIDR
				} else if parsed.remoteSecGrpID != "" {
					remoteTemplate.ID = &parsed.remoteSecGrpID
					remoteTemplateUpdate.ID = &parsed.remoteSecGrpID

					// check if remote is actually a SG identifier
					getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
						ID: &parsed.remoteSecGrpID,
					}
					sg, res, err := sess.GetSecurityGroup(getSecurityGroupOptions)
					if err != nil || sg == nil {
						if res != nil && res.StatusCode == 404 {
							return nil, nil, nil, err
						}
						return nil, nil, nil, fmt.Errorf("Error getting Security Group in remote (%s): %s\n%s", parsed.remoteSecGrpID, err, res)
					}
				}
				sgTemplate.Remote = remoteTemplate
				securityGroupRulePatchModel.Remote = remoteTemplateUpdate
			}
			if err != nil {
				return nil, nil, nil, err
			}
			parsed.protocol = "all"

			if icmpInterface, ok := d.GetOk("icmp"); ok {
				if icmpInterface.([]interface{})[0] != nil {
					haveType := false
					icmp := icmpInterface.([]interface{})[0].(map[string]interface{})
					if value, ok := icmp["type"]; ok {
						parsed.icmpType = int64(value.(int))
						haveType = true
					}
					if value, ok := icmp["code"]; ok {
						if !haveType {
							return nil, nil, nil, fmt.Errorf("icmp code requires icmp type")
						}
						parsed.icmpCode = int64(value.(int))
					}
				}
				parsed.protocol = "icmp"
				if icmpInterface.([]interface{})[0] == nil {
					parsed.icmpType = 0
					parsed.icmpCode = 0
				} else {
					sgTemplate.Type = &parsed.icmpType
					sgTemplate.Code = &parsed.icmpCode
				}
				sgTemplate.Protocol = &parsed.protocol
				securityGroupRulePatchModel.Type = &parsed.icmpType
				securityGroupRulePatchModel.Code = &parsed.icmpCode
			}
			for _, prot := range []string{"tcp", "udp"} {
				if tcpInterface, ok := d.GetOk(prot); ok {
					if tcpInterface.([]interface{})[0] != nil {
						haveMin := false
						haveMax := false
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
					}
					parsed.protocol = prot
					sgTemplate.Protocol = &parsed.protocol
					if tcpInterface.([]interface{})[0] == nil {
						parsed.portMax = 65535
						parsed.portMin = 1
					}
					sgTemplate.PortMax = &parsed.portMax
					sgTemplate.PortMin = &parsed.portMin
					securityGroupRulePatchModel.PortMax = &parsed.portMax
					securityGroupRulePatchModel.PortMin = &parsed.portMin
				}
			}
			if parsed.protocol == "all" {
				sgTemplate.Protocol = &parsed.protocol
			}
			securityGroupRulePatch, err := securityGroupRulePatchModel.AsPatch()
			if err != nil {
				return nil, nil, nil, fmt.Errorf("[ERROR] Error calling asPatch for SecurityGroupRulePatch: %s", err)
			}
			sgTemplateUpdate.SecurityGroupRulePatch = securityGroupRulePatch

			options := &vpcv1.CreateSecurityGroupRuleOptions{
				SecurityGroupID:            &id,
				SecurityGroupRulePrototype: sgTemplate,
			}

			rule, response, err := sess.CreateSecurityGroupRule(options)
			if err != nil {
				return fmt.Errorf("[ERROR] Error while creating Security Group Rule %s\n%s", err, response)
			}
		}

	}
	parsed, sgTemplate, _, err := parseIBMISSecurityGroupRuleDictionary(d, "create", sess)
	if err != nil {
		return err
	}
	isSecurityGroupRuleKey := "security_group_rule_key_" + parsed.secgrpID
	conns.IbmMutexKV.Lock(isSecurityGroupRuleKey)
	defer conns.IbmMutexKV.Unlock(isSecurityGroupRuleKey)

	options := &vpcv1.CreateSecurityGroupRuleOptions{
		SecurityGroupID:            &parsed.secgrpID,
		SecurityGroupRulePrototype: sgTemplate,
	}

	rule, response, err := sess.CreateSecurityGroupRule(options)
	if err != nil {
		return fmt.Errorf("[ERROR] Error while creating Security Group Rule %s\n%s", err, response)
	}
	switch reflect.TypeOf(rule).String() {
	case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp":
		{
			sgrule := rule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp)
			d.Set(isSecurityGroupRuleID, *sgrule.ID)
			tfID := makeTerraformRuleID(parsed.secgrpID, *sgrule.ID)
			d.SetId(tfID)
		}
	case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll":
		{
			sgrule := rule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll)
			d.Set(isSecurityGroupRuleID, *sgrule.ID)
			tfID := makeTerraformRuleID(parsed.secgrpID, *sgrule.ID)
			d.SetId(tfID)
		}
	case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp":
		{
			sgrule := rule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp)
			d.Set(isSecurityGroupRuleID, *sgrule.ID)
			tfID := makeTerraformRuleID(parsed.secgrpID, *sgrule.ID)
			d.SetId(tfID)
		}
	}
	return nil
}
func resourceIBMISVPCDefaultSecurityGroupRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	id := d.Id()

	getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
		ID: &id,
	}
	group, response, err := sess.GetSecurityGroup(getSecurityGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Error getting Security Group : %s\n%s", err, response)
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *group.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"Error getting Security Group tags : %s\n%s", d.Id(), err)
	}
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *group.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of Security Group (%s) access tags: %s", d.Id(), err)
	}
	d.Set(isSecurityGroupTags, tags)
	d.Set(isSecurityGroupAccessTags, accesstags)
	d.Set(isSecurityGroupCRN, *group.CRN)
	d.Set(isSecurityGroupName, *group.Name)
	d.Set(isSecurityGroupVPC, *group.VPC.ID)
	rules := make([]map[string]interface{}, 0)
	if len(group.Rules) > 0 {
		for _, rule := range group.Rules {
			switch reflect.TypeOf(rule).String() {
			case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp":
				{
					rule := rule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp)
					r := make(map[string]interface{})
					if rule.Code != nil {
						r[isSecurityGroupRuleCode] = int(*rule.Code)
					}
					if rule.Type != nil {
						r[isSecurityGroupRuleType] = int(*rule.Type)
					}
					r[isSecurityGroupRuleDirection] = *rule.Direction
					r[isSecurityGroupRuleIPVersion] = *rule.IPVersion
					if rule.Protocol != nil {
						r[isSecurityGroupRuleProtocol] = *rule.Protocol
					}
					remote, ok := rule.Remote.(*vpcv1.SecurityGroupRuleRemote)
					if ok {
						if remote != nil && reflect.ValueOf(remote).IsNil() == false {
							if remote.ID != nil {
								r[isSecurityGroupRuleRemote] = remote.ID
							} else if remote.Address != nil {
								r[isSecurityGroupRuleRemote] = remote.Address
							} else if remote.CIDRBlock != nil {
								r[isSecurityGroupRuleRemote] = remote.CIDRBlock
							}
						}
					}
					rules = append(rules, r)
				}
			case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll":
				{
					rule := rule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll)
					r := make(map[string]interface{})
					r[isSecurityGroupRuleDirection] = *rule.Direction
					r[isSecurityGroupRuleIPVersion] = *rule.IPVersion
					if rule.Protocol != nil {
						r[isSecurityGroupRuleProtocol] = *rule.Protocol
					}
					remote, ok := rule.Remote.(*vpcv1.SecurityGroupRuleRemote)
					if ok {
						if remote != nil && reflect.ValueOf(remote).IsNil() == false {
							if remote.ID != nil {
								r[isSecurityGroupRuleRemote] = remote.ID
							} else if remote.Address != nil {
								r[isSecurityGroupRuleRemote] = remote.Address
							} else if remote.CIDRBlock != nil {
								r[isSecurityGroupRuleRemote] = remote.CIDRBlock
							}
						}
					}
					rules = append(rules, r)
				}
			case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp":
				{
					rule := rule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp)
					r := make(map[string]interface{})
					if rule.PortMin != nil {
						r[isSecurityGroupRulePortMin] = int(*rule.PortMin)
					}
					if rule.PortMax != nil {
						r[isSecurityGroupRulePortMax] = int(*rule.PortMax)
					}
					r[isSecurityGroupRuleDirection] = *rule.Direction
					r[isSecurityGroupRuleIPVersion] = *rule.IPVersion
					if rule.Protocol != nil {
						r[isSecurityGroupRuleProtocol] = *rule.Protocol
					}
					remote, ok := rule.Remote.(*vpcv1.SecurityGroupRuleRemote)
					if ok {
						if remote != nil && reflect.ValueOf(remote).IsNil() == false {
							if remote.ID != nil {
								r[isSecurityGroupRuleRemote] = remote.ID
							} else if remote.Address != nil {
								r[isSecurityGroupRuleRemote] = remote.Address
							} else if remote.CIDRBlock != nil {
								r[isSecurityGroupRuleRemote] = remote.CIDRBlock
							}
						}
					}
					rules = append(rules, r)
				}
			}
		}
	}
	d.Set(isSecurityGroupRules, rules)
	d.SetId(*group.ID)
	if group.ResourceGroup != nil {
		d.Set(isSecurityGroupResourceGroup, group.ResourceGroup.ID)
		d.Set(flex.ResourceGroupName, group.ResourceGroup.Name)
	}
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/securityGroups")
	d.Set(flex.ResourceName, *group.Name)
	d.Set(flex.ResourceCRN, *group.CRN)
	return nil
}

func resourceIBMISVPCDefaultSecurityGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	id := d.Id()
	name := ""
	hasChanged := false

	if d.HasChange(isSecurityGroupTags) {
		oldList, newList := d.GetChange(isSecurityGroupTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get(isSecurityGroupCRN).(string), "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error Updating Security Group tags: %s\n%s", d.Id(), err)
		}
	}
	if d.HasChange(isSecurityGroupAccessTags) {
		oldList, newList := d.GetChange(isSecurityGroupAccessTags)
		err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, d.Get(isSecurityGroupCRN).(string), "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of Security Group (%s) access tags: %s", d.Id(), err)
		}
	}
	if d.HasChange(isSecurityGroupName) {
		name = d.Get(isSecurityGroupName).(string)
		hasChanged = true
	} else {
		return resourceIBMISSecurityGroupRead(d, meta)
	}

	if hasChanged {
		updateSecurityGroupOptions := &vpcv1.UpdateSecurityGroupOptions{
			ID: &id,
		}
		securityGroupPatchModel := &vpcv1.SecurityGroupPatch{
			Name: &name,
		}
		securityGroupPatch, err := securityGroupPatchModel.AsPatch()
		if err != nil {
			return fmt.Errorf("[ERROR] Error calling asPatch for SecurityGroupPatch: %s", err)
		}
		updateSecurityGroupOptions.SecurityGroupPatch = securityGroupPatch
		_, response, err := sess.UpdateSecurityGroup(updateSecurityGroupOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Updating Security Group : %s\n%s", err, response)
		}
	}
	return resourceIBMISSecurityGroupRead(d, meta)
}

func resourceIBMISVPCDefaultSecurityGroupDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	id := d.Id()

	getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
		ID: &id,
	}
	_, response, err := sess.GetSecurityGroup(getSecurityGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Error Getting Security Group (%s): %s\n%s", id, err, response)
	}

	start := ""
	allrecs := []vpcv1.SecurityGroupTargetReferenceIntf{}

	for {
		listSecurityGroupTargetsOptions := sess.NewListSecurityGroupTargetsOptions(id)

		groups, response, err := sess.ListSecurityGroupTargets(listSecurityGroupTargetsOptions)
		if err != nil || groups == nil {
			return fmt.Errorf("[ERROR] Error Getting Security Group Targets %s\n%s", err, response)
		}
		if *groups.TotalCount == int64(0) {
			break
		}

		start = flex.GetNext(groups.Next)
		allrecs = append(allrecs, groups.Targets...)

		if start == "" {
			break
		}

	}

	for _, securityGroupTargetReferenceIntf := range allrecs {
		if securityGroupTargetReferenceIntf != nil {
			securityGroupTargetReference := securityGroupTargetReferenceIntf.(*vpcv1.SecurityGroupTargetReference)
			if securityGroupTargetReference != nil && securityGroupTargetReference.ID != nil {

				deleteSecurityGroupTargetBindingOptions := sess.NewDeleteSecurityGroupTargetBindingOptions(id, *securityGroupTargetReference.ID)
				response, err = sess.DeleteSecurityGroupTargetBinding(deleteSecurityGroupTargetBindingOptions)
				if err != nil {
					if response != nil {
						if response.StatusCode == 404 {
							log.Printf("[DEBUG] Security group target(%s) binding is already deleted", *securityGroupTargetReference.ID)
						} else if response.StatusCode == 409 {
							log.Printf("[DEBUG] Security group target(%s) binding is in deleting status, waiting till target is removed", *securityGroupTargetReference.ID)
							_, err = isWaitForTargetDeleted(sess, id, *securityGroupTargetReference.ID, securityGroupTargetReferenceIntf, d.Timeout(schema.TimeoutDelete))
							if err != nil {
								return err
							}
						}
					} else {
						return fmt.Errorf("[ERROR] Error deleting security group target binding while deleting security group : %s\n%s", err, response)
					}
				}

			}
		}
	}

	deleteSecurityGroupOptions := &vpcv1.DeleteSecurityGroupOptions{
		ID: &id,
	}
	response, err = sess.DeleteSecurityGroup(deleteSecurityGroupOptions)

	if err != nil {
		if response != nil {
			if response.StatusCode == 404 {
				log.Printf("[DEBUG] Security group(%s) target bindings are already deleted", id)
			} else if response.StatusCode == 409 {
				log.Printf("[DEBUG] Security group(%s) has target bindings is in deleting, will wait till target is removed", id)
				_, err = isWaitForSgCleanup(sess, id, allrecs, d.Timeout(schema.TimeoutDelete))
				if err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf("[ERROR] Error Deleting Security Group : %s\n%s", err, response)
		}
	}
	d.SetId("")
	return nil
}

func resourceIBMISVPCDefaultSecurityGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	id := d.Id()

	getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
		ID: &id,
	}
	_, response, err := sess.GetSecurityGroup(getSecurityGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("[ERROR] Error getting Security Group: %s\n%s", err, response)
	}
	return true, nil
}

func makeIBMISVPCDefaultSecurityRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{

		isSecurityGroupRuleDirection: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Direction of traffic to enforce, either inbound or outbound",
		},

		isSecurityGroupRuleIPVersion: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "IP version: ipv4",
		},

		isSecurityGroupRuleRemote: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Security group id: an IP address, a CIDR block, or a single security group identifier",
		},

		isSecurityGroupRuleType: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isSecurityGroupRuleCode: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isSecurityGroupRulePortMin: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isSecurityGroupRulePortMax: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isSecurityGroupRuleProtocol: {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func isWaitForVPCDefaultTargetDeleted(client *vpcv1.VpcV1, sgId, targetId string, target vpcv1.SecurityGroupTargetReferenceIntf, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Security group(%s) target(%s) to be deleted.", sgId, targetId)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting"},
		Target:     []string{"done", ""},
		Refresh:    isTargetRefreshFunc(client, sgId, targetId, target),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isTargetVPCDefaultRefreshFunc(client *vpcv1.VpcV1, sgId, targetId string, target vpcv1.SecurityGroupTargetReferenceIntf) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		targetgetoptions := &vpcv1.GetSecurityGroupTargetOptions{
			SecurityGroupID: &sgId,
			ID:              &targetId,
		}
		sgTarget, response, err := client.GetSecurityGroupTarget(targetgetoptions)
		if err != nil {
			return target, "", fmt.Errorf("[ERROR] Error getting target(%s): %s\n%s", targetId, err, response)
		}
		if response != nil && response.StatusCode == 404 {
			return target, "done", nil
		}
		return sgTarget, "deleting", nil
	}
}
func isWaitForVPCDefaultSgCleanup(client *vpcv1.VpcV1, sgId string, targets []vpcv1.SecurityGroupTargetReferenceIntf, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Security group(%s) target(%s) to be deleted.", sgId, targets)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting"},
		Target:     []string{"done", ""},
		Refresh:    isSgRefreshFunc(client, sgId, targets),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVPCDefaultSgRefreshFunc(client *vpcv1.VpcV1, sgId string, groups []vpcv1.SecurityGroupTargetReferenceIntf) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		start := ""
		allrecs := []vpcv1.SecurityGroupTargetReferenceIntf{}
		for {
			listSecurityGroupTargetsOptions := client.NewListSecurityGroupTargetsOptions(sgId)

			sggroups, response, err := client.ListSecurityGroupTargets(listSecurityGroupTargetsOptions)
			if err != nil || sggroups == nil {
				return groups, "", fmt.Errorf("[ERROR] Error Getting Security Group Targets %s\n%s", err, response)
			}
			if *sggroups.TotalCount == int64(0) {
				return groups, "done", nil
			}

			start = flex.GetNext(sggroups.Next)
			allrecs = append(allrecs, sggroups.Targets...)

			if start == "" {
				break
			}
		}
		return allrecs, "deleting", nil
	}
}
