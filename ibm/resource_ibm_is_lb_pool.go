package ibm

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/lbaas"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/l_baas"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

const (
	isLBPoolName                      = "name"
	isLBID                            = "lb"
	isLBPoolAlgorithm                 = "algorithm"
	isLBPoolProtocol                  = "protocol"
	isLBPoolHealthDelay               = "health_delay"
	isLBPoolHealthRetries             = "health_retries"
	isLBPoolHealthTimeout             = "health_timeout"
	isLBPoolHealthType                = "health_type"
	isLBPoolHealthMonitorURL          = "health_monitor_url"
	isLBPoolHealthMonitorPort         = "health_monitor_port"
	isLBPoolSessPersistenceType       = "session_persistence_type"
	isLBPoolSessPersistenceCookieName = "session_persistence_cookie_name"
	isLBPoolProvisioningStatus        = "provisioning_status"
	isLBPoolActive                    = "active"
	isLBPoolCreatePending             = "create_pending"
	isLBPoolUpdatePending             = "update_pending"
	isLBPoolDeletePending             = "delete_pending"
	isLBPoolMaintainancePending       = "maintenance_pending"
	isLBPoolFailed                    = "failed"
	isLBPoolDeleteDone                = "deleted"
	//isLBPoolActive,isLBPoolCreatePending,isLBPoolUpdatePending,isLBPoolDeletePending,isLBPoolMaintainancePending,isLBPoolFailed
)

func resourceIBMISLBPool() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISLBPoolCreate,
		Read:     resourceIBMISLBPoolRead,
		Update:   resourceIBMISLBPoolUpdate,
		Delete:   resourceIBMISLBPoolDelete,
		Exists:   resourceIBMISLBPoolExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			isLBPoolName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateISName,
				Description:  "Load Balancer Pool name",
			},

			isLBID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Load Balancer ID",
			},

			isLBPoolAlgorithm: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"round_robin", "weighted_round_robin", "least_connections"}),
				Description:  "Load Balancer Pool algorithm",
			},

			isLBPoolProtocol: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"http", "tcp"}),
				Description:  "Load Balancer Protocol",
			},

			isLBPoolHealthDelay: {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Load Blancer health delay time period",
			},

			isLBPoolHealthRetries: {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Load Balancer health retry count",
			},

			isLBPoolHealthTimeout: {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Load Balancer health timeout interval",
			},

			isLBPoolHealthType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"http", "tcp"}),
				Description:  "Load Balancer health type",
			},

			isLBPoolHealthMonitorURL: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			isLBPoolHealthMonitorPort: {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			isLBPoolSessPersistenceType: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"source_ip", "http_cookie", "app_cookie"}),
				Description:  "Load Balancer Pool session persisence type.",
			},

			isLBPoolSessPersistenceCookieName: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Load Balancer Pool session persisence cookie name",
			},

			isLBPoolProvisioningStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMISLBPoolCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] LB Pool create")
	name := d.Get(isLBPoolName).(string)
	lbID := d.Get(isLBID).(string)
	algorithm := d.Get(isLBPoolAlgorithm).(string)
	protocol := d.Get(isLBPoolProtocol).(string)
	healthDelay := d.Get(isLBPoolHealthDelay).(int)
	maxRetries := d.Get(isLBPoolHealthRetries).(int)
	healthTimeOut := d.Get(isLBPoolHealthTimeout).(int)
	healthType := d.Get(isLBPoolHealthType).(string)

	var spType, cName, healthMonitorURL string
	var healthMonitorPort int
	if pt, ok := d.GetOk(isLBPoolSessPersistenceType); ok {
		spType = pt.(string)
	}

	if cn, ok := d.GetOk(isLBPoolSessPersistenceCookieName); ok {
		cName = cn.(string)
	}

	if hmu, ok := d.GetOk(isLBPoolHealthMonitorURL); ok {
		healthMonitorURL = hmu.(string)
	}

	if hmp, ok := d.GetOk(isLBPoolHealthMonitorPort); ok {
		healthMonitorPort = hmp.(int)
	}

	client := lbaas.NewLoadBalancerClient(sess)
	body := &models.PoolTemplate{
		Algorithm: algorithm,
		Name:      name,
		Protocol:  protocol,
		HealthMonitor: &models.HealthMonitorTemplate{
			Delay:      int64(healthDelay),
			MaxRetries: int64(maxRetries),
			Timeout:    int64(healthTimeOut),
			Type:       healthType,
		},
	}

	if healthMonitorURL != "" {
		body.HealthMonitor.URLPath = healthMonitorURL
	}
	if healthMonitorPort > 0 {
		body.HealthMonitor.Port = int64(healthMonitorPort)
	}
	if spType != "" {
		body.SessionPersistence = &models.SessionPersistenceTemplate{
			CookieName: cName,
			Type:       spType,
		}
	}
	isLBPoolKey := "load_balancer_pool_key_" + lbID
	ibmMutexKV.Lock(isLBPoolKey)
	defer ibmMutexKV.Unlock(isLBPoolKey)
	_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	lbPool, err := client.CreatePool(&l_baas.PostLoadBalancersIDPoolsParams{
		Body: body,
		ID:   lbID,
	})
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s", lbID, lbPool.ID.String()))
	log.Printf("[INFO] Ipsec : %s", lbPool.ID.String())

	_, err = isWaitForLBPoolActive(client, lbID, lbPool.ID.String(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer pool (%s) is active: %s", lbPool.ID.String(), err)
	}

	_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	return resourceIBMISLBPoolRead(d, meta)
}

func resourceIBMISLBPoolRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	client := lbaas.NewLoadBalancerClient(sess)

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	lbPoolID := parts[1]

	lbPool, err := client.GetPool(lbID, lbPoolID)
	if err != nil {
		return err
	}

	d.Set(isLBPoolName, lbPool.Name)
	d.Set(isLBID, lbID)
	d.Set(isLBPoolAlgorithm, lbPool.Algorithm)
	d.Set(isLBPoolProtocol, lbPool.Protocol)
	d.Set(isLBPoolHealthDelay, lbPool.HealthMonitor.Delay)
	d.Set(isLBPoolHealthRetries, lbPool.HealthMonitor.MaxRetries)
	d.Set(isLBPoolHealthTimeout, lbPool.HealthMonitor.Timeout)
	d.Set(isLBPoolHealthType, lbPool.HealthMonitor.Type)
	d.Set(isLBPoolHealthMonitorURL, lbPool.HealthMonitor.URLPath)
	d.Set(isLBPoolHealthMonitorPort, lbPool.HealthMonitor.Port)
	if lbPool.SessionPersistence != nil {
		d.Set(isLBPoolSessPersistenceType, lbPool.SessionPersistence.Type)
		d.Set(isLBPoolSessPersistenceCookieName, lbPool.SessionPersistence.CookieName)
	}
	d.Set(isLBPoolProvisioningStatus, lbPool.ProvisioningStatus)

	return nil
}

func resourceIBMISLBPoolUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	client := lbaas.NewLoadBalancerClient(sess)

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	lbPoolID := parts[1]

	hasChanged := false
	var healthMonitorTemplate models.HealthMonitorTemplate
	if d.HasChange(isLBPoolHealthDelay) || d.HasChange(isLBPoolHealthRetries) ||
		d.HasChange(isLBPoolHealthTimeout) || d.HasChange(isLBPoolHealthType) || d.HasChange(isLBPoolHealthMonitorURL) || d.HasChange(isLBPoolHealthMonitorPort) {
		healthMonitorTemplate = models.HealthMonitorTemplate{
			Delay:      int64(d.Get(isLBPoolHealthDelay).(int)),
			MaxRetries: int64(d.Get(isLBPoolHealthRetries).(int)),
			Timeout:    int64(d.Get(isLBPoolHealthTimeout).(int)),
			Type:       d.Get(isLBPoolHealthType).(string),
			URLPath:    d.Get(isLBPoolHealthMonitorURL).(string),
		}
		port := d.Get(isLBPoolHealthMonitorPort).(int)
		if port > 0 {
			healthMonitorTemplate.Port = int64(port)
		}
		hasChanged = true

	}
	var sessionPersistence models.SessionPersistenceTemplate
	if d.HasChange(isLBPoolSessPersistenceType) || d.HasChange(isLBPoolSessPersistenceCookieName) {
		sessionPersistence = models.SessionPersistenceTemplate{
			Type:       d.Get(isLBPoolSessPersistenceType).(string),
			CookieName: d.Get(isLBPoolSessPersistenceCookieName).(string),
		}
		hasChanged = true
	}
	if d.HasChange(isLBPoolName) || d.HasChange(isLBPoolAlgorithm) || d.HasChange(isLBPoolProtocol) || hasChanged {
		name := d.Get(isLBPoolName).(string)
		algorithm := d.Get(isLBPoolAlgorithm).(string)
		protocol := d.Get(isLBPoolProtocol).(string)

		isLBPoolKey := "load_balancer_pool_key_" + lbID + lbPoolID
		ibmMutexKV.Lock(isLBPoolKey)
		defer ibmMutexKV.Unlock(isLBPoolKey)
		_, err := isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error checking for load balancer (%s) is active: %s", lbID, err)
		}

		_, err = isWaitForLBPoolActive(client, lbID, lbPoolID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error checking for load balancer pool (%s) is active: %s", lbPoolID, err)
		}

		_, err = client.UpdatePool(lbID, lbPoolID, algorithm, name, protocol, healthMonitorTemplate, sessionPersistence)
		if err != nil {
			return err
		}

		_, err = isWaitForLBPoolActive(client, lbID, lbPoolID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error checking for load balancer pool (%s) is active: %s", lbPoolID, err)
		}

		_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf(
				"Error checking for load balancer (%s) is active: %s", lbID, err)
		}
	}

	return resourceIBMISLBPoolRead(d, meta)
}

func resourceIBMISLBPoolDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	client := lbaas.NewLoadBalancerClient(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	lbID := parts[0]
	lbPoolID := parts[1]

	isLBPoolKey := "load_balancer_pool_key_" + lbID + lbPoolID
	ibmMutexKV.Lock(isLBPoolKey)
	defer ibmMutexKV.Unlock(isLBPoolKey)
	_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	_, err = isWaitForLBPoolActive(client, lbID, lbPoolID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer pool (%s) is active: %s", lbPoolID, err)
	}

	err = client.DeletePool(lbID, lbPoolID)
	if err != nil {
		return err
	}

	_, err = isWaitForLBPoolDeleted(client, lbID, lbPoolID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer pool (%s) is deleted: %s", lbPoolID, err)
	}

	_, err = isWaitForLBAvailable(client, lbID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf(
			"Error checking for load balancer (%s) is active: %s", lbID, err)
	}

	d.SetId("")
	return nil
}

func resourceIBMISLBPoolExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	client := lbaas.NewLoadBalancerClient(sess)

	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	lbID := parts[0]
	lbPoolID := parts[1]

	_, err = client.GetPool(lbID, lbPoolID)

	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "pool_not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}

func isWaitForLBPoolActive(client *lbaas.LoadBalancerClient, lbId, lbPoolId string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for load balancer pool (%s) to be available.", lbPoolId)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBPoolCreatePending, isLBPoolUpdatePending, isLBPoolMaintainancePending},
		Target:     []string{isLBPoolActive},
		Refresh:    isLBPoolRefreshFunc(client, lbId, lbPoolId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLBPoolRefreshFunc(client *lbaas.LoadBalancerClient, lbId, lbPoolId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		lbPool, err := client.GetPool(lbId, lbPoolId)
		if err != nil {
			return nil, "", err
		}

		if lbPool.ProvisioningStatus == isLBPoolActive || lbPool.ProvisioningStatus == isLBPoolFailed {
			return lbPool, isLBPoolActive, nil
		}

		return lbPool, lbPool.ProvisioningStatus, nil
	}
}

func isWaitForLBPoolDeleted(client *lbaas.LoadBalancerClient, lbId, lbPoolId string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for load balancer pool(%s) to be deleted.", lbPoolId)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{isLBPoolUpdatePending, isLBPoolMaintainancePending, isLBPoolDeletePending},
		Target:     []string{},
		Refresh:    isLBPoolDeleteRefreshFunc(client, lbId, lbPoolId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLBPoolDeleteRefreshFunc(client *lbaas.LoadBalancerClient, lbId, lbPoolId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		lbPool, err := client.GetPool(lbId, lbPoolId)
		if err == nil {
			return lbPool, lbPool.ProvisioningStatus, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s %d %s", iserror.Error(), len(iserror.Payload.Errors), iserror.Payload.Errors[0].Code)
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "pool_not_found" {
				return nil, isLBPoolDeleteDone, nil
			}
		}
		return nil, lbPool.ProvisioningStatus, err
	}
}
