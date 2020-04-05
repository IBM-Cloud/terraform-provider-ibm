package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	//"github.ibm.com/Bluemix/riaas-go-client/clients/lbaas"
	//iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	//"github.ibm.com/Bluemix/riaas-go-client/riaas/client/l_baas"
	//"github.ibm.com/Bluemix/riaas-go-client/riaas/models"

	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isLBListenerPolicyLBID                 = "load_balancer_id"
	isLBListenerPolicyListenerID           = "listener_id"
	isLBListenerPolicyAction               = "action"
	isLBListenerPolicyPriority             = "priority"
	isLBListenerPolicyName                 = "name"
	isLBListenerPolicyID                   = "id"
	isLBListenerPolicyRules                = "rules"
	isLBListenerPolicyTargetID             = "target-id"
	isLBListenerPolicyTargetHTTPStatusCode = "target-http-status-code"
	isLBListenerPolicyTargetURL            = "target-url"
	isLBListenerPolicyStatus               = "provisioning_status"
	isLBListenerPolicyAvailable            = "available"
	isLBListenerPolicyFailed               = "failed"
	isLBListenerPolicyPending              = "pending"
	isLBListenerPolicyDeleting             = "deleting"
	isLBListenerPolicyDeleted              = "done"
	isLBListenerPolicyRetry                = "retry"
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
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

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
				ValidateFunc: validateISName, //Exisiting function
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
						"condition": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateAllowedStringValue([]string{"contains", "equals", "matches_regex"}),
							Description:  "Condition of the rule",
						},

						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateAllowedStringValue([]string{"header", "hostname", "path"}),
							Description:  "Type of the rule",
						},

						"value": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateStringLength,
							Description:  "Value to be matched for rule condition",
						},

						"field": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validateStringLength,
							Description:  "HTTP header field. This is only applicable to rule type.",
						},

						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTTP header field. This is only applicable to header rule type.",
						},
					},
				},
			},

			isLBListenerPolicyTargetID: {
				Type:     schema.TypeInt,
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

	log.Printf("LB-LP: [DEBUG] LB Listener policy create")

	//Get the Load balancer ID
	var lbID string
	if lb, ok := d.GetOk(isLBListenerPolicyLBID); ok {
		lbID = lb.(string)
	}

	//Get the listner ID
	var listenerID string
	if lID, ok := d.GetOk(isLBListenerPolicyListenerID); ok {
		listenerID = lID.(string)
	}

	//Get policy action
	var action string
	if pAction, ok := d.GetOk(isLBListenerPolicyAction); ok {
		action = pAction.(string)
	}

	//Get priority of policy
	var proirity int64
	if prio, ok := d.GetOk(isLBListenerPolicyPriority); ok {
		proirity = prio.(int64)
	}

	if userDetails.generation == 1 {
		err := classicLbListenerPolicyCreate(d, meta, lbID, listenerID, action, proirity)
		if err != nil {
			return err
		}
	} else {
		err := lbListenerPolicyCreate(d, meta, lbID, listenerID, action, proirity)
		if err != nil {
			return err
		}
	}

	return resourceIBMISLBListenerPolicyRead(d, meta)
}

func resourceIBMISLBListenerPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("LB-LP: [DEBUG] LB Listener policy READ")
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
	log.Printf("LB-LP: [DEBUG] LB-LP ID: %v,  LBID:%v, ListenerID:%v policyID: %v", ID, lbID, listenerID, policyID)

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
	_, _, err1 := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

	if err1 != nil {
		iserror, ok := err1.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return nil
			}
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
	_, _, err1 := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)
	if err1 != nil {
		iserror, ok := err1.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return nil
			}
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

	//Update arguments Name, Priority and Target

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	listenerID := parts[1]
	policyID := parts[2]

	name := d.Get(isLBListenerPolicyName).(string)
	priority := d.Get(isLBListenerPolicyPriority).(int64)

	// `LoadBalancerPoolReference` is in the response if `action` is `forward`.
	// `LoadBalancerListenerPolicyRedirectURL` is in the response if `action` is `redirect`.

	//Doubt - is it right way to capture target

	if userDetails.generation == 1 {
		var target vpcclassicv1.LoadBalancerListenerPolicyPatchTargetIntf
		if d.Get(isLBListenerPolicyAction).(string) == "forward" {
			//target = d.Get(isLBListenerPolicyTargetID).(int)
		} else if d.Get(isLBListenerPolicyAction).(string) == "redirect" {
			//target = d.Get(isLBListenerPolicyTargetURL).(string)
		}
		err := classicLbListenerPolicyUpdate(d, meta, lbID, listenerID, priority, target, policyID, name)
		if err != nil {
			return err
		}
	} else {
		var target vpcv1.LoadBalancerListenerPolicyPatchTargetIntf
		if d.Get(isLBListenerPolicyAction).(string) == "forward" {
			//target = d.Get(isLBListenerPolicyTargetID).(int)
		} else if d.Get(isLBListenerPolicyAction).(string) == "redirect" {
			//target = d.Get(isLBListenerPolicyTargetURL).(string)
		}
		err := lbListenerPolicyUpdate(d, meta, lbID, listenerID, priority, target, policyID, name)
		if err != nil {
			return err
		}
	}

	return resourceIBMISLBListenerRead(d, meta)
}

