package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/lbaas"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	isLBPoolID                       = "pool"
	isLBPoolMemberPort               = "port"
	isLBPoolMemberTargetAddress      = "target_address"
	isLBPoolMemberWeight             = "weight"
	isLBPoolMemberProvisioningStatus = "provisioning_status"
	isLBPoolMemberHealth             = "health"
	isLBPoolMemberHref               = "href"
	isLBPoolMemberDeletePending      = "delete_pending"
	isLBPoolMemberDeleted            = "done"
	isLBPoolMemberActive             = "active"
)

func resourceIBMISLBPoolMember() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISLBPoolMemberCreate,
		Read:     resourceIBMISLBPoolMemberRead,
		Update:   resourceIBMISLBPoolMemberUpdate,
		Delete:   resourceIBMISLBPoolMemberDelete,
		Exists:   resourceIBMISLBPoolMemberExists,
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
			},

			isLBID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isLBPoolMemberPort: {
				Type:     schema.TypeInt,
				Required: true,
			},

			isLBPoolMemberTargetAddress: {
				Type:     schema.TypeString,
				Required: true,
			},

			isLBPoolMemberWeight: {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			isLBPoolMemberProvisioningStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isLBPoolMemberHealth: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isLBPoolMemberHref: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMISLBPoolMemberCreate(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()

	log.Printf("[DEBUG] LB Pool create")
	lbPoolID := d.Get(isLBPoolID).(string)
	lbID := d.Get(isLBID).(string)
	port := d.Get(isLBPoolMemberPort).(int)
	targetAddress := d.Get(isLBPoolMemberTargetAddress).(string)

	var weight int
	if w, ok := d.GetOk(isLBPoolMemberWeight); ok {
		weight = w.(int)
	}

	client := lbaas.NewLoadBalancerClient(sess)

	_, err := isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	lbPoolMember, err := client.CreatePoolMember(lbID, lbPoolID, targetAddress, port, weight)
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s/%s", lbID, lbPoolID, lbPoolMember.ID.String()))
	log.Printf("[INFO] lbpool member : %s", lbPoolMember.ID.String())
	_, err = isWaitForLBPoolMemberAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer (%s) is active: %s", lbID, err)
	}
	return resourceIBMISLBPoolMemberRead(d, meta)
}

func resourceIBMISLBPoolMemberRead(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()
	client := lbaas.NewLoadBalancerClient(sess)

	parts, err := idParts(d.Id())
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

	lbPoolMem, err := client.GetPoolMember(lbID, lbPoolID, lbPoolMemID)
	if err != nil {
		return err
	}

	d.Set(isLBPoolID, lbPoolID)
	d.Set(isLBID, lbID)
	d.Set(isLBPoolMemberPort, lbPoolMem.Port)
	d.Set(isLBPoolMemberTargetAddress, lbPoolMem.Target)
	d.Set(isLBPoolMemberWeight, lbPoolMem.Weight)
	d.Set(isLBPoolMemberProvisioningStatus, lbPoolMem.ProvisioningStatus)
	d.Set(isLBPoolMemberHealth, lbPoolMem.Health)
	d.Set(isLBPoolMemberHref, lbPoolMem.Href)

	return nil
}

func resourceIBMISLBPoolMemberUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).ISSession()
	client := lbaas.NewLoadBalancerClient(sess)

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	lbPoolID := parts[1]
	lbPoolMemID := parts[2]

	if d.HasChange(isLBPoolMemberTargetAddress) || d.HasChange(isLBPoolMemberPort) || d.HasChange(isLBPoolMemberWeight) {

		_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error checking for load balancer (%s) is active: %s", lbID, err)
		}
		_, err = client.UpdatePoolMember(lbID, lbPoolID, lbPoolMemID, d.Get(isLBPoolMemberTargetAddress).(string), d.Get(isLBPoolMemberPort).(int), d.Get(isLBPoolMemberWeight).(int))
		if err != nil {
			return err
		}
		_, err = isWaitForLBPoolMemberAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return err
		}
		_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error checking for load balancer (%s) is active: %s", lbID, err)
		}
	}

	return resourceIBMISLBPoolMemberRead(d, meta)
}

func resourceIBMISLBPoolMemberDelete(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).ISSession()
	client := lbaas.NewLoadBalancerClient(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	lbPoolID := parts[1]
	lbPoolMemID := parts[2]

	_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	err = client.DeletePoolMember(lbID, lbPoolID, lbPoolMemID)
	if err != nil {
		return err
	}
	_, err = isWaitForLBPoolMemberDeleted(client, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	d.SetId("")
	return nil
}

func resourceIBMISLBPoolMemberExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, _ := meta.(ClientSession).ISSession()
	client := lbaas.NewLoadBalancerClient(sess)

	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	if len(parts) < 3 {
		return false, fmt.Errorf(
			"The id should contain loadbalancer Id, loadbalancer pool Id and loadbalancer poolmemebr Id")
	}

	lbID := parts[0]
	lbPoolID := parts[1]
	lbPoolMemID := parts[2]

	_, err = client.GetPoolMember(lbID, lbPoolID, lbPoolMemID)

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

func isWaitForLBPoolMemberAvailable(client *lbaas.LoadBalancerClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for load balancer pool member(%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"create_pending", "update_pending", "maintenance_pending"},
		Target:     []string{isLBPoolMemberActive},
		Refresh:    isLBPoolMemberRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLBPoolMemberRefreshFunc(client *lbaas.LoadBalancerClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		parts, err := idParts(id)
		if err != nil {
			return nil, "", err
		}

		lbID := parts[0]
		lbPoolID := parts[1]
		lbPoolMemID := parts[2]

		lbPoolMem, err := client.GetPoolMember(lbID, lbPoolID, lbPoolMemID)
		if err != nil {
			return nil, "", err
		}

		if lbPoolMem.ProvisioningStatus == isLBPoolMemberActive {
			return lbPoolMem, lbPoolMem.ProvisioningStatus, nil
		}

		return lbPoolMem, lbPoolMem.ProvisioningStatus, nil
	}
}

func isWaitForLBPoolMemberDeleted(lbc *lbaas.LoadBalancerClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBPoolMemberDeletePending},
		Target:     []string{},
		Refresh:    isDeleteLBPoolMemberRefreshFunc(lbc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDeleteLBPoolMemberRefreshFunc(lbc *lbaas.LoadBalancerClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		parts, err := idParts(id)
		if err != nil {
			return nil, "", err
		}

		lbID := parts[0]
		lbPoolID := parts[1]
		lbPoolMemID := parts[2]

		lbPoolMem, err := lbc.GetPoolMember(lbID, lbPoolID, lbPoolMemID)
		if err == nil {
			return lbPoolMem, isLBPoolMemberDeletePending, nil
		}
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "member_not_found" {
				return nil, isLBPoolMemberDeleted, nil
			}
		}
		return nil, isLBPoolMemberDeletePending, err
	}
}
