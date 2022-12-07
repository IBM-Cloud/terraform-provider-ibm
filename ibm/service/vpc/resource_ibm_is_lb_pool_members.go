// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isLBPoolMembersPort               = "port"
	isLBPoolMembersTargetAddress      = "target_address"
	isLBPoolMembersTargetID           = "target_id"
	isLBPoolMembersWeight             = "weight"
	isLBPoolMembersProvisioningStatus = "provisioning_status"
	isLBPoolMembersHealth             = "health"
	isLBPoolMembersHref               = "href"
	isLBPoolMembersDeletePending      = "delete_pending"
	isLBPoolMembersDeleted            = "done"
	isLBPoolMembersActive             = "active"
	isLBPoolMembers                   = "members"
	isLBPoolMembersTarget             = "target"
)

func ResourceIBMISLBPoolMemberss() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISLBPoolMembersCreate,
		Read:     resourceIBMISLBPoolMembersRead,
		Update:   resourceIBMISLBPoolMembersUpdate,
		Delete:   resourceIBMISLBPoolMembersDelete,
		Exists:   resourceIBMISLBPoolMembersExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			isLBPoolID: {
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
						new := strings.Split(n, "/")
						if strings.Compare(new[1], o) == 0 {
							return true
						}
					}

					return false
				},
				Description: "Loadblancer Poold ID",
			},

			isLBID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Load balancer ID",
			},
			isLBPoolMembers: {
				Type:     schema.TypeList,
				Required: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isLBPoolMembersPort: {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Load Balancer Pool port",
						},
						isLBPoolMembersTarget: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Load balancer pool member target, Either target id or ip address",
						},
						isLBPoolMembersWeight: {
							Type:         schema.TypeInt,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_lb_pool_member", isLBPoolMembersWeight),
							Description:  "Load balcner pool member weight",
						},
					},
				},
			},

			isLBPoolMembersProvisioningStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Load balancer Pool member provisioning status",
			},

			isLBPoolMembersHealth: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "LB Pool member health",
			},

			isLBPoolMembersHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "LB pool member Href value",
			},

			flex.RelatedCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the LB resource",
			},
		},
	}
}

func ResourceIBMISLBPoolMembersValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isLBPoolMembersWeight,
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			Optional:                   true,
			MinValue:                   "0",
			MaxValue:                   "100"})

	ibmISLBResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_lb_pool_member", Schema: validateSchema}
	return &ibmISLBResourceValidator
}

func resourceIBMISLBPoolMembersCreate(d *schema.ResourceData, meta interface{}) error {

	log.Printf("[DEBUG] LB Pool create")
	lbPoolID, err := getPoolId(d.Get(isLBPoolID).(string))
	if err != nil {
		return err
	}

	lbID := d.Get(isLBID).(string)

	if membersIntf, ok := d.GetOk(isLBPoolMembers); ok {
		nics := nicsintf.([]interface{})
		var intfs []vpcv1.NetworkInterfacePrototype
		for _, resource := range nics {
			nic := resource.(map[string]interface{})
			nwInterface := &vpcv1.NetworkInterfacePrototype{}
			subnetintf, _ := nic[isInstanceNicSubnet]

		}
	}

	port := d.Get(isLBPoolMembersPort).(int)
	port64 := int64(port)

	var weight int64

	isLBKey := "load_balancer_key_" + lbID
	conns.IbmMutexKV.Lock(isLBKey)
	defer conns.IbmMutexKV.Unlock(isLBKey)

	err = lbpMembersCreate(d, meta, lbID, lbPoolID, port64, weight)
	if err != nil {
		return err
	}

	return resourceIBMISLBPoolMembersRead(d, meta)
}

