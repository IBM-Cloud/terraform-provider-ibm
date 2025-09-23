// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isSecurityGroupName          = "name"
	isSecurityGroupVPC           = "vpc"
	isSecurityGroupRules         = "rules"
	isSecurityGroupResourceGroup = "resource_group"
	isSecurityGroupTags          = "tags"
	isSecurityGroupAccessTags    = "access_tags"
	isSecurityGroupCRN           = "crn"
)

func ResourceIBMISSecurityGroup() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceIBMISSecurityGroupCreate,
		ReadContext:   resourceIBMISSecurityGroupRead,
		UpdateContext: resourceIBMISSecurityGroupUpdate,
		DeleteContext: resourceIBMISSecurityGroupDelete,
		Exists:        resourceIBMISSecurityGroupExists,
		Importer:      &schema.ResourceImporter{},

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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Security group name",
				ValidateFunc: validate.InvokeValidator("ibm_is_security_group", isSecurityGroupName),
			},
			isSecurityGroupVPC: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Security group's resource group id",
				ForceNew:    true,
			},

			isSecurityGroupTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_security_group", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags",
			},

			isSecurityGroupAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_security_group", "accesstag")},
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
				Description: "Security Rules",
				Elem: &schema.Resource{
					Schema: makeIBMISSecurityRuleSchema(),
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

func ResourceIBMISSecurityGroupValidator() *validate.ResourceValidator {

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

	ibmISSecurityGroupResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_security_group", Schema: validateSchema}
	return &ibmISSecurityGroupResourceValidator
}

func resourceIBMISSecurityGroupCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	vpc := d.Get(isSecurityGroupVPC).(string)

	createSecurityGroupOptions := &vpcv1.CreateSecurityGroupOptions{
		VPC: &vpcv1.VPCIdentity{
			ID: &vpc,
		},
	}
	var rg, name string
	if grp, ok := d.GetOk(isSecurityGroupResourceGroup); ok {
		rg = grp.(string)
		createSecurityGroupOptions.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	if nm, ok := d.GetOk(isSecurityGroupName); ok {
		name = nm.(string)
		createSecurityGroupOptions.Name = &name
	}
	sg, _, err := sess.CreateSecurityGroup(createSecurityGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSecurityGroupWithContext failed: %s", err.Error()), "ibm_is_security_group", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*sg.ID)
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
	return resourceIBMISSecurityGroupRead(context, d, meta)
}

func resourceIBMISSecurityGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	id := d.Id()

	getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
		ID: &id,
	}
	securityGroup, response, err := sess.GetSecurityGroupWithContext(context, getSecurityGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecurityGroupWithContext failed: %s", err.Error()), "ibm_is_security_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *securityGroup.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"Error getting Security Group tags : %s\n%s", d.Id(), err)
	}
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *securityGroup.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of Security Group (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set(isSecurityGroupTags, tags); err != nil {
		err = fmt.Errorf("Error setting tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "read", "set-tags").GetDiag()
	}
	if err = d.Set(isSecurityGroupAccessTags, accesstags); err != nil {
		err = fmt.Errorf("Error setting access_tags: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "read", "set-access_tags").GetDiag()
	}
	if err = d.Set("crn", securityGroup.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "read", "set-crn").GetDiag()
	}
	if !core.IsNil(securityGroup.Name) {
		if err = d.Set("name", securityGroup.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(securityGroup.VPC) {
		if err = d.Set(isSecurityGroupVPC, *securityGroup.VPC.ID); err != nil {
			err = fmt.Errorf("Error setting vpc: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "read", "set-vpc").GetDiag()
		}
	}
	rules := make([]map[string]interface{}, 0)
	if len(securityGroup.Rules) > 0 {
		for _, rule := range securityGroup.Rules {
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
					local, ok := rule.Local.(*vpcv1.SecurityGroupRuleLocal)
					if ok {
						if local != nil && reflect.ValueOf(local).IsNil() == false {
							if local.Address != nil {
								r[isSecurityGroupRuleLocal] = local.Address
							} else if local.CIDRBlock != nil {
								r[isSecurityGroupRuleLocal] = local.CIDRBlock
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
					local, ok := rule.Local.(*vpcv1.SecurityGroupRuleLocal)
					if ok {
						if local != nil && reflect.ValueOf(local).IsNil() == false {
							if local.Address != nil {
								r[isSecurityGroupRuleLocal] = local.Address
							} else if local.CIDRBlock != nil {
								r[isSecurityGroupRuleLocal] = local.CIDRBlock
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
					local, ok := rule.Local.(*vpcv1.SecurityGroupRuleLocal)
					if ok {
						if local != nil && reflect.ValueOf(local).IsNil() == false {
							if local.Address != nil {
								r[isSecurityGroupRuleLocal] = local.Address
							} else if local.CIDRBlock != nil {
								r[isSecurityGroupRuleLocal] = local.CIDRBlock
							}
						}
					}
					rules = append(rules, r)
				}
			}
		}
	}
	if err = d.Set(isSecurityGroupRules, rules); err != nil {
		err = fmt.Errorf("Error setting rules: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "read", "set-rules").GetDiag()
	}

	d.SetId(*securityGroup.ID)
	if securityGroup.ResourceGroup != nil {
		if err = d.Set(isSecurityGroupResourceGroup, securityGroup.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "read", "set-resource_group").GetDiag()
		}
		if err = d.Set(flex.ResourceGroupName, securityGroup.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "read", "set-resource_group_name").GetDiag()
		}
	}
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBaseController failed: %s", err.Error()), "ibm_is_security_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/network/securityGroups"); err != nil {
		err = fmt.Errorf("Error setting resource_controller_url: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "read", "set-resource_controller_url").GetDiag()
	}
	if err = d.Set(flex.ResourceName, *securityGroup.Name); err != nil {
		err = fmt.Errorf("Error setting resource_name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "read", "set-resource_name").GetDiag()
	}
	if err = d.Set(flex.ResourceCRN, *securityGroup.CRN); err != nil {
		err = fmt.Errorf("Error setting resource_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "read", "set-resource_crn").GetDiag()
	}
	return nil
}

func resourceIBMISSecurityGroupUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
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
		return resourceIBMISSecurityGroupRead(context, d, meta)
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
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("securityGroupPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_security_group", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateSecurityGroupOptions.SecurityGroupPatch = securityGroupPatch
		_, _, err = sess.UpdateSecurityGroupWithContext(context, updateSecurityGroupOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSecurityGroupWithContext failed: %s", err.Error()), "ibm_is_security_group", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIBMISSecurityGroupRead(context, d, meta)
}

func resourceIBMISSecurityGroupDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	id := d.Id()

	getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
		ID: &id,
	}
	_, response, err := sess.GetSecurityGroupWithContext(context, getSecurityGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecurityGroupWithContext failed: %s", err.Error()), "ibm_is_security_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	start := ""
	allrecs := []vpcv1.SecurityGroupTargetReferenceIntf{}

	for {
		listSecurityGroupTargetsOptions := sess.NewListSecurityGroupTargetsOptions(id)

		groups, _, err := sess.ListSecurityGroupTargetsWithContext(context, listSecurityGroupTargetsOptions)
		if err != nil || groups == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListSecurityGroupTargetsWithContext failed: %s", err.Error()), "ibm_is_security_group", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
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
				response, err = sess.DeleteSecurityGroupTargetBindingWithContext(context, deleteSecurityGroupTargetBindingOptions)
				if err != nil {
					if response != nil {
						if response.StatusCode == 404 {
							log.Printf("[DEBUG] Security group target(%s) binding is already deleted", *securityGroupTargetReference.ID)
						} else if response.StatusCode == 409 {
							log.Printf("[DEBUG] Security group target(%s) binding is in deleting status, waiting till target is removed", *securityGroupTargetReference.ID)
							_, err = isWaitForTargetDeleted(sess, id, *securityGroupTargetReference.ID, securityGroupTargetReferenceIntf, d.Timeout(schema.TimeoutDelete))
							if err != nil {
								tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForTargetDeleted failed: %s", err.Error()), "ibm_is_security_group", "delete")
								log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
								return tfErr.GetDiag()
							}
							_, err := isWaitForSecurityGroupTargetDeleteRetry(sess, deleteSecurityGroupTargetBindingOptions, d.Timeout(schema.TimeoutDelete))
							if err != nil {
								tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForSecurityGroupTargetDeleteRetry failed: %s", err.Error()), "ibm_is_security_group", "delete")
								log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
								return tfErr.GetDiag()
							}
						}
					} else {
						tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSecurityGroupTargetBindingWithContext failed: %s", err.Error()), "ibm_is_security_group", "delete")
						log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
						return tfErr.GetDiag()
					}
				}

			}
		}
	}

	deleteSecurityGroupOptions := &vpcv1.DeleteSecurityGroupOptions{
		ID: &id,
	}
	response, err = sess.DeleteSecurityGroupWithContext(context, deleteSecurityGroupOptions)

	if err != nil {
		if response != nil {
			if response.StatusCode == 404 {
				log.Printf("[DEBUG] Security group(%s) target bindings are already deleted", id)
			} else if response.StatusCode == 409 {
				log.Printf("[DEBUG] Security group(%s) has target bindings is in deleting, will wait till target is removed", id)
				_, err = isWaitForSgCleanup(sess, id, allrecs, d.Timeout(schema.TimeoutDelete))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForSgCleanup failed: %s", err.Error()), "ibm_is_security_group", "delete")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
				_, err := isWaitForSecurityGroupDeleteRetry(sess, deleteSecurityGroupOptions, d.Timeout(schema.TimeoutDelete))
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForSecurityGroupDeleteRetry failed: %s", err.Error()), "ibm_is_security_group", "delete")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return tfErr.GetDiag()
				}
			}
		} else {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSecurityGroupWithContext failed: %s", err.Error()), "ibm_is_security_group", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	d.SetId("")
	return nil
}

func resourceIBMISSecurityGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_security_group", "exists", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
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
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecurityGroup failed: %s", err.Error()), "ibm_is_security_group", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}

func makeIBMISSecurityRuleSchema() map[string]*schema.Schema {
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

		isSecurityGroupRuleLocal: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Security group local ip: an IP address, a CIDR block",
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

func isWaitForTargetDeleted(client *vpcv1.VpcV1, sgId, targetId string, target vpcv1.SecurityGroupTargetReferenceIntf, timeout time.Duration) (interface{}, error) {
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

func isTargetRefreshFunc(client *vpcv1.VpcV1, sgId, targetId string, target vpcv1.SecurityGroupTargetReferenceIntf) resource.StateRefreshFunc {
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
func isWaitForSgCleanup(client *vpcv1.VpcV1, sgId string, targets []vpcv1.SecurityGroupTargetReferenceIntf, timeout time.Duration) (interface{}, error) {
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

func isSgRefreshFunc(client *vpcv1.VpcV1, sgId string, groups []vpcv1.SecurityGroupTargetReferenceIntf) resource.StateRefreshFunc {
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

func isWaitForSecurityGroupDeleteRetry(vpcClient *vpcv1.VpcV1, deleteSecurityGroupOptions *vpcv1.DeleteSecurityGroupOptions, timeout time.Duration) (interface{}, error) {
	log.Printf("[DEBUG] Retrying security group (%s) delete", *deleteSecurityGroupOptions.ID)
	stateConf := &resource.StateChangeConf{
		Pending: []string{"security-group-in-use"},
		Target:  []string{"deleted", ""},
		Refresh: func() (interface{}, string, error) {
			log.Printf("[DEBUG] Retrying security group (%s) delete", *deleteSecurityGroupOptions.ID)
			response, err := vpcClient.DeleteSecurityGroup(deleteSecurityGroupOptions)
			if err != nil {
				if response != nil && response.StatusCode == 409 {
					return response, "security-group-in-use", nil
				} else if response != nil && response.StatusCode == 404 {
					return response, "deleted", nil
				}
				return response, "", fmt.Errorf("[ERROR] Error deleting security group: %s\n%s", err, response)
			}
			return response, "deleted", nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isWaitForSecurityGroupTargetDeleteRetry(vpcClient *vpcv1.VpcV1, deleteSecurityGroupTargetBindingOptions *vpcv1.DeleteSecurityGroupTargetBindingOptions, timeout time.Duration) (interface{}, error) {
	log.Printf("[DEBUG] Retrying security group target (%s) delete", *deleteSecurityGroupTargetBindingOptions.ID)
	stateConf := &resource.StateChangeConf{
		Pending: []string{"security-group-target-in-use"},
		Target:  []string{"deleted", ""},
		Refresh: func() (interface{}, string, error) {
			log.Printf("[DEBUG] Retrying security group target(%s) delete", *deleteSecurityGroupTargetBindingOptions.ID)
			response, err := vpcClient.DeleteSecurityGroupTargetBinding(deleteSecurityGroupTargetBindingOptions)
			if err != nil {
				if response != nil && response.StatusCode == 409 {
					return response, "security-group-target-in-use", nil
				} else if response != nil && response.StatusCode == 404 {
					return response, "deleted", nil
				}
				return response, "", fmt.Errorf("[ERROR] Error deleting security group target: %s\n%s", err, response)
			}
			return response, "deleted", nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}
