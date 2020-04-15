package ibm

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isLBListenerPolicyLBID                 = "lb"
	isLBListenerPolicyListenerID           = "listener"
	isLBListenerPolicyAction               = "action"
	isLBListenerPolicyPriority             = "priority"
	isLBListenerPolicyName                 = "name"
	isLBListenerPolicyID                   = "policy_id"
	isLBListenerPolicyRules                = "rules"
	isLBListenerPolicyRulesInfo            = "rule_info"
	isLBListenerPolicyTargetID             = "target_id"
	isLBListenerPolicyTargetHTTPStatusCode = "target_http_status_code"
	isLBListenerPolicyTargetURL            = "target_url"
	isLBListenerPolicyStatus               = "provisioning_status"
	isLBListenerPolicyRuleID               = "rule_id"
	isLBListenerPolicyAvailable            = "active"
	isLBListenerPolicyFailed               = "failed"
	isLBListenerPolicyPending              = "pending"
	isLBListenerPolicyDeleting             = "deleting"
	isLBListenerPolicyDeleted              = "done"
	isLBListenerPolicyRetry                = "retry"
	isLBListenerPolicyRuleCondition        = "condition"
	isLBListenerPolicyRuleType             = "type"
	isLBListenerPolicyRuleValue            = "value"
	isLBListenerPolicyRuleField            = "field"
	isLBListenerPolicyProvisioning         = "provisioning"
	isLBListenerPolicyProvisioningDone     = "done"
)

func resourceIBMISLBListenerPolicy() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISLBListenerPolicyCreate,
		Read:     resourceIBMISLBListenerPolicyRead,
		Update:   resourceIBMISLBListenerPolicyUpdate,
		Delete:   resourceIBMISLBListenerPolicyDelete,
		Exists:   resourceIBMISLBListenerPolicyExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isLBListenerPolicyLBID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isLBListenerPolicyListenerID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isLBListenerPolicyAction: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue([]string{"forward", "redirect", "reject"}),
			},

			isLBListenerPolicyPriority: {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateLBListenerPolicyPriority,
			},

			isLBListenerPolicyName: {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     false,
				Computed:     true,
				ValidateFunc: validateISName,
			},

			isLBListenerPolicyID: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isLBListenerPolicyRules: {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isLBListenerPolicyRuleCondition: {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateAllowedStringValue([]string{"contains", "equals", "matches_regex"}),
							Description:  "Condition of the rule",
						},

						isLBListenerPolicyRuleType: {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateAllowedStringValue([]string{"header", "hostname", "path"}),
							Description:  "Type of the rule",
						},

						isLBListenerPolicyRuleValue: {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateStringLength,
							Description:  "Value to be matched for rule condition",
						},

						isLBListenerPolicyRuleField: {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateStringLength,
							Description:  "HTTP header field. This is only applicable to rule type.",
						},

						isLBListenerPolicyRuleID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rule ID",
						},
					},
				},
			},

			isLBListenerPolicyTargetID: {
				Type:     schema.TypeString,
				ForceNew: false,
				Optional: true,
			},

			isLBListenerPolicyTargetHTTPStatusCode: {
				Type:     schema.TypeInt,
				ForceNew: false,
				Optional: true,
			},

			isLBListenerPolicyTargetURL: {
				Type:     schema.TypeString,
				ForceNew: false,
				Optional: true,
			},

			isLBListenerPolicyStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMISLBListenerPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	//Get the Load balancer ID
	lbID := d.Get(isLBListenerPolicyLBID).(string)
	listenerID := d.Get(isLBListenerPolicyListenerID).(string)
	action := d.Get(isLBListenerPolicyAction).(string)
	priority := int64(d.Get(isLBListenerPolicyPriority).(int))

	//user-defined name for this policy.
	var name string
	if n, ok := d.GetOk(isLBListenerPolicyName); ok {
		name = n.(string)
	}

	if userDetails.generation == 1 {
		err := classicLbListenerPolicyCreate(d, meta, lbID, listenerID, action, name, priority)
		if err != nil {
			return err
		}
	} else {
		err := lbListenerPolicyCreate(d, meta, lbID, listenerID, action, name, priority)
		if err != nil {
			return err
		}
	}

	return resourceIBMISLBListenerPolicyRead(d, meta)
}