func classicLbListenerPolicyUpdate(d *schema.ResourceData, meta interface{}, lbID, listenerID string, priority int64, target vpcclassicv1.LoadBalancerListenerPolicyPatchTargetIntf, ID, name string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	// `LoadBalancerPoolReference` is in the response if `action` is `forward`.
	// `LoadBalancerListenerPolicyRedirectURL` is in the response if `action` is `redirect`.

	var targetChange bool
	if d.Get(isLBListenerPolicyAction).(string) == "forward" {
		targetChange = d.HasChange(isLBListenerPolicyTargetID)
	} else if d.Get(isLBListenerPolicyAction).(string) == "redirect" {
		targetChange = d.HasChange(isLBListenerPolicyTargetURL)
	}

	if d.HasChange(isLBListenerPolicyName) || d.HasChange(isLBListenerPolicyPriority) || targetChange {
		updatePolicyOptions := &vpcclassicv1.UpdateLoadBalancerListenerPolicyOptions{
			LoadBalancerID: &lbID,
			ListenerID:     &listenerID,
			ID:             &ID,
			Name:           &name,
			Priority:       &priority,
			Target:         target,
		}

		_, _, err := sess.UpdateLoadBalancerListenerPolicy(updatePolicyOptions)
		if err != nil {
			return err
		}
	}

	//Getting policy optins
	getLbListenerPolicyOptions := &vpcclassicv1.GetLoadBalancerListenerPolicyOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		ID:             &ID,
	}

	//Getting lb listener policy
	policy, _, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)
	if err != nil {
		return err
	}

	log.Printf("LB-LP: [DEBUG] Updated policy %v", policy)

	return nil
}

func lbListenerPolicyUpdate(d *schema.ResourceData, meta interface{}, lbID, listenerID string, priority int64, target vpcv1.LoadBalancerListenerPolicyPatchTargetIntf, ID, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	// `LoadBalancerPoolReference` is in the response if `action` is `forward`.
	// `LoadBalancerListenerPolicyRedirectURL` is in the response if `action` is `redirect`.
	var targetChange bool
	if d.Get(isLBListenerPolicyAction).(string) == "forward" {
		targetChange = d.HasChange(isLBListenerPolicyTargetID)
	} else if d.Get(isLBListenerPolicyAction).(string) == "redirect" {
		targetChange = d.HasChange(isLBListenerPolicyTargetURL)
	}
	prio := int64(priority)
	//Doubt - change check done only for name, priority and target.
	if d.HasChange(isLBListenerPolicyName) || d.HasChange(isLBListenerPolicyPriority) || targetChange {
		updatePolicyOptions := &vpcv1.UpdateLoadBalancerListenerPolicyOptions{
			LoadBalancerID: &lbID,
			ListenerID:     &listenerID,
			ID:             &ID,
			Name:           &name,
			Priority:       &prio,
			Target:         target,
		}

		_, _, err := sess.UpdateLoadBalancerListenerPolicy(updatePolicyOptions)
		if err != nil {
			return err
		}
	}

	//Getting policy optins
	getLbListenerPolicyOptions := &vpcv1.GetLoadBalancerListenerPolicyOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		ID:             &ID,
	}

	//Getting lb listener policy
	policy, _, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)
	if err != nil {
		return err
	}
	log.Printf("LB-LP: [DEBUG] Updated policy %v", policy)

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
	_, err = sess.DeleteLoadBalancerListenerPolicy(deleteLbListenerPolicyOptions)
	if err != nil {
		return err
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
	_, err = sess.DeleteLoadBalancerListenerPolicy(deleteLbListenerPolicyOptions)
	if err != nil {
		return err
	}
	_, err = isWaitForLbListnerPolicyDeleted(sess, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	return nil
}
func isWaitForLbListnerPolicyDeleted(vpc *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("LB-LP: Waiting for LB-LP (%s) to be deleted.", id)

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
		log.Printf("LB-LP : [DEBUG] delete function ")

		//Retrieve lbId, listenerId and policyID
		parts, err := idParts(id)
		if err != nil {
			return nil, isLBListenerPolicyFailed, nil
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
		policy, _, err := vpc.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)
		if err != nil {
			return nil, isLBListenerPolicyFailed, nil
		}

		failed := isLBListenerPolicyFailed

		if err == nil {
			if policy.ProvisioningStatus == &failed {
				return policy, isLBListenerPolicyFailed, fmt.Errorf("The LB-LP %s failed to delete: %v", *policy.ID, err)
			}
			return policy, isLBListenerPolicyDeleting, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("LB-LP [DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				log.Printf("LB-LP [DEBUG] returning deleted")
				return nil, isLBListenerPolicyDeleted, nil
			}
		}

		return nil, isLBListenerPolicyDeleting, err
	}
}

