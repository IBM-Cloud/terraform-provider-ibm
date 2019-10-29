package ibm

import (
	"fmt"
	"time"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMContainerVpcALB() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerVpcALBCreate,
		Read:     resourceIBMContainerVpcALBRead,
		Update:   resourceIBMContainerVpcALBUpdate,
		Delete:   resourceIBMContainerVpcALBDelete,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"alb_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"alb_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable": {
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"disable_deployment"},
			},
			"disable_deployment": {
				Type:          schema.TypeBool,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"enable"},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"load_balancer_hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resize": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMContainerVpcALBCreate(d *schema.ResourceData, meta interface{}) error {
	albClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	var enable, disableDeployment bool
	albID := d.Get("alb_id").(string)
	if v, ok := d.GetOkExists("enable"); ok {
		enable = v.(bool)
	} else if v, ok := d.GetOkExists("disable_deployment"); ok {
		disableDeployment = v.(bool)
	} else {
		return fmt.Errorf("Provide either `enable` or `disable_deployment`")
	}

	params := v2.AlbConfig{
		AlbID:  albID,
		Enable: enable,
	}

	albAPI := albClient.Albs()
	targetEnv := v2.ClusterTargetHeader{}
	if err != nil {
		return err
	}
	if enable {
		err = albAPI.EnableAlb(params, targetEnv)
		if err != nil {
			return err
		}
	} else {
		err = albAPI.DisableAlb(params, targetEnv)
		if err != nil {
			return err
		}
	}

	d.SetId(albID)
	_, err = waitForVpcContainerALB(d, meta, albID, schema.TimeoutCreate, enable, disableDeployment)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for create resource alb (%s) : %s", d.Id(), err)
	}

	return resourceIBMContainerVpcALBRead(d, meta)
}

func resourceIBMContainerVpcALBRead(d *schema.ResourceData, meta interface{}) error {
	albClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	albID := d.Id()

	albAPI := albClient.Albs()
	targetEnv := v2.ClusterTargetHeader{}

	albConfig, err := albAPI.GetAlb(albID, targetEnv)
	if err != nil {
		return err
	}

	d.Set("alb_type", albConfig.AlbType)
	d.Set("cluster", albConfig.Cluster)
	d.Set("name", albConfig.Name)
	d.Set("enable", albConfig.Enable)
	d.Set("disable_deployment", albConfig.DisableDeployment)
	d.Set("enable", albConfig.Enable)
	d.Set("resize", albConfig.Resize)
	d.Set("zone", albConfig.ZoneAlb)
	d.Set("status", albConfig.Status)
	d.Set("state", albConfig.State)
	d.Set("load_balancer_hostname", albConfig.LoadBalancerHostname)

	return nil
}

func resourceIBMContainerVpcALBUpdate(d *schema.ResourceData, meta interface{}) error {
	albClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	albAPI := albClient.Albs()

	if d.HasChange("enable") {
		enable := d.Get("enable").(bool)
		disableDeployment := d.Get("disable_deployment").(bool)
		albID := d.Id()
		params := v2.AlbConfig{
			AlbID:  albID,
			Enable: enable,
		}

		targetEnv := v2.ClusterTargetHeader{}

		if enable {
			err = albAPI.EnableAlb(params, targetEnv)
			if err != nil {
				return err
			}
		} else {
			err = albAPI.DisableAlb(params, targetEnv)
			if err != nil {
				return err
			}
		}

		_, err = waitForVpcContainerALB(d, meta, albID, schema.TimeoutUpdate, enable, disableDeployment)
		if err != nil {
			return fmt.Errorf(
				"Error waiting for updating resource alb (%s) : %s", d.Id(), err)
		}

	}
	return resourceIBMContainerVpcALBRead(d, meta)
}

func waitForVpcContainerALB(d *schema.ResourceData, meta interface{}, albID, timeout string, enable, disableDeployment bool) (interface{}, error) {
	albClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return false, err
	}

	stateConf := &resource.StateChangeConf{
		Pending: []string{"pending"},
		Target:  []string{"active"},
		Refresh: func() (interface{}, string, error) {
			targetEnv := v2.ClusterTargetHeader{}

			alb, err := albClient.Albs().GetAlb(albID, targetEnv)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return nil, "", fmt.Errorf("The resource alb %s does not exist anymore: %v", d.Id(), err)
				}
				return nil, "", err
			}
			if enable {
				if alb.Enable == false {
					return alb, "pending", nil
				}
			} else if disableDeployment {
				if alb.Enable == true {
					return alb, "pending", nil
				}
			}
			return alb, "active", nil
		},
		Timeout:    d.Timeout(timeout),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func resourceIBMContainerVpcALBDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")

	return nil
}