func classicLbListenerPolicyCreate(d *schema.ResourceData, meta interface{}, lbID, listenerID, action, name string, priority int64) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	// When `action` is `forward`, `LoadBalancerPoolIdentity` is required to specify which
	// pool the load balancer forwards the traffic to. When `action` is `redirect`,
	// `LoadBalancerListenerPolicyRedirectURLPrototype` is required to specify the url and
	// http status code used in the redirect response.

	actionChk := d.Get(isLBListenerPolicyAction)
	tID, targetIDSet := d.GetOk(isLBListenerPolicyTargetID)
	statusCode, statusSet := d.GetOk(isLBListenerPolicyTargetHTTPStatusCode)
	url, urlSet := d.GetOk(isLBListenerPolicyTargetURL)

	var target vpcclassicv1.LoadBalancerListenerPolicyPrototypeTargetIntf

	if actionChk.(string) == "forward" {
		if targetIDSet {
			id := tID.(string)
			target = &vpcclassicv1.LoadBalancerListenerPolicyPrototypeTargetLoadBalancerPoolIdentity{
				ID: &id,
			}
		} else {
			return fmt.Errorf("When action is forward please specify target_id")
		}
	} else if actionChk.(string) == "redirect" {

		urlPrototype := vpcclassicv1.LoadBalancerListenerPolicyPrototypeTargetLoadBalancerListenerPolicyRedirectURLPrototype{}

		if statusSet {
			sc := int64(statusCode.(int))
			urlPrototype.HttpStatusCode = &sc
		} else {
			return fmt.Errorf("When action is redirect please specify target_http_status_code")
		}

		if urlSet {
			link := url.(string)
			urlPrototype.URL = &link
		} else {
			return fmt.Errorf("When action is redirect please specify target_url")
		}

		target = &urlPrototype
	}

	rulesInfo := make([]vpcclassicv1.LoadBalancerListenerPolicyRulePrototype, 0)
	if rules, rulesSet := d.GetOk(isLBListenerPolicyRules); rulesSet {
		policyRules := rules.([]interface{})
		for _, rule := range policyRules {
			rulex := rule.(map[string]interface{})

			//condition, type and value are mandatory params
			var condition string
			if rulex[isLBListenerPolicyRuleCondition] != nil {
				condition = rulex[isLBListenerPolicyRuleCondition].(string)
			}

			var ty string
			if rulex[isLBListenerPolicyRuleType] != nil {
				ty = rulex[isLBListenerPolicyRuleType].(string)
			}

			var value string
			if rulex[isLBListenerPolicyRuleValue] != nil {
				value = rulex[isLBListenerPolicyRuleValue].(string)
			}

			field := rulex[isLBListenerPolicyRuleField].(string)

			r := vpcclassicv1.LoadBalancerListenerPolicyRulePrototype{
				Condition: &condition,
				Field:     &field,
				Type:      &ty,
				Value:     &value,
			}

			rulesInfo = append(rulesInfo, r)
		}
	}

	options := &vpcclassicv1.CreateLoadBalancerListenerPolicyOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		Action:         &action,
		Priority:       &priority,
		Name:           &name,
		Target:         target,
		Rules:          rulesInfo,
	}

	isLBListenerPolicyKey := "load_balancer_listener_policy_key_" + lbID + listenerID
	ibmMutexKV.Lock(isLBListenerPolicyKey)
	defer ibmMutexKV.Unlock(isLBListenerPolicyKey)

	_, err = isWaitForClassicLbAvailable(sess, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf(
			"LB-LP Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	policy, response, err := sess.CreateLoadBalancerListenerPolicy(options)
	if err != nil {
		return fmt.Errorf("Error while creating lb listener policy for LB %s: Error %v Response %v", lbID, err, *response)
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", lbID, listenerID, *(policy.ID)))

	_, err = isWaitForClassicLbListenerPolicyAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	return nil
}

func isWaitForClassicLbAvailable(vpc *vpcclassicv1.VpcClassicV1, id string, timeout time.Duration) (interface{}, error) {

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isLBListenerProvisioning, "create_pending", "update_pending", "maintenance_pending"},
		Target:     []string{isLBProvisioningDone},
		Refresh:    isLbClassicRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLbClassicRefreshFunc(vpc *vpcclassicv1.VpcClassicV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getLbOptions := &vpcclassicv1.GetLoadBalancerOptions{
			ID: &id,
		}

		lb, _, err := vpc.GetLoadBalancer(getLbOptions)
		if err != nil {
			return nil, "", err
		}

		if *(lb.ProvisioningStatus) == isLBListenerPolicyAvailable || *lb.ProvisioningStatus == isLBListenerPolicyFailed {
			return lb, isLBProvisioningDone, nil
		}

		return lb, isLBProvisioning, nil
	}
}

func isWaitForClassicLbListenerPolicyAvailable(vpc *vpcclassicv1.VpcClassicV1, id string, timeout time.Duration) (interface{}, error) {

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isLBListenerProvisioning, "create_pending", "update_pending", "maintenance_pending"},
		Target:     []string{isLBListenerPolicyProvisioningDone},
		Refresh:    isLbListenerPolicyClassicRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLbListenerPolicyClassicRefreshFunc(vpc *vpcclassicv1.VpcClassicV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		parts, err := idParts(id)
		if err != nil {
			return nil, "", err
		}

		lbID := parts[0]
		listenerID := parts[1]
		policyID := parts[2]

		getLbListenerPolicyOptions := &vpcclassicv1.GetLoadBalancerListenerPolicyOptions{
			LoadBalancerID: &lbID,
			ListenerID:     &listenerID,
			ID:             &policyID,
		}

		policy, _, err := vpc.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

		if err != nil {
			return policy, "", err
		}

		if *policy.ProvisioningStatus == isLBListenerPolicyAvailable || *policy.ProvisioningStatus == isLBListenerPolicyFailed {
			return policy, isLBListenerProvisioningDone, nil
		}

		return policy, *policy.ProvisioningStatus, nil
	}
}