func isWaitForLbListenerPolicyClassicDeleted(vpc *vpcclassicv1.VpcClassicV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("LB-LP : Waiting for LB_LP (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBListenerPolicyRetry, isLBListenerPolicyDeleting},
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
		log.Printf("LB-LP: [DEBUG] delete function ")

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
		policy, _, err := vpc.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

		failed := isLBListenerPolicyFailed
		if err != nil {
			if policy.ProvisioningStatus == &failed {
				return policy, isLBListenerPolicyFailed, fmt.Errorf("The LB-LP %s failed to delete: %v", *policy.ID, err)
			}
			return nil, isLBListenerPolicyFailed, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("LB-LP : [DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				log.Printf("LB-LP: [DEBUG] returning deleted")
				return nil, isLBListenerPolicyDeleted, nil
			}
		}
		log.Printf("LB-LP : [DEBUG] returning x")
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
	policy, _, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)
	if err != nil {
		return err
	}

	//Doubt - any params are missing to set and which params need to check for avaialbility nil check
	//setting based on LoadBalancerListenerPolicy params
	d.Set(isLBListenerPolicyAction, policy.Action)
	d.Set(isLBListenerPolicyID, policy.ID)
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
		d.Set(isLBListenerPolicyTargetID, policy.Target)
	} else if *(policy.Action) == "redirect" {
		d.Set(isLBListenerPolicyTargetURL, policy.Target)
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
	d.Set(isLBListenerPolicyAction, policy.Action)
	d.Set(isLBListenerPolicyID, policy.ID)
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
		d.Set(isLBListenerPolicyTargetID, policy.Target)
	} else if *(policy.Action) == "redirect" {
		d.Set(isLBListenerPolicyTargetURL, policy.Target)
	}

	return nil
}

func classicLbListenerPolicyCreate(d *schema.ResourceData, meta interface{}, lbID, listenerID, action string, priority int64) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	options := &vpcclassicv1.CreateLoadBalancerListenerPolicyOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		Action:         &action,
		Priority:       &priority,
	}
	//options := &vpcclassicv1.NewCreateLoadBalancerListenerPolicyOptions(lbID, listenerID, action, priority)

	//Doubt - do we need to wait for avaialability of LB and listener both

	/*_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer (%s) is active: %s", lbID, err)
	}*/

	policy, _, err := sess.CreateLoadBalancerListenerPolicy(options)
	if err != nil {
		return fmt.Errorf("Error while creating lb listener policy for LB %s: %v", lbID, err)
	}

	d.SetId(*policy.ID)
	log.Printf("LB-LP: [INFO] classicLbListenerPolicy ID : %s", *policy.ID)

	_, err = isWaitForClassicLbListenerPolicyAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	log.Printf("LB-LP: [INFO] Load balancer Listener : %s", d.Id())
	return resourceIBMISLBListenerPolicyRead(d, meta)
}

func isWaitForClassicLbListenerPolicyAvailable(vpc *vpcclassicv1.VpcClassicV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("LB-LP: Waiting for LB-LP (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBListenerPolicyPending},
		Target:     []string{isLBListenerPolicyAvailable, isLBListenerPolicyFailed},
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
			return policy, *policy.ProvisioningStatus, nil
		}

		return policy, *policy.ProvisioningStatus, nil
	}
}

func vpcClient(meta interface{}) (*vpcv1.VpcV1, error) {
	sess, err := meta.(ClientSession).VpcV1API()
	return sess, err
}

func lbListenerPolicyCreate(d *schema.ResourceData, meta interface{}, lbID, listenerID, action string, priority int64) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	options := &vpcv1.CreateLoadBalancerListenerPolicyOptions{
		LoadBalancerID: &lbID,
		ListenerID:     &listenerID,
		Action:         &action,
		Priority:       &priority,
	}
	//options := &vpcv1.NewCreateLoadBalancerListenerPolicyOptions(lbID, listenerID, action, priority)

	policy, _, err := sess.CreateLoadBalancerListenerPolicy(options)
	if err != nil {
		return fmt.Errorf("Error while creating lb listener policy for LB %s: %v", lbID, err)
	}

	d.SetId(*policy.ID)
	log.Printf("LB-LP: [INFO] classicLbListenerPolicy ID : %s", *policy.ID)

	_, err = isWaitForLbListenerPolicyAvailable(sess, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	log.Printf(" LB-LP: [INFO] Load balancer Listener Policy : %s", d.Id())
	return nil
}

func isWaitForLbListenerPolicyAvailable(vpc *vpcv1.VpcV1, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("LB-LP: Waiting for Load balancer Listener Policy (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBListenerPolicyPending},
		Target:     []string{isLBListenerPolicyAvailable, isLBListenerPolicyFailed},
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
			return policy, *policy.ProvisioningStatus, nil
		}

		return policy, *policy.ProvisioningStatus, nil
	}
}
