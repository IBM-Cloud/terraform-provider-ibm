package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/lbaas"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/riaas/rias-api/riaas/client/l_baas"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

const (
	isLBListenerLBID                = "lb"
	isLBListenerPort                = "port"
	isLBListenerProtocol            = "protocol"
	isLBListenerCertificateInstance = "certificate_instance"
	isLBListenerConnectionLimit     = "connection_limit"
	isLBListenerDefaultPool         = "default_pool"
	isLBListenerStatus              = "status"
	isLBListenerDeleting            = "deleting"
	isLBListenerDeleted             = "done"
	isLBListenerProvisioning        = "provisioning"
	isLBListenerProvisioningDone    = "done"
)

func resourceIBMISLBListener() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISLBListenerCreate,
		Read:     resourceIBMISLBListenerRead,
		Update:   resourceIBMISLBListenerUpdate,
		Delete:   resourceIBMISLBListenerDelete,
		Exists:   resourceIBMISLBListenerExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isLBListenerLBID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isLBListenerPort: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateLBListenerPort,
			},

			isLBListenerProtocol: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"https", "http", "tcp"}),
			},

			isLBListenerCertificateInstance: {
				Type:     schema.TypeString,
				Optional: true,
			},

			isLBListenerConnectionLimit: {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateLBListenerConnectionLimit,
			},

			isLBListenerDefaultPool: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			isLBListenerStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMISLBListenerCreate(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()

	log.Printf("[DEBUG] LB Listener create")
	lbID := d.Get(isLBListenerLBID).(string)
	port := int64(d.Get(isLBListenerPort).(int))
	protocol := d.Get(isLBListenerProtocol).(string)

	var defPool, certificateCRN string
	if pool, ok := d.GetOk(isLBListenerDefaultPool); ok {
		defPool = pool.(string)
	}

	if crn, ok := d.GetOk(isLBListenerCertificateInstance); ok {
		certificateCRN = crn.(string)
	}

	var connLimit int64

	if limit, ok := d.GetOk(isLBListenerConnectionLimit); ok {
		connLimit = limit.(int64)
	}

	client := lbaas.NewLoadBalancerClient(sess)
	body := &models.ListenerTemplate{
		Port:     &port,
		Protocol: &protocol,
	}

	if certificateCRN != "" {
		body.CertificateInstance = &models.ListenerTemplateCertificateInstance{
			Crn: certificateCRN,
		}
	}
	if defPool != "" {
		body.DefaultPool = &models.ListenerTemplateDefaultPool{
			ID: strfmt.UUID(defPool),
		}
	}

	if connLimit > 0 {
		body.ConnectionLimit = connLimit
	}

	_, err := isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer (%s) is active: %s", lbID, err)
	}
	lbListener, err := client.CreateListeners(&l_baas.PostLoadBalancersIDListenersParams{
		Body: body,
		ID:   lbID,
	})
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", lbID, lbListener.ID.String()))

	_, err = isWaitForLBListenerAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf(
			"Error waiting for load balancer listener(%s) to become ready: %s", d.Id(), err)
	}

	_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf(
			"Error waiting for load balancer (%s) to become ready: %s", lbID, err)
	}

	log.Printf("[INFO] Load balancer Listener : %s", d.Id())
	return resourceIBMISLBListenerRead(d, meta)
}

func resourceIBMISLBListenerRead(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()
	client := lbaas.NewLoadBalancerClient(sess)

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	lbListenerID := parts[1]

	lbListener, err := client.GetListener(lbID, lbListenerID)
	if err != nil {
		return err
	}

	d.Set(isLBListenerLBID, lbID)
	d.Set(isLBListenerPort, lbListener.Port)
	d.Set(isLBListenerProtocol, lbListener.Protocol)

	if lbListener.DefaultPool != nil {
		d.Set(isLBListenerDefaultPool, lbListener.DefaultPool.ID)
	}
	if lbListener.CertificateInstance != nil {
		d.Set(isLBListenerCertificateInstance, lbListener.CertificateInstance.Crn)
	}
	d.Set(isLBListenerConnectionLimit, lbListener.ConnectionLimit)
	d.Set(isLBListenerStatus, lbListener.ProvisioningStatus)

	return nil
}

func resourceIBMISLBListenerUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).ISSession()
	client := lbaas.NewLoadBalancerClient(sess)

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	lbListenerID := parts[1]

	hasChanged := false

	var certificateInstance, defPool, protocol string
	var connLimit, port int
	if d.HasChange(isLBListenerCertificateInstance) {
		certificateInstance = d.Get(isLBListenerCertificateInstance).(string)
		hasChanged = true
	}

	if d.HasChange(isLBListenerDefaultPool) {
		defPool = d.Get(isLBListenerDefaultPool).(string)
		hasChanged = true
	}
	if d.HasChange(isLBListenerPort) {
		port = d.Get(isLBListenerPort).(int)
		hasChanged = true
	}

	if d.HasChange(isLBListenerProtocol) {
		protocol = d.Get(isLBListenerProtocol).(string)
		hasChanged = true
	}

	if d.HasChange(isLBListenerConnectionLimit) {
		connLimit = d.Get(isLBListenerConnectionLimit).(int)
		hasChanged = true
	}

	if hasChanged {

		_, err := isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error checking for load balancer (%s) is active: %s", lbID, err)
		}
		_, err = client.UpdateListener(lbID, lbListenerID, certificateInstance, protocol, defPool, port, connLimit)
		if err != nil {
			return err
		}

		_, err = isWaitForLBListenerAvailable(client, d.Id(), d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error waiting for load balancer listener(%s) to become ready: %s", d.Id(), err)
		}

		_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error waiting for load balancer (%s) to become ready: %s", lbID, err)
		}
	}

	return resourceIBMISLBListenerRead(d, meta)
}

func resourceIBMISLBListenerDelete(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).ISSession()
	client := lbaas.NewLoadBalancerClient(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	lbListenerID := parts[1]
	_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer (%s) is active: %s", lbID, err)
	}
	err = client.DeleteListener(lbID, lbListenerID)
	if err != nil {
		return err
	}
	_, err = isWaitForLBListenerDeleted(client, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}
	_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf(
			"Error waiting for load balancer (%s) to be active: %s", lbID, err)
	}

	d.SetId("")
	return nil
}

func isWaitForLBListenerDeleted(lbc *lbaas.LoadBalancerClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isLBListenerDeleting, "delete_pending"},
		Target:     []string{},
		Refresh:    isLBListenerDeleteRefreshFunc(lbc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLBListenerDeleteRefreshFunc(lbc *lbaas.LoadBalancerClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		parts, err := idParts(id)
		if err != nil {
			return nil, isLBListenerDeleting, err
		}

		lbID := parts[0]
		LBListenerID := parts[1]

		lbLis, err := lbc.GetListener(lbID, LBListenerID)
		if err == nil {
			return lbLis, lbLis.ProvisioningStatus, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "listener_not_found" {
				return nil, isLBListenerDeleted, nil
			}
		}
		return nil, isLBListenerDeleting, err
	}
}

func resourceIBMISLBListenerExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, _ := meta.(ClientSession).ISSession()
	client := lbaas.NewLoadBalancerClient(sess)

	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	lbID := parts[0]
	LBListenerID := parts[1]

	_, err = client.GetListener(lbID, LBListenerID)

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

func isWaitForLBListenerAvailable(client *lbaas.LoadBalancerClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for load balancer listener(%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isLBListenerProvisioning, "create_pending", "update_pending", "maintenance_pending"},
		Target:     []string{isLBListenerProvisioningDone},
		Refresh:    isLBListenerRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLBListenerRefreshFunc(client *lbaas.LoadBalancerClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		parts, err := idParts(id)
		if err != nil {
			return nil, "", err
		}

		lbID := parts[0]
		LBListenerID := parts[1]

		lblis, err := client.GetListener(lbID, LBListenerID)
		if err != nil {
			return nil, "", err
		}

		if lblis.ProvisioningStatus == "active" || lblis.ProvisioningStatus == "failed" {
			return lblis, isLBListenerProvisioningDone, nil
		}

		return lblis, lblis.ProvisioningStatus, nil
	}
}