func vpcClient(meta interface{}) (*vpcv1.VpcV1, error) {
	sess, err := meta.(ClientSession).VpcV1API()
	return sess, err
}

func lbListenerPolicyCreate(d *schema.ResourceData, meta interface{}, lbID, listenerID, action, name string, priority int64) error {

	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	// When `action` is `forward`, `LoadBalancerPoolIdentity` is required to specify which
	// pool the load balancer forwards the traffic to. When `action` is `redirect`,
	// `LoadBalancerListenerPolicyRedirectURLPrototype` is required to specify the url and
	// http status code used in the redirect response.
	actionChk := d.Get(isLBListenerPolicyAction)
	tID, targetIDSet := d.GetOk(isLBListenerPolicyTargetID)
	statusCode, statusSet := d.GetOk(isLBListenerPolicyTargetHTTPStatusCode)
	url, urlSet := d.GetOk(isLBListenerPolicyTargetURL)

	var target vpcv1.LoadBalancerListenerPolicyPrototypeTargetIntf

	if actionChk.(string) == "forward" {
		if targetIDSet {
			id := tID.(string)
			target = &vpcv1.LoadBalancerListenerPolicyPrototypeTargetLoadBalancerPoolIdentity{
				ID: &id,
			}
		} else {
			return fmt.Errorf("When action is forward please specify target_id")
		}
	} else if actionChk.(string) == "redirect" {

		urlPrototype := vpcv1.LoadBalancerListenerPolicyPrototypeTargetLoadBalancerListenerPolicyRedirectURLPrototype{}

		if statusSet {
			sc := int64(statusCode.(int))
			urlPrototype.HttpStatusCode = &sc
		} else {
			return fmt.Errorf("When action is redirect please specify target_http_status_code")
		}

		if urlSet {
			link := url.(string)
			urlPrototype.URL = &link
		} else {
			return fmt.Errorf("When action is redirect please specify target_url")
		}

		target = &urlPrototype
	}

	//Read Rules
	rulesInfo := make([]vpcv1.LoadBalancerListenerPolicyRulePrototype, 0)
	if rules, rulesSet := d.GetOk(isLBListenerPolicyRules); rulesSet {
		policyRules := rules.([]interface{})
		for _, rule := range policyRules {
			rulex := rule.(map[string]interface{})

			//condition, type and value are mandatory params
			var condition string
			if rulex[isLBListenerPolicyRuleCondition] != nil {
				condition = rulex[isLBListenerPolicyRuleCondition].(string)
			}

			var ty string
			if rulex[isLBListenerPolicyRuleType] != nil {
				ty = rulex[isLBListenerPolicyRuleType].(string)
			}

			var value string
			if rulex[isLBListenerPolicyRuleValue] != nil {
				value = rulex[isLBListenerPolicyRuleValue].(string)
			}

			field := rulex[isLBListenerPolicyRuleField].(string)

			r := vpcv1.LoadBalancerListenerPolicyRulePrototype{
				Condition: &condition,
				Field:     &field,
				Type:      &ty,
				Value:     &value,
			}

			rulesInfo = append(rulesInfo, r)
		}
	}

	options := &vpcv1.CreateLoadBalancerListenerPolicyOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		Action:         &action,
		Priority:       &priority,
		Target:         target,
		Name:           &name,
		Rules:          rulesInfo,
	}

	isLBListenerPolicyKey := "load_balancer_listener_policy_key_" + lbID + listenerID
	ibmMutexKV.Lock(isLBListenerPolicyKey)
	defer ibmMutexKV.Unlock(isLBListenerPolicyKey)

	_, err = isWaitForLbAvailable(sess, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf(
			"LB-LP Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	policy, response, err := sess.CreateLoadBalancerListenerPolicy(options)
	if err != nil {
		return fmt.Errorf("Error while creating lb listener policy for LB %s: Error %v Response %v", lbID, err, *response)
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", lbID, listenerID, *(policy.ID)))

	_, err = isWaitForLbListenerPolicyAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}
	return nil
}

func isWaitForLbAvailable(vpc *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBListenerPolicyPending},
		Target:     []string{isLBProvisioningDone},
		Refresh:    isLbRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLbRefreshFunc(vpc *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
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

func isWaitForLbListenerPolicyAvailable(vpc *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isLBListenerProvisioning, "create_pending", "update_pending", "maintenance_pending"},
		Target:     []string{isLBListenerProvisioningDone},
		Refresh:    isLbListenerPolicyRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLbListenerPolicyRefreshFunc(vpc *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		parts, err := idParts(id)
		if err != nil {
			return nil, "", err
		}

		lbID := parts[0]
		listenerID := parts[1]
		policyID := parts[2]

		getLbListenerPolicyOptions := &vpcv1.GetLoadBalancerListenerPolicyOptions{
			LoadBalancerID: &lbID,
			ListenerID:     &listenerID,
			ID:             &policyID,
		}

		policy, _, err := vpc.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

		if err != nil {
			return policy, "", err
		}

		if *policy.ProvisioningStatus == isLBListenerPolicyAvailable || *policy.ProvisioningStatus == isLBListenerPolicyFailed {
			return policy, isLBListenerProvisioningDone, nil
		}

		return policy, *policy.ProvisioningStatus, nil
	}
}

func resourceIBMISLBListenerPolicyRead(d *schema.ResourceData, meta interface{}) error {

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	ID := d.Id()
	parts, err := idParts(ID)
	if err != nil {
		return err
	}

	lbID := parts[0]
	listenerID := parts[1]
	policyID := parts[2]

	if userDetails.generation == 1 {
		err := classicLbListenerPolicyGet(d, meta, lbID, listenerID, policyID)
		if err != nil {
			return err
		}
	} else {
		err := lbListenerPolicyGet(d, meta, lbID, listenerID, policyID)
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceIBMISLBListenerPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	ID := d.Id()
	if userDetails.generation == 1 {
		err := classicLbListenerPolicyExists(d, meta, ID)
		if err != nil {
			return false, err
		}
	} else {
		err := lbListenerPolicyExists(d, meta, ID)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func classicLbListenerPolicyExists(d *schema.ResourceData, meta interface{}, ID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	//Retrieve lbID, listenerID and policyID
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	listenerID := parts[1]
	policyID := parts[2]

	//populate lblistenerpolicyOPtions
	getLbListenerPolicyOptions := &vpcclassicv1.GetLoadBalancerListenerPolicyOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		ID:             &policyID,
	}

	//Getting lb listener policy
	_, response, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

	if err != nil {
		if response.StatusCode != 404 {
			return nil
		}
		return err
	}

	return nil
}

func lbListenerPolicyExists(d *schema.ResourceData, meta interface{}, ID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	listenerID := parts[1]
	policyID := parts[2]

	getLbListenerPolicyOptions := &vpcv1.GetLoadBalancerListenerPolicyOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		ID:             &policyID,
	}

	//Getting lb listener policy
	_, response, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

	if err != nil {
		if response.StatusCode != 404 {
			return nil
		}
		return err
	}
	return nil
}
func resourceIBMISLBListenerPolicyUpdate(d *schema.ResourceData, meta interface{}) error {

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	listenerID := parts[1]
	policyID := parts[2]

	if userDetails.generation == 1 {

		err := classicLbListenerPolicyUpdate(d, meta, lbID, listenerID, policyID)
		if err != nil {
			return err
		}
	} else {

		err := lbListenerPolicyUpdate(d, meta, lbID, listenerID, policyID)
		if err != nil {
			return err
		}
	}

	return resourceIBMISLBListenerRead(d, meta)
}

func classicLbListenerPolicyUpdate(d *schema.ResourceData, meta interface{}, lbID, listenerID, ID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	hasChanged := false
	updatePolicyOptions := vpcclassicv1.UpdateLoadBalancerListenerPolicyOptions{}
	updatePolicyOptions.LoadBalancerID = &lbID
	updatePolicyOptions.ListenerID = &listenerID
	updatePolicyOptions.ID = &ID

	if d.HasChange(isLBListenerPolicyName) {
		policy := d.Get(isLBListenerPolicyName).(string)
		updatePolicyOptions.Name = &policy
		hasChanged = true
	}

	if d.HasChange(isLBListenerPolicyPriority) {
		prio := d.Get(isLBListenerPolicyPriority).(int)
		priority := int64(prio)
		updatePolicyOptions.Priority = &priority
		hasChanged = true
	}

	var target vpcclassicv1.LoadBalancerListenerPolicyPatchTargetIntf

	//If Action is forward and TargetID is changed, set the target to pool ID
	if d.Get(isLBListenerPolicyAction).(string) == "forward" && d.HasChange(isLBListenerPolicyTargetID) {
		id := d.Get(isLBListenerPolicyTargetID).(string)

		target = &vpcclassicv1.LoadBalancerListenerPolicyPatchTargetLoadBalancerPoolIdentity{
			ID: &id,
		}

		updatePolicyOptions.Target = target
		hasChanged = true
	} else if d.Get(isLBListenerPolicyAction).(string) == "redirect" {
		//if Action is redirect and either status code or URL chnaged, set accordingly
		//LoadBalancerListenerPolicyPatchTargetLoadBalancerListenerPolicyRedirectURLPatch

		redirectPatch := vpcclassicv1.LoadBalancerListenerPolicyPatchTargetLoadBalancerListenerPolicyRedirectURLPatch{}

		targetChange := false
		if d.HasChange(isLBListenerPolicyTargetHTTPStatusCode) {
			status := d.Get(isLBListenerPolicyTargetHTTPStatusCode).(int)
			sc := int64(status)
			redirectPatch.HttpStatusCode = &sc
			hasChanged = true
			targetChange = true
		}

		if d.HasChange(isLBListenerPolicyTargetURL) {
			url := d.Get(isLBListenerPolicyTargetURL).(string)
			redirectPatch.URL = &url
			hasChanged = true
			targetChange = true
		}

		//Update the target only if there is a change in either statusCode or URL
		if targetChange {
			target = &redirectPatch
			updatePolicyOptions.Target = target
		}
	}

	isLBListenerPolicyKey := "load_balancer_listener_policy_key_" + lbID + listenerID
	ibmMutexKV.Lock(isLBListenerPolicyKey)
	defer ibmMutexKV.Unlock(isLBListenerPolicyKey)

	if hasChanged {

		_, err = isWaitForClassicLbAvailable(sess, lbID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return fmt.Errorf(
				"LB-LP Error checking for load balancer (%s) is active: %s", lbID, err)
		}
		_, response, err := sess.UpdateLoadBalancerListenerPolicy(&updatePolicyOptions)
		if err != nil {
			return fmt.Errorf("Error Getting Instance: %s\n%s", err, response)
		}

		_, err = isWaitForClassicLbListenerPolicyAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return err
		}
	}
	return nil
}

func lbListenerPolicyUpdate(d *schema.ResourceData, meta interface{}, lbID, listenerID, ID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	hasChanged := false
	updatePolicyOptions := vpcv1.UpdateLoadBalancerListenerPolicyOptions{}
	updatePolicyOptions.LoadBalancerID = &lbID
	updatePolicyOptions.ListenerID = &listenerID
	updatePolicyOptions.ID = &ID

	if d.HasChange(isLBListenerPolicyName) {
		policy := d.Get(isLBListenerPolicyName).(string)
		updatePolicyOptions.Name = &policy
		hasChanged = true
	}

	if d.HasChange(isLBListenerPolicyPriority) {
		prio := d.Get(isLBListenerPolicyPriority).(int)
		priority := int64(prio)
		updatePolicyOptions.Priority = &priority
		hasChanged = true
	}

	var target vpcv1.LoadBalancerListenerPolicyPatchTargetIntf
	//If Action is forward and TargetID is changed, set the target to pool ID
	if d.Get(isLBListenerPolicyAction).(string) == "forward" && d.HasChange(isLBListenerPolicyTargetID) {
		id := d.Get(isLBListenerPolicyTargetID).(string)

		target = &vpcv1.LoadBalancerListenerPolicyPatchTargetLoadBalancerPoolIdentity{
			ID: &id,
		}

		updatePolicyOptions.Target = target
		hasChanged = true
	} else if d.Get(isLBListenerPolicyAction).(string) == "redirect" {
		//if Action is redirect and either status code or URL chnaged, set accordingly
		//LoadBalancerListenerPolicyPatchTargetLoadBalancerListenerPolicyRedirectURLPatch

		redirectPatch := vpcv1.LoadBalancerListenerPolicyPatchTargetLoadBalancerListenerPolicyRedirectURLPatch{}

		targetChange := false
		if d.HasChange(isLBListenerPolicyTargetHTTPStatusCode) {
			status := d.Get(isLBListenerPolicyTargetHTTPStatusCode).(int)
			sc := int64(status)
			redirectPatch.HttpStatusCode = &sc
			hasChanged = true
			targetChange = true
		}

		if d.HasChange(isLBListenerPolicyTargetURL) {
			url := d.Get(isLBListenerPolicyTargetURL).(string)
			redirectPatch.URL = &url
			hasChanged = true
			targetChange = true
		}

		//Update the target only if there is a change in either statusCode or URL
		if targetChange {
			target = &redirectPatch
			updatePolicyOptions.Target = target
		}
	}

	if hasChanged {
		isLBListenerPolicyKey := "load_balancer_listener_policy_key_" + lbID + listenerID
		ibmMutexKV.Lock(isLBListenerPolicyKey)
		defer ibmMutexKV.Unlock(isLBListenerPolicyKey)

		_, err = isWaitForLbAvailable(sess, lbID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return fmt.Errorf(
				"LB-LP Error checking for load balancer (%s) is active: %s", lbID, err)
		}
		_, response, err := sess.UpdateLoadBalancerListenerPolicy(&updatePolicyOptions)
		if err != nil {
			return fmt.Errorf("Error Updating in policy : %s\n%s", err, response)
		}

		_, err = isWaitForLbListenerPolicyAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceIBMISLBListenerPolicyDelete(d *schema.ResourceData, meta interface{}) error {

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	//Retrieve lbId, listenerId and policyID
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	listenerID := parts[1]
	policyID := parts[2]

	if userDetails.generation == 1 {
		err := classicLbListenerPolicycDelete(d, meta, lbID, listenerID, policyID)
		if err != nil {
			return err
		}
	} else {
		err := lbListenerPolicyDelete(d, meta, lbID, listenerID, policyID)
		if err != nil {
			return err
		}
	}

	d.SetId("")
	return nil

}

func classicLbListenerPolicycDelete(d *schema.ResourceData, meta interface{}, lbID, listenerID, ID string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	deleteLbListenerPolicyOptions := &vpcclassicv1.DeleteLoadBalancerListenerPolicyOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		ID:             &ID,
	}

	response, err := sess.DeleteLoadBalancerListenerPolicy(deleteLbListenerPolicyOptions)

	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error in classicLbListenerPolicycDelete: %s\n%s", err, response)
	}
	if response.StatusCode != 404 {
		return nil
	}
	_, err = isWaitForLbListenerPolicyClassicDeleted(sess, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	return nil
}

func lbListenerPolicyDelete(d *schema.ResourceData, meta interface{}, lbID, listenerID, ID string) error {

	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	deleteLbListenerPolicyOptions := &vpcv1.DeleteLoadBalancerListenerPolicyOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		ID:             &ID,
	}
	response, err := sess.DeleteLoadBalancerListenerPolicy(deleteLbListenerPolicyOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error in lbListenerPolicyDelete: %s\n%s", err, response)
	}
	if response.StatusCode != 404 {
		return nil
	}
	_, err = isWaitForLbListnerPolicyDeleted(sess, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	return nil
}
func isWaitForLbListnerPolicyDeleted(vpc *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBListenerPolicyRetry, isLBListenerPolicyDeleting},
		Target:     []string{},
		Refresh:    isLbListenerPolicyDeleteRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLbListenerPolicyDeleteRefreshFunc(vpc *vpcv1.VpcV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		//Retrieve lbId, listenerId and policyID
		parts, err := idParts(id)
		if err != nil {
			return nil, isLBListenerPolicyFailed, nil
		}

		lbID := parts[0]
		listenerID := parts[1]
		//policyID := parts[2]
		policyID := id

		getLbListenerPolicyOptions := &vpcv1.GetLoadBalancerListenerPolicyOptions{
			LoadBalancerID: &lbID,
			ListenerID:     &listenerID,
			ID:             &policyID,
		}

		//Getting lb listener policy
		policy, response, err := vpc.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)
		if err != nil {
			return nil, isLBListenerPolicyFailed, nil
		}

		if err == nil {
			if *(*policy).ProvisioningStatus == isLBListenerPolicyFailed {
				return policy, isLBListenerPolicyFailed, fmt.Errorf("The LB-LP %s failed to delete: %v", *policy.ID, err)
			}
			return policy, isLBListenerPolicyDeleting, nil
		}

		if (*response).StatusCode == 404 {
			return nil, isLBListenerPolicyDeleted, nil
		}

		return nil, isLBListenerPolicyDeleting, err
	}
}

