// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

const (
	isLBListenerPolicyRuleLBID             = "lb"
	isLBListenerPolicyRuleListenerID       = "listener"
	isLBListenerPolicyRulePolicyID         = "policy"
	isLBListenerPolicyRuleid               = "rule"
	isLBListenerPolicyRulecondition        = "condition"
	isLBListenerPolicyRuletype             = "type"
	isLBListenerPolicyRulevalue            = "value"
	isLBListenerPolicyRulefield            = "field"
	isLBListenerPolicyRuleStatus           = "provisioning_status"
	isLBListenerPolicyRuleAvailable        = "active"
	isLBListenerPolicyRuleFailed           = "failed"
	isLBListenerPolicyRulePending          = "pending"
	isLBListenerPolicyRuleDeleting         = "deleting"
	isLBListenerPolicyRuleDeleted          = "done"
	isLBListenerPolicyRuleRetry            = "retry"
	isLBListenerPolicyRuleProvisioning     = "provisioning"
	isLBListenerPolicyRuleProvisioningDone = "done"
)

func ResourceIBMISLBListenerPolicyRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISLBListenerPolicyRuleCreate,
		ReadContext:   resourceIBMISLBListenerPolicyRuleRead,
		UpdateContext: resourceIBMISLBListenerPolicyRuleUpdate,
		DeleteContext: resourceIBMISLBListenerPolicyRuleDelete,
		Exists:        resourceIBMISLBListenerPolicyRuleExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isLBListenerPolicyRuleLBID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Loadbalancer ID",
			},

			isLBListenerPolicyRuleListenerID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
					if o == "" {
						return false
					}
					// if state file entry and tf file entry matches
					if strings.Compare(n, o) == 0 {
						return true
					}

					if strings.Contains(n, "/") {
						//Split lbID/listenerID and fetch listenerID
						new := strings.Split(n, "/")
						if strings.Compare(new[1], o) == 0 {
							return true
						}
					}

					return false
				},
				Description: "Listener ID.",
			},

			isLBListenerPolicyRulePolicyID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
					if o == "" {
						return false
					}
					// if state file entry and tf file entry matches
					if strings.Compare(n, o) == 0 {
						return true
					}

					if strings.Contains(n, "/") {
						//Split lbID/listenerID and fetch listenerID
						new := strings.Split(n, "/")
						if strings.Compare(new[2], o) == 0 {
							return true
						}
					}

					return false
				},
				Description: "Listener Policy ID",
			},

			isLBListenerPolicyRulecondition: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_lb_listener_policy_rule", isLBListenerPolicyRulecondition),
				Description:  "Condition info of the rule.",
			},

			isLBListenerPolicyRuletype: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_lb_listener_policy_rule", isLBListenerPolicyRuletype),
				Description:  "Policy rule type.",
			},

			isLBListenerPolicyRulevalue: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.ValidateStringLength,
				Description:  "policy rule value info",
			},

			isLBListenerPolicyRulefield: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.ValidateStringLength,
			},

			isLBListenerPolicyRuleid: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isLBListenerPolicyStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			flex.RelatedCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the LB resource",
			},
		},
	}
}

func ResourceIBMISLBListenerPolicyRuleValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	condition := "contains, equals, matches_regex"
	ruletype := "header, hostname, path, body, query, sni_hostname"
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isLBListenerPolicyRulecondition,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              condition})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isLBListenerPolicyRuletype,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              ruletype})

	ibmISLBListenerPolicyRuleResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_lb_listener_policy_rule", Schema: validateSchema}
	return &ibmISLBListenerPolicyRuleResourceValidator
}

func resourceIBMISLBListenerPolicyRuleCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	//Read lb, listerner, policy IDs
	var field string
	lbID := d.Get(isLBListenerPolicyRuleLBID).(string)
	listenerID, err := getLbListenerID(d.Get(isLBListenerPolicyRuleListenerID).(string))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "create", "parse-listener_id").GetDiag()
	}

	policyID, err := getLbPolicyID(d.Get(isLBListenerPolicyRulePolicyID).(string))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "create", "parse-policy_id").GetDiag()
	}

	condition := d.Get(isLBListenerPolicyRulecondition).(string)
	ty := d.Get(isLBListenerPolicyRuletype).(string)
	value := d.Get(isLBListenerPolicyRulevalue).(string)
	if n, ok := d.GetOk(isLBListenerPolicyRulefield); ok {
		field = n.(string)
	}

	diagErr := lbListenerPolicyRuleCreate(context, d, meta, lbID, listenerID, policyID, condition, ty, value, field)
	if diagErr != nil {
		return diagErr
	}

	return resourceIBMISLBListenerPolicyRuleRead(context, d, meta)
}

func getLbListenerID(id string) (string, error) {
	if strings.Contains(id, "/") {
		parts, err := flex.IdParts(id)
		if err != nil {
			return "", err
		}

		return parts[1], nil
	} else {
		return id, nil
	}
}

