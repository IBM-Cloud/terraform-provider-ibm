// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"bytes"
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
				Optional:    true,
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
				Type:        schema.TypeSet,
				Computed:    true,
				Optional:    true,
				Description: "Security Group Rules",
				Set:         securityGroupRuleHash,
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
				Computed:    true,
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

	if _, ok := d.GetOk(isSecurityGroupRules); ok {
		err = cleanExistingDefaultSecurityGroupRules(d, sess, sg)
		if err != nil {
			return err
		}
		err = addNewDefaultSecurityGroupRules(d, sess, sg)
		if err != nil {
			return err
		}
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
	return resourceIBMISVPCDefaultSecurityGroupUpdate(d, meta)
}

func cleanExistingDefaultSecurityGroupRules(d *schema.ResourceData, sess *vpcv1.VpcV1, sg *vpcv1.SecurityGroup) error {
	id := d.Id()
	log.Printf("[INFO] UJJK inside cleanup")

	for _, ruleIntf := range sg.Rules {

		log.Printf("[INFO] UJJK removiing rule %v", ruleIntf)
		switch rule := ruleIntf.(type) {
		case *vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll:
			{
				removeSgRuleOptions := &vpcv1.DeleteSecurityGroupRuleOptions{
					SecurityGroupID: &id,
					ID:              rule.ID,
				}
				res, err := sess.DeleteSecurityGroupRule(removeSgRuleOptions)
				if err != nil {
					return fmt.Errorf("[ERROR] Error while removing default security group rule %s/%s", res, err)
				}
			}
		case *vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp:
			{
				removeSgRuleOptions := &vpcv1.DeleteSecurityGroupRuleOptions{
					SecurityGroupID: &id,
					ID:              rule.ID,
				}
				res, err := sess.DeleteSecurityGroupRule(removeSgRuleOptions)
				if err != nil {
					return fmt.Errorf("[ERROR] Error while removing default security group rule %s/%s", res, err)
				}
			}
		case *vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp:
			{
				removeSgRuleOptions := &vpcv1.DeleteSecurityGroupRuleOptions{
					SecurityGroupID: &id,
					ID:              rule.ID,
				}
				res, err := sess.DeleteSecurityGroupRule(removeSgRuleOptions)
				if err != nil {
					return fmt.Errorf("[ERROR] Error while removing default security group rule %s/%s", res, err)
				}
			}
		case *vpcv1.SecurityGroupRule:
			{
				removeSgRuleOptions := &vpcv1.DeleteSecurityGroupRuleOptions{
					SecurityGroupID: &id,
					ID:              rule.ID,
				}
				res, err := sess.DeleteSecurityGroupRule(removeSgRuleOptions)
				if err != nil {
					return fmt.Errorf("[ERROR] Error while removing default security group rule %s/%s", res, err)
				}
			}
		}
	}
	return nil
}
func addNewDefaultSecurityGroupRules(d *schema.ResourceData, sess *vpcv1.VpcV1, sg *vpcv1.SecurityGroup) error {
	id := d.Id()
	log.Printf("[INFO] UJJK inside add")

	if rulesIntf, ok := d.GetOk(isSecurityGroupRules); ok {
		rules := rulesIntf.(*schema.Set).List()
		for _, rule := range rules {
			log.Printf("[INFO] UJJK adding new rule %v", rule)
			rulex := rule.(map[string]interface{})
			direction := rulex[isSecurityGroupRuleDirection].(string)
			direction = strings.ToLower(direction)
			remote := rulex[isSecurityGroupRuleRemote].(string)
			remote = strings.ToLower(direction)
			ruletype := rulex[isSecurityGroupRuleType].(int)
			ruletype64 := int64(ruletype)
			code := rulex[isSecurityGroupRuleCode].(int)
			code64 := int64(code)
			portmin := rulex[isSecurityGroupRulePortMin].(int)
			portmin64 := int64(portmin)

			portmax := rulex[isSecurityGroupRulePortMax].(int)
			portmax64 := int64(portmax)

			protocol := rulex[isSecurityGroupRuleProtocol].(string)
			protocol = strings.ToLower(protocol)

			isSecurityGroupRuleKey := "security_group_rule_key_" + id
			conns.IbmMutexKV.Lock(isSecurityGroupRuleKey)
			defer conns.IbmMutexKV.Unlock(isSecurityGroupRuleKey)

			sgTemplate := &vpcv1.SecurityGroupRulePrototype{}
			if direction != "" {
				sgTemplate.Direction = &direction
			}
			if remote != "" {
				remoteTemplate := &vpcv1.SecurityGroupRuleRemotePrototype{}
				if validate.IsSecurityGroupAddress(remote) {
					remoteTemplate.Address = &remote
				} else if validate.IsSecurityGroupCIDR(remote) {
					remoteTemplate.CIDRBlock = &remote
				} else {
					remoteTemplate.ID = &remote
				}
			}
			if direction != "" {
				sgTemplate.Direction = &direction
			}
			if protocol == "icmp" {
				sgTemplate.Type = &ruletype64
				sgTemplate.Code = &code64
			}
			sgTemplate.PortMin = &portmin64
			sgTemplate.PortMax = &portmax64
			sgTemplate.Protocol = &protocol

			options := &vpcv1.CreateSecurityGroupRuleOptions{
				SecurityGroupID:            &id,
				SecurityGroupRulePrototype: sgTemplate,
			}

			log.Printf("[INFO] UJJK created sg create options %v", options)
			_, response, err := sess.CreateSecurityGroupRule(options)
			if err != nil {
				return fmt.Errorf("[ERROR] Error while creating Security Group Rule %s\n%s", err, response)
			}
			log.Printf("[INFO] UJJK  rule created %v", response)
		}

	}
	// parsed, sgTemplate, _, err := parseIBMISDefaultSecurityGroupRuleDictionary(d, "create", sess)
	// if err != nil {
	// 	return err
	// }
	// isSecurityGroupRuleKey := "security_group_rule_key_" + parsed.secgrpID
	// conns.IbmMutexKV.Lock(isSecurityGroupRuleKey)
	// defer conns.IbmMutexKV.Unlock(isSecurityGroupRuleKey)

	// options := &vpcv1.CreateSecurityGroupRuleOptions{
	// 	SecurityGroupID:            &parsed.secgrpID,
	// 	SecurityGroupRulePrototype: sgTemplate,
	// }

	// _, response, err := sess.CreateSecurityGroupRule(options)
	// if err != nil {
	// 	return fmt.Errorf("[ERROR] Error while creating Security Group Rule %s\n%s", err, response)
	// }

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
	log.Printf("[INFO] UJJK inside update block")
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	id := d.Id()
	name := ""

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

	log.Printf("[INFO] UJJK no change in sg rules")

	if d.HasChange(isSecurityGroupRules) {
		log.Printf("[INFO] UJJK ther eis a change in sg rules")
		sgId := d.Id()
		log.Printf("[INFO] UJJK trying to get sg ")

		getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
			ID: &sgId,
		}

		log.Printf("[INFO] UJJK getting sg")
		sg, response, err := sess.GetSecurityGroup(getSecurityGroupOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error while getting default Security Group of the vpc %s\n%s", err, response)
		}
		log.Printf("[INFO] UJJK sg got %v", sg)

		log.Printf("[INFO] UJJK starting cleanup ")
		err = cleanExistingDefaultSecurityGroupRules(d, sess, sg)
		if err != nil {
			return err
		}
		log.Printf("[INFO] UJJK now adding")
		err = addNewDefaultSecurityGroupRules(d, sess, sg)
		if err != nil {
			return err
		}
	}
	if d.HasChange(isSecurityGroupName) {
		name = d.Get(isSecurityGroupName).(string)
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
	return resourceIBMISVPCDefaultSecurityGroupRead(d, meta)
}