func classicVpcClient(meta interface{}) (*vpcclassicv1.VpcClassicV1, error) {
	sess, err := meta.(ClientSession).VpcClassicV1API()
	return sess, err
}

func classicLbListenerPolicyGet(d *schema.ResourceData, meta interface{}, lbID, listenerID, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	//Getting policy optins
	getLbListenerPolicyOptions := &vpcclassicv1.GetLoadBalancerListenerPolicyOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		ID:             &id,
	}

	//Getting lb listener policy
	policy, response, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)
	if err != nil {
		return fmt.Errorf("Error in classicLbListenerPolicyGet : %s\n%s", err, response)
	}

	d.Set(isLBListenerPolicyLBID, lbID)
	d.Set(isLBListenerPolicyListenerID, listenerID)
	d.Set(isLBListenerPolicyAction, policy.Action)
	d.Set(isLBListenerPolicyID, id)
	d.Set(isLBListenerPolicyPriority, policy.Priority)
	d.Set(isLBListenerPolicyName, policy.Name)
	d.Set(isLBListenerPolicyStatus, policy.ProvisioningStatus)

	if policy.Rules != nil {
		rulesSet := make([]interface{}, 0)
		for _, rule := range policy.Rules {
			getLbListenerPolicyRulesOptions := &vpcclassicv1.GetLoadBalancerListenerPolicyRuleOptions{
				LoadBalancerID: &lbID,
				ListenerID:     &listenerID,
				ID:             rule.ID,
				PolicyID:       &id,
			}
			ruleInfo, response, err := sess.GetLoadBalancerListenerPolicyRule(getLbListenerPolicyRulesOptions)
			if err != nil {
				return fmt.Errorf("Error in classicLbListenerPolicyGet rule: %s\n%s", err, response)
			}

			r := map[string]interface{}{
				isLBListenerPolicyRuleID:        *ruleInfo.ID,
				isLBListenerPolicyRuleCondition: *ruleInfo.Condition,
				isLBListenerPolicyRuleType:      *ruleInfo.Type,
				isLBListenerPolicyRuleField:     *ruleInfo.Field,
				isLBListenerPolicyRuleValue:     *ruleInfo.Value,
			}
			rulesSet = append(rulesSet, r)
		}
		d.Set(isLBListenerPolicyRulesInfo, rulesSet)
	}

	// `LoadBalancerPoolReference` is in the response if `action` is `forward`.
	// `LoadBalancerListenerPolicyRedirectURL` is in the response if `action` is `redirect`.

	if *(policy.Action) == "forward" {
		target := policy.Target.(*vpcclassicv1.LoadBalancerListenerPolicyTargetReference)
		d.Set(isLBListenerPolicyTargetID, target.ID)

	} else if *(policy.Action) == "redirect" {
		target := policy.Target.(*vpcclassicv1.LoadBalancerListenerPolicyTargetReference)
		d.Set(isLBListenerPolicyTargetURL, target.URL)
		d.Set(isLBListenerPolicyTargetHTTPStatusCode, target.HttpStatusCode)

	}

	return nil
}