func getLbPolicyID(id string) (string, error) {
	if strings.Contains(id, "/") {
		parts, err := flex.IdParts(id)
		if err != nil {
			return "", err
		}

		return parts[2], nil
	} else {
		return id, nil
	}
}

func vpcSdkClient(meta interface{}) (*vpcv1.VpcV1, error) {
	sess, err := meta.(conns.ClientSession).VpcV1API()
	return sess, err
}

func lbListenerPolicyRuleCreate(context context.Context, d *schema.ResourceData, meta interface{}, lbID, listenerID, policyID, condition, ty, value, field string) diag.Diagnostics {

	sess, err := vpcSdkClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &vpcv1.CreateLoadBalancerListenerPolicyRuleOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		PolicyID:       &policyID,
		Condition:      &condition,
		Type:           &ty,
		Value:          &value,
		Field:          &field,
	}

	isLBKey := "load_balancer_key_" + lbID
	conns.IbmMutexKV.Lock(isLBKey)
	defer conns.IbmMutexKV.Unlock(isLBKey)

	_, err = isWaitForLoadbalancerAvailable(sess, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLoadbalancerAvailable failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	rule, _, err := sess.CreateLoadBalancerListenerPolicyRuleWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateLoadBalancerListenerPolicyRuleWithContext failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s/%s", lbID, listenerID, policyID, *(rule.ID)))

	_, err = isWaitForLbListenerPolicyRuleAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLbListenerPolicyRuleAvailable failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return nil
}

func isWaitForLoadbalancerAvailable(vpc *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBListenerPolicyRulePending},
		Target:     []string{isLBProvisioningDone},
		Refresh:    isLoadbalancerRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLoadbalancerRefreshFunc(vpc *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getLbOptions := &vpcv1.GetLoadBalancerOptions{
			ID: &id,
		}

		lb, _, err := vpc.GetLoadBalancer(getLbOptions)
		if err != nil {
			return nil, "", err
		}

		if *(lb.ProvisioningStatus) == isLBListenerPolicyAvailable {
			return lb, isLBProvisioningDone, nil
		}

		return lb, isLBProvisioning, nil
	}
}

func isWaitForLbListenerPolicyRuleAvailable(vpc *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isLBListenerPolicyRuleProvisioning, "create_pending", "update_pending", "maintenance_pending"},
		Target:     []string{isLBListenerPolicyRuleProvisioningDone},
		Refresh:    isLbListenerPolicyRuleRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLbListenerPolicyRuleRefreshFunc(vpc *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		parts, err := flex.IdParts(id)
		if err != nil {
			return nil, "", err
		}

		lbID := parts[0]
		listenerID := parts[1]
		policyID := parts[2]
		ruleID := parts[3]

		getLbListenerPolicyRuleOptions := &vpcv1.GetLoadBalancerListenerPolicyRuleOptions{
			LoadBalancerID: &lbID,
			ListenerID:     &listenerID,
			PolicyID:       &policyID,
			ID:             &ruleID,
		}

		rule, _, err := vpc.GetLoadBalancerListenerPolicyRule(getLbListenerPolicyRuleOptions)

		if err != nil {
			return rule, "", err
		}

		if *rule.ProvisioningStatus == isLBListenerPolicyRuleAvailable || *rule.ProvisioningStatus == isLBListenerPolicyRuleFailed {
			return rule, isLBListenerPolicyRuleProvisioningDone, nil
		}

		return rule, *rule.ProvisioningStatus, nil
	}
}

func resourceIBMISLBListenerPolicyRuleRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	ID := d.Id()
	parts, err := flex.IdParts(ID)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "read", "sep-id-parts").GetDiag()
	}

	lbID := parts[0]
	listenerID := parts[1]
	policyID := parts[2]
	ruleID := parts[3]

	diagErr := lbListenerPolicyRuleGet(context, d, meta, lbID, listenerID, policyID, ruleID)
	if diagErr != nil {
		return diagErr
	}

	return nil
}

func resourceIBMISLBListenerPolicyRuleExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	ID := d.Id()

	exists, err := lbListenerPolicyRuleExists(d, meta, ID)
	return exists, err

}

func lbListenerPolicyRuleExists(d *schema.ResourceData, meta interface{}, ID string) (bool, error) {
	sess, err := vpcSdkClient(meta)
	if err != nil {
		return false, err
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return false, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "exists", "sep-id-parts")
	}
	if len(parts) != 4 {
		err = fmt.Errorf("[ERROR] Incorrect ID %s: ID should be a combination of lbID/listenerID/policyID/ruleID", d.Id())
		return false, flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "exists", "sep-id-parts")
	}
	lbID := parts[0]
	listenerID := parts[1]
	policyID := parts[2]
	ruleID := parts[3]

	getLbListenerPolicyRuleOptions := &vpcv1.GetLoadBalancerListenerPolicyRuleOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		PolicyID:       &policyID,
		ID:             &ruleID,
	}

	//Getting lb listener policy
	_, response, err := sess.GetLoadBalancerListenerPolicyRule(getLbListenerPolicyRuleOptions)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerListenerPolicyRule failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "exists")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}
func resourceIBMISLBListenerPolicyRuleUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "update", "sep-id-parts").GetDiag()
	}

	lbID := parts[0]
	listenerID := parts[1]
	policyID := parts[2]
	ruleID := parts[3]

	diagErr := lbListenerPolicyRuleUpdate(context, d, meta, lbID, listenerID, policyID, ruleID)
	if diagErr != nil {
		return diagErr
	}

	return resourceIBMISLBListenerPolicyRuleRead(context, d, meta)
}

func lbListenerPolicyRuleUpdate(context context.Context, d *schema.ResourceData, meta interface{}, lbID, listenerID, policyID, ID string) diag.Diagnostics {
	sess, err := vpcSdkClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	hasChanged := false
	updatePolicyRuleOptions := vpcv1.UpdateLoadBalancerListenerPolicyRuleOptions{}
	updatePolicyRuleOptions.LoadBalancerID = &lbID
	updatePolicyRuleOptions.ListenerID = &listenerID
	updatePolicyRuleOptions.PolicyID = &policyID
	updatePolicyRuleOptions.ID = &ID

	loadBalancerListenerPolicyRulePatchModel := &vpcv1.LoadBalancerListenerPolicyRulePatch{}

	if d.HasChange(isLBListenerPolicyRulecondition) {
		condition := d.Get(isLBListenerPolicyRulecondition).(string)
		loadBalancerListenerPolicyRulePatchModel.Condition = &condition
		hasChanged = true
	}

	if d.HasChange(isLBListenerPolicyRuletype) {
		ty := d.Get(isLBListenerPolicyRuletype).(string)
		loadBalancerListenerPolicyRulePatchModel.Type = &ty
		hasChanged = true
	}

	if d.HasChange(isLBListenerPolicyRulevalue) {
		value := d.Get(isLBListenerPolicyRulevalue).(string)
		loadBalancerListenerPolicyRulePatchModel.Value = &value
		hasChanged = true
	}

	if d.HasChange(isLBListenerPolicyRulefield) {
		field := d.Get(isLBListenerPolicyRulefield).(string)
		loadBalancerListenerPolicyRulePatchModel.Field = &field
		hasChanged = true
	}

	if hasChanged {
		loadBalancerListenerPolicyRulePatch, err := loadBalancerListenerPolicyRulePatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("loadBalancerListenerPolicyRulePatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updatePolicyRuleOptions.LoadBalancerListenerPolicyRulePatch = loadBalancerListenerPolicyRulePatch

		isLBKey := "load_balancer_key_" + lbID
		conns.IbmMutexKV.Lock(isLBKey)
		defer conns.IbmMutexKV.Unlock(isLBKey)

		_, err = isWaitForLoadbalancerAvailable(sess, lbID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLoadbalancerAvailable failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, _, err = sess.UpdateLoadBalancerListenerPolicyRuleWithContext(context, &updatePolicyRuleOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateLoadBalancerListenerPolicyRuleWithContext failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForLbListenerPolicyRuleAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLbListenerPolicyRuleAvailable failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMISLBListenerPolicyRuleDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	//Retrieve lbId, listenerId and policyID
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "delete", "sep-id-parts").GetDiag()
	}

	lbID := parts[0]
	listenerID := parts[1]
	policyID := parts[2]
	ruleID := parts[3]

	isLBKey := "load_balancer_key_" + lbID
	conns.IbmMutexKV.Lock(isLBKey)
	defer conns.IbmMutexKV.Unlock(isLBKey)

	diagErr := lbListenerPolicyRuleDelete(context, d, meta, lbID, listenerID, policyID, ruleID)
	if diagErr != nil {
		return diagErr
	}

	d.SetId("")
	return nil

}

func lbListenerPolicyRuleDelete(context context.Context, d *schema.ResourceData, meta interface{}, lbID, listenerID, policyID, ID string) diag.Diagnostics {

	sess, err := vpcSdkClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	//Getting rule optins
	getLbListenerPolicyRuleOptions := &vpcv1.GetLoadBalancerListenerPolicyRuleOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		PolicyID:       &policyID,
		ID:             &ID,
	}

	//Getting lb listener policy
	_, response, err := sess.GetLoadBalancerListenerPolicyRuleWithContext(context, getLbListenerPolicyRuleOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerListenerPolicyRuleWithContext failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteLbListenerPolicyRuleOptions := &vpcv1.DeleteLoadBalancerListenerPolicyRuleOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		PolicyID:       &policyID,
		ID:             &ID,
	}
	response, err = sess.DeleteLoadBalancerListenerPolicyRuleWithContext(context, deleteLbListenerPolicyRuleOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteLoadBalancerListenerPolicyRuleWithContext failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	_, err = isWaitForLbListnerPolicyRuleDeleted(sess, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForLbListnerPolicyRuleDeleted failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	return nil
}
func isWaitForLbListnerPolicyRuleDeleted(vpc *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBListenerPolicyRuleRetry, isLBListenerPolicyRuleDeleting},
		Target:     []string{isLBListenerPolicyRuleDeleted, isLBListenerPolicyRuleFailed},
		Refresh:    isLbListenerPolicyRuleDeleteRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLbListenerPolicyRuleDeleteRefreshFunc(vpc *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		//Retrieve lbId, listenerId and policyID
		parts, err := flex.IdParts(id)
		if err != nil {
			return nil, isLBListenerPolicyFailed, nil
		}

		lbID := parts[0]
		listenerID := parts[1]
		policyID := parts[2]
		ruleID := parts[3]

		getLbListenerPolicyRuleOptions := &vpcv1.GetLoadBalancerListenerPolicyRuleOptions{
			LoadBalancerID: &lbID,
			ListenerID:     &listenerID,
			PolicyID:       &policyID,
			ID:             &ruleID,
		}

		//Getting lb listener policy
		rule, response, err := vpc.GetLoadBalancerListenerPolicyRule(getLbListenerPolicyRuleOptions)

		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return rule, isLBListenerPolicyRuleDeleted, nil
			}
			return rule, isLBListenerPolicyRuleFailed, err
		}
		return nil, isLBListenerPolicyRuleDeleting, err
	}
}