func resourceIBMISVPCDefaultSecurityGroupDelete(d *schema.ResourceData, meta interface{}) error {

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

func securityGroupRuleHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[isSecurityGroupRuleProtocol].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[isSecurityGroupRuleDirection].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[isSecurityGroupRuleRemote].(string)))
	buf.WriteString(fmt.Sprintf("%d-", m[isSecurityGroupRuleCode].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m[isSecurityGroupRulePortMin].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m[isSecurityGroupRulePortMax].(int)))
	return conns.String(buf.String())
}

func parseIBMISDefaultSecurityGroupRuleDictionary(d *schema.ResourceData, tag string, sess *vpcv1.VpcV1) (*parsedIBMISSecurityGroupRuleDictionary, *vpcv1.SecurityGroupRulePrototype, *vpcv1.UpdateSecurityGroupRuleOptions, error) {
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
		parsed.secgrpID = d.Get(isVPCDefaultSecurityGroup).(string)
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
	//	log.Printf("[DEBUG] parse tag=%s\n\t%v  \n\t%v  \n\t%v  \n\t%v  \n\t%v \n\t%v \n\t%v \n\t%v  \n\t%v  \n\t%v  \n\t%v  \n\t%v ",
	//		tag, parsed.secgrpID, parsed.ruleID, parsed.direction, parsed.ipversion, parsed.protocol, parsed.remoteAddress,
	//		parsed.remoteCIDR, parsed.remoteSecGrpID, parsed.icmpType, parsed.icmpCode, parsed.portMin, parsed.portMax)
	return parsed, sgTemplate, sgTemplateUpdate, nil
}