func lbListenerPolicyGet(d *schema.ResourceData, meta interface{}, lbID, listenerID, id string) error {

	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	//Getting policy optins
	getLbListenerPolicyOptions := &vpcv1.GetLoadBalancerListenerPolicyOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		ID:             &id,
	}

	//Getting lb listener policy
	policy, _, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)
	if err != nil {
		return err
	}

	//set the argument values
	d.Set(isLBListenerPolicyLBID, lbID)
	d.Set(isLBListenerPolicyListenerID, listenerID)
	d.Set(isLBListenerPolicyAction, policy.Action)
	d.Set(isLBListenerPolicyID, id)
	d.Set(isLBListenerPolicyPriority, policy.Priority)
	d.Set(isLBListenerPolicyName, policy.Name)
	d.Set(isLBListenerPolicyStatus, policy.ProvisioningStatus)

	//set rules - Doubt (Rules has condition, type, value, field and id where as SDK has only Href and id, so setting only id)
	if policy.Rules != nil {
		policyRules := policy.Rules
		rulesInfo := make([]map[string]interface{}, 0)
		for _, index := range policyRules {

			l := map[string]interface{}{
				"id": index.ID,
			}
			rulesInfo = append(rulesInfo, l)
		}
		d.Set(isLBListenerPolicyRules, rulesInfo)
	}

	// `LoadBalancerPoolReference` is in the response if `action` is `forward`.
	// `LoadBalancerListenerPolicyRedirectURL` is in the response if `action` is `redirect`.

	if *(policy.Action) == "forward" {
		target := policy.Target.(*vpcv1.LoadBalancerListenerPolicyTargetReference)
		d.Set(isLBListenerPolicyTargetID, target.ID)

	} else if *(policy.Action) == "redirect" {
		target := policy.Target.(*vpcv1.LoadBalancerListenerPolicyTargetReference)
		d.Set(isLBListenerPolicyTargetURL, target.URL)
		d.Set(isLBListenerPolicyTargetHTTPStatusCode, target.HttpStatusCode)

	}

	return nil
}