func lbListenerPolicyRuleGet(context context.Context, d *schema.ResourceData, meta interface{}, lbID, listenerID, policyID, id string) diag.Diagnostics {

	sess, err := vpcSdkClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	//Getting rule optins
	getLbListenerPolicyRuleOptions := &vpcv1.GetLoadBalancerListenerPolicyRuleOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		PolicyID:       &policyID,
		ID:             &id,
	}

	//Getting lb listener policy
	loadBalancerListenerPolicyRule, response, err := sess.GetLoadBalancerListenerPolicyRuleWithContext(context, getLbListenerPolicyRuleOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerListenerPolicyRuleWithContext failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	//set the argument values
	if err = d.Set(isLBListenerPolicyRuleLBID, lbID); err != nil {
		err = fmt.Errorf("Error setting lb: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "read", "set-lb").GetDiag()
	}
	if err = d.Set(isLBListenerPolicyRuleListenerID, listenerID); err != nil {
		err = fmt.Errorf("Error setting listener: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "read", "set-listener").GetDiag()
	}
	if err = d.Set(isLBListenerPolicyRulePolicyID, policyID); err != nil {
		err = fmt.Errorf("Error setting policy: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "read", "set-policy").GetDiag()
	}
	if err = d.Set(isLBListenerPolicyRuleid, id); err != nil {
		err = fmt.Errorf("Error setting rule: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "read", "set-rule").GetDiag()
	}
	if err = d.Set("condition", loadBalancerListenerPolicyRule.Condition); err != nil {
		err = fmt.Errorf("Error setting condition: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "read", "set-condition").GetDiag()
	}
	if err = d.Set("type", loadBalancerListenerPolicyRule.Type); err != nil {
		err = fmt.Errorf("Error setting type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "read", "set-type").GetDiag()
	}
	if err = d.Set("value", loadBalancerListenerPolicyRule.Value); err != nil {
		err = fmt.Errorf("Error setting value: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "read", "set-value").GetDiag()
	}
	if !core.IsNil(loadBalancerListenerPolicyRule.Field) {
		if err = d.Set("field", loadBalancerListenerPolicyRule.Field); err != nil {
			err = fmt.Errorf("Error setting field: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "read", "set-field").GetDiag()
		}
	}
	if err = d.Set("provisioning_status", loadBalancerListenerPolicyRule.ProvisioningStatus); err != nil {
		err = fmt.Errorf("Error setting provisioning_status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "read", "set-provisioning_status").GetDiag()
	}
	getLoadBalancerOptions := &vpcv1.GetLoadBalancerOptions{
		ID: &lbID,
	}
	lb, response, err := sess.GetLoadBalancerWithContext(context, getLoadBalancerOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetLoadBalancerWithContext failed: %s", err.Error()), "ibm_is_lb_listener_policy_rule", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(flex.RelatedCRN, *lb.CRN); err != nil {
		err = fmt.Errorf("Error setting provisioning_status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_lb_listener_policy_rule", "read", "set-provisioning_status").GetDiag()
	}

	return nil
}
