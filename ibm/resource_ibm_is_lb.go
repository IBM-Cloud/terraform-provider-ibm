package ibm

import (
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
	isLBName             = "name"
	isLBStatus           = "status"
	isLBType             = "type"
	isLBSubnets          = "subnets"
	isLBHostName         = "hostname"
	isLBPublicIPs        = "public_ips"
	isLBPrivateIPs       = "private_ips"
	isLBOperatingStatus  = "operating_status"
	isLBDeleting         = "deleting"
	isLBDeleted          = "done"
	isLBProvisioning     = "provisioning"
	isLBProvisioningDone = "done"
	isLBResourceGroup    = "resource_group"
)

func resourceIBMISLB() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISLBCreate,
		Read:     resourceIBMISLBRead,
		Update:   resourceIBMISLBUpdate,
		Delete:   resourceIBMISLBDelete,
		Exists:   resourceIBMISLBExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isLBName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			isLBType: {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				Default:      "public",
				ValidateFunc: validateAllowedStringValue([]string{"public", "private"}),
			},

			isLBStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isLBOperatingStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isLBPublicIPs: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			isLBPrivateIPs: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			isLBSubnets: {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			isVPNGatewayResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},

			isLBHostName: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMISLBCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	name := d.Get(isLBName).(string)
	lbType := d.Get(isLBType).(string)
	subnets := expandStringList((d.Get(isLBSubnets).(*schema.Set)).List())
	var isPublic = true
	if lbType == "private" {
		isPublic = false
	}

	client := lbaas.NewLoadBalancerClient(sess)

	var subnetIdentity []*models.SubnetIdentity
	for _, subnet := range subnets {
		subnetIdentity = append(subnetIdentity, &models.SubnetIdentity{ID: strfmt.UUID(subnet)})
	}
	body := &models.LoadBalancerTemplate{
		IsPublic: &isPublic,
		Name:     name,
		Subnets:  subnetIdentity,
	}

	if rg, ok := d.GetOk(isLBResourceGroup); ok {
		rgref := models.LoadBalancerTemplateResourceGroup{
			ID: strfmt.UUID(rg.(string)),
		}
		body.ResourceGroup = &rgref
	}

	lb, err := client.Create(&l_baas.PostLoadBalancersParams{
		Body: body,
	})
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	d.SetId(lb.ID.String())
	log.Printf("[INFO]  : %s", lb.ID.String())

	_, err = isWaitForLBAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	return resourceIBMISLBRead(d, meta)
}

func resourceIBMISLBRead(d *schema.ResourceData, meta interface{}) error {
	sess, _ := meta.(ClientSession).ISSession()
	client := lbaas.NewLoadBalancerClient(sess)

	lb, err := client.Get(d.Id())
	if err != nil {
		return err
	}

	d.Set("id", lb.ID.String())
	d.Set(isLBName, lb.Name)
	if lb.IsPublic {
		d.Set(isLBType, "public")
	} else {
		d.Set(isLBType, "private")
	}

	d.Set(isLBStatus, lb.ProvisioningStatus)
	d.Set(isLBOperatingStatus, lb.OperatingStatus)
	d.Set(isLBPublicIPs, flattenISLBIPs(lb.PublicIps))
	d.Set(isLBPrivateIPs, flattenISLBIPs(lb.PrivateIps))
	d.Set(isLBSubnets, flattenISLBSubnets(lb.Subnets))
	d.Set(isLBResourceGroup, lb.ResourceGroup.ID)
	d.Set(isLBHostName, lb.Hostname)

	return nil
}

func resourceIBMISLBUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).ISSession()
	client := lbaas.NewLoadBalancerClient(sess)

	if d.HasChange(isLBName) {
		name := d.Get(isLBName).(string)
		_, err := client.Update(d.Id(), name)
		if err != nil {
			return err
		}
	}

	return resourceIBMISLBRead(d, meta)
}

func resourceIBMISLBDelete(d *schema.ResourceData, meta interface{}) error {

	sess, _ := meta.(ClientSession).ISSession()
	client := lbaas.NewLoadBalancerClient(sess)
	err := client.Delete(d.Id())
	if err != nil {
		return err
	}

	_, err = isWaitForDeleted(client, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func isWaitForDeleted(lbc *lbaas.LoadBalancerClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isLBDeleting},
		Target:     []string{},
		Refresh:    isDeleteRefreshFunc(lbc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDeleteRefreshFunc(lbc *lbaas.LoadBalancerClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		lb, err := lbc.Get(id)
		if err == nil {
			return lb, isLBDeleting, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "load_balancer_not_found" {
				return nil, isLBDeleted, nil
			}
		}
		return nil, isLBDeleting, err
	}
}

func resourceIBMISLBExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, _ := meta.(ClientSession).ISSession()
	client := lbaas.NewLoadBalancerClient(sess)

	_, err := client.Get(d.Id())
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

func isWaitForLBAvailable(client *lbaas.LoadBalancerClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for load balancer (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isLBProvisioning},
		Target:     []string{isLBProvisioningDone},
		Refresh:    isLBRefreshFunc(client, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isLBRefreshFunc(client *lbaas.LoadBalancerClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		lb, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if lb.ProvisioningStatus == "active" || lb.ProvisioningStatus == "failed" {
			return lb, isLBProvisioningDone, nil
		}

		return lb, isLBProvisioning, nil
	}
}