func isWaitForLbListenerPolicyClassicDeleted(vpc *vpcclassicv1.VpcClassicV1, id string, timeout time.Duration) (interface{}, error) {

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBListenerPolicyRetry, isLBListenerPolicyDeleting, "delete_pending"},
		Target:     []string{},
		Refresh:    isLbListenerPolicyClassicDeleteRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLbListenerPolicyClassicDeleteRefreshFunc(vpc *vpcclassicv1.VpcClassicV1, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		//Retrieve lbId and listenerId
		parts, err := idParts(id)
		if err != nil {
			return nil, isLBListenerPolicyFailed, nil
		}

		lbID := parts[0]
		listenerID := parts[1]
		policyID := parts[2]

		getLbListenerPolicyOptions := &vpcclassicv1.GetLoadBalancerListenerPolicyOptions{
			LoadBalancerID: &lbID,
			ListenerID:     &listenerID,
			ID:             &policyID,
		}

		//Getting lb listener policy
		policy, response, err := vpc.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

		failed := isLBListenerPolicyFailed
		if err != nil {
			if policy.ProvisioningStatus == &failed {
				return policy, isLBListenerPolicyFailed, fmt.Errorf("The LB-LP %s failed to delete: %v", *policy.ID, err)
			}
			return nil, isLBListenerPolicyFailed, nil
		}

		if (*response).StatusCode == 404 {
			return nil, isLBListenerPolicyDeleted, nil
		}

		return nil, isLBListenerPolicyDeleting, err
	}
}