func lbpMembersCreate(d *schema.ResourceData, meta interface{}, lbID, lbPoolID string, port, weight int64) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("[ERROR] Error checking for load balancer pool (%s) is active: %s", lbPoolID, err)
	}

	_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("[ERROR] Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	options := &vpcv1.CreateLoadBalancerPoolMembersOptions{
		LoadBalancerID: &lbID,
		PoolID:         &lbPoolID,
		Port:           &port,
	}

	if _, ok := d.GetOk(isLBPoolMembersTargetAddress); ok {
		targetAddress := d.Get(isLBPoolMembersTargetAddress).(string)
		target := &vpcv1.LoadBalancerPoolMembersTargetPrototype{
			Address: &targetAddress,
		}
		options.Target = target
	} else {
		targetID := d.Get(isLBPoolMembersTargetID).(string)
		target := &vpcv1.LoadBalancerPoolMembersTargetPrototype{
			ID: &targetID,
		}
		options.Target = target
	}
	if w, ok := d.GetOkExists(isLBPoolMembersWeight); ok {
		weight = int64(w.(int))
		options.Weight = &weight
	}

	lbPoolMembers, response, err := sess.CreateLoadBalancerPoolMembers(options)
	if err != nil {
		return fmt.Errorf("[DEBUG] lbpool member create err: %s\n%s", err, response)
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", lbID, lbPoolID, *lbPoolMembers.ID))
	log.Printf("[INFO] lbpool member : %s", *lbPoolMembers.ID)

	_, err = isWaitForLBPoolMembersAvailable(sess, lbID, lbPoolID, *lbPoolMembers.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("[ERROR] Error checking for load balancer pool (%s) is active: %s", lbPoolID, err)
	}

	_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("[ERROR] Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	return nil
}

func isWaitForLBPoolMembersAvailable(lbc *vpcv1.VpcV1, lbID, lbPoolID, lbPoolMemID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for load balancer pool member(%s) to be available.", lbPoolMemID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"create_pending", "update_pending", "maintenance_pending"},
		Target:     []string{isLBPoolMembersActive, ""},
		Refresh:    isLBPoolMembersRefreshFunc(lbc, lbID, lbPoolID, lbPoolMemID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLBPoolMembersRefreshFunc(lbc *vpcv1.VpcV1, lbID, lbPoolID, lbPoolMemID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getlbpmoptions := &vpcv1.GetLoadBalancerPoolMembersOptions{
			LoadBalancerID: &lbID,
			PoolID:         &lbPoolID,
			ID:             &lbPoolMemID,
		}
		lbPoolMem, response, err := lbc.GetLoadBalancerPoolMembers(getlbpmoptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error Getting Load Balancer Pool Members: %s\n%s", err, response)
		}

		if *lbPoolMem.ProvisioningStatus == isLBPoolMembersActive {
			return lbPoolMem, *lbPoolMem.ProvisioningStatus, nil
		}

		return lbPoolMem, *lbPoolMem.ProvisioningStatus, nil
	}
}

func resourceIBMISLBPoolMembersRead(d *schema.ResourceData, meta interface{}) error {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return err
	}

	if len(parts) < 3 {
		return fmt.Errorf(
			"The id should contain loadbalancer Id, loadbalancer pool Id and loadbalancer poolmemebr Id")
	}

	lbID := parts[0]
	lbPoolID := parts[1]
	lbPoolMemID := parts[2]

	err = lbpmemberGet(d, meta, lbID, lbPoolID, lbPoolMemID)
	if err != nil {
		return err
	}

	return nil
}

func lbpmemberGet(d *schema.ResourceData, meta interface{}, lbID, lbPoolID, lbPoolMemID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getlbpmoptions := &vpcv1.GetLoadBalancerPoolMembersOptions{
		LoadBalancerID: &lbID,
		PoolID:         &lbPoolID,
		ID:             &lbPoolMemID,
	}
	lbPoolMem, response, err := sess.GetLoadBalancerPoolMembers(getlbpmoptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Error Getting Load Balancer Pool Members: %s\n%s", err, response)
	}
	d.Set(isLBPoolID, lbPoolID)
	d.Set(isLBID, lbID)
	d.Set(isLBPoolMembersPort, *lbPoolMem.Port)

	target := lbPoolMem.Target.(*vpcv1.LoadBalancerPoolMembersTarget)
	if target.Address != nil {
		d.Set(isLBPoolMembersTargetAddress, *target.Address)
	}
	if target.ID != nil {
		d.Set(isLBPoolMembersTargetID, *target.ID)
	}
	d.Set(isLBPoolMembersWeight, *lbPoolMem.Weight)
	d.Set(isLBPoolMembersProvisioningStatus, *lbPoolMem.ProvisioningStatus)
	d.Set(isLBPoolMembersHealth, *lbPoolMem.Health)
	d.Set(isLBPoolMembersHref, *lbPoolMem.Href)
	getLoadBalancerOptions := &vpcv1.GetLoadBalancerOptions{
		ID: &lbID,
	}
	lb, response, err := sess.GetLoadBalancer(getLoadBalancerOptions)
	if err != nil {
		return fmt.Errorf("[ERROR] Error Getting Load Balancer : %s\n%s", err, response)
	}
	d.Set(flex.RelatedCRN, *lb.CRN)
	return nil
}

func resourceIBMISLBPoolMembersUpdate(d *schema.ResourceData, meta interface{}) error {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	lbPoolID := parts[1]
	lbPoolMemID := parts[2]

	err = lbpmemberUpdate(d, meta, lbID, lbPoolID, lbPoolMemID)
	if err != nil {
		return err
	}

	return resourceIBMISLBPoolMembersRead(d, meta)
}

func lbpmemberUpdate(d *schema.ResourceData, meta interface{}, lbID, lbPoolID, lbPoolMemID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	if d.HasChange(isLBPoolMembersTargetID) || d.HasChange(isLBPoolMembersTargetAddress) || d.HasChange(isLBPoolMembersPort) || d.HasChange(isLBPoolMembersWeight) {

		port := int64(d.Get(isLBPoolMembersPort).(int))
		weight := int64(d.Get(isLBPoolMembersWeight).(int))

		isLBKey := "load_balancer_key_" + lbID
		conns.IbmMutexKV.Lock(isLBKey)
		defer conns.IbmMutexKV.Unlock(isLBKey)

		_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error checking for load balancer pool (%s) is active: %s", lbPoolID, err)
		}

		_, err = isWaitForLBPoolMembersAvailable(sess, lbID, lbPoolID, lbPoolMemID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return err
		}

		_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error checking for load balancer (%s) is active: %s", lbID, err)
		}

		updatelbpmoptions := &vpcv1.UpdateLoadBalancerPoolMembersOptions{
			LoadBalancerID: &lbID,
			PoolID:         &lbPoolID,
			ID:             &lbPoolMemID,
		}

		loadBalancerPoolMembersPatchModel := &vpcv1.LoadBalancerPoolMembersPatch{
			Port:   &port,
			Weight: &weight,
		}

		if _, ok := d.GetOk(isLBPoolMembersTargetAddress); ok {
			targetAddress := d.Get(isLBPoolMembersTargetAddress).(string)
			target := &vpcv1.LoadBalancerPoolMembersTargetPrototype{
				Address: &targetAddress,
			}
			loadBalancerPoolMembersPatchModel.Target = target
		} else {
			targetID := d.Get(isLBPoolMembersTargetID).(string)
			target := &vpcv1.LoadBalancerPoolMembersTargetPrototype{
				ID: &targetID,
			}
			loadBalancerPoolMembersPatchModel.Target = target
		}

		loadBalancerPoolMembersPatch, err := loadBalancerPoolMembersPatchModel.AsPatch()
		if err != nil {
			return fmt.Errorf("[ERROR] Error calling asPatch for LoadBalancerPoolMembersPatch: %s", err)
		}
		updatelbpmoptions.LoadBalancerPoolMembersPatch = loadBalancerPoolMembersPatch

		_, response, err := sess.UpdateLoadBalancerPoolMembers(updatelbpmoptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Updating Load Balancer Pool Members: %s\n%s", err, response)
		}
		_, err = isWaitForLBPoolMembersAvailable(sess, lbID, lbPoolID, lbPoolMemID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return err
		}

		_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error checking for load balancer pool (%s) is active: %s", lbPoolID, err)
		}

		_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error checking for load balancer (%s) is active: %s", lbID, err)
		}
	}
	return nil
}

func resourceIBMISLBPoolMembersDelete(d *schema.ResourceData, meta interface{}) error {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	lbPoolID := parts[1]
	lbPoolMemID := parts[2]

	isLBKey := "load_balancer_key_" + lbID
	conns.IbmMutexKV.Lock(isLBKey)
	defer conns.IbmMutexKV.Unlock(isLBKey)

	err = lbpmemberDelete(d, meta, lbID, lbPoolID, lbPoolMemID)
	if err != nil {
		return err
	}

	return nil
}

func lbpmemberDelete(d *schema.ResourceData, meta interface{}, lbID, lbPoolID, lbPoolMemID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	getlbpmoptions := &vpcv1.GetLoadBalancerPoolMembersOptions{
		LoadBalancerID: &lbID,
		PoolID:         &lbPoolID,
		ID:             &lbPoolMemID,
	}
	_, response, err := sess.GetLoadBalancerPoolMembers(getlbpmoptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Error Getting Load Balancer Pool Members: %s\n%s", err, response)
	}
	_, err = isWaitForLBPoolMembersAvailable(sess, lbID, lbPoolID, lbPoolMemID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf("[ERROR] Error checking for load balancer pool (%s) is active: %s", lbPoolID, err)
	}

	_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf("[ERROR] Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	dellbpmoptions := &vpcv1.DeleteLoadBalancerPoolMembersOptions{
		LoadBalancerID: &lbID,
		PoolID:         &lbPoolID,
		ID:             &lbPoolMemID,
	}
	response, err = sess.DeleteLoadBalancerPoolMembers(dellbpmoptions)
	if err != nil {
		return fmt.Errorf("[ERROR] Error Deleting Load Balancer Pool Members: %s\n%s", err, response)
	}

	_, err = isWaitForLBPoolMembersDeleted(sess, lbID, lbPoolID, lbPoolMemID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	_, err = isWaitForLBPoolActive(sess, lbID, lbPoolID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf("[ERROR] Error checking for load balancer pool (%s) is active: %s", lbPoolID, err)
	}

	_, err = isWaitForLBAvailable(sess, lbID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf("[ERROR] Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	d.SetId("")
	return nil
}

func isWaitForLBPoolMembersDeleted(lbc *vpcv1.VpcV1, lbID, lbPoolID, lbPoolMemID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", lbPoolMemID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBPoolMembersDeletePending},
		Target:     []string{isLBPoolMembersDeleted, ""},
		Refresh:    isDeleteLBPoolMembersRefreshFunc(lbc, lbID, lbPoolID, lbPoolMemID),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDeleteLBPoolMembersRefreshFunc(lbc *vpcv1.VpcV1, lbID, lbPoolID, lbPoolMemID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		getlbpmoptions := &vpcv1.GetLoadBalancerPoolMembersOptions{
			LoadBalancerID: &lbID,
			PoolID:         &lbPoolID,
			ID:             &lbPoolMemID,
		}
		lbPoolMem, response, err := lbc.GetLoadBalancerPoolMembers(getlbpmoptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return lbPoolMem, isLBPoolMembersDeleted, nil
			}
			return nil, "", fmt.Errorf("[ERROR] Error Deleting Load balancer pool member: %s\n%s", err, response)
		}
		return lbPoolMem, isLBPoolMembersDeletePending, nil
	}
}

func resourceIBMISLBPoolMembersExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return false, err
	}
	if len(parts) != 3 {
		return false, fmt.Errorf(
			"The id should contain loadbalancer Id, loadbalancer pool Id and loadbalancer poolmemebr Id")
	}

	lbID := parts[0]
	lbPoolID := parts[1]
	lbPoolMemID := parts[2]

	exists, err := lbpmemberExists(d, meta, lbID, lbPoolID, lbPoolMemID)
	return exists, err

}

func lbpmemberExists(d *schema.ResourceData, meta interface{}, lbID, lbPoolID, lbPoolMemID string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}

	getlbpmoptions := &vpcv1.GetLoadBalancerPoolMembersOptions{
		LoadBalancerID: &lbID,
		PoolID:         &lbPoolID,
		ID:             &lbPoolMemID,
	}
	_, response, err := sess.GetLoadBalancerPoolMembers(getlbpmoptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("[ERROR] Error getting Load balancer pool member: %s\n%s", err, response)
	}
	return true, nil
}

func getPoolId(id string) (string, error) {
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
